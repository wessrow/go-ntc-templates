package models

type CiscoS300ShowInterfacesSwitchport struct {
	Interface	string	`json:"INTERFACE"`
	Interface_mode	string	`json:"INTERFACE_MODE"`
	Gvrp_status	string	`json:"GVRP_STATUS"`
	Ingress_filtering_status	string	`json:"INGRESS_FILTERING_STATUS"`
	Acceptable_frame_type	string	`json:"ACCEPTABLE_FRAME_TYPE"`
	Native_vlan	string	`json:"NATIVE_VLAN"`
	Vlan	[]string	`json:"VLAN"`
	Name	[]string	`json:"NAME"`
	Egress_rule	[]string	`json:"EGRESS_RULE"`
	Interface_membership_type	[]string	`json:"INTERFACE_MEMBERSHIP_TYPE"`
	Forbidden_vlan	[]string	`json:"FORBIDDEN_VLAN"`
	Forbidden_vlan_name	[]string	`json:"FORBIDDEN_VLAN_NAME"`
	Mac_based_vlan_group_id	[]string	`json:"MAC_BASED_VLAN_GROUP_ID"`
	Mac_based_vlan_id	[]string	`json:"MAC_BASED_VLAN_ID"`
}