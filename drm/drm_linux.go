package drm

import (
	"github.com/mortdeus/egles/support/ioctl"
	"unsafe"
)

type Handle_t uint32

const (
	// Name in kernel, /dev, and /proc
	NAME = "drm"
	// At least 2^5 bytes = 32 bytes
	MIN_ORDER = 5
	// Up to 2^22 bytes = 4MB
	MAX_ORDER = 22
	// How much system ram can we lock?
	RAM_PERCENT = 10
	// Hardware lock is held 
	LOCK_HELD = 0x80000000
	// Hardware lock is contended
	LOCK_CONT = 0x40000000
)

func LockIsHeld(lock uint32) uint32 {
	return (lock & LOCK_HELD)
}
func LockIsCont(lock uint32) uint32 {
	return (lock & LOCK_CONT)
}
func LockingContext(lock uint32) uint32 {
	return lock & ^(uint32(LOCK_HELD) | uint32(LOCK_CONT))
}

type (
	Context_t  uint
	Drawable_t uint32
	Magic_t    uint32
	/*Cliprect.
	  [WARNING] If you change this structure, make sure you change
	  XF86DRIClipRectRec in the server as well  

	  [NOTE] KW: Actually it's illegal to change either for
	  backwards-compatibility reasons.*/
	ClipRect struct {
		X1,
		Y1,
		X2,
		Y2 uint16
	}
	/*Drawable information.*/
	DrawableInfo struct {
		NumRects uint32
		Rects    *ClipRect
	}

	/* Texture region,*/
	TexRegion struct {
		Next,
		Prev,
		InUse,
		Padding byte
		Age uint32
	}

	/*Hardware lock. 
	  The lock structure is a simple cache-line aligned integer.  To avoid
	  processor bus contention on a multiprocessor system, there should not be any
	  other data stored in the same cache line.*/
	HwLock struct {
		// lock variable
		Lock uint32
		// Pad to cache line 
		Padding [60]byte
	}
	/*IOCTL_VERSION [drmGetVersion()]*/
	Version struct {
		VersionMajor      int32
		VersionMinor      int32
		VersionPatchLevel int32
		// Length of name buffer
		NameLen int32
		// Name of driver
		Name *byte
		// Length of date buffer
		DateLen int32
		// User-space buffer to hold date
		Date *byte
		// Length of desc buffer
		DescLen int32
		// User-space buffer to hold desc
		Desc *byte
	}

	/* IOCTL_GET_UNIQUE	[drmGetBusid() and drmSetBusId()]*/
	Unique struct {
		// Length of unique
		UniqueLen int32
		// Unique name for driver instantiation
		Unique *byte
	}
	List struct {
		// Length of user-space structures 
		Count   int32
		Version *Version
	}
	Block struct {
		Unused int32
	}
)

/*IOCTL_CONTROL [drmCtlInstHandler() and drmCtlUninstHandler()]*/
const (
	ADD_COMMAND = iota
	RM_COMMAND
	INST_HANDLER
	UNINST_HANDLER
)

type Control struct {
	Func,
	Irq int32
}

/*Type of memory to map.*/
const (
	// WC (no caching), no core dump
	FRAME_BUFFER = iota
	// no caching, no core dump
	REGISTERS
	// shared, cached
	SHM
	// AGP/GART
	AGP
	// Scatter/gather memory for PCI DMA 
	SCATTER_GATHER
	// Consistent memory for PCI DMA
	CONSISTENT
	/* GEM object */
	GEM
)

type MapType int32

/* Memory mapping flags.*/
const (
	// Cannot be mapped to user-virtual
	RESTRICTED = 0x01
	READ_ONLY  = 0x02
	// shared, cached, locked
	LOCKED = 0x04
	// kernel requires access
	KERNEL = 0x08
	// use write-combining if available
	WRITE_COMBINING = 0x10
	// SHM page that contains lock
	CONTAINS_LOCK = 0x20
	// Removable mapping
	REMOVABLE = 0x40
	// Managed by driver
	DRIVER = 0x80
)

type MapFlags int32

type CtxPrivMap struct {
	// Context requesting private mapping
	CtxID uint32
	// Handle of map
	Handle unsafe.Pointer
}

