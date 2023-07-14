import * as React from 'react';

import { Page } from 'app/core/components/Page/Page';
import { DataSourceAddButton } from 'app/features/datasources/components/DataSourceAddButton';
import { getDataSourcesCount } from 'app/features/datasources/state';
import { StoreState, useSelector } from 'app/types';

import { DataSourcesList } from '../components/DataSourcesList';

export function DataSourcesListPage() {
  const dataSourcesCount = useSelector(({ dataSources }: StoreState) => getDataSourcesCount(dataSources));

  const actions = dataSourcesCount > 0 ? <DataSourceAddButton /> : undefined;
  return (
    <Page navId={'connections-datasources'} actions={actions}>
      <Page.Contents>
        <DataSourcesList />
      </Page.Contents>
    </Page>
  );
}
