//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package proceduralsky

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

	"github.com/shadowapex/godot-go/godot/classes/sky"
)

/*

 */
type ProceduralSky struct {
	sky.Sky
}

func (o *ProceduralSky) baseClass() string {
	return "ProceduralSky"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *ProceduralSky) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *ProceduralSky) getOwner() *C.godot_object {
	return o.owner
}

/*
   Undocumented
*/
func (o *ProceduralSky) X_ThreadDone(image *Image) {
	log.Println("Calling ProceduralSky.X_ThreadDone()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(image)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_thread_done", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ProceduralSky) X_UpdateSky() {
	log.Println("Calling ProceduralSky.X_UpdateSky()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_update_sky", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ProceduralSky) GetGroundBottomColor() *Color {
	log.Println("Calling ProceduralSky.GetGroundBottomColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_ground_bottom_color", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ProceduralSky) GetGroundCurve() float64 {
	log.Println("Calling ProceduralSky.GetGroundCurve()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_ground_curve", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ProceduralSky) GetGroundEnergy() float64 {
	log.Println("Calling ProceduralSky.GetGroundEnergy()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_ground_energy", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ProceduralSky) GetGroundHorizonColor() *Color {
	log.Println("Calling ProceduralSky.GetGroundHorizonColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_ground_horizon_color", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ProceduralSky) GetSkyCurve() float64 {
	log.Println("Calling ProceduralSky.GetSkyCurve()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_sky_curve", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ProceduralSky) GetSkyEnergy() float64 {
	log.Println("Calling ProceduralSky.GetSkyEnergy()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_sky_energy", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ProceduralSky) GetSkyHorizonColor() *Color {
	log.Println("Calling ProceduralSky.GetSkyHorizonColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_sky_horizon_color", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ProceduralSky) GetSkyTopColor() *Color {
	log.Println("Calling ProceduralSky.GetSkyTopColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_sky_top_color", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ProceduralSky) GetSunAngleMax() float64 {
	log.Println("Calling ProceduralSky.GetSunAngleMax()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_sun_angle_max", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ProceduralSky) GetSunAngleMin() float64 {
	log.Println("Calling ProceduralSky.GetSunAngleMin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_sun_angle_min", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ProceduralSky) GetSunColor() *Color {
	log.Println("Calling ProceduralSky.GetSunColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_sun_color", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ProceduralSky) GetSunCurve() float64 {
	log.Println("Calling ProceduralSky.GetSunCurve()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_sun_curve", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ProceduralSky) GetSunEnergy() float64 {
	log.Println("Calling ProceduralSky.GetSunEnergy()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_sun_energy", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ProceduralSky) GetSunLatitude() float64 {
	log.Println("Calling ProceduralSky.GetSunLatitude()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_sun_latitude", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ProceduralSky) GetSunLongitude() float64 {
	log.Println("Calling ProceduralSky.GetSunLongitude()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_sun_longitude", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ProceduralSky) GetTextureSize() int64 {
	log.Println("Calling ProceduralSky.GetTextureSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_texture_size", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ProceduralSky) SetGroundBottomColor(color *Color) {
	log.Println("Calling ProceduralSky.SetGroundBottomColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_ground_bottom_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ProceduralSky) SetGroundCurve(curve float64) {
	log.Println("Calling ProceduralSky.SetGroundCurve()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(curve)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_ground_curve", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ProceduralSky) SetGroundEnergy(energy float64) {
	log.Println("Calling ProceduralSky.SetGroundEnergy()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(energy)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_ground_energy", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ProceduralSky) SetGroundHorizonColor(color *Color) {
	log.Println("Calling ProceduralSky.SetGroundHorizonColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_ground_horizon_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ProceduralSky) SetSkyCurve(curve float64) {
	log.Println("Calling ProceduralSky.SetSkyCurve()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(curve)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_sky_curve", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ProceduralSky) SetSkyEnergy(energy float64) {
	log.Println("Calling ProceduralSky.SetSkyEnergy()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(energy)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_sky_energy", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ProceduralSky) SetSkyHorizonColor(color *Color) {
	log.Println("Calling ProceduralSky.SetSkyHorizonColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_sky_horizon_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ProceduralSky) SetSkyTopColor(color *Color) {
	log.Println("Calling ProceduralSky.SetSkyTopColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_sky_top_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ProceduralSky) SetSunAngleMax(degrees float64) {
	log.Println("Calling ProceduralSky.SetSunAngleMax()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(degrees)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_sun_angle_max", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ProceduralSky) SetSunAngleMin(degrees float64) {
	log.Println("Calling ProceduralSky.SetSunAngleMin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(degrees)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_sun_angle_min", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ProceduralSky) SetSunColor(color *Color) {
	log.Println("Calling ProceduralSky.SetSunColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_sun_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ProceduralSky) SetSunCurve(curve float64) {
	log.Println("Calling ProceduralSky.SetSunCurve()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(curve)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_sun_curve", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ProceduralSky) SetSunEnergy(energy float64) {
	log.Println("Calling ProceduralSky.SetSunEnergy()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(energy)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_sun_energy", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ProceduralSky) SetSunLatitude(degrees float64) {
	log.Println("Calling ProceduralSky.SetSunLatitude()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(degrees)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_sun_latitude", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ProceduralSky) SetSunLongitude(degrees float64) {
	log.Println("Calling ProceduralSky.SetSunLongitude()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(degrees)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_sun_longitude", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ProceduralSky) SetTextureSize(size int64) {
	log.Println("Calling ProceduralSky.SetTextureSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(size)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_texture_size", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   ProceduralSkyImplementer is an interface for ProceduralSky objects.
*/
type ProceduralSkyImplementer interface {
	class.Class
}
