package edgecore

type ShowInterfacesBrief struct {
	Name        string `json:"NAME"`
	Trunk       string `json:"TRUNK"`
	Interface   string `json:"INTERFACE"`
	Status      string `json:"STATUS"`
	Pvid        string `json:"PVID"`
	Priority    string `json:"PRIORITY"`
	Negotiation string `json:"NEGOTIATION"`
	Speed       string `json:"SPEED"`
	Duplex      string `json:"DUPLEX"`
	Type        string `json:"TYPE"`
}

var ShowInterfacesBrief_Template string = `Value INTERFACE (\S+\s+\d+\s*/\s*\d+)
Value NAME (.*?)
Value STATUS (Up|Down|Disabled)
Value PVID (\d+)
Value PRIORITY (\d+)
Value NEGOTIATION (Auto)
Value SPEED (\d+)
Value DUPLEX (half|full)
Value TYPE (.*?)
Value TRUNK (\S+)

Start
  ^\s*Interface\s+Name\s+Status\s+PVID\s+Pri\s+Speed/Duplex\s+Type\s+Trunk\s*$$ -> IfacesTable
  ^\s*$$
  ^. -> Error

IfacesTable
  ^(?:\s*-+)+\s*$$
  ^\s*${INTERFACE}(?:\s+${NAME})?\s+${STATUS}\s+${PVID}\s+${PRIORITY}\s+(?:${NEGOTIATION})?-?(?:${SPEED}${DUPLEX})?\s+${TYPE}\s+${TRUNK}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
