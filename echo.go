package main

import (
	"fmt"
	"bufio"
	"os"
)
func main(){
	var name string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter:")
	name,_ = reader.ReadString('\n')
	// fmt.Scanln(&name) 
	fmt.Println(name)
}