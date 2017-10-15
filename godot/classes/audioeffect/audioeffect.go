//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package audioeffect

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
   Base resource for audio bus. Applies an audio effect on the bus that the resource is applied on.
*/
type AudioEffect struct {
	resource.Resource
}

func (o *AudioEffect) baseClass() string {
	return "AudioEffect"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *AudioEffect) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *AudioEffect) getOwner() *C.godot_object {
	return o.owner
}

/*
   AudioEffectImplementer is an interface for AudioEffect objects.
*/
type AudioEffectImplementer interface {
	class.Class
}
