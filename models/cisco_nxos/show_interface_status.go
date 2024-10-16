package cisco_nxos

type ShowInterfaceStatus struct {
	Port    string `json:"PORT"`
	Name    string `json:"NAME"`
	Status  string `json:"STATUS"`
	Vlan_id string `json:"VLAN_ID"`
	Duplex  string `json:"DUPLEX"`
	Speed   string `json:"SPEED"`
	Type    string `json:"TYPE"`
}

var ShowInterfaceStatus_Template string = `Value PORT (\S+)
Value NAME (.*?)
Value STATUS (\S+)
Value VLAN_ID (\d+|routed|trunk|f-path|--)
Value DUPLEX (\S+)
Value SPEED (\S+)
Value TYPE (\S+(\s\S+)*)

Start
  ^${PORT}\s+${NAME}\s+${STATUS}\s+${VLAN_ID}\s+${DUPLEX}\s+${SPEED}(?:\s+${TYPE})?\s*$$ -> Record
  ^[Pp]ort\s+[Nn]ame\s+[Ss]tatus\s+[Vv]lan\s+[Dd]uplex\s+[Ss]peed\s+[Tt]ype\s*$$
  ^-+
  ^\s*$$
  ^. -> Error
`
