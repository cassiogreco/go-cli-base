package main

import (
	"errors"
	"fmt"
	"os"
)

type FlagDescription struct {
	Flag        string
	Description string
	Valid       bool
	Execute     func(...interface{})
}

var AcceptedFlags map[string]FlagDescription = map[string]FlagDescription{
	"-h": FlagDescription{Flag: "-h", Description: "Help command", Valid: true, Execute: Help},
	"-t": FlagDescription{Flag: "-t", Description: "Test command", Valid: false, Execute: Test},
}

func main() {
	flags, err := validateFlags(getFlags())
	if err == nil {
		for _, flag := range flags {
			if flag == "-h" {
				AcceptedFlags[flag].Execute(getAcceptedFlags())
			} else {
				AcceptedFlags[flag].Execute()
			}
		}
	} else {
		fmt.Println(err)
	}
}

func getFlags() []string {
	return os.Args[1:]
}

func validateFlags(flags []string) ([]string, error) {
	for _, flag := range flags {
		if string(flag[0]) != "-" {
			return []string{}, errors.New("Flag '" + flag + "' is not formatted correctly. Flags must start with a '-' in front. Run command with -h for help.")
		}
		if _, exists := AcceptedFlags[flag]; exists == false || !AcceptedFlags[flag].Valid {
			return []string{}, errors.New("Flag '" + flag + "' does not exist. Run command with -h for help.")
		}
	}
	return flags, nil
}

func getAcceptedFlags() []FlagDescription {
	var accepted []FlagDescription
	for _, flag := range AcceptedFlags {
		if flag.Valid {
			accepted = append(accepted, FlagDescription{Flag: flag.Flag, Description: flag.Description})
		}
	}
	return accepted
}

func Help(flags ...interface{}) {
	fmt.Println("Supported flags: ")
	fmt.Println("\n")
	for _, flag := range flags {
		commands := flag.([]FlagDescription)
		for _, command := range commands {
			fmt.Println(command.Flag + " " + command.Description + "\n")
		}
	}
}

func Test(input ...interface{}) {

}
