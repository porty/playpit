package cmd

import (
	"github.com/porty/playpit/netsend/netsend"
	"github.com/spf13/cobra"
)

var messageBox *bool

// recvCmd represents the recv command
var recvCmd = &cobra.Command{
	Use:   "recv",
	Short: "Receive some text",
	Long: `Receives some text. For example:

netsend recv
  Recieves one bit of text, returns

netsend recv --messagebox
  Recieves one bit of text, sends to a Win32 messagebox
`,
	Run: func(cmd *cobra.Command, args []string) {
		netsend.Recv(*messageBox)
	},
}

func init() {
	RootCmd.AddCommand(recvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	messageBox = recvCmd.Flags().BoolP("messagebox", "m", false, "Opens a Windows messagebox on message receive")
}
