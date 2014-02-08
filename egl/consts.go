package egl

/*
#include <EGL/egl.h>
#include <EGL/eglplatform.h>
*/
import "C"

/* EGL Versioning */
const (
	VERSION_1_0 = 1
	VERSION_1_1 = 1
	VERSION_1_2 = 1
	VERSION_1_3 = 1
	VERSION_1_4 = 1
)

// Out-of-band attribute value
const (
	DONT_CARE = -1
	UNKNOWN   = -1

	/*
		Constant scale factor by which fractional
		display resolutions & aspect ratio are scaled
		when queried as integer values.
	*/
	DISPLAY_SCALING = 10000
)

/* EGL Enumerants. Bitmasks and other exceptional cases aside, most
 * enums are assigned unique values starting at 0x3000.
 */

type Error int

// Errors / GetError return values
const (
	SUCCESS Error = 0x3000 + iota
	NOT_INITIALIZED
	BAD_ACCESS
	BAD_ALLOC
	BAD_ATTRIBUTE
	BAD_CONFIG
	BAD_CONTEXT
	BAD_CURRENT_SURFACE
	BAD_DISPLAY
	BAD_MATCH
	BAD_NATIVE_PIXMAP
	BAD_NATIVE_WINDOW
	BAD_PARAMETER
	BAD_SURFACE
	CONTEXT_LOST //EGL 1.1 - IMG_power_management

	/*
		Reserved= 0x300F=-0x301F
		for additional errors
	*/
)

func (e Error) Error() string {
	pre := "EGLError: "
	switch e {
	case SUCCESS:
		return pre + "success"
	case NOT_INITIALIZED:
		return pre + "not initialized"
	case BAD_ACCESS:
		return pre + "bad access"
	case BAD_ALLOC:
		return pre + "bad alloc"
	case BAD_ATTRIBUTE:
		return pre + "bad attribute"
	case BAD_CONFIG:
		return pre + "bad config"
	case BAD_CONTEXT:
		return pre + "bad context"
	case BAD_CURRENT_SURFACE:
		return pre + "bad current surface"
	case BAD_DISPLAY:
		return pre + "bad display"
	case BAD_MATCH:
		return pre + "bad match"
	case BAD_NATIVE_PIXMAP:
		return pre + "bad_native pixmap"
	case BAD_NATIVE_WINDOW:
		return pre + "bad native window"
	case BAD_PARAMETER:
		return pre + "bad parameter"
	case BAD_SURFACE:
		return pre + "bad surface"
	case CONTEXT_LOST:
		return pre + "context lost"
	default:
		return ""
	}
}

/* Config attributes */
const (
	BUFFER_SIZE = 0x3020 + iota
	ALPHA_SIZE
	BLUE_SIZE
	GREEN_SIZE
	RED_SIZE
	DEPTH_SIZE
	STENCIL_SIZE
	CONFIG_CAVEAT
	CONFIG_ID
	LEVEL
	MAX_PBUFFER_HEIGHT
	MAX_PBUFFER_PIXELS
	MAX_PBUFFER_WIDTH
	NATIVE_RENDERABLE
	NATIVE_VISUAL_ID
	NATIVE_VISUAL_TYPE
	_
	SAMPLES
	SAMPLE_BUFFERS
	SURFACE_TYPE
	TRANSPARENT_TYPE
	TRANSPARENT_BLUE_VALUE
	TRANSPARENT_GREEN_VALUE
	TRANSPARENT_RED_VALUE
	// Attrib list terminator
	NONE
	BIND_TO_TEXTURE_RGB
	BIND_TO_TEXTURE_RGBA
	MIN_SWAP_INTERVAL
	MAX_SWAP_INTERVAL
	LUMINANCE_SIZE
	ALPHA_MASK_SIZE
	COLOR_BUFFER_TYPE
	RENDERABLE_TYPE
	// Pseudo-attribute (not queryable)
	MATCH_NATIVE_PIXMAP
	CONFORMANT

	/*
		Reserved= 0x3041=-0x304F for
		additional config attributes
	*/
)

