// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/etadmin/admin.proto

package etadmin

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = strings.TrimPrefix

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AdminLogoutReq struct {
}

func (m *AdminLogoutReq) Validate() error {
	return nil
}

func (m *AdminLogoutReq) Reset()         { *m = AdminLogoutReq{} }
func (m *AdminLogoutReq) String() string { return proto.CompactTextString(m) }
func (*AdminLogoutReq) ProtoMessage()    {}
func (*AdminLogoutReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f0376662d8a45d, []int{0}
}

func (m *AdminLogoutReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminLogoutReq.Unmarshal(m, b)
}
func (m *AdminLogoutReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminLogoutReq.Marshal(b, m, deterministic)
}
func (m *AdminLogoutReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminLogoutReq.Merge(m, src)
}
func (m *AdminLogoutReq) XXX_Size() int {
	return xxx_messageInfo_AdminLogoutReq.Size(m)
}
func (m *AdminLogoutReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminLogoutReq.DiscardUnknown(m)
}

var xxx_messageInfo_AdminLogoutReq proto.InternalMessageInfo

type AdminLogoutResp struct {
}

func (m *AdminLogoutResp) Validate() error {
	return nil
}

func (m *AdminLogoutResp) Reset()         { *m = AdminLogoutResp{} }
func (m *AdminLogoutResp) String() string { return proto.CompactTextString(m) }
func (*AdminLogoutResp) ProtoMessage()    {}
func (*AdminLogoutResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f0376662d8a45d, []int{1}
}

func (m *AdminLogoutResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminLogoutResp.Unmarshal(m, b)
}
func (m *AdminLogoutResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminLogoutResp.Marshal(b, m, deterministic)
}
func (m *AdminLogoutResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminLogoutResp.Merge(m, src)
}
func (m *AdminLogoutResp) XXX_Size() int {
	return xxx_messageInfo_AdminLogoutResp.Size(m)
}
func (m *AdminLogoutResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminLogoutResp.DiscardUnknown(m)
}

var xxx_messageInfo_AdminLogoutResp proto.InternalMessageInfo

type AdminLoginReq struct {
	//
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username"`
	//
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password"`
	//
	SourceCode string `protobuf:"bytes,3,opt,name=sourceCode,proto3" json:"sourceCode"`
}

func (m *AdminLoginReq) Validate() error {
	return nil
}

func (m *AdminLoginReq) Reset()         { *m = AdminLoginReq{} }
func (m *AdminLoginReq) String() string { return proto.CompactTextString(m) }
func (*AdminLoginReq) ProtoMessage()    {}
func (*AdminLoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f0376662d8a45d, []int{2}
}

func (m *AdminLoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminLoginReq.Unmarshal(m, b)
}
func (m *AdminLoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminLoginReq.Marshal(b, m, deterministic)
}
func (m *AdminLoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminLoginReq.Merge(m, src)
}
func (m *AdminLoginReq) XXX_Size() int {
	return xxx_messageInfo_AdminLoginReq.Size(m)
}
func (m *AdminLoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminLoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_AdminLoginReq proto.InternalMessageInfo

func (m *AdminLoginReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AdminLoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AdminLoginReq) GetSourceCode() string {
	if m != nil {
		return m.SourceCode
	}
	return ""
}

type AdminLoginResp struct {
	//
	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status"`
	//
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`
	//
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message"`
	//
	Data *LoginData `protobuf:"bytes,4,opt,name=data,proto3" json:"data"`
}

func (m *AdminLoginResp) Validate() error {
	return nil
}

func (m *AdminLoginResp) Reset()         { *m = AdminLoginResp{} }
func (m *AdminLoginResp) String() string { return proto.CompactTextString(m) }
func (*AdminLoginResp) ProtoMessage()    {}
func (*AdminLoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f0376662d8a45d, []int{3}
}

func (m *AdminLoginResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminLoginResp.Unmarshal(m, b)
}
func (m *AdminLoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminLoginResp.Marshal(b, m, deterministic)
}
func (m *AdminLoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminLoginResp.Merge(m, src)
}
func (m *AdminLoginResp) XXX_Size() int {
	return xxx_messageInfo_AdminLoginResp.Size(m)
}
func (m *AdminLoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminLoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_AdminLoginResp proto.InternalMessageInfo

func (m *AdminLoginResp) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *AdminLoginResp) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *AdminLoginResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AdminLoginResp) GetData() *LoginData {
	if m != nil {
		return m.Data
	}
	return nil
}

