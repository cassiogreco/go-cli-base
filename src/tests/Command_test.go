package tests

import (
	"flags"
	"testing"
)

func TestRespond(test *testing.T) {
	for _, command := range flags.AcceptedFlags {
		if command.Valid {
			if command.Flag == "-h" {
				err := command.Execute(flags.GetAcceptedFlags())
				if err != nil {
					test.Fail()
				}
			} else {
				err := command.Execute()
				if err != nil {
					test.Fail()
				}
			}
		}
	}
}
