//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package visualscriptcustomnode

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

	"github.com/shadowapex/godot-go/godot/classes/visualscriptnode"
)

/*
   A custom Visual Script node which can be scripted in powerful ways.
*/
type VisualScriptCustomNode struct {
	visualscriptnode.VisualScriptNode
}

func (o *VisualScriptCustomNode) baseClass() string {
	return "VisualScriptCustomNode"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VisualScriptCustomNode) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *VisualScriptCustomNode) getOwner() *C.godot_object {
	return o.owner
}

/*
   Return the node's title.
*/
func (o *VisualScriptCustomNode) X_GetCaption() string {
	log.Println("Calling VisualScriptCustomNode.X_GetCaption()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_caption", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the node's category.
*/
func (o *VisualScriptCustomNode) X_GetCategory() string {
	log.Println("Calling VisualScriptCustomNode.X_GetCategory()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_category", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the count of input value ports.
*/
func (o *VisualScriptCustomNode) X_GetInputValuePortCount() int64 {
	log.Println("Calling VisualScriptCustomNode.X_GetInputValuePortCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_input_value_port_count", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the specified input port's name.
*/
func (o *VisualScriptCustomNode) X_GetInputValuePortName(idx int64) string {
	log.Println("Calling VisualScriptCustomNode.X_GetInputValuePortName()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_input_value_port_name", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the specified input port's type. See the TYPE_* enum in [@GlobalScope].
*/
func (o *VisualScriptCustomNode) X_GetInputValuePortType(idx int64) int64 {
	log.Println("Calling VisualScriptCustomNode.X_GetInputValuePortType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_input_value_port_type", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the amount of output [b]sequence[/b] ports.
*/
func (o *VisualScriptCustomNode) X_GetOutputSequencePortCount() int64 {
	log.Println("Calling VisualScriptCustomNode.X_GetOutputSequencePortCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_output_sequence_port_count", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the specified [b]sequence[/b] output's name.
*/
func (o *VisualScriptCustomNode) X_GetOutputSequencePortText(idx int64) string {
	log.Println("Calling VisualScriptCustomNode.X_GetOutputSequencePortText()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_output_sequence_port_text", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the amount of output value ports.
*/
func (o *VisualScriptCustomNode) X_GetOutputValuePortCount() int64 {
	log.Println("Calling VisualScriptCustomNode.X_GetOutputValuePortCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_output_value_port_count", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the specified output's name.
*/
func (o *VisualScriptCustomNode) X_GetOutputValuePortName(idx int64) string {
	log.Println("Calling VisualScriptCustomNode.X_GetOutputValuePortName()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_output_value_port_name", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the specified output's type. See the TYPE_* enum in [@GlobalScope].
*/
func (o *VisualScriptCustomNode) X_GetOutputValuePortType(idx int64) int64 {
	log.Println("Calling VisualScriptCustomNode.X_GetOutputValuePortType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_output_value_port_type", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the custom node's text, which is shown right next to the input [b]sequence[/b] port (if there is none, on the place that is usually taken by it).
*/
func (o *VisualScriptCustomNode) X_GetText() string {
	log.Println("Calling VisualScriptCustomNode.X_GetText()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_text", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the size of the custom node's working memory. See [method _step] for more details.
*/
func (o *VisualScriptCustomNode) X_GetWorkingMemorySize() int64 {
	log.Println("Calling VisualScriptCustomNode.X_GetWorkingMemorySize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_working_memory_size", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return whether the custom node has an input [b]sequence[/b] port.
*/
func (o *VisualScriptCustomNode) X_HasInputSequencePort() bool {
	log.Println("Calling VisualScriptCustomNode.X_HasInputSequencePort()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_has_input_sequence_port", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VisualScriptCustomNode) X_ScriptChanged() {
	log.Println("Calling VisualScriptCustomNode.X_ScriptChanged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_script_changed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Execute the custom node's logic, returning the index of the output sequence port to use or a [String] when there is an error. The [code]inputs[/code] array contains the values of the input ports. [code]outputs[/code] is an array whose indices should be set to the respective outputs. The [code]start_mode[/code] is usually [code]START_MODE_BEGIN_SEQUENCE[/code], unless you have used the STEP_* constants. [code]working_mem[/code] is an array which can be used to persist information between runs of the custom node. When returning, you can mask the returned value with one of the STEP_* constants.
*/
func (o *VisualScriptCustomNode) X_Step(inputs *Array, outputs *Array, startMode int64, workingMem *Array) *Variant {
	log.Println("Calling VisualScriptCustomNode.X_Step()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(inputs)
	goArguments[1] = reflect.ValueOf(outputs)
	goArguments[2] = reflect.ValueOf(startMode)
	goArguments[3] = reflect.ValueOf(workingMem)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_step", goArguments, "*Variant")

	returnValue := goRet.Interface().(*Variant)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   VisualScriptCustomNodeImplementer is an interface for VisualScriptCustomNode objects.
*/
type VisualScriptCustomNodeImplementer interface {
	class.Class
}
