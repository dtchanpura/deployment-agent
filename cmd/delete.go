package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deleting a project (Not Supported)",
	Long: `Deletes the Project matching provided UUID.

Currently it is not supported. To delete a project just edit the configuration
file and delete the part starting with "- name:"  which is the start of new element.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete is not yet supported.")
	},
}

func init() {
	var name string
	var projectUUID string
	RootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVar(&name, "name", "", "Name of project.")
	deleteCmd.Flags().StringVar(&projectUUID, "uuid", "", "UUID of project.")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
