package command

// initialize a server interface
var Server interface{ Listen(...string) error }

var RunServerCmd = Command{
	Use:   "run",
	Short: "run server",
	Long:  "a command used for starting the server",
	Run: func(cmd *Cmd, args []string) {
		address, _ := cmd.GetFlag("address")
		err := Server.Listen(address)
		if err != nil {
			panic(err)
		}
	},
	Flags: []*Flag{
		{
			Requird:    false,
			Name:       "address",
			Short_Name: "a",
			Help_Text:  "the address on which the server will be running",
		},
	},
}
