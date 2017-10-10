
package cmd

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a macro",
	Long: `Run a macro in your library. Expects a name. For example:

marko run my cool macro`,
	Run: func(cmd *cobra.Command, args []string) {

		keyword := strings.Join(args, " ")
		config_commands := viper.GetStringSlice(keyword)


		fmt.Println("run called:")
		fmt.Println(keyword)

// fixit how do i get each command and execute them synchronously
		fmt.Println(config_commands)

		// for _, cmd := range config_commands {
		//     // element is the element from someSlice for where we are

		// 	fmt.Println(cmd)
		// }

	},
}

func init() {
	RootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
