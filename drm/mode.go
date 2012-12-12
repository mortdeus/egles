package drm

const (
	DISPLAY_INFO_LEN   = 32
	CONNECTOR_NAME_LEN = 32
	DISPLAY_MODE_LEN   = 32
	PROP_NAME_LEN      = 32

	MODE_TYPE_BUILTIN   = 1 << 0
	MODE_TYPE_CLOCK_C   = (1 << 1) | MODE_TYPE_BUILTIN
	MODE_TYPE_CRTC_C    = (1 << 2) | MODE_TYPE_BUILTIN
	MODE_TYPE_PREFERRED = 1 << 3
	MODE_TYPE_DEFAULT   = 1 << 4
	MODE_TYPE_USERDEF   = 1 << 5
	MODE_TYPE_DRIVER    = 1 << 6

	/* Video mode flags */

	// bit compatible with the xorg definitions. 
	MODE_FLAG_PHSYNC    = 1 << 0
	MODE_FLAG_NHSYNC    = 1 << 1
	MODE_FLAG_PVSYNC    = 1 << 2
	MODE_FLAG_NVSYNC    = 1 << 3
	MODE_FLAG_INTERLACE = 1 << 4
	MODE_FLAG_DBLSCAN   = 1 << 5
	MODE_FLAG_CSYNC     = 1 << 6
	MODE_FLAG_PCSYNC    = 1 << 7
	MODE_FLAG_NCSYNC    = 1 << 8

	// hskew provided 
	MODE_FLAG_HSKEW   = 1 << 9
	MODE_FLAG_BCAST   = 1 << 10
	MODE_FLAG_PIXMUX  = 1 << 11
	MODE_FLAG_DBLCLK  = 1 << 12
	MODE_FLAG_CLKDIV2 = 1 << 13

	/* DPMS flags */
	// bit compatible with the xorg definitions. 
	MODE_DPMS_ON      = 0
	MODE_DPMS_STANDBY = 1
	MODE_DPMS_SUSPEND = 2
	MODE_DPMS_OFF     = 3

	/* Scaling mode options */
	// Unmodified timing (display or software can still scale) 
	MODE_SCALE_NONE = 0
	// Full screen, ignore aspect 
	MODE_SCALE_FULLSCREEN = 1
	// Centered, no scaling 
	MODE_SCALE_CENTER = 2
	// Full screen, preserve aspect 
	MODE_SCALE_ASPECT = 3

	/* Dithering mode options */
	MODE_DITHERING_OFF  = 0
	MODE_DITHERING_ON   = 1
	MODE_DITHERING_AUTO = 2

	/* Dirty info options */
	MODE_DIRTY_OFF      = 0
	MODE_DIRTY_ON       = 1
	MODE_DIRTY_ANNOTATE = 2
)

type ModeModeInfo struct {
	Clock uint32
	Hdisplay, HsyncStart, hsyncEnd, Htotal, Hskew,
	Vdisplay, VsyncStart, VsyncEnd, Vtotal, Vscan uint16
	Vrefresh,
	Flags,
	Type uint32
	Name [DISPLAY_MODE_LEN]byte
}

type ModeCardRes struct {
	FbIdPtr,
	CrtcIdPtr,
	ConnectorIdPtr,
	EncoderIdPtr uint64
	CountFbs,
	CountCrtcs,
	CountConnectors,
	CountEncoders,
	MinWidth, MaxWidth,
	MinHeight, MaxHeight uint32
}

type ModeCrtc struct {
	SetConnectorsPtr uint64
	CountConnectors,
	CrtcId,
	// Id of framebuffer 
	FbId,
	// Position on the framebuffer 
	X, Y,
	GammaSize,
	ModeValid uint32
	mode ModeModeInfo
}

const (
	MODE_PRESENT_TOP_FIELD = 1 << iota
	MODE_PRESENT_BOTTOM_FIELD
)

/* Planes blend with or override other bits on the CRTC */
type ModeSetPlane struct {
	PlaneId,
	CrtcId,
	// fb object contains surface format type 
	FbId,
	Flags uint32

	// Signed dest location allows it to be partially off screen 
	CrtcX, CrtcY int32
	CrtcW, CrtcH,

	// Source values are 16.16 fixed point
	SrcX, SrcY,
	SrcH, SrcW uint32
}

type ModeGetPlane struct {
	PlaneId,
	CrtcId,
	FbId,
	PossibleCrtcs,
	GammaSize,
	CountFormatTypes uint32
	FormatTypePtr uint64
}

type ModeGetPlaneRes struct {
	PlaneIdPtr  uint64
	CountPlanes uint32
}

const (
	MODE_ENCODER_NONE  = 0
	MODE_ENCODER_DAC   = 1
	MODE_ENCODER_TMDS  = 2
	MODE_ENCODER_LVDS  = 3
	MODE_ENCODER_TVDAC = 4
)

type ModeGetEncoder struct {
	EncoderId,
	EncoderType,
	//Id of crtc
	CrtcId,
	PossibleCrtcs,
	PossibleClones uint32
}

