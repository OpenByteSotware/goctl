// linux ioctl defines and function definitions explained here
// this file contains converted c header info to golang types, if not able to convert defines, then those are change
// to function calls
// https://www.man7.org/linux/man-pages/man2/ioctl.2.html
// https://docs.kernel.org/6.8/userspace-api/ioctl/ioctl-number.html

// https://docs.kernel.org/6.8/driver-api/ioctl.html

// ioctl magic numbers list
// https://www.kernel.org/doc/Documentation/ioctl/ioctl-number.txt
/*
_IO/_IOR/_IOW/_IOWR
The macro name specifies how the argument will be used. It may be a pointer to data to be passed into the kernel 
(_IOW), out of the kernel (_IOR), or both (_IOWR). _IO can indicate either commands with no argument or those passing 
an integer value instead of a pointer. It is recommended to only use _IO for commands without arguments, and use pointers for passing data.

type
An 8-bit number, often a character literal, specific to a subsystem or driver, and listed in Ioctl Numbers

nr
An 8-bit number identifying the specific command, unique for a give value of 'type'

data_type
The name of the data type pointed to by the argument, the command number encodes the sizeof(data_type) value in
a 13-bit or 14-bit integer, leading to a limit of 8191 bytes for the maximum size of the argument. Note: do not 
pass sizeof(data_type) type into _IOR/_IOW/IOWR, as that will lead to encoding sizeof(sizeof(data_type)), i.e. sizeof(size_t). 
_IO does not have a data_type parameter.
*/

//https://github.com/torvalds/linux/blob/master/include/uapi/asm-generic/ioctl.h

// Note: this is comments from header file for ioctl found above
/* ioctl command encoding: 32 bits total, command in lower 16 bits,
 * size of the parameter structure in the lower 14 bits of the
 * upper 16 bits.
 * Encoding the size of the parameter structure in the ioctl request
 * is useful for catching programs compiled with old versions
 * and to avoid overwriting user space outside the user buffer area.
 * The highest 2 bits are reserved for indicating the ``access mode''.
 * NOTE: This limits the max parameter size to 16kB -1 !
 */
/*
 * The following is for compatibility across the various Linux
 * platforms.  The generic ioctl numbering scheme doesn't really enforce
 * a type field.  De facto, however, the top 8 bits of the lower 16
 * bits are indeed used as a type field, so we might just as well make
 * this explicit here.  Please be sure to use the decoding macros
 * below from now on.
 */

package goctl

import (
	"reflect"
)

type GoctlStruct[T any] struct {
        Value *T
}

const (
	_IOC_NRBITS   = 8
	_IOC_TYPEBITS = 8
	_IOC_SIZEBITS = 14
	_IOC_DIRBITS = 2

	_IOC_NRSHIFT   = 0
	_IOC_TYPESHIFT = _IOC_NRSHIFT + _IOC_NRBITS
	_IOC_SIZESHIFT = _IOC_TYPESHIFT + _IOC_TYPEBITS
	_IOC_DIRSHIFT  = _IOC_SIZESHIFT + _IOC_SIZEBITS

	_IOC_NONE  = uint32(0)
	_IOC_WRITE = uint32(1)
	_IOC_READ  = uint32(2)

	_IOC_NRMASK = (1 << _IOC_NRBITS) - 1
	_IOC_TYPEMASK = (1 << _IOC_TYPEBITS) - 1

	_IOC_SIZEMASK = (1 << _IOC_SIZEBITS) - 1
	_IOC_DIRMASK = (1 << _IOC_DIRBITS) - 1

	IOC_IN = _IOC_WRITE << _IOC_DIRSHIFT
	IOC_OUT = _IOC_READ << _IOC_DIRSHIFT
	IOC_INOUT = (_IOC_WRITE|_IOC_READ) << _IOC_DIRSHIFT
	IOCSIZE_MASK = _IOC_SIZEMASK << _IOC_SIZESHIFT
	IOCSIZE_SHIFT = _IOC_SIZESHIFT
)

func _IOC(dir, typ, nr, size uint32) uint32 { 
	return ((dir << _IOC_DIRSHIFT) |
	(typ << _IOC_TYPESHIFT) |
    (nr << _IOC_NRSHIFT) |
    (size << _IOC_SIZESHIFT))
}

// sizeof doesn't exist in golang, use reflections to get size of type
func _IOC_TYPECHECK(data interface{}) uint32 { return uint32(reflect.TypeOf(data).Size()) }

func _IO(typ uint32, nr uint32) uint32 { return _IOC(_IOC_NONE, typ, nr, 0) }
func _IOR[T any](typ uint32, nr uint32, argtype T) uint32 { return _IOC(_IOC_READ, typ, nr, _IOC_TYPECHECK(argtype)) }
func _IOW[T any](typ uint32, nr uint32, argtype T) uint32 { return _IOC(_IOC_WRITE, typ, nr, _IOC_TYPECHECK(argtype)) }
func _IOWR[T any](typ uint32, nr uint32, argtype T) uint32 { return _IOC(_IOC_READ|_IOC_WRITE, typ, nr, _IOC_TYPECHECK(argtype)) }

func _IOC_DIR(nr uint32) uint32 { return (nr >> _IOC_DIRSHIFT) & _IOC_DIRMASK }
func _IOC_TYPE(nr uint32) uint32 { return (nr >> _IOC_TYPESHIFT) & _IOC_TYPEMASK }
func _IOC_NR(nr uint32) uint32 { return (nr >> _IOC_NRSHIFT) & _IOC_NRMASK }
func _IOC_SIZE(nr uint32) uint32 { return (nr >> _IOC_SIZESHIFT) & _IOC_SIZEMASK }
