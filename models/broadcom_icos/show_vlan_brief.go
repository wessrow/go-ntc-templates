package broadcom_icos

type ShowVlanBrief struct {
	Vlan_id   string `json:"VLAN_ID"`
	Vlan_name string `json:"VLAN_NAME"`
	Type      string `json:"TYPE"`
}

var ShowVlanBrief_Template string = `Value VLAN_ID (\d+)
Value VLAN_NAME (\S+)
Value TYPE (\S+)


Start
  # Captures show vlan brief for:
  # Accton AS4610-54P, Accton AS5610-52X, Quanta LY2R, Quanta LB9, DNI AG3448P-R   
  # Raw data is the same in the case of all those devices
  ^\s*VLAN\s+ID\s+VLAN\s+Name\s+VLAN\s+Type$$
  ^-+
  ^\s*${VLAN_ID}\s+${VLAN_NAME}\s+${TYPE} -> Record
  ^\s*$$
  ^. -> Error`
