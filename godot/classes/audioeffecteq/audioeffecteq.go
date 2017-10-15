//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package audioeffecteq

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
   AudioEffectEQ gives you control over frequencies. Use it to compensate for existing deficiencies in audio. AudioEffectEQ are very useful on the Master Bus to completely master a mix and give it character. They are also very useful when a game is run on a mobile device, to adjust the mix to that kind of speakers (it can be added but disabled when headphones are plugged).
*/
type AudioEffectEQ struct {
	audioeffect.AudioEffect
}

func (o *AudioEffectEQ) baseClass() string {
	return "AudioEffectEQ"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *AudioEffectEQ) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *AudioEffectEQ) getOwner() *C.godot_object {
	return o.owner
}

/*
   Returns the number of bands of the equalizer.
*/
func (o *AudioEffectEQ) GetBandCount() int64 {
	log.Println("Calling AudioEffectEQ.GetBandCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_band_count", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the band's gain at the specified index, in dB.
*/
func (o *AudioEffectEQ) GetBandGainDb(bandIdx int64) float64 {
	log.Println("Calling AudioEffectEQ.GetBandGainDb()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(bandIdx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_band_gain_db", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Sets band's gain at the specified index, in dB.
*/
func (o *AudioEffectEQ) SetBandGainDb(bandIdx int64, volumeDb float64) {
	log.Println("Calling AudioEffectEQ.SetBandGainDb()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(bandIdx)
	goArguments[1] = reflect.ValueOf(volumeDb)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_band_gain_db", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   AudioEffectEQImplementer is an interface for AudioEffectEQ objects.
*/
type AudioEffectEQImplementer interface {
	class.Class
}