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

type Goctl struct {
        FD int
}

func goctlOpenDevice(deviceName string) uintptr {
	if deviceName == "" {
		deviceName = "/dev/crypto" // Default value
	}

	file, err := os.Open(deviceName)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}

	fmt.Printf("Hello, %s!\n", deviceName)

	return file.Fd()
}

func goctlCloseDevice(fd int) {
	unix.Close(fd)
	
}

func goctlPtr(fd int, num uint32, arg unsafe.Pointer) (error) {
	_, _, err := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), uintptr(num), uintptr(arg))

	if err != 0 {
		return syscall.Errno(err)
	}
	
	return nil
}

func goctlGetValue(fd int, num uint32, arg uintptr) (error) {
	_, _, err := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), uintptr(num), uintptr(arg))
	
	if err != 0 {
		return syscall.Errno(err)
	}

	return nil
}
/*
func RetFd(fd int, request uint) (int, error) {
	nsfd, _, errno := unix.Syscall(unix.SYS_IOCTL,
		uintptr(fd), uintptr(request), uintptr(0))
	if errno != 0 {
		return -1, errors.New(errno.Error())
	}
	return int(nsfd), nil
}
*/