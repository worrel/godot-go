//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package visualscriptvariableset

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
   Changes a variable's value to the given input. [b]Input Ports:[/b] - Sequence - Data (variant): [code]set[/code] [b]Output Ports:[/b] - Sequence
*/
type VisualScriptVariableSet struct {
	visualscriptnode.VisualScriptNode
}

func (o *VisualScriptVariableSet) baseClass() string {
	return "VisualScriptVariableSet"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VisualScriptVariableSet) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *VisualScriptVariableSet) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *VisualScriptVariableSet) GetVariable() string {
	log.Println("Calling VisualScriptVariableSet.GetVariable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_variable", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VisualScriptVariableSet) SetVariable(name string) {
	log.Println("Calling VisualScriptVariableSet.SetVariable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_variable", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VisualScriptVariableSetImplementer is an interface for VisualScriptVariableSet objects.
*/
type VisualScriptVariableSetImplementer interface {
	class.Class
}