package cisco_asa

type ShowInterface struct {
	Interface               string `json:"INTERFACE"`
	Speed                   string `json:"SPEED"`
	Netmask                 string `json:"NETMASK"`
	Fivemin_in_rate         string `json:"FIVEMIN_IN_RATE"`
	Tunnel_source_ip        string `json:"TUNNEL_SOURCE_IP"`
	Tunnel_destination_ip   string `json:"TUNNEL_DESTINATION_IP"`
	Delay                   string `json:"DELAY"`
	Bonded_port             string `json:"BONDED_PORT"`
	Link_status             string `json:"LINK_STATUS"`
	Interface_zone          string `json:"INTERFACE_ZONE"`
	Description             string `json:"DESCRIPTION"`
	Mac_address             string `json:"MAC_ADDRESS"`
	Vlan_id                 string `json:"VLAN_ID"`
	Onemin_in_pps           string `json:"ONEMIN_IN_PPS"`
	Onemin_out_rate         string `json:"ONEMIN_OUT_RATE"`
	Bandwidth               string `json:"BANDWIDTH"`
	Duplex                  string `json:"DUPLEX"`
	Bonded_port_members     string `json:"BONDED_PORT_MEMBERS"`
	Mtu                     string `json:"MTU"`
	Ip_address              string `json:"IP_ADDRESS"`
	Onemin_drop_rate        string `json:"ONEMIN_DROP_RATE"`
	Fivemin_in_pps          string `json:"FIVEMIN_IN_PPS"`
	Hardware_type           string `json:"HARDWARE_TYPE"`
	Onemin_in_rate          string `json:"ONEMIN_IN_RATE"`
	Fivemin_out_pps         string `json:"FIVEMIN_OUT_PPS"`
	Fivemin_out_rate        string `json:"FIVEMIN_OUT_RATE"`
	Tunnel_mode             string `json:"TUNNEL_MODE"`
	Tunnel_ipsec_profile    string `json:"TUNNEL_IPSEC_PROFILE"`
	Protocol_status         string `json:"PROTOCOL_STATUS"`
	Onemin_out_pps          string `json:"ONEMIN_OUT_PPS"`
	Fivemin_drop_rate       string `json:"FIVEMIN_DROP_RATE"`
	Tunnel_source_interface string `json:"TUNNEL_SOURCE_INTERFACE"`
}

