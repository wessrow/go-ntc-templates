package fortinet

type GetRouterInfoBgpNeighbors struct {
	Neighbor              string   `json:"NEIGHBOR"`
	Addr_family           []string `json:"ADDR_FAMILY"`
	Nei_table_version     []string `json:"NEI_TABLE_VERSION"`
	Accepted_prefixes_rib []string `json:"ACCEPTED_PREFIXES_RIB"`
	Localhost_ip          string   `json:"LOCALHOST_IP"`
	Remote_port           string   `json:"REMOTE_PORT"`
	Bgp_version           string   `json:"BGP_VERSION"`
	Uptime                string   `json:"UPTIME"`
	Accepted_prefixes     []string `json:"ACCEPTED_PREFIXES"`
	Announced_prefixes    []string `json:"ANNOUNCED_PREFIXES"`
	Remote_asn            string   `json:"REMOTE_ASN"`
	Local_asn             string   `json:"LOCAL_ASN"`
	Bgp_state             string   `json:"BGP_STATE"`
	Localhost_port        string   `json:"LOCALHOST_PORT"`
	Remote_ip             string   `json:"REMOTE_IP"`
	Nexthop_ip            string   `json:"NEXTHOP_IP"`
	Remote_router_id      string   `json:"REMOTE_ROUTER_ID"`
	Table_version         []string `json:"TABLE_VERSION"`
	Inbound_routemap      []string `json:"INBOUND_ROUTEMAP"`
	Outbound_routemap     []string `json:"OUTBOUND_ROUTEMAP"`
	Nexthop_interface     string   `json:"NEXTHOP_INTERFACE"`
}

var GetRouterInfoBgpNeighbors_Template string = `Value NEIGHBOR (\S+)
Value BGP_VERSION (\d+)
Value REMOTE_ASN (\d+)
Value LOCAL_ASN (\d+)
Value REMOTE_ROUTER_ID (\S+)
Value BGP_STATE ([\w\s\(\)]+)
Value UPTIME (\S+)
Value List ADDR_FAMILY (\w+(?:\s+\w+)?)
Value List TABLE_VERSION (\d+)
Value List NEI_TABLE_VERSION (\d+)
Value List INBOUND_ROUTEMAP (\S+)
Value List OUTBOUND_ROUTEMAP (\S+)
Value List ACCEPTED_PREFIXES (\d+)
Value List ACCEPTED_PREFIXES_RIB (\d+)
Value List ANNOUNCED_PREFIXES (\d+)
Value LOCALHOST_IP (\S+)
Value LOCALHOST_PORT (\d+)
Value REMOTE_IP (\S+)
Value REMOTE_PORT (\d+)
Value NEXTHOP_IP (\S+)
Value NEXTHOP_INTERFACE (\S+)

Start
  ^VRF\s+\d+\s+neighbor\s+table:\s*$$
  ^BGP\s+neighbor\s+is -> Continue.Record
  ^BGP\s+neighbor\s+is\s+${NEIGHBOR},\s+remote\s+AS\s+${REMOTE_ASN},\s+local\s+AS\s+${LOCAL_ASN},.*$$
  ^\s+BGP\s+version\s+${BGP_VERSION},\s+remote\s+router\s+ID\s+${REMOTE_ROUTER_ID}\s*$$
  ^\s+BGP\s+state\s+=\s+${BGP_STATE},\s+\w+\s+for\s+${UPTIME}\s*$$
  ^\s+Last\s+read\s+\d{2}:\d{2}:\d{2},\s+hold\s+time\s+is\s+\d+,\s+keepalive\s+interval\s+is\s+\d+\s+seconds\s*$$
  ^\s+Configured\s+hold\s+time\s+is\s+\d+,\s+keepalive\s+interval\s+is\s+\d+\s+seconds\s*$$
  ^\s+Neighbor\s+capabilities:\s*$$
  ^\s+Route\s+refresh:\s+.*$$
  ^\s+Address\s+family\s+IPv4\s+Unicast:\s+.*$$
  ^\s+Address\s+family\s+IPv6\s+Unicast:\s+.*$$
  ^\s+Received\s+\d+\s+messages,\s+\d+\s+notifications,\s+\d+\s+in\s+queue\s*$$
  ^\s+Sent\s+\d+\s+messages,\s+\d+\s+notifications,\s+\d+\s+in\s+queue\s*$$
  ^\s+Route\s+refresh\s+request:\s+received\s+\d+,\s+sent\s+\d+\s*$$
  ^\s+Minimum\s+time\s+between\s+advertisement\s+runs\s+is\s+\d+\s+seconds\s*$$
  ^\s+For\s+address\s+family:\s+${ADDR_FAMILY}\s*$$ -> AddrFamState
  ^\s+Connections\s+established\s+\d+;\s+dropped\s+\d+\s*$$
  ^Local\s+host:\s+${LOCALHOST_IP},\s+Local port:\s+${LOCALHOST_PORT}\s*$$
  ^Foreign\s+host:\s+${REMOTE_IP},\s+Foreign port:\s+${REMOTE_PORT}\s*$$
  ^Nexthop:\s+${NEXTHOP_IP}\s*$$
  ^Nexthop\s+interface:\s+${NEXTHOP_INTERFACE}\s*$$
  ^Nexthop\s+global:\s+.*$$
  ^Nexthop\s+local:\s+.*$$
  ^BGP\s+connection:\s+.*$$
  ^Last\s+Reset:\s+\S+,\s+.*$$
  ^Notification\s+Error\s+Message:\s+.*$$
  ^\s*\%\s+No\s+neighbor\s+exist\s*$$
  ^\s*$$
  ^. -> Error

AddrFamState
  ^\s+BGP\s+table\s+version\s+${TABLE_VERSION},\s+neighbor\s+version\s+${NEI_TABLE_VERSION}\s*$$
  ^\s+Index\s+\d+,\s+Offset\s+\d+,\s+Mask\s+\w+\s*$$
  ^\s+Additional\s+Path:\s*$$
  ^\s+Send-mode:\s+.*$$
  ^\s+Receive-mode:\s+.*$$
  ^\s+NEXT_HOP\s+is\s+always\s+this\s+router\s*$$
  ^\s+Community\s+attribute\s+sent\s+to\s+this\s+neighbor.*$$
  ^\s+Inbound\s+path\s+policy\s+configured\s*$$
  ^\s+Route\s+map\s+for\s+incoming\s+advertisements\s+is\s+${INBOUND_ROUTEMAP}\s*$$
  ^\s+Route\s+map\s+for\s+outgoing\s+advertisements\s+is\s+${OUTBOUND_ROUTEMAP}\s*$$
  ^\s+${ACCEPTED_PREFIXES}\saccepted\sprefixes,\s${ACCEPTED_PREFIXES_RIB}\sprefixes\sin\srib\s*$$
  ^\s+${ANNOUNCED_PREFIXES}\sannounced\sprefixes\s*$$ -> Start
  ^\s*$$
  ^. -> Error
`