/* 
This is for connectors with multiple signal types. 
Try to match DRM_MODE_CONNECTOR_X as closely as possible. 
*/
const (
	MODE_SUBCONNECTOR_Automatic = 0
	MODE_SUBCONNECTOR_Unknown   = 0
	MODE_SUBCONNECTOR_DVID      = 3
	MODE_SUBCONNECTOR_DVIA      = 4
	MODE_SUBCONNECTOR_Composite = 5
	MODE_SUBCONNECTOR_SVIDEO    = 6
	MODE_SUBCONNECTOR_Component = 8
	MODE_SUBCONNECTOR_SCART     = 9

	MODE_CONNECTOR_Unknown     = 0
	MODE_CONNECTOR_VGA         = 1
	MODE_CONNECTOR_DVII        = 2
	MODE_CONNECTOR_DVID        = 3
	MODE_CONNECTOR_DVIA        = 4
	MODE_CONNECTOR_Composite   = 5
	MODE_CONNECTOR_SVIDEO      = 6
	MODE_CONNECTOR_LVDS        = 7
	MODE_CONNECTOR_Component   = 8
	MODE_CONNECTOR_9PinDIN     = 9
	MODE_CONNECTOR_DisplayPort = 10
	MODE_CONNECTOR_HDMIA       = 11
	MODE_CONNECTOR_HDMIB       = 12
	MODE_CONNECTOR_TV          = 13
	MODE_CONNECTOR_eDP         = 14
)

type ModeGetConnector struct {
	EncodersPtr   uint64
	ModesPtr      uint64
	PropsPtr      uint64
	PropValuesPtr uint64

	CountModes,
	CountProps,
	CountEncoders,
	// Current Encoder */
	EncoderId,
	// Id */
	ConnectorId,
	ConnectorType,
	ConnectorTypeId,

	Connection,
	// HxW in millimeters 
	MmWidth, MmHeight,
	Subpixel uint32
}

const (
	MODE_PROP_PENDING   = 1 << 0
	MODE_PROP_RANGE     = 1 << 1
	MODE_PROP_IMMUTABLE = 1 << 2
	// enumerated type with text strings */
	MODE_PROP_ENUM = 1 << 3
	MODE_PROP_BLOB = 1 << 4
	// bitmask of enumerated types */
	MODE_PROP_BITMASK = 1 << 5
)

type ModePropertyEnum struct {
	Value uint64
	Name  [PROP_NAME_LEN]byte
}

type ModeGetProperty struct {
	// values and blob lengths 
	ValuesPtr,
	// enum and blob id ptrs 
	EnumBlobPtr uint64
	PropId,
	Flags uint32
	Name [PROP_NAME_LEN]byte
	CountValues,
	CountEnumBlobs uint32
}

type ModeConnectorSetProperty struct {
	Value uint64
	PropId,
	ConnectorId uint32
}

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

type ModeObjGetProperties struct {
	PropsPtr      uint64
	PropValuesPtr uint64
	CountProps,
	ObjId,
	ObjType uint32
}

type ModeObjSetProperty struct {
	Value uint64
	PropId,
	ObjId,
	ObjType uint32
}

type ModeGetBlob struct {
	BlobId,
	Length uint32
	Data uint64
}

type ModeFbCmd struct {
	FbId,
	Width, Height,
	Pitch,
	Bpp,
	Depth,
	// driver specific handle */
	Handle uint32
}

// for interlaced framebuffers */
const MODE_FB_INTERLACED = 1 << 0

type ModeFbCmd2 struct {
	FbId,
	Width, Height,
	// fourcc code from drm_fourcc.h */
	PixelFormat,
	Flags uint32

	Handles,
	// pitch for each plane */
	Pitches,
	// offset of each plane */
	Offsets [4]uint32
}

const (
	MODE_FB_DIRTY_ANNOTATE_COPY = 0x01
	MODE_FB_DIRTY_ANNOTATE_FILL = 0x02
	MODE_FB_DIRTY_FLAGS         = 0x03
)

type ModeFbDirtyCmd struct {
	FbId,
	Flags,
	Color,
	NumClips uint32
	ClipsPtr uint64
}

type ModeModeCmd struct {
	ConnectorId uint32
	Mode        ModeModeInfo
}

const (
	MODE_CURSOR_BO   = (1 << 0)
	MODE_CURSOR_MOVE = (1 << 1)
)

/* 
depending on the value in flags diffrent members are used.
CURSOR_BO uses crtc width height handle - if 0 turns the cursor of
CURSOR_MOVE uses crtc x y
*/
type ModeCursor struct {
	Flags,
	CrtcId uint32
	X,
	Y int32
	Width,
	Height,
	// driver specific handle */
	Handle uint32
}

type ModeCrtcLut struct {
	CrtcId,
	GammaSize uint32

	// pointers to arrays */
	Red   uint64
	Green uint64
	Blue  uint64
}

const (
	MODE_PAGE_FLIP_EVENT = 0x01
	MODE_PAGE_FLIP_FLAGS = MODE_PAGE_FLIP_EVENT
)

type ModeCrtcPageFlip struct {
	CrtcId,
	FbId,
	Flags,
	Reserved uint32
	UserData uint64
}

// create a dumb scanout buffer */
type ModeCreateDumb struct {
	Height,
	Width,
	Bpp,
	Flags,
	// handle, pitch, size will be returned */
	Handle,
	Pitch uint32
	Size uint64
}

// set up for mmap of a dumb scanout buffer */
type ModeMapDumb struct {
	// Handle for the object being mapped. 
	Handle,
	Pad uint32
	// Fake offset to use for subsequent mmap call 
	// This is a fixed-size type for 32/64 compatibility.
	Offset uint64
}

type ModeDestroyDumb struct {
	Handle uint32
}
