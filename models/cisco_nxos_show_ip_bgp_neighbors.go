package models

type CiscoNxosShowIpBgpNeighbors struct {
	Neighbor	string	`json:"NEIGHBOR"`
	Peer_group	string	`json:"PEER_GROUP"`
	Description	string	`json:"DESCRIPTION"`
	Vrf	string	`json:"VRF"`
	Asn	string	`json:"ASN"`
	Remote_asn	string	`json:"REMOTE_ASN"`
	Local_asn	string	`json:"LOCAL_ASN"`
	Bgp_state	string	`json:"BGP_STATE"`
	Uptime	string	`json:"UPTIME"`
	Source_iface	string	`json:"SOURCE_IFACE"`
	Conn_estab	string	`json:"CONN_ESTAB"`
	Conn_dropped	string	`json:"CONN_DROPPED"`
	Last_reset	string	`json:"LAST_RESET"`
	Last_reset_reason	string	`json:"LAST_RESET_REASON"`
	Last_peer_reset	string	`json:"LAST_PEER_RESET"`
	Last_peer_reset_reason	string	`json:"LAST_PEER_RESET_REASON"`
	Addr_family	[]string	`json:"ADDR_FAMILY"`
	Table_version	[]string	`json:"TABLE_VERSION"`
	Nei_table_version	[]string	`json:"NEI_TABLE_VERSION"`
	Addr_fam_adv	string	`json:"ADDR_FAM_ADV"`
	Addr_fam_rcv	string	`json:"ADDR_FAM_RCV"`
	Localhost_ip	string	`json:"LOCALHOST_IP"`
	Localhost_port	string	`json:"LOCALHOST_PORT"`
	Remote_ip	string	`json:"REMOTE_IP"`
	Remote_port	string	`json:"REMOTE_PORT"`
	Accepted_paths	string	`json:"ACCEPTED_PATHS"`
	Sent_paths	string	`json:"SENT_PATHS"`
	Consumed_mem	string	`json:"CONSUMED_MEM"`
	Rr_old_cap	string	`json:"RR_OLD_CAP"`
	Rr_new_cap	string	`json:"RR_NEW_CAP"`
	Graceful_cap	string	`json:"GRACEFUL_CAP"`
	Fourbyte_cap	string	`json:"FOURBYTE_CAP"`
	Ext_nh_cap	string	`json:"EXT_NH_CAP"`
	Restart_time_adv	string	`json:"RESTART_TIME_ADV"`
	Restart_time_rcv	string	`json:"RESTART_TIME_RCV"`
	Stale_time	string	`json:"STALE_TIME"`
	Opens_count_sent	string	`json:"OPENS_COUNT_SENT"`
	Notifications_count_sent	string	`json:"NOTIFICATIONS_COUNT_SENT"`
	Updates_count_sent	string	`json:"UPDATES_COUNT_SENT"`
	Keepalives_count_sent	string	`json:"KEEPALIVES_COUNT_SENT"`
	Route_refresh_count_sent	string	`json:"ROUTE_REFRESH_COUNT_SENT"`
	Capability_count_sent	string	`json:"CAPABILITY_COUNT_SENT"`
	Total_mess_count_sent	string	`json:"TOTAL_MESS_COUNT_SENT"`
	Total_bytes_count_sent	string	`json:"TOTAL_BYTES_COUNT_SENT"`
	Opens_count_rcvd	string	`json:"OPENS_COUNT_RCVD"`
	Notifications_count_rcvd	string	`json:"NOTIFICATIONS_COUNT_RCVD"`
	Updates_count_rcvd	string	`json:"UPDATES_COUNT_RCVD"`
	Keepalives_count_rcvd	string	`json:"KEEPALIVES_COUNT_RCVD"`
	Route_refresh_count_rcvd	string	`json:"ROUTE_REFRESH_COUNT_RCVD"`
	Capability_count_rcvd	string	`json:"CAPABILITY_COUNT_RCVD"`
	Total_mess_count_rcvd	string	`json:"TOTAL_MESS_COUNT_RCVD"`
	Total_bytes_count_rcvd	string	`json:"TOTAL_BYTES_COUNT_RCVD"`
	Total_bytes_send_queue	string	`json:"TOTAL_BYTES_SEND_QUEUE"`
	Total_bytes_rcvd_queue	string	`json:"TOTAL_BYTES_RCVD_QUEUE"`
	Inbound_routemap	string	`json:"INBOUND_ROUTEMAP"`
	Outbound_routemap	string	`json:"OUTBOUND_ROUTEMAP"`
}

