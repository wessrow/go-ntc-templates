package cisco_viptela

type ShowInterface struct {
	Encap_type     string `json:"ENCAP_TYPE"`
	Mtu            string `json:"MTU"`
	Vpn            string `json:"VPN"`
	Interface      string `json:"INTERFACE"`
	Oper_status    string `json:"OPER_STATUS"`
	Af_type        string `json:"AF_TYPE"`
	Admin_status   string `json:"ADMIN_STATUS"`
	Uptime         string `json:"UPTIME"`
	Ip_address     string `json:"IP_ADDRESS"`
	Tracker_status string `json:"TRACKER_STATUS"`
	Mss_adjust     string `json:"MSS_ADJUST"`
	Duplex         string `json:"DUPLEX"`
	Rx_packets     string `json:"RX_PACKETS"`
	Tx_packets     string `json:"TX_PACKETS"`
	Port_type      string `json:"PORT_TYPE"`
	Mac_address    string `json:"MAC_ADDRESS"`
	Speed          string `json:"SPEED"`
}

var ShowInterface_Template string = `Value VPN (\d+)
Value INTERFACE (\S+)
Value AF_TYPE (\S+)
Value IP_ADDRESS (\S+)
Value ADMIN_STATUS (\S+)
Value OPER_STATUS (\S+)
Value TRACKER_STATUS (\S+)
Value ENCAP_TYPE (\S+)
Value PORT_TYPE (\S+)
Value MTU (\S+)
Value MAC_ADDRESS (\S+)
Value SPEED (\S+)
Value DUPLEX (\S+)
Value MSS_ADJUST (\S+)
Value UPTIME (\S+)
Value RX_PACKETS (\S+)
Value TX_PACKETS (\S+)

Start
  ^\s*IF\s+IF\s+IF\s+TCP
  ^\s*AF\s+ADMIN\s+OPER\s+TRACKER\s+ENCAP\s+SPEED\s+MSS\s+RX\s+TX
  ^VPN\s+INTERFACE\s+TYPE\s+IP\s+ADDRESS\s+STATUS\s+STATUS\s+STATUS\s+TYPE\s+PORT\s+TYPE\s+MTU\s+HWADDR\s+MBPS\s+DUPLEX\s+ADJUST\s+UPTIME\s+PACKETS\s+PACKETS
  ^-+\s*$$
  ^\s*${VPN}\s+${INTERFACE}\s+${AF_TYPE}\s+${IP_ADDRESS}\s+${ADMIN_STATUS}\s+${OPER_STATUS}\s+${TRACKER_STATUS}\s+${ENCAP_TYPE}\s+${PORT_TYPE}\s+${MTU}\s+${MAC_ADDRESS}\s+${SPEED}\s+${DUPLEX}\s+${MSS_ADJUST}\s+${UPTIME}\s+${RX_PACKETS}\s+${TX_PACKETS}\s*$$ -> Record
  ^. -> Error`
