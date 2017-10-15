//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package lightoccluder2d

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

	"github.com/shadowapex/godot-go/godot/classes/node2d"
)

/*
   Occludes light cast by a Light2D, casting shadows. The LightOccluder2D must be provided with an [OccluderPolygon2D] in order for the shadow to be computed.
*/
type LightOccluder2D struct {
	node2d.Node2D
}

func (o *LightOccluder2D) baseClass() string {
	return "LightOccluder2D"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *LightOccluder2D) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *LightOccluder2D) getOwner() *C.godot_object {
	return o.owner
}

/*
   Undocumented
*/
func (o *LightOccluder2D) X_PolyChanged() {
	log.Println("Calling LightOccluder2D.X_PolyChanged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_poly_changed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Return the light mask of the LightOccluder2D.
*/
func (o *LightOccluder2D) GetOccluderLightMask() int64 {
	log.Println("Calling LightOccluder2D.GetOccluderLightMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_occluder_light_mask", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the OccluderPolygon2D that defines the LightOccluder2D.
*/
func (o *LightOccluder2D) GetOccluderPolygon() *OccluderPolygon2D {
	log.Println("Calling LightOccluder2D.GetOccluderPolygon()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_occluder_polygon", goArguments, "*OccluderPolygon2D")

	returnValue := goRet.Interface().(*OccluderPolygon2D)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Set the LightOccluder2D light mask. The LightOccluder2D will cast shadows only from Light2Ds that belong to the same light mask(s).
*/
func (o *LightOccluder2D) SetOccluderLightMask(mask int64) {
	log.Println("Calling LightOccluder2D.SetOccluderLightMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mask)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_occluder_light_mask", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the OccluderPolygon2D that defines the LightOccluder2D.
*/
func (o *LightOccluder2D) SetOccluderPolygon(polygon *OccluderPolygon2D) {
	log.Println("Calling LightOccluder2D.SetOccluderPolygon()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(polygon)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_occluder_polygon", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   LightOccluder2DImplementer is an interface for LightOccluder2D objects.
*/
type LightOccluder2DImplementer interface {
	class.Class
}
