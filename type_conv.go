package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main(){
	var sum int = 1000
	var num int = 33
	fmt.Printf("sum-size = %d\n",unsafe.Sizeof(sum))
	fmt.Printf("num-size = %d\n",unsafe.Sizeof(num))


	var avg_int = sum/num
	fmt.Printf("avg-int-size = %d\n",unsafe.Sizeof(avg_int))
	fmt.Printf("avg-int-type = %s\n",reflect.TypeOf(avg_int))
	fmt.Println("avg-int = ", avg_int)

	var avg_float = float64(sum)/float64(num)
	fmt.Printf("avg-float-size = %d\n",unsafe.Sizeof(avg_float))
	fmt.Printf("avg-float-type = %s\n",reflect.TypeOf(avg_float))
	fmt.Println("avg-float = ", avg_float)

}