package flatten

// Flatten takes in an arbitrarily deep list-like structure
// and returns a flattened structure without any nil values.
func Flatten(list interface{}) []interface{} {
	output := []interface{}{}

	switch el := list.(type) {
	case []interface{}:
		for _, subel := range el {
			if subel == nil {
				continue
			}
			output = append(output, Flatten(subel)...)
		}
	default:
		output = append(output, el)
	}

	return output
}
