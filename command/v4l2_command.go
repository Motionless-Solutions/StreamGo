package v4l2_command

import (
	sys "golang.org/x/sys/unix"
)

const (
	IocOpNone     = 0
	IocOpWrite    = 1
	IocOpRead     = 2
	IocTypeBits   = 8
	IocNumberBits = 8
	IocSizeBits   = 14
	IocOpBits     = 2
	NumberPos     = 0
	TypePos       = NumberPos + IocNumberBits
	SizePos       = TypePos + IocTypeBits
	OpPos         = SizePos + IocSizeBits
)

// Ioctl command encoding
func IoEnc(iocMode, iocType, number, size uintptr) uintptr {
	return (iocMode << OpPos) |
		(iocType << TypePos) |
		(number << NumberPos) |
		(size << SizePos)
}

// Convenience functions
func IoEncR(iocType, number, size uintptr) uintptr {
	return IoEnc(IocOpRead, iocType, number, size)
}

func IoEncW(iocType, number, size uintptr) uintptr {
	return IoEnc(IocOpWrite, iocType, number, size)
}

func IoEncRW(iocType, number, size uintptr) uintptr {
	return IoEnc(IocOpRead|IocOpWrite, iocType, number, size)
}

// Fourcc implements C v4l2_fourcc
func Fourcc(a, b, c, d uint32) uint32 {
	return (a | b<<8) | c<<16 | d<<24
}

func Ioctl(fd, req, arg uintptr) (uintptr, uintptr, error) {
	r1, r2, errno := sys.Syscall(sys.SYS_IOCTL, fd, req, arg)
	if errno != 0 {
		return 0, 0, errno
	}
	return r1, r2, nil
}
