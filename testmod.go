package golang2

import "fmt"

// Hi returns a friendly greeting
func Hi(name string) string {
	fmt.Println("hello world !!!!")
   return fmt.Sprintf("MSG from golangv1.0.5:  %s", name)
}

