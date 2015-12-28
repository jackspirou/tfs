package main

import "github.com/jackspirou/tfs/cmd"

func main() {
	cmd.RootCmd.AddCommand(cmd.AnsibleCmd)
	cmd.RootCmd.AddCommand(cmd.VersionCmd)
	cmd.RootCmd.Execute()
}
