// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notification.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// [START messages]
type Link struct {
	Href                 string   `protobuf:"bytes,1,opt,name=href,proto3" json:"href,omitempty"`
	Methods              []string `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty"`
	Rel                  string   `protobuf:"bytes,3,opt,name=rel,proto3" json:"rel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}
func (*Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{0}
}

func (m *Link) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Link.Unmarshal(m, b)
}
func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Link.Marshal(b, m, deterministic)
}
func (m *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(m, src)
}
func (m *Link) XXX_Size() int {
	return xxx_messageInfo_Link.Size(m)
}
func (m *Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Link proto.InternalMessageInfo

func (m *Link) GetHref() string {
	if m != nil {
		return m.Href
	}
	return ""
}

func (m *Link) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

func (m *Link) GetRel() string {
	if m != nil {
		return m.Rel
	}
	return ""
}

type Order struct {
	OrderId              string            `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Links                *Order_OrderLinks `protobuf:"bytes,2,opt,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{1}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *Order) GetLinks() *Order_OrderLinks {
	if m != nil {
		return m.Links
	}
	return nil
}

type Order_OrderLinks struct {
	Self                 *Link    `protobuf:"bytes,1,opt,name=self,proto3" json:"self,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order_OrderLinks) Reset()         { *m = Order_OrderLinks{} }
func (m *Order_OrderLinks) String() string { return proto.CompactTextString(m) }
func (*Order_OrderLinks) ProtoMessage()    {}
func (*Order_OrderLinks) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{1, 0}
}

func (m *Order_OrderLinks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order_OrderLinks.Unmarshal(m, b)
}
func (m *Order_OrderLinks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order_OrderLinks.Marshal(b, m, deterministic)
}
func (m *Order_OrderLinks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order_OrderLinks.Merge(m, src)
}
func (m *Order_OrderLinks) XXX_Size() int {
	return xxx_messageInfo_Order_OrderLinks.Size(m)
}
func (m *Order_OrderLinks) XXX_DiscardUnknown() {
	xxx_messageInfo_Order_OrderLinks.DiscardUnknown(m)
}

var xxx_messageInfo_Order_OrderLinks proto.InternalMessageInfo

func (m *Order_OrderLinks) GetSelf() *Link {
	if m != nil {
		return m.Self
	}
	return nil
}

type Fulfiller struct {
	Links                *Fulfiller_FulfillerLinks `protobuf:"bytes,1,opt,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Fulfiller) Reset()         { *m = Fulfiller{} }
func (m *Fulfiller) String() string { return proto.CompactTextString(m) }
func (*Fulfiller) ProtoMessage()    {}
func (*Fulfiller) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{2}
}

func (m *Fulfiller) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fulfiller.Unmarshal(m, b)
}
func (m *Fulfiller) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fulfiller.Marshal(b, m, deterministic)
}
func (m *Fulfiller) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fulfiller.Merge(m, src)
}
func (m *Fulfiller) XXX_Size() int {
	return xxx_messageInfo_Fulfiller.Size(m)
}
func (m *Fulfiller) XXX_DiscardUnknown() {
	xxx_messageInfo_Fulfiller.DiscardUnknown(m)
}

var xxx_messageInfo_Fulfiller proto.InternalMessageInfo

func (m *Fulfiller) GetLinks() *Fulfiller_FulfillerLinks {
	if m != nil {
		return m.Links
	}
	return nil
}

type Fulfiller_FulfillerLinks struct {
	Self                 *Link    `protobuf:"bytes,1,opt,name=self,proto3" json:"self,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fulfiller_FulfillerLinks) Reset()         { *m = Fulfiller_FulfillerLinks{} }
func (m *Fulfiller_FulfillerLinks) String() string { return proto.CompactTextString(m) }
func (*Fulfiller_FulfillerLinks) ProtoMessage()    {}
func (*Fulfiller_FulfillerLinks) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{2, 0}
}

