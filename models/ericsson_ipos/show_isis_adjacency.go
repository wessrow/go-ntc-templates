package ericsson_ipos 

type ShowIsisAdjacency struct {
	System_id	string	`json:"SYSTEM_ID"`
	Interface	string	`json:"INTERFACE"`
	L	string	`json:"L"`
	Mt	string	`json:"MT"`
	Stat	string	`json:"STAT"`
	Hold	string	`json:"HOLD"`
	Snpa	string	`json:"SNPA"`
	Uptime	string	`json:"UPTIME"`
}

var ShowIsisAdjacency_Template = `Value SYSTEM_ID ((\d+.\d+.\d+)|\S+)
Value INTERFACE (\S+)
Value L (\S+)
Value MT (\S+)
Value STAT (\S+)
Value HOLD (\d+)
Value SNPA ((\d+.\d+.\d+)|\S+)
Value UPTIME (\S+)

Start
  ^IS-IS\s+Adjacenc\S+\s+for\s+tag\s+\S+:
  ^SystemId\s+Interface\s+L\s+MT\s+Stat\s+Hold\s+SNPA\s+Uptime
  ^${SYSTEM_ID}\s+${INTERFACE}\s+${L}\s+${MT}\s+${STAT}\s+${HOLD}\s+${SNPA}\s+${UPTIME} -> Record
  ^Total\s+IS-IS\s+Adjacenc\S+:\s+\d+
  ^. -> Error
`