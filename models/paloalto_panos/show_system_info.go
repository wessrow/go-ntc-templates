package paloalto_panos

type ShowSystemInfo struct {
	Serial           string `json:"SERIAL"`
	Vm_license       string `json:"VM_LICENSE"`
	Vm_mode          string `json:"VM_MODE"`
	Operational_mode string `json:"OPERATIONAL_MODE"`
	Hostname         string `json:"HOSTNAME"`
	Ip_address       string `json:"IP_ADDRESS"`
	Family           string `json:"FAMILY"`
	Platform_family  string `json:"PLATFORM_FAMILY"`
	Vpn_disable_mode string `json:"VPN_DISABLE_MODE"`
	Mac_address      string `json:"MAC_ADDRESS"`
	Uptime           string `json:"UPTIME"`
	Netmask          string `json:"NETMASK"`
	Model            string `json:"MODEL"`
	Multi_vsys       string `json:"MULTI_VSYS"`
	Gateway          string `json:"GATEWAY"`
	Os               string `json:"OS"`
}

var ShowSystemInfo_Template string = `Value HOSTNAME (\S+)
Value IP_ADDRESS (\S+)
Value NETMASK (\S+)
Value GATEWAY (\S+)
Value MAC_ADDRESS (\S+)
Value UPTIME ([\S+\s+]+)
Value FAMILY (\S+)
Value MODEL (\S+)
Value SERIAL (\S+)
Value VM_LICENSE (\S+)
Value VM_MODE ([\S+\s+]+)
Value OS (\S+)
Value PLATFORM_FAMILY (\S+)
Value VPN_DISABLE_MODE (\S+)
Value MULTI_VSYS (\S+)
Value OPERATIONAL_MODE (\S+)

Start
  ^hostname:\s+${HOSTNAME}
  ^ip-address:\s+${IP_ADDRESS}
  ^netmask:\s+${NETMASK}
  ^default-gateway:\s+${GATEWAY}
  ^mac-address:\s+${MAC_ADDRESS}
  ^uptime:\s+${UPTIME}
  ^family:\s+${FAMILY}
  ^model:\s+${MODEL} 
  ^serial:\s+${SERIAL}
  ^vm-license:\s+${VM_LICENSE}
  ^vm-mode:\s+${VM_MODE}
  ^sw-version:\s+${OS}
  ^platform-family:\s+${PLATFORM_FAMILY}
  ^vpn-disable-mode:\s+${VPN_DISABLE_MODE}
  ^multi-vsys:\s+${MULTI_VSYS}
  ^operational-mode:\s+${OPERATIONAL_MODE} -> Record
`
