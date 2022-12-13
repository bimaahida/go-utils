package utils

import "encoding/json"

// Dump convert interface to json using json marshal
func Dump(i interface{}) string {
	return string(ToByte(i))
}

// ToByte convert interface to byte
func ToByte(i interface{}) []byte {
	bt, _ := json.Marshal(i)
	return bt
}
