package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"

	"github.com/jackspirou/tfs/transformers/ansible"
)

func init() {
	AnsibleCmd.Flags().String("fmt", "ini", "Formats supported are INI, JSON, YAML, TOML")
}

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

		result, err := ansible.Inventory("json", bytes.NewReader(raw))
		if err != nil {
			log.Fatal(err)
		}

		for key, value := range result {
			fmt.Printf("[%s]\n", key)
			for _, ip := range value {
				fmt.Printf("%s\n", ip)
			}
			fmt.Println("")
		}
	},
}
