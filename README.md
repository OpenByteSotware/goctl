# goctl
go ioctl wrapper

# Go IOCTL wrapper

[![Go Reference](<link>)

The `goctl` package implements interface to make safe calls to low level ioctl function.  This allows for access to /dev/<devices>.


## Disclaimer

A program directly or indirectly using this package needs to concider thread safty.  This package is not thread safe as it calls directly to the ioctl function.
But allows for ease of use in go programs.  

## Background

Some devices can be access from the kernel by using ioctrl.  Instead of cgo and c libraries.  This is a strictly go implementation of calling ioctl.  Which allows for static linking if needed.
With cgo the application must be dynamic linked

### Go FIPS compliance

> The status of goctl is not ready for production.  It is in development stage.

golang: https://github.com/golang/go/

## Features

### Building without go versions

