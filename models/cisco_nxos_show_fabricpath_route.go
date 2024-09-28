package models

type CiscoNxosShowFabricpathRoute struct {
	Ftag	string	`json:"FTAG"`
	Switch_id	string	`json:"SWITCH_ID"`
	Subswitch_id	string	`json:"SUBSWITCH_ID"`
	Ports	[]string	`json:"PORTS"`
	Distance	string	`json:"DISTANCE"`
	Metric	string	`json:"METRIC"`
}