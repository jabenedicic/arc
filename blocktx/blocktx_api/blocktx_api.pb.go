// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: blocktx/blocktx_api/blocktx_api.proto

package blocktx_api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// swagger:model HealthResponse
type HealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok        bool                   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Details   string                 `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *HealthResponse) Reset() {
	*x = HealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResponse.ProtoReflect.Descriptor instead.
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return file_blocktx_blocktx_api_blocktx_api_proto_rawDescGZIP(), []int{0}
}

func (x *HealthResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *HealthResponse) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *HealthResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

// swagger:model Block {
type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash         []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`                                     // Little endian
	PreviousHash []byte `protobuf:"bytes,2,opt,name=previous_hash,json=previousHash,proto3" json:"previous_hash,omitempty"` // Little endian
	MerkleRoot   []byte `protobuf:"bytes,3,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`       // Little endian
	Height       uint64 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	Orphaned     bool   `protobuf:"varint,5,opt,name=orphaned,proto3" json:"orphaned,omitempty"`
	Processed    bool   `protobuf:"varint,6,opt,name=processed,proto3" json:"processed,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_blocktx_blocktx_api_blocktx_api_proto_rawDescGZIP(), []int{1}
}

func (x *Block) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Block) GetPreviousHash() []byte {
	if x != nil {
		return x.PreviousHash
	}
	return nil
}

func (x *Block) GetMerkleRoot() []byte {
	if x != nil {
		return x.MerkleRoot
	}
	return nil
}

func (x *Block) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Block) GetOrphaned() bool {
	if x != nil {
		return x.Orphaned
	}
	return false
}

func (x *Block) GetProcessed() bool {
	if x != nil {
		return x.Processed
	}
	return false
}

// swagger:model Transactions
type Transactions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *Transactions) Reset() {
	*x = Transactions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transactions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transactions) ProtoMessage() {}

func (x *Transactions) ProtoReflect() protoreflect.Message {
	mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transactions.ProtoReflect.Descriptor instead.
func (*Transactions) Descriptor() ([]byte, []int) {
	return file_blocktx_blocktx_api_blocktx_api_proto_rawDescGZIP(), []int{2}
}

func (x *Transactions) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type TransactionBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHash       []byte `protobuf:"bytes,1,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"` // Little endian
	BlockHeight     uint64 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	TransactionHash []byte `protobuf:"bytes,3,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"` // Little endian
}

func (x *TransactionBlock) Reset() {
	*x = TransactionBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionBlock) ProtoMessage() {}

func (x *TransactionBlock) ProtoReflect() protoreflect.Message {
	mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionBlock.ProtoReflect.Descriptor instead.
func (*TransactionBlock) Descriptor() ([]byte, []int) {
	return file_blocktx_blocktx_api_blocktx_api_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionBlock) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *TransactionBlock) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *TransactionBlock) GetTransactionHash() []byte {
	if x != nil {
		return x.TransactionHash
	}
	return nil
}

type TransactionBlocks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionBlocks []*TransactionBlock `protobuf:"bytes,1,rep,name=transaction_blocks,json=transactionBlocks,proto3" json:"transaction_blocks,omitempty"`
}

func (x *TransactionBlocks) Reset() {
	*x = TransactionBlocks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionBlocks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionBlocks) ProtoMessage() {}

func (x *TransactionBlocks) ProtoReflect() protoreflect.Message {
	mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionBlocks.ProtoReflect.Descriptor instead.
func (*TransactionBlocks) Descriptor() ([]byte, []int) {
	return file_blocktx_blocktx_api_blocktx_api_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionBlocks) GetTransactionBlocks() []*TransactionBlock {
	if x != nil {
		return x.TransactionBlocks
	}
	return nil
}

// swagger:model MinedTransactions
type MinedTransactions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block        *Block         `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`               // Little endian
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"` // Little endian
}

func (x *MinedTransactions) Reset() {
	*x = MinedTransactions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MinedTransactions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MinedTransactions) ProtoMessage() {}

func (x *MinedTransactions) ProtoReflect() protoreflect.Message {
	mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MinedTransactions.ProtoReflect.Descriptor instead.
func (*MinedTransactions) Descriptor() ([]byte, []int) {
	return file_blocktx_blocktx_api_blocktx_api_proto_rawDescGZIP(), []int{5}
}

func (x *MinedTransactions) GetBlock() *Block {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *MinedTransactions) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

// swagger:model Transaction
type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash   []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`     // Little endian
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"` // This is the metamorph address:port
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_blocktx_blocktx_api_blocktx_api_proto_rawDescGZIP(), []int{6}
}

