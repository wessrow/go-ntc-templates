package models

type CiscoNxosShowFlogiDatabase struct {
	Interface	string	`json:"INTERFACE"`
	Vsan	string	`json:"VSAN"`
	Fcid	string	`json:"FCID"`
	Port_name	string	`json:"PORT_NAME"`
	Port_device_alias	string	`json:"PORT_DEVICE_ALIAS"`
	Node_name	string	`json:"NODE_NAME"`
}

var CiscoNxosShowFlogiDatabase_Template = `Value INTERFACE (\S+)
Value VSAN (\d+)
Value FCID (\w+)
Value PORT_NAME ([a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2})
Value PORT_DEVICE_ALIAS (\w+)
Value NODE_NAME ([a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2})

Start
  ^\S -> Continue.Record
  ^${INTERFACE}\s+${VSAN}\s+${FCID}\s+${PORT_NAME}\s+${NODE_NAME}
  ^\s+\[${PORT_DEVICE_ALIAS}\]

`