var ShowInterface_Template string = `Value Required INTERFACE (\S+)
Value INTERFACE_ZONE (.+?)
Value LINK_STATUS (.+?)
Value PROTOCOL_STATUS (.*)
Value HARDWARE_TYPE (Virtual Tunnel|EtherChannel/LACP|[\w ]+)
Value BANDWIDTH (\d+\s+\w+)
Value DELAY (\d+\s+\w+)
Value DUPLEX (\w+\-\w+)
Value SPEED (\d+\w+\s\w+)
Value DESCRIPTION (.*)
Value BONDED_PORT (.+?)
Value BONDED_PORT_MEMBERS (.*)
Value MAC_ADDRESS ([a-zA-Z0-9]+.[a-zA-Z0-9]+.[a-zA-Z0-9]+)
Value MTU (\d+)
Value VLAN_ID (\d+)
Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value NETMASK (\d+\.\d+\.\d+\.\d+)
Value ONEMIN_IN_PPS (\d+)
Value ONEMIN_IN_RATE (\d+)
Value ONEMIN_OUT_PPS (\d+)
Value ONEMIN_OUT_RATE (\d+)
Value ONEMIN_DROP_RATE (\d+)
Value FIVEMIN_IN_PPS (\d+)
Value FIVEMIN_IN_RATE (\d+)
Value FIVEMIN_OUT_PPS (\d+)
Value FIVEMIN_OUT_RATE (\d+)
Value FIVEMIN_DROP_RATE (\d+)
Value TUNNEL_SOURCE_INTERFACE (\w+|\w+\s\w+)
Value TUNNEL_SOURCE_IP (\d+\.\d+\.\d+\.\d+)
Value TUNNEL_DESTINATION_IP (\d+\.\d+\.\d+\.\d+)
Value TUNNEL_MODE (\w+|\w+\s\w+)
Value TUNNEL_IPSEC_PROFILE (.*)

Start
  ^\s+Tunnel\s+Interface\s+Information:
  ^.*Interface\s+ -> Continue.Record
  ^.*Interface\s+${INTERFACE}\s+"${INTERFACE_ZONE}",\s+is\s+${LINK_STATUS},.*protocol\s+is\s+${PROTOCOL_STATUS}
  ^.*Interface\s+${INTERFACE}.*is\s+${LINK_STATUS},.*protocol\s+is\s+${PROTOCOL_STATUS}
  ^\s+Hardware\s+is\s+${HARDWARE_TYPE} -> Continue
  ^.*BW\s+${BANDWIDTH},\s+DLY\s+${DELAY}
  ^.*\(${DUPLEX}\),\s+Auto-Speed\(${SPEED}\)
  ^.*\(${DUPLEX}\),\s+\d+\s+Mbps\(${SPEED}\)
  ^.*\(${DUPLEX}\),\s+\(${SPEED}\)
  ^.*Duplex,\s+Auto-Speed
  ^.*Description:\s+${DESCRIPTION}
  ^.*VLAN\s+identifier\s+(${VLAN_ID}|(?:none))
  ^.*MAC\s+address\s+${MAC_ADDRESS},\s+MTU\s+${MTU}
  ^.*MAC\s+address\s+${MAC_ADDRESS},\s+MTU\s+not\s+set
  ^\s+Available\s+but\s+not\s+configured\s+(?:with\s+VLAN|via\s+nameif)
  ^.*MAC\s+address\s+N/A,\s+MTU\s+${MTU}
  ^.*IP\s+address\s+${IP_ADDRESS},\s+subnet\s+mask\s+${NETMASK}
  ^.*1\s+minute\s+input\s+rate\s+${ONEMIN_IN_PPS}\s+pkts/sec,\s+${ONEMIN_IN_RATE}\s+bytes/sec
  ^.*1\s+minute\s+output\s+rate\s+${ONEMIN_OUT_PPS}\s+pkts/sec,\s+${ONEMIN_OUT_RATE}\s+bytes/sec
  ^.*1\s+minute\s+drop\s+rate,\s+${ONEMIN_DROP_RATE}
  ^.*5\s+minute\s+input\s+rate\s+${FIVEMIN_IN_PPS}\s+pkts/sec,\s+${FIVEMIN_IN_RATE}\s+bytes/sec
  ^.*5\s+minute\s+output\s+rate\s+${FIVEMIN_OUT_PPS}\s+pkts/sec,\s+${FIVEMIN_OUT_RATE}\s+bytes/sec
  ^.*5\s+minute\s+drop\s+rate,\s+${FIVEMIN_DROP_RATE}
  ^\s+Active\s+member\s+of\s+${BONDED_PORT}\s*$$
  ^\s+Active:\s+${BONDED_PORT_MEMBERS}
  ^\s+Source\s+interface:\s+${TUNNEL_SOURCE_INTERFACE}\s+IP\s+address:\s+${TUNNEL_SOURCE_IP}
  ^\s+Destination\s+IP\s+address:\s+${TUNNEL_DESTINATION_IP}
  ^\s+Mode:\s+${TUNNEL_MODE}\s+IPsec\s+profile:\s+${TUNNEL_IPSEC_PROFILE}
  ^.*Input\s+flow\s+control\s+is\s+unsupported,\s+output\s+flow\s+control\s+is\s+off
  ^.*\d+\s+packets\s+input,\s+\d+\s+bytes,\s+\d+\s+no\s+buffer
  ^.*Received\s+\d+\s+broadcasts,\s+\d+\s+runts,\s+\d+\s+giants
  ^.*\d+\s+input\s+errors,\s+\d+\s+CRC,\s+\d+\s+frame,\s+\d+\s+overrun,\s+\d+\s+ignored,\s+\d+\s+abort
  ^.*\d+\s+pause\s+input,\s+\d+\s+resume\s+input
  ^.*\d+\s+(L2|output)\s+decode\s+drops
  ^.*\d+\s+packets\s+output,\s+\d+\s+bytes,\s+\d+\s+underruns
  ^.*\d+\s+pause\s+output,\s+\d+\s+resume\s+output
  ^.*\d+\s+output\s+errors,\s+\d+\s+collisions,\s+\d+\s+interface\s+resets
  ^.*\d+\s+late\s+collisions,\s+\d+\s+deferred
  ^.*\d+\s+input\s+reset\s+drops,\s+\d+\s+output\s+reset\s+drops
  ^.*input\s+queue\s+\(blocks\s+free\s+curr\/low\):\s+hardware\s+\(\d+\/\d+\)
  ^.*output\s+queue\s+\(blocks\s+free\s+curr\/low\):\s+hardware\s+\(\d+\/\d+\)
  ^.*Traffic\s+Statistics\s+for\s+".+?":
  ^.*\d+\s+packets\s+input,\s+\d+\s+bytes
  ^.*\d+\s+packets\s+output,\s+\d+\s+bytes
  ^.*\d+\s+packets\s+dropped
  ^.*Management-only\sinterface\.\s+Blocked\s+\d+\s+through-the-device\s+packets
  ^.*Input\s+flow\s+control\s+is\s+unsupported,\s+output\s+flow\s+control\s+is\s+unsupported
  ^.*Available\s+but\s+not\s+configured\s+via\s+nameif
  ^.*IP\s+address\s+unassigned
  ^\s*\d+\s+lost\s+carrier,\s+\d+\s+no\s+carrier
  ^\s*(input|output)\s+queue\s+\(curr\/max\s+packets\):\s+hardware\s+\(\d+\/\d+\)\s+software\s+\(\d+\/\d+\)
  ^.*input\s+queue\s+\(blocks\s+free\s+curr\/low\):\s+hardware\s+\(\d+\/\d+\)
  ^\s+Members\s+in\s+this\s+channel:
  ^.*MAC\s+address\s+N/A,\s+MTU\s+not\s+set 
  ^.*\d+\s+pause/resume\s+(input|output)
  ^.*\d+\s+switch\s+(ingress|egress)\s+policy\s+drops
  ^.*\d+\s+rate\s+limit\s+drops
  ^\s*$$
  ^. -> Error
`