/*IOCTL_GET_MAP, IOCTL_ADD_MAP and IOCTL_RM_MAP  [drmAddMap()]*/
type Map struct {
	// Requested physical address (0 for SAREA)
	Offset,
	// Requested physical size (bytes)
	Size uint32
	// Type of memory to map 
	Type  MapType
	Flags MapFlags
	// User-space: "Handle" to pass to mmap()
	Handle unsafe.Pointer

	// Kernel-space: kernel-virtual address 
	// MTRR slot used
	Mtrr int32

	/*Private data*/
}

/*IOCTL_GET_CLIENT */
type Client struct {
	// Which client desired?
	Idx,
	// Is client authenticated?
	Auth int32
	// Process ID
	Pid,
	// User ID
	Uid,
	Magic,
	// Ioctl Count
	Iocs uint32
}

const (
	STAT_LOCK = iota
	STAT_OPENS
	STAT_CLOSES
	STAT_IOCTLS
	STAT_LOCKS
	STAT_UNLOCKS
	// Generic value 
	STAT_VALUE
	// Generic byte counter (1024bytes/K) 
	STAT_BYTE
	// Generic non-byte counter (1000/k) 
	STAT_COUNT
	// IRQ 
	STAT_IRQ
	// Primary DMA bytes
	STAT_PRIMARY
	// Secondary DMA bytes
	STAT_SECONDARY
	// DMA
	STAT_DMA
	// Special DMA (e.g., priority or polled)
	STAT_SPECIAL
	// Missed DMA opportunity
	STAT_MISSED

/* Add to the *END* of the list */
)

type StatType int32

/* IOCTL_GET_STATS */
type Stats struct {
	Count uint32
	Data  [15]struct {
		Value uint32
		Type  StatType
	}
}

/*Hardware locking flags.*/
const (
	//Wait until hardware is ready for DMA 
	LOCK_READY = 0x01
	// Wait until hardware quiescent 
	LOCK_QUIESCENT = 0x02
	// Flush this context's DMA queue first 
	LOCK_FLUSH = 0x04
	// Flush all DMA queues first 
	LOCK_FLUSH_ALL = 0x08

	/* These *HALT* flags aren't supported yet they will 
	be used to support the full-screen DGA-like mode.*/

	// Halt all current and future queues 
	HALT_ALL_QUEUES = 0x10
	// Halt all current queues 
	HALT_CUR_QUEUES = 0x20
)

type LockFlags int32

/*IOCTL_LOCK, IOCTL_UNLOCK and IOCTL_FINISH [drmGetLock() and drmUnlock()]*/
type Lock struct {
	Context int32
	Flags   LockFlags
}

/*DMA Flags WARNING: These values must match xf86drm.h. Flags for DMA buffer dispatch */
const (
	/*	Block until buffer dispatched.	 
		NOTE: The buffer may not yet have been processed by the hardware. 
		Getting a hardware lock with the hardware quiescent will ensure 
		that the buffer has been processed.*/
	DMA_BLOCK = 0x01
	// Dispatch while lock held   
	DMA_WHILE_LOCKED = 0x02
	// High priority dispatch 
	DMA_PRIORITY = 0x04

	// Flags for DMA buffer request 

	//Wait for free buffers */
	DMA_WAIT = 0x10
	// Smaller-than-requested buffers OK */
	DMA_SMALLER_OK = 0x20
	// Larger-than-requested buffers OK */
	DMA_LARGER_OK = 0x40
)

type DmaFlags int32

/*IOCTL_ADD_BUFS and IOCTL_MARK_BUFS [drmAddBufs()]*/
const (
	// Align on page boundaries for DMA 
	PAGE_ALIGN = 0x01
	// Buffer is in AGP space
	AGP_BUFFER = 0x02
	// Scatter/gather memory buffer
	SG_BUFFER = 0x04
	// Buffer is in frame buffer
	FB_BUFFER = 0x08
	// Map PCI DMA buffer read-only
	PCI_BUFFER_RO = 0x10
)

type BufDesc struct {
	// Number of buffers of this size
	Count,
	// Size in bytes 
	Size,
	// Low water mark
	LowMark,
	// High water mark 
	HighMark int
	Flags int32
	//Start address of where the AGP buffers are in the AGP aperture
	AgpStart uint32
}

/*IOCTL_INFO_BUFS */
type BufInfo struct {
	Count int32 // Entries in list 
	List  *BufDesc
}

