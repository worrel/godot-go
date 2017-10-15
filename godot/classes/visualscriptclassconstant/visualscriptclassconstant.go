//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package visualscriptclassconstant

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
   This node returns a constant from a given class, such as [@GlobalScope.TYPE_INT]. See the given class' documentation for available constants. [b]Input Ports:[/b] none [b]Output Ports:[/b] - Data (variant): [code]value[/code]
*/
type VisualScriptClassConstant struct {
	visualscriptnode.VisualScriptNode
}

func (o *VisualScriptClassConstant) baseClass() string {
	return "VisualScriptClassConstant"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VisualScriptClassConstant) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *VisualScriptClassConstant) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *VisualScriptClassConstant) GetBaseType() string {
	log.Println("Calling VisualScriptClassConstant.GetBaseType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_base_type", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VisualScriptClassConstant) GetClassConstant() string {
	log.Println("Calling VisualScriptClassConstant.GetClassConstant()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_class_constant", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VisualScriptClassConstant) SetBaseType(name string) {
	log.Println("Calling VisualScriptClassConstant.SetBaseType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_base_type", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *VisualScriptClassConstant) SetClassConstant(name string) {
	log.Println("Calling VisualScriptClassConstant.SetClassConstant()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_class_constant", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VisualScriptClassConstantImplementer is an interface for VisualScriptClassConstant objects.
*/
type VisualScriptClassConstantImplementer interface {
	class.Class
}
