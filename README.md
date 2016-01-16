# mapop

Regular map[string]interface{} operations. Ruby enumerable inspired package.

[![Build Status](https://travis-ci.org/linkosmos/mapop.svg?branch=master)](https://travis-ci.org/linkosmos/mapop)
[![GoDoc](http://godoc.org/github.com/linkosmos/mapop?status.svg)](http://godoc.org/github.com/linkosmos/mapop)
[![BSD License](http://img.shields.io/badge/license-BSD-blue.svg)](http://opensource.org/licenses/BSD-3-Clause)

### Methods
 - Keys(input map[string]interface{}) (keys []string)
 - Reject(input map[string]interface{}, keys ...string) map[string]interface{}
 - Select(input map[string]interface{}, keys ...string) map[string]interface{}
 - Split(input map[string]interface{}) (keys []string, values []interface{})
 - Values(input map[string]interface{}) (values []interface{})
 - MapKeys(f func(string) string, input map[string]interface{}) (output map[string]interface{})
 - MapValues(f func(interface{}) interface{}, input map[string]interface{}) (output map[string]interface{})
 - Partition(f func(string, interface{}) bool, input map[string]interface{}) (partition []map[string]interface{})
 - Map(f func(key string, value interface{}) (string, interface{}), input map[string]interface{}) (output map[string]interface{})

### License

Copyright (c) 2015, linkosmos
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of mapop nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
