package main

import (
	"fmt"
	"time"

	"github.com/ermanimer/progress_bar"
)

func main() {
	//create new progress bar
	pb := progress_bar.DefaultProgressBar(100)
	//start
	err := pb.Start()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//update
	for value := 1; value <= 100; value++ {
		time.Sleep(20 * time.Millisecond)
		err := pb.Update(float64(value))
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}
