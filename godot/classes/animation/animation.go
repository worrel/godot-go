//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package animation

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

	"github.com/shadowapex/godot-go/godot/classes/resource"
)

/*
   An Animation resource contains data used to animate everything in the engine. Animations are divided into tracks, and each track must be linked to a node. The state of that node can be changed through time, by adding timed keys (events) to the track. Animations are just data containers, and must be added to odes such as an [AnimationPlayer] or [AnimationTreePlayer] to be played back.
*/
type Animation struct {
	resource.Resource
}

func (o *Animation) baseClass() string {
	return "Animation"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *Animation) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *Animation) getOwner() *C.godot_object {
	return o.owner
}

/*
   Add a track to the Animation. The track type must be specified as any of the values in the TYPE_* enumeration.
*/
func (o *Animation) AddTrack(aType int64, atPosition int64) int64 {
	log.Println("Calling Animation.AddTrack()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(aType)
	goArguments[1] = reflect.ValueOf(atPosition)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "add_track", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Clear the animation (clear all tracks and reset all).
*/
func (o *Animation) Clear() {
	log.Println("Calling Animation.Clear()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "clear", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Return the index of the specified track. If the track is not found, return -1.
*/
func (o *Animation) FindTrack(path *NodePath) int64 {
	log.Println("Calling Animation.FindTrack()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(path)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "find_track", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the total length of the animation (in seconds).
*/
func (o *Animation) GetLength() float64 {
	log.Println("Calling Animation.GetLength()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_length", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the animation step value.
*/
func (o *Animation) GetStep() float64 {
	log.Println("Calling Animation.GetStep()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_step", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the amount of tracks in the animation.
*/
func (o *Animation) GetTrackCount() int64 {
	log.Println("Calling Animation.GetTrackCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_track_count", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return whether the animation has the loop flag set.
*/
func (o *Animation) HasLoop() bool {
	log.Println("Calling Animation.HasLoop()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "has_loop", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return all the key indices of a method track, given a position and delta time.
*/
func (o *Animation) MethodTrackGetKeyIndices(idx int64, timeSec float64, delta float64) *PoolIntArray {
	log.Println("Calling Animation.MethodTrackGetKeyIndices()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(timeSec)
	goArguments[2] = reflect.ValueOf(delta)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "method_track_get_key_indices", goArguments, "*PoolIntArray")

	returnValue := goRet.Interface().(*PoolIntArray)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the method name of a method track.
*/
func (o *Animation) MethodTrackGetName(idx int64, keyIdx int64) string {
	log.Println("Calling Animation.MethodTrackGetName()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(keyIdx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "method_track_get_name", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the arguments values to be called on a method track for a given key in a given track.
*/
func (o *Animation) MethodTrackGetParams(idx int64, keyIdx int64) *Array {
	log.Println("Calling Animation.MethodTrackGetParams()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(keyIdx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "method_track_get_params", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Remove a track by specifying the track index.
*/
func (o *Animation) RemoveTrack(idx int64) {
	log.Println("Calling Animation.RemoveTrack()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "remove_track", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the total length of the animation (in seconds). Note that length is not delimited by the last key, as this one may be before or after the end to ensure correct interpolation and looping.
*/
func (o *Animation) SetLength(timeSec float64) {
	log.Println("Calling Animation.SetLength()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(timeSec)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_length", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set a flag indicating that the animation must loop. This is uses for correct interpolation of animation cycles, and for hinting the player that it must restart the animation.
*/
func (o *Animation) SetLoop(enabled bool) {
	log.Println("Calling Animation.SetLoop()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enabled)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_loop", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the animation step value.
*/
func (o *Animation) SetStep(sizeSec float64) {
	log.Println("Calling Animation.SetStep()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(sizeSec)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_step", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Find the key index by time in a given track. Optionally, only find it if the exact time is given.
*/
func (o *Animation) TrackFindKey(idx int64, time float64, exact bool) int64 {
	log.Println("Calling Animation.TrackFindKey()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(time)
	goArguments[2] = reflect.ValueOf(exact)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "track_find_key", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *Animation) TrackGetInterpolationLoopWrap(idx int64) bool {
	log.Println("Calling Animation.TrackGetInterpolationLoopWrap()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "track_get_interpolation_loop_wrap", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the interpolation type of a given track, from the INTERPOLATION_* enum.
*/
func (o *Animation) TrackGetInterpolationType(idx int64) int64 {
	log.Println("Calling Animation.TrackGetInterpolationType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "track_get_interpolation_type", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the amount of keys in a given track.
*/
func (o *Animation) TrackGetKeyCount(idx int64) int64 {
	log.Println("Calling Animation.TrackGetKeyCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "track_get_key_count", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the time at which the key is located.
*/
func (o *Animation) TrackGetKeyTime(idx int64, keyIdx int64) float64 {
	log.Println("Calling Animation.TrackGetKeyTime()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(keyIdx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "track_get_key_time", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the transition curve (easing) for a specific key (see built-in math function "ease").
*/
func (o *Animation) TrackGetKeyTransition(idx int64, keyIdx int64) float64 {
	log.Println("Calling Animation.TrackGetKeyTransition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(keyIdx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "track_get_key_transition", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the value of a given key in a given track.
*/
func (o *Animation) TrackGetKeyValue(idx int64, keyIdx int64) *Variant {
	log.Println("Calling Animation.TrackGetKeyValue()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(keyIdx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "track_get_key_value", goArguments, "*Variant")

	returnValue := goRet.Interface().(*Variant)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the path of a track. for more information on the path format, see [method track_set_path]
*/
func (o *Animation) TrackGetPath(idx int64) *NodePath {
	log.Println("Calling Animation.TrackGetPath()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "track_get_path", goArguments, "*NodePath")

	returnValue := goRet.Interface().(*NodePath)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the type of a track.
*/
func (o *Animation) TrackGetType(idx int64) int64 {
	log.Println("Calling Animation.TrackGetType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "track_get_type", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Insert a generic key in a given track.
*/
func (o *Animation) TrackInsertKey(idx int64, time float64, key *Variant, transition float64) {
	log.Println("Calling Animation.TrackInsertKey()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(time)
	goArguments[2] = reflect.ValueOf(key)
	goArguments[3] = reflect.ValueOf(transition)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "track_insert_key", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Return true if the given track is imported. Else, return false.
*/
func (o *Animation) TrackIsImported(idx int64) bool {
	log.Println("Calling Animation.TrackIsImported()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "track_is_imported", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Move a track down.
*/
func (o *Animation) TrackMoveDown(idx int64) {
	log.Println("Calling Animation.TrackMoveDown()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "track_move_down", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Move a track up.
*/
func (o *Animation) TrackMoveUp(idx int64) {
	log.Println("Calling Animation.TrackMoveUp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "track_move_up", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Remove a key by index in a given track.
*/
func (o *Animation) TrackRemoveKey(idx int64, keyIdx int64) {
	log.Println("Calling Animation.TrackRemoveKey()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(keyIdx)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "track_remove_key", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Remove a key by position (seconds) in a given track.
*/
func (o *Animation) TrackRemoveKeyAtPosition(idx int64, position float64) {
	log.Println("Calling Animation.TrackRemoveKeyAtPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(position)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "track_remove_key_at_position", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the given track as imported or not.
*/
func (o *Animation) TrackSetImported(idx int64, imported bool) {
	log.Println("Calling Animation.TrackSetImported()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(imported)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "track_set_imported", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *Animation) TrackSetInterpolationLoopWrap(idx int64, interpolation bool) {
	log.Println("Calling Animation.TrackSetInterpolationLoopWrap()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(interpolation)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "track_set_interpolation_loop_wrap", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the interpolation type of a given track, from the INTERPOLATION_* enum.
*/
func (o *Animation) TrackSetInterpolationType(idx int64, interpolation int64) {
	log.Println("Calling Animation.TrackSetInterpolationType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(interpolation)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "track_set_interpolation_type", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the transition curve (easing) for a specific key (see built-in math function "ease").
*/
func (o *Animation) TrackSetKeyTransition(idx int64, keyIdx int64, transition float64) {
	log.Println("Calling Animation.TrackSetKeyTransition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(keyIdx)
	goArguments[2] = reflect.ValueOf(transition)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "track_set_key_transition", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the value of an existing key.
*/
func (o *Animation) TrackSetKeyValue(idx int64, key int64, value *Variant) {
	log.Println("Calling Animation.TrackSetKeyValue()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(key)
	goArguments[2] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "track_set_key_value", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the path of a track. Paths must be valid scene-tree paths to a node, and must be specified starting from the parent node of the node that will reproduce the animation. Tracks that control properties or bones must append their name after the path, separated by ":". Example: "character/skeleton:ankle" or "character/mesh:transform/local"
*/
func (o *Animation) TrackSetPath(idx int64, path *NodePath) {
	log.Println("Calling Animation.TrackSetPath()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(path)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "track_set_path", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Insert a transform key for a transform track.
*/
func (o *Animation) TransformTrackInsertKey(idx int64, time float64, location *Vector3, rotation *Quat, scale *Vector3) int64 {
	log.Println("Calling Animation.TransformTrackInsertKey()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 5, 5)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(time)
	goArguments[2] = reflect.ValueOf(location)
	goArguments[3] = reflect.ValueOf(rotation)
	goArguments[4] = reflect.ValueOf(scale)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "transform_track_insert_key", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the interpolated value of a transform track at a given time (in seconds). An array consisting of 3 elements: position ([Vector3]), rotation ([Quat]) and scale ([Vector3]).
*/
func (o *Animation) TransformTrackInterpolate(idx int64, timeSec float64) *Array {
	log.Println("Calling Animation.TransformTrackInterpolate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(timeSec)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "transform_track_interpolate", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return all the key indices of a value track, given a position and delta time.
*/
func (o *Animation) ValueTrackGetKeyIndices(idx int64, timeSec float64, delta float64) *PoolIntArray {
	log.Println("Calling Animation.ValueTrackGetKeyIndices()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(timeSec)
	goArguments[2] = reflect.ValueOf(delta)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "value_track_get_key_indices", goArguments, "*PoolIntArray")

	returnValue := goRet.Interface().(*PoolIntArray)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the update mode of a value track.
*/
func (o *Animation) ValueTrackGetUpdateMode(idx int64) int64 {
	log.Println("Calling Animation.ValueTrackGetUpdateMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "value_track_get_update_mode", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Set the update mode (UPDATE_*) of a value track.
*/
func (o *Animation) ValueTrackSetUpdateMode(idx int64, mode int64) {
	log.Println("Calling Animation.ValueTrackSetUpdateMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "value_track_set_update_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   AnimationImplementer is an interface for Animation objects.
*/
type AnimationImplementer interface {
	class.Class
}