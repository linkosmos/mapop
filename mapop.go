package mapop

// Split - split map string<=>interface to ordered keys and values
func Split(input map[string]interface{}) (keys []string, values []interface{}) {
	size := len(input)
	if size <= 0 {
		return nil, nil
	}
	keys = make([]string, 0, size)
	values = make([]interface{}, 0, size)
	for key, value := range input {
		keys = append(keys, key)
		values = append(values, value)
	}
	return keys, values
}

// Keys - return map keys
func Keys(input map[string]interface{}) (keys []string) {
	keys, _ = Split(input)
	return keys
}

// Values - return map values
func Values(input map[string]interface{}) (values []interface{}) {
	_, values = Split(input)
	return values
}

// Select - select specified keys from map and get new map
func Select(input map[string]interface{}, keys ...string) map[string]interface{} {
	return selectORreject(false, input, keys...)
}

// Reject - reject specified keys from map and get new map
func Reject(input map[string]interface{}, keys ...string) map[string]interface{} {
	return selectORreject(true, input, keys...)
}

// MapKeys - maps map keys, values remain unchanged and associated
func MapKeys(f func(string) string, input map[string]interface{}) (output map[string]interface{}) {
	size := len(input)
	if size == 0 {
		return input
	}
	output = make(map[string]interface{}, size)
	for key, value := range input {
		output[f(key)] = value
	}
	return output
}

// MapValues - maps map values, keys and values association remains unchanged
func MapValues(f func(interface{}) interface{}, input map[string]interface{}) (output map[string]interface{}) {
	size := len(input)
	if size == 0 {
		return input
	}
	output = make(map[string]interface{}, size)
	for key, value := range input {
		output[key] = f(value)
	}
	return output
}

// Partition - returns two maps in array, the first containing the elements
// for which the function evaluates to true, the second containing the rest.
func Partition(f func(string, interface{}) bool, input map[string]interface{}) (partition []map[string]interface{}) {
	partition = make([]map[string]interface{}, 2)
	size := len(input)
	if size == 0 {
		partition[0] = input
		partition[1] = nil
		return partition
	}
	// Assuming half of key values will be partitioned
	partition[0] = make(map[string]interface{}, size/2)
	partition[1] = make(map[string]interface{}, size/2)
	for key, value := range input {
		if f(key, value) {
			partition[0][key] = value
		} else {
			partition[1][key] = value
		}
	}
	return partition
}

// Map - maps key or values as defined in function
func Map(f func(key string, value interface{}) (string, interface{}), input map[string]interface{}) (output map[string]interface{}) {
	size := len(input)
	if size == 0 {
		return input
	}
	output = make(map[string]interface{}, size)
	var mappedKey string
	var mappedValue interface{}
	for key, value := range input {
		mappedKey, mappedValue = f(key, value)
		output[mappedKey] = mappedValue
	}
	return output
}

// Collect - removes all nil interface values and returns clean map
func Collect(input map[string]interface{}) (output map[string]interface{}) {
	size := len(input)
	if size == 0 {
		return input
	}
	output = make(map[string]interface{}, size)
	for key, value := range input {
		if value != nil {
			output[key] = value
		}
	}
	return output
}

// Merge - merges given maps into 1 map
// values will be overridden by last matching key - value
func Merge(maps ...map[string]interface{}) (output map[string]interface{}) {
	size := len(maps)
	if size == 0 {
		return output
	}
	if size == 1 {
		return maps[0]
	}
	output = make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			output[k] = v
		}
	}
	return output
}

// SelectFunc - select map keys that yield true either by key or value
func SelectFunc(f func(key string, value interface{}) bool, input map[string]interface{}) (output map[string]interface{}) {
	size := len(input)
	if size == 0 || f == nil {
		return input
	}
	output = make(map[string]interface{})
	for k, v := range input {
		if f(k, v) {
			output[k] = v
		}
	}
	return output
}
