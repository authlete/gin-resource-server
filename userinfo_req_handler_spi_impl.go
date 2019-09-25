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
	"github.com/authlete/authlete-go-gin/handler/spi"
)

type UserInfoReqHandlerSpiImpl struct {
	spi.UserInfoReqHandlerSpiAdapter
	user  *UserEntity
	tried bool
}

func (self *UserInfoReqHandlerSpiImpl) GetUserClaimValue(
	subject string, claimName string, languageTag string) interface{} {
	user := self.getUserBySubject(subject)

	if user == nil {
		return nil
	}

	return user.GetClaim(claimName, languageTag)
}

func (self *UserInfoReqHandlerSpiImpl) getUserBySubject(subject string) *UserEntity {
	if self.tried == false {
		self.user = UserDatabase_Get().GetBySubject(subject)
		self.tried = true
	}

	return self.user
}
