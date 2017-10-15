//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package instanceplaceholder

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
   Turning on the option [b]Load As Placeholder[/b] for an instanced scene in the editor causes it to be replaced by an InstacePlaceholder when running the game. This makes it possible to delay actually loading the scene until calling [method replace_by_instance]. This is useful to avoid loading large scenes all at once by loading parts of it selectively. The InstancePlaceholder does not have a transform. This causes any child nodes to be positioned relatively to the Viewport from point (0,0), rather than their parent as displayed in the editor. Replacing the placeholder with a scene with a transform will transform children relatively to their parent again.
*/
type InstancePlaceholder struct {
	node.Node
}

func (o *InstancePlaceholder) baseClass() string {
	return "InstancePlaceholder"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *InstancePlaceholder) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *InstancePlaceholder) getOwner() *C.godot_object {
	return o.owner
}

/*
   Retrieve the path to the [PackedScene] resource file that is loaded by default when calling [method replace_by_instance].
*/
func (o *InstancePlaceholder) GetInstancePath() string {
	log.Println("Calling InstancePlaceholder.GetInstancePath()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_instance_path", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *InstancePlaceholder) GetStoredValues(withOrder bool) *Dictionary {
	log.Println("Calling InstancePlaceholder.GetStoredValues()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(withOrder)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_stored_values", goArguments, "*Dictionary")

	returnValue := goRet.Interface().(*Dictionary)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Replace this placeholder by the scene handed as an argument, or the original scene if no argument is given. As for all resources, the scene is loaded only if it's not loaded already. By manually loading the scene beforehand, delays caused by this function can be avoided.
*/
func (o *InstancePlaceholder) ReplaceByInstance(customScene *PackedScene) {
	log.Println("Calling InstancePlaceholder.ReplaceByInstance()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(customScene)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "replace_by_instance", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   InstancePlaceholderImplementer is an interface for InstancePlaceholder objects.
*/
type InstancePlaceholderImplementer interface {
	class.Class
}
