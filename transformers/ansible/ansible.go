package ansible

import (
	"io"

	"github.com/jackspirou/tfs/resources"
	"github.com/jackspirou/tfs/state"
)

// Inventory parses a terraform statefile and returns Ansible Inventory as JSON.
func Inventory(format string, src io.Reader) (string, error) {

	s, err := state.ReadState(src)
	if err != nil {
		return "", err
	}

	var addr string

	for _, m := range s.Modules {
		for _, r := range m.Resources {
			resource, err := resources.New(r)
			if err != nil {
				return "", err
			}
			addr = resource.PublicIP()
		}
	}

	return addr, nil
}
