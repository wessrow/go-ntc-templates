package arista_eos 

type ShowMacSecurityInterface struct {
	Interface	string	`json:"INTERFACE"`
	Sci	string	`json:"SCI"`
	Controlled_port	string	`json:"CONTROLLED_PORT"`
	Key_in_use	string	`json:"KEY_IN_USE"`
}

var ShowMacSecurityInterface_Template = `Value INTERFACE (\S+)
Value SCI (\S+)
Value CONTROLLED_PORT (\S+)
Value KEY_IN_USE (\S+)

Start
  ^Interface.*Use -> Data

Data
  ^${INTERFACE}\s+${SCI}\s+${CONTROLLED_PORT}\s+${KEY_IN_USE} -> Record
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
  
EOF  

`