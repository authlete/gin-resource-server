#
# Copyright (C) 2019 Authlete, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
# either express or implied. See the License for the specific
# language governing permissions and limitations under the
# License.


#==================================================
# VARIABLES
#==================================================
APPLICATION  = gin-resource-server
SOURCE_FILES = $(wildcard *.go)
PORT         = 8081


#==================================================
# TARGETS
#==================================================
.PHONY: default clean run

default: $(APPLICATION)

clean:
	@-rm -f $(APPLICATION)

run: $(APPLICATION)
	PORT=$(PORT) ./$(APPLICATION)

$(APPLICATION): $(SOURCE_FILES)
	go build -o $@ $(SOURCE_FILES)

