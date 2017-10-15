//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package streampeertcp

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

	"github.com/shadowapex/godot-go/godot/classes/streampeer"
)

/*
   TCP Stream peer. This object can be used to connect to TCP servers, or also is returned by a tcp server.
*/
type StreamPeerTCP struct {
	streampeer.StreamPeer
}

func (o *StreamPeerTCP) baseClass() string {
	return "StreamPeerTCP"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *StreamPeerTCP) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *StreamPeerTCP) getOwner() *C.godot_object {
	return o.owner
}

/*
   Connect to the specified host:port pair. A hostname will be resolved if valid. Returns [OK] on success or [FAILED] on failure.
*/
func (o *StreamPeerTCP) ConnectToHost(host string, port int64) int64 {
	log.Println("Calling StreamPeerTCP.ConnectToHost()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(host)
	goArguments[1] = reflect.ValueOf(port)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "connect_to_host", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Disconnect from host.
*/
func (o *StreamPeerTCP) DisconnectFromHost() {
	log.Println("Calling StreamPeerTCP.DisconnectFromHost()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "disconnect_from_host", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Return the IP of this peer.
*/
func (o *StreamPeerTCP) GetConnectedHost() string {
	log.Println("Calling StreamPeerTCP.GetConnectedHost()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_connected_host", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the port of this peer.
*/
func (o *StreamPeerTCP) GetConnectedPort() int64 {
	log.Println("Calling StreamPeerTCP.GetConnectedPort()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_connected_port", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the status of the connection, one of STATUS_* enum.
*/
func (o *StreamPeerTCP) GetStatus() int64 {
	log.Println("Calling StreamPeerTCP.GetStatus()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_status", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *StreamPeerTCP) IsConnectedToHost() bool {
	log.Println("Calling StreamPeerTCP.IsConnectedToHost()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_connected_to_host", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   StreamPeerTCPImplementer is an interface for StreamPeerTCP objects.
*/
type StreamPeerTCPImplementer interface {
	class.Class
}