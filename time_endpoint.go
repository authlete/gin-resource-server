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
	"fmt"
	"time"

	"github.com/authlete/authlete-go-gin/endpoint"
	"github.com/authlete/authlete-go-gin/web"
	"github.com/gin-gonic/gin"
)

type TimeEndpoint struct {
	endpoint.BaseEndpoint
	resutil web.ResponseUtility
}

func TimeEndpoint_Handler() gin.HandlerFunc {
	// Instance of time endpoint
	endpoint := TimeEndpoint{}

	return func(ctx *gin.Context) {
		endpoint.Handle(ctx)
	}
}

func (self *TimeEndpoint) Handle(ctx *gin.Context) {
	// Validate the access token included in the request.
	valid, validator := self.ValidateAccessToken(ctx, nil)

	// If the access token is not valid.
	if valid == false {
		// Generate an error response that conforms to RFC 6750.
		validator.Deny(ctx)
		return
	}

	t := time.Now()

	// HTTP response body
	content := fmt.Sprintf(`{
  "year":   %d,
  "month":  %d,
  "day":    %d,
  "hour":   %d,
  "minute": %d,
  "second": %d
}
`, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	// Response to the API caller.
	self.resutil.OkJson(ctx, content)
}
