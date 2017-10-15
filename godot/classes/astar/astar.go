//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package astar

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

	"github.com/shadowapex/godot-go/godot/classes/reference"
)

/*
   A* (A star) is a computer algorithm that is widely used in pathfinding and graph traversal, the process of plotting an efficiently directed path between multiple points. It enjoys widespread use due to its performance and accuracy. Godot's A* implementation make use of vectors as points. You must add points manually with [method AStar.add_point] and create segments manually with [method AStar.connect_points]. So you can test if there is a path between two points with the [method AStar.are_points_connected] function, get the list of existing ids in the found path with [method AStar.get_id_path], or the points list with [method AStar.get_point_path].
*/
type AStar struct {
	reference.Reference
}

func (o *AStar) baseClass() string {
	return "AStar"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *AStar) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *AStar) getOwner() *C.godot_object {
	return o.owner
}

/*
   Called when computing the cost between two connected points.
*/
func (o *AStar) X_ComputeCost(fromId int64, toId int64) {
	log.Println("Calling AStar.X_ComputeCost()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(fromId)
	goArguments[1] = reflect.ValueOf(toId)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_compute_cost", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Called when estimating the cost between a point and the path's ending point.
*/
func (o *AStar) X_EstimateCost(fromId int64, toId int64) {
	log.Println("Calling AStar.X_EstimateCost()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(fromId)
	goArguments[1] = reflect.ValueOf(toId)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_estimate_cost", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Adds a new point at the given position with the given identifier. The algorithm prefers points with lower [code]weight_scale[/code] to form a path. The [code]id[/code] must be 0 or larger, and the [code]weight_scale[/code] must be 1 or larger. [codeblock] var as = AStar.new() as.add_point(1, Vector3(1,0,0), 4) # Adds the point (1,0,0) with weight_scale=4 and id=1 [/codeblock]
*/
func (o *AStar) AddPoint(id int64, position *Vector3, weightScale float64) {
	log.Println("Calling AStar.AddPoint()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(id)
	goArguments[1] = reflect.ValueOf(position)
	goArguments[2] = reflect.ValueOf(weightScale)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_point", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Returns whether there is a connection/segment between the given points.
*/
func (o *AStar) ArePointsConnected(id int64, toId int64) bool {
	log.Println("Calling AStar.ArePointsConnected()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(id)
	goArguments[1] = reflect.ValueOf(toId)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "are_points_connected", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Clears all the points and segments.
*/
func (o *AStar) Clear() {
	log.Println("Calling AStar.Clear()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "clear", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Creates a segment between the given points. [codeblock] var as = AStar.new() as.add_point(1, Vector3(1,1,0)) as.add_point(2, Vector3(0,5,0)) as.connect_points(1, 2, false) # If bidirectional=false it's only possible to go from point 1 to point 2 # and not from point 2 to point 1. [/codeblock]
*/
func (o *AStar) ConnectPoints(id int64, toId int64, bidirectional bool) {
	log.Println("Calling AStar.ConnectPoints()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(id)
	goArguments[1] = reflect.ValueOf(toId)
	goArguments[2] = reflect.ValueOf(bidirectional)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "connect_points", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Deletes the segment between the given points.
*/
func (o *AStar) DisconnectPoints(id int64, toId int64) {
	log.Println("Calling AStar.DisconnectPoints()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(id)
	goArguments[1] = reflect.ValueOf(toId)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "disconnect_points", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Returns an id with no point associated to it.
*/
func (o *AStar) GetAvailablePointId() int64 {
	log.Println("Calling AStar.GetAvailablePointId()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_available_point_id", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the id of the closest point to [code]to_position[/code]. Returns -1 if there are no points in the points pool.
*/
func (o *AStar) GetClosestPoint(toPosition *Vector3) int64 {
	log.Println("Calling AStar.GetClosestPoint()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(toPosition)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_closest_point", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the closest position to [code]to_position[/code] that resides inside a segment between two connected points. [codeblock] var as = AStar.new() as.add_point(1, Vector3(0,0,0)) as.add_point(2, Vector3(0,5,0)) as.connect_points(1, 2) var res = as.get_closest_position_in_segment(Vector3(3,3,0)) # returns (0, 3, 0) [/codeblock] The result is in the segment that goes from [code]y=0[/code] to [code]y=5[/code]. It's the closest position in the segment to the given point.
*/
func (o *AStar) GetClosestPositionInSegment(toPosition *Vector3) *Vector3 {
	log.Println("Calling AStar.GetClosestPositionInSegment()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(toPosition)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_closest_position_in_segment", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns an array with the ids of the points that form the path found by AStar between the given points. The array is ordered from the starting point to the ending point of the path. [codeblock] var as = AStar.new() as.add_point(1, Vector3(0,0,0)) as.add_point(2, Vector3(0,1,0), 1) # default weight is 1 as.add_point(3, Vector3(1,1,0)) as.add_point(4, Vector3(2,0,0)) as.connect_points(1, 2, false) as.connect_points(2, 3, false) as.connect_points(4, 3, false) as.connect_points(1, 4, false) as.connect_points(5, 4, false) var res = as.get_id_path(1, 3) # returns [1, 2, 3] [/codeblock] If you change the 2nd point's weight to 3, then the result will be [code][1, 4, 3][/code] instead, because now even though the distance is longer, it's "easier" to get through point 4 than through point 2.
*/
func (o *AStar) GetIdPath(fromId int64, toId int64) *PoolIntArray {
	log.Println("Calling AStar.GetIdPath()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(fromId)
	goArguments[1] = reflect.ValueOf(toId)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_id_path", goArguments, "*PoolIntArray")

	returnValue := goRet.Interface().(*PoolIntArray)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns an array with the points that are in the path found by AStar between the given points. The array is ordered from the starting point to the ending point of the path.
*/
func (o *AStar) GetPointPath(fromId int64, toId int64) *PoolVector3Array {
	log.Println("Calling AStar.GetPointPath()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(fromId)
	goArguments[1] = reflect.ValueOf(toId)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_point_path", goArguments, "*PoolVector3Array")

	returnValue := goRet.Interface().(*PoolVector3Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the position of the point associated with the given id.
*/
func (o *AStar) GetPointPosition(id int64) *Vector3 {
	log.Println("Calling AStar.GetPointPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(id)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_point_position", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the weight scale of the point associated with the given id.
*/
func (o *AStar) GetPointWeightScale(id int64) float64 {
	log.Println("Calling AStar.GetPointWeightScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(id)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_point_weight_scale", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *AStar) GetPoints() *Array {
	log.Println("Calling AStar.GetPoints()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_points", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns whether a point associated with the given id exists.
*/
func (o *AStar) HasPoint(id int64) bool {
	log.Println("Calling AStar.HasPoint()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(id)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "has_point", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Removes the point associated with the given id from the points pool.
*/
func (o *AStar) RemovePoint(id int64) {
	log.Println("Calling AStar.RemovePoint()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(id)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "remove_point", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   AStarImplementer is an interface for AStar objects.
*/
type AStarImplementer interface {
	class.Class
}
