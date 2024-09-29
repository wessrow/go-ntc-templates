package models

type CiscoNxosShowIpInterface struct {
	Vrf_name	string	`json:"VRF_NAME"`
	Interface	string	`json:"INTERFACE"`
	Protocol_status	string	`json:"PROTOCOL_STATUS"`
	Link_status	string	`json:"LINK_STATUS"`
	Admin_status	string	`json:"ADMIN_STATUS"`
	Primary_ip_address	string	`json:"PRIMARY_IP_ADDRESS"`
	Primary_ip_subnet	string	`json:"PRIMARY_IP_SUBNET"`
	Primary_broadcast_address	string	`json:"PRIMARY_BROADCAST_ADDRESS"`
	Secondary_ip_address	[]string	`json:"SECONDARY_IP_ADDRESS"`
	Secondary_ip_subnet	[]string	`json:"SECONDARY_IP_SUBNET"`
	Secondary_route_preference	[]string	`json:"SECONDARY_ROUTE_PREFERENCE"`
	Secondary_route_preference_tag	[]string	`json:"SECONDARY_ROUTE_PREFERENCE_TAG"`
	Multicast_groups	string	`json:"MULTICAST_GROUPS"`
	Mtu	string	`json:"MTU"`
	Route_preference	string	`json:"ROUTE_PREFERENCE"`
	Route_preference_tag	string	`json:"ROUTE_PREFERENCE_TAG"`
	Proxy_arp	string	`json:"PROXY_ARP"`
	Local_proxy_arp	string	`json:"LOCAL_PROXY_ARP"`
	Multicast_routing	string	`json:"MULTICAST_ROUTING"`
	Icmp_redirects	string	`json:"ICMP_REDIRECTS"`
	Directed_broadcast	string	`json:"DIRECTED_BROADCAST"`
	Ip_forwarding	string	`json:"IP_FORWARDING"`
	Icmp_unreachables	string	`json:"ICMP_UNREACHABLES"`
	Icmp_port_unreachable	string	`json:"ICMP_PORT_UNREACHABLE"`
	Urpf	string	`json:"URPF"`
	Load_sharing	string	`json:"LOAD_SHARING"`
	Last_reset	string	`json:"LAST_RESET"`
	Acl_out	string	`json:"ACL_OUT"`
	Unicast_packets_sent	string	`json:"UNICAST_PACKETS_SENT"`
	Unicast_packets_received	string	`json:"UNICAST_PACKETS_RECEIVED"`
	Unicast_packets_forwarded	string	`json:"UNICAST_PACKETS_FORWARDED"`
	Unicast_packets_originated	string	`json:"UNICAST_PACKETS_ORIGINATED"`
	Unicast_packets_consumed	string	`json:"UNICAST_PACKETS_CONSUMED"`
	Unicast_bytes_sent	string	`json:"UNICAST_BYTES_SENT"`
	Unicast_bytes_received	string	`json:"UNICAST_BYTES_RECEIVED"`
	Unicast_bytes_forwarded	string	`json:"UNICAST_BYTES_FORWARDED"`
	Unicast_bytes_originated	string	`json:"UNICAST_BYTES_ORIGINATED"`
	Unicast_bytes_consumed	string	`json:"UNICAST_BYTES_CONSUMED"`
	Multicast_packets_sent	string	`json:"MULTICAST_PACKETS_SENT"`
	Multicast_packets_received	string	`json:"MULTICAST_PACKETS_RECEIVED"`
	Multicast_packets_forwarded	string	`json:"MULTICAST_PACKETS_FORWARDED"`
	Multicast_packets_originated	string	`json:"MULTICAST_PACKETS_ORIGINATED"`
	Multicast_packets_consumed	string	`json:"MULTICAST_PACKETS_CONSUMED"`
	Multicast_bytes_sent	string	`json:"MULTICAST_BYTES_SENT"`
	Multicast_bytes_received	string	`json:"MULTICAST_BYTES_RECEIVED"`
	Multicast_bytes_forwarded	string	`json:"MULTICAST_BYTES_FORWARDED"`
	Multicast_bytes_originated	string	`json:"MULTICAST_BYTES_ORIGINATED"`
	Multicast_bytes_consumed	string	`json:"MULTICAST_BYTES_CONSUMED"`
	Broadcast_packets_sent	string	`json:"BROADCAST_PACKETS_SENT"`
	Broadcast_packets_received	string	`json:"BROADCAST_PACKETS_RECEIVED"`
	Broadcast_packets_forwarded	string	`json:"BROADCAST_PACKETS_FORWARDED"`
	Broadcast_packets_originated	string	`json:"BROADCAST_PACKETS_ORIGINATED"`
	Broadcast_packets_consumed	string	`json:"BROADCAST_PACKETS_CONSUMED"`
	Broadcast_bytes_sent	string	`json:"BROADCAST_BYTES_SENT"`
	Broadcast_bytes_received	string	`json:"BROADCAST_BYTES_RECEIVED"`
	Broadcast_bytes_forwarded	string	`json:"BROADCAST_BYTES_FORWARDED"`
	Broadcast_bytes_originated	string	`json:"BROADCAST_BYTES_ORIGINATED"`
	Broadcast_bytes_consumed	string	`json:"BROADCAST_BYTES_CONSUMED"`
	Labeled_packets_sent	string	`json:"LABELED_PACKETS_SENT"`
	Labeled_packets_received	string	`json:"LABELED_PACKETS_RECEIVED"`
	Labeled_packets_forwarded	string	`json:"LABELED_PACKETS_FORWARDED"`
	Labeled_packets_originated	string	`json:"LABELED_PACKETS_ORIGINATED"`
	Labeled_packets_consumed	string	`json:"LABELED_PACKETS_CONSUMED"`
	Labeled_bytes_sent	string	`json:"LABELED_BYTES_SENT"`
	Labeled_bytes_received	string	`json:"LABELED_BYTES_RECEIVED"`
	Labeled_bytes_forwarded	string	`json:"LABELED_BYTES_FORWARDED"`
	Labeled_bytes_originated	string	`json:"LABELED_BYTES_ORIGINATED"`
	Labeled_bytes_consumed	string	`json:"LABELED_BYTES_CONSUMED"`
	Wccp_redirect_outbound	string	`json:"WCCP_REDIRECT_OUTBOUND"`
	Wccp_redirect_inbound	string	`json:"WCCP_REDIRECT_INBOUND"`
	Wccp_redirect_exclude	string	`json:"WCCP_REDIRECT_EXCLUDE"`
}

