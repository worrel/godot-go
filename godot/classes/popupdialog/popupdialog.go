//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package popupdialog

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

	"github.com/shadowapex/godot-go/godot/classes/popup"
)

/*
   PopupDialog is a base class for popup dialogs, along with [WindowDialog].
*/
type PopupDialog struct {
	popup.Popup
}

func (o *PopupDialog) baseClass() string {
	return "PopupDialog"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *PopupDialog) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *PopupDialog) getOwner() *C.godot_object {
	return o.owner
}

/*
   PopupDialogImplementer is an interface for PopupDialog objects.
*/
type PopupDialogImplementer interface {
	class.Class
}
