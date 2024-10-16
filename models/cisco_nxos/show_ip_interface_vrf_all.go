package cisco_nxos

type ShowIpInterfaceVrfAll struct {
	Ip_directed_broadcast  string   `json:"IP_DIRECTED_BROADCAST"`
	Secondary_route_tag    []string `json:"SECONDARY_ROUTE_TAG"`
	Interface              string   `json:"INTERFACE"`
	Mtu                    string   `json:"MTU"`
	Wccp_redirect_exclude  string   `json:"WCCP_REDIRECT_EXCLUDE"`
	Secondary_ip_address   []string `json:"SECONDARY_IP_ADDRESS"`
	Acl_out                string   `json:"ACL_OUT"`
	Unnumbered_child_intfs []string `json:"UNNUMBERED_CHILD_INTFS"`
	Status                 string   `json:"STATUS"`
	Secondary_route_pref   []string `json:"SECONDARY_ROUTE_PREF"`
	Multicast_groups       []string `json:"MULTICAST_GROUPS"`
	Subnet                 string   `json:"SUBNET"`
	Ip_address             string   `json:"IP_ADDRESS"`
	Proto                  string   `json:"PROTO"`
	Route_tag              string   `json:"ROUTE_TAG"`
	Multicast_routing      string   `json:"MULTICAST_ROUTING"`
	Icmp_redirects         string   `json:"ICMP_REDIRECTS"`
	Ip_forwarding          string   `json:"IP_FORWARDING"`
	Vrf                    string   `json:"VRF"`
	Route_pref             string   `json:"ROUTE_PREF"`
	Broadcast              string   `json:"BROADCAST"`
	Icmp_unreachables      string   `json:"ICMP_UNREACHABLES"`
	Icmp_port_unreachables string   `json:"ICMP_PORT_UNREACHABLES"`
	Ip_load_sharing        string   `json:"IP_LOAD_SHARING"`
	Wccp_redirect_inbound  string   `json:"WCCP_REDIRECT_INBOUND"`
	Secondary_ip_subnet    []string `json:"SECONDARY_IP_SUBNET"`
	Link                   string   `json:"LINK"`
	Ip_unicast_rpf         string   `json:"IP_UNICAST_RPF"`
	Unnumbered_parent_intf string   `json:"UNNUMBERED_PARENT_INTF"`
	Proxy_arp              string   `json:"PROXY_ARP"`
	Wccp_redirect_outbound string   `json:"WCCP_REDIRECT_OUTBOUND"`
	Ip_local_proxy_arp     string   `json:"IP_LOCAL_PROXY_ARP"`
}

