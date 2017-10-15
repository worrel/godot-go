//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package groovejoint2d

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

	"github.com/shadowapex/godot-go/godot/classes/joint2d"
)

/*
   Groove constraint for 2D physics. This is useful for making a body "slide" through a segment placed in another.
*/
type GrooveJoint2D struct {
	joint2d.Joint2D
}

func (o *GrooveJoint2D) baseClass() string {
	return "GrooveJoint2D"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *GrooveJoint2D) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *GrooveJoint2D) getOwner() *C.godot_object {
	return o.owner
}

/*
   Set the final offset of the groove on body A.
*/
func (o *GrooveJoint2D) GetInitialOffset() float64 {
	log.Println("Calling GrooveJoint2D.GetInitialOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_initial_offset", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the length of the groove.
*/
func (o *GrooveJoint2D) GetLength() float64 {
	log.Println("Calling GrooveJoint2D.GetLength()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_length", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Set the initial offset of the groove on body A.
*/
func (o *GrooveJoint2D) SetInitialOffset(offset float64) {
	log.Println("Calling GrooveJoint2D.SetInitialOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(offset)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_initial_offset", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the length of the groove.
*/
func (o *GrooveJoint2D) SetLength(length float64) {
	log.Println("Calling GrooveJoint2D.SetLength()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(length)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_length", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   GrooveJoint2DImplementer is an interface for GrooveJoint2D objects.
*/
type GrooveJoint2DImplementer interface {
	class.Class
}
