package paloalto_panos

type ShowInterfaceHardware struct {
	Mac_address string `json:"MAC_ADDRESS"`
	Interface   string `json:"INTERFACE"`
	Id          string `json:"ID"`
	Speed       string `json:"SPEED"`
	Duplex      string `json:"DUPLEX"`
	State       string `json:"STATE"`
}

var ShowInterfaceHardware_Template string = `Value INTERFACE (\S+)
Value ID (\S+)
Value SPEED (\[n/a\]|\S+)
Value DUPLEX (\[n/a\]|\S+)
Value STATE (\S+)
Value MAC_ADDRESS ([a-fA-F0-9]{2}(\:[a-fA-F0-9]{2}){5})

Start
  ^${INTERFACE}\s+${ID}\s+${SPEED}/${DUPLEX}/${STATE}\s+${MAC_ADDRESS} -> Record
`
