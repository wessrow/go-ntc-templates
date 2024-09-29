package models

type CiscoXrShowRedundancySummary struct {
	Primary	string	`json:"PRIMARY"`
	Backup	string	`json:"BACKUP"`
	Status	string	`json:"STATUS"`
	Nsr	string	`json:"NSR"`
}

var CiscoXrShowRedundancySummary_Template = `Value PRIMARY (\S+)
Value BACKUP (\S+)
Value STATUS ((\S+(\s\S+)*))
Value NSR (\S+)

Start
  ^.*\-+ -> Parse

Parse
  ^\s+${PRIMARY}\s+${BACKUP}\s+\(${STATUS},\s+NSR:\s+${NSR}\) -> Record
`