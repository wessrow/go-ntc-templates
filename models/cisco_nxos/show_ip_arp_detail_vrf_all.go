package cisco_nxos

type ShowIpArpDetailVrfAll struct {
	Interface          string `json:"INTERFACE"`
	Physical_interface string `json:"PHYSICAL_INTERFACE"`
	Flags              string `json:"FLAGS"`
	Vrf                string `json:"VRF"`
	Ip_address         string `json:"IP_ADDRESS"`
	Age                string `json:"AGE"`
	Mac_address        string `json:"MAC_ADDRESS"`
}

var ShowIpArpDetailVrfAll_Template string = `Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value AGE (\S+)
Value MAC_ADDRESS (\S+)
Value INTERFACE (\S+)
Value PHYSICAL_INTERFACE (\S+)
Value FLAGS (\*|\+|#|CP|PS|RO)
Value VRF (\S*)

Start
  #Ignore junk
  ^.+\s-\sAdjacencies\s
  ^.+\sCP\s-\s
  ^.+\sPS\s-\s
  ^.+\sRO\s-\s
  ^\s*IP\sARP\sTable\s
  ^\s*Total\snumber\sof\sentries:
  ^\s*Address\s+Age\s+MAC\s+Address\s+Interface\s+Physical\sInterface\s+Flags\s*(VRF\sName)*\s*$$ -> Data
  ^. -> Error

Data
  ^${IP_ADDRESS}\s+${AGE}\s+${MAC_ADDRESS}\s+${INTERFACE}\s+${PHYSICAL_INTERFACE}\s*${FLAGS}?\s*${VRF}\s*$$ -> Record
  ^. -> Error
`
