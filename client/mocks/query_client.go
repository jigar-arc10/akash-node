// Code generated by mockery 2.12.1. DO NOT EDIT.

package mocks

import (
	certtypesv1beta2 "github.com/akash-network/node/x/cert/types/v1beta2"

	context "context"

	deploymenttypesv1beta2 "github.com/akash-network/node/x/deployment/types/v1beta2"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	providertypesv1beta2 "github.com/akash-network/node/x/provider/types/v1beta2"

	testing "testing"

	typesv1beta2 "github.com/akash-network/node/x/market/types/v1beta2"

	v1beta2 "github.com/akash-network/node/x/audit/types/v1beta2"
)

// QueryClient is an autogenerated mock type for the QueryClient type
type QueryClient struct {
	mock.Mock
}

// AllProvidersAttributes provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) AllProvidersAttributes(ctx context.Context, in *v1beta2.QueryAllProvidersAttributesRequest, opts ...grpc.CallOption) (*v1beta2.QueryProvidersResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1beta2.QueryProvidersResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta2.QueryAllProvidersAttributesRequest, ...grpc.CallOption) *v1beta2.QueryProvidersResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta2.QueryProvidersResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta2.QueryAllProvidersAttributesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuditorAttributes provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) AuditorAttributes(ctx context.Context, in *v1beta2.QueryAuditorAttributesRequest, opts ...grpc.CallOption) (*v1beta2.QueryProvidersResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1beta2.QueryProvidersResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta2.QueryAuditorAttributesRequest, ...grpc.CallOption) *v1beta2.QueryProvidersResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta2.QueryProvidersResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta2.QueryAuditorAttributesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Bid provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) Bid(ctx context.Context, in *typesv1beta2.QueryBidRequest, opts ...grpc.CallOption) (*typesv1beta2.QueryBidResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *typesv1beta2.QueryBidResponse
	if rf, ok := ret.Get(0).(func(context.Context, *typesv1beta2.QueryBidRequest, ...grpc.CallOption) *typesv1beta2.QueryBidResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*typesv1beta2.QueryBidResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *typesv1beta2.QueryBidRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Bids provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) Bids(ctx context.Context, in *typesv1beta2.QueryBidsRequest, opts ...grpc.CallOption) (*typesv1beta2.QueryBidsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *typesv1beta2.QueryBidsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *typesv1beta2.QueryBidsRequest, ...grpc.CallOption) *typesv1beta2.QueryBidsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*typesv1beta2.QueryBidsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *typesv1beta2.QueryBidsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Certificates provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) Certificates(ctx context.Context, in *certtypesv1beta2.QueryCertificatesRequest, opts ...grpc.CallOption) (*certtypesv1beta2.QueryCertificatesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *certtypesv1beta2.QueryCertificatesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *certtypesv1beta2.QueryCertificatesRequest, ...grpc.CallOption) *certtypesv1beta2.QueryCertificatesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*certtypesv1beta2.QueryCertificatesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *certtypesv1beta2.QueryCertificatesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Deployment provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) Deployment(ctx context.Context, in *deploymenttypesv1beta2.QueryDeploymentRequest, opts ...grpc.CallOption) (*deploymenttypesv1beta2.QueryDeploymentResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *deploymenttypesv1beta2.QueryDeploymentResponse
	if rf, ok := ret.Get(0).(func(context.Context, *deploymenttypesv1beta2.QueryDeploymentRequest, ...grpc.CallOption) *deploymenttypesv1beta2.QueryDeploymentResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*deploymenttypesv1beta2.QueryDeploymentResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *deploymenttypesv1beta2.QueryDeploymentRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Deployments provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) Deployments(ctx context.Context, in *deploymenttypesv1beta2.QueryDeploymentsRequest, opts ...grpc.CallOption) (*deploymenttypesv1beta2.QueryDeploymentsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *deploymenttypesv1beta2.QueryDeploymentsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *deploymenttypesv1beta2.QueryDeploymentsRequest, ...grpc.CallOption) *deploymenttypesv1beta2.QueryDeploymentsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*deploymenttypesv1beta2.QueryDeploymentsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *deploymenttypesv1beta2.QueryDeploymentsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Group provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) Group(ctx context.Context, in *deploymenttypesv1beta2.QueryGroupRequest, opts ...grpc.CallOption) (*deploymenttypesv1beta2.QueryGroupResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *deploymenttypesv1beta2.QueryGroupResponse
	if rf, ok := ret.Get(0).(func(context.Context, *deploymenttypesv1beta2.QueryGroupRequest, ...grpc.CallOption) *deploymenttypesv1beta2.QueryGroupResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*deploymenttypesv1beta2.QueryGroupResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *deploymenttypesv1beta2.QueryGroupRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Lease provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) Lease(ctx context.Context, in *typesv1beta2.QueryLeaseRequest, opts ...grpc.CallOption) (*typesv1beta2.QueryLeaseResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *typesv1beta2.QueryLeaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *typesv1beta2.QueryLeaseRequest, ...grpc.CallOption) *typesv1beta2.QueryLeaseResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*typesv1beta2.QueryLeaseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *typesv1beta2.QueryLeaseRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Leases provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) Leases(ctx context.Context, in *typesv1beta2.QueryLeasesRequest, opts ...grpc.CallOption) (*typesv1beta2.QueryLeasesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *typesv1beta2.QueryLeasesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *typesv1beta2.QueryLeasesRequest, ...grpc.CallOption) *typesv1beta2.QueryLeasesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*typesv1beta2.QueryLeasesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *typesv1beta2.QueryLeasesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Order provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) Order(ctx context.Context, in *typesv1beta2.QueryOrderRequest, opts ...grpc.CallOption) (*typesv1beta2.QueryOrderResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *typesv1beta2.QueryOrderResponse
	if rf, ok := ret.Get(0).(func(context.Context, *typesv1beta2.QueryOrderRequest, ...grpc.CallOption) *typesv1beta2.QueryOrderResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*typesv1beta2.QueryOrderResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *typesv1beta2.QueryOrderRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Orders provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) Orders(ctx context.Context, in *typesv1beta2.QueryOrdersRequest, opts ...grpc.CallOption) (*typesv1beta2.QueryOrdersResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *typesv1beta2.QueryOrdersResponse
	if rf, ok := ret.Get(0).(func(context.Context, *typesv1beta2.QueryOrdersRequest, ...grpc.CallOption) *typesv1beta2.QueryOrdersResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*typesv1beta2.QueryOrdersResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *typesv1beta2.QueryOrdersRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Provider provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) Provider(ctx context.Context, in *providertypesv1beta2.QueryProviderRequest, opts ...grpc.CallOption) (*providertypesv1beta2.QueryProviderResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *providertypesv1beta2.QueryProviderResponse
	if rf, ok := ret.Get(0).(func(context.Context, *providertypesv1beta2.QueryProviderRequest, ...grpc.CallOption) *providertypesv1beta2.QueryProviderResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*providertypesv1beta2.QueryProviderResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *providertypesv1beta2.QueryProviderRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderAttributes provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) ProviderAttributes(ctx context.Context, in *v1beta2.QueryProviderAttributesRequest, opts ...grpc.CallOption) (*v1beta2.QueryProvidersResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1beta2.QueryProvidersResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta2.QueryProviderAttributesRequest, ...grpc.CallOption) *v1beta2.QueryProvidersResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta2.QueryProvidersResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta2.QueryProviderAttributesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderAuditorAttributes provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) ProviderAuditorAttributes(ctx context.Context, in *v1beta2.QueryProviderAuditorRequest, opts ...grpc.CallOption) (*v1beta2.QueryProvidersResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1beta2.QueryProvidersResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta2.QueryProviderAuditorRequest, ...grpc.CallOption) *v1beta2.QueryProvidersResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta2.QueryProvidersResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta2.QueryProviderAuditorRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Providers provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) Providers(ctx context.Context, in *providertypesv1beta2.QueryProvidersRequest, opts ...grpc.CallOption) (*providertypesv1beta2.QueryProvidersResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *providertypesv1beta2.QueryProvidersResponse
	if rf, ok := ret.Get(0).(func(context.Context, *providertypesv1beta2.QueryProvidersRequest, ...grpc.CallOption) *providertypesv1beta2.QueryProvidersResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*providertypesv1beta2.QueryProvidersResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *providertypesv1beta2.QueryProvidersRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewQueryClient creates a new instance of QueryClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewQueryClient(t testing.TB) *QueryClient {
	mock := &QueryClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