/*IOCTL_FREE_BUFS */
type BufFree struct {
	Count int
	List  *int
}

/* Buffer information */
type BufPub struct {
	// Index into the master buffer list 
	Idx,
	// Buffer size
	Total,
	// Amount of buffer in use (for DMA)
	Used int
	// Address of buffer
	Address unsafe.Pointer
}

/*IOCTL_MAP_BUFS */
type BufMap struct {
	// Length of the buffer list
	Count int32
	// Mmap'd area in user-virtual
	Virtual unsafe.Pointer
	// Buffer information
	List *BufPub
}

/*IOCTL_DMA Indices here refer to the offset into the buffer list in buf_get. */
type Dma struct {
	// Context handle
	Context,
	// Number of buffers to send
	SendCount int32
	// List of handles to buffers
	SendIndices,
	// Lengths of data to send
	SendSizes *int32
	Flags DmaFlags
	// Number of buffers requested
	RequestCount,
	// Desired size for buffers
	RequestSize int32
	// Buffer information
	RequestIndices,
	RequestSizes *int32
	// Number of buffers granted
	GrantedCount int32
}

const (
	CONTEXT_PRESERVED = 0x01
	CONTEXT_2DONLY    = 0x02
)

type CtxFlags int32

/*IOCTL_ADD_CTX */
type Ctx struct {
	Handle Context_t
	Flags  CtxFlags
}

/*IOCTL_RES_CTX */
type CtxRes struct {
	Count    int32
	Contexts *Ctx
}

/*IOCTL_ADD_DRAW and IOCTL_RM_DRAW */
type Draw struct {
	Handle Drawable_t
}

/*IOCTL_UPDATE_DRAW */
const (
	DRAWABLE_CLIPRECTS = 0
)

type DrawableInfoType_t int32

type UpdateDraw struct {
	Handle Drawable_t
	Type,
	Num uint32
	Data uint64
}

/* IOCTL_GET_MAGIC and IOCTL_AUTH_MAGIC */
type Auth struct {
	Magic Magic_t
}

/* IOCTL_IRQ_BUSID [drmGetInterruptFromBusID()]*/
type IrqBusID struct {
	// IRQ number
	Irq,
	// bus number
	BusNum,
	// device number
	DevNum,
	//function number
	FuncNum int32
}

const (
	// Wait for specific vblank sequence number
	VBLANK_ABSOLUTE = 0x0
	// Wait for given number of vblanks
	VBLANK_RELATIVE = 0x1
	// Send event instead of blocking
	VBLANK_EVENT = 0x4000000
	// Scheduled buffer swap should flip
	VBLANK_FLIP = 0x8000000
	// If missed, wait for next vblank
	VBLANK_NEXTONMISS = 0x10000000
	// Secondary display controller
	VBLANK_SECONDARY = 0x20000000
	// Send signal instead of blocking, unsupported
	VBLANK_SIGNAL = 0x40000000
)

type VblankSeqType int32

const (
	VBLANK_TYPES_MASK = (VBLANK_ABSOLUTE | VBLANK_RELATIVE)
	VBLANK_FLAGS_MASK = (VBLANK_EVENT | VBLANK_SIGNAL | VBLANK_SECONDARY | VBLANK_NEXTONMISS)
)

type WaitVblankRequest struct {
	Type     VblankSeqType
	Sequence uint32
	Signal   uint32
}

type WaitVblankReply struct {
	Type     VblankSeqType
	Sequence uint32
	TvalSec,
	TvalUsec int32
}

/*IOCTL_WAIT_VBLANK [drmWaitVBlank()]*/
//This was a union, how do I map this?
type WaitVblank struct {
	Request WaitVblankRequest
	Reply   WaitVblankReply
}

const (
	PRE_MODESET  = 1
	POST_MODESET = 2
)

/*IOCTL_MODESET_CTL[drmModesetCtl()]*/
type ModesetCtl struct {
	Crtc,
	Cmd uint32
}

/* IOCTL_AGP_ENABLE [drmAgpEnable()]*/
type AgpMode struct {
	mode uint32
}

