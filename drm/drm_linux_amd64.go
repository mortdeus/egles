// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs drm.go

package drm

import "syscall"

type (
	Handle   uint32
	Context  uint32
	Drawable uint32
	Magic    uint32
	ClipRect struct {
		X1 uint16
		Y1 uint16
		X2 uint16
		Y2 uint16
	}
	DrawableInfo struct {
		Num_Rects uint32
		Pad_cgo_0 [4]byte
		Rects     *ClipRect
	}
	TexRegion struct {
		Next    uint8
		Prev    uint8
		Use     uint8
		Padding uint8
		Age     uint32
	}
	HwLock struct {
		Lock    uint32
		Padding [60]int8
	}
	Version struct {
		Version_major      int32
		Version_minor      int32
		Version_patchlevel int32
		Pad_cgo_0          [4]byte
		Name_len           uint64
		Name               *int8
		Date_len           uint64
		Date               *int8
		Desc_len           uint64
		Desc               *int8
	}
	Unique struct {
		Len    uint64
		Unique *int8
	}
	List struct {
		Count     int32
		Pad_cgo_0 [4]byte
		Version   *Version
	}
	Block struct {
		Unused int32
	}
	Control struct {
		Func uint32
		Irq  int32
	}
	MapType    uint32
	MapFlags   uint32
	CtxPrivmap [0]byte
	Map        struct {
		Offset    uint64
		Size      uint64
		Type      uint32
		Flags     uint32
		Handle    *byte
		Mtrr      int32
		Pad_cgo_0 [4]byte
	}
	Client struct {
		Idx   int32
		Auth  int32
		Pid   uint64
		Uid   uint64
		Magic uint64
		Iocs  uint64
	}
	StatType uint32
	Stats    struct {
		Count uint64
		Data  [15]struct {
			Value     uint64
			Type      uint32
			Pad_cgo_0 [4]byte
		}
	}
	LockFlags uint32
	Lock      struct {
		Context int32
		Flags   uint32
	}
	DmaFlags uint32
	BufDesk  struct {
		Count     int32
		Size      int32
		Low_mark  int32
		High_mark int32
		Flags     uint32
		Pad_cgo_0 [4]byte
		Agp_start uint64
	}
	BufInfo struct {
		Count     int32
		Pad_cgo_0 [4]byte
		List      *BufDesk
	}
	BufFree struct {
		Count     int32
		Pad_cgo_0 [4]byte
		List      *int32
	}
	BufPub struct {
		Idx       int32
		Total     int32
		Used      int32
		Pad_cgo_0 [4]byte
		Address   *byte
	}
	BufMap struct {
		Count     int32
		Pad_cgo_0 [4]byte
		Virtual   *byte
		List      *BufPub
	}
	Dma struct {
		Context         int32
		Send_count      int32
		Send_indices    *int32
		Send_sizes      *int32
		Flags           uint32
		Request_count   int32
		Request_size    int32
		Pad_cgo_0       [4]byte
		Request_indices *int32
		Request_sizes   *int32
		Granted_count   int32
		Pad_cgo_1       [4]byte
	}
	CtxFlags uint32
	Ctx      struct {
		Handle uint32
		Flags  uint32
	}
	CtxRes struct {
		Count     int32
		Pad_cgo_0 [4]byte
		Contexts  *Ctx
	}
	Draw struct {
		Handle uint32
	}
	DrawableInfoType uint32
	UpdateDraw       struct {
		Handle    uint32
		Type      uint32
		Num       uint32
		Pad_cgo_0 [4]byte
		Data      uint64
	}
	Auth struct {
		Magic uint32
	}
	IrqBusid struct {
		Irq     int32
		Busnum  int32
		Devnum  int32
		Funcnum int32
	}
	VBlankSeqType     uint32
	WaitVBlankRequest struct {
		Type     uint32
		Sequence uint32
		Signal   uint64
	}
	WaitVBlankReply struct {
		Type     uint32
		Sequence uint32
		Sec      int64
		Usec     int64
	}
	WaitVBlank [24]byte
	ModesetCtl struct {
		Crtc uint32
		Cmd  uint32
	}
	AgpMode struct {
		Mode uint64
	}
	AgpBuffer struct {
		Size     uint64
		Handle   uint64
		Type     uint64
		Physical uint64
	}
	AgpBinding struct {
		Handle uint64
		Offset uint64
	}
	AgpInfo struct {
		Agp_version_major int32
		Agp_version_minor int32
		Mode              uint64
		Aperture_base     uint64
		Aperture_size     uint64
		Memory_allowed    uint64
		Memory_used       uint64
		Id_vendor         uint16
		Id_device         uint16
		Pad_cgo_0         [4]byte
	}
	ScatterGather struct {
		Size   uint64
		Handle uint64
	}
	SetVersion struct {
		Di_major int32
		Di_minor int32
		Dd_major int32
		Dd_minor int32
	}
	GemClose struct {
		Handle uint32
		Pad    uint32
	}
	GemFlink struct {
		Handle uint32
		Name   uint32
	}
	GemOpen struct {
		Name   uint32
		Handle uint32
		Size   uint64
	}
	GetCap struct {
		Capability uint64
		Value      uint64
	}
	PrimeHandle struct {
		Handle uint32
		Flags  uint32
		Fd     int32
	}
	Event struct {
		Type   uint32
		Length uint32
	}
	EventVBlank struct {
		Base      Event
		User_data uint64
		Tv_sec    uint32
		Tv_usec   uint32
		Sequence  uint32
		Reserved  uint32
	}
)

