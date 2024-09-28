package models

type CiscoIosShowEtherchannelSummary struct {
	Group	string	`json:"GROUP"`
	Bundle_name	string	`json:"BUNDLE_NAME"`
	Bundle_status	string	`json:"BUNDLE_STATUS"`
	Bundle_protocol	string	`json:"BUNDLE_PROTOCOL"`
	Member_interface	[]string	`json:"MEMBER_INTERFACE"`
	Member_interface_status	[]string	`json:"MEMBER_INTERFACE_STATUS"`
}