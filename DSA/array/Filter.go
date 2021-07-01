package array

// Filter _
func Filter(s []interface{}, fn func(interface{}) bool) []interface{} {
	var p []interface{}
	for _, value := range s {
		if fn(value) {
			p = append(p, value)
		}
	}

	return p
}