var ShowIpInterfaceVrfAll_Template string = `Value Filldown VRF (\S+)
Value Required INTERFACE (\S+)
Value IP_ADDRESS ([a-zA-Z0-9./]+)
Value STATUS (\S+-\S+)
Value LINK (\S+-\S+)
Value PROTO (\S+-\S+)
Value SUBNET (\S+)
Value ROUTE_PREF (\d+)
Value ROUTE_TAG (\d+)
Value BROADCAST ([a-zA-Z0-9./]+)
Value List MULTICAST_GROUPS (\S+)
Value MTU (\d+)
Value PROXY_ARP (\S+)
Value IP_LOCAL_PROXY_ARP (\S+)
Value MULTICAST_ROUTING (\S+)
Value ICMP_REDIRECTS (\S+)
Value IP_DIRECTED_BROADCAST (\S+)
Value IP_FORWARDING (\S+)
Value ICMP_UNREACHABLES (\S+)
Value ICMP_PORT_UNREACHABLES (\S+)
Value IP_UNICAST_RPF (\S+)
Value IP_LOAD_SHARING (\S+)
Value WCCP_REDIRECT_OUTBOUND (\S+)
Value WCCP_REDIRECT_INBOUND (\S+)
Value WCCP_REDIRECT_EXCLUDE (\S+)
Value List SECONDARY_IP_ADDRESS (\S+)
Value List SECONDARY_IP_SUBNET (\S+)
Value List SECONDARY_ROUTE_PREF (\d+)
Value List SECONDARY_ROUTE_TAG (\d+)
Value ACL_OUT (\S+)
Value UNNUMBERED_PARENT_INTF (\S+)
Value List UNNUMBERED_CHILD_INTFS (\S+)

Start
  ^IP\s+Interface\s+Status\s+for\s+VRF -> Continue.Record
  ^IP\s+Interface\s+Status\s+for\s+VRF\s+"${VRF}"
  ^\S+,\s+Interface\s+status: -> Continue.Record
  ^${INTERFACE},\s+Interface\s+status:\s+${PROTO}\/${LINK}\/${STATUS},\s+iod:\s+\d+,
  ^${INTERFACE},\s+Interface\s+status:\s+${PROTO}/${LINK}/${STATUS}
  ^\s+IP\s+address:\s+${IP_ADDRESS},\s+IP\s+subnet:\s+${SUBNET}\s+route-preference:\s+${ROUTE_PREF},\s+tag:\s+${ROUTE_TAG}
  ^\s+IP\s+address:\s+${SECONDARY_IP_ADDRESS},\s+IP\s+subnet:\s+${SECONDARY_IP_SUBNET}\s+secondary\s+route-preference:\s+${SECONDARY_ROUTE_PREF},\s+tag:\s+${SECONDARY_ROUTE_TAG}
  ^\s+IP\s+address:\s+${IP_ADDRESS},\s+IP\s+subnet:\s+${SUBNET}
  ^\s+IP\s+address: none
  ^\s+IP\s+broadcast\s+address:\s+${BROADCAST}
  ^\s+IP\s+multicast\s+groups\s+locally\s+joined:\s+${MULTICAST_GROUPS}
  ^\s+IP\s+multicast\s+groups\s+locally\s+joined:* -> MULTICAST
  ^\s+IP\s+MTU:\s+${MTU}\s+bytes
  ^\s+IP\s+primary\s+address\s+route-preference:\s+${ROUTE_PREF},\s+tag:\s+${ROUTE_TAG}
  ^\s+IP\s+proxy\s+ARP\s+:\s+${PROXY_ARP}
  ^\s+IP\s+Local\s+Proxy\s+ARP\s+:\s+${IP_LOCAL_PROXY_ARP}
  ^\s+IP\s+multicast\s+routing:\s+${MULTICAST_ROUTING}
  ^\s+IP\s+icmp\s+redirects:\s+${ICMP_REDIRECTS}
  ^\s+IP\s+directed-broadcast:\s+${IP_DIRECTED_BROADCAST}
  ^\s+IP\s+Forwarding:\s+${IP_FORWARDING}
  ^\s+IP\s+icmp\s+unreachables.+:\s+${ICMP_UNREACHABLES}
  ^\s+IP\s+icmp\s+port-unreachable:\s+${ICMP_PORT_UNREACHABLES}
  ^\s+IP\s+unicast\s+reverse\s+path\s+forwarding:\s+${IP_UNICAST_RPF}
  ^\s+IP\s+load\s+sharing:\s+${IP_LOAD_SHARING}
  ^\s+(ip|IP)\s+interface\s+statistics\s+last\s+reset:
  ^\s+IP\s+interface\s+software\s+stats:
  ^\s+Unicast\s+packets\s+:\s+\d+\/\d+\/\d+\/\d+\/\d+
  ^\s+Unicast\s+bytes\s+:\s+\d+\/\d+\/\d+\/\d+\/\d+
  ^\s+Multicast\s+packets\s+:\s+\d+\/\d+\/\d+\/\d+\/\d+
  ^\s+Multicast\s+bytes\s+:\s+\d+\/\d+\/\d+\/\d+\/\d+
  ^\s+Broadcast\s+packets\s+:\s+\d+\/\d+\/\d+\/\d+\/\d+
  ^\s+Broadcast\s+bytes\s+:\s+\d+\/\d+\/\d+\/\d+\/\d+
  ^\s+Labeled\s+packets\s+:\s+\d+\/\d+\/\d+\/\d+\/\d+
  ^\s+Labeled\s+bytes\s+:\s+\d+\/\d+\/\d+\/\d+\/\d+
  ^\s+IP\s+outbound\s+access\s+list:\s+${ACL_OUT}
  ^\s+IP\s+unnumbered\s+interface\s+\(${UNNUMBERED_PARENT_INTF}\)
  ^\s*Unnumbered\s+interfaces\s+of -> UNNUMBERED
  ^\s+WCCP\s+Redirect\s+outbound:\s+${WCCP_REDIRECT_OUTBOUND}
  ^\s+WCCP\s+Redirect\s+inbound:\s+${WCCP_REDIRECT_INBOUND}
  ^\s+WCCP\s+Redirect\s+exclude:\s+${WCCP_REDIRECT_EXCLUDE}
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
  ^. -> Error

UNNUMBERED
  ^\s*${UNNUMBERED_CHILD_INTFS}:\s*$$ -> Continue
  ^(?:\S+\s+){0}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){1}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){2}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){3}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){4}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){5}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){6}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){7}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){8}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){9}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){10}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){11}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){12}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){13}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){14}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){15}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){16}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){17}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){18}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){19}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^(?:\S+\s+){20}${UNNUMBERED_CHILD_INTFS}: -> Continue
  ^.* -> Start
  ^. -> Error
`
