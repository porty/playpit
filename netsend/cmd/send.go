package cmd

import (
	"strings"

	"github.com/porty/playpit/netsend/netsend"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Sends some text",
	Long: `Send some text over the network. For example:

netsend send hello
  Sends the text "hello"

netsend send < file.txt
  Sends the contents of the file file.txt

netsend send
  Sends the contents of stdout, until EOF (Ctrl-D on Unix, Ctrl-M on Windows, maybe)
`,
	Run: func(cmd *cobra.Command, args []string) {
		netsend.Send(strings.Join(args, " "))
	},
}

func init() {
	RootCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
