package ansible

import (
	"io"

	"github.com/jackspirou/tfs/resources"
	"github.com/jackspirou/tfs/state"
)

// Inventory represents Ansible Inventory.
type Inventory struct{}

// New returns a new ansible Inventory.
func New() *Inventory {
	return &Inventory{}
}

// Transform parses a terraform statefile and returns Ansible Inventory as JSON.
func (i Inventory) Transform(src io.Reader) (string, error) {

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
			addr = resource.Address()
		}
	}

	return addr, nil
}
