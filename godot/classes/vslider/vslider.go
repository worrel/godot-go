//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package vslider

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

	"github.com/shadowapex/godot-go/godot/classes/slider"
)

/*
   Vertical slider. See [Slider]. This one goes from left (min) to right (max).
*/
type VSlider struct {
	slider.Slider
}

func (o *VSlider) baseClass() string {
	return "VSlider"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VSlider) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *VSlider) getOwner() *C.godot_object {
	return o.owner
}

/*
   VSliderImplementer is an interface for VSlider objects.
*/
type VSliderImplementer interface {
	class.Class
}