package models

type CiscoIosShowIpDeviceTrackingAll struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Vlan_id	string	`json:"VLAN_ID"`
	Interface	string	`json:"INTERFACE"`
	Probe_timeout	string	`json:"PROBE_TIMEOUT"`
	State	string	`json:"STATE"`
	Source	string	`json:"SOURCE"`
}