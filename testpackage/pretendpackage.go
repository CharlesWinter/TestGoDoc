//package testpackage has a package level comment up here
package testpackage

import "fmt"

// ExportedFunction has a comment right here
func ExportedFunction() {
	fmt.Println("vim-go")
}

func AnotherExportedFunction() {
	fmt.Println("its another function")	
}
