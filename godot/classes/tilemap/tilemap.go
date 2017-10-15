//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package tilemap

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

	"github.com/shadowapex/godot-go/godot/classes/node2d"
)

/*
   Node for 2D tile-based maps. Tilemaps use a [TileSet] which contain a list of tiles (textures plus optional collision, navigation, and/or occluder shapes) which are used to create grid-based maps.
*/
type TileMap struct {
	node2d.Node2D
}

func (o *TileMap) baseClass() string {
	return "TileMap"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *TileMap) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *TileMap) getOwner() *C.godot_object {
	return o.owner
}

/*
   Undocumented
*/
func (o *TileMap) X_ClearQuadrants() {
	log.Println("Calling TileMap.X_ClearQuadrants()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_clear_quadrants", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *TileMap) X_GetOldCellSize() int64 {
	log.Println("Calling TileMap.X_GetOldCellSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_old_cell_size", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *TileMap) X_GetTileData() *PoolIntArray {
	log.Println("Calling TileMap.X_GetTileData()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_tile_data", goArguments, "*PoolIntArray")

	returnValue := goRet.Interface().(*PoolIntArray)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *TileMap) X_RecreateQuadrants() {
	log.Println("Calling TileMap.X_RecreateQuadrants()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_recreate_quadrants", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *TileMap) X_SetOldCellSize(size int64) {
	log.Println("Calling TileMap.X_SetOldCellSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(size)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_set_old_cell_size", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *TileMap) X_SetTileData(arg0 *PoolIntArray) {
	log.Println("Calling TileMap.X_SetTileData()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_set_tile_data", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *TileMap) X_UpdateDirtyQuadrants() {
	log.Println("Calling TileMap.X_UpdateDirtyQuadrants()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_update_dirty_quadrants", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Clear all cells.
*/
func (o *TileMap) Clear() {
	log.Println("Calling TileMap.Clear()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "clear", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Return the tile index of the referenced cell.
*/
func (o *TileMap) GetCell(x int64, y int64) int64 {
	log.Println("Calling TileMap.GetCell()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(x)
	goArguments[1] = reflect.ValueOf(y)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_cell", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the cell size.
*/
func (o *TileMap) GetCellSize() *Vector2 {
	log.Println("Calling TileMap.GetCellSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_cell_size", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the tile index of the cell referenced by a Vector2.
*/
func (o *TileMap) GetCellv(position *Vector2) int64 {
	log.Println("Calling TileMap.GetCellv()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(position)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_cellv", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return true if tiles are to be centered in x coordinate (by default this is false and they are drawn from upper left cell corner).
*/
func (o *TileMap) GetCenterX() bool {
	log.Println("Calling TileMap.GetCenterX()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_center_x", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return true if tiles are to be centered in y coordinate (by default this is false and they are drawn from upper left cell corner).
*/
func (o *TileMap) GetCenterY() bool {
	log.Println("Calling TileMap.GetCenterY()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_center_y", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the collision bounce parameter.
*/
func (o *TileMap) GetCollisionBounce() float64 {
	log.Println("Calling TileMap.GetCollisionBounce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_collision_bounce", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the collision friction parameter.
*/
func (o *TileMap) GetCollisionFriction() float64 {
	log.Println("Calling TileMap.GetCollisionFriction()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_collision_friction", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the collision layer.
*/
func (o *TileMap) GetCollisionLayer() int64 {
	log.Println("Calling TileMap.GetCollisionLayer()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_collision_layer", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *TileMap) GetCollisionLayerBit(bit int64) bool {
	log.Println("Calling TileMap.GetCollisionLayerBit()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(bit)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_collision_layer_bit", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the collision mask.
*/
func (o *TileMap) GetCollisionMask() int64 {
	log.Println("Calling TileMap.GetCollisionMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_collision_mask", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *TileMap) GetCollisionMaskBit(bit int64) bool {
	log.Println("Calling TileMap.GetCollisionMaskBit()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(bit)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_collision_mask_bit", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return whether the tilemap handles collisions as a kinematic body.
*/
func (o *TileMap) GetCollisionUseKinematic() bool {
	log.Println("Calling TileMap.GetCollisionUseKinematic()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_collision_use_kinematic", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the custom transform matrix.
*/
func (o *TileMap) GetCustomTransform() *Transform2D {
	log.Println("Calling TileMap.GetCustomTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_custom_transform", goArguments, "*Transform2D")

	returnValue := goRet.Interface().(*Transform2D)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the current half offset configuration.
*/
func (o *TileMap) GetHalfOffset() int64 {
	log.Println("Calling TileMap.GetHalfOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_half_offset", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the orientation mode.
*/
func (o *TileMap) GetMode() int64 {
	log.Println("Calling TileMap.GetMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_mode", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *TileMap) GetOccluderLightMask() int64 {
	log.Println("Calling TileMap.GetOccluderLightMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_occluder_light_mask", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the quadrant size.
*/
func (o *TileMap) GetQuadrantSize() int64 {
	log.Println("Calling TileMap.GetQuadrantSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_quadrant_size", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the tile origin configuration.
*/
func (o *TileMap) GetTileOrigin() int64 {
	log.Println("Calling TileMap.GetTileOrigin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_tile_origin", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the current tileset.
*/
func (o *TileMap) GetTileset() *TileSet {
	log.Println("Calling TileMap.GetTileset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_tileset", goArguments, "*TileSet")

	returnValue := goRet.Interface().(*TileSet)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return an array of all cells containing a tile from the tileset (i.e. a tile index different from -1).
*/
func (o *TileMap) GetUsedCells() *Array {
	log.Println("Calling TileMap.GetUsedCells()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_used_cells", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *TileMap) GetUsedCellsById(id int64) *Array {
	log.Println("Calling TileMap.GetUsedCellsById()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(id)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_used_cells_by_id", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *TileMap) GetUsedRect() *Rect2 {
	log.Println("Calling TileMap.GetUsedRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_used_rect", goArguments, "*Rect2")

	returnValue := goRet.Interface().(*Rect2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return whether the referenced cell is transposed, i.e. the X and Y axes are swapped (mirroring with regard to the (1,1) vector).
*/
func (o *TileMap) IsCellTransposed(x int64, y int64) bool {
	log.Println("Calling TileMap.IsCellTransposed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(x)
	goArguments[1] = reflect.ValueOf(y)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_cell_transposed", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return whether the referenced cell is flipped over the X axis.
*/
func (o *TileMap) IsCellXFlipped(x int64, y int64) bool {
	log.Println("Calling TileMap.IsCellXFlipped()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(x)
	goArguments[1] = reflect.ValueOf(y)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_cell_x_flipped", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return whether the referenced cell is flipped over the Y axis.
*/
func (o *TileMap) IsCellYFlipped(x int64, y int64) bool {
	log.Println("Calling TileMap.IsCellYFlipped()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(x)
	goArguments[1] = reflect.ValueOf(y)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_cell_y_flipped", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the Y sort mode.
*/
func (o *TileMap) IsYSortModeEnabled() bool {
	log.Println("Calling TileMap.IsYSortModeEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_y_sort_mode_enabled", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the absolute world position corresponding to the tilemap (grid-based) coordinates given as an argument. Optionally, the tilemap's potential half offset can be ignored.
*/
func (o *TileMap) MapToWorld(mapPosition *Vector2, ignoreHalfOfs bool) *Vector2 {
	log.Println("Calling TileMap.MapToWorld()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(mapPosition)
	goArguments[1] = reflect.ValueOf(ignoreHalfOfs)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "map_to_world", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Set the tile index for the cell referenced by its grid-based X and Y coordinates. A tile index of -1 clears the cell. Optionally, the tile can also be flipped over the X and Y coordinates or transposed.
*/
func (o *TileMap) SetCell(x int64, y int64, tile int64, flipX bool, flipY bool, transpose bool) {
	log.Println("Calling TileMap.SetCell()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 6, 6)
	goArguments[0] = reflect.ValueOf(x)
	goArguments[1] = reflect.ValueOf(y)
	goArguments[2] = reflect.ValueOf(tile)
	goArguments[3] = reflect.ValueOf(flipX)
	goArguments[4] = reflect.ValueOf(flipY)
	goArguments[5] = reflect.ValueOf(transpose)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_cell", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the cell size.
*/
func (o *TileMap) SetCellSize(size *Vector2) {
	log.Println("Calling TileMap.SetCellSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(size)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_cell_size", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the tile index for the cell referenced by a Vector2 of grid-based coordinates. A tile index of -1 clears the cell. Optionally, the tile can also be flipped over the X and Y axes or transposed.
*/
func (o *TileMap) SetCellv(position *Vector2, tile int64, flipX bool, flipY bool, transpose bool) {
	log.Println("Calling TileMap.SetCellv()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 5, 5)
	goArguments[0] = reflect.ValueOf(position)
	goArguments[1] = reflect.ValueOf(tile)
	goArguments[2] = reflect.ValueOf(flipX)
	goArguments[3] = reflect.ValueOf(flipY)
	goArguments[4] = reflect.ValueOf(transpose)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_cellv", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set tiles to be centered in x coordinate. (by default this is false and they are drawn from upper left cell corner).
*/
func (o *TileMap) SetCenterX(enable bool) {
	log.Println("Calling TileMap.SetCenterX()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_center_x", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set tiles to be centered in y coordinate. (by default this is false and they are drawn from upper left cell corner).
*/
func (o *TileMap) SetCenterY(enable bool) {
	log.Println("Calling TileMap.SetCenterY()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_center_y", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the collision bounce parameter. Allowable values range from 0 to 1.
*/
func (o *TileMap) SetCollisionBounce(value float64) {
	log.Println("Calling TileMap.SetCollisionBounce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_collision_bounce", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the collision friction parameter. Allowable values range from 0 to 1.
*/
func (o *TileMap) SetCollisionFriction(value float64) {
	log.Println("Calling TileMap.SetCollisionFriction()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_collision_friction", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the collision layer. Layers are referenced by binary indexes, so allowable values to describe the 20 available layers range from 0 to 2^20-1.
*/
func (o *TileMap) SetCollisionLayer(layer int64) {
	log.Println("Calling TileMap.SetCollisionLayer()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(layer)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_collision_layer", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *TileMap) SetCollisionLayerBit(bit int64, value bool) {
	log.Println("Calling TileMap.SetCollisionLayerBit()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(bit)
	goArguments[1] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_collision_layer_bit", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the collision masks. Masks are referenced by binary indexes, so allowable values to describe the 20 available masks range from 0 to 2^20-1.
*/
func (o *TileMap) SetCollisionMask(mask int64) {
	log.Println("Calling TileMap.SetCollisionMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mask)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_collision_mask", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *TileMap) SetCollisionMaskBit(bit int64, value bool) {
	log.Println("Calling TileMap.SetCollisionMaskBit()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(bit)
	goArguments[1] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_collision_mask_bit", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the tilemap to handle collisions as a kinematic body (enabled) or a static body (disabled).
*/
func (o *TileMap) SetCollisionUseKinematic(useKinematic bool) {
	log.Println("Calling TileMap.SetCollisionUseKinematic()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(useKinematic)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_collision_use_kinematic", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set custom transform matrix, to use in combination with the custom orientation mode.
*/
func (o *TileMap) SetCustomTransform(customTransform *Transform2D) {
	log.Println("Calling TileMap.SetCustomTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(customTransform)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_custom_transform", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set a half offset on the X coordinate, Y coordinate, or none (use HALF_OFFSET_* constants as argument). Half offset sets every other tile off by a half tile size in the specified direction.
*/
func (o *TileMap) SetHalfOffset(halfOffset int64) {
	log.Println("Calling TileMap.SetHalfOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(halfOffset)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_half_offset", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the orientation mode as square, isometric or custom (use MODE_* constants as argument).
*/
func (o *TileMap) SetMode(mode int64) {
	log.Println("Calling TileMap.SetMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *TileMap) SetOccluderLightMask(mask int64) {
	log.Println("Calling TileMap.SetOccluderLightMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mask)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_occluder_light_mask", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the quadrant size, this optimizes drawing by batching chunks of map at draw/cull time. Allowed values are integers ranging from 1 to 128.
*/
func (o *TileMap) SetQuadrantSize(size int64) {
	log.Println("Calling TileMap.SetQuadrantSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(size)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_quadrant_size", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the tile origin to the tile center or its top-left corner (use TILE_ORIGIN_* constants as argument).
*/
func (o *TileMap) SetTileOrigin(origin int64) {
	log.Println("Calling TileMap.SetTileOrigin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(origin)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_tile_origin", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the current tileset.
*/
func (o *TileMap) SetTileset(tileset *TileSet) {
	log.Println("Calling TileMap.SetTileset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(tileset)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_tileset", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the Y sort mode. Enabled Y sort mode means that children of the tilemap will be drawn in the order defined by their Y coordinate. A tile with a higher Y coordinate will therefore be drawn later, potentially covering up the tile(s) above it if its sprite is higher than its cell size.
*/
func (o *TileMap) SetYSortMode(enable bool) {
	log.Println("Calling TileMap.SetYSortMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_y_sort_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Return the tilemap (grid-based) coordinates corresponding to the absolute world position given as an argument.
*/
func (o *TileMap) WorldToMap(worldPosition *Vector2) *Vector2 {
	log.Println("Calling TileMap.WorldToMap()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(worldPosition)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "world_to_map", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   TileMapImplementer is an interface for TileMap objects.
*/
type TileMapImplementer interface {
	class.Class
}