package cli

import "flag"

type Command struct {
	flags   *flag.FlagSet
	Execute func(cmd *Command, args []string)
}

func (cmd *Command) Init(args []string) error {
	return cmd.flags.Parse(args)
}

func (cmd *Command) Called() bool {
	return cmd.flags.Parsed()
}

func (cmd *Command) Run() {
	cmd.Execute(cmd, cmd.flags.Args())
}