const (
	NAME        = "drm"
	MIN_ORDER   = 0x5
	MAX_ORDER   = 0x16
	RAM_PERCENT = 0xa
	LOCK_HELD   = 0x80000000
	LOCK_CONT   = 0x40000000

	FRAME_BUFFER   = 0x0
	REGISTERS      = 0x1
	SHM            = 0x2
	AGP            = 0x3
	SCATTER_GATHER = 0x4
	CONSISTENT     = 0x5
	GEM            = 0x6

	RESTRICTED      = 0x1
	READ_ONLY       = 0x2
	LOCKED          = 0x4
	KERNEL          = 0x8
	WRITE_COMBINING = 0x10
	CONTAINS_LOCK   = 0x20
	REMOVABLE       = 0x40
	DRIVER          = 0x80

	STAT_LOCK      = 0x0
	STAT_OPENS     = 0x1
	STAT_CLOSES    = 0x2
	STAT_IOCTLS    = 0x3
	STAT_LOCKS     = 0x4
	STAT_UNLOCKS   = 0x5
	STAT_VALUE     = 0x6
	STAT_BYTE      = 0x7
	STAT_COUNT     = 0x8
	STAT_IRQ       = 0x9
	STAT_PRIMARY   = 0xa
	STAT_SECONDARY = 0xb
	STAT_DMA       = 0xc
	STAT_SPECIAL   = 0xd
	STAT_MISSED    = 0xe

	LOCK_READY      = 0x1
	LOCK_QUIESCENT  = 0x2
	LOCK_FLUSH      = 0x4
	LOCK_FLUSH_ALL  = 0x8
	HALT_ALL_QUEUES = 0x10
	HALT_CUR_QUEUES = 0x20

	DMA_BLOCK        = 0x1
	DMA_WHILE_LOCKED = 0x2
	DMA_PRIORITY     = 0x4
	DMA_WaIT         = 0x10
	DMA_SMALLER_OK   = 0x20
	DMA_LARGER_OK    = 0x40

	PAGE_ALIGN    = 0x1
	AGP_BUFFER    = 0x2
	SG_BUFFER     = 0x4
	FB_BUFFER     = 0x8
	PCI_BUFFER_RO = 0x10

	CONTEXT_PRESERVED = 0x1
	CONTEXT_2DONLY    = 0x2

	VBLANK_ABSOLUTE   = 0x0
	VBLANK_RELATIVE   = 0x1
	VBLANK_EVENT      = 0x4000000
	VBLANK_FLIP       = 0x8000000
	VBLANK_NEXTONMISS = 0x10000000
	VBLANK_SECONDARY  = 0x20000000
	VBLANK_SIGNAL     = 0x40000000

	VBLANK_TYPES_MASK = 0x1
	VBLANK_FLAGS_MASK = 0x74000000
	PRE_MODESET       = 0x1
	POST_MODESET      = 0x2
	CLOEXEC           = syscall.O_CLOEXEC
)

