package drm

import (
	"github.com/mortdeus/egles/support/ioctl"
	"unsafe"
)

type Handle_t uint32

const (
	NAME        = "drm" // Name in kernel, /dev, and /proc 
	MIN_ORDER   = 5     // At least 2^5 bytes = 32 bytes 
	MAX_ORDER   = 22    // Up to 2^22 bytes = 4MB 
	RAM_PERCENT = 10    // How much system ram can we lock? 

	LOCK_HELD = 0x80000000 // Hardware lock is held 
	LOCK_CONT = 0x40000000 // Hardware lock is contended 
)

func LockIsHeld(lock uint32) uint32 {
	return (lock & LOCK_HELD)
}
func LockIsCont(lock uint32) uint32 {
	return (lock & LOCK_CONT)
}
func LockingContext(lock uint32) uint32 {
	return lock & -(LOCK_HELD | LOCK_CONT)
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
	drm_clip_rect struct {
		x1,
		y1,
		x2,
		y2 uint16
	}
	/*Drawable information.*/
	drm_drawable_info struct {
		num_rects uint32
		rects     *drm_clip_rect
	}

	/* Texture region,*/
	drm_tex_region struct {
		next,
		prev,
		in_use,
		padding byte
		age uint32
	}

	/*Hardware lock. 
	  The lock structure is a simple cache-line aligned integer.  To avoid
	  processor bus contention on a multiprocessor system, there should not be any
	  other data stored in the same cache line.*/
	drm_hw_lock struct {
		lock    uint32   // lock variable 
		padding [60]byte // Pad to cache line 
	}

	/*DRM_IOCTL_VERSION ioctl argument type.
	[drmGetVersion()]*/
	drm_version struct {
		versionMajor      int32   // Major version 
		versionMinor      int32   // Minor version 
		versionPatchlevel int32   // Patch level 
		name_len          uintptr // Length of name buffer 
		name              *byte   // Name of driver 
		date_len          uintptr // Length of date buffer 
		date              *byte   // User-space buffer to hold date 
		desc_len          uintptr // Length of desc buffer 
		desc              *byte   // User-space buffer to hold desc 
	}

	/* DRM_IOCTL_GET_UNIQUE ioctl argument type. 
	[drmGetBusid() and drmSetBusId()]*/
	drm_unique struct {
		unique_len uintptr // Length of unique 
		unique     *byte   // Unique name for driver instantiation 
	}

	drm_list struct {
		count   int32 // Length of user-space structures 
		version *drm_version
	}

	drm_block struct {
		unused int32
	}
)

/*DRM_IOCTL_CONTROL ioctl argument type. 
  [drmCtlInstHandler() and drmCtlUninstHandler()]*/
const (
	DRM_ADD_COMMAND = iota
	DRM_RM_COMMAND
	DRM_INST_HANDLER
	DRM_UNINST_HANDLER
)

type drm_control struct {
	Func,
	Irq int32
}

/*Type of memory to map.*/
const (
	FRAME_BUFFER   = iota // WC (no caching), no core dump 
	REGISTERS             // no caching, no core dump 
	SHM                   // shared, cached 
	AGP                   // AGP/GART 
	SCATTER_GATHER        // Scatter/gather memory for PCI DMA 
	CONSISTENT            // Consistent memory for PCI DMA 
	GEM
)

type drm_map_type int32 // GEM object 

/* Memory mapping flags.*/
const (
	RESTRICTED      = 0x01 // Cannot be mapped to user-virtual 
	READ_ONLY       = 0x02
	LOCKED          = 0x04 // shared, cached, locked 
	KERNEL          = 0x08 // kernel requires access 
	WRITE_COMBINING = 0x10 // use write-combining if available 
	CONTAINS_LOCK   = 0x20 // SHM page that contains lock 
	REMOVABLE       = 0x40 // Removable mapping 
	DRIVER          = 0x80 // Managed by driver 
)

type drm_map_flags int32

type drm_ctx_priv_map struct {
	ctx_id uint32         // Context requesting private mapping 
	handle unsafe.Pointer // Handle of map 
}