type LoginData struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId"`
	//
	NeedModify bool `protobuf:"varint,2,opt,name=needModify,proto3" json:"needModify"`
	//
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token"`
	//
	TimeToExpire int32 `protobuf:"varint,4,opt,name=timeToExpire,proto3" json:"timeToExpire"`
	//
	MenuFunctionsData *MenuFunctionsData `protobuf:"bytes,5,opt,name=menuFunctionsData,proto3" json:"menuFunctionsData"`
	//
	GroupsData *GroupsData `protobuf:"bytes,6,opt,name=groupsData,proto3" json:"groupsData"`
	//
	RolesData *RolesData `protobuf:"bytes,7,opt,name=rolesData,proto3" json:"rolesData"`
	//
	ModifyMgs string `protobuf:"bytes,8,opt,name=modifyMgs,proto3" json:"modifyMgs"`
}

func (m *LoginData) Validate() error {
	return nil
}

func (m *LoginData) Reset()         { *m = LoginData{} }
func (m *LoginData) String() string { return proto.CompactTextString(m) }
func (*LoginData) ProtoMessage()    {}
func (*LoginData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f0376662d8a45d, []int{4}
}

func (m *LoginData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginData.Unmarshal(m, b)
}
func (m *LoginData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginData.Marshal(b, m, deterministic)
}
func (m *LoginData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginData.Merge(m, src)
}
func (m *LoginData) XXX_Size() int {
	return xxx_messageInfo_LoginData.Size(m)
}
func (m *LoginData) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginData.DiscardUnknown(m)
}

var xxx_messageInfo_LoginData proto.InternalMessageInfo

func (m *LoginData) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *LoginData) GetNeedModify() bool {
	if m != nil {
		return m.NeedModify
	}
	return false
}

func (m *LoginData) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginData) GetTimeToExpire() int32 {
	if m != nil {
		return m.TimeToExpire
	}
	return 0
}

func (m *LoginData) GetMenuFunctionsData() *MenuFunctionsData {
	if m != nil {
		return m.MenuFunctionsData
	}
	return nil
}

func (m *LoginData) GetGroupsData() *GroupsData {
	if m != nil {
		return m.GroupsData
	}
	return nil
}

func (m *LoginData) GetRolesData() *RolesData {
	if m != nil {
		return m.RolesData
	}
	return nil
}

func (m *LoginData) GetModifyMgs() string {
	if m != nil {
		return m.ModifyMgs
	}
	return ""
}

type GroupInfo struct {
	//
	GroupId int64 `protobuf:"varint,1,opt,name=groupId,proto3" json:"groupId"`
	//
	GroupName string `protobuf:"bytes,2,opt,name=groupName,proto3" json:"groupName"`
}

func (m *GroupInfo) Validate() error {
	return nil
}

func (m *GroupInfo) Reset()         { *m = GroupInfo{} }
func (m *GroupInfo) String() string { return proto.CompactTextString(m) }
func (*GroupInfo) ProtoMessage()    {}
func (*GroupInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f0376662d8a45d, []int{5}
}

