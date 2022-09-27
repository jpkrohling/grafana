package grpcserver

import (
	"context"
	"fmt"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/genproto/pluginv2"
	"github.com/grafana/grafana/pkg/components/simplejson"
	"github.com/grafana/grafana/pkg/expr"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/plugins"
	"github.com/grafana/grafana/pkg/plugins/adapters"
	"github.com/grafana/grafana/pkg/services/datasources"
	"github.com/grafana/grafana/pkg/services/user"
	"github.com/grafana/grafana/pkg/setting"
	"github.com/grafana/grafana/pkg/tsdb/grafanads"
)

type DatasourceService struct {
	cfg    *setting.Cfg
	server *DatasourceServer
}

func ProvideDatasourceService(cfg *setting.Cfg, grpcServerProvider Provider, pluginsClient plugins.Client, dataSourceCache datasources.CacheService, dataSourceService datasources.DataSourceService) (*DatasourceService, error) {
	server := &DatasourceServer{
		logger:            log.New("grpc-server-datasource"),
		pluginClient:      pluginsClient,
		dataSourceCache:   dataSourceCache,
		dataSourceService: dataSourceService,
	}
	pluginv2.RegisterDataServer(grpcServerProvider.GetServer(), server)
	return &DatasourceService{
		cfg:    cfg,
		server: server,
	}, nil
}

type DatasourceServer struct {
	logger            log.Logger
	pluginClient      plugins.Client
	dataSourceCache   datasources.CacheService
	dataSourceService datasources.DataSourceService
}

// QueryData handles multiple queries and returns multiple responses.
func (s *DatasourceServer) QueryData(ctx context.Context, in *pluginv2.QueryDataRequest) (*pluginv2.QueryDataResponse, error) {
	req := backend.FromProto().QueryDataRequest(in)
	user := adapters.SignedInUserFromBackendUser(req.PluginContext.User)
	user.OrgID = req.PluginContext.OrgID

	requests, err := s.buildRequests(ctx, user, req)
	if err != nil {
		return nil, err
	}

	response := backend.NewQueryDataResponse()

	for _, req := range requests {
		res, err := s.pluginClient.QueryData(ctx, req)
		if err != nil {
			return nil, err
		}
		for refID, res := range res.Responses {
			response.Responses[refID] = res
		}
	}

	return backend.ToProto().QueryDataResponse(response)
}

// TODO: can this be shared with the query service?
func (s *DatasourceServer) getDataSourceFromQuery(ctx context.Context, signedInUser *user.SignedInUser, skipCache bool, query *simplejson.Json) (*datasources.DataSource, error) {
	uid := query.Get("datasource").Get("uid").MustString()

	// before 8.3 special types could be sent as datasource (expr)
	if uid == "" {
		uid = query.Get("datasource").MustString()
	}

	if expr.IsDataSource(uid) {
		return expr.DataSourceModel(), nil
	}

	if uid == grafanads.DatasourceUID {
		return grafanads.DataSourceModel(signedInUser.OrgID), nil
	}

	// use datasourceId if it exists
	id := query.Get("datasourceId").MustInt64(0)
	if id > 0 {
		ds, err := s.dataSourceCache.GetDatasource(ctx, id, signedInUser, skipCache)
		if err != nil {
			return nil, err
		}
		return ds, nil
	}

	if uid != "" {
		ds, err := s.dataSourceCache.GetDatasourceByUID(ctx, uid, signedInUser, skipCache)
		if err != nil {
			return nil, err
		}
		return ds, nil
	}

	return nil, fmt.Errorf("missing data source ID/UID")
}

func (s *DatasourceServer) buildRequests(ctx context.Context, u *user.SignedInUser, req *backend.QueryDataRequest) ([]*backend.QueryDataRequest, error) {
	requests := make(map[string]*backend.QueryDataRequest)
	for _, query := range req.Queries {
		sj, err := simplejson.NewJson(query.JSON)
		if err != nil {
			return nil, err
		}

		ds, err := s.getDataSourceFromQuery(ctx, u, false, sj)
		if err != nil {
			return nil, err
		}

		r, ok := requests[ds.Uid]
		if !ok {
			instanceSettings, err := adapters.ModelToInstanceSettings(ds, s.decryptSecureJsonDataFn(ctx))
			if err != nil {
				return nil, err
			}
			r = &backend.QueryDataRequest{
				PluginContext: req.PluginContext,
				Headers:       req.Headers,
				Queries:       make([]backend.DataQuery, 0),
			}
			r.PluginContext.PluginID = ds.Type
			r.PluginContext.DataSourceInstanceSettings = instanceSettings
			requests[ds.Uid] = r
		}

		requests[ds.Uid].Queries = append(r.Queries, query)
	}

	reqs := make([]*backend.QueryDataRequest, 0)
	for _, req := range requests {
		reqs = append(reqs, req)
	}

	return reqs, nil
}

// TODO: can this be shared with the query service?
func (s *DatasourceServer) decryptSecureJsonDataFn(ctx context.Context) func(ds *datasources.DataSource) (map[string]string, error) {
	return func(ds *datasources.DataSource) (map[string]string, error) {
		return s.dataSourceService.DecryptedValues(ctx, ds)
	}
}
