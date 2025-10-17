# goctl
go ioctl wrapper

# Go IOCTL wrapper

[![Go Reference](<link>)

The `goctl` package implements interface to make safe calls to low level ioctl function.  This allows for access to /dev/<devices>.

## Disclaimer

A program directly or indirectly using this package needs to concider thread safty.  This package is not thread safe as it calls directly to the ioctl function.
But allows for ease of use in go programs.  This is provided as is with no liabilities.  See license for details  

## Background

Some devices can be access from the kernel by using ioctrl.  Instead of using cgo and c libraries.  This is a strictly go implementation of calling ioctl.  Which allows for static linking if needed.
With cgo the application must be dynamic linked.
Source code for ioctl call is converted from linux header and c files to golang.  Should be able to access all ioctl devices.  This is for linux only.  This is not compatible with Windows.

## Example
To use goctl for structured data in ioctl you must incapsulate data struct in a GoctlStruct.
This makes use of reflection and ease of higher level call not concerned with pointers.
Just create a struct for the ioctl call then create the GoctlStruct like GoctlStruct[<your struct>]{}

Then pass the GoctlStruct instance by reference into the function goctlGetStruct
Note:  You must have the correct struct for the ioctl call command or you could get a panic.  This is because the structure needs to be correct size for the command.


```go
   	ws := winsize{}
        ws.ws_row = 0
    	ws.ws_col = 0
    	ws.ws_xpixel = 0
    	ws.ws_ypixel = 0

        goctlwinStruct := GoctlStruct[winsize]{}
        goctlwinStruct.Value = &ws

        stdoutFd := int(os.Stdout.Fd())
        // Pass a pointer to the intStruct to the generic function
        err := goctlGetStruct(stdoutFd, unix.TIOCGWINSZ, &goctlwinStruct)
        if err != nil {
            t.Error("cannot get value from device")
        } 

		
		