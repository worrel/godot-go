//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package visualscriptwhile

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
   Loops while a condition is [code]true[/code]. Execution continues out the [code]exit[/code] Sequence port when the loop terminates. [b]Input Ports:[/b] - Sequence: [code]while(cond)[/code] - Data (bool): [code]cond[/code] [b]Output Ports:[/b] - Sequence: [code]repeat[/code] - Sequence: [code]exit[/code]
*/
type VisualScriptWhile struct {
	visualscriptnode.VisualScriptNode
}

func (o *VisualScriptWhile) baseClass() string {
	return "VisualScriptWhile"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VisualScriptWhile) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *VisualScriptWhile) getOwner() *C.godot_object {
	return o.owner
}

/*
   VisualScriptWhileImplementer is an interface for VisualScriptWhile objects.
*/
type VisualScriptWhileImplementer interface {
	class.Class
}