func (m *GroupInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupInfo.Unmarshal(m, b)
}
func (m *GroupInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupInfo.Marshal(b, m, deterministic)
}
func (m *GroupInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupInfo.Merge(m, src)
}
func (m *GroupInfo) XXX_Size() int {
	return xxx_messageInfo_GroupInfo.Size(m)
}
func (m *GroupInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GroupInfo proto.InternalMessageInfo

func (m *GroupInfo) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *GroupInfo) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

type GroupsData struct {
	//
	GroupInfo []*GroupInfo `protobuf:"bytes,1,rep,name=groupInfo,proto3" json:"groupInfo"`
}

func (m *GroupsData) Validate() error {
	return nil
}

func (m *GroupsData) Reset()         { *m = GroupsData{} }
func (m *GroupsData) String() string { return proto.CompactTextString(m) }
func (*GroupsData) ProtoMessage()    {}
func (*GroupsData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f0376662d8a45d, []int{6}
}

func (m *GroupsData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupsData.Unmarshal(m, b)
}
func (m *GroupsData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupsData.Marshal(b, m, deterministic)
}
func (m *GroupsData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupsData.Merge(m, src)
}
func (m *GroupsData) XXX_Size() int {
	return xxx_messageInfo_GroupsData.Size(m)
}
func (m *GroupsData) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupsData.DiscardUnknown(m)
}

var xxx_messageInfo_GroupsData proto.InternalMessageInfo

func (m *GroupsData) GetGroupInfo() []*GroupInfo {
	if m != nil {
		return m.GroupInfo
	}
	return nil
}

type RoleInfo struct {
	//
	RoleId int64 `protobuf:"varint,1,opt,name=roleId,proto3" json:"roleId"`
	//
	RoleName string `protobuf:"bytes,2,opt,name=roleName,proto3" json:"roleName"`
}

func (m *RoleInfo) Validate() error {
	return nil
}

func (m *RoleInfo) Reset()         { *m = RoleInfo{} }
func (m *RoleInfo) String() string { return proto.CompactTextString(m) }
func (*RoleInfo) ProtoMessage()    {}
func (*RoleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f0376662d8a45d, []int{7}
}

func (m *RoleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleInfo.Unmarshal(m, b)
}
func (m *RoleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleInfo.Marshal(b, m, deterministic)
}
func (m *RoleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleInfo.Merge(m, src)
}
func (m *RoleInfo) XXX_Size() int {
	return xxx_messageInfo_RoleInfo.Size(m)
}
func (m *RoleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RoleInfo proto.InternalMessageInfo

func (m *RoleInfo) GetRoleId() int64 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *RoleInfo) GetRoleName() string {
	if m != nil {
		return m.RoleName
	}
	return ""
}

type RolesData struct {
	//
	RoleInfo []*RoleInfo `protobuf:"bytes,1,rep,name=roleInfo,proto3" json:"roleInfo"`
}

func (m *RolesData) Validate() error {
	return nil
}

func (m *RolesData) Reset()         { *m = RolesData{} }
func (m *RolesData) String() string { return proto.CompactTextString(m) }
func (*RolesData) ProtoMessage()    {}
func (*RolesData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f0376662d8a45d, []int{8}
}

func (m *RolesData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RolesData.Unmarshal(m, b)
}
func (m *RolesData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RolesData.Marshal(b, m, deterministic)
}
func (m *RolesData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolesData.Merge(m, src)
}
func (m *RolesData) XXX_Size() int {
	return xxx_messageInfo_RolesData.Size(m)
}
func (m *RolesData) XXX_DiscardUnknown() {
	xxx_messageInfo_RolesData.DiscardUnknown(m)
}

var xxx_messageInfo_RolesData proto.InternalMessageInfo

func (m *RolesData) GetRoleInfo() []*RoleInfo {
	if m != nil {
		return m.RoleInfo
	}
	return nil
}

type MenuFunctionsData struct {
	//
	Menus []*Menus `protobuf:"bytes,1,rep,name=menus,proto3" json:"menus"`
	//
	Functions []*Function `protobuf:"bytes,2,rep,name=functions,proto3" json:"functions"`
}

