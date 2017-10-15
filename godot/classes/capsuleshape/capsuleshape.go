//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package capsuleshape

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

	"github.com/shadowapex/godot-go/godot/classes/shape"
)

/*
   Capsule shape for collisions.
*/
type CapsuleShape struct {
	shape.Shape
}

func (o *CapsuleShape) baseClass() string {
	return "CapsuleShape"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *CapsuleShape) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *CapsuleShape) getOwner() *C.godot_object {
	return o.owner
}

/*
   Return the capsule height.
*/
func (o *CapsuleShape) GetHeight() float64 {
	log.Println("Calling CapsuleShape.GetHeight()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_height", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the capsule radius.
*/
func (o *CapsuleShape) GetRadius() float64 {
	log.Println("Calling CapsuleShape.GetRadius()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_radius", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Set the capsule height.
*/
func (o *CapsuleShape) SetHeight(height float64) {
	log.Println("Calling CapsuleShape.SetHeight()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(height)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_height", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the capsule radius.
*/
func (o *CapsuleShape) SetRadius(radius float64) {
	log.Println("Calling CapsuleShape.SetRadius()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(radius)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_radius", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   CapsuleShapeImplementer is an interface for CapsuleShape objects.
*/
type CapsuleShapeImplementer interface {
	class.Class
}
