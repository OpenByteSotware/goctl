package goctl

import "unsafe"
import "testing"

/*
struct winsize
{
  ws_row uint16	
  ws_col uint16	
  ws_xpixel uint16	
  ws_ypixel uint16	
}
*/

//func MY_IOCTL_COMMAND() uint32 { return _IOR('k', 1, int) }

func TestNSUser(t *testing.T) {
    t.Log("Say bye")
    /*
    result := Add(1, 2)
	if result != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", result)
	}
    */
}

func TestNSIOCTL(t *testing.T) {
    const NSIO = 0xb7
	var NS_GET_USERNS = _IO(NSIO, 0x1)
    var val = 5
     varPtr := unsafe.Pointer(&val)
    uptr := uintptr(varPtr)
    
    netnsf := goctlOpenDevice("/proc/self/ns/net")
	defer goctlCloseDevice(int(netnsf))
	goctlGetValue(int(netnsf), NS_GET_USERNS, uptr)
    t.Logf("value = %d", val)
    t.Log("done test")
 }
