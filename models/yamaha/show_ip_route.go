package yamaha

type ShowIpRoute struct {
	Nexthop_if      string `json:"NEXTHOP_IF"`
	Protocol        string `json:"PROTOCOL"`
	Additional_info string `json:"ADDITIONAL_INFO"`
	Network         string `json:"NETWORK"`
	Prefix_length   string `json:"PREFIX_LENGTH"`
	Nexthop_ip      string `json:"NEXTHOP_IP"`
}

var ShowIpRoute_Template string = `# You need to set "console character ascci".
Value NETWORK (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}|default)
Value PREFIX_LENGTH (\d{1,2})
Value NEXTHOP_IP (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}|-)
Value NEXTHOP_IF (\S+)
Value PROTOCOL (\S+)
Value ADDITIONAL_INFO (\S*)

Start
  ^Destination\s+Gateway\s+Interface\s+Kind\s+Additional\s+Info -> Routes

Routes
  ^${NETWORK}(/${PREFIX_LENGTH})?\s+${NEXTHOP_IP}\s+${NEXTHOP_IF}\s+${PROTOCOL}\s+${ADDITIONAL_INFO} -> Record
  ^\s*$$
  ^. -> Error

EOF
`
