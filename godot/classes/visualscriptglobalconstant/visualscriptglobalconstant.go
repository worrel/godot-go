//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package visualscriptglobalconstant

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

 */
type VisualScriptGlobalConstant struct {
	visualscriptnode.VisualScriptNode
}

func (o *VisualScriptGlobalConstant) baseClass() string {
	return "VisualScriptGlobalConstant"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VisualScriptGlobalConstant) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *VisualScriptGlobalConstant) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *VisualScriptGlobalConstant) GetGlobalConstant() int64 {
	log.Println("Calling VisualScriptGlobalConstant.GetGlobalConstant()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_global_constant", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VisualScriptGlobalConstant) SetGlobalConstant(index int64) {
	log.Println("Calling VisualScriptGlobalConstant.SetGlobalConstant()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(index)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_global_constant", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VisualScriptGlobalConstantImplementer is an interface for VisualScriptGlobalConstant objects.
*/
type VisualScriptGlobalConstantImplementer interface {
	class.Class
}
