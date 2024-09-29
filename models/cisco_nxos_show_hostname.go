package models

type CiscoNxosShowHostname struct {
	Hostname	string	`json:"HOSTNAME"`
}

var CiscoNxosShowHostname_Template = `Value HOSTNAME (.+?)

Start
  ^${HOSTNAME}\s*$$ -> Record
`