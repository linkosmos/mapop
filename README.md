# mapop

Regular map[string]interface{} operations. Ruby enumerable inspired package.

[![Build Status](https://travis-ci.org/linkosmos/mapop.svg?branch=master)](https://travis-ci.org/linkosmos/mapop)
[![Coverage Status](https://coveralls.io/repos/github/linkosmos/mapop/badge.svg?branch=master)](https://coveralls.io/github/linkosmos/mapop?branch=master)
[![GoDoc](http://godoc.org/github.com/linkosmos/mapop?status.svg)](http://godoc.org/github.com/linkosmos/mapop)
[![Go Report Card](http://goreportcard.com/badge/linkosmos/mapop)](http://goreportcard.com/report/linkosmos/mapop)
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
 - Collect(input map[string]interface{}) (output map[string]interface{})
 - Merge(maps ...map[string]interface{}) (output map[string]interface{})
 - SelectFunc(f func(key string, value interface{}) bool, input map[string]interface{}) (output map[string]interface{})


### Usage

```go
    input :=  map[string]interface{}{
      "Key1": 2,
      "key3": nil,
      "val": 2,
      "val2": "str",
      "val3": 4,
    }
```

#### Keys

```go
  keys := mapop.Keys(input)

  > keys["Key1", "key3", "val", "val2", "val3"]

```

#### Reject

```go
  input = mapop.Reject(input, "val", "val2", "val3")

  > input{"Key1": 2, "key3": nil}
```

#### Select

```go
  input = mapop.Select(input, "val", "val2", "val3")

  > input{"val": 2, "val2": "str", "val3": 4}

```

#### Split

```go
  keys, values := mapop.Split(input, "val", "val2", "val3")

  > keys["Key1", "key3", "val", "val2", "val3"]
  > values[2,nil,2,"str",4]

```

#### Values

```go
  values := mapop.Values(input)

  > values[2,nil,2,"str",4]

```

#### MapKeys

```go
  input = mapop.MapKeys(strins.ToUpper, input)

  > input{"KEY1": 2, "KEY3": nil, "VAL": 2, "VAL2": "str", "VAL3": 4}

```

#### MapValues

```go
  input = mapop.MapValues(func(val interface{}) interface{} {
      return "-10"
  }, input)

  > input{"Key1": -10, "key3": -10, "val": -10, "val2": -10, "val3": -10}

```

#### Partition

```go
  partitioned := mapop.Partition(func(key string, value interface{}) bool {
    return strings.Contains(key, "val")
  }, input)

  > partitioned[0]{"Key1": 2, "key3": nil}
  > partitioned[1]{"val": 2, "val2": "str", "val3": 4}
```

#### Map

```go
  input = mapop.Map(func(key string, value interface{}) (string, interface{}) {
    if strings.Contains(key, "val") {
      return key, key
    } else {
      return key, value
    }
  }, input)

  > input{"Key1": 2, "key3": nil, "val": "val", "val2": "val2", "val3": "val3"}
```

#### Collect

```go
  input = mapop.Collect(input)

  > input{"Key1": 2, "val": 2, "val2": "str", "val3": 4}
```

#### Merge

```go
  input2 := map[string]interface{}{
    "a2": "str",
    "a3": 4,
  }

  input = mapop.Merge(input, input2)

  > input{"Key1": 2, "key3": nil, "val": 2, "val2": "str", "val3": 4, "a2": "str", "a3": 4}
```

#### SelectFunc

```go
  input = mapop.SelectFunc(func(key string, value interface{}) bool {
    return key == "val"
  }, input)

  > input{"val": 2}

```

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
