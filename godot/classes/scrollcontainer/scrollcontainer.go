//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package scrollcontainer

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

	"github.com/shadowapex/godot-go/godot/classes/container"
)

/*
   A ScrollContainer node with a [Control] child and scrollbar child ([HScrollbar], [VScrollBar], or both) will only draw the Control within the ScrollContainer area. Scrollbars will automatically be drawn at the right (for vertical) or bottom (for horizontal) and will enable dragging to move the viewable Control (and its children) within the ScrollContainer. Scrollbars will also automatically resize the grabber based on the minimum_size of the Control relative to the ScrollContainer. Works great with a [Panel] control. You can set EXPAND on children size flags, so they will upscale to ScrollContainer size if ScrollContainer size is bigger (scroll is invisible for chosen dimension).
*/
type ScrollContainer struct {
	container.Container
}

func (o *ScrollContainer) baseClass() string {
	return "ScrollContainer"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *ScrollContainer) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *ScrollContainer) getOwner() *C.godot_object {
	return o.owner
}

/*
   Undocumented
*/
func (o *ScrollContainer) X_GuiInput(arg0 *InputEvent) {
	log.Println("Calling ScrollContainer.X_GuiInput()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_gui_input", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScrollContainer) X_ScrollMoved(arg0 float64) {
	log.Println("Calling ScrollContainer.X_ScrollMoved()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_scroll_moved", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScrollContainer) X_UpdateScrollbarPosition() {
	log.Println("Calling ScrollContainer.X_UpdateScrollbarPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_update_scrollbar_position", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Return current horizontal scroll value.
*/
func (o *ScrollContainer) GetHScroll() int64 {
	log.Println("Calling ScrollContainer.GetHScroll()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_h_scroll", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return current vertical scroll value.
*/
func (o *ScrollContainer) GetVScroll() int64 {
	log.Println("Calling ScrollContainer.GetVScroll()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_v_scroll", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return true if horizontal scroll is allowed.
*/
func (o *ScrollContainer) IsHScrollEnabled() bool {
	log.Println("Calling ScrollContainer.IsHScrollEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_h_scroll_enabled", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return true if vertical scroll is allowed.
*/
func (o *ScrollContainer) IsVScrollEnabled() bool {
	log.Println("Calling ScrollContainer.IsVScrollEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_v_scroll_enabled", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Set allows horizontal scroll.
*/
func (o *ScrollContainer) SetEnableHScroll(enable bool) {
	log.Println("Calling ScrollContainer.SetEnableHScroll()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_enable_h_scroll", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set allows vertical scroll.
*/
func (o *ScrollContainer) SetEnableVScroll(enable bool) {
	log.Println("Calling ScrollContainer.SetEnableVScroll()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_enable_v_scroll", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set horizontal scroll value.
*/
func (o *ScrollContainer) SetHScroll(val int64) {
	log.Println("Calling ScrollContainer.SetHScroll()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(val)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_h_scroll", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set vertical scroll value.
*/
func (o *ScrollContainer) SetVScroll(val int64) {
	log.Println("Calling ScrollContainer.SetVScroll()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(val)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_v_scroll", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   ScrollContainerImplementer is an interface for ScrollContainer objects.
*/
type ScrollContainerImplementer interface {
	class.Class
}