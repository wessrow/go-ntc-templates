package juniper_screenos

type GetRoute struct {
	Vr            string `json:"VR"`
	Prefix_length string `json:"PREFIX_LENGTH"`
	Nexthop       string `json:"NEXTHOP"`
	Pref          string `json:"PREF"`
	Vsys          string `json:"VSYS"`
	Best          string `json:"BEST"`
	Id            string `json:"ID"`
	Ip_address    string `json:"IP_ADDRESS"`
	Outinterface  string `json:"OUTINTERFACE"`
	Protocol      string `json:"PROTOCOL"`
	Metric        string `json:"METRIC"`
}

var GetRoute_Template string = `Value Filldown VR (\S+)
Value BEST (\S+)
Value ID (\d+)
Value Required IP_ADDRESS (\d+(\.\d+){3})
Value Required PREFIX_LENGTH (\d{1,2})
Value OUTINTERFACE (\S+)
Value NEXTHOP (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value PROTOCOL (\S+)
Value PREF (\d+)
Value METRIC (\d+)
Value VSYS (\S+)

Start
  # Match the VR
  ^IPv\d\s+Dest-Routes\s+for\s+<${VR}>
  # Match route line
  ^${BEST}\s+${ID}\s+${IP_ADDRESS}\/${PREFIX_LENGTH}\s+${OUTINTERFACE}\s+${NEXTHOP}\s+${PROTOCOL}\s+${PREF}\s+${METRIC}\s+${VSYS} -> Record
  # Match table separator
  ^-+\s*$$
  # Match route type keys
  ^\S+:\s+
  # Match lines that are just space
  ^\s+$$
  # Match table header
  ^\s+ID\s+IP-Prefix\s+Interface\s+Gateway\s+P\s+Pref\s+Mtr\s+Vsys\s*$$
  # Error out if line does not match any expressions
  ^.+ -> Error "Line not found"

EOF
`
