package cala

import (
	"fmt"
	"time"
)

func Timeshow() {
	for {
		time.Sleep(5 * time.Second)
		fmt.Println(time.Now())

	}

}