func (m *Fulfiller_FulfillerLinks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fulfiller_FulfillerLinks.Unmarshal(m, b)
}
func (m *Fulfiller_FulfillerLinks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fulfiller_FulfillerLinks.Marshal(b, m, deterministic)
}
func (m *Fulfiller_FulfillerLinks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fulfiller_FulfillerLinks.Merge(m, src)
}
func (m *Fulfiller_FulfillerLinks) XXX_Size() int {
	return xxx_messageInfo_Fulfiller_FulfillerLinks.Size(m)
}
func (m *Fulfiller_FulfillerLinks) XXX_DiscardUnknown() {
	xxx_messageInfo_Fulfiller_FulfillerLinks.DiscardUnknown(m)
}

var xxx_messageInfo_Fulfiller_FulfillerLinks proto.InternalMessageInfo

func (m *Fulfiller_FulfillerLinks) GetSelf() *Link {
	if m != nil {
		return m.Self
	}
	return nil
}

type NotificationItem struct {
	ItemId               string                                  `protobuf:"bytes,1,opt,name=itemId,proto3" json:"itemId,omitempty"`
	Status               string                                  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Links                *NotificationItem_NotificationItemLinks `protobuf:"bytes,3,opt,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *NotificationItem) Reset()         { *m = NotificationItem{} }
func (m *NotificationItem) String() string { return proto.CompactTextString(m) }
func (*NotificationItem) ProtoMessage()    {}
func (*NotificationItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{3}
}

func (m *NotificationItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationItem.Unmarshal(m, b)
}
func (m *NotificationItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationItem.Marshal(b, m, deterministic)
}
func (m *NotificationItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationItem.Merge(m, src)
}
func (m *NotificationItem) XXX_Size() int {
	return xxx_messageInfo_NotificationItem.Size(m)
}
func (m *NotificationItem) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationItem.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationItem proto.InternalMessageInfo

func (m *NotificationItem) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

func (m *NotificationItem) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *NotificationItem) GetLinks() *NotificationItem_NotificationItemLinks {
	if m != nil {
		return m.Links
	}
	return nil
}

type NotificationItem_NotificationItemLinks struct {
	Self                 *Link    `protobuf:"bytes,1,opt,name=self,proto3" json:"self,omitempty"`
	Accept               *Link    `protobuf:"bytes,2,opt,name=accept,proto3" json:"accept,omitempty"`
	Reject               *Link    `protobuf:"bytes,3,opt,name=reject,proto3" json:"reject,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationItem_NotificationItemLinks) Reset() {
	*m = NotificationItem_NotificationItemLinks{}
}
func (m *NotificationItem_NotificationItemLinks) String() string { return proto.CompactTextString(m) }
func (*NotificationItem_NotificationItemLinks) ProtoMessage()    {}
func (*NotificationItem_NotificationItemLinks) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{3, 0}
}

func (m *NotificationItem_NotificationItemLinks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationItem_NotificationItemLinks.Unmarshal(m, b)
}
func (m *NotificationItem_NotificationItemLinks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationItem_NotificationItemLinks.Marshal(b, m, deterministic)
}
func (m *NotificationItem_NotificationItemLinks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationItem_NotificationItemLinks.Merge(m, src)
}
func (m *NotificationItem_NotificationItemLinks) XXX_Size() int {
	return xxx_messageInfo_NotificationItem_NotificationItemLinks.Size(m)
}
func (m *NotificationItem_NotificationItemLinks) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationItem_NotificationItemLinks.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationItem_NotificationItemLinks proto.InternalMessageInfo

func (m *NotificationItem_NotificationItemLinks) GetSelf() *Link {
	if m != nil {
		return m.Self
	}
	return nil
}

func (m *NotificationItem_NotificationItemLinks) GetAccept() *Link {
	if m != nil {
		return m.Accept
	}
	return nil
}

func (m *NotificationItem_NotificationItemLinks) GetReject() *Link {
	if m != nil {
		return m.Reject
	}
	return nil
}

