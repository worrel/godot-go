//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package audiostreamrandompitch

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
   Randomly varies pitch on each start.
*/
type AudioStreamRandomPitch struct {
	audiostream.AudioStream
}

func (o *AudioStreamRandomPitch) baseClass() string {
	return "AudioStreamRandomPitch"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *AudioStreamRandomPitch) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *AudioStreamRandomPitch) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *AudioStreamRandomPitch) GetAudioStream() *AudioStream {
	log.Println("Calling AudioStreamRandomPitch.GetAudioStream()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_audio_stream", goArguments, "*AudioStream")

	returnValue := goRet.Interface().(*AudioStream)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *AudioStreamRandomPitch) GetRandomPitch() float64 {
	log.Println("Calling AudioStreamRandomPitch.GetRandomPitch()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_random_pitch", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *AudioStreamRandomPitch) SetAudioStream(stream *AudioStream) {
	log.Println("Calling AudioStreamRandomPitch.SetAudioStream()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(stream)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_audio_stream", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *AudioStreamRandomPitch) SetRandomPitch(scale float64) {
	log.Println("Calling AudioStreamRandomPitch.SetRandomPitch()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(scale)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_random_pitch", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   AudioStreamRandomPitchImplementer is an interface for AudioStreamRandomPitch objects.
*/
type AudioStreamRandomPitchImplementer interface {
	class.Class
}
