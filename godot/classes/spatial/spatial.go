//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package spatial

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
   Most basic 3D game object, with a 3D [Transform] and visibility settings. All other 3D game objects inherit from Spatial. Use Spatial as a parent node to move, scale, rotate and show/hide children in a 3D project.
*/
type Spatial struct {
	node.Node
}

func (o *Spatial) baseClass() string {
	return "Spatial"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *Spatial) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *Spatial) getOwner() *C.godot_object {
	return o.owner
}

/*
   Undocumented
*/
func (o *Spatial) X_GetRotationDeg() *Vector3 {
	log.Println("Calling Spatial.X_GetRotationDeg()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_rotation_deg", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Spatial) X_SetRotationDeg(rotationDeg *Vector3) {
	log.Println("Calling Spatial.X_SetRotationDeg()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(rotationDeg)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_set_rotation_deg", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Spatial) X_UpdateGizmo() {
	log.Println("Calling Spatial.X_UpdateGizmo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_update_gizmo", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Returns the SpatialGizmo for this node. Used for example in [EditorSpatialGizmo] as custom visualization and editing handles in Editor.
*/
func (o *Spatial) GetGizmo() *SpatialGizmo {
	log.Println("Calling Spatial.GetGizmo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_gizmo", goArguments, "*SpatialGizmo")

	returnValue := goRet.Interface().(*SpatialGizmo)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the global transform, relative to worldspace.
*/
func (o *Spatial) GetGlobalTransform() *Transform {
	log.Println("Calling Spatial.GetGlobalTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_global_transform", goArguments, "*Transform")

	returnValue := goRet.Interface().(*Transform)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the parent [Spatial], or an empty [Object] if no parent exists or parent is not of type [Spatial].
*/
func (o *Spatial) GetParentSpatial() *Spatial {
	log.Println("Calling Spatial.GetParentSpatial()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_parent_spatial", goArguments, "*Spatial")

	returnValue := goRet.Interface().(*Spatial)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the rotation (in radians).
*/
func (o *Spatial) GetRotation() *Vector3 {
	log.Println("Calling Spatial.GetRotation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_rotation", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the rotation (in degrees).
*/
func (o *Spatial) GetRotationDeg() *Vector3 {
	log.Println("Calling Spatial.GetRotationDeg()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_rotation_deg", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *Spatial) GetScale() *Vector3 {
	log.Println("Calling Spatial.GetScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_scale", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the local transform, relative to the bone parent.
*/
func (o *Spatial) GetTransform() *Transform {
	log.Println("Calling Spatial.GetTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_transform", goArguments, "*Transform")

	returnValue := goRet.Interface().(*Transform)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *Spatial) GetTranslation() *Vector3 {
	log.Println("Calling Spatial.GetTranslation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_translation", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the current [World] resource this Spatial node is registered to.
*/
func (o *Spatial) GetWorld() *World {
	log.Println("Calling Spatial.GetWorld()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_world", goArguments, "*World")

	returnValue := goRet.Interface().(*World)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Rotates the current node along normal [Vector3] by angle in radians in Global space.
*/
func (o *Spatial) GlobalRotate(normal *Vector3, radians float64) {
	log.Println("Calling Spatial.GlobalRotate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(normal)
	goArguments[1] = reflect.ValueOf(radians)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "global_rotate", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Moves the node by [Vector3] offset in Global space.
*/
func (o *Spatial) GlobalTranslate(offset *Vector3) {
	log.Println("Calling Spatial.GlobalTranslate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(offset)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "global_translate", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Disables rendering of this node. Change Spatial Visible property to false.
*/
func (o *Spatial) Hide() {
	log.Println("Calling Spatial.Hide()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "hide", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Returns whether node notifies about its local transformation changes. Spatial will not propagate this by default.
*/
func (o *Spatial) IsLocalTransformNotificationEnabled() bool {
	log.Println("Calling Spatial.IsLocalTransformNotificationEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_local_transform_notification_enabled", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns whether this node is set as Toplevel, that is whether it ignores its parent nodes transformations.
*/
func (o *Spatial) IsSetAsToplevel() bool {
	log.Println("Calling Spatial.IsSetAsToplevel()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_set_as_toplevel", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns whether the node notifies about its global and local transformation changes. Spatial will not propagate this by default.
*/
func (o *Spatial) IsTransformNotificationEnabled() bool {
	log.Println("Calling Spatial.IsTransformNotificationEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_transform_notification_enabled", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns whether the node is set to be visible.
*/
func (o *Spatial) IsVisible() bool {
	log.Println("Calling Spatial.IsVisible()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_visible", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns whether the node is visible, taking into consideration that its parents visibility.
*/
func (o *Spatial) IsVisibleInTree() bool {
	log.Println("Calling Spatial.IsVisibleInTree()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_visible_in_tree", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Rotates itself to point into direction of target position. Operations take place in global space.
*/
func (o *Spatial) LookAt(target *Vector3, up *Vector3) {
	log.Println("Calling Spatial.LookAt()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(target)
	goArguments[1] = reflect.ValueOf(up)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "look_at", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Moves the node to specified position and then rotates itself to point into direction of target position. Operations take place in global space.
*/
func (o *Spatial) LookAtFromPosition(position *Vector3, target *Vector3, up *Vector3) {
	log.Println("Calling Spatial.LookAtFromPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(position)
	goArguments[1] = reflect.ValueOf(target)
	goArguments[2] = reflect.ValueOf(up)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "look_at_from_position", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Resets this node's transformations (like scale, skew and taper) preserving its rotation and translation. Performs orthonormalization on this node [Transform3D].
*/
func (o *Spatial) Orthonormalize() {
	log.Println("Calling Spatial.Orthonormalize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "orthonormalize", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Rotates the node in local space on given normal [Vector3] by angle in radians.
*/
func (o *Spatial) Rotate(normal *Vector3, radians float64) {
	log.Println("Calling Spatial.Rotate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(normal)
	goArguments[1] = reflect.ValueOf(radians)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "rotate", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Rotates the node in local space on X axis by angle in radians.
*/
func (o *Spatial) RotateX(radians float64) {
	log.Println("Calling Spatial.RotateX()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(radians)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "rotate_x", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Rotates the node in local space on Y axis by angle in radians.
*/
func (o *Spatial) RotateY(radians float64) {
	log.Println("Calling Spatial.RotateY()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(radians)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "rotate_y", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Rotates the node in local space on Z axis by angle in radians.
*/
func (o *Spatial) RotateZ(radians float64) {
	log.Println("Calling Spatial.RotateZ()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(radians)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "rotate_z", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Makes the node ignore its parents transformations. Node transformations are only in global space.
*/
func (o *Spatial) SetAsToplevel(enable bool) {
	log.Println("Calling Spatial.SetAsToplevel()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_as_toplevel", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set [SpatialGizmo] for this node. Used for example in [EditorSpatialGizmo] as custom visualization and editing handles in Editor.
*/
func (o *Spatial) SetGizmo(gizmo *SpatialGizmo) {
	log.Println("Calling Spatial.SetGizmo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(gizmo)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_gizmo", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the transform globally, relative to world space.
*/
func (o *Spatial) SetGlobalTransform(global *Transform) {
	log.Println("Calling Spatial.SetGlobalTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(global)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_global_transform", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Reset all transformations for this node. Set its [Transform3D] to identity matrix.
*/
func (o *Spatial) SetIdentity() {
	log.Println("Calling Spatial.SetIdentity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_identity", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set whether the node ignores notification that its transformation (global or local) changed.
*/
func (o *Spatial) SetIgnoreTransformNotification(enabled bool) {
	log.Println("Calling Spatial.SetIgnoreTransformNotification()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enabled)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_ignore_transform_notification", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set whether the node notifies about its local transformation changes. Spatial will not propagate this by default.
*/
func (o *Spatial) SetNotifyLocalTransform(enable bool) {
	log.Println("Calling Spatial.SetNotifyLocalTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_notify_local_transform", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set whether the node notifies about its global and local transformation changes. Spatial will not propagate this by default.
*/
func (o *Spatial) SetNotifyTransform(enable bool) {
	log.Println("Calling Spatial.SetNotifyTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_notify_transform", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the rotation (in radians).
*/
func (o *Spatial) SetRotation(rotationRad *Vector3) {
	log.Println("Calling Spatial.SetRotation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(rotationRad)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_rotation", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the rotation (in degrees).
*/
func (o *Spatial) SetRotationDeg(rotationDeg *Vector3) {
	log.Println("Calling Spatial.SetRotationDeg()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(rotationDeg)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_rotation_deg", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the scale.
*/
func (o *Spatial) SetScale(scale *Vector3) {
	log.Println("Calling Spatial.SetScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(scale)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_scale", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the transform locally, relative to the parent spatial node.
*/
func (o *Spatial) SetTransform(local *Transform) {
	log.Println("Calling Spatial.SetTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(local)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_transform", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *Spatial) SetTranslation(translation *Vector3) {
	log.Println("Calling Spatial.SetTranslation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(translation)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_translation", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *Spatial) SetVisible(visible bool) {
	log.Println("Calling Spatial.SetVisible()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(visible)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_visible", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Enables rendering of this node. Change Spatial Visible property to "True".
*/
func (o *Spatial) Show() {
	log.Println("Calling Spatial.Show()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "show", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Transforms [Vector3] "local_point" from this node's local space to world space.
*/
func (o *Spatial) ToGlobal(localPoint *Vector3) *Vector3 {
	log.Println("Calling Spatial.ToGlobal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(localPoint)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "to_global", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Transforms [Vector3] "global_point" from world space to this node's local space.
*/
func (o *Spatial) ToLocal(globalPoint *Vector3) *Vector3 {
	log.Println("Calling Spatial.ToLocal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(globalPoint)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "to_local", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Changes the node's position by given offset [Vector3].
*/
func (o *Spatial) Translate(offset *Vector3) {
	log.Println("Calling Spatial.Translate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(offset)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "translate", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Updates the [SpatialGizmo] of this node.
*/
func (o *Spatial) UpdateGizmo() {
	log.Println("Calling Spatial.UpdateGizmo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "update_gizmo", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   SpatialImplementer is an interface for Spatial objects.
*/
type SpatialImplementer interface {
	class.Class
}