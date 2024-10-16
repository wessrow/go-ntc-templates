package cisco_xr

type ShowIpv4Interface struct {
	Protocol          string   `json:"PROTOCOL"`
	Vrf_id            string   `json:"VRF_ID"`
	Interface         string   `json:"INTERFACE"`
	Protocol_status   string   `json:"PROTOCOL_STATUS"`
	Sec_ip_address    []string `json:"SEC_IP_ADDRESS"`
	Mtu               string   `json:"MTU"`
	Link_status       string   `json:"LINK_STATUS"`
	Prefix_length     string   `json:"PREFIX_LENGTH"`
	Sec_prefix_length []string `json:"SEC_PREFIX_LENGTH"`
	Multicast_groups  []string `json:"MULTICAST_GROUPS"`
	Out_acl           string   `json:"OUT_ACL"`
	Rpf_check         string   `json:"RPF_CHECK"`
	Vrf               string   `json:"VRF"`
	Ip_mtu            string   `json:"IP_MTU"`
	In_acl            string   `json:"IN_ACL"`
	Ip_address        string   `json:"IP_ADDRESS"`
}

var ShowIpv4Interface_Template string = `Value Required INTERFACE (\S+)
Value LINK_STATUS (\S+)
Value PROTOCOL (\S+)
Value PROTOCOL_STATUS (\S+)
Value VRF (\S+)
Value VRF_ID (\S+)
Value IP_ADDRESS (\S+?)
Value List SEC_IP_ADDRESS (\S+?)
Value PREFIX_LENGTH (\d*)
Value List SEC_PREFIX_LENGTH (\d*)
Value MTU (\d+)
Value IP_MTU (\d+)
Value List MULTICAST_GROUPS (\d+\.\d+\.\d+\.\d+)
Value OUT_ACL (.*)
Value IN_ACL (.*)
Value RPF_CHECK (\S+)


Start
  ^(Mon?)|(Tue?)|(Wed?)|(Thu?)|(Fri?)|(Sat?)|(Sun?)\s.*$$
  ^\S+\s+is -> Continue.Record
  ^${INTERFACE}\sis\s${LINK_STATUS},\s+${PROTOCOL}\sprotocol\sis\s${PROTOCOL_STATUS}.*$$
  ^\s+Vrf\s+is\s+${VRF}\s+\(vrfid\s+${VRF_ID}\)
  ^\s+Internet\s+address\s+is\s+${IP_ADDRESS}/?${PREFIX_LENGTH}\s*$$
  ^\s+Secondary\s+address\s+${SEC_IP_ADDRESS}/?${SEC_PREFIX_LENGTH}\s*$$
  ^\s+Interface\s+is\s+unnumbered\S*
  ^\s+MTU\s+is\s+${MTU}\s+\(${IP_MTU}\s+is\s+available\s+to\s+IP\)
  ^\s+Helper.*$$
  ^\s+Multicast\s+reserved\s+groups\s+joined\:\s+${MULTICAST_GROUPS} -> Continue
  ^\s+Multicast\s+reserved\s+groups\s+joined\:\s+\S+\s+${MULTICAST_GROUPS} -> Continue
  ^\s+Multicast\s+reserved\s+groups\s+joined\:\s+\S+\s+\S+\s+${MULTICAST_GROUPS} -> Continue
  ^\s+Multicast\s+reserved\s+groups\s+joined\:\s+ -> MULTICAST
  ^\s+Internet\s+protocol.*$$
  ^\s+Directed.*$$
  ^\s+Outgoing\s+access\s+list\s+is\s+${OUT_ACL}
  ^\s+Inbound\s+(?:common)?\s+access\s+list\s+is\s+${IN_ACL}
  ^\s+Proxy.*$$
  ^\s+ICMP redirects.*$$
  ^\s+ICMP unreachables.*$$
  ^\s+ICMP mask.*$$
  ^\s+Table.*$$
  ^\s+mLACP.*$$
  ^IP\s+unicast\s+RPF\s+check\s+is\s+${RPF_CHECK}
  ^RPF\s+mode.*$$
  ^\s*$$
  ^. -> Error

MULTICAST
  ^\s+${MULTICAST_GROUPS}\s*$$ -> Start
  ^\s+(?:\S+\s+){0}${MULTICAST_GROUPS} -> Continue
  ^\s+(?:\S+\s+){1}${MULTICAST_GROUPS} -> Continue
  ^\s+(?:\S+\s+){2}${MULTICAST_GROUPS} -> Continue
  ^\s+(?:\S+\s+){3}${MULTICAST_GROUPS} -> Continue
  ^\s+(?:\S+\s+){4}${MULTICAST_GROUPS} -> Continue
  ^\s+(?:\S+\s+){5}${MULTICAST_GROUPS} -> Continue
  ^\s+(?:\S+\s+){6}${MULTICAST_GROUPS} -> Continue
  ^\s+(?:\S+\s+){7}${MULTICAST_GROUPS} -> Continue
  ^\s+(?:\S+\s+){8}${MULTICAST_GROUPS} -> Continue
  ^\s+(?:\S+\s+){9}${MULTICAST_GROUPS} -> Continue
  ^\s+(?:\S+\s+){10}${MULTICAST_GROUPS} -> Continue
  ^.* -> Start
  ^. -> Error MULTICAST
`
