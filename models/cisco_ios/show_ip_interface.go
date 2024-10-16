package cisco_ios

type ShowIpInterface struct {
	Link_status     string   `json:"LINK_STATUS"`
	Ip_address      []string `json:"IP_ADDRESS"`
	Vrf             string   `json:"VRF"`
	Ip_helper       []string `json:"IP_HELPER"`
	Inbound_acl     string   `json:"INBOUND_ACL"`
	Interface       string   `json:"INTERFACE"`
	Protocol_status string   `json:"PROTOCOL_STATUS"`
	Prefix_length   []string `json:"PREFIX_LENGTH"`
	Mtu             string   `json:"MTU"`
	Outgoing_acl    string   `json:"OUTGOING_ACL"`
}

var ShowIpInterface_Template string = `Value Required INTERFACE (\S+)
Value LINK_STATUS (.+?)
Value PROTOCOL_STATUS (.+?)
Value List IP_ADDRESS (\S+?)
Value List PREFIX_LENGTH (\d*)
Value VRF (\S+)
Value MTU (\d+)
Value List IP_HELPER (\d+\.\d+\.\d+\.\d+)
Value OUTGOING_ACL (.*?)
Value INBOUND_ACL (.*?)


Start
  ^\S -> Continue.Record
  ^${INTERFACE}\s+is\s+${LINK_STATUS},\s+line\s+protocol\s+is\s+${PROTOCOL_STATUS}\s*$$
  ^\s+Internet\s+address\s+is\s+${IP_ADDRESS}/?${PREFIX_LENGTH}\s*$$
  ^\s+Internet\s+address\s+will\s+be\s+negotiated
  ^\s+Secondary\s+address\s+${IP_ADDRESS}/?${PREFIX_LENGTH}\s*$$
  ^.+VPN\s+Routing/Forwarding\s+"${VRF}"
  ^\s+MTU\s+is\s+${MTU}\s+bytes
  ^\s+Helper\s+address(es|)\s(is|are)\s+${IP_HELPER}\s*$$ -> HELPERS
  ^\s+Outgoing\s+(?:Common\s+)?access\s+list\s+is\s+not\s+set
  ^\s+Outgoing\s+(?:Common\s+)?access\s+list\s+is\s+${OUTGOING_ACL}\s*$$
  ^\s+Inbound\s+(?:Common\s+)?access\s+list\s+is\s+not\s+set
  ^\s+Inbound\s+(?:Common\s+)?access\s+list\s+is\s+${INBOUND_ACL}\s*$$
  ^\s+Internet\s+protocol
  ^\s+Interface\s+is\s+unnumbered
  ^\s+Peer\s+address
  ^\s+Dialer
  ^\s+Broadcast
  ^\s+Multicast
  ^\s+2[2-5]\d\.
  ^\s+Address\s+determined
  ^\s+Peer
  ^\s+Directed
  ^\s+MTU
  ^\s+Helper
  ^\s+Directed
  ^\s+Outgoing
  ^\s+Inbound
  ^\s+.*Proxy
  ^\s+Security
  ^\s+Split
  ^\s+ICMP
  ^\s+IP\s+(?:fast|Fast|Flow|Normal|CEF|Null|multicast|route|output|access|Clear)
  ^\s+Downstream
  ^\s+Associated
  ^\s+Topology
  ^\s+Router
  ^\s+TCP
  ^\s+RTP
  ^\s+Probe
  ^\s+Policy
  ^\s+Network\s+address\s+
  ^\s+BGP
  ^\s+Sampled\s+Netflow
  ^\s+IP\s+(Routed|Bridged)\s+Flow
  ^\s+(Input|Output|Post)\s+.*features
  ^\s+(IPv4\s+)?WCCP
  ^\s+IP\s+verify\s+source\s+reachable
  ^\s+\d+\s+(suppressed\s+)?verification\s+drop(s)?
  ^\s*IP\s+Distributed\s+switching\s+is
  ^\s*$$
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^. -> Error

HELPERS
  ^\s+${IP_HELPER}\s*$$
  ^\s+Directed -> Start
  ^.* -> Error`
