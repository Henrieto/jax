package command

import "github.com/spf13/cobra"

type Flag struct {
	Requird    bool
	Name       string
	Short_Name string
	Help_Text  string
}

type Cmd struct {
	cobraCommand *cobra.Command
}

func (flag *Cmd) GetFlag(key string) (string, error) {
	value := flag.cobraCommand.Flags().Lookup(key).Value
	if value != nil {
		return value.String(), nil
	}
	return "", nil
}

type Command struct {
	Use   string
	Short string
	Long  string
	Run   func(*Cmd, []string)
	Flags []*Flag
}

func (cmd *Command) RunFunc() func(*cobra.Command, []string) {
	return func(cobra_cmd *cobra.Command, args []string) {
		_cmd := &Cmd{cobra_cmd}
		cmd.Run(_cmd, args)
	}
}

func (cmd *Command) Cmd() *cobra.Command {
	cobra_cmd := &cobra.Command{
		Use:   cmd.Use,
		Short: cmd.Short,
		Long:  cmd.Long,
		Run:   cmd.RunFunc(),
	}
	for _, fl := range cmd.Flags {
		cobra_cmd.Flags().StringP(fl.Name, fl.Short_Name, "", fl.Help_Text)
		if fl.Requird {
			cobra_cmd.MarkFlagRequired(fl.Name)
		}
	}
	return cobra_cmd
}

type commandList struct {
	Commands []*Command
}
