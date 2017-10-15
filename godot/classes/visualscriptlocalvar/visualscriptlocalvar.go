//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package visualscriptlocalvar

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
   Returns a local variable's value. "Var Name" must be supplied, with an optional type. [b]Input Ports:[/b] none [b]Output Ports:[/b] - Data (variant): [code]get[/code]
*/
type VisualScriptLocalVar struct {
	visualscriptnode.VisualScriptNode
}

func (o *VisualScriptLocalVar) baseClass() string {
	return "VisualScriptLocalVar"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VisualScriptLocalVar) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *VisualScriptLocalVar) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *VisualScriptLocalVar) GetVarName() string {
	log.Println("Calling VisualScriptLocalVar.GetVarName()")

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
func (o *VisualScriptLocalVar) GetVarType() int64 {
	log.Println("Calling VisualScriptLocalVar.GetVarType()")

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
func (o *VisualScriptLocalVar) SetVarName(name string) {
	log.Println("Calling VisualScriptLocalVar.SetVarName()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_var_name", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *VisualScriptLocalVar) SetVarType(aType int64) {
	log.Println("Calling VisualScriptLocalVar.SetVarType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(aType)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_var_type", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VisualScriptLocalVarImplementer is an interface for VisualScriptLocalVar objects.
*/
type VisualScriptLocalVarImplementer interface {
	class.Class
}
