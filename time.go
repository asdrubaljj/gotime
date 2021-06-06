package main

import (
	"fmt"
	"time"
)

func main(){	
	t := time.Now()
	fmt.Println(t.Day(), t.Month(), t.Year(),"\n", t.Hour(),":",t.Minute())
}
