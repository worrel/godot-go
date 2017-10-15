//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package pinjoint

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

	"github.com/shadowapex/godot-go/godot/classes/joint"
)

/*
   Pin Joint for 3D Rigid Bodies. It pins 2 bodies (rigid or static) together.
*/
type PinJoint struct {
	joint.Joint
}

func (o *PinJoint) baseClass() string {
	return "PinJoint"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *PinJoint) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *PinJoint) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *PinJoint) GetParam(param int64) float64 {
	log.Println("Calling PinJoint.GetParam()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(param)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_param", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *PinJoint) SetParam(param int64, value float64) {
	log.Println("Calling PinJoint.SetParam()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(param)
	goArguments[1] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_param", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   PinJointImplementer is an interface for PinJoint objects.
*/
type PinJointImplementer interface {
	class.Class
}
