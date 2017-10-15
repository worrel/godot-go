//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package undoredo

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

/*
   Helper to manage UndoRedo in the editor or custom tools. It works by storing calls to functions in both 'do' an 'undo' lists. Common behavior is to create an action, then add do/undo calls to functions or property changes, then committing the action.
*/
type UndoRedo struct {
	object.Object
}

func (o *UndoRedo) baseClass() string {
	return "UndoRedo"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *UndoRedo) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *UndoRedo) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *UndoRedo) AddDoMethod(object *Object, method string) *Variant {
	log.Println("Calling UndoRedo.AddDoMethod()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(object)
	goArguments[1] = reflect.ValueOf(method)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "add_do_method", goArguments, "*Variant")

	returnValue := goRet.Interface().(*Variant)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Set a property with a custom value.
*/
func (o *UndoRedo) AddDoProperty(object *Object, property string, value *Variant) {
	log.Println("Calling UndoRedo.AddDoProperty()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(object)
	goArguments[1] = reflect.ValueOf(property)
	goArguments[2] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_do_property", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Add a 'do' reference that will be erased if the 'do' history is lost. This is useful mostly for new nodes created for the 'do' call. Do not use for resources.
*/
func (o *UndoRedo) AddDoReference(object *Object) {
	log.Println("Calling UndoRedo.AddDoReference()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(object)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_do_reference", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *UndoRedo) AddUndoMethod(object *Object, method string) *Variant {
	log.Println("Calling UndoRedo.AddUndoMethod()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(object)
	goArguments[1] = reflect.ValueOf(method)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "add_undo_method", goArguments, "*Variant")

	returnValue := goRet.Interface().(*Variant)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undo setting of a property with a custom value.
*/
func (o *UndoRedo) AddUndoProperty(object *Object, property string, value *Variant) {
	log.Println("Calling UndoRedo.AddUndoProperty()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(object)
	goArguments[1] = reflect.ValueOf(property)
	goArguments[2] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_undo_property", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Add an 'undo' reference that will be erased if the 'undo' history is lost. This is useful mostly for nodes removed with the 'do' call (not the 'undo' call!).
*/
func (o *UndoRedo) AddUndoReference(object *Object) {
	log.Println("Calling UndoRedo.AddUndoReference()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(object)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_undo_reference", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Clear the undo/redo history and associated references.
*/
func (o *UndoRedo) ClearHistory() {
	log.Println("Calling UndoRedo.ClearHistory()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "clear_history", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Commit the action. All 'do' methods/properties are called/set when this function is called.
*/
func (o *UndoRedo) CommitAction() {
	log.Println("Calling UndoRedo.CommitAction()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "commit_action", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Create a new action. After this is called, do all your calls to [method add_do_method], [method add_undo_method], [method add_do_property] and [method add_undo_property].
*/
func (o *UndoRedo) CreateAction(name string, mergeMode int64) {
	log.Println("Calling UndoRedo.CreateAction()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(name)
	goArguments[1] = reflect.ValueOf(mergeMode)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "create_action", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Get the name of the current action.
*/
func (o *UndoRedo) GetCurrentActionName() string {
	log.Println("Calling UndoRedo.GetCurrentActionName()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_current_action_name", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *UndoRedo) GetMaxSteps() int64 {
	log.Println("Calling UndoRedo.GetMaxSteps()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_max_steps", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the version, each time a new action is committed, the version number of the UndoRedo is increased automatically. This is useful mostly to check if something changed from a saved version.
*/
func (o *UndoRedo) GetVersion() int64 {
	log.Println("Calling UndoRedo.GetVersion()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_version", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *UndoRedo) Redo() {
	log.Println("Calling UndoRedo.Redo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "redo", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *UndoRedo) SetMaxSteps(maxSteps int64) {
	log.Println("Calling UndoRedo.SetMaxSteps()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(maxSteps)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_max_steps", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *UndoRedo) Undo() {
	log.Println("Calling UndoRedo.Undo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "undo", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   UndoRedoImplementer is an interface for UndoRedo objects.
*/
type UndoRedoImplementer interface {
	class.Class
}
