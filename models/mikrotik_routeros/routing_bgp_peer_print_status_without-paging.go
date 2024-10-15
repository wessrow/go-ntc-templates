package mikrotik_routeros

type RoutingBgpPeerPrintStatusWithoutPaging struct {
	Name                string `json:"NAME"`
	Out_filter          string `json:"OUT_FILTER"`
	Address_families    string `json:"ADDRESS_FAMILIES"`
	Remove_private_as   string `json:"REMOVE_PRIVATE_AS"`
	Use_bfd             string `json:"USE_BFD"`
	Remote_hold_time    string `json:"REMOTE_HOLD_TIME"`
	Index               string `json:"INDEX"`
	Passive             string `json:"PASSIVE"`
	Remote_id           string `json:"REMOTE_ID"`
	Withdrawn_received  string `json:"WITHDRAWN_RECEIVED"`
	Flag                string `json:"FLAG"`
	Prefix_count        string `json:"PREFIX_COUNT"`
	Withdrawn_sent      string `json:"WITHDRAWN_SENT"`
	Used_keepalive_time string `json:"USED_KEEPALIVE_TIME"`
	Local_address       string `json:"LOCAL_ADDRESS"`
	Remote_as           string `json:"REMOTE_AS"`
	Remote_address      string `json:"REMOTE_ADDRESS"`
	Route_reflect       string `json:"ROUTE_REFLECT"`
	Hold_time           string `json:"HOLD_TIME"`
	Update_source       string `json:"UPDATE_SOURCE"`
	Multihop            string `json:"MULTIHOP"`
	Updates_sent        string `json:"UPDATES_SENT"`
	Refresh_capability  string `json:"REFRESH_CAPABILITY"`
	Instance            string `json:"INSTANCE"`
	Ttl                 string `json:"TTL"`
	Default_originate   string `json:"DEFAULT_ORIGINATE"`
	As_override         string `json:"AS_OVERRIDE"`
	Used_hold_time      string `json:"USED_HOLD_TIME"`
	Tcp_md5_key         string `json:"TCP_MD5_KEY"`
	Keepalive_time      string `json:"KEEPALIVE_TIME"`
	In_filter           string `json:"IN_FILTER"`
	Uptime              string `json:"UPTIME"`
	Updates_received    string `json:"UPDATES_RECEIVED"`
	As4_capability      string `json:"AS4_CAPABILITY"`
	State               string `json:"STATE"`
	Nexthop_choice      string `json:"NEXTHOP_CHOICE"`
}