const (
	DISPLAY_INFO_LEN   = 0x20
	CONNECTOR_NAME_LEN = 0x20
	DISPLAY_MODE_LEN   = 0x20
	PROP_NAME_LEN      = 0x20

	MODE_TYPE_BUILTIN   = 0x1
	MODE_TYPE_CLOCK_C   = 0x3
	MODE_TYPE_CRTC_C    = 0x5
	MODE_TYPE_PREFERRED = 0x8
	MODE_TYPE_DEFAULT   = 0x10
	MODE_TYPE_USERDEF   = 0x20
	MODE_TYPE_DRIVER    = 0x40

	MODE_FLAG_PHSYNC    = 0x1
	MODE_FLAG_NHSYNC    = 0x2
	MODE_FLAG_PVSYNC    = 0x4
	MODE_FLAG_NVSYNC    = 0x8
	MODE_FLAG_INTERLACE = 0x10
	MODE_FLAG_DBLSCAN   = 0x20
	MODE_FLAG_CSYNC     = 0x40
	MODE_FLAG_PCSYNC    = 0x80
	MODE_FLAG_NCSYNC    = 0x100

	MODE_FLAG_HSKEW   = 0x200
	MODE_FLAG_BCAST   = 0x400
	MODE_FLAG_PIXMUX  = 0x800
	MODE_FLAG_DBLCLK  = 0x1000
	MODE_FLAG_CLKDIV2 = 0x2000

	MODE_DPMS_ON      = 0x0
	MODE_DPMS_STANDBY = 0x1
	MODE_DPMS_SUSPEND = 0x2
	MODE_DPMS_OFF     = 0x3

	MODE_SCALE_NONE = 0x0

	MODE_SCALE_FULLSCREEN = 0x1

	MODE_SCALE_CENTER = 0x2

	MODE_SCALE_ASPECT = 0x3

	MODE_DITHERING_OFF  = 0x0
	MODE_DITHERING_ON   = 0x1
	MODE_DITHERING_AUTO = 0x2

	MODE_DIRTY_OFF      = 0x0
	MODE_DIRTY_ON       = 0x1
	MODE_DIRTY_ANNOTATE = 0x2
)

type (
	ModeModeInfo struct {
		Clock       uint32
		Hdisplay    uint16
		Hsync_start uint16
		Hsync_end   uint16
		Htotal      uint16
		Hskew       uint16
		Vdisplay    uint16
		Vsync_start uint16
		Vsync_end   uint16
		Vtotal      uint16
		Vscan       uint16
		Vrefresh    uint32
		Flags       uint32
		Type        uint32
		Name        [32]int8
	}
	ModeCardRes struct {
		Fb_id_ptr        uint64
		Crtc_id_ptr      uint64
		Connector_id_ptr uint64
		Encoder_id_ptr   uint64
		Count_fbs        uint32
		Count_crtcs      uint32
		Count_connectors uint32
		Count_encoders   uint32
		Min_width        uint32
		Max_width        uint32
		Min_height       uint32
		Max_height       uint32
	}
	ModeCrtc struct {
		Set_connectors_ptr uint64
		Count_connectors   uint32
		Crtc_id            uint32
		Fb_id              uint32
		X                  uint32
		Y                  uint32
		Gamma_size         uint32
		Mode_valid         uint32
		Mode               ModeModeInfo
	}
)

const (
	MODE_PRESENT_TOP_FIELD    = 0x1
	MODE_PRESENT_BOTTOM_FIELD = 0x2
)

type (
	ModeSetPlane struct {
		Plane_id uint32
		Crtc_id  uint32
		Fb_id    uint32
		Flags    uint32
		Crtc_x   int32
		Crtc_y   int32
		Crtc_w   uint32
		Crtc_h   uint32
		Src_x    uint32
		Src_y    uint32
		Src_h    uint32
		Src_w    uint32
	}
	ModeGetPlane struct {
		Plane_id           uint32
		Crtc_id            uint32
		Fb_id              uint32
		Possible_crtcs     uint32
		Gamma_size         uint32
		Count_format_types uint32
		Format_type_ptr    uint64
	}
	ModeGetPlaneRes struct {
		Plane_id_ptr uint64
		Count_planes uint32
		Pad_cgo_0    [4]byte
	}
)

