//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package quadmesh

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

	"github.com/shadowapex/godot-go/godot/classes/primitivemesh"
)

/*
   Class representing a square mesh with size (2,2,0). Consider using a [PlaneMesh] if you require a differently sized plane.
*/
type QuadMesh struct {
	primitivemesh.PrimitiveMesh
}

func (o *QuadMesh) baseClass() string {
	return "QuadMesh"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *QuadMesh) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *QuadMesh) getOwner() *C.godot_object {
	return o.owner
}

/*
   QuadMeshImplementer is an interface for QuadMesh objects.
*/
type QuadMeshImplementer interface {
	class.Class
}
