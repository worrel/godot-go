//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package audioeffectchorus

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

	"github.com/shadowapex/godot-go/godot/classes/audioeffect"
)

/*
   Adds a chorus audio effect. The effect applies a filter with voices to duplicate the audio source and manipulate it through the filter.
*/
type AudioEffectChorus struct {
	audioeffect.AudioEffect
}

func (o *AudioEffectChorus) baseClass() string {
	return "AudioEffectChorus"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *AudioEffectChorus) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *AudioEffectChorus) getOwner() *C.godot_object {
	return o.owner
}

/*
   Returns the set dry ratio.
*/
func (o *AudioEffectChorus) GetDry() float64 {
	log.Println("Calling AudioEffectChorus.GetDry()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_dry", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the set voice count.
*/
func (o *AudioEffectChorus) GetVoiceCount() int64 {
	log.Println("Calling AudioEffectChorus.GetVoiceCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_voice_count", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the voice's set cutoff frequency.
*/
func (o *AudioEffectChorus) GetVoiceCutoffHz(voiceIdx int64) float64 {
	log.Println("Calling AudioEffectChorus.GetVoiceCutoffHz()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(voiceIdx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_voice_cutoff_hz", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the voice's set delay.
*/
func (o *AudioEffectChorus) GetVoiceDelayMs(voiceIdx int64) float64 {
	log.Println("Calling AudioEffectChorus.GetVoiceDelayMs()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(voiceIdx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_voice_delay_ms", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the voice's set filter depth.
*/
func (o *AudioEffectChorus) GetVoiceDepthMs(voiceIdx int64) float64 {
	log.Println("Calling AudioEffectChorus.GetVoiceDepthMs()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(voiceIdx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_voice_depth_ms", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the voice's set maximum volume.
*/
func (o *AudioEffectChorus) GetVoiceLevelDb(voiceIdx int64) float64 {
	log.Println("Calling AudioEffectChorus.GetVoiceLevelDb()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(voiceIdx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_voice_level_db", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the voice's set pan.
*/
func (o *AudioEffectChorus) GetVoicePan(voiceIdx int64) float64 {
	log.Println("Calling AudioEffectChorus.GetVoicePan()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(voiceIdx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_voice_pan", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the voice filter's set rate in cycles.
*/
func (o *AudioEffectChorus) GetVoiceRateHz(voiceIdx int64) float64 {
	log.Println("Calling AudioEffectChorus.GetVoiceRateHz()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(voiceIdx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_voice_rate_hz", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the set applied wetness of the effect.
*/
func (o *AudioEffectChorus) GetWet() float64 {
	log.Println("Calling AudioEffectChorus.GetWet()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_wet", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the set applied dryness of the effect.
*/
func (o *AudioEffectChorus) SetDry(amount float64) {
	log.Println("Calling AudioEffectChorus.SetDry()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(amount)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_dry", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the number of voices in the effect's filter.
*/
func (o *AudioEffectChorus) SetVoiceCount(voices int64) {
	log.Println("Calling AudioEffectChorus.SetVoiceCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(voices)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_voice_count", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the cutoff frequency of the voice. The maximum frequency the voice may affect.
*/
func (o *AudioEffectChorus) SetVoiceCutoffHz(voiceIdx int64, cutoffHz float64) {
	log.Println("Calling AudioEffectChorus.SetVoiceCutoffHz()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(voiceIdx)
	goArguments[1] = reflect.ValueOf(cutoffHz)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_voice_cutoff_hz", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the delay of the voice's signal.
*/
func (o *AudioEffectChorus) SetVoiceDelayMs(voiceIdx int64, delayMs float64) {
	log.Println("Calling AudioEffectChorus.SetVoiceDelayMs()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(voiceIdx)
	goArguments[1] = reflect.ValueOf(delayMs)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_voice_delay_ms", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the filter depth of the voice's signal.
*/
func (o *AudioEffectChorus) SetVoiceDepthMs(voiceIdx int64, depthMs float64) {
	log.Println("Calling AudioEffectChorus.SetVoiceDepthMs()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(voiceIdx)
	goArguments[1] = reflect.ValueOf(depthMs)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_voice_depth_ms", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the volume level of the voice.
*/
func (o *AudioEffectChorus) SetVoiceLevelDb(voiceIdx int64, levelDb float64) {
	log.Println("Calling AudioEffectChorus.SetVoiceLevelDb()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(voiceIdx)
	goArguments[1] = reflect.ValueOf(levelDb)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_voice_level_db", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the pan level of the voice.
*/
func (o *AudioEffectChorus) SetVoicePan(voiceIdx int64, pan float64) {
	log.Println("Calling AudioEffectChorus.SetVoicePan()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(voiceIdx)
	goArguments[1] = reflect.ValueOf(pan)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_voice_pan", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the voice filter's rate.
*/
func (o *AudioEffectChorus) SetVoiceRateHz(voiceIdx int64, rateHz float64) {
	log.Println("Calling AudioEffectChorus.SetVoiceRateHz()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(voiceIdx)
	goArguments[1] = reflect.ValueOf(rateHz)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_voice_rate_hz", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the amount of effect.
*/
func (o *AudioEffectChorus) SetWet(amount float64) {
	log.Println("Calling AudioEffectChorus.SetWet()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(amount)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_wet", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   AudioEffectChorusImplementer is an interface for AudioEffectChorus objects.
*/
type AudioEffectChorusImplementer interface {
	class.Class
}