const (
	MODE_ENCODER_NONE  = 0x0
	MODE_ENCODER_DAC   = 0x1
	MODE_ENCODER_TMDS  = 0x2
	MODE_ENCODER_LVDS  = 0x3
	MODE_ENCODER_TVDAC = 0x4
)

type ModeGetEncoder struct {
	Encoder_id      uint32
	Encoder_type    uint32
	Crtc_id         uint32
	Possible_crtcs  uint32
	Possible_clones uint32
}

const (
	MODE_SUBCONNECTOR_Automatic = 0x0
	MODE_SUBCONNECTOR_Unknown   = 0x0
	MODE_SUBCONNECTOR_DVID      = 0x3
	MODE_SUBCONNECTOR_DVIA      = 0x4
	MODE_SUBCONNECTOR_Composite = 0x5
	MODE_SUBCONNECTOR_SVIDEO    = 0x6
	MODE_SUBCONNECTOR_Component = 0x8
	MODE_SUBCONNECTOR_SCART     = 0x9

	MODE_CONNECTOR_Unknown     = 0x0
	MODE_CONNECTOR_VGA         = 0x1
	MODE_CONNECTOR_DVII        = 0x2
	MODE_CONNECTOR_DVID        = 0x3
	MODE_CONNECTOR_DVIA        = 0x4
	MODE_CONNECTOR_Composite   = 0x5
	MODE_CONNECTOR_SVIDEO      = 0x6
	MODE_CONNECTOR_LVDS        = 0x7
	MODE_CONNECTOR_Component   = 0x8
	MODE_CONNECTOR_9PinDIN     = 0x9
	MODE_CONNECTOR_DisplayPort = 0xa
	MODE_CONNECTOR_HDMIA       = 0xb
	MODE_CONNECTOR_HDMIB       = 0xc
	MODE_CONNECTOR_TV          = 0xd
	MODE_CONNECTOR_eDP         = 0xe
)

type ModeGetConnector struct {
	Encoders_ptr      uint64
	Modes_ptr         uint64
	Props_ptr         uint64
	Prop_values_ptr   uint64
	Count_modes       uint32
	Count_props       uint32
	Count_encoders    uint32
	Encoder_id        uint32
	Connector_id      uint32
	Connector_type    uint32
	Connector_type_id uint32
	Connection        uint32
	Mm_width          uint32
	Mm_height         uint32
	Subpixel          uint32
	Pad_cgo_0         [4]byte
}

const (
	MODE_PROP_PENDING   = 0x1
	MODE_PROP_RANGE     = 0x2
	MODE_PROP_IMMUTABLE = 0x4

	MODE_PROP_ENUM = 0x8
	MODE_PROP_BLOB = 0x10

	MODE_PROP_BITMASK = 0x20
)

type (
	ModePropertyEnum struct {
		Value uint64
		Name  [32]int8
	}
	ModeGetProperty struct {
		Values_ptr       uint64
		Enum_blob_ptr    uint64
		Prop_id          uint32
		Flags            uint32
		Name             [32]int8
		Count_values     uint32
		Count_enum_blobs uint32
	}
	ModeConnectorSetProperty struct {
		Value        uint64
		Prop_id      uint32
		Connector_id uint32
	}
)

const (
	MODE_OBJECT_CRTC      = 0xcccccccc
	MODE_OBJECT_CONNECTOR = 0xc0c0c0c0
	MODE_OBJECT_ENCODER   = 0xe0e0e0e0
	MODE_OBJECT_MODE      = 0xdededede
	MODE_OBJECT_PROPERTY  = 0xb0b0b0b0
	MODE_OBJECT_FB        = 0xfbfbfbfb
	MODE_OBJECT_BLOB      = 0xbbbbbbbb
	MODE_OBJECT_PLANE     = 0xeeeeeeee
)

