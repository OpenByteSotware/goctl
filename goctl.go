package goctl

import (
	"fmt"
	//"bytes"
	//"flag"
	//"log"
	"os"
	"syscall"
	"unsafe"
	"golang.org/x/sys/unix"
)
// from https://github.com/torvalds/linux/blob/master/include/uapi/asm-generic/ioctl.h
// 

func goctlOpenDevice(deviceName string) (int, error) {
	if deviceName == "" {
		deviceName = "/dev/crypto" // Default value
	}

	file, err := os.Open(deviceName)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, err
	}

	fmt.Printf("Hello, %s!\n", deviceName)

	fd := int(file.Fd())
	return fd, err
}

func goctlCloseDevice(fd int) {
	unix.Close(fd)
	
}

func goctlGetStruct[T any](fd int, num uint32, s *GoctlStruct[T]) (error) {
	// the hidden trick to safely use pointer that will not get garbage collected until not used anymore
	// since calls are not atomic, convert to unsafe pointer in this function to system call
	// then after system call data will be stored in s
	p := unsafe.Pointer(s.Value)
	_, _, err := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), uintptr(num), uintptr(p))
	
	if err != 0 {
		return syscall.Errno(err)
	}

	return nil
}

/*
func _IOR[T any](typ uint32, nr uint32, argtype T) uint32 { return _IOC(_IOC_READ, typ, nr, _IOC_TYPECHECK(argtype)) }
*/

func goctlGetValue(fd int, num uint32, retVal *int) (error) {
	p := unsafe.Pointer(retVal)
	_, _, err := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), uintptr(num), uintptr(p))
	
	if err != 0 {
		return syscall.Errno(err)
	}

	return nil
}
