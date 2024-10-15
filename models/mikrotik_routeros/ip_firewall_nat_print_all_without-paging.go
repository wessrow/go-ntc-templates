package mikrotik_routeros

type IpFirewallNatPrintAllWithoutPaging struct {
	Dst_address_list   string `json:"DST_ADDRESS_LIST"`
	Flags              string `json:"FLAGS"`
	To_ports           string `json:"TO_PORTS"`
	Protocol           string `json:"PROTOCOL"`
	Src_port           string `json:"SRC_PORT"`
	Dst_address        string `json:"DST_ADDRESS"`
	Dst_port           string `json:"DST_PORT"`
	Log                string `json:"LOG"`
	Index              string `json:"INDEX"`
	Chain              string `json:"CHAIN"`
	Src_address_list   string `json:"SRC_ADDRESS_LIST"`
	Ipsec_policy       string `json:"IPSEC_POLICY"`
	Comment            string `json:"COMMENT"`
	Action             string `json:"ACTION"`
	To_addresses       string `json:"TO_ADDRESSES"`
	Out_interface_list string `json:"OUT_INTERFACE_LIST"`
	Log_prefix         string `json:"LOG_PREFIX"`
	Src_address        string `json:"SRC_ADDRESS"`
	In_interface       string `json:"IN_INTERFACE"`
	Out_interface      string `json:"OUT_INTERFACE"`
	In_interface_list  string `json:"IN_INTERFACE_LIST"`
}

var IpFirewallNatPrintAllWithoutPaging_Template string = `Value Key INDEX (\d+)
Value FLAGS ([XID]+)
Value COMMENT (\S+[\S ]*\S)
Value CHAIN (dstnat|srcnat)
Value ACTION (accept|add-dst-to-address-list|add-src-to-address-list|dst-nat|jump|log|masquerade|netmap|passthrough|redirect|return|same|src-nat)
Value TO_ADDRESSES ([\w!,./:\d]+)
Value TO_PORTS ([\d,-]+)
Value PROTOCOL (!?(dccp|egp|etherip|gre|icmp|idpr-cmtp|ipencap|ipsec-ah|ipv6-encap|ipv6-nonxt|ipv6-route|l2tp|pim|rdp|rsvp|st|udp|vmtp|xns-idp|ddp|encap|ggp|hmp|icmpv6|igmp|ipip|ipsec-esp|ipv6-frag|ipv6-opts|iso-tp4|ospf|pup|rspf|sctp|tcp|udp-lite|vrrp|xtp))
Value SRC_ADDRESS ([\w!.:/\d]+)
Value SRC_ADDRESS_LIST (\S+)
Value SRC_PORT (\S+)
Value IN_INTERFACE (\S+)
Value OUT_INTERFACE (\S+)
Value IN_INTERFACE_LIST (\S+)
Value OUT_INTERFACE_LIST (\S+)
Value IPSEC_POLICY (\S+)
Value DST_ADDRESS ([\w!.:/\d]+)
Value DST_ADDRESS_LIST (\S+)
Value DST_PORT (\S+)
Value LOG (yes|no)
Value LOG_PREFIX (\S+[\S ]*\S)

Start
  ^Flags:.*$$ -> NATTable

NATTable
  ^\s?${INDEX}\s+(${FLAGS})?\s+;;;\s${COMMENT}
  ^\s+chain=${CHAIN}\saction=${ACTION}(\sto-addresses=${TO_ADDRESSES})?(\sto-ports=${TO_PORTS})?(\sprotocol=${PROTOCOL})?(\ssrc-address=${SRC_ADDRESS})?(\sdst-address=${DST_ADDRESS})?(\ssrc-address-list=${SRC_ADDRESS_LIST})?(\sdst-address-list=${DST_ADDRESS_LIST})?(\sin-interface=${IN_INTERFACE})?(\sin-interface-list=${IN_INTERFACE_LIST})?(\sout-interface=${OUT_INTERFACE})?(\sout-interface-list=${OUT_INTERFACE_LIST})?(\ssrc-port=${SRC_PORT})?(\sdst-port=${DST_PORT})?(\slog=${LOG})?(\slog-prefix="(${LOG_PREFIX})?")?(\sipsec-policy=${IPSEC_POLICY})?\s*$$ -> Record
  ^\s?${INDEX}\s+(${FLAGS})?\s+chain=${CHAIN}\saction=${ACTION}(\sto-addresses=${TO_ADDRESSES})?(\sto-ports=${TO_PORTS})?(\sprotocol=${PROTOCOL})?(\ssrc-address=${SRC_ADDRESS})?(\sdst-address=${DST_ADDRESS})?(\ssrc-address-list=${SRC_ADDRESS_LIST})?(\sdst-address-list=${DST_ADDRESS_LIST})?(\sin-interface=${IN_INTERFACE})?(\sin-interface-list=${IN_INTERFACE_LIST})?(\sout-interface=${OUT_INTERFACE})?(\sout-interface-list=${OUT_INTERFACE_LIST})?(\ssrc-port=${SRC_PORT})?(\sdst-port=${DST_PORT})?(\slog=${LOG})?(\slog-prefix="(${LOG_PREFIX})?")?(\sipsec-policy=${IPSEC_POLICY})?\s*$$ -> Record
  ^\s*(?:\d{2}:){2}\d{2}\s+echo:\s*.*$$ -> Next
  ^. -> Error
`
