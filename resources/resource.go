// Package resources describes a terraform provider.
package resources

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jackspirou/tfs/resources/openstack"
	"github.com/jackspirou/tfs/state"
)

// Resource describes a terraform resource provider.
type Resource interface {
	Address() string
	Name() string
	Count() int
}

// New takes a terraform resource type, as well as a terraform ResourceState
// and returns an object that fufills the Resource interface.
func New(r *state.ResourceState) (Resource, error) {
	typ := r.Type
	types := strings.Split(typ, ".")
	if types[0] == "" {
		return nil, errors.New("unable to determine resource type")
	}

	switch types[0] {
	case "openstack_compute_instance_v2":
		return openstack.NewComputeInstanceV2(r), nil
	default:
		return nil, fmt.Errorf("%s is a supported resource type", types[0])
	}
}
