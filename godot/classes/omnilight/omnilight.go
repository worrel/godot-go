//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package omnilight

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
   An OmniDirectional light is a type of [Light] node that emits lights in all directions. The light is attenuated through the distance and this attenuation can be configured by changing the energy, radius and attenuation parameters of [Light].
*/
type OmniLight struct {
	light.Light
}

func (o *OmniLight) baseClass() string {
	return "OmniLight"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *OmniLight) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *OmniLight) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *OmniLight) GetShadowDetail() int64 {
	log.Println("Calling OmniLight.GetShadowDetail()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_shadow_detail", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *OmniLight) GetShadowMode() int64 {
	log.Println("Calling OmniLight.GetShadowMode()")

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
func (o *OmniLight) SetShadowDetail(detail int64) {
	log.Println("Calling OmniLight.SetShadowDetail()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(detail)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_shadow_detail", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *OmniLight) SetShadowMode(mode int64) {
	log.Println("Calling OmniLight.SetShadowMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_shadow_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   OmniLightImplementer is an interface for OmniLight objects.
*/
type OmniLightImplementer interface {
	class.Class
}
