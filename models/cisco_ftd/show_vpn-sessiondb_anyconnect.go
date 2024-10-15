package cisco_ftd

type ShowVpnSessiondbAnyconnect struct {
	Tunnel_group    string `json:"TUNNEL_GROUP"`
	Inactivity      string `json:"INACTIVITY"`
	Protocol        string `json:"PROTOCOL"`
	Bytes_tx        string `json:"BYTES_TX"`
	Login_month     string `json:"LOGIN_MONTH"`
	Security_grp    string `json:"SECURITY_GRP"`
	Tunnel_zone     string `json:"TUNNEL_ZONE"`
	Public_ip       string `json:"PUBLIC_IP"`
	Encryption      string `json:"ENCRYPTION"`
	Hashing         string `json:"HASHING"`
	Login_time      string `json:"LOGIN_TIME"`
	Assigned_ipv6   string `json:"ASSIGNED_IPV6"`
	Index           string `json:"INDEX"`
	Login_weekday   string `json:"LOGIN_WEEKDAY"`
	Login_day       string `json:"LOGIN_DAY"`
	Login_year      string `json:"LOGIN_YEAR"`
	Login_time_zone string `json:"LOGIN_TIME_ZONE"`
	Duration        string `json:"DURATION"`
	License         string `json:"LICENSE"`
	Group_policy    string `json:"GROUP_POLICY"`
	Vlan_mapping    string `json:"VLAN_MAPPING"`
	Username        string `json:"USERNAME"`
	Assigned_ip     string `json:"ASSIGNED_IP"`
	Session_type    string `json:"SESSION_TYPE"`
	Audt_sess_id    string `json:"AUDT_SESS_ID"`
	Bytes_rx        string `json:"BYTES_RX"`
	Vlan_id         string `json:"VLAN_ID"`
}

var ShowVpnSessiondbAnyconnect_Template string = `Value Filldown,Required SESSION_TYPE (\S+)
Value USERNAME (\S+)
Value Required INDEX (\d+)
Value ASSIGNED_IP (\d+\.\d+\.\d+\.\d+)
Value PUBLIC_IP (\S+)
Value ASSIGNED_IPV6 (\S+)
Value PROTOCOL (.+?)
Value LICENSE (.+?)
Value ENCRYPTION (.+?)
Value HASHING (.+?)
Value BYTES_TX (\d+)
Value BYTES_RX (\d+)
Value GROUP_POLICY (\S+)
Value TUNNEL_GROUP (\S+)
Value LOGIN_TIME (\d+:\d+:\d+)
Value LOGIN_TIME_ZONE (\S+)
Value LOGIN_WEEKDAY (\w+)
Value LOGIN_MONTH (\w+)
Value LOGIN_DAY (\d+)
Value LOGIN_YEAR (\d+)
Value DURATION (.+?)
Value INACTIVITY (.+?)
Value VLAN_MAPPING (\S+)
Value VLAN_ID (.+?)
Value AUDT_SESS_ID (.+?)
Value SECURITY_GRP (\S+)
Value TUNNEL_ZONE (\d+)

Start
  ^Session\s+Type:\s+${SESSION_TYPE}$$ -> Connection

Connection
  ^\s*Username\s*:\s+${USERNAME}\s+Index\s+:\s*${INDEX}$$
  ^\s*Username\s*:\s+${USERNAME}$$
  ^\s*Index\s*:\s+${INDEX}$$
  ^\s*Assigned\s+IP\s*:\s+${ASSIGNED_IP}\s+Public\s*IP\s*:\s*${PUBLIC_IP}$$
  ^\s*Assigned\s+IP\s*:\s+${ASSIGNED_IP}$$
  ^\s*Assigned\s+IPv6\s*:\s+${ASSIGNED_IPV6}$$
  ^\s*Public\s+IP\s*:\s+${PUBLIC_IP}$$
  ^\s*Protocol\s+:\s+${PROTOCOL}$$
  ^\s*License\s+:\s*${LICENSE}$$
  ^\s*Encryption\s+:\s*${ENCRYPTION}$$
  ^\s*Hashing\s+:\s*${HASHING}$$
  ^\s*Bytes\s+Tx\s+:\s+${BYTES_TX}\s+Bytes\s+Rx\s+:\s+${BYTES_RX}$$
  ^\s*Group\s+Policy\s+:\s+${GROUP_POLICY}\s+Tunnel\s+Group\s+:\s+${TUNNEL_GROUP}$$
  ^\s*Group\s+Policy\s+:\s+${GROUP_POLICY}$$
  ^\s*Tunnel\s+Group\s+:\s+${TUNNEL_GROUP}$$
  ^\s*Login\s+Time\s+:\s+${LOGIN_TIME}\s+${LOGIN_TIME_ZONE}\s+${LOGIN_WEEKDAY}\s+${LOGIN_MONTH}\s+${LOGIN_DAY}\s+${LOGIN_YEAR}$$
  ^\s*Duration\s+:\s+${DURATION}$$
  ^\s*Inactivity\s+:\s+${INACTIVITY}$$
  ^\s*VLAN\s+Mapping\s+:\s+${VLAN_MAPPING}\s+VLAN\s+:\s+${VLAN_ID}$$
  ^\s*Audt\s+Sess\s+ID\s+:\s+${AUDT_SESS_ID}$$
  ^\s*Security\s+Grp\s+:\s+${SECURITY_GRP}\s+Tunnel\s+Zone\s+:\s+${TUNNEL_ZONE}$$ -> Record
  ^\s*$$
  ^. -> Error
`
