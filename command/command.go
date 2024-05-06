package command

var CommandList = commandList{[]*Command{}}

func LoadCommands() {
	for _, cmd := range CommandList.Commands {
		rootCmd.AddCommand(cmd.Cmd())
	}
}

func Register(cmds ...*Command) {
	CommandList.Commands = append(CommandList.Commands, cmds...)
}
