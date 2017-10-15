//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package audiobuslayout

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

	"github.com/shadowapex/godot-go/godot/classes/resource"
)

/*
   Stores position, muting, solo, bypass, effects, effect position, volume, and the connections between busses. See [AudioServer] for usage.
*/
type AudioBusLayout struct {
	resource.Resource
}

func (o *AudioBusLayout) baseClass() string {
	return "AudioBusLayout"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *AudioBusLayout) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *AudioBusLayout) getOwner() *C.godot_object {
	return o.owner
}

/*
   AudioBusLayoutImplementer is an interface for AudioBusLayout objects.
*/
type AudioBusLayoutImplementer interface {
	class.Class
}