/*DRM_IOCTL_GET_MAP, DRM_IOCTL_ADD_MAP and DRM_IOCTL_RM_MAP ioctls
  argument type.
  [drmAddMap()]*/
type drm_map struct {
	Offset, // Requested physical address (0 for SAREA)
	Size uint32 // Requested physical size (bytes) 
	Type   drm_map_type   // Type of memory to map 
	Flags  drm_map_flags  // Flags 
	handle unsafe.Pointer // User-space: "Handle" to pass to mmap() 
	// Kernel-space: kernel-virtual address 
	mtrr int32 // MTRR slot used 
	/*   Private data */
}

/**
 * DRM_IOCTL_GET_CLIENT ioctl argument type.
 */
type drm_client struct {
	idx, // Which client desired? 
	auth int32 // Is client authenticated? 
	pid, // Process ID 
	uid, // User ID 
	magic, // Magic 
	iocs uint32 // Ioctl count 
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

	STAT_IRQ       // IRQ 
	STAT_PRIMARY   // Primary DMA bytes 
	STAT_SECONDARY // Secondary DMA bytes 
	STAT_DMA       // DMA 
	STAT_SPECIAL   // Special DMA (e.g., priority or polled) 
	STAT_MISSED    // Missed DMA opportunity 
	/* Add to the *END* of the list */)

type drm_stat_type int32

