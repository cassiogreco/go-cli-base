package flags

import (
	"errors"
	"fmt"
	"os"
)

type FlagDescription struct {
	Flag        string
	Description string
	Valid       bool
	Execute     func(...interface{}) error
}

var AcceptedFlags map[string]FlagDescription = map[string]FlagDescription{
	"-h": FlagDescription{Flag: "-h", Description: "Help command", Valid: true, Execute: Help},
	"-t": FlagDescription{Flag: "-t", Description: "Test command", Valid: false, Execute: Test}, // Toggle the Valid bool to test
}

func GetFlags() []string {
	return os.Args[1:]
}

func ValidateFlags(flags []string) ([]string, error) {
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

func GetAcceptedFlags() []FlagDescription {
	var accepted []FlagDescription
	for _, flag := range AcceptedFlags {
		if flag.Valid {
			accepted = append(accepted, FlagDescription{Flag: flag.Flag, Description: flag.Description})
		}
	}
	return accepted
}

func Help(input ...interface{}) error {
	fmt.Println("\nSupported flags: ")
	fmt.Println("\n")
	if len(input) > 0 {
		flags := input[0].([]FlagDescription)
		for _, flag := range flags {
			fmt.Println(flag.Flag + " " + flag.Description)
		}
		fmt.Println("\n")
		return nil
	}
	return errors.New("There was an unexpexted error.")
}

func Test(input ...interface{}) error {
	return nil
}
