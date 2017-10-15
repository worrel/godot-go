//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package vsplitcontainer

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

	"github.com/shadowapex/godot-go/godot/classes/splitcontainer"
)

/*
   Vertical split container. See [SplitContainer]. This goes from left to right.
*/
type VSplitContainer struct {
	splitcontainer.SplitContainer
}

func (o *VSplitContainer) baseClass() string {
	return "VSplitContainer"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VSplitContainer) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *VSplitContainer) getOwner() *C.godot_object {
	return o.owner
}

/*
   VSplitContainerImplementer is an interface for VSplitContainer objects.
*/
type VSplitContainerImplementer interface {
	class.Class
}