func (m *MenuFunctionsData) Validate() error {
	return nil
}

func (m *MenuFunctionsData) Reset()         { *m = MenuFunctionsData{} }
func (m *MenuFunctionsData) String() string { return proto.CompactTextString(m) }
func (*MenuFunctionsData) ProtoMessage()    {}
func (*MenuFunctionsData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f0376662d8a45d, []int{9}
}

func (m *MenuFunctionsData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MenuFunctionsData.Unmarshal(m, b)
}
func (m *MenuFunctionsData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MenuFunctionsData.Marshal(b, m, deterministic)
}
func (m *MenuFunctionsData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MenuFunctionsData.Merge(m, src)
}
func (m *MenuFunctionsData) XXX_Size() int {
	return xxx_messageInfo_MenuFunctionsData.Size(m)
}
func (m *MenuFunctionsData) XXX_DiscardUnknown() {
	xxx_messageInfo_MenuFunctionsData.DiscardUnknown(m)
}

var xxx_messageInfo_MenuFunctionsData proto.InternalMessageInfo

func (m *MenuFunctionsData) GetMenus() []*Menus {
	if m != nil {
		return m.Menus
	}
	return nil
}

func (m *MenuFunctionsData) GetFunctions() []*Function {
	if m != nil {
		return m.Functions
	}
	return nil
}

type Menus struct {
	//
	MenuId int64 `protobuf:"varint,1,opt,name=menuId,proto3" json:"menuId"`
	//
	MenuName string `protobuf:"bytes,2,opt,name=menuName,proto3" json:"menuName"`
	//
	MenuLevel int32 `protobuf:"varint,3,opt,name=menuLevel,proto3" json:"menuLevel"`
	//
	Url string `protobuf:"bytes,4,opt,name=url,proto3" json:"url"`
	//
	IsUnfold bool `protobuf:"varint,5,opt,name=isUnfold,proto3" json:"isUnfold"`
	//
	Order int32 `protobuf:"varint,6,opt,name=order,proto3" json:"order"`
	//
	NewWindow bool `protobuf:"varint,7,opt,name=newWindow,proto3" json:"newWindow"`
	//
	Icon string `protobuf:"bytes,8,opt,name=icon,proto3" json:"icon"`
	//
	ChildMenus []*Menus `protobuf:"bytes,9,rep,name=childMenus,proto3" json:"childMenus"`
}

func (m *Menus) Validate() error {
	return nil
}

func (m *Menus) Reset()         { *m = Menus{} }
func (m *Menus) String() string { return proto.CompactTextString(m) }
func (*Menus) ProtoMessage()    {}
func (*Menus) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f0376662d8a45d, []int{10}
}

func (m *Menus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Menus.Unmarshal(m, b)
}
func (m *Menus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Menus.Marshal(b, m, deterministic)
}
func (m *Menus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Menus.Merge(m, src)
}
func (m *Menus) XXX_Size() int {
	return xxx_messageInfo_Menus.Size(m)
}
func (m *Menus) XXX_DiscardUnknown() {
	xxx_messageInfo_Menus.DiscardUnknown(m)
}

var xxx_messageInfo_Menus proto.InternalMessageInfo

func (m *Menus) GetMenuId() int64 {
	if m != nil {
		return m.MenuId
	}
	return 0
}

func (m *Menus) GetMenuName() string {
	if m != nil {
		return m.MenuName
	}
	return ""
}

func (m *Menus) GetMenuLevel() int32 {
	if m != nil {
		return m.MenuLevel
	}
	return 0
}

func (m *Menus) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Menus) GetIsUnfold() bool {
	if m != nil {
		return m.IsUnfold
	}
	return false
}

func (m *Menus) GetOrder() int32 {
	if m != nil {
		return m.Order
	}
	return 0
}

