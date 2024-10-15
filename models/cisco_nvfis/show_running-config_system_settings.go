package cisco_nvfis

type ShowRunningConfigSystemSettings struct {
	Ip_address      string `json:"IP_ADDRESS"`
	Netmask         string `json:"NETMASK"`
	Default_gateway string `json:"DEFAULT_GATEWAY"`
	Hostname        string `json:"HOSTNAME"`
	Dpdk            string `json:"DPDK"`
}

var ShowRunningConfigSystemSettings_Template string = `Value HOSTNAME ([\w-]+)
Value DPDK (enable|disable)
Value IP_ADDRESS ([A-F0-9\.\:]+)
Value NETMASK ([A-F0-9\.\:]+)
Value DEFAULT_GATEWAY ([A-F0-9\.\:]+)


Start
  ^system\ssettings\shostname\s+${HOSTNAME}
  ^system\ssettings\sdpdk\s${DPDK}
  ^system\s+settings\s+mgmt\s+ip\s+address\s+${IP_ADDRESS}\s+${NETMASK}
  ^system\s+settings\s+default-gw\s+${DEFAULT_GATEWAY} -> Record

EOF`
