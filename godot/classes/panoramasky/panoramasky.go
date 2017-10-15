//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package panoramasky

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

	"github.com/shadowapex/godot-go/godot/classes/sky"
)

/*

 */
type PanoramaSky struct {
	sky.Sky
}

func (o *PanoramaSky) baseClass() string {
	return "PanoramaSky"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *PanoramaSky) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *PanoramaSky) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *PanoramaSky) GetPanorama() *Texture {
	log.Println("Calling PanoramaSky.GetPanorama()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_panorama", goArguments, "*Texture")

	returnValue := goRet.Interface().(*Texture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *PanoramaSky) SetPanorama(texture *Texture) {
	log.Println("Calling PanoramaSky.SetPanorama()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(texture)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_panorama", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   PanoramaSkyImplementer is an interface for PanoramaSky objects.
*/
type PanoramaSkyImplementer interface {
	class.Class
}