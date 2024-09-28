package models

type CiscoAsaShowVpnSessiondbAnyconnect struct {
	Session_type	string	`json:"SESSION_TYPE"`
	Username	string	`json:"USERNAME"`
	Index	string	`json:"INDEX"`
	Assigned_ip	string	`json:"ASSIGNED_IP"`
	Public_ip	string	`json:"PUBLIC_IP"`
	Assigned_ipv6	string	`json:"ASSIGNED_IPV6"`
	Protocol	string	`json:"PROTOCOL"`
	License	string	`json:"LICENSE"`
	Encryption	string	`json:"ENCRYPTION"`
	Hashing	string	`json:"HASHING"`
	Bytes_tx	string	`json:"BYTES_TX"`
	Bytes_rx	string	`json:"BYTES_RX"`
	Group_policy	string	`json:"GROUP_POLICY"`
	Tunnel_group	string	`json:"TUNNEL_GROUP"`
	Login_time	string	`json:"LOGIN_TIME"`
	Login_time_zone	string	`json:"LOGIN_TIME_ZONE"`
	Login_weekday	string	`json:"LOGIN_WEEKDAY"`
	Login_month	string	`json:"LOGIN_MONTH"`
	Login_day	string	`json:"LOGIN_DAY"`
	Login_year	string	`json:"LOGIN_YEAR"`
	Duration	string	`json:"DURATION"`
	Inactivity	string	`json:"INACTIVITY"`
	Vlan_mapping	string	`json:"VLAN_MAPPING"`
	Vlan_id	string	`json:"VLAN_ID"`
	Audt_sess_id	string	`json:"AUDT_SESS_ID"`
	Security_grp	string	`json:"SECURITY_GRP"`
}