func (x *Transaction) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Transaction) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

// swagger:model Height
type Height struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *Height) Reset() {
	*x = Height{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Height) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Height) ProtoMessage() {}

func (x *Height) ProtoReflect() protoreflect.Message {
	mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Height.ProtoReflect.Descriptor instead.
func (*Height) Descriptor() ([]byte, []int) {
	return file_blocktx_blocktx_api_blocktx_api_proto_rawDescGZIP(), []int{7}
}

func (x *Height) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type Hash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *Hash) Reset() {
	*x = Hash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hash) ProtoMessage() {}

func (x *Hash) ProtoReflect() protoreflect.Message {
	mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hash.ProtoReflect.Descriptor instead.
func (*Hash) Descriptor() ([]byte, []int) {
	return file_blocktx_blocktx_api_blocktx_api_proto_rawDescGZIP(), []int{8}
}

func (x *Hash) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type MerklePath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MerklePath string `protobuf:"bytes,1,opt,name=merklePath,proto3" json:"merklePath,omitempty"`
}

func (x *MerklePath) Reset() {
	*x = MerklePath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerklePath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerklePath) ProtoMessage() {}

func (x *MerklePath) ProtoReflect() protoreflect.Message {
	mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerklePath.ProtoReflect.Descriptor instead.
func (*MerklePath) Descriptor() ([]byte, []int) {
	return file_blocktx_blocktx_api_blocktx_api_proto_rawDescGZIP(), []int{9}
}

func (x *MerklePath) GetMerklePath() string {
	if x != nil {
		return x.MerklePath
	}
	return ""
}

type TransactionAndSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash   []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *TransactionAndSource) Reset() {
	*x = TransactionAndSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionAndSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionAndSource) ProtoMessage() {}

func (x *TransactionAndSource) ProtoReflect() protoreflect.Message {
	mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionAndSource.ProtoReflect.Descriptor instead.
func (*TransactionAndSource) Descriptor() ([]byte, []int) {
	return file_blocktx_blocktx_api_blocktx_api_proto_rawDescGZIP(), []int{10}
}

func (x *TransactionAndSource) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *TransactionAndSource) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

// swagger:model BlockAndSource
type BlockAndSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash   []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *BlockAndSource) Reset() {
	*x = BlockAndSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockAndSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockAndSource) ProtoMessage() {}

func (x *BlockAndSource) ProtoReflect() protoreflect.Message {
	mi := &file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockAndSource.ProtoReflect.Descriptor instead.
func (*BlockAndSource) Descriptor() ([]byte, []int) {
	return file_blocktx_blocktx_api_blocktx_api_proto_rawDescGZIP(), []int{11}
}

func (x *BlockAndSource) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *BlockAndSource) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

var File_blocktx_blocktx_api_blocktx_api_proto protoreflect.FileDescriptor

var file_blocktx_blocktx_api_blocktx_api_proto_rawDesc = []byte{
	0x0a, 0x25, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x78, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x74,
	0x78, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x78, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x78,
	0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x74, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x02, 0x6f, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x38,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xb3, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x6d,
	0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x70, 0x68, 0x61, 0x6e, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6f, 0x72, 0x70, 0x68, 0x61, 0x6e, 0x65, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x22, 0x4c,
	0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3c,
	0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x78, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x7f, 0x0a, 0x10,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x22, 0x61, 0x0a,
	0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x12, 0x4c, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x78, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x11, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x22, 0x7b, 0x0a, 0x11, 0x4d, 0x69, 0x6e, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x78, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x3c, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x78, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x39, 0x0a,
	0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x20, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x1a, 0x0a, 0x04, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x2c, 0x0a, 0x0a, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x22, 0x42, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x3c, 0x0a, 0x0e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x41, 0x6e, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x32, 0xc7, 0x02, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x54, 0x78, 0x41, 0x50, 0x49, 0x12, 0x3f, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x74,
	0x78, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x78, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x72, 0x6b,
	0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x18, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x78,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x78, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4d,
	0x65, 0x72, 0x6b, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x73, 0x12, 0x19, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x78, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1e,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x78, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x00,
	0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x3b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x78, 0x5f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blocktx_blocktx_api_blocktx_api_proto_rawDescOnce sync.Once
	file_blocktx_blocktx_api_blocktx_api_proto_rawDescData = file_blocktx_blocktx_api_blocktx_api_proto_rawDesc
)