type ChangeRequest struct {
	DeliveryChangeDetails *ChangeRequest_DeliveryChange `protobuf:"bytes,1,opt,name=deliveryChangeDetails,proto3" json:"deliveryChangeDetails,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                      `json:"-"`
	XXX_unrecognized      []byte                        `json:"-"`
	XXX_sizecache         int32                         `json:"-"`
}

func (m *ChangeRequest) Reset()         { *m = ChangeRequest{} }
func (m *ChangeRequest) String() string { return proto.CompactTextString(m) }
func (*ChangeRequest) ProtoMessage()    {}
func (*ChangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{4}
}

func (m *ChangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeRequest.Unmarshal(m, b)
}
func (m *ChangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeRequest.Marshal(b, m, deterministic)
}
func (m *ChangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeRequest.Merge(m, src)
}
func (m *ChangeRequest) XXX_Size() int {
	return xxx_messageInfo_ChangeRequest.Size(m)
}
func (m *ChangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeRequest proto.InternalMessageInfo

func (m *ChangeRequest) GetDeliveryChangeDetails() *ChangeRequest_DeliveryChange {
	if m != nil {
		return m.DeliveryChangeDetails
	}
	return nil
}

type ChangeRequest_DeliveryChange struct {
	DestinationAddress   *ChangeRequest_DeliveryChange_AddressChange `protobuf:"bytes,1,opt,name=destinationAddress,proto3" json:"destinationAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *ChangeRequest_DeliveryChange) Reset()         { *m = ChangeRequest_DeliveryChange{} }
func (m *ChangeRequest_DeliveryChange) String() string { return proto.CompactTextString(m) }
func (*ChangeRequest_DeliveryChange) ProtoMessage()    {}
func (*ChangeRequest_DeliveryChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{4, 0}
}

func (m *ChangeRequest_DeliveryChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeRequest_DeliveryChange.Unmarshal(m, b)
}
func (m *ChangeRequest_DeliveryChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeRequest_DeliveryChange.Marshal(b, m, deterministic)
}
func (m *ChangeRequest_DeliveryChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeRequest_DeliveryChange.Merge(m, src)
}
func (m *ChangeRequest_DeliveryChange) XXX_Size() int {
	return xxx_messageInfo_ChangeRequest_DeliveryChange.Size(m)
}
func (m *ChangeRequest_DeliveryChange) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeRequest_DeliveryChange.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeRequest_DeliveryChange proto.InternalMessageInfo

func (m *ChangeRequest_DeliveryChange) GetDestinationAddress() *ChangeRequest_DeliveryChange_AddressChange {
	if m != nil {
		return m.DestinationAddress
	}
	return nil
}

type ChangeRequest_DeliveryChange_AddressChange struct {
	Country              string   `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	PostalCode           string   `protobuf:"bytes,2,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
	StateOrProvince      string   `protobuf:"bytes,3,opt,name=stateOrProvince,proto3" json:"stateOrProvince,omitempty"`
	City                 string   `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Company              string   `protobuf:"bytes,5,opt,name=company,proto3" json:"company,omitempty"`
	FirstName            string   `protobuf:"bytes,6,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,7,opt,name=lastName,proto3" json:"lastName,omitempty"`
	MiddleName           string   `protobuf:"bytes,8,opt,name=middleName,proto3" json:"middleName,omitempty"`
	Phone                string   `protobuf:"bytes,9,opt,name=phone,proto3" json:"phone,omitempty"`
	PhoneExt             string   `protobuf:"bytes,10,opt,name=phoneExt,proto3" json:"phoneExt,omitempty"`
	Street1              string   `protobuf:"bytes,11,opt,name=street1,proto3" json:"street1,omitempty"`
	Street2              string   `protobuf:"bytes,12,opt,name=street2,proto3" json:"street2,omitempty"`
	Email                string   `protobuf:"bytes,13,opt,name=email,proto3" json:"email,omitempty"`
	DoorCode             string   `protobuf:"bytes,14,opt,name=doorCode,proto3" json:"doorCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeRequest_DeliveryChange_AddressChange) Reset() {
	*m = ChangeRequest_DeliveryChange_AddressChange{}
}
func (m *ChangeRequest_DeliveryChange_AddressChange) String() string {
	return proto.CompactTextString(m)
}
func (*ChangeRequest_DeliveryChange_AddressChange) ProtoMessage() {}
func (*ChangeRequest_DeliveryChange_AddressChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{4, 0, 0}
}

func (m *ChangeRequest_DeliveryChange_AddressChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeRequest_DeliveryChange_AddressChange.Unmarshal(m, b)
}
func (m *ChangeRequest_DeliveryChange_AddressChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeRequest_DeliveryChange_AddressChange.Marshal(b, m, deterministic)
}
func (m *ChangeRequest_DeliveryChange_AddressChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeRequest_DeliveryChange_AddressChange.Merge(m, src)
}
func (m *ChangeRequest_DeliveryChange_AddressChange) XXX_Size() int {
	return xxx_messageInfo_ChangeRequest_DeliveryChange_AddressChange.Size(m)
}
func (m *ChangeRequest_DeliveryChange_AddressChange) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeRequest_DeliveryChange_AddressChange.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeRequest_DeliveryChange_AddressChange proto.InternalMessageInfo

func (m *ChangeRequest_DeliveryChange_AddressChange) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *ChangeRequest_DeliveryChange_AddressChange) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

