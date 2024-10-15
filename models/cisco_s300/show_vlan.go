package cisco_s300

type ShowVlan struct {
	Interfaces    string `json:"INTERFACES"`
	Type          string `json:"TYPE"`
	Authorization string `json:"AUTHORIZATION"`
	Created_by    string `json:"CREATED_BY"`
	Vlan_id       string `json:"VLAN_ID"`
	Vlan_name     string `json:"VLAN_NAME"`
}

var ShowVlan_Template string = `Value Required VLAN_ID (\d+)
Value VLAN_NAME (.*?)
Value INTERFACES ((?:fa|gi|te|Po)\d+(?:-\d+)?(?:,(?:fa|gi|te|Po)\d+(?:-\d+)?)*,?)
Value TYPE (\S+)
Value AUTHORIZATION (\S+)
Value CREATED_BY (\S+)

# Broken. Does not support wrapped lines. There is no good way to differ ports
# and names on the wrapped lines

Start
  ^\s*Vlan\s+Name\s+Ports\s+Type\s+Authorization\s*$$ -> Column5
  ^\s*Created\s+by:\s+D-Default,\s+S-Static,\s+G-GVRP,\s+R-Radius\s+Assigned\s+VLAN\s*$$
  ^\s*Vlan\s+Name\s+Ports\s+Created\s+by\s*$$ -> Column4
  ^\s*$$
  ^. -> Error

Column5
  ^(\s*-*)*\s*$$
  ^\s*\d+.* -> Continue.Record
  ^\s*${VLAN_ID}\s+${VLAN_NAME}(?:\s+${INTERFACES})?\s+${TYPE}\s+${AUTHORIZATION}\s*$$

Column4
  ^(\s*-*)*\s*$$
  ^\s*\d+.* -> Continue.Record
  ^\s*${VLAN_ID}\s+${VLAN_NAME}(?:\s+${INTERFACES})?\s+${CREATED_BY}\s*$$
`
