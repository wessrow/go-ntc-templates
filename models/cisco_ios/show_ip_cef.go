package cisco_ios

type ShowIpCef struct {
	Ip_address    string   `json:"IP_ADDRESS"`
	Prefix_length string   `json:"PREFIX_LENGTH"`
	Nexthop       []string `json:"NEXTHOP"`
	Interface     []string `json:"INTERFACE"`
}

var ShowIpCef_Template string = `#
# This template is only compatible with the following command
#   - show ip cef [vrf NAME]
# This is enforced by INDEX file
# Check cisco_ios_show_ip_cef_detail.textfsm for more details
#
Value Required IP_ADDRESS ((?:\d{1,3}\.){3}\d{1,3})
Value Required PREFIX_LENGTH (\d{1,2})
# To support ECMP, NEXTHOP and INTERFACE are defined as lists
# When no ECMP is available, these will be single item lists
# However, if there are multiple paths available
#     then expect to see one item per route/path
Value List NEXTHOP ((?:no\s)?(?:\S+))
Value List INTERFACE ([A-Za-z][A-Za-z0-9\.\/-]+)

Start
  # >>> Parse EXCEPTIONS
  # CEF is not enabled
  ^%IPv4\s+CEF\s+not\s+running$$ -> End
  # Invalid prefix
  ^\s+Invalid\s+prefix/mask -> End
  #
  # >>> Parse HEADING
  ^Prefix\s+Next\s+Hop\s+Interface$$ -> Entries
  #
  # >>> Parse SPECIAL
  ^\s*$$
  ^. -> Error

Entries
  # >>> Parse CEF ENTRIES
  # Entry detected
  ^(?:\d{1,3}\.){3}\d{1,3}\/\d{1,2} -> Continue.Record
  # Parse prefix/nexthop
  ^${IP_ADDRESS}\/${PREFIX_LENGTH}\s+${NEXTHOP}\s*$$
  # Parse prefix, nexthop and interface
  ^${IP_ADDRESS}\/${PREFIX_LENGTH}\s+${NEXTHOP}\s+${INTERFACE}$$
  # Parse nexthop and interface for ECMP prefixes
  ^\s+${NEXTHOP}\s+${INTERFACE}$$
  #
  # >>> Parse SPECIAL
  ^\s*$$
  ^. -> Error
`
