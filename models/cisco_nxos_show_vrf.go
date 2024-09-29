package models

type CiscoNxosShowVrf struct {
	Name	string	`json:"NAME"`
	Id	string	`json:"ID"`
	State	string	`json:"STATE"`
	Reason	string	`json:"REASON"`
}

var CiscoNxosShowVrf_Template = `Value Required NAME (\S+)
Value Required ID (\S+)
Value Required STATE (\S+)
Value Required REASON (\S+)

Start
  ^VRF-Name\s+VRF-ID\s+State\s+Reason -> Start_record

Start_record
  ^${NAME}\s+${ID}\s+${STATE}\s+${REASON} -> Record
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"

`