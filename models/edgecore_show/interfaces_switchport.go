package edgecore_show 

type InterfacesSwitchport struct {
	Interface	string	`json:"INTERFACE"`
	Broadcast_threshold	string	`json:"BROADCAST_THRESHOLD"`
	Multicast_threshold	string	`json:"MULTICAST_THRESHOLD"`
	Unicast_threshold	string	`json:"UNICAST_THRESHOLD"`
	Unknown_unicast_threshold	string	`json:"UNKNOWN_UNICAST_THRESHOLD"`
	Lacp_status	string	`json:"LACP_STATUS"`
	Ingress_rate_limit	string	`json:"INGRESS_RATE_LIMIT"`
	Egress_rate_limit	string	`json:"EGRESS_RATE_LIMIT"`
	Vlan_membership_mode	string	`json:"VLAN_MEMBERSHIP_MODE"`
	Ingress_rule	string	`json:"INGRESS_RULE"`
	Acceptable_frame_type	string	`json:"ACCEPTABLE_FRAME_TYPE"`
	Native_vlan	string	`json:"NATIVE_VLAN"`
	Priority_for_untagged_traffic	string	`json:"PRIORITY_FOR_UNTAGGED_TRAFFIC"`
	Gvrp_status	string	`json:"GVRP_STATUS"`
	Allowed_vlan	[]string	`json:"ALLOWED_VLAN"`
	Allowed_vlan_tag	[]string	`json:"ALLOWED_VLAN_TAG"`
	Forbidden_vlan	[]string	`json:"FORBIDDEN_VLAN"`
	Forbidden_vlan_tag	[]string	`json:"FORBIDDEN_VLAN_TAG"`
	Private_vlan_mode	string	`json:"PRIVATE_VLAN_MODE"`
	Private_vlan_host_association	string	`json:"PRIVATE_VLAN_HOST_ASSOCIATION"`
	Private_vlan_mapping	string	`json:"PRIVATE_VLAN_MAPPING"`
	Tunnel_status_802_1q	string	`json:"TUNNEL_STATUS_802_1Q"`
	Tunnel_mode_802_1q	string	`json:"TUNNEL_MODE_802_1Q"`
	Tunnel_tpid_802_1q	string	`json:"TUNNEL_TPID_802_1Q"`
	Layer_2_protocol_tunnel	string	`json:"LAYER_2_PROTOCOL_TUNNEL"`
}

