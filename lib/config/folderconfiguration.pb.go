// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/config/folderconfiguration.proto

package config

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	fs "github.com/syncthing/syncthing/lib/fs"
	github_com_syncthing_syncthing_lib_protocol "github.com/syncthing/syncthing/lib/protocol"
	_ "github.com/syncthing/syncthing/proto/ext"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type FolderDeviceConfiguration struct {
	DeviceID     github_com_syncthing_syncthing_lib_protocol.DeviceID `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3,customtype=github.com/syncthing/syncthing/lib/protocol.DeviceID" json:"deviceID" xml:"id,attr"`
	IntroducedBy github_com_syncthing_syncthing_lib_protocol.DeviceID `protobuf:"bytes,2,opt,name=introduced_by,json=introducedBy,proto3,customtype=github.com/syncthing/syncthing/lib/protocol.DeviceID" json:"introducedBy" xml:"introducedBy,attr"`
}

func (m *FolderDeviceConfiguration) Reset()         { *m = FolderDeviceConfiguration{} }
func (m *FolderDeviceConfiguration) String() string { return proto.CompactTextString(m) }
func (*FolderDeviceConfiguration) ProtoMessage()    {}
func (*FolderDeviceConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_44a9785876ed3afa, []int{0}
}
func (m *FolderDeviceConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FolderDeviceConfiguration.Unmarshal(m, b)
}
func (m *FolderDeviceConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FolderDeviceConfiguration.Marshal(b, m, deterministic)
}
func (m *FolderDeviceConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FolderDeviceConfiguration.Merge(m, src)
}
func (m *FolderDeviceConfiguration) XXX_Size() int {
	return xxx_messageInfo_FolderDeviceConfiguration.Size(m)
}
func (m *FolderDeviceConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_FolderDeviceConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_FolderDeviceConfiguration proto.InternalMessageInfo

type FolderConfiguration struct {
	ID                      string                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id" xml:"id,attr"`
	Label                   string                      `protobuf:"bytes,2,opt,name=label,proto3" json:"label" xml:"label,attr" restart:"false"`
	FilesystemType          fs.FilesystemType           `protobuf:"varint,3,opt,name=filesystem_type,json=filesystemType,proto3,enum=fs.FilesystemType" json:"filesystemType" xml:"filesystemType"`
	Path                    string                      `protobuf:"bytes,4,opt,name=path,proto3" json:"path" xml:"path,attr"`
	Type                    FolderType                  `protobuf:"varint,5,opt,name=type,proto3,enum=config.FolderType" json:"type" xml:"type,attr"`
	Devices                 []FolderDeviceConfiguration `protobuf:"bytes,6,rep,name=devices,proto3" json:"devices" xml:"device"`
	RescanIntervalS         int                         `protobuf:"varint,7,opt,name=rescan_interval_s,json=rescanIntervalS,proto3,casttype=int" json:"rescanIntervalS" xml:"rescanIntervalS,attr" default:"3600"`
	FSWatcherEnabled        bool                        `protobuf:"varint,8,opt,name=fs_watcher_enabled,json=fsWatcherEnabled,proto3" json:"fsWatcherEnabled" xml:"fsWatcherEnabled,attr" default:"true"`
	FSWatcherDelayS         int                         `protobuf:"varint,9,opt,name=fs_watcher_delay_s,json=fsWatcherDelayS,proto3,casttype=int" json:"fsWatcherDelayS" xml:"fsWatcherDelayS,attr" default:"10"`
	IgnorePerms             bool                        `protobuf:"varint,10,opt,name=ignore_perms,json=ignorePerms,proto3" json:"ignorePerms" xml:"ignorePerms,attr"`
	AutoNormalize           bool                        `protobuf:"varint,11,opt,name=auto_normalize,json=autoNormalize,proto3" json:"autoNormalize" xml:"autoNormalize,attr" default:"true"`
	MinDiskFree             Size                        `protobuf:"bytes,12,opt,name=min_disk_free,json=minDiskFree,proto3" json:"minDiskFree" xml:"minDiskFree"`
	Versioning              VersioningConfiguration     `protobuf:"bytes,13,opt,name=versioning,proto3" json:"versioning" xml:"versioning"`
	Copiers                 int                         `protobuf:"varint,14,opt,name=copiers,proto3,casttype=int" json:"copiers" xml:"copiers"`
	PullerMaxPendingKiB     int                         `protobuf:"varint,15,opt,name=puller_max_pending_kib,json=pullerMaxPendingKib,proto3,casttype=int" json:"pullerMaxPendingKiB" xml:"pullerMaxPendingKiB"`
	Hashers                 int                         `protobuf:"varint,16,opt,name=hashers,proto3,casttype=int" json:"hashers" xml:"hashers"`
	Order                   PullOrder                   `protobuf:"varint,17,opt,name=order,proto3,enum=config.PullOrder" json:"order" xml:"order"`
	IgnoreDelete            bool                        `protobuf:"varint,18,opt,name=ignore_delete,json=ignoreDelete,proto3" json:"ignoreDelete" xml:"ignoreDelete"`
	ScanProgressIntervalS   int                         `protobuf:"varint,19,opt,name=scan_progress_interval_s,json=scanProgressIntervalS,proto3,casttype=int" json:"scanProgressIntervalS" xml:"scanProgressIntervalS"`
	PullerPauseS            int                         `protobuf:"varint,20,opt,name=puller_pause_s,json=pullerPauseS,proto3,casttype=int" json:"pullerPauseS" xml:"pullerPauseS"`
	MaxConflicts            int                         `protobuf:"varint,21,opt,name=max_conflicts,json=maxConflicts,proto3,casttype=int" json:"maxConflicts" xml:"maxConflicts" default:"-1"`
	DisableSparseFiles      bool                        `protobuf:"varint,22,opt,name=disable_sparse_files,json=disableSparseFiles,proto3" json:"disableSparseFiles" xml:"disableSparseFiles"`
	DisableTempIndexes      bool                        `protobuf:"varint,23,opt,name=disable_temp_indexes,json=disableTempIndexes,proto3" json:"disableTempIndexes" xml:"disableTempIndexes"`
	Paused                  bool                        `protobuf:"varint,24,opt,name=paused,proto3" json:"paused" xml:"paused"`
	WeakHashThresholdPct    int                         `protobuf:"varint,25,opt,name=weak_hash_threshold_pct,json=weakHashThresholdPct,proto3,casttype=int" json:"weakHashThresholdPct" xml:"weakHashThresholdPct"`
	MarkerName              string                      `protobuf:"bytes,26,opt,name=marker_name,json=markerName,proto3" json:"markerName" xml:"markerName"`
	CopyOwnershipFromParent bool                        `protobuf:"varint,27,opt,name=copy_ownership_from_parent,json=copyOwnershipFromParent,proto3" json:"copyOwnershipFromParent" xml:"copyOwnershipFromParent"`
	RawModTimeWindowS       int                         `protobuf:"varint,28,opt,name=mod_time_window_s,json=modTimeWindowS,proto3,casttype=int" json:"modTimeWindowS" xml:"modTimeWindowS"`
	MaxConcurrentWrites     int                         `protobuf:"varint,29,opt,name=max_concurrent_writes,json=maxConcurrentWrites,proto3,casttype=int" json:"maxConcurrentWrites" xml:"maxConcurrentWrites" default:"2"`
	DisableFsync            bool                        `protobuf:"varint,30,opt,name=disable_fsync,json=disableFsync,proto3" json:"disableFsync" xml:"disableFsync"`
	BlockPullOrder          BlockPullOrder              `protobuf:"varint,31,opt,name=block_pull_order,json=blockPullOrder,proto3,enum=config.BlockPullOrder" json:"blockPullOrder" xml:"blockPullOrder"`
	CopyRangeMethod         fs.CopyRangeMethod          `protobuf:"varint,32,opt,name=copy_range_method,json=copyRangeMethod,proto3,enum=fs.CopyRangeMethod" json:"copyRangeMethod" xml:"copyRangeMethod" default:"standard"`
	CaseSensitiveFS         bool                        `protobuf:"varint,33,opt,name=case_sensitive_fs,json=caseSensitiveFs,proto3" json:"caseSensitiveFS" xml:"caseSensitiveFS"`
	JunctionsAsDirs         bool                        `protobuf:"varint,34,opt,name=follow_junctions,json=followJunctions,proto3" json:"junctionsAsDirs" xml:"junctionsAsDirs"`
	// Legacy deprecated
	DeprecatedReadOnly       bool    `protobuf:"varint,9000,opt,name=read_only,json=readOnly,proto3" json:"-" xml:"ro,attr,omitempty"`                       // Deprecated: Do not use.
	DeprecatedMinDiskFreePct float64 `protobuf:"fixed64,9001,opt,name=min_disk_free_pct,json=minDiskFreePct,proto3" json:"-" xml:"minDiskFreePct,omitempty"` // Deprecated: Do not use.
	DeprecatedPullers        int     `protobuf:"varint,9002,opt,name=pullers,proto3,casttype=int" json:"-" xml:"pullers,omitempty"`                          // Deprecated: Do not use.
}

func (m *FolderConfiguration) Reset()         { *m = FolderConfiguration{} }
func (m *FolderConfiguration) String() string { return proto.CompactTextString(m) }
func (*FolderConfiguration) ProtoMessage()    {}
func (*FolderConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_44a9785876ed3afa, []int{1}
}
func (m *FolderConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FolderConfiguration.Unmarshal(m, b)
}
func (m *FolderConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FolderConfiguration.Marshal(b, m, deterministic)
}
func (m *FolderConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FolderConfiguration.Merge(m, src)
}
func (m *FolderConfiguration) XXX_Size() int {
	return xxx_messageInfo_FolderConfiguration.Size(m)
}
func (m *FolderConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_FolderConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_FolderConfiguration proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FolderDeviceConfiguration)(nil), "config.FolderDeviceConfiguration")
	proto.RegisterType((*FolderConfiguration)(nil), "config.FolderConfiguration")
}

