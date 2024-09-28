package models

type CiscoNxosShowIpInterfaceVrfAll struct {
	Vrf	string	`json:"VRF"`
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Status	string	`json:"STATUS"`
	Link	string	`json:"LINK"`
	Proto	string	`json:"PROTO"`
	Subnet	string	`json:"SUBNET"`
	Route_pref	string	`json:"ROUTE_PREF"`
	Route_tag	string	`json:"ROUTE_TAG"`
	Broadcast	string	`json:"BROADCAST"`
	Multicast_groups	[]string	`json:"MULTICAST_GROUPS"`
	Mtu	string	`json:"MTU"`
	Proxy_arp	string	`json:"PROXY_ARP"`
	Ip_local_proxy_arp	string	`json:"IP_LOCAL_PROXY_ARP"`
	Multicast_routing	string	`json:"MULTICAST_ROUTING"`
	Icmp_redirects	string	`json:"ICMP_REDIRECTS"`
	Ip_directed_broadcast	string	`json:"IP_DIRECTED_BROADCAST"`
	Ip_forwarding	string	`json:"IP_FORWARDING"`
	Icmp_unreachables	string	`json:"ICMP_UNREACHABLES"`
	Icmp_port_unreachables	string	`json:"ICMP_PORT_UNREACHABLES"`
	Ip_unicast_rpf	string	`json:"IP_UNICAST_RPF"`
	Ip_load_sharing	string	`json:"IP_LOAD_SHARING"`
	Wccp_redirect_outbound	string	`json:"WCCP_REDIRECT_OUTBOUND"`
	Wccp_redirect_inbound	string	`json:"WCCP_REDIRECT_INBOUND"`
	Wccp_redirect_exclude	string	`json:"WCCP_REDIRECT_EXCLUDE"`
	Secondary_ip_address	[]string	`json:"SECONDARY_IP_ADDRESS"`
	Secondary_ip_subnet	[]string	`json:"SECONDARY_IP_SUBNET"`
	Secondary_route_pref	[]string	`json:"SECONDARY_ROUTE_PREF"`
	Secondary_route_tag	[]string	`json:"SECONDARY_ROUTE_TAG"`
	Acl_out	string	`json:"ACL_OUT"`
	Unnumbered_parent_intf	string	`json:"UNNUMBERED_PARENT_INTF"`
	Unnumbered_child_intfs	[]string	`json:"UNNUMBERED_CHILD_INTFS"`
}