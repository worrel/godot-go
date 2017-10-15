//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package inputdefault

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

	"github.com/shadowapex/godot-go/godot/classes/input"
)

/*
   Default implementation of the [Input] class, used internally by the editor and games for default input management.
*/
type InputDefault struct {
	input.input
}

func (o *InputDefault) baseClass() string {
	return "InputDefault"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *InputDefault) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *InputDefault) getOwner() *C.godot_object {
	return o.owner
}

/*
   InputDefaultImplementer is an interface for InputDefault objects.
*/
type InputDefaultImplementer interface {
	class.Class
}
