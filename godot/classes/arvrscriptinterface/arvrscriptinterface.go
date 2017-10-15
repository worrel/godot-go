//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package arvrscriptinterface

/*
#include <stdio.h>
#include <stdlib.h>
#include <gdnative/gdnative.h>
#include <nativescript/godot_nativescript.h>
*/
import "C"

import (
	"github.com/shadowapex/godot-go/godot/classes/class"
	"log"
	"reflect"

	"github.com/shadowapex/godot-go/godot/classes/arvrinterface"
)

/*
   Undocumented
*/
type ARVRScriptInterface struct {
	arvrinterface.ARVRInterface
}

func (o *ARVRScriptInterface) baseClass() string {
	return "ARVRScriptInterface"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *ARVRScriptInterface) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *ARVRScriptInterface) getOwner() *C.godot_object {
	return o.owner
}

/*
   Undocumented
*/
func (o *ARVRScriptInterface) X_GetProjectionForEye() {
	log.Println("Calling ARVRScriptInterface.X_GetProjectionForEye()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_get_projection_for_eye", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ARVRScriptInterface) CommitForEye(eye int64, renderTarget *RID) {
	log.Println("Calling ARVRScriptInterface.CommitForEye()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(eye)
	goArguments[1] = reflect.ValueOf(renderTarget)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "commit_for_eye", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ARVRScriptInterface) GetRecommendedRenderTargetsize() *Vector2 {
	log.Println("Calling ARVRScriptInterface.GetRecommendedRenderTargetsize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_recommended_render_targetsize", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ARVRScriptInterface) GetTransformForEye(eye int64, camTransform *Transform) *Transform {
	log.Println("Calling ARVRScriptInterface.GetTransformForEye()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(eye)
	goArguments[1] = reflect.ValueOf(camTransform)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_transform_for_eye", goArguments, "*Transform")

	returnValue := goRet.Interface().(*Transform)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ARVRScriptInterface) HmdIsPresent() bool {
	log.Println("Calling ARVRScriptInterface.HmdIsPresent()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "hmd_is_present", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ARVRScriptInterface) Initialize() bool {
	log.Println("Calling ARVRScriptInterface.Initialize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "initialize", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ARVRScriptInterface) IsInitialized() bool {
	log.Println("Calling ARVRScriptInterface.IsInitialized()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_initialized", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ARVRScriptInterface) IsInstalled() bool {
	log.Println("Calling ARVRScriptInterface.IsInstalled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_installed", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ARVRScriptInterface) IsStereo() bool {
	log.Println("Calling ARVRScriptInterface.IsStereo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_stereo", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ARVRScriptInterface) Process() {
	log.Println("Calling ARVRScriptInterface.Process()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "process", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ARVRScriptInterface) SupportsHmd() bool {
	log.Println("Calling ARVRScriptInterface.SupportsHmd()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "supports_hmd", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ARVRScriptInterface) Uninitialize() {
	log.Println("Calling ARVRScriptInterface.Uninitialize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "uninitialize", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   ARVRScriptInterfaceImplementer is an interface for ARVRScriptInterface objects.
*/
type ARVRScriptInterfaceImplementer interface {
	class.Class
}