func init() {
	proto.RegisterFile("lib/config/folderconfiguration.proto", fileDescriptor_44a9785876ed3afa)
}

var fileDescriptor_44a9785876ed3afa = []byte{
	// 1981 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x58, 0xcd, 0x8f, 0x1c, 0x47,
	0x15, 0xdf, 0x5e, 0x7f, 0xed, 0xd6, 0x7e, 0x97, 0xbd, 0x76, 0x79, 0x9d, 0x6c, 0x6d, 0x9a, 0x21,
	0x6c, 0xa2, 0x78, 0x6d, 0x6f, 0x10, 0x07, 0x2b, 0x09, 0x64, 0xbc, 0x59, 0x30, 0xc6, 0xf1, 0xa8,
	0xd7, 0x60, 0x61, 0x90, 0x9a, 0x9e, 0xee, 0x9a, 0x99, 0xca, 0xf6, 0x17, 0x55, 0x3d, 0xde, 0x1d,
	0x9f, 0xcc, 0x05, 0x81, 0xc8, 0x01, 0x2d, 0x07, 0xae, 0x48, 0x20, 0x04, 0xf9, 0x07, 0x90, 0xf8,
	0x0b, 0x72, 0x41, 0x3b, 0x47, 0xc4, 0xa1, 0xa4, 0xac, 0x2f, 0x68, 0x6e, 0xf4, 0xd1, 0x27, 0x54,
	0x55, 0xdd, 0x3d, 0xdd, 0x3d, 0x13, 0x84, 0x94, 0xdb, 0xd4, 0xef, 0xf7, 0xab, 0xf7, 0x5e, 0xbf,
	0x7e, 0xef, 0x55, 0xf5, 0x80, 0x86, 0x4f, 0xdb, 0xb7, 0xdc, 0x28, 0xec, 0xd0, 0xee, 0xad, 0x4e,
	0xe4, 0x7b, 0x84, 0xe9, 0x45, 0x9f, 0x39, 0x09, 0x8d, 0xc2, 0x9d, 0x98, 0x45, 0x49, 0x04, 0x2f,
	0x6a, 0x70, 0xe3, 0xc6, 0x84, 0x3a, 0x19, 0xc4, 0x44, 0x8b, 0x36, 0xd6, 0x4b, 0x24, 0xa7, 0xcf,
	0x73, 0x78, 0xa3, 0x04, 0xc7, 0x7d, 0xdf, 0x8f, 0x98, 0x47, 0x58, 0xc6, 0x6d, 0x97, 0xb8, 0x67,
	0x84, 0x71, 0x1a, 0x85, 0x34, 0xec, 0x4e, 0x89, 0x60, 0x03, 0x97, 0x94, 0x6d, 0x3f, 0x72, 0x0f,
	0xeb, 0xa6, 0xa0, 0x14, 0x74, 0xf8, 0x2d, 0x19, 0x10, 0xcf, 0xb0, 0xd7, 0x32, 0xcc, 0x8d, 0xe2,
	0x01, 0x73, 0xc2, 0x2e, 0x09, 0x48, 0xd2, 0x8b, 0xbc, 0x8c, 0x9d, 0x27, 0xc7, 0x89, 0xfe, 0x69,
	0xfe, 0x67, 0x16, 0x5c, 0xdf, 0x57, 0xcf, 0xb3, 0x47, 0x9e, 0x51, 0x97, 0xdc, 0x2b, 0x47, 0x00,
	0x3f, 0x33, 0xc0, 0xbc, 0xa7, 0x70, 0x9b, 0x7a, 0xc8, 0xd8, 0x32, 0xb6, 0x17, 0x9b, 0x9f, 0x1a,
	0x9f, 0x0b, 0x3c, 0xf3, 0x2f, 0x81, 0xbf, 0xd9, 0xa5, 0x49, 0xaf, 0xdf, 0xde, 0x71, 0xa3, 0xe0,
	0x16, 0x1f, 0x84, 0x6e, 0xd2, 0xa3, 0x61, 0xb7, 0xf4, 0x4b, 0x86, 0xa0, 0x9c, 0xb8, 0x91, 0xbf,
	0xa3, 0xad, 0xdf, 0xdf, 0x3b, 0x13, 0x78, 0x2e, 0xff, 0x3d, 0x12, 0x78, 0xce, 0xcb, 0x7e, 0xa7,
	0x02, 0x2f, 0x1d, 0x07, 0xfe, 0x5d, 0x93, 0x7a, 0xef, 0x38, 0x49, 0xc2, 0xcc, 0xd1, 0x69, 0xe3,
	0x52, 0xf6, 0x3b, 0x3d, 0x6d, 0x14, 0xba, 0x5f, 0x0d, 0x1b, 0xc6, 0xc9, 0xb0, 0x51, 0xd8, 0xb0,
	0x72, 0xc6, 0x83, 0x7f, 0x36, 0xc0, 0x12, 0x0d, 0x13, 0x16, 0x79, 0x7d, 0x97, 0x78, 0x76, 0x7b,
	0x80, 0x66, 0x55, 0xc0, 0x2f, 0xbe, 0x52, 0xc0, 0x23, 0x81, 0x17, 0xc7, 0x56, 0x9b, 0x83, 0x54,
	0xe0, 0x6b, 0x3a, 0xd0, 0x12, 0x58, 0x84, 0xbc, 0x36, 0x81, 0xca, 0x80, 0xad, 0x8a, 0x05, 0xf3,
	0x25, 0x06, 0x97, 0x75, 0xce, 0xab, 0xd9, 0xfe, 0x00, 0xcc, 0x66, 0x59, 0x9e, 0x6f, 0xee, 0x9c,
	0x09, 0x3c, 0xab, 0xbc, 0xcf, 0x52, 0xef, 0x7f, 0x25, 0xe7, 0x64, 0xd8, 0x98, 0xbd, 0xbf, 0x67,
	0xcd, 0x52, 0x0f, 0xfe, 0x10, 0x5c, 0xf0, 0x9d, 0x36, 0xf1, 0xd5, 0x73, 0xcf, 0x37, 0xbf, 0x3d,
	0x12, 0x58, 0x03, 0xa9, 0xc0, 0x5b, 0x6a, 0xbf, 0x5a, 0x69, 0x13, 0x5b, 0x8c, 0xf0, 0xc4, 0x61,
	0xc9, 0x5d, 0xb3, 0xe3, 0xf8, 0x9c, 0x48, 0x93, 0x60, 0x4c, 0xbf, 0x18, 0x36, 0x66, 0x2c, 0xbd,
	0x19, 0x76, 0xc1, 0x4a, 0x87, 0xfa, 0x84, 0x0f, 0x78, 0x42, 0x02, 0x5b, 0x56, 0x19, 0x3a, 0xb7,
	0x65, 0x6c, 0x2f, 0xef, 0xc2, 0x9d, 0x0e, 0xdf, 0xd9, 0x2f, 0xa8, 0xc7, 0x83, 0x98, 0x34, 0xdf,
	0x1e, 0x09, 0xbc, 0xdc, 0xa9, 0x60, 0xa9, 0xc0, 0x57, 0x94, 0xf7, 0x2a, 0x6c, 0x5a, 0x35, 0x1d,
	0x7c, 0x0f, 0x9c, 0x8f, 0x9d, 0xa4, 0x87, 0xce, 0xab, 0xf0, 0xb7, 0x47, 0x02, 0xab, 0x75, 0x2a,
	0xf0, 0x8a, 0xda, 0x2f, 0x17, 0xc5, 0xf3, 0xcf, 0x17, 0x2b, 0x4b, 0xa9, 0x60, 0x0b, 0x9c, 0x57,
	0xb1, 0x5d, 0xc8, 0x62, 0xd3, 0x2d, 0xb3, 0xa3, 0x13, 0xad, 0x62, 0x53, 0x16, 0x13, 0x1d, 0x91,
	0xb6, 0x28, 0x17, 0x63, 0x8b, 0xc5, 0xca, 0x52, 0x2a, 0xf8, 0x53, 0x70, 0x49, 0x17, 0x17, 0x47,
	0x17, 0xb7, 0xce, 0x6d, 0x2f, 0xec, 0xbe, 0x51, 0x35, 0x3a, 0xa5, 0x63, 0x9a, 0x58, 0xd6, 0xda,
	0x48, 0xe0, 0x7c, 0x67, 0x2a, 0xf0, 0xa2, 0x72, 0xa5, 0xd7, 0xa6, 0x95, 0x13, 0xf0, 0x77, 0x06,
	0x58, 0x63, 0x84, 0xbb, 0x4e, 0x68, 0xd3, 0x30, 0x21, 0xec, 0x99, 0xe3, 0xdb, 0x1c, 0x5d, 0xda,
	0x32, 0xb6, 0x2f, 0x34, 0xbb, 0x23, 0x81, 0x57, 0x34, 0x79, 0x3f, 0xe3, 0x0e, 0x52, 0x81, 0xdf,
	0x52, 0x96, 0x6a, 0x78, 0xf6, 0x3a, 0x3d, 0xd2, 0x71, 0xfa, 0x7e, 0x72, 0xd7, 0x7c, 0xf7, 0x5b,
	0xb7, 0x6f, 0x9b, 0xaf, 0x04, 0x3e, 0x47, 0xc3, 0x64, 0x74, 0xda, 0xb8, 0x32, 0x4d, 0xfe, 0xea,
	0xb4, 0x71, 0x5e, 0xea, 0xac, 0xba, 0x13, 0xf8, 0x77, 0x03, 0xc0, 0x0e, 0xb7, 0x8f, 0x9c, 0xc4,
	0xed, 0x11, 0x66, 0x93, 0xd0, 0x69, 0xfb, 0xc4, 0x43, 0x73, 0x5b, 0xc6, 0xf6, 0x5c, 0xf3, 0x37,
	0xc6, 0x99, 0xc0, 0xab, 0xfb, 0x07, 0x4f, 0x34, 0xfb, 0x91, 0x26, 0x47, 0x02, 0xaf, 0x76, 0x78,
	0x15, 0x4b, 0x05, 0x7e, 0x5b, 0xbf, 0xf3, 0x1a, 0x51, 0x8f, 0x36, 0x61, 0x7d, 0x55, 0x7b, 0xeb,
	0x53, 0x85, 0x32, 0x4e, 0xa9, 0x38, 0x19, 0x36, 0x26, 0xdc, 0x5a, 0x13, 0x4e, 0xe1, 0xdf, 0xaa,
	0xc1, 0x7b, 0xc4, 0x77, 0x06, 0x36, 0x47, 0xf3, 0x2a, 0xa7, 0xbf, 0x96, 0xc1, 0xaf, 0x14, 0x56,
	0xf6, 0x24, 0x79, 0x20, 0xf3, 0x5c, 0x98, 0xd1, 0x50, 0x2a, 0xf0, 0x37, 0xaa, 0xa1, 0x6b, 0xbc,
	0x1e, 0xf9, 0x9d, 0x4a, 0x96, 0xa7, 0x89, 0x5f, 0x9d, 0x36, 0x66, 0xef, 0xdc, 0x3e, 0x19, 0x36,
	0xea, 0x5e, 0xad, 0xba, 0x4f, 0xf8, 0x33, 0xb0, 0x48, 0xbb, 0x61, 0xc4, 0x88, 0x1d, 0x13, 0x16,
	0x70, 0x04, 0x54, 0xbe, 0xdf, 0x1f, 0x09, 0xbc, 0xa0, 0xf1, 0x96, 0x84, 0x53, 0x81, 0xaf, 0xea,
	0x39, 0x30, 0xc6, 0x8a, 0xf2, 0x5d, 0xad, 0x83, 0x56, 0x79, 0x2b, 0xfc, 0x85, 0x01, 0x96, 0x9d,
	0x7e, 0x12, 0xd9, 0x61, 0xc4, 0x02, 0xc7, 0xa7, 0xcf, 0x09, 0x5a, 0x50, 0x4e, 0x9e, 0x8e, 0x04,
	0x5e, 0x92, 0xcc, 0xc7, 0x39, 0x51, 0x64, 0xa0, 0x82, 0x7e, 0xd9, 0x9b, 0x83, 0x93, 0xaa, 0xfc,
	0xb5, 0x59, 0x55, 0xbb, 0xf0, 0x29, 0x58, 0x0a, 0x68, 0x68, 0x7b, 0x94, 0x1f, 0xda, 0x1d, 0x46,
	0x08, 0x5a, 0xdc, 0x32, 0xb6, 0x17, 0x76, 0x17, 0xf3, 0xb6, 0x3a, 0xa0, 0xcf, 0x49, 0x73, 0x3b,
	0xeb, 0xa0, 0x85, 0x80, 0x86, 0x7b, 0x94, 0x1f, 0xee, 0x33, 0x22, 0x23, 0x5a, 0x53, 0x11, 0x95,
	0x30, 0xd3, 0x2a, 0x2b, 0x60, 0x17, 0x80, 0xf1, 0x39, 0x8a, 0x96, 0x94, 0x61, 0x9c, 0x1b, 0xfe,
	0x51, 0xc1, 0x54, 0xbb, 0xf5, 0xcd, 0xcc, 0x57, 0x69, 0x6b, 0x2a, 0xf0, 0xaa, 0x72, 0x35, 0x86,
	0x4c, 0xab, 0xc4, 0xc3, 0xf7, 0xc1, 0x25, 0x37, 0x8a, 0x29, 0x61, 0x1c, 0x2d, 0xab, 0xc2, 0xfa,
	0x9a, 0x6c, 0xf7, 0x0c, 0x2a, 0x26, 0x75, 0xb6, 0xce, 0x4b, 0xc4, 0xca, 0x05, 0xf0, 0x1f, 0x06,
	0xb8, 0x2a, 0x4f, 0x70, 0xc2, 0xec, 0xc0, 0x39, 0xb6, 0x63, 0x12, 0x7a, 0x34, 0xec, 0xda, 0x87,
	0xb4, 0x8d, 0x56, 0x94, 0xb9, 0xdf, 0xcb, 0x3a, 0xbd, 0xdc, 0x52, 0x92, 0x87, 0xce, 0x71, 0x4b,
	0x0b, 0x1e, 0xd0, 0xe6, 0x48, 0xe0, 0xcb, 0xf1, 0x24, 0x9c, 0x0a, 0x7c, 0x5d, 0x8f, 0xc7, 0x49,
	0xae, 0x54, 0xa1, 0x53, 0xb7, 0x4e, 0x87, 0x4f, 0x86, 0x8d, 0x69, 0xfe, 0xad, 0x29, 0xda, 0xb6,
	0x4c, 0x47, 0xcf, 0xe1, 0x3d, 0x99, 0x8e, 0xd5, 0x71, 0x3a, 0x32, 0xa8, 0x48, 0x47, 0xb6, 0x1e,
	0xa7, 0x23, 0x03, 0xe0, 0x87, 0xe0, 0x82, 0xba, 0xcb, 0xa0, 0x35, 0x35, 0xb6, 0xd7, 0xf2, 0x37,
	0x26, 0xfd, 0x3f, 0x92, 0x44, 0x13, 0xc9, 0x63, 0x4c, 0x69, 0x52, 0x81, 0x17, 0x94, 0x35, 0xb5,
	0x32, 0x2d, 0x8d, 0xc2, 0x07, 0x60, 0x29, 0xeb, 0x1d, 0x8f, 0xf8, 0x24, 0x21, 0x08, 0xaa, 0xba,
	0x7e, 0x53, 0x9d, 0xdc, 0x8a, 0xd8, 0x53, 0x78, 0x2a, 0x30, 0x2c, 0x75, 0x8f, 0x06, 0x4d, 0xab,
	0xa2, 0x81, 0xc7, 0x00, 0xa9, 0x91, 0x1c, 0xb3, 0xa8, 0xcb, 0x08, 0xe7, 0xe5, 0xd9, 0x7c, 0x59,
	0x3d, 0x9f, 0x3c, 0x56, 0xd7, 0xa5, 0xa6, 0x95, 0x49, 0xca, 0x13, 0xfa, 0x86, 0x72, 0x30, 0x95,
	0x2d, 0x9e, 0x7d, 0xfa, 0x66, 0x78, 0x00, 0x96, 0xb3, 0xba, 0x88, 0x9d, 0x3e, 0x27, 0x36, 0x47,
	0x57, 0x94, 0xbf, 0x9b, 0xf2, 0x39, 0x34, 0xd3, 0x92, 0xc4, 0x41, 0xf1, 0x1c, 0x65, 0xb0, 0xb0,
	0x5e, 0x91, 0x42, 0x02, 0x96, 0x64, 0x95, 0xc9, 0xa4, 0xfa, 0xd4, 0x4d, 0x38, 0x5a, 0x57, 0x36,
	0xbf, 0x23, 0x6d, 0x06, 0xce, 0xf1, 0xbd, 0x1c, 0x4f, 0x05, 0xc6, 0xba, 0xc1, 0x4a, 0x60, 0xa9,
	0xd9, 0x6f, 0xde, 0xc9, 0x1d, 0xc8, 0xa1, 0x76, 0xf3, 0x8e, 0x55, 0xd9, 0x0d, 0x3d, 0x70, 0xc5,
	0xa3, 0x5c, 0x0e, 0x61, 0x9b, 0xc7, 0x0e, 0xe3, 0xc4, 0x56, 0x47, 0x3b, 0xba, 0xaa, 0xde, 0xc4,
	0xee, 0x48, 0x60, 0x98, 0xf1, 0x07, 0x8a, 0x56, 0x97, 0x86, 0x54, 0x60, 0xa4, 0x8f, 0xc6, 0x09,
	0xca, 0xb4, 0xa6, 0xe8, 0xcb, 0x5e, 0x12, 0x12, 0xc4, 0x36, 0x0d, 0x3d, 0x72, 0x4c, 0x38, 0xba,
	0x36, 0xe1, 0xe5, 0x31, 0x09, 0xe2, 0xfb, 0x9a, 0xad, 0x7b, 0x29, 0x51, 0x63, 0x2f, 0x25, 0x10,
	0xee, 0x82, 0x8b, 0xea, 0x05, 0x78, 0x08, 0x29, 0xbb, 0x1b, 0x23, 0x81, 0x33, 0xa4, 0x38, 0xcc,
	0xf5, 0xd2, 0xb4, 0x32, 0x1c, 0x26, 0xe0, 0xda, 0x11, 0x71, 0x0e, 0x6d, 0x59, 0xd5, 0x76, 0xd2,
	0x63, 0x84, 0xf7, 0x22, 0xdf, 0xb3, 0x63, 0x37, 0x41, 0xd7, 0x55, 0xc2, 0xe5, 0x24, 0xbf, 0x22,
	0x25, 0xdf, 0x73, 0x78, 0xef, 0x71, 0x2e, 0x68, 0xb9, 0x49, 0x2a, 0xf0, 0x86, 0x32, 0x39, 0x8d,
	0x2c, 0x5e, 0xea, 0xd4, 0xad, 0xf0, 0x1e, 0x58, 0x08, 0x1c, 0x76, 0x48, 0x98, 0x1d, 0x3a, 0x01,
	0x41, 0x1b, 0xea, 0xda, 0x64, 0xca, 0x71, 0xa6, 0xe1, 0x8f, 0x9d, 0x80, 0x14, 0xe3, 0x6c, 0x0c,
	0x99, 0x56, 0x89, 0x87, 0x03, 0xb0, 0x21, 0x3f, 0x12, 0xec, 0xe8, 0x28, 0x24, 0x8c, 0xf7, 0x68,
	0x6c, 0x77, 0x58, 0x14, 0xd8, 0xb1, 0xc3, 0x48, 0x98, 0xa0, 0x1b, 0x2a, 0x05, 0xef, 0x8d, 0x04,
	0xbe, 0x26, 0x55, 0x8f, 0x72, 0xd1, 0x3e, 0x8b, 0x82, 0x96, 0x92, 0xa4, 0x02, 0xbf, 0x9e, 0x4f,
	0xbc, 0x69, 0xbc, 0x69, 0x7d, 0xd9, 0x4e, 0xf8, 0x4b, 0x03, 0xac, 0x05, 0x91, 0x67, 0x27, 0x34,
	0x20, 0xf6, 0x11, 0x0d, 0xbd, 0xe8, 0xc8, 0xe6, 0xe8, 0x35, 0x95, 0xb0, 0x9f, 0x9c, 0x09, 0xbc,
	0x66, 0x39, 0x47, 0x0f, 0x23, 0xef, 0x31, 0x0d, 0xc8, 0x13, 0xc5, 0xca, 0xe3, 0x7a, 0x39, 0xa8,
	0x20, 0xc5, 0xe5, 0xb2, 0x0a, 0xe7, 0x99, 0x3b, 0x19, 0x36, 0x26, 0xad, 0x58, 0x35, 0x1b, 0xf0,
	0x85, 0x01, 0xd6, 0xb3, 0x36, 0x71, 0xfb, 0x4c, 0xc6, 0x66, 0x1f, 0x31, 0x9a, 0x10, 0x8e, 0x5e,
	0x57, 0xc1, 0xfc, 0x40, 0x8e, 0x5e, 0x5d, 0xf0, 0x19, 0xff, 0x44, 0xd1, 0xa9, 0xc0, 0x5f, 0x2f,
	0x75, 0x4d, 0x85, 0x2b, 0x35, 0xcf, 0x6e, 0xa9, 0x77, 0x8c, 0x5d, 0x6b, 0x9a, 0x25, 0x39, 0xc4,
	0xf2, 0xda, 0xee, 0xc8, 0x2f, 0x12, 0xb4, 0x39, 0x1e, 0x62, 0x19, 0xb1, 0x2f, 0xf1, 0xa2, 0xf9,
	0xcb, 0xa0, 0x69, 0x55, 0x34, 0xd0, 0x07, 0xab, 0xea, 0x4b, 0xd1, 0x96, 0xb3, 0xc0, 0xd6, 0xf3,
	0x15, 0xab, 0xf9, 0x7a, 0x35, 0x9f, 0xaf, 0x4d, 0xc9, 0x8f, 0x87, 0xac, 0xba, 0xb6, 0xb7, 0x2b,
	0x58, 0x91, 0xd9, 0x2a, 0x6c, 0x5a, 0x35, 0x1d, 0xfc, 0xd4, 0x00, 0x6b, 0xaa, 0x84, 0xd4, 0x87,
	0xa6, 0xad, 0xbf, 0x34, 0xd1, 0x96, 0xf2, 0x77, 0x59, 0x7e, 0x22, 0xdc, 0x8b, 0xe2, 0x81, 0x25,
	0xb9, 0x87, 0x8a, 0x6a, 0x3e, 0x90, 0xb7, 0x2e, 0xb7, 0x0a, 0xa6, 0x02, 0x6f, 0x17, 0x65, 0x54,
	0xc2, 0x4b, 0x69, 0xe4, 0x89, 0x13, 0x7a, 0x0e, 0xf3, 0xcc, 0x57, 0xa7, 0x8d, 0xb9, 0x7c, 0x61,
	0xd5, 0x0d, 0xc1, 0x3f, 0xc9, 0x70, 0x1c, 0x39, 0x40, 0x49, 0xc8, 0x69, 0x42, 0x9f, 0xc9, 0x8c,
	0xa2, 0x37, 0x54, 0x3a, 0x8f, 0xe5, 0x15, 0xf0, 0x9e, 0xc3, 0xc9, 0x41, 0xce, 0xed, 0xab, 0x2b,
	0xa0, 0x5b, 0x85, 0x52, 0x81, 0xd7, 0x75, 0x30, 0x55, 0x5c, 0x5e, 0x77, 0x26, 0xb4, 0x93, 0x90,
	0xbc, 0xf1, 0xd5, 0x9c, 0x58, 0x35, 0x0d, 0x87, 0x7f, 0x34, 0xc0, 0x6a, 0x27, 0xf2, 0xfd, 0xe8,
	0xc8, 0xfe, 0xa4, 0x1f, 0xba, 0xf2, 0x3a, 0xc2, 0x91, 0x39, 0x8e, 0xf2, 0xfb, 0x39, 0xf8, 0x21,
	0xdf, 0xa3, 0x8c, 0xcb, 0x28, 0x3f, 0xa9, 0x42, 0x45, 0x94, 0x35, 0x5c, 0x45, 0x59, 0xd7, 0x4e,
	0x42, 0x32, 0xca, 0x9a, 0x13, 0x6b, 0x45, 0x47, 0x54, 0xc0, 0xf0, 0x10, 0xcc, 0x33, 0xe2, 0x78,
	0x76, 0x14, 0xfa, 0x03, 0xf4, 0x97, 0x7d, 0x15, 0xde, 0xc3, 0x33, 0x81, 0xe1, 0x1e, 0x89, 0x19,
	0x71, 0x9d, 0x84, 0x78, 0x16, 0x71, 0xbc, 0x47, 0xa1, 0x3f, 0x18, 0x09, 0x6c, 0xdc, 0x2c, 0xbe,
	0x8e, 0x59, 0xa4, 0x6e, 0x82, 0xef, 0x44, 0x01, 0x95, 0xb3, 0x3a, 0x19, 0xa8, 0xaf, 0xe3, 0x09,
	0x14, 0x19, 0xd6, 0x1c, 0xcb, 0x0c, 0xc0, 0x9f, 0x83, 0xb5, 0xca, 0xf5, 0x50, 0xcd, 0xcf, 0xbf,
	0x4a, 0xa7, 0x46, 0xf3, 0xa3, 0x33, 0x81, 0xd1, 0xd8, 0xe9, 0xc3, 0xf1, 0xcd, 0xaf, 0xe5, 0x26,
	0xb9, 0xeb, 0xcd, 0xfa, 0x1d, 0xb1, 0xe5, 0x26, 0xa5, 0x08, 0x90, 0x61, 0x2d, 0x57, 0x49, 0xf8,
	0x63, 0x70, 0x49, 0x9f, 0x97, 0x1c, 0x7d, 0xb6, 0xaf, 0x7a, 0xfd, 0x03, 0x39, 0x78, 0xc6, 0x8e,
	0xf4, 0x3d, 0x88, 0x57, 0x1f, 0x2e, 0xdb, 0x52, 0x32, 0x9d, 0x35, 0x38, 0x32, 0xac, 0xdc, 0x5e,
	0xf3, 0xbb, 0x9f, 0x7f, 0xb1, 0x39, 0x33, 0xfc, 0x62, 0x73, 0xe6, 0xdf, 0x67, 0x9b, 0x33, 0xbf,
	0x7d, 0xb9, 0x39, 0xf3, 0x87, 0x97, 0x9b, 0xc6, 0xf0, 0xe5, 0xe6, 0xcc, 0x3f, 0x5f, 0x6e, 0xce,
	0x3c, 0x7d, 0xeb, 0xff, 0xf8, 0x2f, 0x42, 0xb7, 0x6a, 0xfb, 0xa2, 0xfa, 0x4f, 0xe2, 0xdd, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0xe9, 0xb3, 0xf1, 0x6d, 0xb1, 0x12, 0x00, 0x00,
}

