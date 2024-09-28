package models

type CiscoNxosShowIpMroutesVrfAll struct {
	Group	string	`json:"GROUP"`
	Source	string	`json:"SOURCE"`
	Vrf	string	`json:"VRF"`
	Oif_list	[]string	`json:"OIF_LIST"`
	Incoming_interface	string	`json:"INCOMING_INTERFACE"`
	Rpf_neighbor	string	`json:"RPF_NEIGHBOR"`
	Last_uptime	string	`json:"LAST_UPTIME"`
}