// Package openstack describes an openstack terraform provider.
package openstack

var Arguments map[string][]string

func init() {
	Arguments = map[string][]string{
		"network": []string{"access_ip_v4", "floating_ip"},
	}
}
