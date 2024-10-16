package eltex

type ShowVlan struct {
	Vlan_id       string `json:"VLAN_ID"`
	Vlan_name     string `json:"VLAN_NAME"`
	Type          string `json:"TYPE"`
	Authorization string `json:"AUTHORIZATION"`
	Created_by    string `json:"CREATED_BY"`
}

var ShowVlan_Template string = `Value Required VLAN_ID (\d+)
Value VLAN_NAME (\S+)
Value TYPE (\S+)
Value AUTHORIZATION (\S+)
Value CREATED_BY (\S+)

# Broken. There is no good way to differ ports and names when they are wrapped (for
# 'TableType3' in general). So ports are not parsed at all and 'Name' only partially

Start
  ^\s*Vlan\s+mode:.*$$
  ^\s*Created by:.*$$
  ^\s*Vlan\s+Name\s+Tagged\s+ports\s+Untagged\s+ports\s+Type\s+Authorization\s*$$ -> TableType1
  ^\s*Vlan\s+Name\s+Ports\s+Type\s+Authorization\s*$$ -> TableType2
  ^\s*Vlan\s+Name\s+Tagged\s+Ports\s+UnTagged\s+Ports\s+Created\s+by\s*$$ -> TableType3
  ^\s*$$
  ^. -> Error

TableType1
  ^(?:\s*-+)+\s*$$
  ^\s*${VLAN_ID}\s+(?:${VLAN_NAME}|-)\s+\S+\s+\S+\s+${TYPE}\s+${AUTHORIZATION}\s*$$ -> Record
  # Skip wrapped lines
  ^\s{5}\s*\S+(?:\s+\S+)*\s*$$
  ^\s*$$
  ^. -> Error

TableType2
  ^(?:\s*-+)+\s*$$
  ^\s*${VLAN_ID}\s+${VLAN_NAME}(?:\s+\S+)?\s+${TYPE}\s+${AUTHORIZATION}\s*$$ -> Record
  # Skip wrapped lines
  ^\s{5}\s*\S+(?:\s+\S+)*\s*$$
  ^\s*$$
  ^. -> Error

TableType3
  ^(?:\s*-+)+\s*$$
  ^\s*${VLAN_ID}\s+(?:${VLAN_NAME}|-)(?:\s+\S+)?(?:\s+\S+)?\s+${CREATED_BY}\s*$$ -> Record
  # Skip wrapped lines
  ^\s{5}\s*\S+(?:\s+\S+)*\s*$$
  ^\s*$$
  ^. -> Error
`
