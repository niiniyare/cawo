// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/flight.proto

package flight

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Flight service

func NewFlightEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Flight service

type FlightService interface {
	CreateAircraft(ctx context.Context, in *CreateAircraftParams, opts ...client.CallOption) (*AircraftsDatum, error)
	CreateAirlineCompany(ctx context.Context, in *CreateAirlineCompanyParams, opts ...client.CallOption) (*AirlineCompany, error)
	CreateAirportList(ctx context.Context, in *CreateAirportListParams, opts ...client.CallOption) (*CreateAirportListResponse, error)
	CreateAirports(ctx context.Context, in *CreateAirportsParams, opts ...client.CallOption) (*AirportsDatum, error)
	DeleteAircraft(ctx context.Context, in *DeleteAircraftParams, opts ...client.CallOption) (*emptypb.Empty, error)
	DeleteAirlineCompany(ctx context.Context, in *DeleteAirlineCompanyParams, opts ...client.CallOption) (*emptypb.Empty, error)
	DeleteAirports(ctx context.Context, in *DeleteAirportsParams, opts ...client.CallOption) (*emptypb.Empty, error)
	GetAircraft(ctx context.Context, in *GetAircraftParams, opts ...client.CallOption) (*AircraftsDatum, error)
	GetAirlineCompany(ctx context.Context, in *GetAirlineCompanyParams, opts ...client.CallOption) (*AirlineCompany, error)
	GetAirports(ctx context.Context, in *GetAirportsParams, opts ...client.CallOption) (*AirportsDatum, error)
	ListAircraft(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*ListAircraftResponse, error)
	ListAirlineCompany(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*ListAirlineCompanyResponse, error)
	ListAirports(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*ListAirportsResponse, error)
}

type flightService struct {
	c    client.Client
	name string
}

func NewFlightService(name string, c client.Client) FlightService {
	return &flightService{
		c:    c,
		name: name,
	}
}

