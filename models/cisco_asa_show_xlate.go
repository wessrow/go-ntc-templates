package models

type CiscoAsaShowXlate struct {
	Source_intf	string	`json:"SOURCE_INTF"`
	Source	[]string	`json:"SOURCE"`
	Destination_intf	string	`json:"DESTINATION_INTF"`
	Destination	[]string	`json:"DESTINATION"`
}