var RoutingBgpPeerPrintStatusWithoutPaging_Template string = `Value INDEX (\d+)
Value FLAG (X|E)
Value NAME (.*)
Value INSTANCE (\S+)
Value REMOTE_ADDRESS (\S+)
Value REMOTE_AS (\d+)
Value TCP_MD5_KEY (.*)
Value NEXTHOP_CHOICE (\S+)
Value MULTIHOP (yes|no)
Value ROUTE_REFLECT (yes|no)
Value HOLD_TIME (\S+)
Value KEEPALIVE_TIME (\S+)
Value TTL (\S+)
Value IN_FILTER (\S+)
Value OUT_FILTER (\S+)
Value ADDRESS_FAMILIES (\S+)
Value UPDATE_SOURCE (\S+)
Value DEFAULT_ORIGINATE (\S+)
Value REMOVE_PRIVATE_AS (yes|no)
Value AS_OVERRIDE (yes|no)
Value PASSIVE (yes|no)
Value USE_BFD (yes|no)
Value REMOTE_ID (\S+)
Value LOCAL_ADDRESS (\S+)
Value UPTIME (\S+)
Value PREFIX_COUNT (\d+)
Value UPDATES_SENT (\d+)
Value UPDATES_RECEIVED (\d+)
Value WITHDRAWN_SENT (\d+)
Value WITHDRAWN_RECEIVED (\d+)
Value REMOTE_HOLD_TIME (\S+)
Value USED_HOLD_TIME (\S+)
Value USED_KEEPALIVE_TIME (\S+)
Value REFRESH_CAPABILITY (yes|no)
Value AS4_CAPABILITY (yes|no)
Value STATE (\S+)

Start
  ^\s*Flags:\s+X\s+-\s+disabled,\s+E\s+-\s+established\s*$$
  ^\s*\d+\s+(X|E\s+)?name= -> Continue.Record
  ^\s*${INDEX}\s+(${FLAG}\s+)?name="${NAME}"\s+instance=${INSTANCE}\s+remote-address=${REMOTE_ADDRESS}\s+remote-as=${REMOTE_AS}(\s+tcp-md5-key=(?:\"?(\b${TCP_MD5_KEY}\b)?\"?))?(\s+nexthop-choice=${NEXTHOP_CHOICE})?(\s+multihop=${MULTIHOP})?(\s+route-reflect=${ROUTE_REFLECT})?(\s+hold-time=${HOLD_TIME})?(\s+keepalive-time=${KEEPALIVE_TIME})?(\s+ttl=${TTL})?(\s+in-filter=(?:\"?(\b${IN_FILTER}\b)?\"?))?(\s+out-filter=(?:\"?(\b${OUT_FILTER}\b)?\"?))?(\s+address-families=${ADDRESS_FAMILIES})?(\s+update-source=${UPDATE_SOURCE})?(\s+default-originate=${DEFAULT_ORIGINATE})?(\s+remove-private-as=${REMOVE_PRIVATE_AS})?(\s+as-override=${AS_OVERRIDE})?(\s+passive=${PASSIVE})?(\s+use-bfd=${USE_BFD})?(\s+remote-id=${REMOTE_ID})?(\s+local-address=${LOCAL_ADDRESS})?(\s+uptime=${UPTIME})?(\s+prefix-count=${PREFIX_COUNT})?(\s+updates-sent=${UPDATES_SENT})?(\s+updates-received=${UPDATES_RECEIVED})?(\s+withdrawn-sent=${WITHDRAWN_SENT})?(\s+withdrawn-received=${WITHDRAWN_RECEIVED})?(\s+remote-hold-time=${REMOTE_HOLD_TIME})?(\s+used-hold-time=${USED_HOLD_TIME})?(\s+used-keepalive-time=${USED_KEEPALIVE_TIME})?(\s+refresh-capability=${REFRESH_CAPABILITY})?(\s+as4-capability=${AS4_CAPABILITY})?(\s+state=${STATE})?\s*$$
  ^\s+updates-sent=${UPDATES_SENT}(\s+updates-received=${UPDATES_RECEIVED})?(\s+withdrawn-sent=${WITHDRAWN_SENT})?(\s+withdrawn-received=${WITHDRAWN_RECEIVED})?(\s+remote-hold-time=${REMOTE_HOLD_TIME})?(\s+used-hold-time=${USED_HOLD_TIME})?(\s+used-keepalive-time=${USED_KEEPALIVE_TIME})?(\s+refresh-capability=${REFRESH_CAPABILITY})?(\s+as4-capability=${AS4_CAPABILITY})?(\s+state=${STATE})?\s*$$
  ^\s+updates-received=${UPDATES_RECEIVED}(\s+withdrawn-sent=${WITHDRAWN_SENT})?(\s+withdrawn-received=${WITHDRAWN_RECEIVED})?(\s+remote-hold-time=${REMOTE_HOLD_TIME})?(\s+used-hold-time=${USED_HOLD_TIME})?(\s+used-keepalive-time=${USED_KEEPALIVE_TIME})?(\s+refresh-capability=${REFRESH_CAPABILITY})?(\s+as4-capability=${AS4_CAPABILITY})?(\s+state=${STATE})?\s*$$
  ^\s+withdrawn-sent=${WITHDRAWN_SENT}(\s+withdrawn-received=${WITHDRAWN_RECEIVED})?(\s+remote-hold-time=${REMOTE_HOLD_TIME})?(\s+used-hold-time=${USED_HOLD_TIME})?(\s+used-keepalive-time=${USED_KEEPALIVE_TIME})?(\s+refresh-capability=${REFRESH_CAPABILITY})?(\s+as4-capability=${AS4_CAPABILITY})?(\s+state=${STATE})?\s*$$
  ^\s+withdrawn-received=${WITHDRAWN_RECEIVED}(\s+remote-hold-time=${REMOTE_HOLD_TIME})?(\s+used-hold-time=${USED_HOLD_TIME})?(\s+used-keepalive-time=${USED_KEEPALIVE_TIME})?(\s+refresh-capability=${REFRESH_CAPABILITY})?(\s+as4-capability=${AS4_CAPABILITY})?(\s+state=${STATE})?\s*$$
  ^\s+remote-hold-time=${REMOTE_HOLD_TIME}(\s+used-hold-time=${USED_HOLD_TIME})?(\s+used-keepalive-time=${USED_KEEPALIVE_TIME})?(\s+refresh-capability=${REFRESH_CAPABILITY})?(\s+as4-capability=${AS4_CAPABILITY})?(\s+state=${STATE})?\s*$$
  ^\s+used-hold-time=${USED_HOLD_TIME}(\s+used-keepalive-time=${USED_KEEPALIVE_TIME})?(\s+refresh-capability=${REFRESH_CAPABILITY})?(\s+as4-capability=${AS4_CAPABILITY})?(\s+state=${STATE})?\s*$$
  ^\s+used-keepalive-time=${USED_KEEPALIVE_TIME}(\s+refresh-capability=${REFRESH_CAPABILITY})?(\s+as4-capability=${AS4_CAPABILITY})?(\s+state=${STATE})?\s*$$
  ^\s+refresh-capability=${REFRESH_CAPABILITY}(\s+as4-capability=${AS4_CAPABILITY})?(\s+state=${STATE})?\s*$$
  ^\s+as4-capability=${AS4_CAPABILITY}(\s+state=${STATE})?\s*$$
  ^\s+state=${STATE}\s*$$
  ^\s*$$
  ^. -> Error
`
