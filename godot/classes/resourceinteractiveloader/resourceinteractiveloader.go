//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package resourceinteractiveloader

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

	"github.com/shadowapex/godot-go/godot/classes/reference"
)

/*
   Interactive Resource Loader. This object is returned by ResourceLoader when performing an interactive load. It allows to load with high granularity, so this is mainly useful for displaying load bars/percentages.
*/
type ResourceInteractiveLoader struct {
	reference.Reference
}

func (o *ResourceInteractiveLoader) baseClass() string {
	return "ResourceInteractiveLoader"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *ResourceInteractiveLoader) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *ResourceInteractiveLoader) getOwner() *C.godot_object {
	return o.owner
}

/*
   Return the loaded resource (only if loaded). Otherwise, returns null.
*/
func (o *ResourceInteractiveLoader) GetResource() *Resource {
	log.Println("Calling ResourceInteractiveLoader.GetResource()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_resource", goArguments, "*Resource")

	returnValue := goRet.Interface().(*Resource)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the load stage. The total amount of stages can be queried with [method get_stage_count]
*/
func (o *ResourceInteractiveLoader) GetStage() int64 {
	log.Println("Calling ResourceInteractiveLoader.GetStage()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_stage", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the total amount of stages (calls to [method poll]) needed to completely load this resource.
*/
func (o *ResourceInteractiveLoader) GetStageCount() int64 {
	log.Println("Calling ResourceInteractiveLoader.GetStageCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_stage_count", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Poll the load. If OK is returned, this means poll will have to be called again. If ERR_FILE_EOF is returned, them the load has finished and the resource can be obtained by calling [method get_resource].
*/
func (o *ResourceInteractiveLoader) Poll() int64 {
	log.Println("Calling ResourceInteractiveLoader.Poll()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "poll", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ResourceInteractiveLoader) Wait() int64 {
	log.Println("Calling ResourceInteractiveLoader.Wait()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "wait", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   ResourceInteractiveLoaderImplementer is an interface for ResourceInteractiveLoader objects.
*/
type ResourceInteractiveLoaderImplementer interface {
	class.Class
}
