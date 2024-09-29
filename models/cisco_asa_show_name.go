package models

type CiscoAsaShowName struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Name	string	`json:"NAME"`
}

var CiscoAsaShowName_Template = `Value IP_ADDRESS (\d+.\d+.\d+.\d+)
Value NAME (\S+)

Start
  ^name\s+${IP_ADDRESS}\s+${NAME}\s* -> Record
`