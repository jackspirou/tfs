package openstack

import (
	"strconv"
	"strings"

	"github.com/jackspirou/tfs/state"
)

// ComputeInstanceV2 represents a openstack compute resource.
type ComputeInstanceV2 struct {
	count int
	name  string
	*state.ResourceState
}

// NewComputeInstanceV2 returns a OpenStackComputeInstanceV2.
func NewComputeInstanceV2(r *state.ResourceState) *ComputeInstanceV2 {
	types := strings.Split(r.Type, ".")
	switch len(types) {
	case 2:
		return &ComputeInstanceV2{ResourceState: r, name: types[1]}
	case 3:
		i, err := strconv.Atoi(types[2])
		if err != nil {
			return &ComputeInstanceV2{ResourceState: r, name: types[1]}
		}
		return &ComputeInstanceV2{ResourceState: r, name: types[1], count: i}
	default:
		return &ComputeInstanceV2{ResourceState: r}
	}
}

// Address returns the bound network addresses of this compute resource.
func (c ComputeInstanceV2) Address() string {

	for _, arg := range Attributes["network"] {
		if ip := c.Primary.Attributes[arg]; ip != "" {
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
