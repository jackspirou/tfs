// Package openstack describes an openstack terraform provider.
package openstack

var Attributes map[string][]string

func init() {
	Attributes = map[string][]string{
		"network": []string{"access_ip_v4", "floating_ip"},
	}
}
