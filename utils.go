package mapop

func includes(k string, keys ...string) bool {
	for _, key := range keys {
		if key == k {
			return true
		}
	}
	return false
}

func selectORreject(reject bool, input map[string]interface{}, keys ...string) (output map[string]interface{}) {
	size := len(input)
	keysSize := len(keys)
	if size <= 0 {
		return nil
	}
	if keysSize <= 0 {
		return input
	}
	if size >= keysSize {
		size = size - keysSize
	}
	output = make(map[string]interface{}, size)
	for key, value := range input {
		if reject {
			if !includes(key, keys...) {
				output[key] = value
			}
		} else {
			if includes(key, keys...) {
				output[key] = value
			}
		}
	}
	return output
}
