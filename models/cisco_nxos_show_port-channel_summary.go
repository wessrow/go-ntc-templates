package models

type CiscoNxosShowPortChannelSummary struct {
	Bundle_name	string	`json:"BUNDLE_NAME"`
	Bundle_status	string	`json:"BUNDLE_STATUS"`
	Bundle_protocol	string	`json:"BUNDLE_PROTOCOL"`
	Bundle_protocol_state	string	`json:"BUNDLE_PROTOCOL_STATE"`
	Member_interface	[]string	`json:"MEMBER_INTERFACE"`
	Member_interface_status	[]string	`json:"MEMBER_INTERFACE_STATUS"`
}