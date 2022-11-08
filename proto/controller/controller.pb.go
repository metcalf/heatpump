// Protocol buffers for the controller program that interfaces
// with the heat pump and fan coil units.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: controller.proto

package controller

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	fancoil "github.com/gonzojive/heatpump/proto/fancoil"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Requests the state of a device .
type GetDeviceStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the device.
	//
	// For fancoils, looks like "fan-coils/[id]".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetDeviceStateRequest) Reset() {
	*x = GetDeviceStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceStateRequest) ProtoMessage() {}

func (x *GetDeviceStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceStateRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceStateRequest) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{0}
}

func (x *GetDeviceStateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Describes the state of a fan coil unit or other device managed by the
// controller.
type DeviceState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the device.
	//
	// For fancoils, looks like "fan-coils/[id]".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// If the device is a fan coil unit, the state of the fan coil unit.
	FancoilState *fancoil.State `protobuf:"bytes,2,opt,name=fancoil_state,json=fancoilState,proto3" json:"fancoil_state,omitempty"`
}

func (x *DeviceState) Reset() {
	*x = DeviceState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceState) ProtoMessage() {}

func (x *DeviceState) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceState.ProtoReflect.Descriptor instead.
func (*DeviceState) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceState) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceState) GetFancoilState() *fancoil.State {
	if x != nil {
		return x.FancoilState
	}
	return nil
}

// Sets the state of the device.
type SetDeviceStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State *DeviceState `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *SetDeviceStateRequest) Reset() {
	*x = SetDeviceStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDeviceStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDeviceStateRequest) ProtoMessage() {}

func (x *SetDeviceStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDeviceStateRequest.ProtoReflect.Descriptor instead.
func (*SetDeviceStateRequest) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{2}
}

func (x *SetDeviceStateRequest) GetState() *DeviceState {
	if x != nil {
		return x.State
	}
	return nil
}

type SetDeviceStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetDeviceStateResponse) Reset() {
	*x = SetDeviceStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDeviceStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDeviceStateResponse) ProtoMessage() {}

func (x *SetDeviceStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDeviceStateResponse.ProtoReflect.Descriptor instead.
func (*SetDeviceStateResponse) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{3}
}

type Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudIotDeviceName string `protobuf:"bytes,1,opt,name=cloud_iot_device_name,json=cloudIotDeviceName,proto3" json:"cloud_iot_device_name,omitempty"`
	// Path to USB-to-RS485 device connected to modbus to use for talking to the heat pump (CX34).
	HeatpumpModbusDevicePath string `protobuf:"bytes,2,opt,name=heatpump_modbus_device_path,json=heatpumpModbusDevicePath,proto3" json:"heatpump_modbus_device_path,omitempty"`
	// Path to USB-to-RS485 device connected to modbus to use for talking to the heat pump (CX34).
	FanCoilsModbusDevicePath string `protobuf:"bytes,3,opt,name=fan_coils_modbus_device_path,json=fanCoilsModbusDevicePath,proto3" json:"fan_coils_modbus_device_path,omitempty"`
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{4}
}

func (x *Configuration) GetCloudIotDeviceName() string {
	if x != nil {
		return x.CloudIotDeviceName
	}
	return ""
}

func (x *Configuration) GetHeatpumpModbusDevicePath() string {
	if x != nil {
		return x.HeatpumpModbusDevicePath
	}
	return ""
}

func (x *Configuration) GetFanCoilsModbusDevicePath() string {
	if x != nil {
		return x.FanCoilsModbusDevicePath
	}
	return ""
}

// Command is used to issue commands to the controller from pub/sub.
type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Command:
	//	*Command_SetStateRequest
	Command isCommand_Command `protobuf_oneof:"command"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{5}
}

func (m *Command) GetCommand() isCommand_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (x *Command) GetSetStateRequest() *fancoil.SetStateRequest {
	if x, ok := x.GetCommand().(*Command_SetStateRequest); ok {
		return x.SetStateRequest
	}
	return nil
}

type isCommand_Command interface {
	isCommand_Command()
}

type Command_SetStateRequest struct {
	SetStateRequest *fancoil.SetStateRequest `protobuf:"bytes,1,opt,name=set_state_request,json=setStateRequest,proto3,oneof"`
}

func (*Command_SetStateRequest) isCommand_Command() {}

// A wire-encoded version of this proto is passed from the device to the server
// to authenticate.
type DeviceAccessToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string               `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Expiration *timestamp.Timestamp `protobuf:"bytes,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	// Signature provided by the server of the wire-encoded message without the
	// signature field.
	Signature string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *DeviceAccessToken) Reset() {
	*x = DeviceAccessToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceAccessToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceAccessToken) ProtoMessage() {}

func (x *DeviceAccessToken) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceAccessToken.ProtoReflect.Descriptor instead.
func (*DeviceAccessToken) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{6}
}

func (x *DeviceAccessToken) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeviceAccessToken) GetExpiration() *timestamp.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

func (x *DeviceAccessToken) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type ExtendTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ExtendTokenRequest) Reset() {
	*x = ExtendTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendTokenRequest) ProtoMessage() {}

func (x *ExtendTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendTokenRequest.ProtoReflect.Descriptor instead.
func (*ExtendTokenRequest) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{7}
}

