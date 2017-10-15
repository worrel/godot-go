//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package canvasitem

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
   Base class of anything 2D. Canvas items are laid out in a tree and children inherit and extend the transform of their parent. CanvasItem is extended by [Control], for anything GUI related, and by [Node2D] for anything 2D engine related. Any CanvasItem can draw. For this, the "update" function must be called, then NOTIFICATION_DRAW will be received on idle time to request redraw. Because of this, canvas items don't need to be redraw on every frame, improving the performance significantly. Several functions for drawing on the CanvasItem are provided (see draw_* functions). They can only be used inside the notification, signal or _draw() overrides function, though. Canvas items are draw in tree order. By default, children are on top of their parents so a root CanvasItem will be drawn behind everything (this can be changed per item though). Canvas items can also be hidden (hiding also their subtree). They provide many means for changing standard parameters such as opacity (for it and the subtree) and self opacity, blend mode. Ultimately, a transform notification can be requested, which will notify the node that its global position changed in case the parent tree changed.
*/
type CanvasItem struct {
	node.Node
}

func (o *CanvasItem) baseClass() string {
	return "CanvasItem"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *CanvasItem) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *CanvasItem) getOwner() *C.godot_object {
	return o.owner
}

/*
   Called (if exists) to draw the canvas item.
*/
func (o *CanvasItem) X_Draw() {
	log.Println("Calling CanvasItem.X_Draw()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_draw", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *CanvasItem) X_IsOnTop() bool {
	log.Println("Calling CanvasItem.X_IsOnTop()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_is_on_top", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *CanvasItem) X_SetOnTop(onTop bool) {
	log.Println("Calling CanvasItem.X_SetOnTop()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(onTop)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_set_on_top", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *CanvasItem) X_ToplevelRaiseSelf() {
	log.Println("Calling CanvasItem.X_ToplevelRaiseSelf()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_toplevel_raise_self", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *CanvasItem) X_UpdateCallback() {
	log.Println("Calling CanvasItem.X_UpdateCallback()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_update_callback", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Draw a string character using a custom font. Returns the advance, depending on the char width and kerning with an optional next char.
*/
func (o *CanvasItem) DrawChar(font *Font, position *Vector2, char string, next string, modulate *Color) float64 {
	log.Println("Calling CanvasItem.DrawChar()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 5, 5)
	goArguments[0] = reflect.ValueOf(font)
	goArguments[1] = reflect.ValueOf(position)
	goArguments[2] = reflect.ValueOf(char)
	goArguments[3] = reflect.ValueOf(next)
	goArguments[4] = reflect.ValueOf(modulate)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "draw_char", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Draw a colored circle.
*/
func (o *CanvasItem) DrawCircle(position *Vector2, radius float64, color *Color) {
	log.Println("Calling CanvasItem.DrawCircle()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(position)
	goArguments[1] = reflect.ValueOf(radius)
	goArguments[2] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "draw_circle", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Draw a colored polygon of any amount of points, convex or concave.
*/
func (o *CanvasItem) DrawColoredPolygon(points *PoolVector2Array, color *Color, uvs *PoolVector2Array, texture *Texture, normalMap *Texture, antialiased bool) {
	log.Println("Calling CanvasItem.DrawColoredPolygon()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 6, 6)
	goArguments[0] = reflect.ValueOf(points)
	goArguments[1] = reflect.ValueOf(color)
	goArguments[2] = reflect.ValueOf(uvs)
	goArguments[3] = reflect.ValueOf(texture)
	goArguments[4] = reflect.ValueOf(normalMap)
	goArguments[5] = reflect.ValueOf(antialiased)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "draw_colored_polygon", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Draw a line from a 2D point to another, with a given color and width. It can be optionally antialiased.
*/
func (o *CanvasItem) DrawLine(from *Vector2, to *Vector2, color *Color, width float64, antialiased bool) {
	log.Println("Calling CanvasItem.DrawLine()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 5, 5)
	goArguments[0] = reflect.ValueOf(from)
	goArguments[1] = reflect.ValueOf(to)
	goArguments[2] = reflect.ValueOf(color)
	goArguments[3] = reflect.ValueOf(width)
	goArguments[4] = reflect.ValueOf(antialiased)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "draw_line", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Draw a polygon of any amount of points, convex or concave.
*/
func (o *CanvasItem) DrawPolygon(points *PoolVector2Array, colors *PoolColorArray, uvs *PoolVector2Array, texture *Texture, normalMap *Texture, antialiased bool) {
	log.Println("Calling CanvasItem.DrawPolygon()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 6, 6)
	goArguments[0] = reflect.ValueOf(points)
	goArguments[1] = reflect.ValueOf(colors)
	goArguments[2] = reflect.ValueOf(uvs)
	goArguments[3] = reflect.ValueOf(texture)
	goArguments[4] = reflect.ValueOf(normalMap)
	goArguments[5] = reflect.ValueOf(antialiased)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "draw_polygon", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *CanvasItem) DrawPolyline(points *PoolVector2Array, color *Color, width float64, antialiased bool) {
	log.Println("Calling CanvasItem.DrawPolyline()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(points)
	goArguments[1] = reflect.ValueOf(color)
	goArguments[2] = reflect.ValueOf(width)
	goArguments[3] = reflect.ValueOf(antialiased)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "draw_polyline", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *CanvasItem) DrawPolylineColors(points *PoolVector2Array, colors *PoolColorArray, width float64, antialiased bool) {
	log.Println("Calling CanvasItem.DrawPolylineColors()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(points)
	goArguments[1] = reflect.ValueOf(colors)
	goArguments[2] = reflect.ValueOf(width)
	goArguments[3] = reflect.ValueOf(antialiased)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "draw_polyline_colors", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Draw a custom primitive, 1 point for a point, 2 points for a line, 3 points for a triangle and 4 points for a quad.
*/
func (o *CanvasItem) DrawPrimitive(points *PoolVector2Array, colors *PoolColorArray, uvs *PoolVector2Array, texture *Texture, width float64, normalMap *Texture) {
	log.Println("Calling CanvasItem.DrawPrimitive()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 6, 6)
	goArguments[0] = reflect.ValueOf(points)
	goArguments[1] = reflect.ValueOf(colors)
	goArguments[2] = reflect.ValueOf(uvs)
	goArguments[3] = reflect.ValueOf(texture)
	goArguments[4] = reflect.ValueOf(width)
	goArguments[5] = reflect.ValueOf(normalMap)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "draw_primitive", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Draw a colored rectangle.
*/
func (o *CanvasItem) DrawRect(rect *Rect2, color *Color, filled bool) {
	log.Println("Calling CanvasItem.DrawRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(rect)
	goArguments[1] = reflect.ValueOf(color)
	goArguments[2] = reflect.ValueOf(filled)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "draw_rect", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set a custom transform for drawing. Anything drawn afterwards will be transformed by this.
*/
func (o *CanvasItem) DrawSetTransform(position *Vector2, rotation float64, scale *Vector2) {
	log.Println("Calling CanvasItem.DrawSetTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(position)
	goArguments[1] = reflect.ValueOf(rotation)
	goArguments[2] = reflect.ValueOf(scale)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "draw_set_transform", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *CanvasItem) DrawSetTransformMatrix(xform *Transform2D) {
	log.Println("Calling CanvasItem.DrawSetTransformMatrix()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(xform)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "draw_set_transform_matrix", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Draw a string using a custom font.
*/
func (o *CanvasItem) DrawString(font *Font, position *Vector2, text string, modulate *Color, clipW int64) {
	log.Println("Calling CanvasItem.DrawString()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 5, 5)
	goArguments[0] = reflect.ValueOf(font)
	goArguments[1] = reflect.ValueOf(position)
	goArguments[2] = reflect.ValueOf(text)
	goArguments[3] = reflect.ValueOf(modulate)
	goArguments[4] = reflect.ValueOf(clipW)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "draw_string", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Draw a styled rectangle.
*/
func (o *CanvasItem) DrawStyleBox(styleBox *StyleBox, rect *Rect2) {
	log.Println("Calling CanvasItem.DrawStyleBox()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(styleBox)
	goArguments[1] = reflect.ValueOf(rect)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "draw_style_box", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Draw a texture at a given position.
*/
func (o *CanvasItem) DrawTexture(texture *Texture, position *Vector2, modulate *Color, normalMap *Texture) {
	log.Println("Calling CanvasItem.DrawTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(texture)
	goArguments[1] = reflect.ValueOf(position)
	goArguments[2] = reflect.ValueOf(modulate)
	goArguments[3] = reflect.ValueOf(normalMap)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "draw_texture", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Draw a textured rectangle at a given position, optionally modulated by a color. Transpose swaps the x and y coordinates when reading the texture.
*/
func (o *CanvasItem) DrawTextureRect(texture *Texture, rect *Rect2, tile bool, modulate *Color, transpose bool, normalMap *Texture) {
	log.Println("Calling CanvasItem.DrawTextureRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 6, 6)
	goArguments[0] = reflect.ValueOf(texture)
	goArguments[1] = reflect.ValueOf(rect)
	goArguments[2] = reflect.ValueOf(tile)
	goArguments[3] = reflect.ValueOf(modulate)
	goArguments[4] = reflect.ValueOf(transpose)
	goArguments[5] = reflect.ValueOf(normalMap)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "draw_texture_rect", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Draw a textured rectangle region at a given position, optionally modulated by a color. Transpose swaps the x and y coordinates when reading the texture.
*/
func (o *CanvasItem) DrawTextureRectRegion(texture *Texture, rect *Rect2, srcRect *Rect2, modulate *Color, transpose bool, normalMap *Texture, clipUv bool) {
	log.Println("Calling CanvasItem.DrawTextureRectRegion()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 7, 7)
	goArguments[0] = reflect.ValueOf(texture)
	goArguments[1] = reflect.ValueOf(rect)
	goArguments[2] = reflect.ValueOf(srcRect)
	goArguments[3] = reflect.ValueOf(modulate)
	goArguments[4] = reflect.ValueOf(transpose)
	goArguments[5] = reflect.ValueOf(normalMap)
	goArguments[6] = reflect.ValueOf(clipUv)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "draw_texture_rect_region", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Used for editing, returns an opaque value representing the transform state.
*/
func (o *CanvasItem) EditGetState() *Variant {
	log.Println("Calling CanvasItem.EditGetState()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "edit_get_state", goArguments, "*Variant")

	returnValue := goRet.Interface().(*Variant)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Used for editing, handle rotation.
*/
func (o *CanvasItem) EditRotate(degrees float64) {
	log.Println("Calling CanvasItem.EditRotate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(degrees)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "edit_rotate", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *CanvasItem) EditSetRect(rect *Rect2) {
	log.Println("Calling CanvasItem.EditSetRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(rect)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "edit_set_rect", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the transform state of this CanvasItem. For [Node2D], this is an [Array] with (in order) a [Vector2] for position, a float for rotation (radians) and another [Vector2] for scale. For [Control] this is a [Rect2] with the position and size.
*/
func (o *CanvasItem) EditSetState(state *Variant) {
	log.Println("Calling CanvasItem.EditSetState()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(state)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "edit_set_state", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Return the [RID] of the [World2D] canvas where this item is in.
*/
func (o *CanvasItem) GetCanvas() *RID {
	log.Println("Calling CanvasItem.GetCanvas()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_canvas", goArguments, "*RID")

	returnValue := goRet.Interface().(*RID)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the canvas item RID used by [VisualServer] for this item.
*/
func (o *CanvasItem) GetCanvasItem() *RID {
	log.Println("Calling CanvasItem.GetCanvasItem()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_canvas_item", goArguments, "*RID")

	returnValue := goRet.Interface().(*RID)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the transform matrix of this item's canvas.
*/
func (o *CanvasItem) GetCanvasTransform() *Transform2D {
	log.Println("Calling CanvasItem.GetCanvasTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_canvas_transform", goArguments, "*Transform2D")

	returnValue := goRet.Interface().(*Transform2D)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the global position of the mouse.
*/
func (o *CanvasItem) GetGlobalMousePosition() *Vector2 {
	log.Println("Calling CanvasItem.GetGlobalMousePosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_global_mouse_position", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the global transform matrix of this item.
*/
func (o *CanvasItem) GetGlobalTransform() *Transform2D {
	log.Println("Calling CanvasItem.GetGlobalTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_global_transform", goArguments, "*Transform2D")

	returnValue := goRet.Interface().(*Transform2D)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the global transform matrix of this item in relation to the canvas.
*/
func (o *CanvasItem) GetGlobalTransformWithCanvas() *Transform2D {
	log.Println("Calling CanvasItem.GetGlobalTransformWithCanvas()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_global_transform_with_canvas", goArguments, "*Transform2D")

	returnValue := goRet.Interface().(*Transform2D)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get a [Rect2] with the boundaries of this item and its children.
*/
func (o *CanvasItem) GetItemAndChildrenRect() *Rect2 {
	log.Println("Calling CanvasItem.GetItemAndChildrenRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_item_and_children_rect", goArguments, "*Rect2")

	returnValue := goRet.Interface().(*Rect2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return a rect containing the editable boundaries of the item.
*/
func (o *CanvasItem) GetItemRect() *Rect2 {
	log.Println("Calling CanvasItem.GetItemRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_item_rect", goArguments, "*Rect2")

	returnValue := goRet.Interface().(*Rect2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get this item's light mask number.
*/
func (o *CanvasItem) GetLightMask() int64 {
	log.Println("Calling CanvasItem.GetLightMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_light_mask", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the mouse position relative to this item's position.
*/
func (o *CanvasItem) GetLocalMousePosition() *Vector2 {
	log.Println("Calling CanvasItem.GetLocalMousePosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_local_mouse_position", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the material of this item.
*/
func (o *CanvasItem) GetMaterial() *Material {
	log.Println("Calling CanvasItem.GetMaterial()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_material", goArguments, "*Material")

	returnValue := goRet.Interface().(*Material)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the modulate of the CanvasItem, which affects children items too.
*/
func (o *CanvasItem) GetModulate() *Color {
	log.Println("Calling CanvasItem.GetModulate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_modulate", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the self-modulate of the CanvasItem.
*/
func (o *CanvasItem) GetSelfModulate() *Color {
	log.Println("Calling CanvasItem.GetSelfModulate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_self_modulate", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the transform matrix of this item.
*/
func (o *CanvasItem) GetTransform() *Transform2D {
	log.Println("Calling CanvasItem.GetTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_transform", goArguments, "*Transform2D")

	returnValue := goRet.Interface().(*Transform2D)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get whether this item uses its parent's material.
*/
func (o *CanvasItem) GetUseParentMaterial() bool {
	log.Println("Calling CanvasItem.GetUseParentMaterial()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_use_parent_material", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the viewport's boundaries as a [Rect2].
*/
func (o *CanvasItem) GetViewportRect() *Rect2 {
	log.Println("Calling CanvasItem.GetViewportRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_viewport_rect", goArguments, "*Rect2")

	returnValue := goRet.Interface().(*Rect2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get this item's transform in relation to the viewport.
*/
func (o *CanvasItem) GetViewportTransform() *Transform2D {
	log.Println("Calling CanvasItem.GetViewportTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_viewport_transform", goArguments, "*Transform2D")

	returnValue := goRet.Interface().(*Transform2D)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the [World2D] where this item is in.
*/
func (o *CanvasItem) GetWorld2D() *World2D {
	log.Println("Calling CanvasItem.GetWorld2D()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_world_2d", goArguments, "*World2D")

	returnValue := goRet.Interface().(*World2D)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Hide the CanvasItem currently visible.
*/
func (o *CanvasItem) Hide() {
	log.Println("Calling CanvasItem.Hide()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "hide", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Return whether the item is drawn behind its parent.
*/
func (o *CanvasItem) IsDrawBehindParentEnabled() bool {
	log.Println("Calling CanvasItem.IsDrawBehindParentEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_draw_behind_parent_enabled", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *CanvasItem) IsLocalTransformNotificationEnabled() bool {
	log.Println("Calling CanvasItem.IsLocalTransformNotificationEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_local_transform_notification_enabled", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return if set as toplevel. See [method set_as_toplevel].
*/
func (o *CanvasItem) IsSetAsToplevel() bool {
	log.Println("Calling CanvasItem.IsSetAsToplevel()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_set_as_toplevel", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *CanvasItem) IsTransformNotificationEnabled() bool {
	log.Println("Calling CanvasItem.IsTransformNotificationEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_transform_notification_enabled", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return true if this CanvasItem is visible. It may be invisible because itself or a parent canvas item is hidden.
*/
func (o *CanvasItem) IsVisible() bool {
	log.Println("Calling CanvasItem.IsVisible()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_visible", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *CanvasItem) IsVisibleInTree() bool {
	log.Println("Calling CanvasItem.IsVisibleInTree()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_visible_in_tree", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *CanvasItem) MakeCanvasPositionLocal(screenPoint *Vector2) *Vector2 {
	log.Println("Calling CanvasItem.MakeCanvasPositionLocal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(screenPoint)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "make_canvas_position_local", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *CanvasItem) MakeInputLocal(event *InputEvent) *InputEvent {
	log.Println("Calling CanvasItem.MakeInputLocal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(event)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "make_input_local", goArguments, "*InputEvent")

	returnValue := goRet.Interface().(*InputEvent)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Set as top level. This means that it will not inherit transform from parent canvas items.
*/
func (o *CanvasItem) SetAsToplevel(enable bool) {
	log.Println("Calling CanvasItem.SetAsToplevel()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_as_toplevel", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set whether the canvas item is drawn behind its parent.
*/
func (o *CanvasItem) SetDrawBehindParent(enable bool) {
	log.Println("Calling CanvasItem.SetDrawBehindParent()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_draw_behind_parent", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the ligtht mask number of this item.
*/
func (o *CanvasItem) SetLightMask(lightMask int64) {
	log.Println("Calling CanvasItem.SetLightMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(lightMask)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_light_mask", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the material of this item.
*/
func (o *CanvasItem) SetMaterial(material *Material) {
	log.Println("Calling CanvasItem.SetMaterial()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(material)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_material", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the modulate of the CanvasItem. This [i]affects[/i] the modulation of children items.
*/
func (o *CanvasItem) SetModulate(modulate *Color) {
	log.Println("Calling CanvasItem.SetModulate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(modulate)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_modulate", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *CanvasItem) SetNotifyLocalTransform(enable bool) {
	log.Println("Calling CanvasItem.SetNotifyLocalTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_notify_local_transform", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *CanvasItem) SetNotifyTransform(enable bool) {
	log.Println("Calling CanvasItem.SetNotifyTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_notify_transform", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the self-modulate of the CanvasItem. This does not affect the modulation of children items.
*/
func (o *CanvasItem) SetSelfModulate(selfModulate *Color) {
	log.Println("Calling CanvasItem.SetSelfModulate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(selfModulate)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_self_modulate", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set whether or not this item should use its parent's material.
*/
func (o *CanvasItem) SetUseParentMaterial(enable bool) {
	log.Println("Calling CanvasItem.SetUseParentMaterial()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_use_parent_material", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set whether this item should be visible or not. Note that a hidden CanvasItem will make all children hidden too, so no matter what is set here this item won't be shown if its parent or grandparents nodes are hidden.
*/
func (o *CanvasItem) SetVisible(visible bool) {
	log.Println("Calling CanvasItem.SetVisible()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(visible)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_visible", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Show the CanvasItem currently hidden.
*/
func (o *CanvasItem) Show() {
	log.Println("Calling CanvasItem.Show()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "show", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Queue the CanvasItem for update. [code]NOTIFICATION_DRAW[/code] will be called on idle time to request redraw.
*/
func (o *CanvasItem) Update() {
	log.Println("Calling CanvasItem.Update()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "update", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   CanvasItemImplementer is an interface for CanvasItem objects.
*/
type CanvasItemImplementer interface {
	class.Class
}
