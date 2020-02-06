package main

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Data struct {
	UserId   int
	UserName string `yaml:"user_name"`
	TextT    string `yaml:"Text"`
}

func main() {
	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	var m Data
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		fmt.Println(err)
	}
}

	fmt.Println(m)
