// Package class contains binding methods for godot. This generally should
// only by imported by godot-go packages only.
package class

/*
#include <stdio.h>
#include <stdlib.h>
#include <gdnative/gdnative.h>
#include <nativescript/godot_nativescript.h>

void **build_array(int length);
void **build_array(int length) {
	void *ptr;
	void **arr = malloc(sizeof(void *) * length);
	for (int i = 0; i < length; i++) {
	    arr[i] = ptr;
	}

	return arr;
}

void add_element(void**, void*, int);
void add_element(void **array, void *element, int index) {
	printf("CGO: Array %p %p %p %p %p\n", &array, array, &array[index], *array, array[index]);
    array[index] = element;
	printf("CGO: Index %i %p\n", index, element);
	printf("CGO: Array %p %p %p %p %p\n", &array, array, &array[index], *array, array[index]);
}
*/
import "C"
import (
	"log"
	"reflect"
	"unsafe"
)

// Class is an interface for any objects that can have Godot
// inheritance.
type Class interface {
	baseClass() string
	setOwner(object *C.godot_object)
	getOwner() *C.godot_object
}

// CallParentMethod will call the given object's method with the given method name.
// This should generally only be called by godot-go functions.
func CallParentMethod(o Class, baseClass, methodName string, args []reflect.Value, returns string) reflect.Value {
	log.Println("Calling parent method!")

	// Convert the base class and method names to C strings.
	log.Println("  Using base class: ", baseClass)
	classCString := C.CString(baseClass)
	log.Println("  Using method name: ", methodName)
	methodCString := C.CString(methodName)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Get the Godot method bind pointer so we can pass it to godot_method_bind_ptrcall.
	var methodBind *C.godot_method_bind
	methodBind = C.godot_method_bind_get_method(classCString, methodCString)
	log.Println("  Using method bind pointer: ", methodBind)

	// Loop through the given arguments and see what type they are. When we know what
	// type it is, we need to convert them to the correct godot objects.
	// TODO: Probably pull this out into its own function?
	variantArgs := []unsafe.Pointer{}
	for _, arg := range args {
		log.Println("  Argument type: ", arg.Type().String())

		// Look up our conversion function in our map of conversion functions
		// based on the Go type. This is essentially a more optimal case/switch
		// statement on the type of Go object, so we can know how to convert it
		// to a Godot object.
		if convert, ok := goToGodotConversionMap[arg.Type().String()]; ok {
			argValue := convert(arg.Interface())
			variantArgs = append(variantArgs, argValue)
		} else {
			err := "Unknown type of argument value when calling parent method: " + arg.Type().String()
			Log.Error(err)
			panic(err)
		}
	}
	log.Println("  Built variant arguments: ", variantArgs)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(len(variantArgs)))
	log.Println("    C Array: ", cArgsArray)

	// Loop through and add each argument to our C args array.
	for i, arg := range variantArgs {
		C.add_element(cArgsArray, arg, C.int(i))
	}
	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	// Here we're just using a CString
	log.Println("  Building return value.")
	ret := unsafe.Pointer(C.CString(""))

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		ret,        // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	var retValue reflect.Value
	if _, ok := godotToGoConversionMap[returns]; ok {
		retValue = godotToGoConversionMap[returns](ret)
	} else {
		panic("Return type not found when calling parent method: " + returns)
	}

	// Return the converted variant.
	return retValue
}