/*IOCTL_AGP_ALLOC and IOCTL_AGP_FREE [drmAgpAlloc() and drmAgpFree()]*/
type AgpBuffer struct {
	// In bytes -- will round to page boundary
	Size,
	// Used for binding / unbinding
	Handle,
	// Type of memory to allocate
	Type,
	// Physical used by i810
	Physical uint32
}

/*IOCTL_AGP_BIND and IOCTL_AGP_UNBIND [drmAgpBind() and drmAgpUnbind()].*/
type AgpBinding struct {
	// From agp_buffer
	Handle,
	// In bytes -- will round to page boundary
	Offset uint32
}

/*IOCTL_AGP_INFO drmAgpVersionMajor(), drmAgpVersionMinor(), drmAgpGetMode(), drmAgpBase(), 
drmAgpSize(), drmAgpMemoryUsed(), drmAgpMemoryAvail(),drmAgpVendorId() and drmAgpDeviceId().*/
type AgpInfo struct {
	AgpVersionMajor,
	AgpVersionMinor int32
	Mode,
	// physical address 
	ApertureBase,
	// bytes 
	ApertureSize,
	// bytes
	MemoryAllowed,
	MemoryUsed uint32

	// PCI information 
	IdVendor,
	IdDevice uint16
}

/*IOCTL_SG_ALLOC */
type ScatterGather struct {
	// In bytes -- will round to page boundary 
	Size,
	// Used for mapping / unmapping 
	Handle uint32
}

/*IOCTL_SET_VERSION */
type SetVersion struct {
	DiMajor,
	DiMinor,
	DdMajor,
	DdMinor int32
}

/* IOCTL_GEM_CLOSE */
type GemClose struct {
	// Handle of the object to be closed. 
	Handle,
	Pad uint32
}

/* IOCTL_GEM_FLINK */
type GemFlink struct {
	//Handle for the object being named 
	Handle,

	// Returned global name 
	Name uint32
}

/* IOCTL_GEM_OPEN */
type GemOpen struct {
	// Name of object being opened 
	Name,

	// Returned handle for the object 
	Handle uint32

	// Returned size of the object 
	Size uint64
}

/* IOCTL_GET_CAP */
type GetCap struct {
	Capability,
	Value uint64
}

//What is this?
//const CLOEXEC = ioctl.O_CLOEXEC

type PrimeHandle struct {
	Handle,
	// Flags.. only applicable for handle->fd 
	Flags uint32
	// Returned dmabuf file descriptor 
	Fd int32
}

const IOCTL_BASE = 'd'

func IO(nr int32) int32          { return ioctl.IO(IOCTL_BASE, nr) }
func IOR(nr, _type int32) int32  { return ioctl.IOR(IOCTL_BASE, nr, _type) }
func IOW(nr, _type int32) int32  { return ioctl.IOW(IOCTL_BASE, nr, _type) }
func IOWR(nr, _type int32) int32 { return ioctl.IOWR(IOCTL_BASE, nr, _type) }

