//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package visualscriptemitsignal

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
   Emits a specified signal when it is executed. [b]Input Ports:[/b] - Sequence: [code]emit[/code] [b]Output Ports:[/b] - Sequence
*/
type VisualScriptEmitSignal struct {
	visualscriptnode.VisualScriptNode
}

func (o *VisualScriptEmitSignal) baseClass() string {
	return "VisualScriptEmitSignal"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VisualScriptEmitSignal) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *VisualScriptEmitSignal) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *VisualScriptEmitSignal) GetSignal() string {
	log.Println("Calling VisualScriptEmitSignal.GetSignal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_signal", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VisualScriptEmitSignal) SetSignal(name string) {
	log.Println("Calling VisualScriptEmitSignal.SetSignal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_signal", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VisualScriptEmitSignalImplementer is an interface for VisualScriptEmitSignal objects.
*/
type VisualScriptEmitSignalImplementer interface {
	class.Class
}