type (
	ModeObjGetProperties struct {
		Props_ptr       uint64
		Prop_values_ptr uint64
		Count_props     uint32
		Obj_id          uint32
		Obj_type        uint32
		Pad_cgo_0       [4]byte
	}

	ModeObjSetProperty struct {
		Value     uint64
		Prop_id   uint32
		Obj_id    uint32
		Obj_type  uint32
		Pad_cgo_0 [4]byte
	}

	ModeGetBlob struct {
		Id     uint32
		Length uint32
		Data   uint64
	}

	ModeFbCmd struct {
		Id     uint32
		Width  uint32
		Height uint32
		Pitch  uint32
		Bpp    uint32
		Depth  uint32
		Handle uint32
	}
)

const MODE_FB_INTERLACED = 0x1

type ModeFbCmd2 struct {
	Fb_id        uint32
	Width        uint32
	Height       uint32
	Pixel_format uint32
	Flags        uint32
	Handles      [4]uint32
	Pitches      [4]uint32
	Offsets      [4]uint32
}

const (
	MODE_FB_DIRTY_ANNOTATE_COPY = 0x1
	MODE_FB_DIRTY_ANNOTATE_FILL = 0x2
	MODE_FB_DIRTY_FLAGS         = 0x3
)

type ModeFbDirtyCmd struct {
	Fb_id     uint32
	Flags     uint32
	Color     uint32
	Num_clips uint32
	Clips_ptr uint64
}

type ModeModeCmd struct {
	Id   uint32
	Mode ModeModeInfo
}

const (
	MODE_CURSOR_BO   = 0x1
	MODE_CURSOR_MOVE = 0x2
)

type (
	ModeCursor struct {
		Flags  uint32
		Id     uint32
		X      int32
		Y      int32
		Width  uint32
		Height uint32
		Handle uint32
	}
	ModeCrtcLut struct {
		Crtc_id    uint32
		Gamma_size uint32
		Red        uint64
		Green      uint64
		Blue       uint64
	}
)

const (
	MODE_PAGE_FLIP_EVENT = 0x1
	MODE_PAGE_FLIP_FLAGS = 0x1
)

type (
	ModeCrtcPageFlip struct {
		Crtc_id   uint32
		Fb_id     uint32
		Flags     uint32
		Reserved  uint32
		User_data uint64
	}

	ModeCreateDumb struct {
		Height uint32
		Width  uint32
		Bpp    uint32
		Flags  uint32
		Handle uint32
		Pitch  uint32
		Size   uint64
	}

	ModeMapDumb struct {
		Handle uint32
		Pad    uint32
		Offset uint64
	}

	ModeDestroyDumb struct {
		Handle uint32
	}
)

