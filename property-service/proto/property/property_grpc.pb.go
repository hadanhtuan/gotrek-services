// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: proto/property/property.proto

package property

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	sdk "property-service/proto/sdk"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PropertyServiceClient is the client API for PropertyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PropertyServiceClient interface {
	// Booking
	GetBooking(ctx context.Context, in *MsgQueryBooking, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	CountBookingStatus(ctx context.Context, in *MsgBooking, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	CreateBooking(ctx context.Context, in *MsgBooking, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	CancelBooking(ctx context.Context, in *MsgBooking, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	// Property
	GetProperty(ctx context.Context, in *MsgQueryProperty, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	CountPropertyStatus(ctx context.Context, in *MsgProperty, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	CreateProperty(ctx context.Context, in *MsgProperty, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	UpdateProperty(ctx context.Context, in *MsgProperty, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	DeleteProperty(ctx context.Context, in *MsgDeleteProperty, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	// Review
	CreateReview(ctx context.Context, in *MsgCreateReview, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	UpdateReview(ctx context.Context, in *MsgUpdateReview, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	DeleteReview(ctx context.Context, in *MsgDeleteReview, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	GetReview(ctx context.Context, in *MsgQueryReview, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	// Amenity
	CreateAmenity(ctx context.Context, in *MsgAmenity, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	UpdateAmenity(ctx context.Context, in *MsgAmenity, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	DeleteAmenity(ctx context.Context, in *MsgId, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	GetAmenity(ctx context.Context, in *MsgQueryAmenity, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	// Favorite
	CreateFavorite(ctx context.Context, in *MsgFavorite, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	DeleteFavorite(ctx context.Context, in *MsgId, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	GetFavorite(ctx context.Context, in *MsgQueryFavorite, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
}

type propertyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPropertyServiceClient(cc grpc.ClientConnInterface) PropertyServiceClient {
	return &propertyServiceClient{cc}
}

func (c *propertyServiceClient) GetBooking(ctx context.Context, in *MsgQueryBooking, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/GetBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) CountBookingStatus(ctx context.Context, in *MsgBooking, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/CountBookingStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) CreateBooking(ctx context.Context, in *MsgBooking, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/CreateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) CancelBooking(ctx context.Context, in *MsgBooking, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/CancelBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) GetProperty(ctx context.Context, in *MsgQueryProperty, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/GetProperty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) CountPropertyStatus(ctx context.Context, in *MsgProperty, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/CountPropertyStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) CreateProperty(ctx context.Context, in *MsgProperty, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/CreateProperty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) UpdateProperty(ctx context.Context, in *MsgProperty, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/UpdateProperty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) DeleteProperty(ctx context.Context, in *MsgDeleteProperty, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/DeleteProperty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) CreateReview(ctx context.Context, in *MsgCreateReview, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/CreateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) UpdateReview(ctx context.Context, in *MsgUpdateReview, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/UpdateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) DeleteReview(ctx context.Context, in *MsgDeleteReview, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/DeleteReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) GetReview(ctx context.Context, in *MsgQueryReview, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/GetReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) CreateAmenity(ctx context.Context, in *MsgAmenity, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/CreateAmenity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) UpdateAmenity(ctx context.Context, in *MsgAmenity, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/UpdateAmenity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) DeleteAmenity(ctx context.Context, in *MsgId, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/DeleteAmenity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) GetAmenity(ctx context.Context, in *MsgQueryAmenity, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/GetAmenity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) CreateFavorite(ctx context.Context, in *MsgFavorite, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/CreateFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) DeleteFavorite(ctx context.Context, in *MsgId, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/DeleteFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) GetFavorite(ctx context.Context, in *MsgQueryFavorite, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/propertyService.propertyService/GetFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PropertyServiceServer is the server API for PropertyService service.
// All implementations must embed UnimplementedPropertyServiceServer
// for forward compatibility
type PropertyServiceServer interface {
	// Booking
	GetBooking(context.Context, *MsgQueryBooking) (*sdk.BaseResponse, error)
	CountBookingStatus(context.Context, *MsgBooking) (*sdk.BaseResponse, error)
	CreateBooking(context.Context, *MsgBooking) (*sdk.BaseResponse, error)
	CancelBooking(context.Context, *MsgBooking) (*sdk.BaseResponse, error)
	// Property
	GetProperty(context.Context, *MsgQueryProperty) (*sdk.BaseResponse, error)
	CountPropertyStatus(context.Context, *MsgProperty) (*sdk.BaseResponse, error)
	CreateProperty(context.Context, *MsgProperty) (*sdk.BaseResponse, error)
	UpdateProperty(context.Context, *MsgProperty) (*sdk.BaseResponse, error)
	DeleteProperty(context.Context, *MsgDeleteProperty) (*sdk.BaseResponse, error)
	// Review
	CreateReview(context.Context, *MsgCreateReview) (*sdk.BaseResponse, error)
	UpdateReview(context.Context, *MsgUpdateReview) (*sdk.BaseResponse, error)
	DeleteReview(context.Context, *MsgDeleteReview) (*sdk.BaseResponse, error)
	GetReview(context.Context, *MsgQueryReview) (*sdk.BaseResponse, error)
	// Amenity
	CreateAmenity(context.Context, *MsgAmenity) (*sdk.BaseResponse, error)
	UpdateAmenity(context.Context, *MsgAmenity) (*sdk.BaseResponse, error)
	DeleteAmenity(context.Context, *MsgId) (*sdk.BaseResponse, error)
	GetAmenity(context.Context, *MsgQueryAmenity) (*sdk.BaseResponse, error)
	// Favorite
	CreateFavorite(context.Context, *MsgFavorite) (*sdk.BaseResponse, error)
	DeleteFavorite(context.Context, *MsgId) (*sdk.BaseResponse, error)
	GetFavorite(context.Context, *MsgQueryFavorite) (*sdk.BaseResponse, error)
	mustEmbedUnimplementedPropertyServiceServer()
}

// UnimplementedPropertyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPropertyServiceServer struct {
}

func (UnimplementedPropertyServiceServer) GetBooking(context.Context, *MsgQueryBooking) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooking not implemented")
}
func (UnimplementedPropertyServiceServer) CountBookingStatus(context.Context, *MsgBooking) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountBookingStatus not implemented")
}
func (UnimplementedPropertyServiceServer) CreateBooking(context.Context, *MsgBooking) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedPropertyServiceServer) CancelBooking(context.Context, *MsgBooking) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBooking not implemented")
}
func (UnimplementedPropertyServiceServer) GetProperty(context.Context, *MsgQueryProperty) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProperty not implemented")
}
func (UnimplementedPropertyServiceServer) CountPropertyStatus(context.Context, *MsgProperty) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountPropertyStatus not implemented")
}
func (UnimplementedPropertyServiceServer) CreateProperty(context.Context, *MsgProperty) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProperty not implemented")
}
func (UnimplementedPropertyServiceServer) UpdateProperty(context.Context, *MsgProperty) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProperty not implemented")
}
func (UnimplementedPropertyServiceServer) DeleteProperty(context.Context, *MsgDeleteProperty) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProperty not implemented")
}
func (UnimplementedPropertyServiceServer) CreateReview(context.Context, *MsgCreateReview) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReview not implemented")
}
func (UnimplementedPropertyServiceServer) UpdateReview(context.Context, *MsgUpdateReview) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReview not implemented")
}
func (UnimplementedPropertyServiceServer) DeleteReview(context.Context, *MsgDeleteReview) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReview not implemented")
}
func (UnimplementedPropertyServiceServer) GetReview(context.Context, *MsgQueryReview) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReview not implemented")
}
func (UnimplementedPropertyServiceServer) CreateAmenity(context.Context, *MsgAmenity) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAmenity not implemented")
}
func (UnimplementedPropertyServiceServer) UpdateAmenity(context.Context, *MsgAmenity) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAmenity not implemented")
}
func (UnimplementedPropertyServiceServer) DeleteAmenity(context.Context, *MsgId) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAmenity not implemented")
}
func (UnimplementedPropertyServiceServer) GetAmenity(context.Context, *MsgQueryAmenity) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAmenity not implemented")
}
func (UnimplementedPropertyServiceServer) CreateFavorite(context.Context, *MsgFavorite) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFavorite not implemented")
}
func (UnimplementedPropertyServiceServer) DeleteFavorite(context.Context, *MsgId) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFavorite not implemented")
}
func (UnimplementedPropertyServiceServer) GetFavorite(context.Context, *MsgQueryFavorite) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavorite not implemented")
}
func (UnimplementedPropertyServiceServer) mustEmbedUnimplementedPropertyServiceServer() {}

// UnsafePropertyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PropertyServiceServer will
// result in compilation errors.
type UnsafePropertyServiceServer interface {
	mustEmbedUnimplementedPropertyServiceServer()
}

func RegisterPropertyServiceServer(s grpc.ServiceRegistrar, srv PropertyServiceServer) {
	s.RegisterService(&PropertyService_ServiceDesc, srv)
}

func _PropertyService_GetBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgQueryBooking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).GetBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/GetBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).GetBooking(ctx, req.(*MsgQueryBooking))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_CountBookingStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBooking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).CountBookingStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/CountBookingStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).CountBookingStatus(ctx, req.(*MsgBooking))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBooking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/CreateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).CreateBooking(ctx, req.(*MsgBooking))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_CancelBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBooking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).CancelBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/CancelBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).CancelBooking(ctx, req.(*MsgBooking))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_GetProperty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgQueryProperty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).GetProperty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/GetProperty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).GetProperty(ctx, req.(*MsgQueryProperty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_CountPropertyStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgProperty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).CountPropertyStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/CountPropertyStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).CountPropertyStatus(ctx, req.(*MsgProperty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_CreateProperty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgProperty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).CreateProperty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/CreateProperty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).CreateProperty(ctx, req.(*MsgProperty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_UpdateProperty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgProperty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).UpdateProperty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/UpdateProperty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).UpdateProperty(ctx, req.(*MsgProperty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_DeleteProperty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteProperty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).DeleteProperty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/DeleteProperty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).DeleteProperty(ctx, req.(*MsgDeleteProperty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_CreateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateReview)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).CreateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/CreateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).CreateReview(ctx, req.(*MsgCreateReview))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_UpdateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateReview)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).UpdateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/UpdateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).UpdateReview(ctx, req.(*MsgUpdateReview))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_DeleteReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteReview)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).DeleteReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/DeleteReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).DeleteReview(ctx, req.(*MsgDeleteReview))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_GetReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgQueryReview)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).GetReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/GetReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).GetReview(ctx, req.(*MsgQueryReview))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_CreateAmenity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAmenity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).CreateAmenity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/CreateAmenity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).CreateAmenity(ctx, req.(*MsgAmenity))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_UpdateAmenity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAmenity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).UpdateAmenity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/UpdateAmenity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).UpdateAmenity(ctx, req.(*MsgAmenity))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_DeleteAmenity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).DeleteAmenity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/DeleteAmenity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).DeleteAmenity(ctx, req.(*MsgId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_GetAmenity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgQueryAmenity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).GetAmenity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/GetAmenity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).GetAmenity(ctx, req.(*MsgQueryAmenity))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_CreateFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFavorite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).CreateFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/CreateFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).CreateFavorite(ctx, req.(*MsgFavorite))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_DeleteFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).DeleteFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/DeleteFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).DeleteFavorite(ctx, req.(*MsgId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_GetFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgQueryFavorite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).GetFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/propertyService.propertyService/GetFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).GetFavorite(ctx, req.(*MsgQueryFavorite))
	}
	return interceptor(ctx, in, info, handler)
}