func (m *Menus) GetNewWindow() bool {
	if m != nil {
		return m.NewWindow
	}
	return false
}

func (m *Menus) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Menus) GetChildMenus() []*Menus {
	if m != nil {
		return m.ChildMenus
	}
	return nil
}

type Function struct {
	//
	MenuId int32 `protobuf:"varint,1,opt,name=menuId,proto3" json:"menuId"`
	//
	FunctionName string `protobuf:"bytes,2,opt,name=functionName,proto3" json:"functionName"`
	//
	FunctionCode string `protobuf:"bytes,3,opt,name=functionCode,proto3" json:"functionCode"`
	//
	InterfaceUrl string `protobuf:"bytes,4,opt,name=interfaceUrl,proto3" json:"interfaceUrl"`
}

func (m *Function) Validate() error {
	return nil
}

func (m *Function) Reset()         { *m = Function{} }
func (m *Function) String() string { return proto.CompactTextString(m) }
func (*Function) ProtoMessage()    {}
func (*Function) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f0376662d8a45d, []int{11}
}

func (m *Function) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Function.Unmarshal(m, b)
}
func (m *Function) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Function.Marshal(b, m, deterministic)
}
func (m *Function) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Function.Merge(m, src)
}
func (m *Function) XXX_Size() int {
	return xxx_messageInfo_Function.Size(m)
}
func (m *Function) XXX_DiscardUnknown() {
	xxx_messageInfo_Function.DiscardUnknown(m)
}

var xxx_messageInfo_Function proto.InternalMessageInfo

func (m *Function) GetMenuId() int32 {
	if m != nil {
		return m.MenuId
	}
	return 0
}

func (m *Function) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *Function) GetFunctionCode() string {
	if m != nil {
		return m.FunctionCode
	}
	return ""
}

func (m *Function) GetInterfaceUrl() string {
	if m != nil {
		return m.InterfaceUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*AdminLogoutReq)(nil), "etadmin.AdminLogoutReq")
	proto.RegisterType((*AdminLogoutResp)(nil), "etadmin.AdminLogoutResp")
	proto.RegisterType((*AdminLoginReq)(nil), "etadmin.AdminLoginReq")
	proto.RegisterType((*AdminLoginResp)(nil), "etadmin.AdminLoginResp")
	proto.RegisterType((*LoginData)(nil), "etadmin.LoginData")
	proto.RegisterType((*GroupInfo)(nil), "etadmin.GroupInfo")
	proto.RegisterType((*GroupsData)(nil), "etadmin.GroupsData")
	proto.RegisterType((*RoleInfo)(nil), "etadmin.RoleInfo")
	proto.RegisterType((*RolesData)(nil), "etadmin.RolesData")
	proto.RegisterType((*MenuFunctionsData)(nil), "etadmin.MenuFunctionsData")
	proto.RegisterType((*Menus)(nil), "etadmin.Menus")
	proto.RegisterType((*Function)(nil), "etadmin.Function")
}

func init() { proto.RegisterFile("proto/etadmin/admin.proto", fileDescriptor_f0f0376662d8a45d) }

