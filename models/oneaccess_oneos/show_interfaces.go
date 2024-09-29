package oneaccess_oneos 

type ShowInterfaces struct {
	Interface	string	`json:"INTERFACE"`
	Flags	string	`json:"FLAGS"`
	If_index	string	`json:"IF_INDEX"`
	Description	string	`json:"DESCRIPTION"`
	Link_status	string	`json:"LINK_STATUS"`
	Protocol_status	string	`json:"PROTOCOL_STATUS"`
	Promiscuous_mode	string	`json:"PROMISCUOUS_MODE"`
	Encapsulation	string	`json:"ENCAPSULATION"`
	Vlan_id	string	`json:"VLAN_ID"`
	Ipv4_mtu	string	`json:"IPV4_MTU"`
	Ipv6_mtu	string	`json:"IPV6_MTU"`
	Uptime	string	`json:"UPTIME"`
	Downtime	string	`json:"DOWNTIME"`
	Status_changes	string	`json:"STATUS_CHANGES"`
	Negotiation	string	`json:"NEGOTIATION"`
	Duplex	string	`json:"DUPLEX"`
	Ipv4	string	`json:"IPV4"`
	Ipv4_unnumbered_intf	string	`json:"IPV4_UNNUMBERED_INTF"`
	Ipv4_unnumbered_ip	string	`json:"IPV4_UNNUMBERED_IP"`
	Ipv4_gateway	string	`json:"IPV4_GATEWAY"`
	Ipv4_broadcast	string	`json:"IPV4_BROADCAST"`
	Ipv4_secondary	[]string	`json:"IPV4_SECONDARY"`
	Ipv6	[]string	`json:"IPV6"`
	Tunnel_source	string	`json:"TUNNEL_SOURCE"`
	Tunnel_destination	string	`json:"TUNNEL_DESTINATION"`
	Tunnel_protocol	string	`json:"TUNNEL_PROTOCOL"`
	Tunnel_transport	string	`json:"TUNNEL_TRANSPORT"`
	Tunnel_protection_mode	string	`json:"TUNNEL_PROTECTION_MODE"`
	Tunnel_protection_profile	string	`json:"TUNNEL_PROTECTION_PROFILE"`
	Tunnel_vrf	string	`json:"TUNNEL_VRF"`
	Firewall_zone	string	`json:"FIREWALL_ZONE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Arp_timeout	string	`json:"ARP_TIMEOUT"`
	Bandwidth_kbps	string	`json:"BANDWIDTH_KBPS"`
	Bandwidth_limit_kbps	string	`json:"BANDWIDTH_LIMIT_KBPS"`
	Sfp_mediatype	string	`json:"SFP_MEDIATYPE"`
	Input_rate	string	`json:"INPUT_RATE"`
	Output_rate	string	`json:"OUTPUT_RATE"`
	Input_pps	string	`json:"INPUT_PPS"`
	Output_pps	string	`json:"OUTPUT_PPS"`
	Input_load_pct	string	`json:"INPUT_LOAD_PCT"`
	Output_load_pct	string	`json:"OUTPUT_LOAD_PCT"`
	Congestion_dropped_rx	string	`json:"CONGESTION_DROPPED_RX"`
	Congestion_dropped_cpu	string	`json:"CONGESTION_DROPPED_CPU"`
	Congestion_dropped_total	string	`json:"CONGESTION_DROPPED_TOTAL"`
	Bridge_group	string	`json:"BRIDGE_GROUP"`
	Output_queue_strategy	string	`json:"OUTPUT_QUEUE_STRATEGY"`
	Output_queue_length	string	`json:"OUTPUT_QUEUE_LENGTH"`
	Output_queue_depth	string	`json:"OUTPUT_QUEUE_DEPTH"`
	Shaper_packets_dequeued	string	`json:"SHAPER_PACKETS_DEQUEUED"`
	Shaper_burst_current	string	`json:"SHAPER_BURST_CURRENT"`
	Shaper_burst_max	string	`json:"SHAPER_BURST_MAX"`
	Reliability_value	string	`json:"RELIABILITY_VALUE"`
	Reliability_max	string	`json:"RELIABILITY_MAX"`
	Input_packets	string	`json:"INPUT_PACKETS"`
	Input_bytes	string	`json:"INPUT_BYTES"`
	Input_qdrops	string	`json:"INPUT_QDROPS"`
	Input_broadcasts	string	`json:"INPUT_BROADCASTS"`
	Input_multicasts	string	`json:"INPUT_MULTICASTS"`
	Input_errors	string	`json:"INPUT_ERRORS"`
	Input_discards	string	`json:"INPUT_DISCARDS"`
	Input_mac_acl_discards	string	`json:"INPUT_MAC_ACL_DISCARDS"`
	Input_unknown_protocols	string	`json:"INPUT_UNKNOWN_PROTOCOLS"`
	Output_packets	string	`json:"OUTPUT_PACKETS"`
	Output_bytes	string	`json:"OUTPUT_BYTES"`
	Output_qdrops	string	`json:"OUTPUT_QDROPS"`
	Output_broadcasts	string	`json:"OUTPUT_BROADCASTS"`
	Output_multicasts	string	`json:"OUTPUT_MULTICASTS"`
	Output_errors	string	`json:"OUTPUT_ERRORS"`
	Output_discards	string	`json:"OUTPUT_DISCARDS"`
	Vrf	string	`json:"VRF"`
}

var ShowInterfaces_Template = `Value Required INTERFACE (\S+(?:\s\S+)?)
Value FLAGS (.*)
Value IF_INDEX (\d+)
Value DESCRIPTION (.*)
Value LINK_STATUS (up|down)
Value PROTOCOL_STATUS (up|down)
Value PROMISCUOUS_MODE (\S+)
Value ENCAPSULATION (.+?)
Value VLAN_ID (\d+)
Value IPV4_MTU (\d+)
Value IPV6_MTU (\d+)
Value UPTIME (\S+)
Value DOWNTIME (\S+)
Value STATUS_CHANGES (\d+)
Value NEGOTIATION ([aA]uto|no|[Uu]nknown)
Value DUPLEX (full|half|[Uu]nknown)
Value IPV4 (\d+\.\d+\.\d+\.\d+\/\d+|unnumbered)
Value IPV4_UNNUMBERED_INTF (.*)
Value IPV4_UNNUMBERED_IP (\S+)
Value IPV4_GATEWAY (\d+\.\d+\.\d+\.\d+)
Value IPV4_BROADCAST (\d+\.\d+\.\d+\.\d+)
Value List IPV4_SECONDARY (\d+\.\d+\.\d+\.\d+\/\d+)
Value List IPV6 (\S+)
Value TUNNEL_SOURCE (\S+)
Value TUNNEL_DESTINATION (\S+)
Value TUNNEL_PROTOCOL (\S+)
Value TUNNEL_TRANSPORT (\S+)
Value TUNNEL_PROTECTION_MODE (\S+)
Value TUNNEL_PROTECTION_PROFILE (\S+)
Value TUNNEL_VRF (\S+)
Value FIREWALL_ZONE (.*)
Value MAC_ADDRESS ([a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2})
Value ARP_TIMEOUT (\d+)
Value BANDWIDTH_KBPS (\d+|unknown)
Value BANDWIDTH_LIMIT_KBPS (\d+)
Value SFP_MEDIATYPE (\S+)
Value INPUT_RATE (\d+)
Value OUTPUT_RATE (\d+)
Value INPUT_PPS (\d+)
Value OUTPUT_PPS (\d+)
Value INPUT_LOAD_PCT (\d+(\.\d+)?)
Value OUTPUT_LOAD_PCT (\d+(\.\d+)?)
Value CONGESTION_DROPPED_RX (\d+)
Value CONGESTION_DROPPED_CPU (\d+)
Value CONGESTION_DROPPED_TOTAL (\d+)
Value BRIDGE_GROUP (\d+)
Value OUTPUT_QUEUE_STRATEGY (.+)
Value OUTPUT_QUEUE_LENGTH (\d+)
Value OUTPUT_QUEUE_DEPTH (\d+)
Value SHAPER_PACKETS_DEQUEUED (\-?\d+)
Value SHAPER_BURST_CURRENT (\d+)
Value SHAPER_BURST_MAX (\d+)
Value RELIABILITY_VALUE (\d+)
Value RELIABILITY_MAX (\d+)
Value INPUT_PACKETS (\d+)
Value INPUT_BYTES (\d+)
Value INPUT_QDROPS (\d+)
Value INPUT_BROADCASTS (\d+)
Value INPUT_MULTICASTS (\d+)
Value INPUT_ERRORS (\d+)
Value INPUT_DISCARDS (\d+)
Value INPUT_MAC_ACL_DISCARDS (\d+)
Value INPUT_UNKNOWN_PROTOCOLS (\d+)
Value OUTPUT_PACKETS (\d+)
Value OUTPUT_BYTES (\d+)
Value OUTPUT_QDROPS (\d+)
Value OUTPUT_BROADCASTS (\d+)
Value OUTPUT_MULTICASTS (\d+)
Value OUTPUT_ERRORS (\d+)
Value OUTPUT_DISCARDS (\d+)
Value VRF (\S+)
					

Start
  ^\S+(?:\s\S+)?\s+is\s+(?:up|down),\s+line\s+protocol.*$$ -> Continue.Record
  ^${INTERFACE}\s+is\s+${LINK_STATUS},\s+line\s+protocol\s+is\s+${PROTOCOL_STATUS}\s*$$
  ^\s+Flags:\s\(\S+\)\s${FLAGS},\sinterface\sindex\sis\s${IF_INDEX}
  ^\s+Description:\s*${DESCRIPTION}
  ^\s+Promiscuous\smode\s${PROMISCUOUS_MODE}
  ^\s+Encapsulation:\s${ENCAPSULATION}\sVirtual\sLAN,\sVLAN\sID\s${VLAN_ID} -> Continue
  ^\s+Encapsulation:\s${ENCAPSULATION},\s -> Continue
  ^.*\sIPv4\sMTU\s${IPV4_MTU}\sbytes -> Continue
  ^.*\sIPv6\sMTU\s${IPV6_MTU}\sbytes -> Continue
  ^.*\s+MTU\s${IPV4_MTU}\sbytes
  ^\s+Firewalls?\szone:\s${FIREWALL_ZONE}
  ^\s+This\sinterface\sis\sassociated\sto\sVRF\s${VRF}
  ^\s+Up-time\s${UPTIME},\sstatus\schange\scount\s${STATUS_CHANGES}
  ^\s+Down-time\s${DOWNTIME},\sstatus\schange\scount\s${STATUS_CHANGES}
  ^\s+${NEGOTIATION}[\-\s]negotiation,\s+${DUPLEX}[\-\s]duplex
  ^\s+${NEGOTIATION}[\-\s]negotiation,\s+duplex[\-\s]${DUPLEX}
  ^\s+negotiation\s+${NEGOTIATION},\s+duplex\s+${DUPLEX}
  ^\s+Internet\saddress\sis\s${IPV4},\sdestination\saddress\sis\s${IPV4_GATEWAY}
  ^\s+Internet\saddress\sis\s${IPV4},\sbroadcast\saddress\sis\s${IPV4_BROADCAST}
  ^\s+Internet\saddress\sis\s${IPV4}
  ^\s+Secondary\saddress\sis\s${IPV4_SECONDARY}
  ^\s+Interface\sis\s${IPV4},\suse\saddress\sof\s${IPV4_UNNUMBERED_INTF}\(${IPV4_UNNUMBERED_IP}\)
  ^\s+IPv6\saddress\sis\s${IPV6}
  ^\s+Hardware\saddress\sis\s${MAC_ADDRESS},\sARP\stimeout\s${ARP_TIMEOUT}\ssec
  ^\s+Line\sspeed\s${BANDWIDTH_KBPS}(\skbps,\sbandwidth\slimit\s${BANDWIDTH_LIMIT_KBPS}\skbps)?
  ^\s+[Mm]edia\-type\s${SFP_MEDIATYPE}
  ^\s+Mean\sinput/output\srate\s+${INPUT_RATE}/${OUTPUT_RATE}\sbits/s,\s${INPUT_PPS}/${OUTPUT_PPS}\s.*
  ^\s+Mean\sinput/output\sload\spercentage\s+${INPUT_LOAD_PCT}/${OUTPUT_LOAD_PCT}\spercent.*
  ^\s+Congestion\sManagement\sDropped\spackets:\s+RX:${CONGESTION_DROPPED_RX},\s*CPU:${CONGESTION_DROPPED_CPU},\s*Total:${CONGESTION_DROPPED_TOTAL}
  ^\s+Bridged\sto\sgroup\s${BRIDGE_GROUP}
  ^\s+Tunnel\ssource\s${TUNNEL_SOURCE},\sdestination\s${TUNNEL_DESTINATION}
  ^\s+Tunnel\sprotocol/transport\s${TUNNEL_PROTOCOL}/${TUNNEL_TRANSPORT}
  ^\s+Tunnel\sprotection\svia\s${TUNNEL_PROTECTION_MODE}\s\(profile\s\"${TUNNEL_PROTECTION_PROFILE}\"\)
  ^\s+This\sinterface\sis\sassociated\sto\sTunnel\sVRF\s${TUNNEL_VRF}
  ^\s+Output\squeuing\sstrategy:\s*${OUTPUT_QUEUE_STRATEGY},\s*output\squeue\slength/depth\s${OUTPUT_QUEUE_LENGTH}/${OUTPUT_QUEUE_DEPTH}
  ^\s+Output\squeuing\sstrategy:\s*${OUTPUT_QUEUE_STRATEGY}
  ^\s+Shaper:\spackets\sdequeued\s${SHAPER_PACKETS_DEQUEUED},\sburst\(current/max\)\s${SHAPER_BURST_CURRENT}/${SHAPER_BURST_MAX}
  ^\s+Reliability:\s*${RELIABILITY_VALUE}/${RELIABILITY_MAX}
  ^\s+IN:\s*${INPUT_PACKETS}\spackets,\s${INPUT_BYTES}\sbytes,\s${INPUT_QDROPS}\squeue\sdrops -> IN
  ^\s+OUT:\s*${OUTPUT_PACKETS}\spackets,\s${OUTPUT_BYTES}\sbytes,\s${OUTPUT_QDROPS}\squeue\sdrops -> OUT
  ^\s+Keepalive
  ^\s+\d+\s+bad\s+checksum,
  ^\s*$$
  ^. -> Error

IN
  ^\s+${INPUT_BROADCASTS}\sbroadcasts,\s*${INPUT_MULTICASTS}\smulticasts,\s*${INPUT_ERRORS}\serrors,\s*${INPUT_DISCARDS}\sdiscards(,\s*${INPUT_MAC_ACL_DISCARDS}\smac\sacl\sdiscards)?
  ^\s+${INPUT_UNKNOWN_PROTOCOLS}\sunknown\sprotocols -> Start
  ^\s*$$
  ^. -> Error

OUT
  ^\s+${OUTPUT_BROADCASTS}\sbroadcasts,\s*${OUTPUT_MULTICASTS}\smulticasts,\s*${OUTPUT_ERRORS}\serrors,\s*${OUTPUT_DISCARDS}\sdiscards -> Start
  ^\s*$$
  ^. -> Error
`