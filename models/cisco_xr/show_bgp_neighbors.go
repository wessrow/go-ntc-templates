package cisco_xr 

type ShowBgpNeighbors struct {
	Neighbor	string	`json:"NEIGHBOR"`
	Remote_as	string	`json:"REMOTE_AS"`
	Local_as	string	`json:"LOCAL_AS"`
	Type	string	`json:"TYPE"`
	Description	string	`json:"DESCRIPTION"`
	Remote_router_id	string	`json:"REMOTE_ROUTER_ID"`
	Cluster_id	string	`json:"CLUSTER_ID"`
	State	string	`json:"STATE"`
	Uptime	string	`json:"UPTIME"`
	Nsr_state	string	`json:"NSR_STATE"`
	Holdtime	string	`json:"HOLDTIME"`
	Keepalive	string	`json:"KEEPALIVE"`
	Nsr	string	`json:"NSR"`
	Gr	string	`json:"GR"`
	Afi	[]string	`json:"AFI"`
	Version	[]string	`json:"VERSION"`
	Route_reflector_role	[]string	`json:"ROUTE_REFLECTOR_ROLE"`
	Policy_incoming	[]string	`json:"POLICY_INCOMING"`
	Policy_outgoing	[]string	`json:"POLICY_OUTGOING"`
	Prefixes_in	[]string	`json:"PREFIXES_IN"`
	Prefixes_in_best	[]string	`json:"PREFIXES_IN_BEST"`
	Prefixes_out	[]string	`json:"PREFIXES_OUT"`
	Prefixes_out_suppressed	[]string	`json:"PREFIXES_OUT_SUPPRESSED"`
	Prefixes_out_withdrawn	[]string	`json:"PREFIXES_OUT_WITHDRAWN"`
	Max_prefixes	[]string	`json:"MAX_PREFIXES"`
	Max_prefixes_warn_percent	[]string	`json:"MAX_PREFIXES_WARN_PERCENT"`
	Default_originate	[]string	`json:"DEFAULT_ORIGINATE"`
	Last_ack_version	[]string	`json:"LAST_ACK_VERSION"`
	Last_synced_ack_version	[]string	`json:"LAST_SYNCED_ACK_VERSION"`
	Connections_established	string	`json:"CONNECTIONS_ESTABLISHED"`
	Connections_dropped	string	`json:"CONNECTIONS_DROPPED"`
	Local_address	string	`json:"LOCAL_ADDRESS"`
	Local_port	string	`json:"LOCAL_PORT"`
	Remote_address	string	`json:"REMOTE_ADDRESS"`
	Remote_port	string	`json:"REMOTE_PORT"`
	Last_reset	string	`json:"LAST_RESET"`
	Last_reset_reason	string	`json:"LAST_RESET_REASON"`
	Last_notification_sent	string	`json:"LAST_NOTIFICATION_SENT"`
	Error_code	string	`json:"ERROR_CODE"`
	Last_notification_received	string	`json:"LAST_NOTIFICATION_RECEIVED"`
	Peer_error_code	string	`json:"PEER_ERROR_CODE"`
	Peer_reset_reason	string	`json:"PEER_RESET_REASON"`
	Max_hops	string	`json:"MAX_HOPS"`
}

