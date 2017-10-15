//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package hseparator

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

	"github.com/shadowapex/godot-go/godot/classes/separator"
)

/*
   Horizontal separator. See [Separator]. It is used to separate objects vertically, though (but it looks horizontal!).
*/
type HSeparator struct {
	separator.Separator
}

func (o *HSeparator) baseClass() string {
	return "HSeparator"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *HSeparator) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *HSeparator) getOwner() *C.godot_object {
	return o.owner
}

/*
   HSeparatorImplementer is an interface for HSeparator objects.
*/
type HSeparatorImplementer interface {
	class.Class
}
