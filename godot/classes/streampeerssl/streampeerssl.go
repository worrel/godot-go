//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package streampeerssl

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
   SSL Stream peer. This object can be used to connect to SSL servers.
*/
type StreamPeerSSL struct {
	streampeer.StreamPeer
}

func (o *StreamPeerSSL) baseClass() string {
	return "StreamPeerSSL"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *StreamPeerSSL) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *StreamPeerSSL) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *StreamPeerSSL) AcceptStream(stream *StreamPeer) int64 {
	log.Println("Calling StreamPeerSSL.AcceptStream()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(stream)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "accept_stream", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Connect to a peer using an underlying [StreamPeer] "stream", when "validate_certs" is true, [StreamPeerSSL] will validate that the certificate presented by the peer matches the "for_hostname".
*/
func (o *StreamPeerSSL) ConnectToStream(stream *StreamPeer, validateCerts bool, forHostname string) int64 {
	log.Println("Calling StreamPeerSSL.ConnectToStream()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(stream)
	goArguments[1] = reflect.ValueOf(validateCerts)
	goArguments[2] = reflect.ValueOf(forHostname)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "connect_to_stream", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Disconnect from host.
*/
func (o *StreamPeerSSL) DisconnectFromStream() {
	log.Println("Calling StreamPeerSSL.DisconnectFromStream()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "disconnect_from_stream", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Return the status of the connection, one of STATUS_* enum.
*/
func (o *StreamPeerSSL) GetStatus() int64 {
	log.Println("Calling StreamPeerSSL.GetStatus()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_status", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   StreamPeerSSLImplementer is an interface for StreamPeerSSL objects.
*/
type StreamPeerSSLImplementer interface {
	class.Class
}
