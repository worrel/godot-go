//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package audioeffecthighshelffilter

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

	"github.com/shadowapex/godot-go/godot/classes/audioeffectfilter"
)

/*

 */
type AudioEffectHighShelfFilter struct {
	audioeffectfilter.AudioEffectFilter
}

func (o *AudioEffectHighShelfFilter) baseClass() string {
	return "AudioEffectHighShelfFilter"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *AudioEffectHighShelfFilter) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *AudioEffectHighShelfFilter) getOwner() *C.godot_object {
	return o.owner
}

/*
   AudioEffectHighShelfFilterImplementer is an interface for AudioEffectHighShelfFilter objects.
*/
type AudioEffectHighShelfFilterImplementer interface {
	class.Class
}
