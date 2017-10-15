//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package visualscriptlocalvarset

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
   Changes a local variable's value to the given input. The new value is also provided on an output Data port. [b]Input Ports:[/b] - Sequence - Data (variant): [code]set[/code] [b]Output Ports:[/b] - Sequence - Data (variant): [code]get[/code]
*/
type VisualScriptLocalVarSet struct {
	visualscriptnode.VisualScriptNode
}

func (o *VisualScriptLocalVarSet) baseClass() string {
	return "VisualScriptLocalVarSet"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VisualScriptLocalVarSet) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *VisualScriptLocalVarSet) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *VisualScriptLocalVarSet) GetVarName() string {
	log.Println("Calling VisualScriptLocalVarSet.GetVarName()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_var_name", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VisualScriptLocalVarSet) GetVarType() int64 {
	log.Println("Calling VisualScriptLocalVarSet.GetVarType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_var_type", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VisualScriptLocalVarSet) SetVarName(name string) {
	log.Println("Calling VisualScriptLocalVarSet.SetVarName()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_var_name", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *VisualScriptLocalVarSet) SetVarType(aType int64) {
	log.Println("Calling VisualScriptLocalVarSet.SetVarType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(aType)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_var_type", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VisualScriptLocalVarSetImplementer is an interface for VisualScriptLocalVarSet objects.
*/
type VisualScriptLocalVarSetImplementer interface {
	class.Class
}
