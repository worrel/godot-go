//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package kinematicbody2d

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

	"github.com/shadowapex/godot-go/godot/classes/physicsbody2d"
)

/*
   Kinematic bodies are special types of bodies that are meant to be user-controlled. They are not affected by physics at all (to other types of bodies, such a character or a rigid body, these are the same as a static body). They have however, two main uses: Simulated Motion: When these bodies are moved manually, either from code or from an AnimationPlayer (with process mode set to fixed), the physics will automatically compute an estimate of their linear and angular velocity. This makes them very useful for moving platforms or other AnimationPlayer-controlled objects (like a door, a bridge that opens, etc). Kinematic Characters: KinematicBody2D also has an API for moving objects (the [method move_and_collide] and [method move_and_slide] methods) while performing collision tests. This makes them really useful to implement characters that collide against a world, but that don't require advanced physics.
*/
type KinematicBody2D struct {
	physicsbody2d.PhysicsBody2D
}

func (o *KinematicBody2D) baseClass() string {
	return "KinematicBody2D"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *KinematicBody2D) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *KinematicBody2D) getOwner() *C.godot_object {
	return o.owner
}

/*
   Returns the velocity of the floor. Only updates when calling [method move_and_slide].
*/
func (o *KinematicBody2D) GetFloorVelocity() *Vector2 {
	log.Println("Calling KinematicBody2D.GetFloorVelocity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_floor_velocity", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *KinematicBody2D) GetSafeMargin() float64 {
	log.Println("Calling KinematicBody2D.GetSafeMargin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_safe_margin", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns a [KinematicCollision2D], which contains information about a collision that occured during the last [method move_and_slide] call. Since the body can collide several times in a single call to [method move_and_slide], you must specify the index of the collision in the range 0 to ([method get_slide_count]()-1).
*/
func (o *KinematicBody2D) GetSlideCollision(slideIdx int64) *KinematicCollision2D {
	log.Println("Calling KinematicBody2D.GetSlideCollision()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(slideIdx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_slide_collision", goArguments, "*KinematicCollision2D")

	returnValue := goRet.Interface().(*KinematicCollision2D)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the number of times the body collided and changed direction during the last call to [method move_and_slide].
*/
func (o *KinematicBody2D) GetSlideCount() int64 {
	log.Println("Calling KinematicBody2D.GetSlideCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_slide_count", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if the body is on the ceiling. Only updates when calling [method move_and_slide].
*/
func (o *KinematicBody2D) IsOnCeiling() bool {
	log.Println("Calling KinematicBody2D.IsOnCeiling()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_on_ceiling", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if the body is on the floor. Only updates when calling [method move_and_slide].
*/
func (o *KinematicBody2D) IsOnFloor() bool {
	log.Println("Calling KinematicBody2D.IsOnFloor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_on_floor", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if the body is on a wall. Only updates when calling [method move_and_slide].
*/
func (o *KinematicBody2D) IsOnWall() bool {
	log.Println("Calling KinematicBody2D.IsOnWall()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_on_wall", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Moves the body along the vector [code]rel_vec[/code]. The body will stop if it collides. Returns a [KinematicCollision2D], which contains information about the collision.
*/
func (o *KinematicBody2D) MoveAndCollide(relVec *Vector2) *KinematicCollision2D {
	log.Println("Calling KinematicBody2D.MoveAndCollide()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(relVec)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "move_and_collide", goArguments, "*KinematicCollision2D")

	returnValue := goRet.Interface().(*KinematicCollision2D)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Moves the body along a vector. If the body collides with another, it will slide along the other body rather than stop immediately. If the other body is a [KinematicBody2D] or [RigidBody2D], it will also be affected by the motion of the other body. You can use this to make moving or rotating platforms, or to make nodes push other nodes. [code]linear_velocity[/code] is a value in pixels per second. Unlike in for example [method move_and_collide], you should [i]not[/i] multiply it with [code]delta[/code] — this is done by the method. [code]floor_normal[/code] is the up direction, used to determine what is a wall and what is a floor or a ceiling. If set to the default value of [code]Vector2(0, 0)[/code], everything is considered a wall. This is useful for topdown games. If the body is standing on a slope and the horizontal speed (relative to the floor's speed) goes below [code]slope_stop_min_velocity[/code], the body will stop completely. This prevents the body from sliding down slopes when you include gravity in [code]linear_velocity[/code]. When set to lower values, the body will not be able to stand still on steep slopes. If the body collides, it will change direction a maximum of [code]max_bounces[/code] times before it stops. [code]floor_max_angle[/code] is the maximum angle (in radians) where a slope is still considered a floor (or a ceiling), rather than a wall. The default value equals 45 degrees. Returns the movement that remained when the body stopped. To get more detailed information about collisions that occured, use [method get_slide_collision].
*/
func (o *KinematicBody2D) MoveAndSlide(linearVelocity *Vector2, floorNormal *Vector2, slopeStopMinVelocity float64, maxBounces int64, floorMaxAngle float64) *Vector2 {
	log.Println("Calling KinematicBody2D.MoveAndSlide()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 5, 5)
	goArguments[0] = reflect.ValueOf(linearVelocity)
	goArguments[1] = reflect.ValueOf(floorNormal)
	goArguments[2] = reflect.ValueOf(slopeStopMinVelocity)
	goArguments[3] = reflect.ValueOf(maxBounces)
	goArguments[4] = reflect.ValueOf(floorMaxAngle)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "move_and_slide", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *KinematicBody2D) SetSafeMargin(pixels float64) {
	log.Println("Calling KinematicBody2D.SetSafeMargin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(pixels)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_safe_margin", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Checks for collisions without moving the body. Virtually sets the node's position, scale and rotation to that of the given [Transform2D], then tries to move the body along the vector [code]rel_vec[/code]. Returns [code]true[/code] if a collision would occur.
*/
func (o *KinematicBody2D) TestMove(from *Transform2D, relVec *Vector2) bool {
	log.Println("Calling KinematicBody2D.TestMove()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(from)
	goArguments[1] = reflect.ValueOf(relVec)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "test_move", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   KinematicBody2DImplementer is an interface for KinematicBody2D objects.
*/
type KinematicBody2DImplementer interface {
	class.Class
}
