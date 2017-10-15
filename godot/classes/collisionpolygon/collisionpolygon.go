//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package collisionpolygon

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

	"github.com/shadowapex/godot-go/godot/classes/spatial"
)

/*
   Allows editing a collision polygon's vertices on a selected plane. Can also set a depth perpendicular to that plane. This class is only available in the editor. It will not appear in the scene tree at runtime. Creates a [Shape] for gameplay. Properties modified during gameplay will have no effect.
*/
type CollisionPolygon struct {
	spatial.Spatial
}

func (o *CollisionPolygon) baseClass() string {
	return "CollisionPolygon"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *CollisionPolygon) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *CollisionPolygon) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *CollisionPolygon) GetDepth() float64 {
	log.Println("Calling CollisionPolygon.GetDepth()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_depth", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *CollisionPolygon) GetPolygon() *PoolVector2Array {
	log.Println("Calling CollisionPolygon.GetPolygon()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_polygon", goArguments, "*PoolVector2Array")

	returnValue := goRet.Interface().(*PoolVector2Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *CollisionPolygon) IsDisabled() bool {
	log.Println("Calling CollisionPolygon.IsDisabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_disabled", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *CollisionPolygon) SetDepth(depth float64) {
	log.Println("Calling CollisionPolygon.SetDepth()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(depth)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_depth", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *CollisionPolygon) SetDisabled(disabled bool) {
	log.Println("Calling CollisionPolygon.SetDisabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(disabled)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_disabled", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *CollisionPolygon) SetPolygon(polygon *PoolVector2Array) {
	log.Println("Calling CollisionPolygon.SetPolygon()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(polygon)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_polygon", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   CollisionPolygonImplementer is an interface for CollisionPolygon objects.
*/
type CollisionPolygonImplementer interface {
	class.Class
}