func (m *ChangeRequest_DeliveryChange_AddressChange) GetStateOrProvince() string {
	if m != nil {
		return m.StateOrProvince
	}
	return ""
}

func (m *ChangeRequest_DeliveryChange_AddressChange) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *ChangeRequest_DeliveryChange_AddressChange) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *ChangeRequest_DeliveryChange_AddressChange) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *ChangeRequest_DeliveryChange_AddressChange) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *ChangeRequest_DeliveryChange_AddressChange) GetMiddleName() string {
	if m != nil {
		return m.MiddleName
	}
	return ""
}

func (m *ChangeRequest_DeliveryChange_AddressChange) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *ChangeRequest_DeliveryChange_AddressChange) GetPhoneExt() string {
	if m != nil {
		return m.PhoneExt
	}
	return ""
}

func (m *ChangeRequest_DeliveryChange_AddressChange) GetStreet1() string {
	if m != nil {
		return m.Street1
	}
	return ""
}

func (m *ChangeRequest_DeliveryChange_AddressChange) GetStreet2() string {
	if m != nil {
		return m.Street2
	}
	return ""
}

func (m *ChangeRequest_DeliveryChange_AddressChange) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ChangeRequest_DeliveryChange_AddressChange) GetDoorCode() string {
	if m != nil {
		return m.DoorCode
	}
	return ""
}

type RetryDetails struct {
	Reason               string   `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetryDetails) Reset()         { *m = RetryDetails{} }
func (m *RetryDetails) String() string { return proto.CompactTextString(m) }
func (*RetryDetails) ProtoMessage()    {}
func (*RetryDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{5}
}

func (m *RetryDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryDetails.Unmarshal(m, b)
}
func (m *RetryDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryDetails.Marshal(b, m, deterministic)
}
func (m *RetryDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryDetails.Merge(m, src)
}
func (m *RetryDetails) XXX_Size() int {
	return xxx_messageInfo_RetryDetails.Size(m)
}
func (m *RetryDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryDetails.DiscardUnknown(m)
}

var xxx_messageInfo_RetryDetails proto.InternalMessageInfo

func (m *RetryDetails) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *RetryDetails) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Notification struct {
	NotificationId       string                          `protobuf:"bytes,1,opt,name=notificationId,proto3" json:"notificationId,omitempty"`
	Type                 string                          `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Status               string                          `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	CreatedDate          string                          `protobuf:"bytes,4,opt,name=createdDate,proto3" json:"createdDate,omitempty"`
	Order                *Order                          `protobuf:"bytes,5,opt,name=order,proto3" json:"order,omitempty"`
	Fulfiller            *Fulfiller                      `protobuf:"bytes,6,opt,name=fulfiller,proto3" json:"fulfiller,omitempty"`
	Items                []*NotificationItem             `protobuf:"bytes,7,rep,name=items,proto3" json:"items,omitempty"`
	Links                *Notification_NotificationLinks `protobuf:"bytes,8,opt,name=links,proto3" json:"links,omitempty"`
	ChangeRequest        *ChangeRequest                  `protobuf:"bytes,9,opt,name=changeRequest,proto3" json:"changeRequest,omitempty"`
	RetryOrderRequest    *RetryDetails                   `protobuf:"bytes,10,opt,name=retryOrderRequest,proto3" json:"retryOrderRequest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{6}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetNotificationId() string {
	if m != nil {
		return m.NotificationId
	}
	return ""
}

func (m *Notification) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Notification) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Notification) GetCreatedDate() string {
	if m != nil {
		return m.CreatedDate
	}
	return ""
}

