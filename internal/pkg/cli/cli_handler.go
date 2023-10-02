package cli

import (
	"flag"
	"fmt"
)

const HelpText string = `
Yet Another CLI based Mock ReSTFul Server. 

Usage: imposter -c config.json 

OPTIONS 
	-c		Path to Config.json 
	-l		Enable logging
	-h 		Show this help message
			
`

var helpFunc = func(cmd *Command, args []string) {
	fmt.Printf(HelpText)
}

func NewHelpCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("help", flag.ExitOnError),
		Execute: helpFunc,
	}
	return cmd
}
