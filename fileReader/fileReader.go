package fileReader

import (
	"io/ioutil"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFile(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	checkErr(err)

	return string(data)
}
