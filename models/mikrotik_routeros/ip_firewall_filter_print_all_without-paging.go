package mikrotik_routeros

type IpFirewallFilterPrintAllWithoutPaging struct {
	In_interface         string `json:"IN_INTERFACE"`
	Dst_port             string `json:"DST_PORT"`
	Flags                string `json:"FLAGS"`
	Connection_state     string `json:"CONNECTION_STATE"`
	Src_address_list     string `json:"SRC_ADDRESS_LIST"`
	In_interface_list    string `json:"IN_INTERFACE_LIST"`
	Out_interface        string `json:"OUT_INTERFACE"`
	Ipsec_policy         string `json:"IPSEC_POLICY"`
	Dst_address_list     string `json:"DST_ADDRESS_LIST"`
	Comment              string `json:"COMMENT"`
	Hw_offload           string `json:"HW_OFFLOAD"`
	Connection_nat_state string `json:"CONNECTION_NAT_STATE"`
	Src_port             string `json:"SRC_PORT"`
	Out_interface_list   string `json:"OUT_INTERFACE_LIST"`
	Dst_address          string `json:"DST_ADDRESS"`
	Log_prefix           string `json:"LOG_PREFIX"`
	Index                string `json:"INDEX"`
	Chain                string `json:"CHAIN"`
	Action               string `json:"ACTION"`
	Protocol             string `json:"PROTOCOL"`
	Src_address          string `json:"SRC_ADDRESS"`
	Src_mac_address      string `json:"SRC_MAC_ADDRESS"`
	Log                  string `json:"LOG"`
}

var IpFirewallFilterPrintAllWithoutPaging_Template string = `Value Key INDEX (\d+)
Value FLAGS ([XID]+)
Value COMMENT (\S+[\S ]*\S)
Value CHAIN (input|forward|output)
Value ACTION (accept|add-dst-to-address-list|add-src-to-address-list|drop|fasttrack-connection|jump|log|passthrough|reject|return|tarpit)
Value HW_OFFLOAD (yes|no)
Value CONNECTION_STATE (\S+)
Value CONNECTION_NAT_STATE (\S+)
Value PROTOCOL (!?(dccp|egp|etherip|gre|icmp|idpr-cmtp|ipencap|ipsec-ah|ipv6-encap|ipv6-nonxt|ipv6-route|l2tp|pim|rdp|rsvp|st|udp|vmtp|xns-idp|ddp|encap|ggp|hmp|icmpv6|igmp|ipip|ipsec-esp|ipv6-frag|ipv6-opts|iso-tp4|ospf|pup|rspf|sctp|tcp|udp-lite|vrrp|xtp))
Value SRC_ADDRESS ([\w!.:/\d]+)
Value SRC_ADDRESS_LIST (\S+)
Value SRC_PORT (\S+)
Value IN_INTERFACE_LIST (\S+)
Value OUT_INTERFACE_LIST (\S+)
Value IN_INTERFACE (\S+)
Value OUT_INTERFACE (\S+)
Value IPSEC_POLICY (\S+)
Value DST_ADDRESS ([\w!.:/\d]+)
Value DST_ADDRESS_LIST (\S+)
Value DST_PORT (\S+)
Value SRC_MAC_ADDRESS (\S+)
Value LOG (yes|no)
Value LOG_PREFIX ((\S+[\S ]+)?)

Start
  ^Flags:.*$$ -> FirewallTable

FirewallTable
  ^\s?${INDEX}\s+(${FLAGS})?\s+;;;\s${COMMENT}
  ^\s+chain=${CHAIN}(\saction=${ACTION})?(\shw-offload=${HW_OFFLOAD})?(\sconnection-state=${CONNECTION_STATE})?(\sconnection-nat-state=${CONNECTION_NAT_STATE})?\s*(\sprotocol=${PROTOCOL})?(\ssrc-address=${SRC_ADDRESS})?(\sdst-address=${DST_ADDRESS})?(\ssrc-address-list=${SRC_ADDRESS_LIST})?(\sdst-address-list=${DST_ADDRESS_LIST})?(\sin-interface-list=${IN_INTERFACE_LIST})?(\sout-interface-list=${OUT_INTERFACE_LIST})?(\ssrc-port=${SRC_PORT})?(\sdst-port=${DST_PORT})?(\ssrc-mac-address=${SRC_MAC_ADDRESS})?(\sin-interface=${IN_INTERFACE})?(\sout-interface=${OUT_INTERFACE})?(\sipsec-policy=${IPSEC_POLICY})?(\slog=${LOG})?(\slog-prefix="${LOG_PREFIX}")?\s*$$ -> Record
  ^\s?${INDEX}\s+(${FLAGS})?\s+chain=${CHAIN}(\saction=${ACTION})?(\shw-offload=${HW_OFFLOAD})?(\sconnection-state=${CONNECTION_STATE})?(\sconnection-nat-state=${CONNECTION_NAT_STATE})?(\sprotocol=${PROTOCOL})?(\ssrc-address=${SRC_ADDRESS})?(\sdst-address=${DST_ADDRESS})?(\ssrc-address-list=${SRC_ADDRESS_LIST})?(\sdst-address-list=${DST_ADDRESS_LIST})?(\sin-interface-list=${IN_INTERFACE_LIST})?(\sout-interface-list=${OUT_INTERFACE_LIST})?(\ssrc-port=${SRC_PORT})?(\sdst-port=${DST_PORT})?(\ssrc-mac-address=${SRC_MAC_ADDRESS})?(\sin-interface=${IN_INTERFACE})?(\sout-interface=${OUT_INTERFACE})?(\sipsec-policy=${IPSEC_POLICY})?(\slog=${LOG})?(\slog-prefix="${LOG_PREFIX}")?\s*$$ -> Record
  ^\s*(?:\d{2}:){2}\d{2}\s+echo:\s*.*$$ -> Next
  ^. -> Error
`