const (
	IOCTL_BASE                   = 'd'
	IOCTL_VERSION                = 0xc0406400
	IOCTL_GET_UNIQUE             = 0xc0106401
	IOCTL_GET_MAGIC              = 0x80046402
	IOCTL_IRQ_BUSID              = 0xc0106403
	IOCTL_GET_MAP                = 0xc0286404
	IOCTL_GET_CLIENT             = 0xc0286405
	IOCTL_GET_STATS              = 0x80f86406
	IOCTL_SET_VERSION            = 0xc0106407
	IOCTL_MODESET_CTL            = 0x40086408
	IOCTL_GEM_CLOSE              = 0x40086409
	IOCTL_GEM_FLINK              = 0xc008640a
	IOCTL_GEM_OPEN               = 0xc010640b
	IOCTL_GET_CAP                = 0xc010640c
	IOCTL_SET_UNIQUE             = 0x40106410
	IOCTL_AUTH_MAGIC             = 0x40046411
	IOCTL_BLOCK                  = 0xc0046412
	IOCTL_UNBLOCK                = 0xc0046413
	IOCTL_CONTROL                = 0x40086414
	IOCTL_ADD_MAP                = 0xc0286415
	IOCTL_ADD_BUFS               = 0xc0206416
	IOCTL_MARK_BUFS              = 0x40206417
	IOCTL_INFO_BUFS              = 0xc0106418
	IOCTL_MAP_BUFS               = 0xc0186419
	IOCTL_FREE_BUFS              = 0x4010641a
	IOCTL_RM_MAP                 = 0x4028641b
	IOCTL_SET_SAREA_CTX          = 0x4010641c
	IOCTL_GET_SAREA_CTX          = 0xc010641d
	IOCTL_SET_MASTER             = 0x641e
	IOCTL_DROP_MASTER            = 0x641f
	IOCTL_ADD_CTX                = 0xc0086420
	IOCTL_RM_CTX                 = 0xc0086421
	IOCTL_MOD_CTX                = 0x40086422
	IOCTL_GET_CTX                = 0xc0086423
	IOCTL_SWITCH_CTX             = 0x40086424
	IOCTL_NEW_CTX                = 0x40086425
	IOCTL_RES_CTX                = 0xc0106426
	IOCTL_ADD_DRAW               = 0xc0046427
	IOCTL_RM_DRAW                = 0xc0046428
	IOCTL_DMA                    = 0xc0406429
	IOCTL_LOCK                   = 0x4008642a
	IOCTL_UNLOCK                 = 0x4008642b
	IOCTL_FINISH                 = 0x4008642c
	IOCTL_PRIME_HANDLE_TO_FD     = 0xc00c642d
	IOCTL_PRIME_FD_TO_HANDLE     = 0xc00c642e
	IOCTL_AGP_AQUIRE             = 0x6430
	IOCTL_AGP_RELEASE            = 0x6431
	IOCTL_AGP_ENABLE             = 0x40086432
	IOCTL_AGP_INFO               = 0x80386433
	IOCTL_AGP_ALLOC              = 0xc0206434
	IOCTL_AGP_FREE               = 0x40206435
	IOCTL_AGP_BIND               = 0x40106436
	IOCTL_AGP_UNBIND             = 0x40106437
	IOCTL_SG_ALLOC               = 0xc0106438
	IOCTL_SG_FREE                = 0x40106439
	IOCTL_WAIT_VBLANK            = 0xc018643a
	IOCTL_UPDATE_DRAW            = 0x4018643f
	IOCTL_MODE_GETRESOURCES      = 0xc04064a0
	IOCTL_MODE_GETCRTC           = 0xc06864a1
	IOCTL_MODE_SETCRTC           = 0xc06864a2
	IOCTL_MODE_CURSOR            = 0xc01c64a3
	IOCTL_MODE_GETGAMMA          = 0xc02064a4
	IOCTL_MODE_SETGAMMA          = 0xc02064a5
	IOCTL_MODE_GETENCODER        = 0xc01464a6
	IOCTL_MODE_GETCONNECTOR      = 0xc05064a7
	IOCTL_MODE_ATTACHMODE        = 0xc04864a8
	IOCTL_MODE_DETACHMODE        = 0xc04864a9
	IOCTL_MODE_GETPROPERTY       = 0xc04064aa
	IOCTL_MODE_SETPROPERTY       = 0xc01064ab
	IOCTL_MODE_GETPROPBLOB       = 0xc01064ac
	IOCTL_MODE_GETFB             = 0xc01c64ad
	IOCTL_MODE_ADDFB             = 0xc01c64ae
	IOCTL_MODE_RMFB              = 0xc00464af
	IOCTL_MODE_PAGE_FLIP         = 0xc01864b0
	IOCTL_MODE_DIRTYFB           = 0xc01864b1
	IOCTL_MODE_CREATE_DUMB       = 0xc02064b2
	IOCTL_MODE_MAP_DUMB          = 0xc01064b3
	IOCTL_MODE_DESTROY_DUM       = 0xc00464b4
	IOCTL_MODE_GETPLANERESOURCES = 0xc01064b5
	IOCTL_MODE_GETPLANE          = 0xc02064b6
	IOCTL_MODE_SETPLANE          = 0xc03064b7
	IOCTL_MODE_ADDFB2            = 0xc04464b8
	IOCTL_MODE_OBJ_GETPROPERTIES = 0xc02064b9
	IOCTL_MODE_OBJ_SETPROPERTY   = 0xc01864ba

	COMMAND_BASE = 0x40
	COMMAND_END  = 0xa0

	EVENT_VBLANK             = 0x1
	EVENT_FLIP_COMPLETE      = 0x2
	CAP_DUMB_BUFFER          = 0x1
	CAP_VBLANK_HIGH_CRTC     = 0x2
	CAP_DUMB_PREFERRED_DEPTH = 0x3
	CAP_DUMB_PREFER_SHADOW   = 0x4
	CAP_PRIME                = 0x5
	PRIME_CAP_IMPORT         = 0x1
	PRIME_CAP_EXPORT         = 0x2
)

