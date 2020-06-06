<<<<<<< HEAD
package system
=======
package hehe
>>>>>>> 9997f684c210492b7ca121eb08ec10cc974a5a34

import (
	"fmt"
	"os"
)

func Path() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("starting dir:", wd)

	if err := ps.Chdir("/"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("final dir:", wd)
}
