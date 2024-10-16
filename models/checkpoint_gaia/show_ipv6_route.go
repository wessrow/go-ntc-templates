package checkpoint_gaia

type ShowIpv6Route struct {
	Comment       string `json:"COMMENT"`
	Protocol      string `json:"PROTOCOL"`
	Network       string `json:"NETWORK"`
	Prefix_length string `json:"PREFIX_LENGTH"`
	Nexthopip     string `json:"NEXTHOPIP"`
	Interface     string `json:"INTERFACE"`
}

var ShowIpv6Route_Template string = `Value PROTOCOL (C|S|R|B|O|A|K|H|P|U|i)
Value NETWORK ([0-9a-f:]*)
Value PREFIX_LENGTH (\d{1,3})
Value NEXTHOPIP ([0-9a-f:]*)
Value INTERFACE (\S+)
Value COMMENT (\S.*)


Start
  ^\S\s -> Continue.Record
  # Match regular routes
  ^${PROTOCOL}\s+${NETWORK}/${PREFIX_LENGTH}\s+via\s${NEXTHOPIP},\s${INTERFACE},
  # Match directly connected routes
  ^${PROTOCOL}\s+${NETWORK}/${PREFIX_LENGTH}\s+is directly connected,\s${INTERFACE}
  # Match optional comment, ignore trailing whitespace
  ^\s{39}${COMMENT}\s*$$
`
