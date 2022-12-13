// Package utils :nodoc:
package utils

// Int32PointerToInt64 :nodoc:
func Int32PointerToInt64(i *int32) int64 {
	if i != nil {
		return int64(*i)
	}
	return int64(0)
}
