package models

type CiscoIosShowNvePeers struct {
	Interface	string	`json:"INTERFACE"`
	Vni	string	`json:"VNI"`
	Type	string	`json:"TYPE"`
	Peer	string	`json:"PEER"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	State	string	`json:"STATE"`
}

var CiscoIosShowNvePeers_Template = `Value INTERFACE (\S+)
Value VNI (\d+)
Value TYPE (L2CP|L3CP)
Value PEER (\d+.\d+.\d+.\d+)
Value MAC_ADDRESS ([0-9a-fA-F]{4}\.[0-9a-fA-F]{4}\.[0-9a-fA-F]{4})
Value STATE (UP|DOWN)

Start
  ^Interface\s+VNI\s+Type\s+Peer-IP\s+RMAC/Num_RTs\s+eVNI\s+state\s+flags\s+UP\s+time\s*$$
  ^${INTERFACE}\s+${VNI}\s+${TYPE}\s+${PEER}\s+${MAC_ADDRESS}\s+\d+\s+${STATE} -> Record
  ^\s*$$
  ^. -> Error
  
`