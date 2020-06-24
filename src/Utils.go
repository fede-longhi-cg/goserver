package utils

import (
	"io/ioutil"
	"os"
)

//Check panics if error is not nil
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

//ReadFile reads file and returns its content
func ReadFile(filename string) []byte {
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	Check(err)

	return data
}
