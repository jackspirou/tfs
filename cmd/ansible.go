package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/jackspirou/tfs/ansible"
)

// AnsibleCmd transforms a terraform state file to ansible inventory.
var AnsibleCmd = &cobra.Command{
	Use:     "ansible",
	Aliases: []string{"a"},
	Short:   "Terraform state file to ansible inventory.",
	Long:    "Parse and transform a terraform state file into ansible dynamic inventory.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("ansible needs a terraform state to transform")
		}
		if err := ansible.Inventory(args[0]); err != nil {
			log.Fatal(err)
		}
	},
}
