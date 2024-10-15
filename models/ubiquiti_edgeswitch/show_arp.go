package ubiquiti_edgeswitch

type ShowArp struct {
	Ip_address  string `json:"IP_ADDRESS"`
	Mac_address string `json:"MAC_ADDRESS"`
	Interface   string `json:"INTERFACE"`
	Type        string `json:"TYPE"`
	Age         string `json:"AGE"`
}

var ShowArp_Template string = `Value Required IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value MAC_ADDRESS (\S+)
Value INTERFACE (\S+)
Value TYPE (\S+)
Value AGE (.*)

Start
  ^${IP_ADDRESS}\s+${MAC_ADDRESS}\s+${INTERFACE}\s+${TYPE}\s+${AGE}.*$$ -> Record
  ^.+\.
  ^\s+IP\s+Address\s+MAC\s+Address\s+Interface\s+Type\s+Age
  ^-+\s+-+\s+-+\s+-+\s+-+\s*$$
  ^\s*$$
  ^. -> Error
`