var InterfacesSwitchport_Template = `Value INTERFACE (.*)
Value BROADCAST_THRESHOLD (.*)
Value MULTICAST_THRESHOLD (.*)
Value UNICAST_THRESHOLD (.*)
Value UNKNOWN_UNICAST_THRESHOLD (.*)
Value LACP_STATUS (.*)
Value INGRESS_RATE_LIMIT (.*)
Value EGRESS_RATE_LIMIT (.*)
Value VLAN_MEMBERSHIP_MODE (.*)
Value INGRESS_RULE (.*)
Value ACCEPTABLE_FRAME_TYPE (.*)
Value NATIVE_VLAN (.*)
Value PRIORITY_FOR_UNTAGGED_TRAFFIC (.*)
Value GVRP_STATUS (.*)
Value List ALLOWED_VLAN (\d+)
Value List ALLOWED_VLAN_TAG (\w+)
Value List FORBIDDEN_VLAN (\d+)
Value List FORBIDDEN_VLAN_TAG (\w+)
Value PRIVATE_VLAN_MODE (.*)
Value PRIVATE_VLAN_HOST_ASSOCIATION (.*)
Value PRIVATE_VLAN_MAPPING (.*)
Value TUNNEL_STATUS_802_1Q (.*)
Value TUNNEL_MODE_802_1Q (.*)
Value TUNNEL_TPID_802_1Q (.*)
Value LAYER_2_PROTOCOL_TUNNEL (.*)

Start
  ^\s*Information\sof\s.*$$ -> Continue.Record
  ^\s*Information\sof\s${INTERFACE}\s*$$
  ^\s*Broadcast\s+Threshold:\s*${BROADCAST_THRESHOLD}\s*$$
  ^\s*Multicast\s+Threshold:\s*${MULTICAST_THRESHOLD}\s*$$
  ^\s*Unicast\s+Threshold:\s*${UNICAST_THRESHOLD}\s*$$
  ^\s*Unknown-unicast\s+Threshold:\s*${UNKNOWN_UNICAST_THRESHOLD}\s*$$
  ^\s*LACP\s+Status:\s*${LACP_STATUS}\s*$$
  ^\s*Ingress\s+Rate\s+Limit:\s*${INGRESS_RATE_LIMIT}\s*$$
  ^\s*Egress\s+Rate\s+Limit:\s*${EGRESS_RATE_LIMIT}\s*$$
  ^\s*VLAN\sMembership\sMode:\s*${VLAN_MEMBERSHIP_MODE}\s*$$
  ^\s*Ingress\s+Rule:\s*${INGRESS_RULE}\s*$$
  ^\s*Acceptable\s+Frame\s+Type:\s*${ACCEPTABLE_FRAME_TYPE}\s*$$
  ^\s*Native\sVLAN:\s*${NATIVE_VLAN}\s*$$
  ^\s*Priority\s+for\s+Untagged\s+Traffic:\s*${PRIORITY_FOR_UNTAGGED_TRAFFIC}\s*$$
  ^\s*GVRP\s+Status:\s*${GVRP_STATUS}\s*$$
  ^\s*Allowed\sVLAN:\s*${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s*Allowed\sVLAN:\s*(\d+\(\w+\),\s*){1}${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s*Allowed\sVLAN:\s*(\d+\(\w+\),\s*){2}${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s*Allowed\sVLAN:\s*(\d+\(\w+\),\s*){3}${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s*Allowed\sVLAN:\s*(\d+\(\w+\),\s*){4}${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s*Allowed\sVLAN:\s*(\d+\(\w+\),\s*){5}${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s*Allowed\sVLAN:\s*(\d+\(\w+\),\s*){6}${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s*Allowed\sVLAN:\s*(\d+\(\w+\),\s*){7}${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s*Allowed\sVLAN:\s*(\d+\(\w+\),\s*){8}${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s*Allowed\sVLAN:\s*(\d+\(\w+\),\s*){9}${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s*Allowed\sVLAN:\s*(\d+\(\w+\),?\s*)*$$
  ^\s+${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),\s*){1}${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),\s*){2}${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),\s*){3}${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),\s*){4}${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),\s*){5}${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),\s*){6}${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),\s*){7}${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),\s*){8}${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),\s*){9}${ALLOWED_VLAN}\(${ALLOWED_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),?\s*)*$$
  ^\s*Forbidden\sVLAN:\s*${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s*Forbidden\sVLAN:\s*(\d+\(\w+\),\s*){1}${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s*Forbidden\sVLAN:\s*(\d+\(\w+\),\s*){2}${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s*Forbidden\sVLAN:\s*(\d+\(\w+\),\s*){3}${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s*Forbidden\sVLAN:\s*(\d+\(\w+\),\s*){4}${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s*Forbidden\sVLAN:\s*(\d+\(\w+\),\s*){5}${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s*Forbidden\sVLAN:\s*(\d+\(\w+\),\s*){6}${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s*Forbidden\sVLAN:\s*(\d+\(\w+\),\s*){7}${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s*Forbidden\sVLAN:\s*(\d+\(\w+\),\s*){8}${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s*Forbidden\sVLAN:\s*(\d+\(\w+\),\s*){9}${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s*Forbidden\sVLAN:\s*(\d+\(\w+\),?\s*)*$$
  ^\s+${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),\s*){1}${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),\s*){2}${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),\s*){3}${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),\s*){4}${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),\s*){5}${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),\s*){6}${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),\s*){7}${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),\s*){8}${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),\s*){9}${FORBIDDEN_VLAN}\(${FORBIDDEN_VLAN_TAG}\).*$$ -> Continue
  ^\s+(\d+\(\w+\),?\s*)*$$
  ^\s*Private-VLAN\s+Mode:\s*${PRIVATE_VLAN_MODE}\s*$$
  ^\s*Private-VLAN\s+host-association:\s*${PRIVATE_VLAN_HOST_ASSOCIATION}\s*$$
  ^\s*Private-VLAN\s+Mapping:\s*${PRIVATE_VLAN_MAPPING}\s*$$
  ^\s*802.1Q-tunnel\s+Status:\s*${TUNNEL_STATUS_802_1Q}\s*$$
  ^\s*802.1Q-tunnel\s+Mode:\s*${TUNNEL_MODE_802_1Q}\s*$$
  ^\s*802.1Q-tunnel\s+TPID:\s*${TUNNEL_TPID_802_1Q}\s*$$
  ^\s*Layer\s+2\s+Protocol\s+Tunnel\s*:\s*${LAYER_2_PROTOCOL_TUNNEL}\s*$$
  ^\s*$$
  ^. -> Error
`