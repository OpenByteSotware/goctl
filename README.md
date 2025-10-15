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
