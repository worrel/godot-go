//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package worldenvironment

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

	"github.com/shadowapex/godot-go/godot/classes/node"
)

/*
   The [WorldEnvironment] node can be added to a scene in order to set default [Environment] variables for the scene. The [WorldEnvironment] can be overridden by an [Environment] node set on the current [Camera]. Additionally, only one [WorldEnvironment] may be instanced in a given scene at a time. The [WorldEnvironment] allows the user to specify default lighting parameters (e.g. ambient lighting), various post-processing effects (e.g. SSAO, DOF, Tonemapping), and how to draw the background (e.g. solid color, skybox).
*/
type WorldEnvironment struct {
	node.Node
}

func (o *WorldEnvironment) baseClass() string {
	return "WorldEnvironment"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *WorldEnvironment) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *WorldEnvironment) getOwner() *C.godot_object {
	return o.owner
}

/*
   Return the [Environment] currently bound.
*/
func (o *WorldEnvironment) GetEnvironment() *Environment {
	log.Println("Calling WorldEnvironment.GetEnvironment()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_environment", goArguments, "*Environment")

	returnValue := goRet.Interface().(*Environment)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Set the currently bound [Environment] to the one specified.
*/
func (o *WorldEnvironment) SetEnvironment(env *Environment) {
	log.Println("Calling WorldEnvironment.SetEnvironment()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(env)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_environment", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   WorldEnvironmentImplementer is an interface for WorldEnvironment objects.
*/
type WorldEnvironmentImplementer interface {
	class.Class
}