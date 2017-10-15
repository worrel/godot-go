//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package particlesmaterial

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

	"github.com/shadowapex/godot-go/godot/classes/material"
)

/*
   ParticlesMaterial defines particle properties and behavior. It is used in the [code]process_material[/code] of [Particles] and [Particles2D] emitter nodes. Some of this material's properties are applied to each particle when emitted, while others can have a [CurveTexture] applied to vary values over the lifetime of the particle.
*/
type ParticlesMaterial struct {
	material.Material
}

func (o *ParticlesMaterial) baseClass() string {
	return "ParticlesMaterial"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *ParticlesMaterial) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *ParticlesMaterial) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *ParticlesMaterial) GetColor() *Color {
	log.Println("Calling ParticlesMaterial.GetColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_color", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) GetColorRamp() *Texture {
	log.Println("Calling ParticlesMaterial.GetColorRamp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_color_ramp", goArguments, "*Texture")

	returnValue := goRet.Interface().(*Texture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) GetEmissionBoxExtents() *Vector3 {
	log.Println("Calling ParticlesMaterial.GetEmissionBoxExtents()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_emission_box_extents", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) GetEmissionColorTexture() *Texture {
	log.Println("Calling ParticlesMaterial.GetEmissionColorTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_emission_color_texture", goArguments, "*Texture")

	returnValue := goRet.Interface().(*Texture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) GetEmissionNormalTexture() *Texture {
	log.Println("Calling ParticlesMaterial.GetEmissionNormalTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_emission_normal_texture", goArguments, "*Texture")

	returnValue := goRet.Interface().(*Texture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) GetEmissionPointCount() int64 {
	log.Println("Calling ParticlesMaterial.GetEmissionPointCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_emission_point_count", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) GetEmissionPointTexture() *Texture {
	log.Println("Calling ParticlesMaterial.GetEmissionPointTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_emission_point_texture", goArguments, "*Texture")

	returnValue := goRet.Interface().(*Texture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) GetEmissionShape() int64 {
	log.Println("Calling ParticlesMaterial.GetEmissionShape()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_emission_shape", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) GetEmissionSphereRadius() float64 {
	log.Println("Calling ParticlesMaterial.GetEmissionSphereRadius()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_emission_sphere_radius", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) GetFlag(flag int64) bool {
	log.Println("Calling ParticlesMaterial.GetFlag()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(flag)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_flag", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) GetFlatness() float64 {
	log.Println("Calling ParticlesMaterial.GetFlatness()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_flatness", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) GetGravity() *Vector3 {
	log.Println("Calling ParticlesMaterial.GetGravity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_gravity", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) GetParam(param int64) float64 {
	log.Println("Calling ParticlesMaterial.GetParam()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(param)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_param", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) GetParamRandomness(param int64) float64 {
	log.Println("Calling ParticlesMaterial.GetParamRandomness()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(param)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_param_randomness", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) GetParamTexture(param int64) *Texture {
	log.Println("Calling ParticlesMaterial.GetParamTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(param)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_param_texture", goArguments, "*Texture")

	returnValue := goRet.Interface().(*Texture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) GetSpread() float64 {
	log.Println("Calling ParticlesMaterial.GetSpread()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_spread", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) GetTrailColorModifier() *GradientTexture {
	log.Println("Calling ParticlesMaterial.GetTrailColorModifier()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_trail_color_modifier", goArguments, "*GradientTexture")

	returnValue := goRet.Interface().(*GradientTexture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) GetTrailDivisor() int64 {
	log.Println("Calling ParticlesMaterial.GetTrailDivisor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_trail_divisor", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) GetTrailSizeModifier() *CurveTexture {
	log.Println("Calling ParticlesMaterial.GetTrailSizeModifier()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_trail_size_modifier", goArguments, "*CurveTexture")

	returnValue := goRet.Interface().(*CurveTexture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ParticlesMaterial) SetColor(color *Color) {
	log.Println("Calling ParticlesMaterial.SetColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ParticlesMaterial) SetColorRamp(ramp *Texture) {
	log.Println("Calling ParticlesMaterial.SetColorRamp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(ramp)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_color_ramp", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ParticlesMaterial) SetEmissionBoxExtents(extents *Vector3) {
	log.Println("Calling ParticlesMaterial.SetEmissionBoxExtents()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(extents)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_emission_box_extents", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ParticlesMaterial) SetEmissionColorTexture(texture *Texture) {
	log.Println("Calling ParticlesMaterial.SetEmissionColorTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(texture)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_emission_color_texture", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ParticlesMaterial) SetEmissionNormalTexture(texture *Texture) {
	log.Println("Calling ParticlesMaterial.SetEmissionNormalTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(texture)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_emission_normal_texture", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ParticlesMaterial) SetEmissionPointCount(pointCount int64) {
	log.Println("Calling ParticlesMaterial.SetEmissionPointCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(pointCount)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_emission_point_count", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ParticlesMaterial) SetEmissionPointTexture(texture *Texture) {
	log.Println("Calling ParticlesMaterial.SetEmissionPointTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(texture)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_emission_point_texture", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ParticlesMaterial) SetEmissionShape(shape int64) {
	log.Println("Calling ParticlesMaterial.SetEmissionShape()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(shape)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_emission_shape", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ParticlesMaterial) SetEmissionSphereRadius(radius float64) {
	log.Println("Calling ParticlesMaterial.SetEmissionSphereRadius()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(radius)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_emission_sphere_radius", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ParticlesMaterial) SetFlag(flag int64, enable bool) {
	log.Println("Calling ParticlesMaterial.SetFlag()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(flag)
	goArguments[1] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_flag", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ParticlesMaterial) SetFlatness(amount float64) {
	log.Println("Calling ParticlesMaterial.SetFlatness()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(amount)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_flatness", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ParticlesMaterial) SetGravity(accelVec *Vector3) {
	log.Println("Calling ParticlesMaterial.SetGravity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(accelVec)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_gravity", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ParticlesMaterial) SetParam(param int64, value float64) {
	log.Println("Calling ParticlesMaterial.SetParam()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(param)
	goArguments[1] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_param", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ParticlesMaterial) SetParamRandomness(param int64, randomness float64) {
	log.Println("Calling ParticlesMaterial.SetParamRandomness()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(param)
	goArguments[1] = reflect.ValueOf(randomness)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_param_randomness", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ParticlesMaterial) SetParamTexture(param int64, texture *Texture) {
	log.Println("Calling ParticlesMaterial.SetParamTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(param)
	goArguments[1] = reflect.ValueOf(texture)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_param_texture", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ParticlesMaterial) SetSpread(degrees float64) {
	log.Println("Calling ParticlesMaterial.SetSpread()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(degrees)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_spread", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ParticlesMaterial) SetTrailColorModifier(texture *GradientTexture) {
	log.Println("Calling ParticlesMaterial.SetTrailColorModifier()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(texture)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_trail_color_modifier", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ParticlesMaterial) SetTrailDivisor(divisor int64) {
	log.Println("Calling ParticlesMaterial.SetTrailDivisor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(divisor)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_trail_divisor", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ParticlesMaterial) SetTrailSizeModifier(texture *CurveTexture) {
	log.Println("Calling ParticlesMaterial.SetTrailSizeModifier()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(texture)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_trail_size_modifier", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   ParticlesMaterialImplementer is an interface for ParticlesMaterial objects.
*/
type ParticlesMaterialImplementer interface {
	class.Class
}