const (
	FORMAT_BIG_ENDIAN = -0x80000000
	FORMAT_C8         = 0x20203843
	FORMAT_RGB332     = 0x38424752
	FORMAT_BGR233     = 0x38524742
	FORMAT_XRGB4444   = 0x32315258
	FORMAT_XBGR4444   = 0x32314258
	FORMAT_RGBX4444   = 0x32315852
	FORMAT_BGRX4444   = 0x32315842

	FORMAT_ARGB4444 = 0x32315241
	FORMAT_ABGR4444 = 0x32314241
	FORMAT_RGBA4444 = 0x32314152
	FORMAT_BGRA4444 = 0x32314142

	FORMAT_XRGB1555 = 0x35315258
	FORMAT_XBGR1555 = 0x35314258
	FORMAT_RGBX5551 = 0x35315852
	FORMAT_BGRX5551 = 0x35315842

	FORMAT_ARGB1555 = 0x35315241
	FORMAT_ABGR1555 = 0x35314241
	FORMAT_RGBA5551 = 0x35314152
	FORMAT_BGRA5551 = 0x35314142
	FORMAT_RGB565   = 0x36314752
	FORMAT_BGR565   = 0x36314742
	FORMAT_RGB888   = 0x34324752
	FORMAT_BGR888   = 0x34324742
	FORMAT_XRGB8888 = 0x34325258
	FORMAT_XBGR8888 = 0x34324258
	FORMAT_RGBX8888 = 0x34325852
	FORMAT_BGRX8888 = 0x34325842

	FORMAT_ARGB8888 = 0x34325241
	FORMAT_ABGR8888 = 0x34324241
	FORMAT_RGBA8888 = 0x34324152
	FORMAT_BGRA8888 = 0x34324142

	FORMAT_XRGB2101010 = 0x30335258
	FORMAT_XBGR2101010 = 0x30334258
	FORMAT_RGBX1010102 = 0x30335852
	FORMAT_BGRX1010102 = 0x30335842

	FORMAT_ARGB2101010 = 0x30335241
	FORMAT_ABGR2101010 = 0x30334241
	FORMAT_RGBA1010102 = 0x30334152
	FORMAT_BGRA1010102 = 0x30334142
	FORMAT_YUYV        = 0x56595559
	FORMAT_YVYU        = 0x55595659
	FORMAT_UYVY        = 0x59565955
	FORMAT_VYUY        = 0x59555956
	FORMAT_AYUV        = 0x56555941
	FORMAT_NV12        = 0x3231564e
	FORMAT_NV21        = 0x3132564e
	FORMAT_NV16        = 0x3631564e
	FORMAT_NV61        = 0x3136564e
	FORMAT_YUV410      = 0x39565559
	FORMAT_YVU410      = 0x39555659
	FORMAT_YUV411      = 0x31315559
	FORMAT_YVU411      = 0x31315659
	FORMAT_YUV420      = 0x32315559
	FORMAT_YVU420      = 0x32315659
	FORMAT_YUV422      = 0x36315559
	FORMAT_YVU422      = 0x36315659
	FORMAT_YUV444      = 0x34325559
	FORMAT_YVU444      = 0x34325659
)

const (
	SAREA_MAX = 0x2000

	SAREA_MAX_DRAWABLES = 0x100

	SAREA_DRAWABLE_CLAIMED_ENTRY = 0x80000000
)

type (
	SareaDrawable struct {
		Stamp uint32
		Flags uint32
	}

	SareaFrame struct {
		X          uint32
		Y          uint32
		Width      uint32
		Height     uint32
		Fullscreen uint32
	}

	Sarea struct {
		Lock          HwLock
		Drawable_lock HwLock
		DrawableTable [256]SareaDrawable
		Frame         SareaFrame
		Dummy_context uint32
	}
)
