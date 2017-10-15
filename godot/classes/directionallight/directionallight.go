//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package directionallight

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

	"github.com/shadowapex/godot-go/godot/classes/light"
)

/*
   A DirectionalLight is a type of [Light] node that emits light constantly in one direction (the negative z axis of the node). It is used lights with strong intensity that are located far away from the scene to model sunlight or moonlight. The worldspace location of the DirectionalLight transform (origin) is ignored, only the basis is used do determine light direction.
*/
type DirectionalLight struct {
	light.Light
}

func (o *DirectionalLight) baseClass() string {
	return "DirectionalLight"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *DirectionalLight) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *DirectionalLight) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *DirectionalLight) GetShadowDepthRange() int64 {
	log.Println("Calling DirectionalLight.GetShadowDepthRange()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_shadow_depth_range", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *DirectionalLight) GetShadowMode() int64 {
	log.Println("Calling DirectionalLight.GetShadowMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_shadow_mode", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *DirectionalLight) IsBlendSplitsEnabled() bool {
	log.Println("Calling DirectionalLight.IsBlendSplitsEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_blend_splits_enabled", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *DirectionalLight) SetBlendSplits(enabled bool) {
	log.Println("Calling DirectionalLight.SetBlendSplits()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enabled)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_blend_splits", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *DirectionalLight) SetShadowDepthRange(mode int64) {
	log.Println("Calling DirectionalLight.SetShadowDepthRange()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_shadow_depth_range", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *DirectionalLight) SetShadowMode(mode int64) {
	log.Println("Calling DirectionalLight.SetShadowMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_shadow_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   DirectionalLightImplementer is an interface for DirectionalLight objects.
*/
type DirectionalLightImplementer interface {
	class.Class
}