func (m *Notification) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *Notification) GetFulfiller() *Fulfiller {
	if m != nil {
		return m.Fulfiller
	}
	return nil
}

func (m *Notification) GetItems() []*NotificationItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Notification) GetLinks() *Notification_NotificationLinks {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *Notification) GetChangeRequest() *ChangeRequest {
	if m != nil {
		return m.ChangeRequest
	}
	return nil
}

func (m *Notification) GetRetryOrderRequest() *RetryDetails {
	if m != nil {
		return m.RetryOrderRequest
	}
	return nil
}

type Notification_NotificationLinks struct {
	Self                 *Link    `protobuf:"bytes,1,opt,name=self,proto3" json:"self,omitempty"`
	Accept               *Link    `protobuf:"bytes,2,opt,name=accept,proto3" json:"accept,omitempty"`
	Reject               *Link    `protobuf:"bytes,3,opt,name=reject,proto3" json:"reject,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notification_NotificationLinks) Reset()         { *m = Notification_NotificationLinks{} }
func (m *Notification_NotificationLinks) String() string { return proto.CompactTextString(m) }
func (*Notification_NotificationLinks) ProtoMessage()    {}
func (*Notification_NotificationLinks) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{6, 0}
}

func (m *Notification_NotificationLinks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification_NotificationLinks.Unmarshal(m, b)
}
func (m *Notification_NotificationLinks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification_NotificationLinks.Marshal(b, m, deterministic)
}
func (m *Notification_NotificationLinks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification_NotificationLinks.Merge(m, src)
}
func (m *Notification_NotificationLinks) XXX_Size() int {
	return xxx_messageInfo_Notification_NotificationLinks.Size(m)
}
func (m *Notification_NotificationLinks) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification_NotificationLinks.DiscardUnknown(m)
}

var xxx_messageInfo_Notification_NotificationLinks proto.InternalMessageInfo

func (m *Notification_NotificationLinks) GetSelf() *Link {
	if m != nil {
		return m.Self
	}
	return nil
}

func (m *Notification_NotificationLinks) GetAccept() *Link {
	if m != nil {
		return m.Accept
	}
	return nil
}

func (m *Notification_NotificationLinks) GetReject() *Link {
	if m != nil {
		return m.Reject
	}
	return nil
}

type Notifications struct {
	Notifications        []*Notification                  `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
	Count                int32                            `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Links                *Notifications_NotificationLinks `protobuf:"bytes,3,opt,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *Notifications) Reset()         { *m = Notifications{} }
func (m *Notifications) String() string { return proto.CompactTextString(m) }
func (*Notifications) ProtoMessage()    {}
func (*Notifications) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{7}
}

func (m *Notifications) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notifications.Unmarshal(m, b)
}
func (m *Notifications) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notifications.Marshal(b, m, deterministic)
}
func (m *Notifications) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notifications.Merge(m, src)
}
func (m *Notifications) XXX_Size() int {
	return xxx_messageInfo_Notifications.Size(m)
}
func (m *Notifications) XXX_DiscardUnknown() {
	xxx_messageInfo_Notifications.DiscardUnknown(m)
}

var xxx_messageInfo_Notifications proto.InternalMessageInfo

func (m *Notifications) GetNotifications() []*Notification {
	if m != nil {
		return m.Notifications
	}
	return nil
}

func (m *Notifications) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Notifications) GetLinks() *Notifications_NotificationLinks {
	if m != nil {
		return m.Links
	}
	return nil
}

type Notifications_NotificationLinks struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notifications_NotificationLinks) Reset()         { *m = Notifications_NotificationLinks{} }
func (m *Notifications_NotificationLinks) String() string { return proto.CompactTextString(m) }
func (*Notifications_NotificationLinks) ProtoMessage()    {}
func (*Notifications_NotificationLinks) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{7, 0}
}

func (m *Notifications_NotificationLinks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notifications_NotificationLinks.Unmarshal(m, b)
}
func (m *Notifications_NotificationLinks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notifications_NotificationLinks.Marshal(b, m, deterministic)
}
func (m *Notifications_NotificationLinks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notifications_NotificationLinks.Merge(m, src)
}
func (m *Notifications_NotificationLinks) XXX_Size() int {
	return xxx_messageInfo_Notifications_NotificationLinks.Size(m)
}
func (m *Notifications_NotificationLinks) XXX_DiscardUnknown() {
	xxx_messageInfo_Notifications_NotificationLinks.DiscardUnknown(m)
}

