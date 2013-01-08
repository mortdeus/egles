// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs drm.go

package drm

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
		NumRects  uint32
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

	Data struct {
		Value uint64
		Type  StatType
	}
	Stats struct {
		Count uint64
		Data  [15]Data
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
