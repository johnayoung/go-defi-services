// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/coingecko.proto

package coingecko

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Api Endpoints for Coingecko service

func NewCoingeckoEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Coingecko service

type CoingeckoService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Coingecko_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Coingecko_PingPongService, error)
}

type coingeckoService struct {
	c    client.Client
	name string
}

func NewCoingeckoService(name string, c client.Client) CoingeckoService {
	return &coingeckoService{
		c:    c,
		name: name,
	}
}

func (c *coingeckoService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Coingecko.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coingeckoService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Coingecko_StreamService, error) {
	req := c.c.NewRequest(c.name, "Coingecko.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &coingeckoServiceStream{stream}, nil
}

type Coingecko_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type coingeckoServiceStream struct {
	stream client.Stream
}

func (x *coingeckoServiceStream) Close() error {
	return x.stream.Close()
}

func (x *coingeckoServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *coingeckoServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *coingeckoServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *coingeckoServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *coingeckoService) PingPong(ctx context.Context, opts ...client.CallOption) (Coingecko_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Coingecko.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &coingeckoServicePingPong{stream}, nil
}

type Coingecko_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type coingeckoServicePingPong struct {
	stream client.Stream
}

func (x *coingeckoServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *coingeckoServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *coingeckoServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *coingeckoServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *coingeckoServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *coingeckoServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Coingecko service

type CoingeckoHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Coingecko_StreamStream) error
	PingPong(context.Context, Coingecko_PingPongStream) error
}

func RegisterCoingeckoHandler(s server.Server, hdlr CoingeckoHandler, opts ...server.HandlerOption) error {
	type coingecko interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Coingecko struct {
		coingecko
	}
	h := &coingeckoHandler{hdlr}
	return s.Handle(s.NewHandler(&Coingecko{h}, opts...))
}

type coingeckoHandler struct {
	CoingeckoHandler
}

func (h *coingeckoHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.CoingeckoHandler.Call(ctx, in, out)
}

func (h *coingeckoHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.CoingeckoHandler.Stream(ctx, m, &coingeckoStreamStream{stream})
}

type Coingecko_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type coingeckoStreamStream struct {
	stream server.Stream
}

func (x *coingeckoStreamStream) Close() error {
	return x.stream.Close()
}

func (x *coingeckoStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *coingeckoStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *coingeckoStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *coingeckoStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *coingeckoHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.CoingeckoHandler.PingPong(ctx, &coingeckoPingPongStream{stream})
}

type Coingecko_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type coingeckoPingPongStream struct {
	stream server.Stream
}

func (x *coingeckoPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *coingeckoPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *coingeckoPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *coingeckoPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *coingeckoPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *coingeckoPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
