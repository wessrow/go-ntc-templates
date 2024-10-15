package hp_comware

type DisplayIpVpnInstance struct {
	Name string `json:"NAME"`
	Rd   string `json:"RD"`
}

var DisplayIpVpnInstance_Template string = `Value Required NAME (\S+)
Value RD (\d+:\d+)

Start
  ^\s*VPN-Instance Name -> VPNInstances

VPNInstances
  ^\s*${NAME}\s+${RD}\s+.* -> Record
  ^\s*${NAME}\s+.* -> Record
  ^\s*$$
  ^. -> Error
`