var xxx_messageInfo_Notifications_NotificationLinks proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Link)(nil), "pb.Link")
	proto.RegisterType((*Order)(nil), "pb.Order")
	proto.RegisterType((*Order_OrderLinks)(nil), "pb.Order.OrderLinks")
	proto.RegisterType((*Fulfiller)(nil), "pb.Fulfiller")
	proto.RegisterType((*Fulfiller_FulfillerLinks)(nil), "pb.Fulfiller.FulfillerLinks")
	proto.RegisterType((*NotificationItem)(nil), "pb.NotificationItem")
	proto.RegisterType((*NotificationItem_NotificationItemLinks)(nil), "pb.NotificationItem.NotificationItemLinks")
	proto.RegisterType((*ChangeRequest)(nil), "pb.ChangeRequest")
	proto.RegisterType((*ChangeRequest_DeliveryChange)(nil), "pb.ChangeRequest.DeliveryChange")
	proto.RegisterType((*ChangeRequest_DeliveryChange_AddressChange)(nil), "pb.ChangeRequest.DeliveryChange.AddressChange")
	proto.RegisterType((*RetryDetails)(nil), "pb.RetryDetails")
	proto.RegisterType((*Notification)(nil), "pb.Notification")
	proto.RegisterType((*Notification_NotificationLinks)(nil), "pb.Notification.NotificationLinks")
	proto.RegisterType((*Notifications)(nil), "pb.Notifications")
	proto.RegisterType((*Notifications_NotificationLinks)(nil), "pb.Notifications.NotificationLinks")
}

func init() { proto.RegisterFile("notification.proto", fileDescriptor_736a457d4a5efa07) }