func file_blocktx_blocktx_api_blocktx_api_proto_rawDescGZIP() []byte {
	file_blocktx_blocktx_api_blocktx_api_proto_rawDescOnce.Do(func() {
		file_blocktx_blocktx_api_blocktx_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_blocktx_blocktx_api_blocktx_api_proto_rawDescData)
	})
	return file_blocktx_blocktx_api_blocktx_api_proto_rawDescData
}

var file_blocktx_blocktx_api_blocktx_api_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_blocktx_blocktx_api_blocktx_api_proto_goTypes = []interface{}{
	(*HealthResponse)(nil),        // 0: blocktx_api.HealthResponse
	(*Block)(nil),                 // 1: blocktx_api.Block
	(*Transactions)(nil),          // 2: blocktx_api.Transactions
	(*TransactionBlock)(nil),      // 3: blocktx_api.TransactionBlock
	(*TransactionBlocks)(nil),     // 4: blocktx_api.TransactionBlocks
	(*MinedTransactions)(nil),     // 5: blocktx_api.MinedTransactions
	(*Transaction)(nil),           // 6: blocktx_api.Transaction
	(*Height)(nil),                // 7: blocktx_api.Height
	(*Hash)(nil),                  // 8: blocktx_api.Hash
	(*MerklePath)(nil),            // 9: blocktx_api.MerklePath
	(*TransactionAndSource)(nil),  // 10: blocktx_api.TransactionAndSource
	(*BlockAndSource)(nil),        // 11: blocktx_api.BlockAndSource
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 13: google.protobuf.Empty
}
var file_blocktx_blocktx_api_blocktx_api_proto_depIdxs = []int32{
	12, // 0: blocktx_api.HealthResponse.timestamp:type_name -> google.protobuf.Timestamp
	6,  // 1: blocktx_api.Transactions.transactions:type_name -> blocktx_api.Transaction
	3,  // 2: blocktx_api.TransactionBlocks.transaction_blocks:type_name -> blocktx_api.TransactionBlock
	1,  // 3: blocktx_api.MinedTransactions.block:type_name -> blocktx_api.Block
	6,  // 4: blocktx_api.MinedTransactions.transactions:type_name -> blocktx_api.Transaction
	13, // 5: blocktx_api.BlockTxAPI.Health:input_type -> google.protobuf.Empty
	10, // 6: blocktx_api.BlockTxAPI.RegisterTransaction:input_type -> blocktx_api.TransactionAndSource
	6,  // 7: blocktx_api.BlockTxAPI.GetTransactionMerklePath:input_type -> blocktx_api.Transaction
	2,  // 8: blocktx_api.BlockTxAPI.GetTransactionBlocks:input_type -> blocktx_api.Transactions
	0,  // 9: blocktx_api.BlockTxAPI.Health:output_type -> blocktx_api.HealthResponse
	13, // 10: blocktx_api.BlockTxAPI.RegisterTransaction:output_type -> google.protobuf.Empty
	9,  // 11: blocktx_api.BlockTxAPI.GetTransactionMerklePath:output_type -> blocktx_api.MerklePath
	4,  // 12: blocktx_api.BlockTxAPI.GetTransactionBlocks:output_type -> blocktx_api.TransactionBlocks
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_blocktx_blocktx_api_blocktx_api_proto_init() }
func file_blocktx_blocktx_api_blocktx_api_proto_init() {
	if File_blocktx_blocktx_api_blocktx_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthResponse); i {
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
		file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transactions); i {
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
		file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionBlock); i {
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
		file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionBlocks); i {
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
		file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MinedTransactions); i {
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
		file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Height); i {
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
		file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hash); i {
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
		file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerklePath); i {
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
		file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionAndSource); i {
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
		file_blocktx_blocktx_api_blocktx_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockAndSource); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blocktx_blocktx_api_blocktx_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blocktx_blocktx_api_blocktx_api_proto_goTypes,
		DependencyIndexes: file_blocktx_blocktx_api_blocktx_api_proto_depIdxs,
		MessageInfos:      file_blocktx_blocktx_api_blocktx_api_proto_msgTypes,
	}.Build()
	File_blocktx_blocktx_api_blocktx_api_proto = out.File
	file_blocktx_blocktx_api_blocktx_api_proto_rawDesc = nil
	file_blocktx_blocktx_api_blocktx_api_proto_goTypes = nil
	file_blocktx_blocktx_api_blocktx_api_proto_depIdxs = nil
}