// PropertyService_ServiceDesc is the grpc.ServiceDesc for PropertyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PropertyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "propertyService.propertyService",
	HandlerType: (*PropertyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBooking",
			Handler:    _PropertyService_GetBooking_Handler,
		},
		{
			MethodName: "CountBookingStatus",
			Handler:    _PropertyService_CountBookingStatus_Handler,
		},
		{
			MethodName: "CreateBooking",
			Handler:    _PropertyService_CreateBooking_Handler,
		},
		{
			MethodName: "CancelBooking",
			Handler:    _PropertyService_CancelBooking_Handler,
		},
		{
			MethodName: "GetProperty",
			Handler:    _PropertyService_GetProperty_Handler,
		},
		{
			MethodName: "CountPropertyStatus",
			Handler:    _PropertyService_CountPropertyStatus_Handler,
		},
		{
			MethodName: "CreateProperty",
			Handler:    _PropertyService_CreateProperty_Handler,
		},
		{
			MethodName: "UpdateProperty",
			Handler:    _PropertyService_UpdateProperty_Handler,
		},
		{
			MethodName: "DeleteProperty",
			Handler:    _PropertyService_DeleteProperty_Handler,
		},
		{
			MethodName: "CreateReview",
			Handler:    _PropertyService_CreateReview_Handler,
		},
		{
			MethodName: "UpdateReview",
			Handler:    _PropertyService_UpdateReview_Handler,
		},
		{
			MethodName: "DeleteReview",
			Handler:    _PropertyService_DeleteReview_Handler,
		},
		{
			MethodName: "GetReview",
			Handler:    _PropertyService_GetReview_Handler,
		},
		{
			MethodName: "CreateAmenity",
			Handler:    _PropertyService_CreateAmenity_Handler,
		},
		{
			MethodName: "UpdateAmenity",
			Handler:    _PropertyService_UpdateAmenity_Handler,
		},
		{
			MethodName: "DeleteAmenity",
			Handler:    _PropertyService_DeleteAmenity_Handler,
		},
		{
			MethodName: "GetAmenity",
			Handler:    _PropertyService_GetAmenity_Handler,
		},
		{
			MethodName: "CreateFavorite",
			Handler:    _PropertyService_CreateFavorite_Handler,
		},
		{
			MethodName: "DeleteFavorite",
			Handler:    _PropertyService_DeleteFavorite_Handler,
		},
		{
			MethodName: "GetFavorite",
			Handler:    _PropertyService_GetFavorite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/property/property.proto",
}
