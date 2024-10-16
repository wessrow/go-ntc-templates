package broadcom_icos

type ShowMacAddrTable struct {
	Vlan_id     string `json:"VLAN_ID"`
	Mac_address string `json:"MAC_ADDRESS"`
	Interface   string `json:"INTERFACE"`
	Index       string `json:"INDEX"`
	Status      string `json:"STATUS"`
}

var ShowMacAddrTable_Template string = `Value VLAN_ID (\S+)
Value MAC_ADDRESS (\S+)
Value INTERFACE (\S+)
Value INDEX (\S+)
Value STATUS (\S+)

Start
  # Captures show mac-address-table for:
  # Accton AS4610-54P, Accton AS5610-52X, Quanta LY2R, Quanta LB9, DNI AG3448P-R   
  # Raw data is the same in the case of all those devices
  ^\s*VLAN\s+ID\s+MAC\s+Address\s+Interface\s+IfIndex\s+Status
  ^-+
  ^\s*${VLAN_ID}\s+${MAC_ADDRESS}\s+${INTERFACE}\s+${INDEX}\s+${STATUS} -> Record
  ^\s*$$
  ^. -> Error`
