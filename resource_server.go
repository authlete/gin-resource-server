//
// Copyright (C) 2019 Authlete, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific
// language governing permissions and limitations under the
// License.

package main

import (
	"github.com/authlete/authlete-go-gin/endpoint"
	"github.com/authlete/authlete-go-gin/middleware"
	"github.com/gin-gonic/gin"
)

type ResourceServer struct {
	Engine *gin.Engine
}

func ResourceServer_New() *ResourceServer {
	server := ResourceServer{}
	server.init()

	return &server
}

func (self *ResourceServer) Run(addr ...string) error {
	return self.Engine.Run(addr...)
}

func (self *ResourceServer) init() {
	self.Engine = gin.Default()

	self.setupAuthleteApi()
	self.setupUserInfoEndpoint(`/api/userinfo`)
	self.setupTimeEndpoint(`/api/time`)
}

func (self *ResourceServer) setupAuthleteApi() {
	// Register middleware that creates an instance of api.AuthleteApi and
	// sets the instance to the given gin contxt with the key `AuthleteApi`.
	//
	// middleware.AuthleteApi_Toml(file string) loads settings from a TOML
	// file. middleware.AuthleteApi_Env() reads settings from the environment.
	// middleware.AuthleteApi_Conf(conf.AuthleteConfiguration) reads settings
	// from a given AuthleteConfiguration.
	//
	// The following code loads `authlete.toml`.
	self.Engine.Use(middleware.AuthleteApi_Toml(`authlete.toml`))
}

func (self *ResourceServer) setupUserInfoEndpoint(path string) {
	spi := UserInfoReqHandlerSpiImpl{}
	handler := endpoint.UserInfoEndpoint_Handler(&spi)

	// UserInfo endpoint (OpenID Connect Core 1.0 Section 5.3)
	self.Engine.GET(path, handler)
	self.Engine.POST(path, handler)
}

func (self *ResourceServer) setupTimeEndpoint(path string) {
	// Example of a protected resource endpoint
	self.Engine.GET(path, TimeEndpoint_Handler())
}
