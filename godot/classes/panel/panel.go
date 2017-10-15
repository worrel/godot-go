//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package panel

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

	"github.com/shadowapex/godot-go/godot/classes/control"
)

/*
   Panel is a [Control] that displays an opaque background. It's commonly used as a parent and container for other types of [Control] nodes.
*/
type Panel struct {
	control.Control
}

func (o *Panel) baseClass() string {
	return "Panel"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *Panel) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *Panel) getOwner() *C.godot_object {
	return o.owner
}

/*
   PanelImplementer is an interface for Panel objects.
*/
type PanelImplementer interface {
	class.Class
}