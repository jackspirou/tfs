package openstack

import (
	"strconv"
	"strings"

	"github.com/jackspirou/tfs/state"
)

// ComputeInstanceV2 represents a openstack compute resource.
type ComputeInstanceV2 struct {
	name     string
	count    int
	resource *state.ResourceState
}

// NewComputeInstanceV2 returns a OpenStackComputeInstanceV2.
func NewComputeInstanceV2(r *state.ResourceState) *ComputeInstanceV2 {
	typ := r.Type
	types := strings.Split(typ, ".")
	switch len(types) {
	case 2:
		return &ComputeInstanceV2{resource: r, name: types[1]}
	case 3:
		i, err := strconv.Atoi(types[2])
		if err != nil {
			return &ComputeInstanceV2{resource: r, name: types[1]}
		}
		return &ComputeInstanceV2{resource: r, name: types[1], count: i}
	default:
		return &ComputeInstanceV2{resource: r}
	}
}

// Address returns the bound network addresses of this compute resource.
func (c ComputeInstanceV2) Address() string {

	for _, arg := range Arguments["network"] {
		if ip := c.resource.Primary.Attributes[arg]; ip != "" {
			return ip
		}
	}

	return ""
}

// Name returns the name for this compute resource.
func (c ComputeInstanceV2) Name() string {
	return c.name
}

// Count returns the current count for this compute resource.
func (c ComputeInstanceV2) Count() int {
	return c.count
}