var ShowBgpNeighbors_Template = `Value NEIGHBOR (\S+)
Value REMOTE_AS (\d+)
Value LOCAL_AS (\d+)
Value TYPE (\w+)
Value DESCRIPTION (.+?)
Value REMOTE_ROUTER_ID (\S+)
Value CLUSTER_ID (\S+)
Value STATE (.+?)
Value UPTIME (.+?)
Value NSR_STATE (.+?)
Value HOLDTIME (\d+)
Value KEEPALIVE (\d+)
Value NSR (.+?)
Value GR (.+?)
Value List AFI (.+?)
Value List VERSION (\d+)
Value List ROUTE_REFLECTOR_ROLE (\S+)
Value List POLICY_INCOMING (\S+)
Value List POLICY_OUTGOING (\S+)
Value List PREFIXES_IN (\d+)
Value List PREFIXES_IN_BEST (\d+)
Value List PREFIXES_OUT (\d+)
Value List PREFIXES_OUT_SUPPRESSED (\d+)
Value List PREFIXES_OUT_WITHDRAWN (\d+)
Value List MAX_PREFIXES (\d+)
Value List MAX_PREFIXES_WARN_PERCENT (\d+)
Value List DEFAULT_ORIGINATE (.+?)
Value List LAST_ACK_VERSION (\d+)
Value List LAST_SYNCED_ACK_VERSION (\d+)
Value CONNECTIONS_ESTABLISHED (\d+)
Value CONNECTIONS_DROPPED (\d+)
Value LOCAL_ADDRESS (.*)
Value LOCAL_PORT (\d+)
Value REMOTE_ADDRESS (.*)
Value REMOTE_PORT (\d+)
Value LAST_RESET (.+?)
Value LAST_RESET_REASON (.+?)
Value LAST_NOTIFICATION_SENT (\S+)
Value ERROR_CODE (.+?)
Value LAST_NOTIFICATION_RECEIVED (\S+)
Value PEER_ERROR_CODE (.+?)
Value PEER_RESET_REASON (.+?)
Value MAX_HOPS (\d+)


Start
  ^\S+\s+\S+\s+\d+\s+\d+:\d+:\d+\.\d+\s+\S+\s*$$
  ^BGP\s+neighbor -> Continue.Record
  ^BGP\s+neighbor\s+is\s+${NEIGHBOR}\s*$$
  ^\s+Remote\s+AS\s+${REMOTE_AS}\,\s+local\s+AS\s+${LOCAL_AS}\,\s+(?:no\-prepend,\s+)?(?:replace\-as\,\s+)?${TYPE}\s+link
  ^\s+Administratively\s+shut\s+down
  ^\s+Description:\s+${DESCRIPTION}\s*$$
  ^\s+Remote\s+router\s+ID\s+${REMOTE_ROUTER_ID}
  ^\s+Cluster\s+ID\s+${CLUSTER_ID}\s*$$
  ^\s+BGP\s+[Ss]tate\s+=\s+${STATE}(?:,\s+up\s+for\s+${UPTIME})?\s*$$
  ^\s+NSR\s+[Ss]tate:\s+(?:NSR\s+)?${NSR_STATE}\s*$$
  ^\s+BFD\s+[Ee]nabled
  ^\s+Last\s+[Rr]ead
  ^\s+Hold\s+time\s+is\s+${HOLDTIME},\s+keepalive\s+interval\s+is\s+${KEEPALIVE}
  ^\s+Configured hold time\:\s+\d+\,\s+keepalive\:
  ^\s+(?:Second\s+)?[Ll]ast\s+write
  ^\s+Socket\s+(?:not\s+)?armed
  ^\s+Last\s+KA
  ^\s+Precedence:
  ^\s+Non-stop\s+routing\s+is\s+${NSR}\s*$$
  ^\s+Graceful\s+restart\s+is\s+${GR}\s*$$
  ^\s+Restart\s+time\s+is
  ^\s+Stale\s+path\s+timeout\s+time\s+is
  ^\s+Enforcing\s+first\s+AS
  ^\s+Multi-protocol\s+capability
  ^\s+Neighbor\s+capabilities:
  ^\s+Graceful\s+Restart\s+\(GR\s+Awareness\)\:
  ^\s+Route\s+refresh:
  ^\s+4-byte\s+AS:
  ^\s+Address\s+family
  ^\s+Received\s+\d+\s+
  ^\s+Sent\s+\d+\s+messages
  ^\s+Inbound\s+message\s+logging
  ^\s+Outbound\s+message\s+logging
  ^\s+Minimum\s+time\s+between\s+advertisement
  ^\s+For\s+Address\s+Family:\s+${AFI}\s*$$ -> AFI
  ^\s+Connections\s+established\s+${CONNECTIONS_ESTABLISHED};\s+dropped\s+${CONNECTIONS_DROPPED}\s*$$ -> Connection
  ^\s*$$
  ^. -> Error

AFI
  ^\s+For\s+Address\s+Family:\s+${AFI}\s*$$
  ^\s+BGP\s+neighbor\s+version\s+${VERSION}\s*$$
  ^\s+Update\s+group:
  ^\s+Route-Reflector\s+${ROUTE_REFLECTOR_ROLE}\s*$$
  ^\s+Default\s+information\s+originate\:\s+${DEFAULT_ORIGINATE}\s*$$
  ^\s+AF-dependent capabilities
  ^\s+Graceful\s+Restart\s+capability
  ^\s+Local\s+restart\s+time\s+is
  ^\s+Maximum\s+stalepath\s+time\s+is
  ^\s+Remote\s+Restart\s+time\s+is
  ^\s+NEXT_HOP
  ^\s+Route\s+refresh\s+request:
  ^\s+Policy\s+for\s+incoming\s+advertisements\s+is\s+${POLICY_INCOMING}\s*$$
  ^\s+Policy\s+for\s+outgoing\s+advertisements\s+is\s+${POLICY_OUTGOING}\s*$$
  ^\s+${PREFIXES_IN}\s+accepted\s+prefixes,\s+${PREFIXES_IN_BEST}\s+are\s+bestpaths
  ^\s+Cumulative\s+no\.\s+of\s+prefixes\s+denied:
  ^\s+No\s+policy:
  ^\s+By\s+ORF\s+policy:
  ^\s+Prefix\s+advertised\s+${PREFIXES_OUT},\s+suppressed\s+${PREFIXES_OUT_SUPPRESSED},\s+withdrawn\s+${PREFIXES_OUT_WITHDRAWN}
  ^\s+Maximum\s+prefixes\s+allowed\s+${MAX_PREFIXES}
  ^\s+Threshold\s+for\s+warning\s+message\s+${MAX_PREFIXES_WARN_PERCENT}
  ^\s+AIGP
  ^\s+An\s+EoR
  ^\s+Private\s+AS\s+number\s+removed
  ^\s+Community\s+attribute
  ^\s+Advertise\s+VPNv[46]\s+routes\s+(is\s)?enabled\s+with
  ^\s+Last\s+ack\s+version\s+${LAST_ACK_VERSION},\s+Last\s+synced\s+ack\s+version\s+${LAST_SYNCED_ACK_VERSION}
  ^\s+Outstanding\s+version\s+objects:
  ^\s+Additional-paths\s+operation:
  ^\s+AS\s+[oO]verride\s+is
  ^\s+Send\s+Multicast\s+Attributes
  ^\s+Connections\s+established\s+${CONNECTIONS_ESTABLISHED};\s+dropped\s+${CONNECTIONS_DROPPED}\s*$$ -> Connection
  ^BGP\s+neighbor -> Continue.Record
  ^BGP\s+neighbor\s+is\s+${NEIGHBOR}\s*$$ -> Start
  ^\s*$$
  ^. -> Error

Connection
  ^\s+Local\s+host\:\s+${LOCAL_ADDRESS}\,\s+Local\s+port\:\s+${LOCAL_PORT}
  ^\s+Foreign\s+host\:\s+${REMOTE_ADDRESS}\,\s+Foreign\s+port\:\s+${REMOTE_PORT}
  ^\s+Last\s+reset\s+${LAST_RESET}(?:\,\s+due\s+to\s+${LAST_RESET_REASON})?\s*$$
  ^\s+External\s+BGP\s+neighbor\s+not\s+directly
  ^\s+Time\s+since\s+last\s+notification\s+sent\s+to\s+neighbor:\s+${LAST_NOTIFICATION_SENT}
  ^\s+Error\s+Code:\s+${ERROR_CODE}\s*$$
  ^\s+Notification\s+data\s+(?:sent|received):
  ^\s+(?:[0-9A-F]{8}|None)\s*$$
  ^\s+Time\s+since\s+last\s+notification\s+received\s+from\s+neighbor:\s+${LAST_NOTIFICATION_RECEIVED} -> PeerError
  ^\s+Peer\s+reset\s+reason:\s+${PEER_RESET_REASON}\s*$$
  ^\s+.+up\s+to\s+${MAX_HOPS}\s+hops\s+away
  ^BGP\s+neighbor -> Continue.Record
  ^BGP\s+neighbor\s+is\s+${NEIGHBOR}\s*$$ -> Start
  ^\s+For\s+Address\s+Family:\s+${AFI}\s*$$ -> AFI
  ^\s*$$
  ^. -> Error

PeerError
  ^\s+Error\s+Code:\s+${PEER_ERROR_CODE}\s*$$ -> Connection
  ^. -> Error
`