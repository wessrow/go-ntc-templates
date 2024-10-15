package cisco_viptela

type ShowArp struct {
	Mac_address string `json:"MAC_ADDRESS"`
	State       string `json:"STATE"`
	Idle_timer  string `json:"IDLE_TIMER"`
	Uptime      string `json:"UPTIME"`
	Vpn         string `json:"VPN"`
	Name        string `json:"NAME"`
	Ip_address  string `json:"IP_ADDRESS"`
}

var ShowArp_Template string = `Value VPN (\d+)
Value NAME (\S+)
Value IP_ADDRESS (\S+)
Value MAC_ADDRESS (\S+)
Value STATE (\S+)
Value IDLE_TIMER (\S+)
Value UPTIME (\S+)

Start
  ^\s*IF
  ^\s*VPN\s+(IF\s+)?NAME\s+IP\s+MAC\s+STATE\s+IDLE\s+TIMER\s+UPTIME
  ^-+
  ^\s*${VPN}\s+${NAME}\s+${IP_ADDRESS}\s+${MAC_ADDRESS}\s+${STATE}\s+${IDLE_TIMER}\s+${UPTIME} -> Record
  ^. -> Error
`
