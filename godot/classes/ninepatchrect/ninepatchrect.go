//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package ninepatchrect

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
   Better known as 9-slice panels, NinePatchRect produces clean panels of any size, based on a small texture. To do so, it splits the texture in a 3 by 3 grid. When you scale the node, it tiles the texture's sides horizontally or vertically, the center on both axes but it doesn't scale or tile the corners.
*/
type NinePatchRect struct {
	control.Control
}

func (o *NinePatchRect) baseClass() string {
	return "NinePatchRect"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *NinePatchRect) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *NinePatchRect) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *NinePatchRect) GetHAxisStretchMode() int64 {
	log.Println("Calling NinePatchRect.GetHAxisStretchMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_h_axis_stretch_mode", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *NinePatchRect) GetPatchMargin(margin int64) int64 {
	log.Println("Calling NinePatchRect.GetPatchMargin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(margin)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_patch_margin", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *NinePatchRect) GetRegionRect() *Rect2 {
	log.Println("Calling NinePatchRect.GetRegionRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_region_rect", goArguments, "*Rect2")

	returnValue := goRet.Interface().(*Rect2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *NinePatchRect) GetTexture() *Texture {
	log.Println("Calling NinePatchRect.GetTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_texture", goArguments, "*Texture")

	returnValue := goRet.Interface().(*Texture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *NinePatchRect) GetVAxisStretchMode() int64 {
	log.Println("Calling NinePatchRect.GetVAxisStretchMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_v_axis_stretch_mode", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *NinePatchRect) IsDrawCenterEnabled() bool {
	log.Println("Calling NinePatchRect.IsDrawCenterEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_draw_center_enabled", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *NinePatchRect) SetDrawCenter(drawCenter bool) {
	log.Println("Calling NinePatchRect.SetDrawCenter()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(drawCenter)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_draw_center", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *NinePatchRect) SetHAxisStretchMode(mode int64) {
	log.Println("Calling NinePatchRect.SetHAxisStretchMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_h_axis_stretch_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *NinePatchRect) SetPatchMargin(margin int64, value int64) {
	log.Println("Calling NinePatchRect.SetPatchMargin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(margin)
	goArguments[1] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_patch_margin", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *NinePatchRect) SetRegionRect(rect *Rect2) {
	log.Println("Calling NinePatchRect.SetRegionRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(rect)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_region_rect", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *NinePatchRect) SetTexture(texture *Texture) {
	log.Println("Calling NinePatchRect.SetTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(texture)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_texture", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *NinePatchRect) SetVAxisStretchMode(mode int64) {
	log.Println("Calling NinePatchRect.SetVAxisStretchMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_v_axis_stretch_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   NinePatchRectImplementer is an interface for NinePatchRect objects.
*/
type NinePatchRectImplementer interface {
	class.Class
}