var fileDescriptor_f0f0376662d8a45d = []byte{
	// 623 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x04, 0x03, 0x02, 0x01, 0x02, 0xff, 0x6c, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0x96, 0x9b, 0x3a, 0xb5, 0xe7, 0xd7, 0x5f, 0x69, 0x16, 0x54, 0x99, 0x0a, 0xa1, 0x6a, 0x85,
	0x50, 0x2f, 0x24, 0x88, 0xde, 0x38, 0x54, 0x42, 0x05, 0x4a, 0xa5, 0x86, 0xc3, 0x8a, 0x8a, 0xb3,
	0xb1, 0x37, 0xc6, 0x10, 0xef, 0xba, 0x5e, 0x9b, 0xc0, 0x85, 0x47, 0xe0, 0xc9, 0x78, 0x27, 0xd0,
	0x8c, 0xbd, 0x6b, 0x3b, 0xe5, 0x12, 0xcd, 0x7c, 0x33, 0x3b, 0xdf, 0x37, 0x7f, 0x62, 0x78, 0x58,
	0x56, 0xba, 0xd6, 0x0b, 0x59, 0xc7, 0x69, 0x91, 0xab, 0x05, 0xfd, 0xce, 0x09, 0x63, 0x7b, 0x1d,
	0xc8, 0x0f, 0xe1, 0xe0, 0x15, 0x1a, 0xd7, 0x3a, 0xd3, 0x4d, 0x2d, 0xe4, 0x2d, 0x9f, 0xc1, 0xbd,
	0x11, 0x62, 0x4a, 0x9e, 0xc1, 0xff, 0x16, 0xca, 0x95, 0x90, 0xb7, 0xec, 0x18, 0x82, 0xc6, 0xc8,
	0x4a, 0xc5, 0x85, 0x8c, 0xbc, 0x13, 0xef, 0x34, 0x14, 0xce, 0xc7, 0x58, 0x19, 0x1b, 0xb3, 0xd1,
	0x55, 0x1a, 0xed, 0xb4, 0x31, 0xeb, 0xb3, 0xc7, 0x00, 0x46, 0x37, 0x55, 0x22, 0x2f, 0x74, 0x2a,
	0xa3, 0x09, 0x45, 0x07, 0x08, 0xff, 0xd9, 0xab, 0x41, 0x22, 0x53, 0xb2, 0x23, 0x98, 0x9a, 0x3a,
	0xae, 0x1b, 0x43, 0x3c, 0x81, 0xe8, 0x3c, 0xc6, 0x60, 0x37, 0xc1, 0x1a, 0x2d, 0x03, 0xd9, 0x2c,
	0x82, 0xbd, 0x42, 0x1a, 0x13, 0x67, 0xb6, 0xb4, 0x75, 0xd9, 0x53, 0xd8, 0x4d, 0xe3, 0x3a, 0x8e,
	0x76, 0x4f, 0xbc, 0xd3, 0xff, 0x5e, 0xb0, 0x79, 0xd7, 0xfd, 0x9c, 0x78, 0x5e, 0xc7, 0x75, 0x2c,
	0x28, 0xce, 0x7f, 0xef, 0x40, 0xe8, 0x30, 0xe4, 0xc6, 0xae, 0xae, 0x52, 0xe2, 0x9e, 0x88, 0xce,
	0xc3, 0x2e, 0x94, 0x94, 0xe9, 0x52, 0xa7, 0xf9, 0xea, 0x07, 0x29, 0x08, 0xc4, 0x00, 0x61, 0x0f,
	0xc0, 0xaf, 0xf5, 0x57, 0xa9, 0x3a, 0x15, 0xad, 0xc3, 0x38, 0xec, 0xd7, 0x79, 0x21, 0x3f, 0xe8,
	0x37, 0xdf, 0xcb, 0xbc, 0x92, 0xa4, 0xc5, 0x17, 0x23, 0x8c, 0xbd, 0x83, 0x59, 0x21, 0x55, 0xf3,
	0xb6, 0x51, 0x49, 0x9d, 0x6b, 0x65, 0x50, 0x46, 0xe4, 0x93, 0xe8, 0x63, 0x27, 0x7a, 0xb9, 0x9d,
	0x21, 0xee, 0x3e, 0x62, 0x67, 0x00, 0x59, 0xa5, 0x9b, 0xb2, 0x2d, 0x31, 0xa5, 0x12, 0xf7, 0x5d,
	0x89, 0x4b, 0x17, 0x12, 0x83, 0x34, 0xf6, 0x1c, 0xc2, 0x4a, 0xaf, 0x65, 0xfb, 0x66, 0x6f, 0x6b,
	0x56, 0xc2, 0x46, 0x44, 0x9f, 0xc4, 0x1e, 0x41, 0x58, 0x50, 0xd3, 0xcb, 0xcc, 0x44, 0x01, 0xb5,
	0xdb, 0x03, 0xfc, 0x02, 0x42, 0x62, 0xba, 0x52, 0x2b, 0x8d, 0xdb, 0x21, 0x2a, 0x37, 0x4e, 0xeb,
	0x62, 0x11, 0x32, 0xdf, 0xe3, 0x39, 0xb5, 0x0b, 0xed, 0x01, 0x7e, 0x0e, 0x70, 0x39, 0x92, 0x98,
	0xd9, 0x92, 0x91, 0x77, 0x32, 0x19, 0x49, 0x74, 0x64, 0xa2, 0x4f, 0xe2, 0xe7, 0x10, 0xa0, 0x74,
	0xd2, 0x70, 0x04, 0x53, 0xd4, 0xde, 0x6f, 0xb4, 0xf5, 0xf0, 0x66, 0xd1, 0x1a, 0x08, 0x70, 0x3e,
	0x7f, 0x09, 0xa1, 0x6b, 0x9d, 0x3d, 0x6b, 0x13, 0x07, 0xec, 0xb3, 0xd1, 0x80, 0x88, 0xdc, 0xa5,
	0xf0, 0x2f, 0x30, 0xbb, 0xb3, 0x2d, 0xf6, 0x04, 0x7c, 0xdc, 0x97, 0xe9, 0x0a, 0x1c, 0x8c, 0x16,
	0x6b, 0x44, 0x1b, 0x64, 0x0b, 0x08, 0x57, 0xf6, 0x59, 0xb4, 0xb3, 0x45, 0x65, 0x0b, 0x8a, 0x3e,
	0x87, 0xff, 0xf1, 0xc0, 0xa7, 0x0a, 0xd8, 0x25, 0xd6, 0xe8, 0xbb, 0x6c, 0x3d, 0xec, 0x12, 0xad,
	0x61, 0x97, 0xd6, 0xa7, 0x45, 0x4a, 0xd5, 0x5c, 0xcb, 0x6f, 0x72, 0x4d, 0x77, 0xeb, 0x8b, 0x1e,
	0x60, 0x87, 0x30, 0x69, 0xaa, 0x35, 0x9d, 0x6c, 0x28, 0xd0, 0xc4, 0x5a, 0xb9, 0xb9, 0x51, 0x2b,
	0xbd, 0x4e, 0xe9, 0x40, 0x03, 0xe1, 0x7c, 0xbc, 0x7f, 0x5d, 0xa5, 0xb2, 0xa2, 0xb3, 0xf3, 0x45,
	0xeb, 0x20, 0x83, 0x92, 0x9b, 0x8f, 0xb9, 0x4a, 0xf5, 0x86, 0x8e, 0x2b, 0x10, 0x3d, 0x80, 0xff,
	0xe7, 0x3c, 0xd1, 0xaa, 0xbb, 0x21, 0xb2, 0xd9, 0x1c, 0x20, 0xf9, 0x9c, 0xaf, 0x53, 0xea, 0x2a,
	0x0a, 0xff, 0x39, 0xad, 0x41, 0x06, 0xff, 0xe5, 0x41, 0x60, 0x27, 0xb3, 0x35, 0x04, 0xdf, 0x0d,
	0x81, 0xc3, 0xbe, 0x9d, 0xd9, 0x60, 0x10, 0x23, 0x6c, 0x98, 0x33, 0xf8, 0x50, 0x8d, 0x30, 0xcc,
	0xc9, 0x55, 0x2d, 0xab, 0x55, 0x9c, 0xc8, 0x1b, 0x37, 0x9b, 0x11, 0xf6, 0x69, 0x4a, 0x1f, 0xdb,
	0xb3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x72, 0xa0, 0x0d, 0x89, 0x05, 0x00, 0x00,
}