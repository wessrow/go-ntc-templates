package huawei_vrp

type DisplayVlan struct {
	Vlan_type   string `json:"VLAN_TYPE"`
	Mode        string `json:"MODE"`
	Interface   string `json:"INTERFACE"`
	Vlan_name   string `json:"VLAN_NAME"`
	Vlan_status string `json:"VLAN_STATUS"`
	Vlan_id     string `json:"VLAN_ID"`
}

var DisplayVlan_Template string = `Value Filldown VLAN_ID (\d+)
Value Filldown VLAN_TYPE (\w+)
Value Filldown MODE ((TG|UT))
Value INTERFACE ([A-Z0-9/th\-runk]+)
Value VLAN_NAME (.+)
Value VLAN_STATUS (\S+)

Start
  ^The\s+total\s+number\s+of\s+VLANs
  ^-+$$
  ^U:\s+Up;\s+D:\s+Down;\s+TG:\s+Tagged;\s+UT:\s+Untagged;
  ^MP:\s+Vlan-mapping;\s+ST:\s+Vlan-stacking;
  ^#:\s+ProtocolTransparent-vlan;\s+\*:\s+Management-vlan;
  ^VID\s+Type\s+Ports -> VLANS
  ^VID\s+Status\s+Property -> STATUS
  ^. -> Error

VLANS
  ^-+$$
  ^${VLAN_ID}\s+${VLAN_TYPE}\s*$$ -> Record
  ^${VLAN_ID}\s+${VLAN_TYPE}\s+${MODE}:${INTERFACE}\(\w+\)\s* -> Continue.Record
  ^${VLAN_ID}\s+${VLAN_TYPE}\s+${MODE}:(?:\S+\s+){1}\s+${INTERFACE}\(\w+\)\s* -> Continue.Record
  ^${VLAN_ID}\s+${VLAN_TYPE}\s+${MODE}:(?:\S+\s+){2}\s+${INTERFACE}\(\w+\)\s* -> Continue.Record
  ^${VLAN_ID}\s+${VLAN_TYPE}\s+${MODE}:(?:\S+\s+){3}\s+${INTERFACE}\(\w+\)\s* -> Continue.Record
  ^${VLAN_ID}\s+${VLAN_TYPE}\s+${MODE}:(?:\S+\s+){4}\s+${INTERFACE}\(\w+\)\s* -> Continue.Record
  ^\d+
  ^\s{6,}${MODE}:${INTERFACE}\(\w+\)\s* -> Continue.Record
  ^\s{6,}${INTERFACE}\(\w+\)\s* -> Continue.Record
  ^\s{6,}(?:\S+\s+){1}${INTERFACE}\(\w+\)\s* -> Continue.Record
  ^\s{6,}(?:\S+\s+){2}${INTERFACE}\(\w+\)\s* -> Continue.Record
  ^\s{6,}(?:\S+\s+){3}${INTERFACE}\(\w+\)\s* -> Continue.Record
  ^\s{6,}(?:\S+\s+){4}${INTERFACE}\(\w+\)\s* -> Continue.Record
  ^\s{6,}
  ^\s*$$ -> Start
  ^. -> Error


STATUS
  ^-+$$ -> Clearall
  ^${VLAN_ID}\s+${VLAN_STATUS}\s+\S+\s+\S+\s+\S+\s+${VLAN_NAME} -> Continue.Record
  ^\d+ -> Clearall
  ^. -> Error
`
