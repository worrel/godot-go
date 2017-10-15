//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package pckpacker

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

 */
type PCKPacker struct {
	reference.Reference
}

func (o *PCKPacker) baseClass() string {
	return "PCKPacker"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *PCKPacker) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *PCKPacker) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *PCKPacker) AddFile(pckPath string, sourcePath string) int64 {
	log.Println("Calling PCKPacker.AddFile()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(pckPath)
	goArguments[1] = reflect.ValueOf(sourcePath)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "add_file", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *PCKPacker) Flush(verbose bool) int64 {
	log.Println("Calling PCKPacker.Flush()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(verbose)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "flush", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *PCKPacker) PckStart(pckName string, alignment int64) int64 {
	log.Println("Calling PCKPacker.PckStart()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(pckName)
	goArguments[1] = reflect.ValueOf(alignment)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "pck_start", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   PCKPackerImplementer is an interface for PCKPacker objects.
*/
type PCKPackerImplementer interface {
	class.Class
}
