//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package inputeventjoypadbutton

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

	"github.com/shadowapex/godot-go/godot/classes/inputevent"
)

/*

 */
type InputEventJoypadButton struct {
	inputevent.InputEvent
}

func (o *InputEventJoypadButton) baseClass() string {
	return "InputEventJoypadButton"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *InputEventJoypadButton) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *InputEventJoypadButton) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *InputEventJoypadButton) GetButtonIndex() int64 {
	log.Println("Calling InputEventJoypadButton.GetButtonIndex()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_button_index", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *InputEventJoypadButton) GetPressure() float64 {
	log.Println("Calling InputEventJoypadButton.GetPressure()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_pressure", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *InputEventJoypadButton) SetButtonIndex(buttonIndex int64) {
	log.Println("Calling InputEventJoypadButton.SetButtonIndex()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(buttonIndex)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_button_index", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *InputEventJoypadButton) SetPressed(pressed bool) {
	log.Println("Calling InputEventJoypadButton.SetPressed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(pressed)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_pressed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *InputEventJoypadButton) SetPressure(pressure float64) {
	log.Println("Calling InputEventJoypadButton.SetPressure()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(pressure)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_pressure", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   InputEventJoypadButtonImplementer is an interface for InputEventJoypadButton objects.
*/
type InputEventJoypadButtonImplementer interface {
	class.Class
}
