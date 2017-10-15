//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package hboxcontainer

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

	"github.com/shadowapex/godot-go/godot/classes/boxcontainer"
)

/*
   Horizontal box container. See [BoxContainer].
*/
type HBoxContainer struct {
	boxcontainer.BoxContainer
}

func (o *HBoxContainer) baseClass() string {
	return "HBoxContainer"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *HBoxContainer) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *HBoxContainer) getOwner() *C.godot_object {
	return o.owner
}

/*
   HBoxContainerImplementer is an interface for HBoxContainer objects.
*/
type HBoxContainerImplementer interface {
	class.Class
}
