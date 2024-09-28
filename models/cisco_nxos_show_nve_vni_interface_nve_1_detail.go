package models

type CiscoNxosShowNveVniInterfaceNve1Detail struct {
	Vni	string	`json:"VNI"`
	Nve_interface	string	`json:"NVE_INTERFACE"`
	Multicast_address	string	`json:"MULTICAST_ADDRESS"`
	Vni_state	string	`json:"VNI_STATE"`
	Mode	string	`json:"MODE"`
	Vni_type	string	`json:"VNI_TYPE"`
	Vni_flags	string	`json:"VNI_FLAGS"`
	Dci_multicast_address	string	`json:"DCI_MULTICAST_ADDRESS"`
	Provision_state	string	`json:"PROVISION_STATE"`
	Vlan_bd	string	`json:"VLAN_BD"`
	Svi_state	string	`json:"SVI_STATE"`
}