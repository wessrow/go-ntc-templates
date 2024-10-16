package cisco_ios

type ShowDhcpLease struct {
	Server         string `json:"SERVER"`
	State_id       string `json:"STATE_ID"`
	State          string `json:"STATE"`
	Time_lease     string `json:"TIME_LEASE"`
	Time_rebound   string `json:"TIME_REBOUND"`
	Ip_address     string `json:"IP_ADDRESS"`
	Hostname       string `json:"HOSTNAME"`
	Interface      string `json:"INTERFACE"`
	Transaction_id string `json:"TRANSACTION_ID"`
	Time_renewal   string `json:"TIME_RENEWAL"`
	Time_next_fire string `json:"TIME_NEXT_FIRE"`
	Gateway        string `json:"GATEWAY"`
	Retry_count    string `json:"RETRY_COUNT"`
	Client_id      string `json:"CLIENT_ID"`
	Netmask        string `json:"NETMASK"`
}

var ShowDhcpLease_Template string = `Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value NETMASK (\d+\.\d+\.\d+\.\d+)
Value INTERFACE (\S+)
Value SERVER (\d+\.\d+\.\d+\.\d+)
Value STATE_ID (\d+)
Value STATE (\w+)
Value TRANSACTION_ID (\S+)
Value TIME_LEASE (\d+)
Value TIME_RENEWAL (\d+)
Value TIME_REBOUND (\d+)
Value TIME_NEXT_FIRE (\S+)
Value GATEWAY (\d+\.\d+\.\d+\.\d+)
Value RETRY_COUNT (\d+)
Value CLIENT_ID (\S+)
Value HOSTNAME (\S+)

Start
  ^\s*Temp\s+IP\s+addr:\s+${IP_ADDRESS}\s+for\s+peer\s+on\s+Interface: ${INTERFACE}
  ^\s*Temp\s+sub\s+net\s+mask:\s+${NETMASK}
  ^\s*DHCP\s+Lease\s+server:\s+${SERVER},\s+state:\s+${STATE_ID}\s+${STATE}
  ^\s*DHCP\s+transaction\s+id:\s+${TRANSACTION_ID}
  ^\s*Lease:\s+${TIME_LEASE}\s+secs,\s+Renewal:\s+${TIME_RENEWAL}\s+secs,\s+Rebind:\s+${TIME_REBOUND}\s+secs
  ^\s*Next\s+timer\s+fires\s+after:\s+${TIME_NEXT_FIRE}
  ^\s*Client\-ID\s+hex\s+dump.*$$
  ^\s*\S+$$
  ^\s*Temp\s+default-gateway\s+addr:\s+${GATEWAY}
  ^\s*Retry\s+count:\s+${RETRY_COUNT}\s+Client-ID:\s+${CLIENT_ID}
  ^\s*Hostname: ${HOSTNAME} -> Record
  ^\s*$$
  ^. -> Error
`
