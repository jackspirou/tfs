package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"

	"github.com/jackspirou/tfs/transformations/ansible"
)

// AnsibleCmd transforms a terraform state file to ansible inventory.
var AnsibleCmd = &cobra.Command{
	Use:     "ansible",
	Aliases: []string{"a"},
	Short:   "Transform Terraform state file into ansible inventory.",
	Long:    "Parse and transform a terraform state file into ansible dynamic inventory.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("ansible needs a terraform state to transform")
		}

		raw, err := ioutil.ReadFile(args[0])
		if err != nil {
			log.Fatal(err)
		}

		inventory := ansible.New()

		result, err := inventory.Transform(bytes.NewReader(raw))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result)
	},
}
