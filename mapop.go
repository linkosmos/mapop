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
