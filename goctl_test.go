package goctl

import (
//import "unsafe"
"testing"
"os"
"golang.org/x/sys/unix"
"syscall"
//import "net"
"reflect"
)


// make odd number
// in go all struct data is always even # of bytes.  It will add padding byte if not even number
type MixedData struct {
	v1 uint8
	v2 uint16
	v3 uint32
    v4 uint64
}

type winsize struct {
    	ws_row    uint16
    	ws_col    uint16
    	ws_xpixel uint16
    	ws_ypixel uint16
    }

func Test_IOC(t *testing.T) {
    // test zero
    val := _IOC(0, 0, 0, 0)

    if val != 0 {
        t.Errorf("_IOC(0, 0, 0, 0) = %d, want 0", val)
    }

    val = _IOC(1, 0, 0, 0)

    if val != 0x40000000 {
        t.Errorf("_IOC(1, 0, 0, 0) = %x, want 0x40000000", val)
    }

    val = _IOC(0, 1, 0, 0)

    if val != 0x100 {
        t.Errorf("_IOC(0, 1, 0, 0) = %x, want 0x100", val)
    }

   val = _IOC(0, 0, 1, 0)

    if val != 0x1 {
        t.Errorf("_IOC(0, 0, 1, 0) = %x, want 0x1", val)
    }

   val = _IOC(0, 0, 0, 1)

    if val != 0x010000 {
        t.Errorf("_IOC(0, 0, 0, 1) = %x, want 0x010000", val)
    }
}

//func MY_IOCTL_COMMAND() uint32 { return _IOR('k', 1, int) }
func Test_Struct_IOC_TYPECHECK(t *testing.T) {
    data := MixedData {}
    size := _IOC_TYPECHECK(data)

    // size should be 15, but it pads with to make even
    if size != (15 + 1) {
        t.Errorf("_IOC_TYPECHECK(MixedData) = %d; want 16", size)
    }
}

func Test_byte_IOC_TYPECHECK(t *testing.T) {
    data := MixedData {}
    size := _IOC_TYPECHECK(data.v1)

    // size should be 15, but it pads with to make even
    if size != (1) {
        t.Errorf("_IOC_TYPECHECK(MixedData) = %d; want 1", size)
    }
}

func TestNSUser(t *testing.T) {
    t.Log("Say bye")
    /*
    result := Add(1, 2)
	if result != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", result)
	}
    */
}

/*
// this is from the nsfs.h linux header file
#define NSIO	0xb7

// Returns a file descriptor that refers to an owning user namespace
#define NS_GET_USERNS		_IO(NSIO, 0x1)
// Returns a file descriptor that refers to a parent namespace
#define NS_GET_PARENT		_IO(NSIO, 0x2)
// Returns the type of namespace (CLONE_NEW* value) referred to by
   file descriptor 
#define NS_GET_NSTYPE		_IO(NSIO, 0x3)
// Get owner UID (in the caller's user namespace) for a user namespace
#define NS_GET_OWNER_UID	_IO(NSIO, 0x4)
*/
func TestNSIOCTL(t *testing.T) {
    const NSIO = 0xb7
    val := int(0x100)
    var NS_GET_OWNER_UID = IO(NSIO, 0x4)
    
    netnsf, err := GoctlOpenDevice("/proc/self/ns/user")

    if err != nil {
        t.Error("cannot opne device")
    }

	defer GoctlCloseDevice(netnsf)
	err = GoctlGetValue(netnsf, NS_GET_OWNER_UID, &val)
    if err != nil {
        t.Error("cannot get value from device")
    }
    t.Logf("value = %d", val)
    t.Log("done test")
 }

 func TestWindowSize(t *testing.T) {
    	ws := winsize{}
        ws.ws_row = 0
    	ws.ws_col = 0
    	ws.ws_xpixel = 0
    	ws.ws_ypixel = 0

        goctlwinStruct := GoctlStruct[winsize]{}
        goctlwinStruct.Value = &ws

        stdoutFd := int(os.Stdout.Fd())
        // Pass a pointer to the intStruct to the generic function
        err := GoctlGetStruct(stdoutFd, unix.TIOCGWINSZ, &goctlwinStruct)
        if err != nil {
            t.Error("cannot get value from device")
        } 
        
        t.Logf("ws_row=%d", ws.ws_row)
        t.Logf("ws_col=%d", ws.ws_col)
        t.Logf("ws_xpixel=%d", ws.ws_xpixel)
        t.Logf("ws_ypixel=%d", ws.ws_ypixel)
 }

 func displayObjectVariables(obj interface{}, t *testing.T) {
	v := reflect.ValueOf(obj)

	// If it's a pointer, get the element it points to
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		t.Logf("Cannot display variables for non-struct type: %s\n", v.Kind())
		return
	}

	t.Logf("Variables of type %s:\n", v.Type())
 
    for i := 0; i < v.NumField(); i++ {
        field := v.Type().Field(i)
        //if field.CanSet() {
            fieldValue := v.Field(i)
            t.Logf(" %s: %v (Type: %s)\n", field.Name, fieldValue.Interface(), field.Type)
        //}
    }
 }

 func TestNetData(t *testing.T) {
    ifaceName := "eth0" 
    // Create a socket
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)

	if err != nil {
		t.Error("cannot get value from device")
	}

	defer syscall.Close(fd)

	// Prepare ifreq struct
	ifr, err := unix.NewIfreq(ifaceName)
    //ifr.ifr.mtu = 0;
	//copy(ifr.ifr_name[:], []byte(ifaceName))

	if err != nil {
		t.Error("Error creating Ifreq")
	}
	
    goctlStruct := GoctlStruct[unix.Ifreq]{}
    goctlStruct.Value = ifr

	err = GoctlGetStruct(fd, syscall.SIOCGIFFLAGS, &goctlStruct)

	if err != nil {
		t.Errorf("ioctl SIOCGIFMTU failed:%s ", err)
	}

    flags := ifr.Uint16()
    t.Logf("Interface %s flags: 0x%x\n", ifaceName, flags)

    //var byteArray [5]byte

    ipAddr, _ := ifr.Inet4Addr()
    for _, n := range ipAddr {
        t.Logf("interface ip=%08b", n)
	//	fmt.Printf("%08b ", n) // Prints each byte as an 8-bit binary string
	}

    //displayObjectVariables(ifr, t)
	// Extract MTU value
	//mtu := int(ifr.raw.Ifrn)
    
    //t.Logf("mtu=%d", mtu)	
 }
