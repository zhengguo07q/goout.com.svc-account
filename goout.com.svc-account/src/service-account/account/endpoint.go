package account

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"

	"github.com/go-kit/kit/examples/shipping/cargo"
	"github.com/go-kit/kit/examples/shipping/location"
)

//购物车的请求数据
type bookCargoRequest struct {
	Origin          location.UNLocode
	Destination     location.UNLocode
	ArrivalDeadline time.Time
}

//购物车的响应数据ff
type bookCargoResponse struct {
	ID  cargo.TrackingID `json:"tracking_id,omitempty"`
	Err error            `json:"error,omitempty"`
}

//返回购物车的报错信息
func (r bookCargoResponse) error() error { return r.Err }

//制作购物车的监听点
func makeBookCargoEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(bookCargoRequest)
		id, err := s.BookNewCargo(req.Origin, req.Destination, req.ArrivalDeadline)
		return bookCargoResponse{ID: id, Err: err}, nil
	}
}

//载入购物车请求
type loadCargoRequest struct {
	ID cargo.TrackingID
}

//返回购物车载入请求
type loadCargoResponse struct {
	Cargo *Cargo `json:"cargo,omitempty"`
	Err   error  `json:"error,omitempty"`
}

//错误
func (r loadCargoResponse) error() error { return r.Err }

//制载入购物车的监听点
func makeLoadCargoEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loadCargoRequest)
		c, err := s.LoadCargo(req.ID)
		return loadCargoResponse{Cargo: &c, Err: err}, nil
	}
}

//请求路径路由请求
type requestRoutesRequest struct {
	ID cargo.TrackingID
}

//响应路径路由请求
type requestRoutesResponse struct {
	Routes []cargo.Itinerary `json:"routes,omitempty"`
	Err    error             `json:"error,omitempty"`
}

func (r requestRoutesResponse) error() error { return r.Err }

//制作请求路由的监听点
func makeRequestRoutesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requestRoutesRequest)
		itin := s.RequestPossibleRoutesForCargo(req.ID)
		return requestRoutesResponse{Routes: itin, Err: nil}, nil
	}
}

//关联到路由请求
type assignToRouteRequest struct {
	ID        cargo.TrackingID
	Itinerary cargo.Itinerary
}

//关联到路由响应
type assignToRouteResponse struct {
	Err error `json:"error,omitempty"`
}

func (r assignToRouteResponse) error() error { return r.Err }

//制作关联路由的监听点
func makeAssignToRouteEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(assignToRouteRequest)
		err := s.AssignCargoToRoute(req.ID, req.Itinerary)
		return assignToRouteResponse{Err: err}, nil
	}
}

//改变目的地请求
type changeDestinationRequest struct {
	ID          cargo.TrackingID
	Destination location.UNLocode
}

//改变目的地响应
type changeDestinationResponse struct {
	Err error `json:"error,omitempty"`
}

func (r changeDestinationResponse) error() error { return r.Err }

//制作改变目的地的监听点
func makeChangeDestinationEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(changeDestinationRequest)
		err := s.ChangeDestination(req.ID, req.Destination)
		return changeDestinationResponse{Err: err}, nil
	}
}

//列出购物车请求
type listCargosRequest struct{}

//列出购物车响应
type listCargosResponse struct {
	Cargos []Cargo `json:"cargos,omitempty"`
	Err    error   `json:"error,omitempty"`
}

func (r listCargosResponse) error() error { return r.Err }

//制作列出购物车的监听点
func makeListCargosEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listCargosRequest)
		return listCargosResponse{Cargos: s.Cargos(), Err: nil}, nil
	}
}

//列出支持的区域请求
type listLocationsRequest struct {
}

//列出支持的区域响应
type listLocationsResponse struct {
	Locations []Location `json:"locations,omitempty"`
	Err       error      `json:"error,omitempty"`
}

//制作支持的区域的监听点
func makeListLocationsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listLocationsRequest)
		return listLocationsResponse{Locations: s.Locations(), Err: nil}, nil
	}
}
