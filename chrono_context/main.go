package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	madnessState:= 0
	ctx, cancel := context.WithCancel(context.Background())

	for madnessState == 5{
		go func(ctx context.Context){
			for{
				select{
				case <- ctx.Done():
					fmt.Println("Magia parada por força maior")
					madnessState++
					return
				default:
					fmt.Println("Azazel conseguiu fazer Chrono Seal")
					time.Sleep(3* time.Second)
				}
			}
		}(ctx)
		time.Sleep(3* time.Second)
		fmt.Println("Azazel não conseguiu fazer Chrono Seal")
		cancel()
		time.Sleep(3*time.Second)
		}
}