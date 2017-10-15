//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package acceptdialog

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

	"github.com/shadowapex/godot-go/godot/classes/windowdialog"
)

/*
   This dialog is useful for small notifications to the user about an event. It can only be accepted or closed, with the same result.
*/
type AcceptDialog struct {
	windowdialog.WindowDialog
}

func (o *AcceptDialog) baseClass() string {
	return "AcceptDialog"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *AcceptDialog) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *AcceptDialog) getOwner() *C.godot_object {
	return o.owner
}

/*
   Undocumented
*/
func (o *AcceptDialog) X_BuiltinTextEntered(arg0 string) {
	log.Println("Calling AcceptDialog.X_BuiltinTextEntered()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_builtin_text_entered", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AcceptDialog) X_CustomAction(arg0 string) {
	log.Println("Calling AcceptDialog.X_CustomAction()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_custom_action", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *AcceptDialog) X_Ok() {
	log.Println("Calling AcceptDialog.X_Ok()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_ok", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Adds a button with label [i]text[/i] and a custom [i]action[/i] to the dialog and returns the created button. [i]action[/i] will be passed to the [custom_action] signal when pressed. If [code]true[/code], [i]right[/i] will place the button to the right of any sibling buttons. Default value: [code]false[/code].
*/
func (o *AcceptDialog) AddButton(text string, right bool, action string) *Button {
	log.Println("Calling AcceptDialog.AddButton()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(text)
	goArguments[1] = reflect.ValueOf(right)
	goArguments[2] = reflect.ValueOf(action)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "add_button", goArguments, "*Button")

	returnValue := goRet.Interface().(*Button)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Adds a button with label [i]name[/i] and a cancel action to the dialog and returns the created button.
*/
func (o *AcceptDialog) AddCancel(name string) *Button {
	log.Println("Calling AcceptDialog.AddCancel()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "add_cancel", goArguments, "*Button")

	returnValue := goRet.Interface().(*Button)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return true if the dialog will be hidden when accepted (default true).
*/
func (o *AcceptDialog) GetHideOnOk() bool {
	log.Println("Calling AcceptDialog.GetHideOnOk()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_hide_on_ok", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the label used for built-in text.
*/
func (o *AcceptDialog) GetLabel() *Label {
	log.Println("Calling AcceptDialog.GetLabel()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_label", goArguments, "*Label")

	returnValue := goRet.Interface().(*Label)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the OK Button.
*/
func (o *AcceptDialog) GetOk() *Button {
	log.Println("Calling AcceptDialog.GetOk()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_ok", goArguments, "*Button")

	returnValue := goRet.Interface().(*Button)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the built-in label text.
*/
func (o *AcceptDialog) GetText() string {
	log.Println("Calling AcceptDialog.GetText()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_text", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Registers a [LineEdit] in the dialog. When the enter key is pressed, the dialog will be accepted.
*/
func (o *AcceptDialog) RegisterTextEnter(lineEdit *Object) {
	log.Println("Calling AcceptDialog.RegisterTextEnter()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(lineEdit)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "register_text_enter", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set whether the dialog is hidden when accepted (default true).
*/
func (o *AcceptDialog) SetHideOnOk(enabled bool) {
	log.Println("Calling AcceptDialog.SetHideOnOk()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enabled)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_hide_on_ok", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the built-in label text.
*/
func (o *AcceptDialog) SetText(text string) {
	log.Println("Calling AcceptDialog.SetText()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(text)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_text", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   AcceptDialogImplementer is an interface for AcceptDialog objects.
*/
type AcceptDialogImplementer interface {
	class.Class
}