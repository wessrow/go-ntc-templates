package cisco_nxos 

type ShowNvePeers struct {
	Interface	string	`json:"INTERFACE"`
	Peer	string	`json:"PEER"`
	State	string	`json:"STATE"`
	Type	string	`json:"TYPE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
}

var ShowNvePeers_Template = `Value INTERFACE (\S+)
Value PEER (\d+.\d+.\d+.\d+)
Value STATE (Up|Down)
Value TYPE (\S+)
Value MAC_ADDRESS (.*)

Start
  ^Interface\s+Peer-IP\s+State\s+LearnType\s+Uptime\s+Router-Mac\s*$$
  ^-+
  ^${INTERFACE}\s+${PEER}\s+${STATE}\s+${TYPE}\s+\S+\s+${MAC_ADDRESS}\s+ -> Record
  ^\s*$$
  ^. -> Error
`