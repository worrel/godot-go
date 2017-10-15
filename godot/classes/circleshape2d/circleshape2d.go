//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package circleshape2d

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

	"github.com/shadowapex/godot-go/godot/classes/shape2d"
)

/*
   Circular shape for 2D collisions. This shape is useful for modeling balls or small characters and its collision detection with everything else is very fast.
*/
type CircleShape2D struct {
	shape2d.Shape2D
}

func (o *CircleShape2D) baseClass() string {
	return "CircleShape2D"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *CircleShape2D) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *CircleShape2D) getOwner() *C.godot_object {
	return o.owner
}

/*
   Return the radius of the circle shape.
*/
func (o *CircleShape2D) GetRadius() float64 {
	log.Println("Calling CircleShape2D.GetRadius()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_radius", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Set the radius of the circle shape.
*/
func (o *CircleShape2D) SetRadius(radius float64) {
	log.Println("Calling CircleShape2D.SetRadius()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(radius)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_radius", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   CircleShape2DImplementer is an interface for CircleShape2D objects.
*/
type CircleShape2DImplementer interface {
	class.Class
}
