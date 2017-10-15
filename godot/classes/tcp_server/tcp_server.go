//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package tcp_server

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
   TCP Server class. Listens to connections on a port and returns a [StreamPeerTCP] when got a connection.
*/
type TCP_Server struct {
	reference.Reference
}

func (o *TCP_Server) baseClass() string {
	return "TCP_Server"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *TCP_Server) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *TCP_Server) getOwner() *C.godot_object {
	return o.owner
}

/*
   Return true if a connection is available for taking.
*/
func (o *TCP_Server) IsConnectionAvailable() bool {
	log.Println("Calling TCP_Server.IsConnectionAvailable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_connection_available", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Listen on the "port" binding to "bind_address". If "bind_address" is set as "*" (default), the server will listen on all available addresses (both IPv4 and IPv6). If "bind_address" is set as "0.0.0.0" (for IPv4) or "::" (for IPv6), the server will listen on all available addresses matching that IP type. If "bind_address" is set to any valid address (e.g. "192.168.1.101", "::1", etc), the server will only listen on the interface with that addresses (or fail if no interface with the given address exists).
*/
func (o *TCP_Server) Listen(port int64, bindAddress string) int64 {
	log.Println("Calling TCP_Server.Listen()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(port)
	goArguments[1] = reflect.ValueOf(bindAddress)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "listen", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Stop listening.
*/
func (o *TCP_Server) Stop() {
	log.Println("Calling TCP_Server.Stop()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "stop", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   If a connection is available, return a StreamPeerTCP with the connection/
*/
func (o *TCP_Server) TakeConnection() *StreamPeerTCP {
	log.Println("Calling TCP_Server.TakeConnection()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "take_connection", goArguments, "*StreamPeerTCP")

	returnValue := goRet.Interface().(*StreamPeerTCP)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   TCP_ServerImplementer is an interface for TCP_Server objects.
*/
type TCP_ServerImplementer interface {
	class.Class
}
