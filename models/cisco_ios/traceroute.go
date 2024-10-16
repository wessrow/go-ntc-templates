package cisco_ios

type Traceroute struct {
	Details      string   `json:"DETAILS"`
	Hop_num      string   `json:"HOP_NUM"`
	Ip_address   string   `json:"IP_ADDRESS"`
	Fqdn         string   `json:"FQDN"`
	Rtt_response []string `json:"RTT_RESPONSE"`
}

var Traceroute_Template string = `Value Required,Filldown HOP_NUM (\d+)
Value IP_ADDRESS ((?:\d{1,3}\.){3}\d{1,3})
Value FQDN (\S+)
# Captures RTT (in milliseconds), or faults, like
#   - * (probe timed out)
#   - !H (destination host unreachable)
#   - !N (destination network unreachable)
#   - !P (destination protocol unreachable)
#   - !A (administratively prohibited, aka ACL deny)
#   - !Q (source quench, aka destination too busy)
#   - !I (user interrupted test)
#   - !T (timeout)
#   - ? (unknown packet type)
Value List RTT_RESPONSE (\d+(?=\s+msec)|![A-Z]|\*|\?)
# Some intermediate devices provide additional information, such as
#   - ASN
#   - MPLS Exp Bits / Labels
# This is captured into details field
Value DETAILS (.*)

Start
  ^Type\s+escape
  ^Tracing\s+the\s+route
  ^VRF\s+info: -> Entries
  ^\s*$$
  ^\s*$$
  ^. -> Error

Entries
  ^\s+${HOP_NUM}\s+ -> Continue
  ^.*?${FQDN}\s+\(${IP_ADDRESS}\)\s+(?:\[${DETAILS}\])? -> Continue
  ^.*?${IP_ADDRESS}\s+(?:\[${DETAILS}\])? -> Continue
  ^.*?${RTT_RESPONSE} -> Continue
  ^.*?(?:(?:\d+\s+msec|![A-Z]|\*|\?)\s+){1}${RTT_RESPONSE} -> Continue
  ^.*?(?:(?:\d+\s+msec|![A-Z]|\*|\?)\s+){2}${RTT_RESPONSE} -> Continue
  ^.*?(?:(?:\d+\s+msec|![A-Z]|\*|\?)\s+){3}${RTT_RESPONSE} -> Continue
  ^.*?(?:(?:\d+\s+msec|![A-Z]|\*|\?)\s+){4}${RTT_RESPONSE} -> Continue
  ^.*?(?:(?:\d+\s+msec|![A-Z]|\*|\?)\s+){5}${RTT_RESPONSE} -> Continue
  ^.*?(?:(?:\d+\s+msec|![A-Z]|\*|\?)\s+){6}${RTT_RESPONSE} -> Continue
  ^.*?(?:(?:\d+\s+msec|![A-Z]|\*|\?)\s+){7}${RTT_RESPONSE} -> Continue
  ^.*?(?:(?:\d+\s+msec|![A-Z]|\*|\?)\s+){8}${RTT_RESPONSE} -> Continue
  ^.*?(?:(?:\d+\s+msec|![A-Z]|\*|\?)\s+){9}${RTT_RESPONSE} -> Continue
  ^.* -> Record

EOF
`
