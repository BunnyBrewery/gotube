package hehe

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
