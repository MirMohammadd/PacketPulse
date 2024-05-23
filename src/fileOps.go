package main

import (
	"io/ioutil"
	"log"
)

func WriteFile(data []byte) {
	err := ioutil.WriteFile("output.txt", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
