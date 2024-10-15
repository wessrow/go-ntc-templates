package ruckus_fastiron

type ShowInterfacesBrief struct {
	Trunk       string `json:"TRUNK"`
	Pvid        string `json:"PVID"`
	Priority    string `json:"PRIORITY"`
	Name        string `json:"NAME"`
	Speed       string `json:"SPEED"`
	Link        string `json:"LINK"`
	State       string `json:"STATE"`
	Duplex      string `json:"DUPLEX"`
	Tag         string `json:"TAG"`
	Mac_address string `json:"MAC_ADDRESS"`
	Port        string `json:"PORT"`
}

var ShowInterfacesBrief_Template string = `Value Required PORT (\S+)
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