/* SURFACE_TYPE mask bits */
const (
	PBUFFER_BIT = 0x0001 << iota
	PIXMAP_BIT
	WINDOW_BIT
	_
	_
	VG_COLORSPACE_LINEAR_BIT
	VG_ALPHA_FORMAT_PRE_BIT
	_
	_
	MULTISAMPLE_RESOLVE_BOX_BIT
	SWAP_BEHAVIOR_PRESERVED_BIT
)

/* RENDERABLE_TYPE mask bits */
const (
	OPENGL_ES_BIT = 0x0001 << iota
	OPENVG_BIT
	OPENGL_ES2_BIT
	OPENGL_BIT
)
const (
	// CONFIG_CAVEAT values
	SLOW_CONFIG = 0x3050 + iota
	NON_CONFORMANT_CONFIG

	// TRANSPARENT_TYPE value
	TRANSPARENT_RGB
	/* QueryString targets */
	VENDOR
	VERSION
	EXTENSIONS

	/* QuerySurface & SurfaceAttrib & CreatePbufferSurface targets */
	HEIGHT
	WIDTH
	LARGEST_PBUFFER

	/* GetCurrentSurface targets */
	DRAW
	READ

	/* WaitNative engines */
	CORE_NATIVE_ENGINE

	/* Config attribute values for EGL_TEXTURE_FORMAT*/
	NO_TEXTURE
	TEXTURE_RGB
	TEXTURE_RGBA
	TEXTURE_2D
)
const (

	// more QuerySurface & SurfaceAttrib & CreatePbufferSurface targets
	TEXTURE_FORMAT = 0x3080 + iota
	TEXTURE_TARGET
	MIPMAP_TEXTURE
	MIPMAP_LEVEL

	/* BindTexImage & ReleaseTexImage buffer targets */
	BACK_BUFFER
	SINGLE_BUFFER
	RENDER_BUFFER

	/* OpenVG color spaces and alpha formats */
	VG_COLORSPACE
	VG_ALPHA_FORMAT
	VG_COLORSPACE_sRGB
	VG_COLORSPACE_LINEAR
	VG_ALPHA_FORMAT_NONPRE
	VG_ALPHA_FORMAT_PRE

	/* QueryString target */
	CLIENT_APIS

	/* Config attribute values */
	RGB_BUFFER
	LUMINANCE_BUFFER

	HORIZONTAL_RESOLUTION
	VERTICAL_RESOLUTION
	PIXEL_ASPECT_RATIO
	SWAP_BEHAVIOR

	/* Back buffer swap behaviors */
	BUFFER_PRESERVED
	BUFFER_DESTROYED

	/* CreatePbufferFromClientBuffer buffer types */
	OPENVG_IMAGE

	/* QueryContext targets */
	CONTEXT_CLIENT_TYPE

	/* CreateContext attributes */
	CONTEXT_CLIENT_VERSION

	MULTISAMPLE_RESOLVE

	/* Multisample resolution behaviors */
	MULTISAMPLE_RESOLVE_DEFAULT
	MULTISAMPLE_RESOLVE_BOX
	_
	_
	_
	_
	/* BindAPI & QueryAPI targets */
	OPENGL_ES_API
	OPENVG_API
	OPENGL_API
)

/* EGL 1.2 tokens renamed for consistency in EGL 1.3 */
const (
	COLORSPACE          = VG_COLORSPACE
	ALPHA_FORMAT        = VG_ALPHA_FORMAT
	COLORSPACE_sRGB     = VG_COLORSPACE_sRGB
	COLORSPACE_LINEAR   = VG_COLORSPACE_LINEAR
	ALPHA_FORMAT_NONPRE = VG_ALPHA_FORMAT_NONPRE
	ALPHA_FORMAT_PRE    = VG_ALPHA_FORMAT_PRE
)