var CiscoNxosShowIpInterface_Template = `Value Filldown VRF_NAME (\S+)
Value Required INTERFACE (\S+)
Value PROTOCOL_STATUS (\w+)
Value LINK_STATUS (\w+)
Value ADMIN_STATUS (\w+)
Value PRIMARY_IP_ADDRESS (\S+)
Value PRIMARY_IP_SUBNET (\S+)
Value PRIMARY_BROADCAST_ADDRESS (\S+)
Value List SECONDARY_IP_ADDRESS (\S+)
Value List SECONDARY_IP_SUBNET (\S+)
Value List SECONDARY_ROUTE_PREFERENCE (\d+)
Value List SECONDARY_ROUTE_PREFERENCE_TAG (\d+)
Value MULTICAST_GROUPS (224\..*?)
Value MTU (\d+)
Value ROUTE_PREFERENCE (\d+)
Value ROUTE_PREFERENCE_TAG (\d+)
Value PROXY_ARP (\S+)
Value LOCAL_PROXY_ARP (\S+)
Value MULTICAST_ROUTING (\S+)
Value ICMP_REDIRECTS (\S+)
Value DIRECTED_BROADCAST (\S+)
Value IP_FORWARDING (\S+)
Value ICMP_UNREACHABLES (\S+)
Value ICMP_PORT_UNREACHABLE (\S+)
Value URPF (.*)
Value LOAD_SHARING (\S+)
Value LAST_RESET (\S+)
Value ACL_OUT (\S+)
Value UNICAST_PACKETS_SENT (\d+)
Value UNICAST_PACKETS_RECEIVED (\d+)
Value UNICAST_PACKETS_FORWARDED (\d+)
Value UNICAST_PACKETS_ORIGINATED (\d+)
Value UNICAST_PACKETS_CONSUMED (\d+)
Value UNICAST_BYTES_SENT (\d+)
Value UNICAST_BYTES_RECEIVED (\d+)
Value UNICAST_BYTES_FORWARDED (\d+)
Value UNICAST_BYTES_ORIGINATED (\d+)
Value UNICAST_BYTES_CONSUMED (\d+)
Value MULTICAST_PACKETS_SENT (\d+)
Value MULTICAST_PACKETS_RECEIVED (\d+)
Value MULTICAST_PACKETS_FORWARDED (\d+)
Value MULTICAST_PACKETS_ORIGINATED (\d+)
Value MULTICAST_PACKETS_CONSUMED (\d+)
Value MULTICAST_BYTES_SENT (\d+)
Value MULTICAST_BYTES_RECEIVED (\d+)
Value MULTICAST_BYTES_FORWARDED (\d+)
Value MULTICAST_BYTES_ORIGINATED (\d+)
Value MULTICAST_BYTES_CONSUMED (\d+)
Value BROADCAST_PACKETS_SENT (\d+)
Value BROADCAST_PACKETS_RECEIVED (\d+)
Value BROADCAST_PACKETS_FORWARDED (\d+)
Value BROADCAST_PACKETS_ORIGINATED (\d+)
Value BROADCAST_PACKETS_CONSUMED (\d+)
Value BROADCAST_BYTES_SENT (\d+)
Value BROADCAST_BYTES_RECEIVED (\d+)
Value BROADCAST_BYTES_FORWARDED (\d+)
Value BROADCAST_BYTES_ORIGINATED (\d+)
Value BROADCAST_BYTES_CONSUMED (\d+)
Value LABELED_PACKETS_SENT (\d+)
Value LABELED_PACKETS_RECEIVED (\d+)
Value LABELED_PACKETS_FORWARDED (\d+)
Value LABELED_PACKETS_ORIGINATED (\d+)
Value LABELED_PACKETS_CONSUMED (\d+)
Value LABELED_BYTES_SENT (\d+)
Value LABELED_BYTES_RECEIVED (\d+)
Value LABELED_BYTES_FORWARDED (\d+)
Value LABELED_BYTES_ORIGINATED (\d+)
Value LABELED_BYTES_CONSUMED (\d+)
Value WCCP_REDIRECT_OUTBOUND (\S+)
Value WCCP_REDIRECT_INBOUND (\S+)
Value WCCP_REDIRECT_EXCLUDE (\S+)

Start
  ^IP\s+Interface\s+Status.*$$ -> Continue.Record
  ^IP\s+Interface\s+Status\s+for\s+VRF\s+\"${VRF_NAME}\"\s*$$
  ^\S+\s+Interface\s+status.*$$ -> Continue.Record
  ^${INTERFACE},.*protocol-${PROTOCOL_STATUS}\/link-${LINK_STATUS}\/admin-${ADMIN_STATUS},.*$$
  ^\s*IP\s+address:\s+${PRIMARY_IP_ADDRESS},\s+IP\s+subnet:\s+${PRIMARY_IP_SUBNET}(\s+route-preference.+)?$$
  # Secondary IP addresses are stored as a list
  ^\s*IP\s+address:\s+${SECONDARY_IP_ADDRESS},\s+IP\s+subnet:\s+${SECONDARY_IP_SUBNET}\s+secondary\s*(route-preference:\s+${SECONDARY_ROUTE_PREFERENCE},\s+tag:\s+${SECONDARY_ROUTE_PREFERENCE_TAG}\s*)?$$
#  (,\s+route-preference:\s+${SECONDARY_ROUTE_PREFERENCE},\s+tag:\d+)?
  ^\s*IP\s+broadcast\s+address:\s+${PRIMARY_BROADCAST_ADDRESS}\s*$$
  ^\s*IP\s+multicast\s+groups\s+locally\s+joined:.*$$
  # Multiple multicast groups are on the same line so we can't make a list out of those
  ^\s*${MULTICAST_GROUPS}\s*$$
  ^\s*IP\s+MTU:\s+${MTU}.*$$
  ^\s*IP\s+primary\s+address\s+route-preference:\s+${ROUTE_PREFERENCE}, tag:\s+${ROUTE_PREFERENCE_TAG}\s*$$
  ^\s*IP\s+proxy\s+ARP\s*:\s+${PROXY_ARP}\s*$$
  ^\s*IP\s+Local\s+Proxy\s+ARP\s*:\s+${LOCAL_PROXY_ARP}\s*$$
  ^\s*IP\s+multicast\s+routing:\s+${MULTICAST_ROUTING}\s*$$
  ^\s*IP\s+icmp\s+redirects:\s+${ICMP_REDIRECTS}\s*$$
  ^\s*IP\s+directed-broadcast:\s+${DIRECTED_BROADCAST}(,.*|\s*)$$
  ^\s*IP\s+Forwarding:\s+${IP_FORWARDING}(,.*|\s*)$$
  ^\s*IP\s+icmp\s+unreachables \(except port\):\s+${ICMP_UNREACHABLES}\s*$$
  ^\s*IP\s+icmp\s+port-unreachable:\s+${ICMP_PORT_UNREACHABLE}\s*$$
  ^\s*IP\s+unicast\s+reverse\s+path\s+forwarding:\s+${URPF}\s*$$
  ^\s*IP\s+load\s+sharing:\s+${LOAD_SHARING}\s*$$
  ^\s*IP\s+interface\s+statistics\s+last\s+reset:\s+${LAST_RESET}\s*$$
  ^\s*IP\s+interface\s+software\s+stats:.*$$
  ^\s*Unnumbered\s+interfaces\s+of
  ^\s*mti\d+:
  ^\s*IP\s+outbound\s+access\s+list:\s+${ACL_OUT}
  ^\s*Unicast\s+packets\s+:\s+${UNICAST_PACKETS_SENT}\/${UNICAST_PACKETS_RECEIVED}\/${UNICAST_PACKETS_FORWARDED}\/${UNICAST_PACKETS_ORIGINATED}\/${UNICAST_PACKETS_CONSUMED}\s*$$
  ^\s*Unicast\s+bytes\s+:\s+${UNICAST_BYTES_SENT}\/${UNICAST_BYTES_RECEIVED}\/${UNICAST_BYTES_FORWARDED}\/${UNICAST_BYTES_ORIGINATED}\/${UNICAST_BYTES_CONSUMED}\s*$$
  ^\s*Multicast\s+packets\s+:\s+${MULTICAST_PACKETS_SENT}\/${MULTICAST_PACKETS_RECEIVED}\/${MULTICAST_PACKETS_FORWARDED}\/${MULTICAST_PACKETS_ORIGINATED}\/${MULTICAST_PACKETS_CONSUMED}\s*$$
  ^\s*Multicast\s+bytes\s+:\s+${MULTICAST_BYTES_SENT}\/${MULTICAST_BYTES_RECEIVED}\/${MULTICAST_BYTES_FORWARDED}\/${MULTICAST_BYTES_ORIGINATED}\/${MULTICAST_BYTES_CONSUMED}\s*$$
  ^\s*Broadcast\s+packets\s+:\s+${BROADCAST_PACKETS_SENT}\/${BROADCAST_PACKETS_RECEIVED}\/${BROADCAST_PACKETS_FORWARDED}\/${BROADCAST_PACKETS_ORIGINATED}\/${BROADCAST_PACKETS_CONSUMED}\s*$$
  ^\s*Broadcast\s+bytes\s+:\s+${BROADCAST_BYTES_SENT}\/${BROADCAST_BYTES_RECEIVED}\/${BROADCAST_BYTES_FORWARDED}\/${BROADCAST_BYTES_ORIGINATED}\/${BROADCAST_BYTES_CONSUMED}\s*$$
  ^\s*Labeled\s+packets\s+:\s+${LABELED_PACKETS_SENT}\/${LABELED_PACKETS_RECEIVED}\/${LABELED_PACKETS_FORWARDED}\/${LABELED_PACKETS_ORIGINATED}\/${LABELED_PACKETS_CONSUMED}\s*$$
  ^\s*Labeled\s+bytes\s+:\s+${LABELED_BYTES_SENT}\/${LABELED_BYTES_RECEIVED}\/${LABELED_BYTES_FORWARDED}\/${LABELED_BYTES_ORIGINATED}\/${LABELED_BYTES_CONSUMED}\s*$$
  ^\s*WCCP\s+Redirect\s+outbound:\s+${WCCP_REDIRECT_OUTBOUND}\s*$$
  ^\s*WCCP\s+Redirect\s+inbound:\s+${WCCP_REDIRECT_INBOUND}\s*$$
  ^\s*WCCP\s+Redirect\s+exclude:\s+${WCCP_REDIRECT_EXCLUDE}\s*$$
  ^\s*$$
  ^. -> Error

`