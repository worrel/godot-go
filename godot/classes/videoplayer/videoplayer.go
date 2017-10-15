//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package videoplayer

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

	"github.com/shadowapex/godot-go/godot/classes/control"
)

/*
   This control has the ability to play video streams. The only format accepted is the OGV Theora, so any other format must be converted before using in a project.
*/
type VideoPlayer struct {
	control.Control
}

func (o *VideoPlayer) baseClass() string {
	return "VideoPlayer"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VideoPlayer) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *VideoPlayer) getOwner() *C.godot_object {
	return o.owner
}

/*
   Get the selected audio track (for multitrack videos).
*/
func (o *VideoPlayer) GetAudioTrack() int64 {
	log.Println("Calling VideoPlayer.GetAudioTrack()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_audio_track", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the amount of milliseconds to store in buffer while playing.
*/
func (o *VideoPlayer) GetBufferingMsec() int64 {
	log.Println("Calling VideoPlayer.GetBufferingMsec()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_buffering_msec", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the video stream.
*/
func (o *VideoPlayer) GetStream() *VideoStream {
	log.Println("Calling VideoPlayer.GetStream()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_stream", goArguments, "*VideoStream")

	returnValue := goRet.Interface().(*VideoStream)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the name of the video stream.
*/
func (o *VideoPlayer) GetStreamName() string {
	log.Println("Calling VideoPlayer.GetStreamName()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_stream_name", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the current position of the stream, in seconds.
*/
func (o *VideoPlayer) GetStreamPosition() float64 {
	log.Println("Calling VideoPlayer.GetStreamPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_stream_position", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the current frame of the video as a [Texture].
*/
func (o *VideoPlayer) GetVideoTexture() *Texture {
	log.Println("Calling VideoPlayer.GetVideoTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_video_texture", goArguments, "*Texture")

	returnValue := goRet.Interface().(*Texture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the volume of the audio track as a linear value.
*/
func (o *VideoPlayer) GetVolume() float64 {
	log.Println("Calling VideoPlayer.GetVolume()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_volume", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the volume of the audio track in decibels.
*/
func (o *VideoPlayer) GetVolumeDb() float64 {
	log.Println("Calling VideoPlayer.GetVolumeDb()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_volume_db", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get whether or not the video is set as autoplay.
*/
func (o *VideoPlayer) HasAutoplay() bool {
	log.Println("Calling VideoPlayer.HasAutoplay()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "has_autoplay", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get whether or not the expand property is set.
*/
func (o *VideoPlayer) HasExpand() bool {
	log.Println("Calling VideoPlayer.HasExpand()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "has_expand", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get whether or not the video is paused.
*/
func (o *VideoPlayer) IsPaused() bool {
	log.Println("Calling VideoPlayer.IsPaused()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_paused", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get whether or not the video is playing.
*/
func (o *VideoPlayer) IsPlaying() bool {
	log.Println("Calling VideoPlayer.IsPlaying()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_playing", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Start the video playback.
*/
func (o *VideoPlayer) Play() {
	log.Println("Calling VideoPlayer.Play()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "play", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the audio track (for multitrack videos).
*/
func (o *VideoPlayer) SetAudioTrack(track int64) {
	log.Println("Calling VideoPlayer.SetAudioTrack()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(track)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_audio_track", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set whether this node should start playing automatically.
*/
func (o *VideoPlayer) SetAutoplay(enabled bool) {
	log.Println("Calling VideoPlayer.SetAutoplay()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enabled)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_autoplay", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the amount of milliseconds to buffer during playback.
*/
func (o *VideoPlayer) SetBufferingMsec(msec int64) {
	log.Println("Calling VideoPlayer.SetBufferingMsec()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(msec)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_buffering_msec", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the expand property. If enabled, the video will grow or shrink to fit the player size, otherwise it will play at the stream resolution.
*/
func (o *VideoPlayer) SetExpand(enable bool) {
	log.Println("Calling VideoPlayer.SetExpand()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_expand", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set whether the video should pause the playback.
*/
func (o *VideoPlayer) SetPaused(paused bool) {
	log.Println("Calling VideoPlayer.SetPaused()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(paused)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_paused", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the video stream for this player.
*/
func (o *VideoPlayer) SetStream(stream *VideoStream) {
	log.Println("Calling VideoPlayer.SetStream()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(stream)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_stream", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the audio volume as a linear value.
*/
func (o *VideoPlayer) SetVolume(volume float64) {
	log.Println("Calling VideoPlayer.SetVolume()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(volume)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_volume", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the audio volume in decibels.
*/
func (o *VideoPlayer) SetVolumeDb(db float64) {
	log.Println("Calling VideoPlayer.SetVolumeDb()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(db)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_volume_db", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Stop the video playback.
*/
func (o *VideoPlayer) Stop() {
	log.Println("Calling VideoPlayer.Stop()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "stop", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VideoPlayerImplementer is an interface for VideoPlayer objects.
*/
type VideoPlayerImplementer interface {
	class.Class
}
