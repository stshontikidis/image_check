package util

import (
	"os"
)

// CheckErr is boiler plate err check
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// WipeAndWrite will remove old contents and write new value
func WipeAndWrite(f *os.File, contents string) {
	f.Truncate(0)
	f.Seek(0, 0)
	f.WriteString(contents)
}
