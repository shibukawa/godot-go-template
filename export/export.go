package export

import "C"
import (
	"godottemplate"
	"unsafe"

	"github.com/godot-go/godot-go/pkg/gdnative"
)

//export godot_gdnative_init
func godot_gdnative_init(options unsafe.Pointer) {
	gdnative.GodotGdnativeInit((*gdnative.GdnativeInitOptions)(options))
}

//export godot_gdnative_terminate
func godot_gdnative_terminate(options unsafe.Pointer) {
	gdnative.GodotGdnativeTerminate((*gdnative.GdnativeTerminateOptions)(options))
}

//export godot_nativescript_init
func godot_nativescript_init(handle unsafe.Pointer) {
	gdnative.GodotNativescriptInit(handle)

	godottemplate.MyCounterNativescriptInit()
	gdnative.RegisterClass(&godottemplate.MyCounter{}, godottemplate.MycounterCreateFunc)
}

//export godot_nativescript_terminate
func godot_nativescript_terminate(handle unsafe.Pointer) {
	godottemplate.MyCounterNativescriptTerminate()

	gdnative.GodotNativescriptTerminate(handle)
}
