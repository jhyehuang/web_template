/*
Copyright © 2021 NAME HERE rickiey@qq.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/jhyehuang/web_template/intern"
	"net/http/httptest"
)

func main() {
	// create a new server instance
	rpcServer := jsonrpc.NewServer()
	//
	//// create a handler instance and register it
	//serverHandler := &intern.SimpleServerHandler{}
	//rpcServer.Register("SimpleServerHandler", serverHandler)

	// rpcServer is now http.Handler which will serve jsonrpc calls to SimpleServerHandler.AddGet
	// a method with a single int param, and an int response. The server supports both http and websockets.

	// serve the api
	testServ := httptest.NewServer(rpcServer)
	defer testServ.Close()
	//
	fmt.Println("URL: ", "ws://"+testServ.Listener.Addr().String())
	//
	intern.ServeAPI()

	return

}
