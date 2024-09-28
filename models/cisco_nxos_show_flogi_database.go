package models

type CiscoNxosShowFlogiDatabase struct {
	Interface	string	`json:"INTERFACE"`
	Vsan	string	`json:"VSAN"`
	Fcid	string	`json:"FCID"`
	Port_name	string	`json:"PORT_NAME"`
	Port_device_alias	string	`json:"PORT_DEVICE_ALIAS"`
	Node_name	string	`json:"NODE_NAME"`
}