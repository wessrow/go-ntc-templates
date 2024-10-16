package checkpoint_gaia

type ShowRoute struct {
	Network   string `json:"NETWORK"`
	Mask      string `json:"MASK"`
	Nexthopip string `json:"NEXTHOPIP"`
	Interface string `json:"INTERFACE"`
	Comment   string `json:"COMMENT"`
	Protocol  string `json:"PROTOCOL"`
}

var ShowRoute_Template string = `Value PROTOCOL (C|S|R|B|O|A|K|H|P|U|i)
Value NETWORK ([0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3})
Value MASK (\d{1,2})
Value NEXTHOPIP ([0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3})
Value INTERFACE (\S+)
Value COMMENT (\S.*)


Start
  ^\S\s -> Continue.Record
  # Match regular routes
  ^${PROTOCOL}\s+${NETWORK}/${MASK}\s+via\s${NEXTHOPIP},\s${INTERFACE},
  # Match directly connected networks
  ^${PROTOCOL}\s+${NETWORK}/${MASK}\s+is directly connected,\s${INTERFACE}
  # Match optional comment, ignore trailing whitespace
  ^\s{34}${COMMENT}\s*$$
`
