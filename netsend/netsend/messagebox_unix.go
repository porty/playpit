// +build !windows

package netsend

import "fmt"

func MessageBox(caption string, text string, flags int) {
	fmt.Println("=========================")
	fmt.Println(caption)
	fmt.Println("-------------------------")
	fmt.Println(text)
	fmt.Println("    [OK]")
	fmt.Println("=========================")
}
