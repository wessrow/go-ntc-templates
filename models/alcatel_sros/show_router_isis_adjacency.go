package alcatel_sros 

type ShowRouterIsisAdjacency struct {
	System_id	string	`json:"SYSTEM_ID"`
	Usage	string	`json:"USAGE"`
	State	string	`json:"STATE"`
	Hold	string	`json:"HOLD"`
	Interface	string	`json:"INTERFACE"`
	Mt_id	string	`json:"MT_ID"`
}

var ShowRouterIsisAdjacency_Template = `Value SYSTEM_ID (\S+)
Value USAGE (\S+)
Value STATE (Up|Down)
Value HOLD (\d+)
Value INTERFACE (\S+)
Value MT_ID (\S+)

Start
  ^----------- -> Adjacency

Adjacency
  ^${SYSTEM_ID}\s+${USAGE}\s+${STATE}\s+${HOLD}\s+${INTERFACE}\s+${MT_ID} -> Record
  ^\s*$$
  ^-----------
  ^===========
  ^Adjacencies
  ^. -> Error

`