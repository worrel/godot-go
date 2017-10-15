//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package ip_unix

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

	"github.com/shadowapex/godot-go/godot/classes/ip"
)

/*

 */
type IP_Unix struct {
	ip.ip
}

func (o *IP_Unix) baseClass() string {
	return "IP_Unix"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *IP_Unix) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *IP_Unix) getOwner() *C.godot_object {
	return o.owner
}

/*
   IP_UnixImplementer is an interface for IP_Unix objects.
*/
type IP_UnixImplementer interface {
	class.Class
}
