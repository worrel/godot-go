//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package visualscriptpropertyset

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
type VisualScriptPropertySet struct {
	visualscriptnode.VisualScriptNode
}

func (o *VisualScriptPropertySet) baseClass() string {
	return "VisualScriptPropertySet"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VisualScriptPropertySet) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *VisualScriptPropertySet) getOwner() *C.godot_object {
	return o.owner
}

/*
   Undocumented
*/
func (o *VisualScriptPropertySet) X_GetTypeCache() *Dictionary {
	log.Println("Calling VisualScriptPropertySet.X_GetTypeCache()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_type_cache", goArguments, "*Dictionary")

	returnValue := goRet.Interface().(*Dictionary)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VisualScriptPropertySet) X_SetTypeCache(typeCache *Dictionary) {
	log.Println("Calling VisualScriptPropertySet.X_SetTypeCache()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(typeCache)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_set_type_cache", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *VisualScriptPropertySet) GetAssignOp() int64 {
	log.Println("Calling VisualScriptPropertySet.GetAssignOp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_assign_op", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VisualScriptPropertySet) GetBasePath() *NodePath {
	log.Println("Calling VisualScriptPropertySet.GetBasePath()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_base_path", goArguments, "*NodePath")

	returnValue := goRet.Interface().(*NodePath)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VisualScriptPropertySet) GetBaseScript() string {
	log.Println("Calling VisualScriptPropertySet.GetBaseScript()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_base_script", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VisualScriptPropertySet) GetBaseType() string {
	log.Println("Calling VisualScriptPropertySet.GetBaseType()")

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
func (o *VisualScriptPropertySet) GetBasicType() int64 {
	log.Println("Calling VisualScriptPropertySet.GetBasicType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_basic_type", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VisualScriptPropertySet) GetCallMode() int64 {
	log.Println("Calling VisualScriptPropertySet.GetCallMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_call_mode", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VisualScriptPropertySet) GetIndex() string {
	log.Println("Calling VisualScriptPropertySet.GetIndex()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_index", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VisualScriptPropertySet) GetProperty() string {
	log.Println("Calling VisualScriptPropertySet.GetProperty()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_property", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VisualScriptPropertySet) SetAssignOp(assignOp int64) {
	log.Println("Calling VisualScriptPropertySet.SetAssignOp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(assignOp)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_assign_op", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *VisualScriptPropertySet) SetBasePath(basePath *NodePath) {
	log.Println("Calling VisualScriptPropertySet.SetBasePath()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(basePath)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_base_path", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *VisualScriptPropertySet) SetBaseScript(baseScript string) {
	log.Println("Calling VisualScriptPropertySet.SetBaseScript()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(baseScript)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_base_script", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *VisualScriptPropertySet) SetBaseType(baseType string) {
	log.Println("Calling VisualScriptPropertySet.SetBaseType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(baseType)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_base_type", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *VisualScriptPropertySet) SetBasicType(basicType int64) {
	log.Println("Calling VisualScriptPropertySet.SetBasicType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(basicType)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_basic_type", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *VisualScriptPropertySet) SetCallMode(mode int64) {
	log.Println("Calling VisualScriptPropertySet.SetCallMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_call_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *VisualScriptPropertySet) SetIndex(index string) {
	log.Println("Calling VisualScriptPropertySet.SetIndex()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(index)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_index", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *VisualScriptPropertySet) SetProperty(property string) {
	log.Println("Calling VisualScriptPropertySet.SetProperty()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(property)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_property", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VisualScriptPropertySetImplementer is an interface for VisualScriptPropertySet objects.
*/
type VisualScriptPropertySetImplementer interface {
	class.Class
}
