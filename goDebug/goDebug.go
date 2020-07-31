/*
When debugging your code in goLang you may need a pretty print function not just for int or strings,
but capable to render any kind of data structure. Here it is, simple and useful. Enjoy it!
*/

package goDebug

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func interface2string(i interface{}) string {
	pretty, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		fmt.Println("error: ", err)
	}
	return string(pretty)
}

func print_static(s string, i interface{}) {
	pretty := interface2string(i)
	fmt.Printf("\ngoDebug: %s (%s) => %s\n", s, reflect.TypeOf(i).String(), string(pretty))
	// fmt.Printf("\n\ngoDebug: %s RAW => %s", s, i)
	// fmt.Printf("\n\ngoDebugPlain: %+v\n", i)
}

// Debugging function to show a data structure w/ pretty style
// If called w/ 1 parameter, it just show its value
// If call w/ 2 parameters, the first one should be the name of the structure to enhance debug visibility
func Print(params ...interface{}) {
	if len(params) == 0 {
		fmt.Println("\nError: not enough parameters when calling Print function")
	} else if len(params) == 1 {
		print_static("", params[0])
	} else if len(params) == 2 {
		switch params[0].(type) {
		case string:
			print_static(params[0].(string), params[1])
		default:
			print_static("", params[0])
			print_static("", params[1])
		}
	} else {
		fmt.Println("\nError: too many parameters when calling Print function")
	}
}
