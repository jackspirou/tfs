package ansible

import (
	"io"

	"github.com/jackspirou/tfs/resources"
	"github.com/jackspirou/tfs/state"
)

// Inventory parses a terraform statefile and returns Ansible Inventory as JSON.
func Inventory(format string, src io.Reader) (map[string][]string, error) {

	s, err := state.ReadState(src)
	if err != nil {
		return nil, err
	}

	groups := make(map[string][]string, 0)

	for _, m := range s.Modules {
		for _, r := range m.Resources {
			resource, err := resources.New(r)
			if err != nil {
				return nil, err
			}
			g := resource.Groups()
			for _, i := range g {
				if i != "" {
					groups[i] = append(groups[i], resource.PublicIP())
				}
			}
		}
	}

	return groups, nil
}
