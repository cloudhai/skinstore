package main

import (
	"fmt"
	"reflect"
	"time"
)

func main(){
	var t = time.Now()
	v := reflect.ValueOf(t)
	//k := v.Kind()
	fmt.Println(v.Uint())

}
