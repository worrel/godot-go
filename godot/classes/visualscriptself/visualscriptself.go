//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package visualscriptself

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

	"github.com/shadowapex/godot-go/godot/classes/visualscriptnode"
)

/*
   Provides a reference to the node running the visual script. [b]Input Ports:[/b] none [b]Output Ports:[/b] - Data (object): [code]instance[/code]
*/
type VisualScriptSelf struct {
	visualscriptnode.VisualScriptNode
}

func (o *VisualScriptSelf) baseClass() string {
	return "VisualScriptSelf"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VisualScriptSelf) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *VisualScriptSelf) getOwner() *C.godot_object {
	return o.owner
}

/*
   VisualScriptSelfImplementer is an interface for VisualScriptSelf objects.
*/
type VisualScriptSelfImplementer interface {
	class.Class
}
