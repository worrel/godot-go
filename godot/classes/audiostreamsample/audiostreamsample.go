//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package audiostreamsample

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

	"github.com/shadowapex/godot-go/godot/classes/audiostream"
)

/*
   Plays audio, can loop.
*/
type AudioStreamSample struct {
	audiostream.AudioStream
}

func (o *AudioStreamSample) baseClass() string {
	return "AudioStreamSample"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *AudioStreamSample) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *AudioStreamSample) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *AudioStreamSample) GetData() *PoolByteArray {
	log.Println("Calling AudioStreamSample.GetData()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_data", goArguments, "*PoolByteArray")

	returnValue := goRet.Interface().(*PoolByteArray)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *AudioStreamSample) GetFormat() int64 {
	log.Println("Calling AudioStreamSample.GetFormat()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_format", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *AudioStreamSample) GetLoopBegin() int64 {
	log.Println("Calling AudioStreamSample.GetLoopBegin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_loop_begin", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *AudioStreamSample) GetLoopEnd() int64 {
	log.Println("Calling AudioStreamSample.GetLoopEnd()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_loop_end", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *AudioStreamSample) GetLoopMode() int64 {
	log.Println("Calling AudioStreamSample.GetLoopMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_loop_mode", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *AudioStreamSample) GetMixRate() int64 {
	log.Println("Calling AudioStreamSample.GetMixRate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_mix_rate", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *AudioStreamSample) IsStereo() bool {
	log.Println("Calling AudioStreamSample.IsStereo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_stereo", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *AudioStreamSample) SetData(data *PoolByteArray) {
	log.Println("Calling AudioStreamSample.SetData()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(data)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_data", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *AudioStreamSample) SetFormat(format int64) {
	log.Println("Calling AudioStreamSample.SetFormat()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(format)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_format", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *AudioStreamSample) SetLoopBegin(loopBegin int64) {
	log.Println("Calling AudioStreamSample.SetLoopBegin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(loopBegin)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_loop_begin", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *AudioStreamSample) SetLoopEnd(loopEnd int64) {
	log.Println("Calling AudioStreamSample.SetLoopEnd()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(loopEnd)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_loop_end", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *AudioStreamSample) SetLoopMode(loopMode int64) {
	log.Println("Calling AudioStreamSample.SetLoopMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(loopMode)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_loop_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *AudioStreamSample) SetMixRate(mixRate int64) {
	log.Println("Calling AudioStreamSample.SetMixRate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mixRate)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_mix_rate", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *AudioStreamSample) SetStereo(stereo bool) {
	log.Println("Calling AudioStreamSample.SetStereo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(stereo)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_stereo", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   AudioStreamSampleImplementer is an interface for AudioStreamSample objects.
*/
type AudioStreamSampleImplementer interface {
	class.Class
}