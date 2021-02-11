package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/instrumenta/kubeval/kubeval"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("./cmd/playground/sample.yaml")
	if err != nil {
		panic(err)
	}
	//t := make(map[interface{}]interface{}, 1)
	//err = yaml.Unmarshal(content, &t)
	//if err != nil {
	//	panic(err)
	//}
	result, err := kubeval.Validate(content)
	if err != nil {
		fmt.Println(err)
	}
	spew.Dump(result)

}