var CiscoNxosShowIpBgpNeighbors_Template = `Value NEIGHBOR (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value PEER_GROUP (\S+)
Value DESCRIPTION (.*)
Value VRF (\S+)
Value ASN (\d+|\d+\.\d+)
Value REMOTE_ASN (\d+|\d+\.\d+)
Value LOCAL_ASN (\d+|\d+\.\d+)
Value BGP_STATE ([\(\)\w\s]+)
Value UPTIME ([\w:]+)
Value SOURCE_IFACE (\w+[\/\d+]+)
Value CONN_ESTAB (\w+)
Value CONN_DROPPED (\w+)
Value LAST_RESET (\S+)
Value LAST_RESET_REASON (.*)
Value LAST_PEER_RESET (\S+)
Value LAST_PEER_RESET_REASON (.*)
Value List ADDR_FAMILY ([\w\s]+)
Value List TABLE_VERSION (\d+)
Value List NEI_TABLE_VERSION (\d+)
Value ADDR_FAM_ADV ([\s\w]+$)
Value ADDR_FAM_RCV ([\s\w]+)
Value LOCALHOST_IP (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value LOCALHOST_PORT (\d+)
Value REMOTE_IP (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value REMOTE_PORT (\d+)
Value ACCEPTED_PATHS (\d+)
Value SENT_PATHS (\d+)
Value CONSUMED_MEM (\d+)
Value RR_OLD_CAP ([\w\s]+)
Value RR_NEW_CAP ([\w\s]+)
Value GRACEFUL_CAP ([\w\s]+)
Value FOURBYTE_CAP ([\w\s]+)
Value EXT_NH_CAP ([\w\s]+)
Value RESTART_TIME_ADV (\d+)
Value RESTART_TIME_RCV (\d+)
Value STALE_TIME (\d+)
Value OPENS_COUNT_SENT (\d+)
Value NOTIFICATIONS_COUNT_SENT (\d+)
Value UPDATES_COUNT_SENT (\d+)
Value KEEPALIVES_COUNT_SENT (\d+)
Value ROUTE_REFRESH_COUNT_SENT (\d+)
Value CAPABILITY_COUNT_SENT (\d+)
Value TOTAL_MESS_COUNT_SENT (\d+)
Value TOTAL_BYTES_COUNT_SENT (\d+)
Value OPENS_COUNT_RCVD (\d+)
Value NOTIFICATIONS_COUNT_RCVD (\d+)
Value UPDATES_COUNT_RCVD (\d+)
Value KEEPALIVES_COUNT_RCVD (\d+)
Value ROUTE_REFRESH_COUNT_RCVD (\d+)
Value CAPABILITY_COUNT_RCVD (\d+)
Value TOTAL_MESS_COUNT_RCVD (\d+)
Value TOTAL_BYTES_COUNT_RCVD (\d+)
Value TOTAL_BYTES_SEND_QUEUE (\d+)
Value TOTAL_BYTES_RCVD_QUEUE (\d+)
Value INBOUND_ROUTEMAP (\S+)
Value OUTBOUND_ROUTEMAP (\S+)

Start
  ^BGP\s+neighbor\s+is -> Continue.Record
  ^BGP\s+neighbor\s+is\s+${NEIGHBOR},\s+remote\s+AS\s+${ASN},.* -> Continue
  ^BGP\s+neighbor\s+is\s+${NEIGHBOR},\s+remote\s+AS\s+${REMOTE_ASN},\s+local\s+AS\s+${LOCAL_ASN}.*
  ^BGP\s+neighbor\s+is\s+${NEIGHBOR},\s+remote\s+AS\s+${REMOTE_ASN},.*
  ^\s+Inherits.+\s${PEER_GROUP}
  ^\s+Description:\s+${DESCRIPTION}
  ^\s+BGP state = ${BGP_STATE}, \w+ for ${UPTIME}
  ^\s+Neighbor\s+vrf:\s+${VRF}
  ^\s+Using ${SOURCE_IFACE} as [\w\s]+
  ^\s+[\w\s]+, interface ${SOURCE_IFACE}
  ^\s+Connections established ${CONN_ESTAB}, dropped ${CONN_DROPPED}
  ^\s+Last reset by us ${LAST_RESET}, due to ${LAST_RESET_REASON}
  ^\s+Last reset by peer ${LAST_PEER_RESET}, due to ${LAST_PEER_RESET_REASON}
  ^\s+Route refresh capability \(new\):\s+${RR_NEW_CAP}
  ^\s+Route refresh capability \(old\):\s+${RR_OLD_CAP}
  ^\s+4-Byte AS capability:\s+${FOURBYTE_CAP}
  ^\s+Graceful Restart capability:\s+${GRACEFUL_CAP}
  ^\s+Address families advertised to peer: -> AddrFamCap
  ^\s+Address families received from peer: -> AddrFamCap
  ^\s+Restart time advertised to peer:\s+${RESTART_TIME_ADV} seconds
  ^\s+Stale time for routes advertised by peer:\s+${STALE_TIME} seconds
  ^\s+Restart time advertised by peer:\s+${RESTART_TIME_RCV} seconds
  ^\s+Extended Next Hop Encoding Capability:\s+${EXT_NH_CAP}
  ^\s+Opens:\s+${OPENS_COUNT_SENT}\s+${OPENS_COUNT_RCVD}
  ^\s+Notifications:\s+${NOTIFICATIONS_COUNT_SENT}\s+${NOTIFICATIONS_COUNT_RCVD}
  ^\s+Updates:\s+${UPDATES_COUNT_SENT}\s+${UPDATES_COUNT_RCVD}
  ^\s+Keepalives:\s+${KEEPALIVES_COUNT_SENT}\s+${KEEPALIVES_COUNT_RCVD}
  ^\s+Route Refresh:\s+${ROUTE_REFRESH_COUNT_SENT}\s+${ROUTE_REFRESH_COUNT_RCVD}
  ^\s+Capability:\s+${CAPABILITY_COUNT_SENT}\s+${CAPABILITY_COUNT_RCVD}
  ^\s+Total:\s+${TOTAL_MESS_COUNT_SENT}\s+${TOTAL_MESS_COUNT_RCVD}
  ^\s+Total bytes:\s+${TOTAL_BYTES_COUNT_SENT}\s+${TOTAL_BYTES_COUNT_RCVD}
  ^\s+Bytes in queue:\s+${TOTAL_BYTES_SEND_QUEUE}\s+${TOTAL_BYTES_RCVD_QUEUE} -> AddrFamState

AddrFamCap
  ^\s+${ADDR_FAM_ADV} -> Start
  ^BGP\s+neighbor\s+is -> Continue.Record
  ^BGP\s+neighbor\s+is\s+${NEIGHBOR},\s+remote\s+AS\s+${ASN},.* -> Start

AddrFamState
  ^\s+For address family:\s+${ADDR_FAMILY}
  ^\s+BGP table version ${TABLE_VERSION}, neighbor version ${NEI_TABLE_VERSION}
  ^\s+${ACCEPTED_PATHS} accepted paths consume ${CONSUMED_MEM} bytes of memory
  ^\s+${SENT_PATHS} sent paths
  ^\s+Inbound\s+route-map\s+configured\s+is\s+${INBOUND_ROUTEMAP},\s+handle\s+obtained
  ^\s+Outbound\s+route-map\s+configured\s+is\s+${OUTBOUND_ROUTEMAP},\s+handle\s+obtained
  ^\s+Last End-of-RIB received [\d:]+ after session start -> AddrFamState
  ^\s+Local host:\s+${LOCALHOST_IP}, Local port:\s+${LOCALHOST_PORT}
  ^\s+Foreign host:\s+${REMOTE_IP}, Foreign port:\s+${REMOTE_PORT}
  ^\s+No\s+established\s+BGP\s+session\s+with\s+peer
  ^BGP\s+neighbor\s+is -> Continue.Record
  ^BGP\s+neighbor\s+is\s+${NEIGHBOR},\s+remote\s+AS\s+${ASN},.* -> Continue
  ^BGP\s+neighbor\s+is\s+${NEIGHBOR},\s+remote\s+AS\s+${REMOTE_ASN},\s+local\s+AS\s+${LOCAL_ASN}.* -> Start
  ^BGP\s+neighbor\s+is\s+${NEIGHBOR},\s+remote\s+AS\s+${REMOTE_ASN},.* -> Start

`