package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	dll := syscall.NewLazyDLL("SDL2-2_0_9-386.dll")
	addHintCallback := dll.NewProc("SDL_AddHintCallback")
	setHint := dll.NewProc("SDL_SetHint")
	hintCallbackPtr := syscall.NewCallback(theHintCallback)
	name := append([]byte("SDL_RENDER_SCALE_QUALITY"), 0)
	value := append([]byte("1"), 0)
	addHintCallback.Call(uintptr(unsafe.Pointer(&name[0])), hintCallbackPtr, 0)
	setHint.Call(uintptr(unsafe.Pointer(&name[0])), uintptr(unsafe.Pointer(&value[0])))
}

func theHintCallback(userdata, name, oldValue, newValue uintptr) uintptr {
	fmt.Println("theHintCallback", userdata, name, oldValue, newValue)
	return 0
}
