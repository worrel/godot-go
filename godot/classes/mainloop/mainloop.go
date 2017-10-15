//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package mainloop

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
   Main loop is the abstract main loop base class. All other main loop classes are derived from it. Upon application start, a [MainLoop] has to be provided to OS, else the application will exit. This happens automatically (and a [SceneTree] is created), unless a main [Script] is supplied, which may or not create and return a [MainLoop].
*/
type MainLoop struct {
	object.Object
}

func (o *MainLoop) baseClass() string {
	return "MainLoop"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *MainLoop) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *MainLoop) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *MainLoop) X_DropFiles(files *PoolStringArray, screen int64) {
	log.Println("Calling MainLoop.X_DropFiles()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(files)
	goArguments[1] = reflect.ValueOf(screen)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_drop_files", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Called before the program exits.
*/
func (o *MainLoop) X_Finalize() {
	log.Println("Calling MainLoop.X_Finalize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_finalize", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Called each idle frame with time since last call as an only argument.
*/
func (o *MainLoop) X_Idle(delta float64) {
	log.Println("Calling MainLoop.X_Idle()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(delta)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_idle", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Called once during initialization.
*/
func (o *MainLoop) X_Initialize() {
	log.Println("Calling MainLoop.X_Initialize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_initialize", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *MainLoop) X_InputEvent(ev *InputEvent) {
	log.Println("Calling MainLoop.X_InputEvent()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(ev)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_input_event", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *MainLoop) X_InputText(text string) {
	log.Println("Calling MainLoop.X_InputText()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(text)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_input_text", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *MainLoop) X_Iteration(delta float64) {
	log.Println("Calling MainLoop.X_Iteration()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(delta)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_iteration", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *MainLoop) Finish() {
	log.Println("Calling MainLoop.Finish()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "finish", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *MainLoop) Idle(delta float64) bool {
	log.Println("Calling MainLoop.Idle()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(delta)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "idle", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *MainLoop) Init() {
	log.Println("Calling MainLoop.Init()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "init", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *MainLoop) InputEvent(ev *InputEvent) {
	log.Println("Calling MainLoop.InputEvent()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(ev)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "input_event", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *MainLoop) InputText(text string) {
	log.Println("Calling MainLoop.InputText()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(text)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "input_text", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *MainLoop) Iteration(delta float64) bool {
	log.Println("Calling MainLoop.Iteration()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(delta)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "iteration", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   MainLoopImplementer is an interface for MainLoop objects.
*/
type MainLoopImplementer interface {
	class.Class
}