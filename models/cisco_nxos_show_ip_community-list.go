package models

type CiscoNxosShowIpCommunityList struct {
	Type	string	`json:"TYPE"`
	Name	string	`json:"NAME"`
	Seq	string	`json:"SEQ"`
	Action	string	`json:"ACTION"`
	As	[]string	`json:"AS"`
	Tag	[]string	`json:"TAG"`
	Community	string	`json:"COMMUNITY"`
}