var (
	IOCTL_VERSION     = IOWR(0x00, sizeof["Version"])
	IOCTL_GET_UNIQUE  = IOWR(0x01, sizeof["Unique"])
	IOCTL_GET_MAGIC   = IOR(0x02, sizeof["Auth"])
	IOCTL_AUTH_MAGIC  = IOW(0x11, sizeof["Auth"])
	IOCTL_IRQ_BUSID   = IOWR(0x03, sizeof["IrqBusId"])
	IOCTL_GET_MAP     = IOWR(0x04, sizeof["Map"])
	IOCTL_GET_CLIENT  = IOWR(0x05, sizeof["Client"])
	IOCTL_GET_STATS   = IOR(0x06, sizeof["Stats"])
	IOCTL_SET_VERSION = IOWR(0x07, sizeof["SetVersion"])
	IOCTL_MODESET_CTL = IOW(0x08, sizeof["ModesetCtl"])

	IOCTL_GEM_CLOSE = IOW(0x09, sizeof["GemClose"])
	IOCTL_GEM_FLINK = IOWR(0x0a, sizeof["GemFlink"])
	IOCTL_GEM_OPEN  = IOWR(0x0b, sizeof["GemOpen"])

	IOCTL_GET_CAP    = IOWR(0x0c, sizeof["GetCap"])
	IOCTL_SET_UNIQUE = IOW(0x10, sizeof["Unique"])

	IOCTL_BLOCK   = IOWR(0x12, sizeof["Block"])
	IOCTL_UNBLOCK = IOWR(0x13, sizeof["Block"])

	IOCTL_SET_MASTER  = IO(0x1e)
	IOCTL_DROP_MASTER = IO(0x1f)

	IOCTL_CONTROL = IOW(0x14, sizeof["Control"])

	IOCTL_ADD_MAP   = IOWR(0x15, sizeof["Map"])
	IOCTL_ADD_BUFS  = IOWR(0x16, sizeof["BufDesc"])
	IOCTL_MARK_BUFS = IOW(0x17, sizeof["BufDesc"])
	IOCTL_INFO_BUFS = IOWR(0x18, sizeof["BufInfo"])
	IOCTL_MAP_BUFS  = IOWR(0x19, sizeof["BufMap"])
	IOCTL_FREE_BUFS = IOW(0x1a, sizeof["BufFree"])
	IOCTL_RM_MAP    = IOW(0x1b, sizeof["Map"])

	IOCTL_SET_SAREA_CTX = IOW(0x1c, sizeof["CtxPrivMap"])
	IOCTL_GET_SAREA_CTX = IOWR(0x1d, sizeof["CtxPrivMap"])
	IOCTL_ADD_CTX       = IOWR(0x20, sizeof["Ctx"])
	IOCTL_RM_CTX        = IOWR(0x21, sizeof["Ctx"])
	IOCTL_MOD_CTX       = IOW(0x22, sizeof["Ctx"])
	IOCTL_GET_CTX       = IOWR(0x23, sizeof["Ctx"])
	IOCTL_SWITCH_CTX    = IOW(0x24, sizeof["Ctx"])
	IOCTL_NEW_CTX       = IOW(0x25, sizeof["Ctx"])
	IOCTL_RES_CTX       = IOWR(0x26, sizeof["CtxRes"])

	IOCTL_ADD_DRAW = IOWR(0x27, sizeof["Draw"])
	IOCTL_RM_DRAW  = IOWR(0x28, sizeof["Draw"])
	IOCTL_DMA      = IOWR(0x29, sizeof["Dma"])

	IOCTL_LOCK   = IOW(0x2a, sizeof["Lock"])
	IOCTL_UNLOCK = IOW(0x2b, sizeof["Lock"])
	IOCTL_FINISH = IOW(0x2c, sizeof["Lock"])

	IOCTL_PRIME_HANDLE_TO_FD = IOWR(0x2d, sizeof["PrimeHandle"])
	IOCTL_PRIME_FD_TO_HANDLE = IOWR(0x2e, sizeof["PrimeHandle"])

	IOCTL_AGP_ACQUIRE = IO(0x30)
	IOCTL_AGP_RELEASE = IO(0x31)
	IOCTL_AGP_ENABLE  = IOW(0x32, sizeof["AgpMode"])
	IOCTL_AGP_INFO    = IOR(0x33, sizeof["AgpInfo"])
	IOCTL_AGP_ALLOC   = IOWR(0x34, sizeof["AgpBuffer"])
	IOCTL_AGP_FREE    = IOW(0x35, sizeof["AgpBuffer"])
	IOCTL_AGP_BIND    = IOW(0x36, sizeof["AgpBinding"])
	IOCTL_AGP_UNBIND  = IOW(0x37, sizeof["AgpBinding"])

	IOCTL_SG_ALLOC = IOWR(0x38, sizeof["ScatterGather"])
	IOCTL_SG_FREE  = IOW(0x39, sizeof["ScatterGather"])

	IOCTL_WAIT_VBLANK = IOWR(0x3a, sizeof["WaitVblank"])
	IOCTL_UPDATE_DRAW = IOW(0x3f, sizeof["UpdateDraw"])

	IOCTL_MODE_GETRESOURCES      = IOWR(0xA0, sizeof["ModeCardRes"])
	IOCTL_MODE_GETCRTC           = IOWR(0xA1, sizeof["ModeCrtc"])
	IOCTL_MODE_SETCRTC           = IOWR(0xA2, sizeof["ModeCrtc"])
	IOCTL_MODE_CURSOR            = IOWR(0xA3, sizeof["ModeCursor"])
	IOCTL_MODE_GETGAMMA          = IOWR(0xA4, sizeof["ModeCrtcLut"])
	IOCTL_MODE_SETGAMMA          = IOWR(0xA5, sizeof["ModeCrtcLut"])
	IOCTL_MODE_GETENCODER        = IOWR(0xA6, sizeof["ModeGetEncoder"])
	IOCTL_MODE_GETCONNECTOR      = IOWR(0xA7, sizeof["ModeGetConnector"])
	IOCTL_MODE_ATTACHMODE        = IOWR(0xA8, sizeof["ModeModeCmd"])
	IOCTL_MODE_DETACHMODE        = IOWR(0xA9, sizeof["ModeModeCmd"])
	IOCTL_MODE_GETPROPERTY       = IOWR(0xAA, sizeof["ModeGetProperty"])
	IOCTL_MODE_SETPROPERTY       = IOWR(0xAB, sizeof["ModeConnectorSetProperty"])
	IOCTL_MODE_GETPROPBLOB       = IOWR(0xAC, sizeof["ModeGetBlob"])
	IOCTL_MODE_GETFB             = IOWR(0xAD, sizeof["ModeFbCmd"])
	IOCTL_MODE_ADDFB             = IOWR(0xAE, sizeof["ModeFbCmd"])
	IOCTL_MODE_RMFB              = IOWR(0xAF, int32(4))
	IOCTL_MODE_PAGE_FLIP         = IOWR(0xB0, sizeof["ModeCrtcPageFlip"])
	IOCTL_MODE_DIRTYFB           = IOWR(0xB1, sizeof["ModeFbDirtyCmd"])
	IOCTL_MODE_CREATE_DUMB       = IOWR(0xB2, sizeof["ModeCreateDumb"])
	IOCTL_MODE_MAP_DUMB          = IOWR(0xB3, sizeof["ModeMapDumb"])
	IOCTL_MODE_DESTROY_DUMB      = IOWR(0xB4, sizeof["ModeDestroyDumb"])
	IOCTL_MODE_GETPLANERESOURCES = IOWR(0xB5, sizeof["ModeGetPlaneRes"])
	IOCTL_MODE_GETPLANE          = IOWR(0xB6, sizeof["ModeGetPlane"])
	IOCTL_MODE_SETPLANE          = IOWR(0xB7, sizeof["ModeSetPlane"])
	IOCTL_MODE_ADDFB2            = IOWR(0xB8, sizeof["ModeFbCmd2"])
	IOCTL_MODE_OBJ_GETPROPERTIES = IOWR(0xB9, sizeof["ModeObjGetProperties"])
	IOCTL_MODE_OBJ_SETPROPERTY   = IOWR(0xBA, sizeof["ModeObjSetProperty"])
)