var fileDescriptor_736a457d4a5efa07 = []byte{
	// 781 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xdd, 0x6a, 0xdb, 0x48,
	0x14, 0xc6, 0x91, 0xed, 0x58, 0xc7, 0xb1, 0x37, 0x99, 0x4d, 0x16, 0x61, 0xc2, 0xae, 0xf1, 0xc2,
	0x62, 0xb2, 0x60, 0xa8, 0x0a, 0xfd, 0xb9, 0x29, 0x29, 0x49, 0x03, 0x81, 0x92, 0x14, 0x5d, 0xf4,
	0xb2, 0xa0, 0x48, 0xc7, 0xf1, 0x34, 0xb2, 0x46, 0x9d, 0x19, 0x87, 0xea, 0x05, 0xfa, 0x02, 0x7d,
	0x89, 0xde, 0xf4, 0xdd, 0xfa, 0x06, 0x2d, 0xf3, 0x67, 0x4b, 0xb6, 0xdb, 0xdc, 0xf5, 0xc6, 0xcc,
	0x77, 0xce, 0x77, 0x7e, 0x46, 0x67, 0xce, 0x67, 0x20, 0x39, 0x93, 0x74, 0x4a, 0x93, 0x58, 0x52,
	0x96, 0x4f, 0x0a, 0xce, 0x24, 0x23, 0x3b, 0xc5, 0xcd, 0xe8, 0x02, 0x9a, 0xaf, 0x69, 0x7e, 0x47,
	0x08, 0x34, 0x67, 0x1c, 0xa7, 0x41, 0x63, 0xd8, 0x18, 0xfb, 0x91, 0x3e, 0x93, 0x00, 0x76, 0xe7,
	0x28, 0x67, 0x2c, 0x15, 0xc1, 0xce, 0xd0, 0x1b, 0xfb, 0x91, 0x83, 0x64, 0x1f, 0x3c, 0x8e, 0x59,
	0xe0, 0x69, 0xb2, 0x3a, 0x8e, 0x4a, 0x68, 0x5d, 0xf3, 0x14, 0xb9, 0x0a, 0x62, 0xea, 0x70, 0x99,
	0xda, 0x5c, 0x0e, 0x92, 0x13, 0x68, 0x65, 0x34, 0xbf, 0x53, 0xc9, 0x1a, 0xe3, 0x6e, 0x78, 0x38,
	0x29, 0x6e, 0x26, 0x3a, 0xc6, 0xfc, 0xaa, 0x36, 0x44, 0x64, 0x28, 0x83, 0x13, 0x80, 0x95, 0x91,
	0x1c, 0x43, 0x53, 0x60, 0x66, 0x9a, 0xeb, 0x86, 0x1d, 0x15, 0xa8, 0x1c, 0x91, 0xb6, 0x8e, 0x18,
	0xf8, 0x17, 0x8b, 0x6c, 0x4a, 0xb3, 0x0c, 0x39, 0x09, 0x5d, 0x11, 0xc3, 0x3d, 0x56, 0xdc, 0xa5,
	0x77, 0x75, 0xaa, 0x15, 0x9b, 0x40, 0xbf, 0xee, 0x78, 0xa0, 0xe0, 0xf7, 0x06, 0xec, 0x5f, 0x55,
	0x3e, 0xe7, 0xa5, 0xc4, 0x39, 0xf9, 0x0b, 0xda, 0x54, 0xe2, 0x7c, 0x79, 0x6d, 0x8b, 0x94, 0x5d,
	0xc8, 0x58, 0x2e, 0xcc, 0xb5, 0xfd, 0xc8, 0x22, 0x72, 0xea, 0x1a, 0xf5, 0x74, 0x8d, 0x13, 0x55,
	0x63, 0x3d, 0xe9, 0x86, 0xa1, 0xd6, 0x76, 0x09, 0x47, 0x5b, 0xfd, 0xbf, 0xee, 0x9e, 0x0c, 0xa1,
	0x1d, 0x27, 0x09, 0x16, 0xd2, 0xce, 0x61, 0xe5, 0xb7, 0x76, 0xc5, 0xe0, 0xf8, 0x1e, 0x13, 0x69,
	0x7b, 0xab, 0x30, 0x8c, 0x7d, 0xf4, 0xad, 0x09, 0xbd, 0xb3, 0x59, 0x9c, 0xdf, 0x62, 0x84, 0x1f,
	0x16, 0x28, 0x24, 0x79, 0x0b, 0x47, 0x29, 0x66, 0xf4, 0x1e, 0x79, 0x69, 0x1c, 0xe7, 0x28, 0x63,
	0x9a, 0xb9, 0x39, 0x0c, 0x55, 0x8a, 0x5a, 0xc4, 0xe4, 0xbc, 0x46, 0x8f, 0xb6, 0x87, 0x0f, 0x3e,
	0x35, 0xa1, 0x5f, 0x67, 0x92, 0x77, 0x40, 0x52, 0x14, 0x92, 0xe6, 0xfa, 0xda, 0x2f, 0xd3, 0x94,
	0xa3, 0x70, 0x75, 0x26, 0x0f, 0xd5, 0x99, 0x58, 0xbe, 0xe5, 0x6c, 0xc9, 0x34, 0xf8, 0xec, 0x41,
	0xaf, 0xc6, 0x52, 0x6f, 0x3a, 0x61, 0x8b, 0x5c, 0xf2, 0xd2, 0xbd, 0x69, 0x0b, 0xc9, 0xdf, 0x00,
	0x05, 0x13, 0x32, 0xce, 0xce, 0x58, 0x8a, 0x76, 0xc2, 0x15, 0x0b, 0x19, 0xc3, 0x1f, 0x6a, 0xde,
	0x78, 0xcd, 0xdf, 0x70, 0x76, 0x4f, 0xf3, 0x04, 0xed, 0xd2, 0xac, 0x9b, 0xd5, 0x02, 0x26, 0x54,
	0x96, 0x41, 0xd3, 0x2c, 0xa0, 0x3a, 0x9b, 0xba, 0xf3, 0x22, 0xce, 0xcb, 0xa0, 0xe5, 0xea, 0x6a,
	0x48, 0x8e, 0xc1, 0x9f, 0x52, 0x2e, 0xe4, 0x55, 0x3c, 0xc7, 0xa0, 0xad, 0x7d, 0x2b, 0x03, 0x19,
	0x40, 0x27, 0x8b, 0xad, 0x73, 0x57, 0x3b, 0x97, 0x58, 0x75, 0x3c, 0xa7, 0x69, 0x9a, 0xa1, 0xf6,
	0x76, 0x4c, 0xc7, 0x2b, 0x0b, 0x39, 0x84, 0x56, 0x31, 0x63, 0x39, 0x06, 0xbe, 0x76, 0x19, 0xa0,
	0x32, 0xea, 0xc3, 0xab, 0x8f, 0x32, 0x00, 0x93, 0xd1, 0x61, 0xd5, 0xa5, 0x90, 0x1c, 0x51, 0x3e,
	0x0a, 0xba, 0xa6, 0x4b, 0x0b, 0x57, 0x9e, 0x30, 0xd8, 0xab, 0x7a, 0x42, 0x55, 0x05, 0xe7, 0x31,
	0xcd, 0x82, 0x9e, 0xa9, 0xa2, 0x81, 0xaa, 0x92, 0x32, 0xc6, 0xf5, 0xb7, 0xec, 0x9b, 0x2a, 0x0e,
	0x8f, 0x4e, 0x61, 0x2f, 0x42, 0xc9, 0x4b, 0xfb, 0x30, 0xd4, 0x5e, 0x71, 0x8c, 0x05, 0xcb, 0xdd,
	0xbe, 0x19, 0x64, 0x44, 0x4b, 0x88, 0xf8, 0xd6, 0x8d, 0xc3, 0xc1, 0xd1, 0x97, 0x26, 0xec, 0x55,
	0x17, 0x86, 0xfc, 0x07, 0xfd, 0xaa, 0x2a, 0x2e, 0x57, 0x77, 0xcd, 0xaa, 0x46, 0x23, 0xcb, 0xc2,
	0xe5, 0xd3, 0xe7, 0xca, 0x5a, 0x7b, 0xb5, 0xb5, 0x1e, 0x42, 0x37, 0xe1, 0x18, 0x4b, 0x4c, 0xcf,
	0x63, 0x89, 0x76, 0x9a, 0x55, 0x13, 0xf9, 0x07, 0x5a, 0x5a, 0x11, 0xf5, 0x48, 0xbb, 0xa1, 0xbf,
	0x94, 0xc1, 0xc8, 0xd8, 0xc9, 0xff, 0xe0, 0x4f, 0x9d, 0x1c, 0xe9, 0xd9, 0x76, 0xc3, 0x5e, 0x4d,
	0xc6, 0xa2, 0x95, 0x5f, 0x89, 0xaa, 0x12, 0x1a, 0x11, 0xec, 0x0e, 0x3d, 0x27, 0xaa, 0xeb, 0xaa,
	0x10, 0x19, 0x0a, 0x79, 0xe6, 0x24, 0xa7, 0xa3, 0x93, 0x8e, 0xd6, 0xb9, 0x35, 0x50, 0x95, 0x1a,
	0xf2, 0x14, 0x7a, 0x49, 0x75, 0xa9, 0xf4, 0xe3, 0xe8, 0x86, 0x07, 0x1b, 0xdb, 0x16, 0xd5, 0x79,
	0xe4, 0x05, 0x1c, 0x70, 0x35, 0x35, 0x73, 0x41, 0x1b, 0x0c, 0x3a, 0x78, 0x5f, 0x05, 0x57, 0x47,
	0x1a, 0x6d, 0x52, 0x07, 0x0b, 0x38, 0xd8, 0x68, 0xea, 0x37, 0xe8, 0xdb, 0xd7, 0x06, 0xf4, 0xaa,
	0x75, 0x05, 0x79, 0x02, 0xbd, 0xea, 0xab, 0x50, 0x7a, 0xe3, 0xb9, 0x4b, 0x54, 0x99, 0x51, 0x9d,
	0xa6, 0x1e, 0xba, 0xd6, 0x0a, 0xdd, 0x4c, 0x2b, 0x32, 0x80, 0x3c, 0xaf, 0x8b, 0xff, 0xbf, 0xeb,
	0x59, 0xc4, 0x4f, 0x47, 0x31, 0xf8, 0x73, 0xcb, 0x17, 0xb9, 0x69, 0xeb, 0x3f, 0xf4, 0xc7, 0x3f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x70, 0x9a, 0x39, 0x1a, 0xe6, 0x07, 0x00, 0x00,
}
