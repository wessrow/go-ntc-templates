package brocade_fastiron

type ShowArp struct {
	Protocol    string `json:"PROTOCOL"`
	Ip_address  string `json:"IP_ADDRESS"`
	Age         string `json:"AGE"`
	Mac_address string `json:"MAC_ADDRESS"`
	Type        string `json:"TYPE"`
	Port        string `json:"PORT"`
}

var ShowArp_Template string = `Value PROTOCOL (\S+)
Value Required IP_ADDRESS ([A-Fa-f0-9:\.]+)
Value AGE (\S+)
Value Required MAC_ADDRESS ([A-Fa-f0-9\.]{14})
Value TYPE (\S+)
Value Required PORT (\S+)

Start
  ^Protocol\s+Address\s+Age \(min\)\s+Hardware Addr\s+Type\s+Interface
  ^${PROTOCOL}\s+${IP_ADDRESS}\s+${AGE}\s+${MAC_ADDRESS}\s+${TYPE}\s+${PORT} -> Record
  ^\s*$$
  ^. -> Error`
