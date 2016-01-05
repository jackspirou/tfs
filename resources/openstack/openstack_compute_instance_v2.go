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

// PublicIP returns the public bound network addresses of the compute resource.
func (c ComputeInstanceV2) PublicIP() string {

	for _, arg := range Attributes["network"] {
		if ip := c.Primary.Attributes[arg]; ip != "" {
			return ip
		}
	}

	return ""
}

// Groups returns the group names of the metadata.groups attribute.
func (c ComputeInstanceV2) Groups() []string {
	groups := c.Primary.Attributes["metadata.groups"]
	return strings.Split(groups, ",")
}

// Name returns the name for this compute resource.
func (c ComputeInstanceV2) Name() string {
	return c.name
}

// Count returns the current count for this compute resource.
func (c ComputeInstanceV2) Count() int {
	return c.count
}
