//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package scripteditor

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

	"github.com/shadowapex/godot-go/godot/classes/panelcontainer"
)

/*

 */
type ScriptEditor struct {
	panelcontainer.PanelContainer
}

func (o *ScriptEditor) baseClass() string {
	return "ScriptEditor"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *ScriptEditor) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *ScriptEditor) getOwner() *C.godot_object {
	return o.owner
}

/*
   Undocumented
*/
func (o *ScriptEditor) X_AddCallback(arg0 *Object, arg1 string, arg2 *PoolStringArray) {
	log.Println("Calling ScriptEditor.X_AddCallback()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(arg0)
	goArguments[1] = reflect.ValueOf(arg1)
	goArguments[2] = reflect.ValueOf(arg2)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_add_callback", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_AutosaveScripts() {
	log.Println("Calling ScriptEditor.X_AutosaveScripts()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_autosave_scripts", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_Breaked(arg0 bool, arg1 bool) {
	log.Println("Calling ScriptEditor.X_Breaked()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(arg0)
	goArguments[1] = reflect.ValueOf(arg1)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_breaked", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_CloseAllTabs() {
	log.Println("Calling ScriptEditor.X_CloseAllTabs()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_close_all_tabs", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_CloseCurrentTab() {
	log.Println("Calling ScriptEditor.X_CloseCurrentTab()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_close_current_tab", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_CloseDiscardCurrentTab(arg0 string) {
	log.Println("Calling ScriptEditor.X_CloseDiscardCurrentTab()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_close_discard_current_tab", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_CloseDocsTab() {
	log.Println("Calling ScriptEditor.X_CloseDocsTab()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_close_docs_tab", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_EditorPause() {
	log.Println("Calling ScriptEditor.X_EditorPause()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_editor_pause", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_EditorPlay() {
	log.Println("Calling ScriptEditor.X_EditorPlay()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_editor_play", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_EditorSettingsChanged() {
	log.Println("Calling ScriptEditor.X_EditorSettingsChanged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_editor_settings_changed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_EditorStop() {
	log.Println("Calling ScriptEditor.X_EditorStop()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_editor_stop", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_FileDialogAction(arg0 string) {
	log.Println("Calling ScriptEditor.X_FileDialogAction()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_file_dialog_action", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_GetDebugTooltip(arg0 string, arg1 *Object) string {
	log.Println("Calling ScriptEditor.X_GetDebugTooltip()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(arg0)
	goArguments[1] = reflect.ValueOf(arg1)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_debug_tooltip", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_GotoScriptLine(arg0 *Reference, arg1 int64) {
	log.Println("Calling ScriptEditor.X_GotoScriptLine()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(arg0)
	goArguments[1] = reflect.ValueOf(arg1)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_goto_script_line", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_GotoScriptLine2(arg0 int64) {
	log.Println("Calling ScriptEditor.X_GotoScriptLine2()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_goto_script_line2", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_HelpClassGoto(arg0 string) {
	log.Println("Calling ScriptEditor.X_HelpClassGoto()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_help_class_goto", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_HelpClassOpen(arg0 string) {
	log.Println("Calling ScriptEditor.X_HelpClassOpen()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_help_class_open", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_HelpIndex(arg0 string) {
	log.Println("Calling ScriptEditor.X_HelpIndex()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_help_index", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_HelpOverviewSelected(arg0 int64) {
	log.Println("Calling ScriptEditor.X_HelpOverviewSelected()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_help_overview_selected", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_HelpSearch(arg0 string) {
	log.Println("Calling ScriptEditor.X_HelpSearch()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_help_search", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_HistoryBack() {
	log.Println("Calling ScriptEditor.X_HistoryBack()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_history_back", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_HistoryForward() {
	log.Println("Calling ScriptEditor.X_HistoryForward()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_history_forward", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_LiveAutoReloadRunningScripts() {
	log.Println("Calling ScriptEditor.X_LiveAutoReloadRunningScripts()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_live_auto_reload_running_scripts", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_MembersOverviewSelected(arg0 int64) {
	log.Println("Calling ScriptEditor.X_MembersOverviewSelected()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_members_overview_selected", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_MenuOption(arg0 int64) {
	log.Println("Calling ScriptEditor.X_MenuOption()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_menu_option", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_OpenRecentScript(arg0 int64) {
	log.Println("Calling ScriptEditor.X_OpenRecentScript()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_open_recent_script", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_ReloadScripts() {
	log.Println("Calling ScriptEditor.X_ReloadScripts()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_reload_scripts", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_RequestHelp(arg0 string) {
	log.Println("Calling ScriptEditor.X_RequestHelp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_request_help", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_ResSavedCallback(arg0 *Resource) {
	log.Println("Calling ScriptEditor.X_ResSavedCallback()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_res_saved_callback", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_ResaveScripts(arg0 string) {
	log.Println("Calling ScriptEditor.X_ResaveScripts()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_resave_scripts", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_SaveHistory() {
	log.Println("Calling ScriptEditor.X_SaveHistory()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_save_history", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_ScriptChanged() {
	log.Println("Calling ScriptEditor.X_ScriptChanged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_script_changed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_ScriptCreated(arg0 *Script) {
	log.Println("Calling ScriptEditor.X_ScriptCreated()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_script_created", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_ScriptSelected(arg0 int64) {
	log.Println("Calling ScriptEditor.X_ScriptSelected()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_script_selected", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_ScriptSplitDragged(arg0 float64) {
	log.Println("Calling ScriptEditor.X_ScriptSplitDragged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_script_split_dragged", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_ShowDebugger(arg0 bool) {
	log.Println("Calling ScriptEditor.X_ShowDebugger()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_show_debugger", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_TabChanged(arg0 int64) {
	log.Println("Calling ScriptEditor.X_TabChanged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_tab_changed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_TreeChanged() {
	log.Println("Calling ScriptEditor.X_TreeChanged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_tree_changed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_UnhandledInput(arg0 *InputEvent) {
	log.Println("Calling ScriptEditor.X_UnhandledInput()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_unhandled_input", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ScriptEditor) X_UpdateScriptNames() {
	log.Println("Calling ScriptEditor.X_UpdateScriptNames()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_update_script_names", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Returns a [Script] that is currently active in editor.
*/
func (o *ScriptEditor) GetCurrentScript() *Script {
	log.Println("Calling ScriptEditor.GetCurrentScript()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_current_script", goArguments, "*Script")

	returnValue := goRet.Interface().(*Script)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns an array with all [Script] objects which are currently open in editor.
*/
func (o *ScriptEditor) GetOpenScripts() *Array {
	log.Println("Calling ScriptEditor.GetOpenScripts()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_open_scripts", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   ScriptEditorImplementer is an interface for ScriptEditor objects.
*/
type ScriptEditorImplementer interface {
	class.Class
}
