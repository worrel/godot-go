//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package tree

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
   This shows a tree of items that can be selected, expanded and collapsed. The tree can have multiple columns with custom controls like text editing, buttons and popups. It can be useful for structured displays and interactions. Trees are built via code, using [TreeItem] objects to create the structure. They have a single root but multiple roots can be simulated if a dummy hidden root is added. [codeblock] func _ready(): var tree = Tree.new() var root = tree.create_item() tree.set_hide_root(true) var child1 = tree.create_item(root) var child2 = tree.create_item(root) var subchild1 = tree.create_item(child1) subchild1.set_text(0, "Subchild1") [/codeblock]
*/
type Tree struct {
	control.Control
}

func (o *Tree) baseClass() string {
	return "Tree"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *Tree) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *Tree) getOwner() *C.godot_object {
	return o.owner
}

/*
   Undocumented
*/
func (o *Tree) X_GuiInput(arg0 *InputEvent) {
	log.Println("Calling Tree.X_GuiInput()")

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
func (o *Tree) X_PopupSelect(arg0 int64) {
	log.Println("Calling Tree.X_PopupSelect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_popup_select", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Tree) X_RangeClickTimeout() {
	log.Println("Calling Tree.X_RangeClickTimeout()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_range_click_timeout", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Tree) X_ScrollMoved(arg0 float64) {
	log.Println("Calling Tree.X_ScrollMoved()")

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
func (o *Tree) X_TextEditorEnter(arg0 string) {
	log.Println("Calling Tree.X_TextEditorEnter()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_text_editor_enter", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Tree) X_TextEditorModalClose() {
	log.Println("Calling Tree.X_TextEditorModalClose()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_text_editor_modal_close", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Tree) X_ValueEditorChanged(arg0 float64) {
	log.Println("Calling Tree.X_ValueEditorChanged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_value_editor_changed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Returns [code]true[/code] if the column titles are being shown.
*/
func (o *Tree) AreColumnTitlesVisible() bool {
	log.Println("Calling Tree.AreColumnTitlesVisible()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "are_column_titles_visible", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Clears the tree. This removes all items.
*/
func (o *Tree) Clear() {
	log.Println("Calling Tree.Clear()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "clear", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Create an item in the tree and add it as the last child of [code]parent[/code]. If parent is not given, it will be added as the root's last child, or it'll the be the root itself if the tree is empty.
*/
func (o *Tree) CreateItem(parent *Object) *Object {
	log.Println("Calling Tree.CreateItem()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(parent)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "create_item", goArguments, "*Object")

	returnValue := goRet.Interface().(*Object)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Makes the currently selected item visible. This will scroll the tree to make sure the selected item is visible.
*/
func (o *Tree) EnsureCursorIsVisible() {
	log.Println("Calling Tree.EnsureCursorIsVisible()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "ensure_cursor_is_visible", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Returns [code]true[/code] if a cell that is currently already selected may be selected again.
*/
func (o *Tree) GetAllowReselect() bool {
	log.Println("Calling Tree.GetAllowReselect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_allow_reselect", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if a right click can select items.
*/
func (o *Tree) GetAllowRmbSelect() bool {
	log.Println("Calling Tree.GetAllowRmbSelect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_allow_rmb_select", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the column index under the given point.
*/
func (o *Tree) GetColumnAtPosition(position *Vector2) int64 {
	log.Println("Calling Tree.GetColumnAtPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(position)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_column_at_position", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the column's title.
*/
func (o *Tree) GetColumnTitle(column int64) string {
	log.Println("Calling Tree.GetColumnTitle()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_column_title", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the column's width in pixels.
*/
func (o *Tree) GetColumnWidth(column int64) int64 {
	log.Println("Calling Tree.GetColumnWidth()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_column_width", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the amount of columns.
*/
func (o *Tree) GetColumns() int64 {
	log.Println("Calling Tree.GetColumns()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_columns", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the rectangle for custom popups. Helper to create custom cell controls that display a popup. See [method TreeItem.set_cell_mode].
*/
func (o *Tree) GetCustomPopupRect() *Rect2 {
	log.Println("Calling Tree.GetCustomPopupRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_custom_popup_rect", goArguments, "*Rect2")

	returnValue := goRet.Interface().(*Rect2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the current drop mode's flags.
*/
func (o *Tree) GetDropModeFlags() int64 {
	log.Println("Calling Tree.GetDropModeFlags()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_drop_mode_flags", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *Tree) GetDropSectionAtPosition(position *Vector2) int64 {
	log.Println("Calling Tree.GetDropSectionAtPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(position)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_drop_section_at_position", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the currently edited item. This is only available for custom cell mode.
*/
func (o *Tree) GetEdited() *TreeItem {
	log.Println("Calling Tree.GetEdited()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_edited", goArguments, "*TreeItem")

	returnValue := goRet.Interface().(*TreeItem)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the column for the currently edited item. This is only available for custom cell mode.
*/
func (o *Tree) GetEditedColumn() int64 {
	log.Println("Calling Tree.GetEditedColumn()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_edited_column", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the rectangle area for the specified item. If column is specified, only get the position and size of that column, otherwise get the rectangle containing all columns.
*/
func (o *Tree) GetItemAreaRect(item *Object, column int64) *Rect2 {
	log.Println("Calling Tree.GetItemAreaRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(item)
	goArguments[1] = reflect.ValueOf(column)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_item_area_rect", goArguments, "*Rect2")

	returnValue := goRet.Interface().(*Rect2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the tree item at the specified position (relative to the tree origin position).
*/
func (o *Tree) GetItemAtPosition(position *Vector2) *TreeItem {
	log.Println("Calling Tree.GetItemAtPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(position)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_item_at_position", goArguments, "*TreeItem")

	returnValue := goRet.Interface().(*TreeItem)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the next selected item after the given one.
*/
func (o *Tree) GetNextSelected(from *Object) *TreeItem {
	log.Println("Calling Tree.GetNextSelected()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(from)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_next_selected", goArguments, "*TreeItem")

	returnValue := goRet.Interface().(*TreeItem)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the last pressed button's index.
*/
func (o *Tree) GetPressedButton() int64 {
	log.Println("Calling Tree.GetPressedButton()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_pressed_button", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the tree's root item.
*/
func (o *Tree) GetRoot() *TreeItem {
	log.Println("Calling Tree.GetRoot()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_root", goArguments, "*TreeItem")

	returnValue := goRet.Interface().(*TreeItem)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the current scrolling position.
*/
func (o *Tree) GetScroll() *Vector2 {
	log.Println("Calling Tree.GetScroll()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_scroll", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the currently selected item.
*/
func (o *Tree) GetSelected() *TreeItem {
	log.Println("Calling Tree.GetSelected()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_selected", goArguments, "*TreeItem")

	returnValue := goRet.Interface().(*TreeItem)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the current selection's column.
*/
func (o *Tree) GetSelectedColumn() int64 {
	log.Println("Calling Tree.GetSelectedColumn()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_selected_column", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if the folding arrow is hidden.
*/
func (o *Tree) IsFoldingHidden() bool {
	log.Println("Calling Tree.IsFoldingHidden()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_folding_hidden", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   If [code]true[/code] the currently selected cell may be selected again.
*/
func (o *Tree) SetAllowReselect(allow bool) {
	log.Println("Calling Tree.SetAllowReselect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(allow)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_allow_reselect", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   If [code]true[/code] a right mouse button click can select items.
*/
func (o *Tree) SetAllowRmbSelect(allow bool) {
	log.Println("Calling Tree.SetAllowRmbSelect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(allow)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_allow_rmb_select", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   If [code]true[/code] the column will have the "Expand" flag of [Control].
*/
func (o *Tree) SetColumnExpand(column int64, expand bool) {
	log.Println("Calling Tree.SetColumnExpand()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(expand)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_column_expand", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the minimum width of a column.
*/
func (o *Tree) SetColumnMinWidth(column int64, minWidth int64) {
	log.Println("Calling Tree.SetColumnMinWidth()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(minWidth)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_column_min_width", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the title of a column.
*/
func (o *Tree) SetColumnTitle(column int64, title string) {
	log.Println("Calling Tree.SetColumnTitle()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(column)
	goArguments[1] = reflect.ValueOf(title)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_column_title", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   If [code]true[/code] column titles are visible.
*/
func (o *Tree) SetColumnTitlesVisible(visible bool) {
	log.Println("Calling Tree.SetColumnTitlesVisible()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(visible)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_column_titles_visible", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the amount of columns.
*/
func (o *Tree) SetColumns(amount int64) {
	log.Println("Calling Tree.SetColumns()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(amount)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_columns", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the drop mode as an OR combination of flags. See [code]DROP_MODE_*[/code] constants.
*/
func (o *Tree) SetDropModeFlags(flags int64) {
	log.Println("Calling Tree.SetDropModeFlags()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(flags)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_drop_mode_flags", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   If [code]true[/code] the folding arrow is hidden.
*/
func (o *Tree) SetHideFolding(hide bool) {
	log.Println("Calling Tree.SetHideFolding()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(hide)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_hide_folding", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   If [code]true[/code] the tree's root is hidden.
*/
func (o *Tree) SetHideRoot(enable bool) {
	log.Println("Calling Tree.SetHideRoot()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_hide_root", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Allow single or multiple selection. See the [code]SELECT_*[/code] constants.
*/
func (o *Tree) SetSelectMode(mode int64) {
	log.Println("Calling Tree.SetSelectMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_select_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   TreeImplementer is an interface for Tree objects.
*/
type TreeImplementer interface {
	class.Class
}
