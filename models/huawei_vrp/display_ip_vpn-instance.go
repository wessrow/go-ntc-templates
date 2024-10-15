package huawei_vrp

type DisplayIpVpnInstance struct {
	Rd             string `json:"RD"`
	Address_family string `json:"ADDRESS_FAMILY"`
	Name           string `json:"NAME"`
}

var DisplayIpVpnInstance_Template string = `Value Required NAME (\S+)
Value RD (\d+:\d+)
Value ADDRESS_FAMILY (\S+)

Start
  ^\s*VPN-Instance Name -> VPNInstances
  ^\s*Total
  ^\s*$$
  ^. -> Error

VPNInstances
  ^\s*${NAME}\s+${RD}\s+${ADDRESS_FAMILY}\s* -> Record
  ^\s*$$
  ^. -> Error
`