func (m *FolderDeviceConfiguration) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DeviceID.ProtoSize()
	n += 1 + l + sovFolderconfiguration(uint64(l))
	l = m.IntroducedBy.ProtoSize()
	n += 1 + l + sovFolderconfiguration(uint64(l))
	return n
}

func (m *FolderConfiguration) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovFolderconfiguration(uint64(l))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovFolderconfiguration(uint64(l))
	}
	if m.FilesystemType != 0 {
		n += 1 + sovFolderconfiguration(uint64(m.FilesystemType))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovFolderconfiguration(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovFolderconfiguration(uint64(m.Type))
	}
	if len(m.Devices) > 0 {
		for _, e := range m.Devices {
			l = e.ProtoSize()
			n += 1 + l + sovFolderconfiguration(uint64(l))
		}
	}
	if m.RescanIntervalS != 0 {
		n += 1 + sovFolderconfiguration(uint64(m.RescanIntervalS))
	}
	if m.FSWatcherEnabled {
		n += 2
	}
	if m.FSWatcherDelayS != 0 {
		n += 1 + sovFolderconfiguration(uint64(m.FSWatcherDelayS))
	}
	if m.IgnorePerms {
		n += 2
	}
	if m.AutoNormalize {
		n += 2
	}
	l = m.MinDiskFree.ProtoSize()
	n += 1 + l + sovFolderconfiguration(uint64(l))
	l = m.Versioning.ProtoSize()
	n += 1 + l + sovFolderconfiguration(uint64(l))
	if m.Copiers != 0 {
		n += 1 + sovFolderconfiguration(uint64(m.Copiers))
	}
	if m.PullerMaxPendingKiB != 0 {
		n += 1 + sovFolderconfiguration(uint64(m.PullerMaxPendingKiB))
	}
	if m.Hashers != 0 {
		n += 2 + sovFolderconfiguration(uint64(m.Hashers))
	}
	if m.Order != 0 {
		n += 2 + sovFolderconfiguration(uint64(m.Order))
	}
	if m.IgnoreDelete {
		n += 3
	}
	if m.ScanProgressIntervalS != 0 {
		n += 2 + sovFolderconfiguration(uint64(m.ScanProgressIntervalS))
	}
	if m.PullerPauseS != 0 {
		n += 2 + sovFolderconfiguration(uint64(m.PullerPauseS))
	}
	if m.MaxConflicts != 0 {
		n += 2 + sovFolderconfiguration(uint64(m.MaxConflicts))
	}
	if m.DisableSparseFiles {
		n += 3
	}
	if m.DisableTempIndexes {
		n += 3
	}
	if m.Paused {
		n += 3
	}
	if m.WeakHashThresholdPct != 0 {
		n += 2 + sovFolderconfiguration(uint64(m.WeakHashThresholdPct))
	}
	l = len(m.MarkerName)
	if l > 0 {
		n += 2 + l + sovFolderconfiguration(uint64(l))
	}
	if m.CopyOwnershipFromParent {
		n += 3
	}
	if m.RawModTimeWindowS != 0 {
		n += 2 + sovFolderconfiguration(uint64(m.RawModTimeWindowS))
	}
	if m.MaxConcurrentWrites != 0 {
		n += 2 + sovFolderconfiguration(uint64(m.MaxConcurrentWrites))
	}
	if m.DisableFsync {
		n += 3
	}
	if m.BlockPullOrder != 0 {
		n += 2 + sovFolderconfiguration(uint64(m.BlockPullOrder))
	}
	if m.CopyRangeMethod != 0 {
		n += 2 + sovFolderconfiguration(uint64(m.CopyRangeMethod))
	}
	if m.CaseSensitiveFS {
		n += 3
	}
	if m.JunctionsAsDirs {
		n += 3
	}
	if m.DeprecatedReadOnly {
		n += 4
	}
	if m.DeprecatedMinDiskFreePct != 0 {
		n += 11
	}
	if m.DeprecatedPullers != 0 {
		n += 3 + sovFolderconfiguration(uint64(m.DeprecatedPullers))
	}
	return n
}

func sovFolderconfiguration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFolderconfiguration(x uint64) (n int) {
	return sovFolderconfiguration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}