/* DRM_IOCTL_GET_STATS ioctl argument type.*/
type drm_stats struct {
	count uint32
	data  [15]struct {
		Value uint32
		Type  drm_stat_type
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

	/* These *HALT* flags aren't supported yet
	   they will be used to support the
	   full-screen DGA-like mode.*/

	// Halt all current and future queues 
	HALT_ALL_QUEUES = 0x10
	// Halt all current queues 
	HALT_CUR_QUEUES = 0x20
)

type drm_lock_flags int32

/*DRM_IOCTL_LOCK, 
  DRM_IOCTL_UNLOCK and 
  DRM_IOCTL_FINISH ioctl argument type. 
  drmGetLock() and drmUnlock().*/

type drm_lock struct {
	context int32
	flags   drm_lock_flags
}

/*DMA flags  
  [WARNING]
  These values must match xf86drm.h.*/

//Flags for DMA buffer dispatch 
const (
	/*
	 Block until buffer dispatched.
	 note The buffer may not yet have
	 been processed by the hardware --
	 getting a hardware lock with the
	 hardware quiescent will ensure
	 that the buffer has been processed.
	*/
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

type drm_dma_flags int32

/**
 * DRM_IOCTL_ADD_BUFS and DRM_IOCTL_MARK_BUFS ioctl argument type.
 *
 * \sa drmAddBufs().
 */
const (
	PAGE_ALIGN    = 0x01 // Align on page boundaries for DMA */
	AGP_BUFFER    = 0x02 // Buffer is in AGP space */
	SG_BUFFER     = 0x04 // Scatter/gather memory buffer */
	FB_BUFFER     = 0x08 // Buffer is in frame buffer */
	PCI_BUFFER_RO = 0x10 // Map PCI DMA buffer read-only */
)

type drm_buf_desc struct {
	count, // Number of buffers of this size */
	size, // Size in bytes */
	low_mark, // Low water mark */
	high_mark int // High water mark */
	flags int32
	//Start address of where the AGP buffers are in the AGP aperture
	agp_start uint32
}

/**
 * DRM_IOCTL_INFO_BUFS ioctl argument type.
 */
type drm_buf_info struct {
	count int32 // Entries in list */
	list  *drm_buf_desc
}

/**
 * DRM_IOCTL_FREE_BUFS ioctl argument type.
 */
type drm_buf_free struct {
	count int
	list  *int
}

/* Buffer information */
type drm_buf_pub struct {
	idx, // Index into the master buffer list */
	total, // Buffer size */
	used int // Amount of buffer in use (for DMA) */
	address unsafe.Pointer // Address of buffer */
}

/*DRM_IOCTL_MAP_BUFS ioctl argument type.*/
type drm_buf_map struct {
	count   int32          // Length of the buffer list */
	virtual unsafe.Pointer // Mmap'd area in user-virtual */
	list    *drm_buf_pub   // Buffer information */
}

/*DRM_IOCTL_DMA ioctl argument type. 
  Indices here refer to the offset into the buffer list in drm_buf_get. */
type drm_dma struct {
	context, // Context handle */
	send_count int32 // Number of buffers to send */
	send_indices, // List of handles to buffers */
	send_sizes *int32 // Lengths of data to send */
	flags          drm_dma_flags // Flags */
	request_count, // Number of buffers requested */
	request_size int32 // Desired size for buffers */
	request_indices, // Buffer information */
	request_sizes *int32
	granted_count int32 // Number of buffers granted */
}

const (
	CONTEXT_PRESERVED = 0x01
	CONTEXT_2DONLY    = 0x02
)

type drm_ctx_flags int32

/*DRM_IOCTL_ADD_CTX ioctl argument type.*/
type drm_ctx struct {
	handle drm_context_t
	flags  drm_ctx_flags
}

/**
 * DRM_IOCTL_RES_CTX ioctl argument type.
 */
type drm_ctx_res struct {
	count    int32
	contexts *drm_ctx
}

/**
 * DRM_IOCTL_ADD_DRAW and DRM_IOCTL_RM_DRAW ioctl argument type.
 */
type drm_draw struct {
	drm_drawable_t handle
}

/*DRM_IOCTL_UPDATE_DRAW ioctl argument type.*/
const (
	DRM_DRAWABLE_CLIPRECTS = 0
)

type drm_drawable_info_type_t int32

type drm_update_draw struct {
	handle drm_drawable_t
	Type,
	num uint32
	data uint64
}

/* DRM_IOCTL_GET_MAGIC and DRM_IOCTL_AUTH_MAGIC ioctl argument type.*/
type drm_auth struct {
	magic drm_magic_t
}

/**
 * DRM_IOCTL_IRQ_BUSID ioctl argument type.
 *
 * \sa drmGetInterruptFromBusID().
 */
type drm_irq_busid struct {
	irq, // IRQ number */
	busnum, // bus number */
	devnum, // device number */
	funcnum int32 // function number */
}

// enum drm_vblank_seq_type {
// 	VBLANK_ABSOLUTE = 0x0,	// Wait for specific vblank sequence number */
// 	VBLANK_RELATIVE = 0x1,	// Wait for given number of vblanks */
// 	VBLANK_EVENT = 0x4000000,   // Send event instead of blocking */
// 	VBLANK_FLIP = 0x8000000,   // Scheduled buffer swap should flip */
// 	VBLANK_NEXTONMISS = 0x10000000,	*< If missed, wait for next vblank 
// 	VBLANK_SECONDARY = 0x20000000,	// Secondary display controller */
// 	VBLANK_SIGNAL = 0x40000000	// Send signal instead of blocking, unsupported */
// }

// #define VBLANK_TYPES_MASK (VBLANK_ABSOLUTE | VBLANK_RELATIVE)
// #define VBLANK_FLAGS_MASK (VBLANK_EVENT | VBLANK_SIGNAL | \
// 				VBLANK_SECONDARY | VBLANK_NEXTONMISS)

type drm_wait_vblank_request struct {
	Type     drm_vblank_seq_type
	sequence uint32
	signal   uint32
}

type drm_wait_vblank_reply struct {
	Type     drm_vblank_seq_type
	sequence uint32
	tval_sec,
	tval_usec int32
}

/**
 * DRM_IOCTL_WAIT_VBLANK ioctl argument type.
 *
 * \sa drmWaitVBlank().
 */
//This was a union.
type drm_wait_vblank struct {
	request drm_wait_vblank_request
	reply   drm_wait_vblank_reply
}

const (
	PRE_MODESET  = 1
	POST_MODESET = 2
)

/**
 * DRM_IOCTL_MODESET_CTL ioctl argument type
 *
 * \sa drmModesetCtl().
 */
type drm_modeset_ctl struct {
	crtc,
	cmd uint32
}

/**
 * DRM_IOCTL_AGP_ENABLE ioctl argument type.
 *
 * \sa drmAgpEnable().
 */
type drm_agp_mode struct {
	mode uint32 // AGP mode */
}

/**
 * DRM_IOCTL_AGP_ALLOC and DRM_IOCTL_AGP_FREE ioctls argument type.
 *
 * \sa drmAgpAlloc() and drmAgpFree().
 */
type drm_agp_buffer struct {
	size, // In bytes -- will round to page boundary */
	handle, // Used for binding / unbinding */
	Type, // Type of memory to allocate */
	physical uint32 // Physical used by i810 */
}

/**
 * DRM_IOCTL_AGP_BIND and DRM_IOCTL_AGP_UNBIND ioctls argument type.
 *
 * \sa drmAgpBind() and drmAgpUnbind().
 */
type drm_agp_binding struct {
	handle, // From drm_agp_buffer */
	offset uint32 // In bytes -- will round to page boundary */
}

/**
 * DRM_IOCTL_AGP_INFO ioctl argument type.
 *
 * \sa drmAgpVersionMajor(), drmAgpVersionMinor(), drmAgpGetMode(),
 * drmAgpBase(), drmAgpSize(), drmAgpMemoryUsed(), drmAgpMemoryAvail(),
 * drmAgpVendorId() and drmAgpDeviceId().
 */
type drm_agp_info struct {
	agp_version_major,
	agp_version_minor int32
	mode,
	aperture_base, /* physical address */
	aperture_size, /* bytes */
	memory_allowed, /* bytes */
	memory_used uint32

	/* PCI information */
	id_vendor,
	id_device uint16
}

/**
 * DRM_IOCTL_SG_ALLOC ioctl argument type.
 */
type drm_scatter_gather struct {
	size, // In bytes -- will round to page boundary */
	handle uint32 // Used for mapping / unmapping */
}

/**
 * DRM_IOCTL_SET_VERSION ioctl argument type.
 */
type drm_set_version struct {
	drm_di_major,
	drm_di_minor,
	drm_dd_major,
	drm_dd_minor int32
}

/** DRM_IOCTL_GEM_CLOSE ioctl argument type */
type drm_gem_close struct {
	/** Handle of the object to be closed. */
	handle,
	pad uint32
}

/** DRM_IOCTL_GEM_FLINK ioctl argument type */
type drm_gem_flink struct {
	/** Handle for the object being named */
	handle,

	/** Returned global name */
	name uint32
}

/** DRM_IOCTL_GEM_OPEN ioctl argument type */
type drm_gem_open struct {
	/** Name of object being opened */
	name,

	/** Returned handle for the object */
	handle uint32

	/** Returned size of the object */
	size uint64
}

/* DRM_IOCTL_GET_CAP ioctl argument type */
type drm_get_cap struct {
	capability,
	value uint64
}

//#define DRM_CLOEXEC O_CLOEXEC

type drm_prime_handle struct {
	handle,

	/** Flags.. only applicable for handle->fd */
	flags uint32

	/** Returned dmabuf file descriptor */
	fd int32
}

/*
#include "drm_mode.h"

#define DRM_IOCTL_BASE			'd'
#define DRM_IO(nr)			_IO(DRM_IOCTL_BASE,nr)
#define DRM_IOR(nr,type)		_IOR(DRM_IOCTL_BASE,nr,type)
#define DRM_IOW(nr,type)		_IOW(DRM_IOCTL_BASE,nr,type)
#define DRM_IOWR(nr,type)		_IOWR(DRM_IOCTL_BASE,nr,type)

#define DRM_IOCTL_VERSION		DRM_IOWR(0x00, struct drm_version)
#define DRM_IOCTL_GET_UNIQUE		DRM_IOWR(0x01, struct drm_unique)
#define DRM_IOCTL_GET_MAGIC		DRM_IOR( 0x02, struct drm_auth)
#define DRM_IOCTL_IRQ_BUSID		DRM_IOWR(0x03, struct drm_irq_busid)
#define DRM_IOCTL_GET_MAP               DRM_IOWR(0x04, struct drm_map)
#define DRM_IOCTL_GET_CLIENT            DRM_IOWR(0x05, struct drm_client)
#define DRM_IOCTL_GET_STATS             DRM_IOR( 0x06, struct drm_stats)
#define DRM_IOCTL_SET_VERSION		DRM_IOWR(0x07, struct drm_set_version)
#define DRM_IOCTL_MODESET_CTL           DRM_IOW(0x08, struct drm_modeset_ctl)
#define DRM_IOCTL_GEM_CLOSE		DRM_IOW (0x09, struct drm_gem_close)
#define DRM_IOCTL_GEM_FLINK		DRM_IOWR(0x0a, struct drm_gem_flink)
#define DRM_IOCTL_GEM_OPEN		DRM_IOWR(0x0b, struct drm_gem_open)
#define DRM_IOCTL_GET_CAP		DRM_IOWR(0x0c, struct drm_get_cap)

#define DRM_IOCTL_SET_UNIQUE		DRM_IOW( 0x10, struct drm_unique)
#define DRM_IOCTL_AUTH_MAGIC		DRM_IOW( 0x11, struct drm_auth)
#define DRM_IOCTL_BLOCK			DRM_IOWR(0x12, struct drm_block)
#define DRM_IOCTL_UNBLOCK		DRM_IOWR(0x13, struct drm_block)
#define DRM_IOCTL_CONTROL		DRM_IOW( 0x14, struct drm_control)
#define DRM_IOCTL_ADD_MAP		DRM_IOWR(0x15, struct drm_map)
#define DRM_IOCTL_ADD_BUFS		DRM_IOWR(0x16, struct drm_buf_desc)
#define DRM_IOCTL_MARK_BUFS		DRM_IOW( 0x17, struct drm_buf_desc)
#define DRM_IOCTL_INFO_BUFS		DRM_IOWR(0x18, struct drm_buf_info)
#define DRM_IOCTL_MAP_BUFS		DRM_IOWR(0x19, struct drm_buf_map)
#define DRM_IOCTL_FREE_BUFS		DRM_IOW( 0x1a, struct drm_buf_free)

#define DRM_IOCTL_RM_MAP		DRM_IOW( 0x1b, struct drm_map)

#define DRM_IOCTL_SET_SAREA_CTX		DRM_IOW( 0x1c, struct drm_ctx_priv_map)
#define DRM_IOCTL_GET_SAREA_CTX 	DRM_IOWR(0x1d, struct drm_ctx_priv_map)

#define DRM_IOCTL_SET_MASTER            DRM_IO(0x1e)
#define DRM_IOCTL_DROP_MASTER           DRM_IO(0x1f)

#define DRM_IOCTL_ADD_CTX		DRM_IOWR(0x20, struct drm_ctx)
#define DRM_IOCTL_RM_CTX		DRM_IOWR(0x21, struct drm_ctx)
#define DRM_IOCTL_MOD_CTX		DRM_IOW( 0x22, struct drm_ctx)
#define DRM_IOCTL_GET_CTX		DRM_IOWR(0x23, struct drm_ctx)
#define DRM_IOCTL_SWITCH_CTX		DRM_IOW( 0x24, struct drm_ctx)
#define DRM_IOCTL_NEW_CTX		DRM_IOW( 0x25, struct drm_ctx)
#define DRM_IOCTL_RES_CTX		DRM_IOWR(0x26, struct drm_ctx_res)
#define DRM_IOCTL_ADD_DRAW		DRM_IOWR(0x27, struct drm_draw)
#define DRM_IOCTL_RM_DRAW		DRM_IOWR(0x28, struct drm_draw)
#define DRM_IOCTL_DMA			DRM_IOWR(0x29, struct drm_dma)
#define DRM_IOCTL_LOCK			DRM_IOW( 0x2a, struct drm_lock)
#define DRM_IOCTL_UNLOCK		DRM_IOW( 0x2b, struct drm_lock)
#define DRM_IOCTL_FINISH		DRM_IOW( 0x2c, struct drm_lock)

#define DRM_IOCTL_PRIME_HANDLE_TO_FD    DRM_IOWR(0x2d, struct drm_prime_handle)
#define DRM_IOCTL_PRIME_FD_TO_HANDLE    DRM_IOWR(0x2e, struct drm_prime_handle)

#define DRM_IOCTL_AGP_ACQUIRE		DRM_IO(  0x30)
#define DRM_IOCTL_AGP_RELEASE		DRM_IO(  0x31)
#define DRM_IOCTL_AGP_ENABLE		DRM_IOW( 0x32, struct drm_agp_mode)
#define DRM_IOCTL_AGP_INFO		DRM_IOR( 0x33, struct drm_agp_info)
#define DRM_IOCTL_AGP_ALLOC		DRM_IOWR(0x34, struct drm_agp_buffer)
#define DRM_IOCTL_AGP_FREE		DRM_IOW( 0x35, struct drm_agp_buffer)
#define DRM_IOCTL_AGP_BIND		DRM_IOW( 0x36, struct drm_agp_binding)
#define DRM_IOCTL_AGP_UNBIND		DRM_IOW( 0x37, struct drm_agp_binding)

#define DRM_IOCTL_SG_ALLOC		DRM_IOWR(0x38, struct drm_scatter_gather)
#define DRM_IOCTL_SG_FREE		DRM_IOW( 0x39, struct drm_scatter_gather)

#define DRM_IOCTL_WAIT_VBLANK		DRM_IOWR(0x3a, union drm_wait_vblank)

#define DRM_IOCTL_UPDATE_DRAW		DRM_IOW(0x3f, struct drm_update_draw)

#define DRM_IOCTL_MODE_GETRESOURCES	DRM_IOWR(0xA0, struct drm_mode_card_res)
#define DRM_IOCTL_MODE_GETCRTC		DRM_IOWR(0xA1, struct drm_mode_crtc)
#define DRM_IOCTL_MODE_SETCRTC		DRM_IOWR(0xA2, struct drm_mode_crtc)
#define DRM_IOCTL_MODE_CURSOR		DRM_IOWR(0xA3, struct drm_mode_cursor)
#define DRM_IOCTL_MODE_GETGAMMA		DRM_IOWR(0xA4, struct drm_mode_crtc_lut)
#define DRM_IOCTL_MODE_SETGAMMA		DRM_IOWR(0xA5, struct drm_mode_crtc_lut)
#define DRM_IOCTL_MODE_GETENCODER	DRM_IOWR(0xA6, struct drm_mode_get_encoder)
#define DRM_IOCTL_MODE_GETCONNECTOR	DRM_IOWR(0xA7, struct drm_mode_get_connector)
#define DRM_IOCTL_MODE_ATTACHMODE	DRM_IOWR(0xA8, struct drm_mode_mode_cmd)
#define DRM_IOCTL_MODE_DETACHMODE	DRM_IOWR(0xA9, struct drm_mode_mode_cmd)

#define DRM_IOCTL_MODE_GETPROPERTY	DRM_IOWR(0xAA, struct drm_mode_get_property)
#define DRM_IOCTL_MODE_SETPROPERTY	DRM_IOWR(0xAB, struct drm_mode_connector_set_property)
#define DRM_IOCTL_MODE_GETPROPBLOB	DRM_IOWR(0xAC, struct drm_mode_get_blob)
#define DRM_IOCTL_MODE_GETFB		DRM_IOWR(0xAD, struct drm_mode_fb_cmd)
#define DRM_IOCTL_MODE_ADDFB		DRM_IOWR(0xAE, struct drm_mode_fb_cmd)
#define DRM_IOCTL_MODE_RMFB		DRM_IOWR(0xAF, unsigned int)
#define DRM_IOCTL_MODE_PAGE_FLIP	DRM_IOWR(0xB0, struct drm_mode_crtc_page_flip)
#define DRM_IOCTL_MODE_DIRTYFB		DRM_IOWR(0xB1, struct drm_mode_fb_dirty_cmd)

#define DRM_IOCTL_MODE_CREATE_DUMB DRM_IOWR(0xB2, struct drm_mode_create_dumb)
#define DRM_IOCTL_MODE_MAP_DUMB    DRM_IOWR(0xB3, struct drm_mode_map_dumb)
#define DRM_IOCTL_MODE_DESTROY_DUMB    DRM_IOWR(0xB4, struct drm_mode_destroy_dumb)
#define DRM_IOCTL_MODE_GETPLANERESOURCES DRM_IOWR(0xB5, struct drm_mode_get_plane_res)
#define DRM_IOCTL_MODE_GETPLANE	DRM_IOWR(0xB6, struct drm_mode_get_plane)
#define DRM_IOCTL_MODE_SETPLANE	DRM_IOWR(0xB7, struct drm_mode_set_plane)
#define DRM_IOCTL_MODE_ADDFB2		DRM_IOWR(0xB8, struct drm_mode_fb_cmd2)
#define DRM_IOCTL_MODE_OBJ_GETPROPERTIES	DRM_IOWR(0xB9, struct drm_mode_obj_get_properties)
#define DRM_IOCTL_MODE_OBJ_SETPROPERTY	DRM_IOWR(0xBA, struct drm_mode_obj_set_property)
*/
/**
 * Device specific ioctls should only be in their respective headers
 * The device specific ioctl range is from 0x40 to 0x99.
 * Generic IOCTLS restart at 0xA0.
 *
 * \sa drmCommandNone(), drmCommandRead(), drmCommandWrite(), and
 * drmCommandReadWrite().
 */
/*
#define DRM_COMMAND_BASE                0x40
#define DRM_COMMAND_END			0xA0
*/
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
/*
struct drm_event {
	__u32 type
	__u32 length
}

#define DRM_EVENT_VBLANK 0x01
#define DRM_EVENT_FLIP_COMPLETE 0x02

struct drm_event_vblank {
	struct drm_event base
	__u64 user_data
	__u32 tv_sec
	__u32 tv_usec
	__u32 sequence
	__u32 reserved
}

#define DRM_CAP_DUMB_BUFFER 0x1
#define DRM_CAP_VBLANK_HIGH_CRTC   0x2
#define DRM_CAP_DUMB_PREFERRED_DEPTH 0x3
#define DRM_CAP_DUMB_PREFER_SHADOW 0x4
#define DRM_CAP_PRIME 0x5

#define DRM_PRIME_CAP_IMPORT 0x1
#define DRM_PRIME_CAP_EXPORT 0x2
*/
/* typedef area */
/*
typedef struct drm_clip_rect drm_clip_rect_t
typedef struct drm_drawable_info drm_drawable_info_t
typedef struct drm_tex_region drm_tex_region_t
typedef struct drm_hw_lock drm_hw_lock_t
typedef struct drm_version drm_version_t
typedef struct drm_unique drm_unique_t
typedef struct drm_list drm_list_t
typedef struct drm_block drm_block_t
typedef struct drm_control drm_control_t
typedef enum drm_map_type drm_map_type_t
typedef enum drm_map_flags drm_map_flags_t
typedef struct drm_ctx_priv_map drm_ctx_priv_map_t
typedef struct drm_map drm_map_t
typedef struct drm_client drm_client_t
typedef enum drm_stat_type drm_stat_type_t
typedef struct drm_stats drm_stats_t
typedef enum drm_lock_flags drm_lock_flags_t
typedef struct drm_lock drm_lock_t
typedef enum drm_dma_flags drm_dma_flags_t
typedef struct drm_buf_desc drm_buf_desc_t
typedef struct drm_buf_info drm_buf_info_t
typedef struct drm_buf_free drm_buf_free_t
typedef struct drm_buf_pub drm_buf_pub_t
typedef struct drm_buf_map drm_buf_map_t
typedef struct drm_dma drm_dma_t
typedef union drm_wait_vblank drm_wait_vblank_t
typedef struct drm_agp_mode drm_agp_mode_t
typedef enum drm_ctx_flags drm_ctx_flags_t
typedef struct drm_ctx drm_ctx_t
typedef struct drm_ctx_res drm_ctx_res_t
typedef struct drm_draw drm_draw_t
typedef struct drm_update_draw drm_update_draw_t
typedef struct drm_auth drm_auth_t
typedef struct drm_irq_busid drm_irq_busid_t
typedef enum drm_vblank_seq_type drm_vblank_seq_type_t

typedef struct drm_agp_buffer drm_agp_buffer_t
typedef struct drm_agp_binding drm_agp_binding_t
typedef struct drm_agp_info drm_agp_info_t
typedef struct drm_scatter_gather drm_scatter_gather_t
typedef struct drm_set_version drm_set_version_t
*/
