//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package visibilityenabler

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

	"github.com/shadowapex/godot-go/godot/classes/visibilitynotifier"
)

/*
   The VisibilityEnabler will disable [RigidBody] and [AnimationPlayer] nodes when they are not visible. It will only affect other nodes within the same scene as the VisibilityEnabler itself.
*/
type VisibilityEnabler struct {
	visibilitynotifier.VisibilityNotifier
}

func (o *VisibilityEnabler) baseClass() string {
	return "VisibilityEnabler"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VisibilityEnabler) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *VisibilityEnabler) getOwner() *C.godot_object {
	return o.owner
}

/*
   Undocumented
*/
func (o *VisibilityEnabler) X_NodeRemoved(arg0 *Object) {
	log.Println("Calling VisibilityEnabler.X_NodeRemoved()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_node_removed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Returns whether the specified enabler was set to true or not.
*/
func (o *VisibilityEnabler) IsEnablerEnabled(enabler int64) bool {
	log.Println("Calling VisibilityEnabler.IsEnablerEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enabler)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_enabler_enabled", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Set an enabler to true for all nodes of its type to be disabled when the VisibilityEnabler is not in view. See the constants for enablers and what they affect.
*/
func (o *VisibilityEnabler) SetEnabler(enabler int64, enabled bool) {
	log.Println("Calling VisibilityEnabler.SetEnabler()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(enabler)
	goArguments[1] = reflect.ValueOf(enabled)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_enabler", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VisibilityEnablerImplementer is an interface for VisibilityEnabler objects.
*/
type VisibilityEnablerImplementer interface {
	class.Class
}