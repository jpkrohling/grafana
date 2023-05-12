// Code generated by mockery v2.10.0. DO NOT EDIT.

package eval

import (
	context "context"

	backend "github.com/grafana/grafana-plugin-sdk-go/backend"

	mock "github.com/stretchr/testify/mock"

	models "github.com/grafana/grafana/pkg/services/ngalert/models"

	time "time"

	user "github.com/grafana/grafana/pkg/services/user"
)

// FakeEvaluator is an autogenerated mock type for the Evaluator type
type FakeEvaluator struct {
	mock.Mock
}

type FakeEvaluator_Expecter struct {
	mock *mock.Mock
}

func (_m *FakeEvaluator) EXPECT() *FakeEvaluator_Expecter {
	return &FakeEvaluator_Expecter{mock: &_m.Mock}
}

// ConditionEval provides a mock function with given fields: ctx, _a1, condition, now
func (_m *FakeEvaluator) ConditionEval(ctx context.Context, _a1 *user.SignedInUser, condition models.Condition, now time.Time) Results {
	ret := _m.Called(ctx, _a1, condition, now)

	var r0 Results
	if rf, ok := ret.Get(0).(func(context.Context, *user.SignedInUser, models.Condition, time.Time) Results); ok {
		r0 = rf(ctx, _a1, condition, now)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Results)
		}
	}

	return r0
}

// FakeEvaluator_ConditionEval_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConditionEval'
type FakeEvaluator_ConditionEval_Call struct {
	*mock.Call
}

// ConditionEval is a helper method to define mock.On call
//  - ctx context.Context
//  - _a1 *user.SignedInUser
//  - condition models.Condition
//  - now time.Time
func (_e *FakeEvaluator_Expecter) ConditionEval(ctx interface{}, _a1 interface{}, condition interface{}, now interface{}) *FakeEvaluator_ConditionEval_Call {
	return &FakeEvaluator_ConditionEval_Call{Call: _e.mock.On("ConditionEval", ctx, _a1, condition, now)}
}

func (_c *FakeEvaluator_ConditionEval_Call) Run(run func(ctx context.Context, _a1 *user.SignedInUser, condition models.Condition, now time.Time)) *FakeEvaluator_ConditionEval_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*user.SignedInUser), args[2].(models.Condition), args[3].(time.Time))
	})
	return _c
}

func (_c *FakeEvaluator_ConditionEval_Call) Return(_a0 Results) *FakeEvaluator_ConditionEval_Call {
	_c.Call.Return(_a0)
	return _c
}

// QueriesAndExpressionsEval provides a mock function with given fields: ctx, _a1, data, now
func (_m *FakeEvaluator) QueriesAndExpressionsEval(ctx context.Context, _a1 *user.SignedInUser, data []models.AlertQuery, now time.Time) (*backend.QueryDataResponse, error) {
	ret := _m.Called(ctx, _a1, data, now)

	var r0 *backend.QueryDataResponse
	if rf, ok := ret.Get(0).(func(context.Context, *user.SignedInUser, []models.AlertQuery, time.Time) *backend.QueryDataResponse); ok {
		r0 = rf(ctx, _a1, data, now)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backend.QueryDataResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.SignedInUser, []models.AlertQuery, time.Time) error); ok {
		r1 = rf(ctx, _a1, data, now)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FakeEvaluator_QueriesAndExpressionsEval_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueriesAndExpressionsEval'
type FakeEvaluator_QueriesAndExpressionsEval_Call struct {
	*mock.Call
}

// QueriesAndExpressionsEval is a helper method to define mock.On call
//  - ctx context.Context
//  - _a1 *user.SignedInUser
//  - data []models.AlertQuery
//  - now time.Time
func (_e *FakeEvaluator_Expecter) QueriesAndExpressionsEval(ctx interface{}, _a1 interface{}, data interface{}, now interface{}) *FakeEvaluator_QueriesAndExpressionsEval_Call {
	return &FakeEvaluator_QueriesAndExpressionsEval_Call{Call: _e.mock.On("QueriesAndExpressionsEval", ctx, _a1, data, now)}
}

func (_c *FakeEvaluator_QueriesAndExpressionsEval_Call) Run(run func(ctx context.Context, _a1 *user.SignedInUser, data []models.AlertQuery, now time.Time)) *FakeEvaluator_QueriesAndExpressionsEval_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*user.SignedInUser), args[2].([]models.AlertQuery), args[3].(time.Time))
	})
	return _c
}

func (_c *FakeEvaluator_QueriesAndExpressionsEval_Call) Return(_a0 *backend.QueryDataResponse, _a1 error) *FakeEvaluator_QueriesAndExpressionsEval_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Validate provides a mock function with given fields: ctx, _a1, condition
func (_m *FakeEvaluator) Validate(ctx context.Context, _a1 *user.SignedInUser, condition models.Condition) error {
	ret := _m.Called(ctx, _a1, condition)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.SignedInUser, models.Condition) error); ok {
		r0 = rf(ctx, _a1, condition)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FakeEvaluator_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type FakeEvaluator_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//  - ctx context.Context
//  - _a1 *user.SignedInUser
//  - condition models.Condition
func (_e *FakeEvaluator_Expecter) Validate(ctx interface{}, _a1 interface{}, condition interface{}) *FakeEvaluator_Validate_Call {
	return &FakeEvaluator_Validate_Call{Call: _e.mock.On("Validate", ctx, _a1, condition)}
}

func (_c *FakeEvaluator_Validate_Call) Run(run func(ctx context.Context, _a1 *user.SignedInUser, condition models.Condition)) *FakeEvaluator_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*user.SignedInUser), args[2].(models.Condition))
	})
	return _c
}

func (_c *FakeEvaluator_Validate_Call) Return(_a0 error) *FakeEvaluator_Validate_Call {
	_c.Call.Return(_a0)
	return _c
}