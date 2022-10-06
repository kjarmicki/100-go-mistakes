package module_with_init

import "fmt"

/*
 * Init function downsides:
 * - no error management other than panic - especially third party libraries should not panic an application
 * - not testable directly
 * - output = global variables
 */

func init() {
	fmt.Println("this will get called")
}

func init() {
	fmt.Println("curiously, this will get called too")
}

func WontWork() {
	// init(); - nope, can't call init directly
}
