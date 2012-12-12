package ioctl

import (
	"syscall"
	"unsafe"
)

/* ioctl command encoding: 32 bits total, command in lower 16 bits,
 * size of the parameter structure in the lower 14 bits of the
 * upper 16 bits.
 * Encoding the size of the parameter structure in the ioctl request
 * is useful for catching programs compiled with old versions
 * and to avoid overwriting user space outside the user buffer area.
 * The highest 2 bits are reserved for indicating the ``access mode''.
 * NOTE: This limits the max parameter size to 16kB -1 !
 */
func Call(fd uintptr, name int, data unsafe.Pointer) syscall.Errno {
	_, _, err := syscall.RawSyscall(syscall.SYS_IOCTL, fd, uintptr(name), uintptr(data))
	return err
}

/*
 * The following is for compatibility across the various Linux
 * platforms.  The generic ioctl numbering scheme doesn't really enforce
 * a type field.  De facto, however, the top 8 bits of the lower 16
 * bits are indeed used as a type field, so we might just as well make
 * this explicit here.  Please be sure to use the decoding macros
 * below from now on.
 */
const (
	NRBITS   = 8
	TYPEBITS = 8

	/*
	 * Let any architecture override either of the following before
	 * including this file.
	 */

	SIZEBITS = 14

	DIRBITS = 2

	NRMASK   = (1 << NRBITS) - 1
	TYPEMASK = (1 << TYPEBITS) - 1
	SIZEMASK = (1 << SIZEBITS) - 1
	DIRMASK  = (1 << DIRBITS) - 1

	NRSHIFT   = 0
	TYPESHIFT = NRSHIFT + NRBITS
	SIZESHIFT = TYPESHIFT + TYPEBITS
	DIRSHIFT  = SIZESHIFT + SIZEBITS

	/*
	 * Direction bits, which any architecture can choose to override
	 * before including this file.
	 */

	NONE = 0x0

	WRITE = 0x1

	READ = 0x2

	/* ...and for the drivers/sound files... */
	IN         = WRITE << DIRSHIFT
	OUT        = READ << DIRSHIFT
	INOUT      = (WRITE | READ) << DIRSHIFT
	SIZE_MASK  = SIZEMASK << SIZESHIFT
	SIZE_SHIFT = SIZESHIFT
)

func IOC(dir, _type, nr, size int32) int32 {
	return (dir << DIRSHIFT) | (_type << TYPESHIFT) |
		(nr << NRSHIFT) | (size << SIZESHIFT)
}

/* used to create numbers */
func IO(_type, nr int32) int32 {
	return IOC(NONE, _type, nr, 0)
}
func IOR(_type, nr, size int32) int32 {
	return IOC(READ, _type, nr, size)
}
func IOW(_type, nr, size int32) int32 {
	return IOC(WRITE, _type, nr, size)
}
func IOWR(_type, nr, size int32) int32 {
	return IOC((READ | WRITE), _type, nr, size)
}
func IOR_BAD(_type, nr, size int32) int32 {
	return IOC(READ, _type, nr, size)
}
func IOW_BAD(_type, nr, size int32) int32 {
	return IOC(WRITE, _type, nr, size)
}
func IOWR_BAD(_type, nr, size int32) int32 {
	return IOC(READ|WRITE, _type, nr, size)

}

/* used to decode ioctl numbers.. */
func DIR(nr int32) int32 {
	return ((nr >> DIRSHIFT) & DIRMASK)
}
func TYPE(nr int32) int32 {
	return ((nr >> TYPESHIFT) & TYPEMASK)
}
func NR(nr int32) int32 {
	return ((nr >> NRSHIFT) & NRMASK)
}
func SIZE(nr int32) int32 {
	return ((nr >> SIZESHIFT) & SIZEMASK)
}
