package main

import (
	"fmt"
	"github.com/jeffail/gabs"
)

func main() {
	jsonParsed, _ := gabs.ParseJSON([]byte(`{
			"outter":{
				"inner":{
					"value1":10,
					"value2":22
					},
					"alsoInner":{
						"value1":20
					}
				}
				}`))

	var value float64
	var ok bool

	if value, ok = jsonParsed.Path("outter.inner.value1").Data().(float64); ok {
		fmt.Printf("value: %v, ok: %v \n", value, ok)
		// value == 10.0, ok == true
	}

	if value, ok = jsonParsed.Search("outter", "inner", "value1").Data().(float64); ok {
		fmt.Printf("value: %v, ok: %v \n", value, ok)
		// value == 10.0, ok == true
	}

	if value, ok = jsonParsed.Path("does.not.exist").Data().(float64); !ok {
		fmt.Printf("value: %v, ok: %v \n", value, ok)
		// value == 0.0, ok == false
	}

	var containerObj *gabs.Container
	containerObj = jsonParsed.Path("does.not.exist")
	fmt.Println(containerObj)

	/*
	* Working with arrays
	 */

}