func (x *ExtendTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ExtendTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshedToken string               `protobuf:"bytes,1,opt,name=refreshed_token,json=refreshedToken,proto3" json:"refreshed_token,omitempty"`
	Expiration     *timestamp.Timestamp `protobuf:"bytes,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *ExtendTokenResponse) Reset() {
	*x = ExtendTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendTokenResponse) ProtoMessage() {}

func (x *ExtendTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendTokenResponse.ProtoReflect.Descriptor instead.
func (*ExtendTokenResponse) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{8}
}

func (x *ExtendTokenResponse) GetRefreshedToken() string {
	if x != nil {
		return x.RefreshedToken
	}
	return ""
}

func (x *ExtendTokenResponse) GetExpiration() *timestamp.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

var File_controller_proto protoreflect.FileDescriptor

var file_controller_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x1a, 0x1b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x61, 0x6e, 0x63, 0x6f, 0x69, 0x6c, 0x2f, 0x66, 0x61,
	0x6e, 0x63, 0x6f, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x0b, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x0d,
	0x66, 0x61, 0x6e, 0x63, 0x6f, 0x69, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x61, 0x6e, 0x63, 0x6f, 0x69, 0x6c, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x0c, 0x66, 0x61, 0x6e, 0x63, 0x6f, 0x69, 0x6c, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x46, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x53, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0xc1, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x15, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x69,
	0x6f, 0x74, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6f, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x1b, 0x68, 0x65, 0x61, 0x74,
	0x70, 0x75, 0x6d, 0x70, 0x5f, 0x6d, 0x6f, 0x64, 0x62, 0x75, 0x73, 0x5f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x68,
	0x65, 0x61, 0x74, 0x70, 0x75, 0x6d, 0x70, 0x4d, 0x6f, 0x64, 0x62, 0x75, 0x73, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x3e, 0x0a, 0x1c, 0x66, 0x61, 0x6e, 0x5f, 0x63,
	0x6f, 0x69, 0x6c, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x62, 0x75, 0x73, 0x5f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x66,
	0x61, 0x6e, 0x43, 0x6f, 0x69, 0x6c, 0x73, 0x4d, 0x6f, 0x64, 0x62, 0x75, 0x73, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x5c, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x46, 0x0a, 0x11, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x66, 0x61, 0x6e, 0x63, 0x6f, 0x69, 0x6c, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x2a,
	0x0a, 0x12, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7a, 0x0a, 0x13, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3a, 0x0a, 0x0a, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xb9, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x32, 0x5f, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x50, 0x0a, 0x0b, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x6e, 0x7a, 0x6f, 0x6a, 0x69, 0x76, 0x65, 0x2f, 0x68, 0x65, 0x61, 0x74,
	0x70, 0x75, 0x6d, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_proto_rawDescOnce sync.Once
	file_controller_proto_rawDescData = file_controller_proto_rawDesc
)

func file_controller_proto_rawDescGZIP() []byte {
	file_controller_proto_rawDescOnce.Do(func() {
		file_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_proto_rawDescData)
	})
	return file_controller_proto_rawDescData
}

var file_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_controller_proto_goTypes = []interface{}{
	(*GetDeviceStateRequest)(nil),   // 0: controller.GetDeviceStateRequest
	(*DeviceState)(nil),             // 1: controller.DeviceState
	(*SetDeviceStateRequest)(nil),   // 2: controller.SetDeviceStateRequest
	(*SetDeviceStateResponse)(nil),  // 3: controller.SetDeviceStateResponse
	(*Configuration)(nil),           // 4: controller.Configuration
	(*Command)(nil),                 // 5: controller.Command
	(*DeviceAccessToken)(nil),       // 6: controller.DeviceAccessToken
	(*ExtendTokenRequest)(nil),      // 7: controller.ExtendTokenRequest
	(*ExtendTokenResponse)(nil),     // 8: controller.ExtendTokenResponse
	(*fancoil.State)(nil),           // 9: fancoil.State
	(*fancoil.SetStateRequest)(nil), // 10: fancoil.SetStateRequest
	(*timestamp.Timestamp)(nil),     // 11: google.protobuf.Timestamp
}
var file_controller_proto_depIdxs = []int32{
	9,  // 0: controller.DeviceState.fancoil_state:type_name -> fancoil.State
	1,  // 1: controller.SetDeviceStateRequest.state:type_name -> controller.DeviceState
	10, // 2: controller.Command.set_state_request:type_name -> fancoil.SetStateRequest
	11, // 3: controller.DeviceAccessToken.expiration:type_name -> google.protobuf.Timestamp
	11, // 4: controller.ExtendTokenResponse.expiration:type_name -> google.protobuf.Timestamp
	0,  // 5: controller.StateService.GetDeviceState:input_type -> controller.GetDeviceStateRequest
	2,  // 6: controller.StateService.SetDeviceState:input_type -> controller.SetDeviceStateRequest
	7,  // 7: controller.AuthService.ExtendToken:input_type -> controller.ExtendTokenRequest
	1,  // 8: controller.StateService.GetDeviceState:output_type -> controller.DeviceState
	3,  // 9: controller.StateService.SetDeviceState:output_type -> controller.SetDeviceStateResponse
	8,  // 10: controller.AuthService.ExtendToken:output_type -> controller.ExtendTokenResponse
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_controller_proto_init() }
func file_controller_proto_init() {
	if File_controller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceStateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDeviceStateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDeviceStateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configuration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceAccessToken); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendTokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendTokenResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_controller_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Command_SetStateRequest)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_controller_proto_goTypes,
		DependencyIndexes: file_controller_proto_depIdxs,
		MessageInfos:      file_controller_proto_msgTypes,
	}.Build()
	File_controller_proto = out.File
	file_controller_proto_rawDesc = nil
	file_controller_proto_goTypes = nil
	file_controller_proto_depIdxs = nil
}