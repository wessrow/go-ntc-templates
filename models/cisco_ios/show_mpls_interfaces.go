package cisco_ios

type ShowMplsInterfaces struct {
	Protocol    string `json:"PROTOCOL"`
	Tunnel      string `json:"TUNNEL"`
	Bgp         string `json:"BGP"`
	Static      string `json:"STATIC"`
	Operational string `json:"OPERATIONAL"`
	Interface   string `json:"INTERFACE"`
	Mpls_ip     string `json:"MPLS_IP"`
}

var ShowMplsInterfaces_Template string = `Value Required INTERFACE (\S+)
Value MPLS_IP (\w+)
Value PROTOCOL (\w+)
Value TUNNEL (\w+)
Value BGP (\w+)
Value STATIC (\w+)
Value OPERATIONAL (\w+)

Start
  ^Interface\s+IP\s+Tunnel\s+BGP\s+Static\s+Operational$$
  ^${INTERFACE}\s+${MPLS_IP}\s+\(${PROTOCOL}\)\s+${TUNNEL}\s+${BGP}\s+${STATIC}\s+${OPERATIONAL}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
