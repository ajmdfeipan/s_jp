package cala

import (
	"fmt"
	"time"
)

func Timeshow() {
	for {
		time.Sleep(20 * time.Second)
		fmt.Println(time.Now())

	}

}
