package models

type CiscoAsaShowResourceUsage struct {
	Resource	string	`json:"RESOURCE"`
	Current	string	`json:"CURRENT"`
	Peak	string	`json:"PEAK"`
	Limit	string	`json:"LIMIT"`
	Denied	string	`json:"DENIED"`
	Context	string	`json:"CONTEXT"`
}

var CiscoAsaShowResourceUsage_Template = `Value Required RESOURCE (.+?)
Value CURRENT (\d+)
Value PEAK (\d+)
Value LIMIT (\S+)
Value DENIED (\d+)
Value CONTEXT (.+?)


Start
  ^Resource\s+Current\s+Peak\s+Limit\s+Denied\s+Context\s*$$ -> Start_record
  ^\s*$$
  ^. -> Error

Start_record
  ^${RESOURCE}\s+${CURRENT}\s+${PEAK}\s+${LIMIT}\s+${DENIED}\s+${CONTEXT}\s*$$ -> Record
  ^S\s+=\s+System:\s+Combined\s+context\s+limits
  ^\s*$$
  ^.* -> Error "LINE NOT FOUND"

`