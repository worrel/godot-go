//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package projectsettings

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

	"github.com/shadowapex/godot-go/godot/classes/object"
)

func newSingletonProjectSettings() *projectSettings {
	obj := &projectSettings{}
	ptr := C.godot_global_get_singleton(C.CString("ProjectSettings"))
	obj.owner = (*C.godot_object)(ptr)
	return obj
}

/*
   Contains global variables accessible from everywhere. Use the normal [Object] API, such as "ProjectSettings.get(variable)", "ProjectSettings.set(variable,value)" or "ProjectSettings.has(variable)" to access them. Variables stored in project.godot are also loaded into ProjectSettings, making this object very useful for reading custom game configuration options.
*/
var ProjectSettings = newSingletonProjectSettings()

/*
   Contains global variables accessible from everywhere. Use the normal [Object] API, such as "ProjectSettings.get(variable)", "ProjectSettings.set(variable,value)" or "ProjectSettings.has(variable)" to access them. Variables stored in project.godot are also loaded into ProjectSettings, making this object very useful for reading custom game configuration options.
*/
type projectSettings struct {
	object.Object
}

func (o *projectSettings) baseClass() string {
	return "ProjectSettings"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *projectSettings) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *projectSettings) getOwner() *C.godot_object {
	return o.owner
}

/*
   Add a custom property info to a property. The dictionary must contain: name:[String](the name of the property) and type:[int](see TYPE_* in [@Global Scope]), and optionally hint:[int](see PROPERTY_HINT_* in [@Global Scope]), hint_string:[String]. Example: [codeblock] ProjectSettings.set("category/property_name", 0) var property_info = { "name": "category/property_name", "type": TYPE_INT, "hint": PROPERTY_HINT_ENUM, "hint_string": "one,two,three" } ProjectSettings.add_property_info(property_info) [/codeblock]
*/
func (o *projectSettings) AddPropertyInfo(hint *Dictionary) {
	log.Println("Calling ProjectSettings.AddPropertyInfo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(hint)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_property_info", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Clear the whole configuration (not recommended, may break things).
*/
func (o *projectSettings) Clear(name string) {
	log.Println("Calling ProjectSettings.Clear()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "clear", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Return the order of a configuration value (influences when saved to the config file).
*/
func (o *projectSettings) GetOrder(name string) int64 {
	log.Println("Calling ProjectSettings.GetOrder()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_order", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *projectSettings) GetSingleton(name string) *Object {
	log.Println("Calling ProjectSettings.GetSingleton()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_singleton", goArguments, "*Object")

	returnValue := goRet.Interface().(*Object)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Convert a localized path (res://) to a full native OS path.
*/
func (o *projectSettings) GlobalizePath(path string) string {
	log.Println("Calling ProjectSettings.GlobalizePath()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(path)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "globalize_path", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *projectSettings) Has(name string) bool {
	log.Println("Calling ProjectSettings.Has()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "has", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *projectSettings) HasSingleton(name string) bool {
	log.Println("Calling ProjectSettings.HasSingleton()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "has_singleton", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *projectSettings) LoadResourcePack(pack string) bool {
	log.Println("Calling ProjectSettings.LoadResourcePack()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(pack)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "load_resource_pack", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Convert a path to a localized path (res:// path).
*/
func (o *projectSettings) LocalizePath(path string) string {
	log.Println("Calling ProjectSettings.LocalizePath()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(path)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "localize_path", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *projectSettings) PropertyCanRevert(name string) bool {
	log.Println("Calling ProjectSettings.PropertyCanRevert()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "property_can_revert", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *projectSettings) PropertyGetRevert(name string) *Variant {
	log.Println("Calling ProjectSettings.PropertyGetRevert()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "property_get_revert", goArguments, "*Variant")

	returnValue := goRet.Interface().(*Variant)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *projectSettings) Save() int64 {
	log.Println("Calling ProjectSettings.Save()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "save", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *projectSettings) SaveCustom(file string) int64 {
	log.Println("Calling ProjectSettings.SaveCustom()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(file)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "save_custom", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *projectSettings) SetInitialValue(name string, value *Variant) {
	log.Println("Calling ProjectSettings.SetInitialValue()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(name)
	goArguments[1] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_initial_value", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the order of a configuration value (influences when saved to the config file).
*/
func (o *projectSettings) SetOrder(name string, position int64) {
	log.Println("Calling ProjectSettings.SetOrder()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(name)
	goArguments[1] = reflect.ValueOf(position)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_order", goArguments, "")

	log.Println("  Function successfully completed.")

}