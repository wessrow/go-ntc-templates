package models

type CiscoNxosShowIpArpDetailVrfAll struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Age	string	`json:"AGE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
	Physical_interface	string	`json:"PHYSICAL_INTERFACE"`
	Flags	string	`json:"FLAGS"`
	Vrf	string	`json:"VRF"`
}