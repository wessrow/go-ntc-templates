package ruckus_fastiron 

type ShowInterfacesBrief struct {
	Port	string	`json:"PORT"`
	Link	string	`json:"LINK"`
	State	string	`json:"STATE"`
	Duplex	string	`json:"DUPLEX"`
	Speed	string	`json:"SPEED"`
	Trunk	string	`json:"TRUNK"`
	Tag	string	`json:"TAG"`
	Pvid	string	`json:"PVID"`
	Priority	string	`json:"PRIORITY"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Name	string	`json:"NAME"`
}

var ShowInterfacesBrief_Template = `Value Required PORT (\S+)
Value LINK (\S+)
Value STATE (\S+)
Value DUPLEX (\S+)
Value SPEED (\S+)
Value TRUNK (\S+)
Value TAG (\S+)
Value PVID (\S+)
Value PRIORITY (\S+)
Value MAC_ADDRESS (([A-Fa-f0-9\.]{14}|None))
Value NAME (\S+)

Start
  ^Port\s+Link\s+State\s+Dupl\s+Speed\s+Trunk\s+Tag\s+Pvid\s+Pri\s+MAC\s+Name
  ^${PORT}\s+${LINK}\s+${STATE}\s+${DUPLEX}\s+${SPEED}\s+${TRUNK}\s+${TAG}\s+${PVID}\s+${PRIORITY}\s+${MAC_ADDRESS}\s+${NAME} -> Record
  ^${PORT}\s+${LINK}\s+${STATE}\s+${DUPLEX}\s+${SPEED}\s+${TRUNK}\s+${TAG}\s+${PVID}\s+${PRIORITY}\s+${MAC_ADDRESS} -> Record
  ^\s*$$
  ^. -> Error `