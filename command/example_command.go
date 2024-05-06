package command

import "fmt"

var StartCmd = Command{
	Use:   "start",
	Short: "start server",
	Long:  "a command used for starting the server",
	Run: func(cmd *Cmd, args []string) {
		port, _ := cmd.GetFlag("port")
		fmt.Println(port, args)
	},
	Flags: []*Flag{
		{
			Requird:    false,
			Name:       "port",
			Short_Name: "p",
			Help_Text:  "the port on which the server will be running",
		},
	},
}