func (c *flightService) CreateAircraft(ctx context.Context, in *CreateAircraftParams, opts ...client.CallOption) (*AircraftsDatum, error) {
	req := c.c.NewRequest(c.name, "Flight.CreateAircraft", in)
	out := new(AircraftsDatum)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightService) CreateAirlineCompany(ctx context.Context, in *CreateAirlineCompanyParams, opts ...client.CallOption) (*AirlineCompany, error) {
	req := c.c.NewRequest(c.name, "Flight.CreateAirlineCompany", in)
	out := new(AirlineCompany)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightService) CreateAirportList(ctx context.Context, in *CreateAirportListParams, opts ...client.CallOption) (*CreateAirportListResponse, error) {
	req := c.c.NewRequest(c.name, "Flight.CreateAirportList", in)
	out := new(CreateAirportListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightService) CreateAirports(ctx context.Context, in *CreateAirportsParams, opts ...client.CallOption) (*AirportsDatum, error) {
	req := c.c.NewRequest(c.name, "Flight.CreateAirports", in)
	out := new(AirportsDatum)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightService) DeleteAircraft(ctx context.Context, in *DeleteAircraftParams, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Flight.DeleteAircraft", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightService) DeleteAirlineCompany(ctx context.Context, in *DeleteAirlineCompanyParams, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Flight.DeleteAirlineCompany", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightService) DeleteAirports(ctx context.Context, in *DeleteAirportsParams, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Flight.DeleteAirports", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightService) GetAircraft(ctx context.Context, in *GetAircraftParams, opts ...client.CallOption) (*AircraftsDatum, error) {
	req := c.c.NewRequest(c.name, "Flight.GetAircraft", in)
	out := new(AircraftsDatum)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightService) GetAirlineCompany(ctx context.Context, in *GetAirlineCompanyParams, opts ...client.CallOption) (*AirlineCompany, error) {
	req := c.c.NewRequest(c.name, "Flight.GetAirlineCompany", in)
	out := new(AirlineCompany)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightService) GetAirports(ctx context.Context, in *GetAirportsParams, opts ...client.CallOption) (*AirportsDatum, error) {
	req := c.c.NewRequest(c.name, "Flight.GetAirports", in)
	out := new(AirportsDatum)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightService) ListAircraft(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*ListAircraftResponse, error) {
	req := c.c.NewRequest(c.name, "Flight.ListAircraft", in)
	out := new(ListAircraftResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightService) ListAirlineCompany(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*ListAirlineCompanyResponse, error) {
	req := c.c.NewRequest(c.name, "Flight.ListAirlineCompany", in)
	out := new(ListAirlineCompanyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightService) ListAirports(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*ListAirportsResponse, error) {
	req := c.c.NewRequest(c.name, "Flight.ListAirports", in)
	out := new(ListAirportsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Flight service

type FlightHandler interface {
	CreateAircraft(context.Context, *CreateAircraftParams, *AircraftsDatum) error
	CreateAirlineCompany(context.Context, *CreateAirlineCompanyParams, *AirlineCompany) error
	CreateAirportList(context.Context, *CreateAirportListParams, *CreateAirportListResponse) error
	CreateAirports(context.Context, *CreateAirportsParams, *AirportsDatum) error
	DeleteAircraft(context.Context, *DeleteAircraftParams, *emptypb.Empty) error
	DeleteAirlineCompany(context.Context, *DeleteAirlineCompanyParams, *emptypb.Empty) error
	DeleteAirports(context.Context, *DeleteAirportsParams, *emptypb.Empty) error
	GetAircraft(context.Context, *GetAircraftParams, *AircraftsDatum) error
	GetAirlineCompany(context.Context, *GetAirlineCompanyParams, *AirlineCompany) error
	GetAirports(context.Context, *GetAirportsParams, *AirportsDatum) error
	ListAircraft(context.Context, *emptypb.Empty, *ListAircraftResponse) error
	ListAirlineCompany(context.Context, *emptypb.Empty, *ListAirlineCompanyResponse) error
	ListAirports(context.Context, *emptypb.Empty, *ListAirportsResponse) error
}

func RegisterFlightHandler(s server.Server, hdlr FlightHandler, opts ...server.HandlerOption) error {
	type flight interface {
		CreateAircraft(ctx context.Context, in *CreateAircraftParams, out *AircraftsDatum) error
		CreateAirlineCompany(ctx context.Context, in *CreateAirlineCompanyParams, out *AirlineCompany) error
		CreateAirportList(ctx context.Context, in *CreateAirportListParams, out *CreateAirportListResponse) error
		CreateAirports(ctx context.Context, in *CreateAirportsParams, out *AirportsDatum) error
		DeleteAircraft(ctx context.Context, in *DeleteAircraftParams, out *emptypb.Empty) error
		DeleteAirlineCompany(ctx context.Context, in *DeleteAirlineCompanyParams, out *emptypb.Empty) error
		DeleteAirports(ctx context.Context, in *DeleteAirportsParams, out *emptypb.Empty) error
		GetAircraft(ctx context.Context, in *GetAircraftParams, out *AircraftsDatum) error
		GetAirlineCompany(ctx context.Context, in *GetAirlineCompanyParams, out *AirlineCompany) error
		GetAirports(ctx context.Context, in *GetAirportsParams, out *AirportsDatum) error
		ListAircraft(ctx context.Context, in *emptypb.Empty, out *ListAircraftResponse) error
		ListAirlineCompany(ctx context.Context, in *emptypb.Empty, out *ListAirlineCompanyResponse) error
		ListAirports(ctx context.Context, in *emptypb.Empty, out *ListAirportsResponse) error
	}
	type Flight struct {
		flight
	}
	h := &flightHandler{hdlr}
	return s.Handle(s.NewHandler(&Flight{h}, opts...))
}

type flightHandler struct {
	FlightHandler
}

func (h *flightHandler) CreateAircraft(ctx context.Context, in *CreateAircraftParams, out *AircraftsDatum) error {
	return h.FlightHandler.CreateAircraft(ctx, in, out)
}

func (h *flightHandler) CreateAirlineCompany(ctx context.Context, in *CreateAirlineCompanyParams, out *AirlineCompany) error {
	return h.FlightHandler.CreateAirlineCompany(ctx, in, out)
}

func (h *flightHandler) CreateAirportList(ctx context.Context, in *CreateAirportListParams, out *CreateAirportListResponse) error {
	return h.FlightHandler.CreateAirportList(ctx, in, out)
}

func (h *flightHandler) CreateAirports(ctx context.Context, in *CreateAirportsParams, out *AirportsDatum) error {
	return h.FlightHandler.CreateAirports(ctx, in, out)
}

func (h *flightHandler) DeleteAircraft(ctx context.Context, in *DeleteAircraftParams, out *emptypb.Empty) error {
	return h.FlightHandler.DeleteAircraft(ctx, in, out)
}

func (h *flightHandler) DeleteAirlineCompany(ctx context.Context, in *DeleteAirlineCompanyParams, out *emptypb.Empty) error {
	return h.FlightHandler.DeleteAirlineCompany(ctx, in, out)
}

func (h *flightHandler) DeleteAirports(ctx context.Context, in *DeleteAirportsParams, out *emptypb.Empty) error {
	return h.FlightHandler.DeleteAirports(ctx, in, out)
}

func (h *flightHandler) GetAircraft(ctx context.Context, in *GetAircraftParams, out *AircraftsDatum) error {
	return h.FlightHandler.GetAircraft(ctx, in, out)
}

func (h *flightHandler) GetAirlineCompany(ctx context.Context, in *GetAirlineCompanyParams, out *AirlineCompany) error {
	return h.FlightHandler.GetAirlineCompany(ctx, in, out)
}

func (h *flightHandler) GetAirports(ctx context.Context, in *GetAirportsParams, out *AirportsDatum) error {
	return h.FlightHandler.GetAirports(ctx, in, out)
}

func (h *flightHandler) ListAircraft(ctx context.Context, in *emptypb.Empty, out *ListAircraftResponse) error {
	return h.FlightHandler.ListAircraft(ctx, in, out)
}

func (h *flightHandler) ListAirlineCompany(ctx context.Context, in *emptypb.Empty, out *ListAirlineCompanyResponse) error {
	return h.FlightHandler.ListAirlineCompany(ctx, in, out)
}

func (h *flightHandler) ListAirports(ctx context.Context, in *emptypb.Empty, out *ListAirportsResponse) error {
	return h.FlightHandler.ListAirports(ctx, in, out)
}