/*Device specific ioctls should only be in their respective headers
  The device specific ioctl range is from 0x40 to 0x99.
  Generic IOCTLS restart at 0xA0. 
  [drmCommandNone(), drmCommandRead(), 
  drmCommandWrite(),drmCommandReadWrite()].*/
const (
	COMMAND_BASE = 0x40
	COMMAND_END  = 0xA0
)

/**
 * Header for events written back to userspace on the drm fd.  The
 * type defines the type of event, the length specifies the total
 * length of the event (including the header), and user_data is
 * typically a 64 bit value passed with the ioctl that triggered the
 * event.  A read on the drm fd will always only return complete
 * events, that is, if for example the read buffer is 100 bytes, and
 * there are two 64 byte events pending, only one will be returned.
 *
 * Event types 0 - 0x7fffffff are generic drm events, 0x80000000 and
 * up are chipset specific.
 */

type Event struct {
	Type,
	Length uint32
}

const (
	EVENT_VBLANK        = 0x01
	EVENT_FLIP_COMPLETE = 0x02
)

type EventVblank struct {
	EventBase struct {
		UserData uint64
		TvSec,
		TvUsec,
		Sequence,
		Reserved uint32
	}
}

const (
	CAP_DUMB_BUFFER          = 0x1
	CAP_VBLANK_HIGH_CRTC     = 0x2
	CAP_DUMB_PREFERRED_DEPTH = 0x3
	CAP_DUMB_PREFER_SHADOW   = 0x4
	CAP_PRIME                = 0x5

	PRIME_CAP_IMPORT = 0x1
	PRIME_CAP_EXPORT = 0x2
)
