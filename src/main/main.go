package main

import (
	"flags"
	"fmt"
)

func main() {
	inputFlags, err := flags.ValidateFlags(flags.GetFlags())
	if err == nil {
		for _, flag := range inputFlags {
			if flag == "-h" {
				problem := flags.AcceptedFlags[flag].Execute(flags.GetAcceptedFlags())
				if problem != nil {
					fmt.Println("\nThere was an error executing your command: " + problem.Error())
				}
			} else {
				problem := flags.AcceptedFlags[flag].Execute()
				if problem != nil {
					fmt.Println("\nThere was an error executing your command: " + problem.Error())
				}
			}
		}
	} else {
		fmt.Println(err.Error())
	}
}
