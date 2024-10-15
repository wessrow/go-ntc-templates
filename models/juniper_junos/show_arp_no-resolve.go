package juniper_junos

type ShowArpNoResolve struct {
	Interface   string `json:"INTERFACE"`
	Flags       string `json:"FLAGS"`
	Mac_address string `json:"MAC_ADDRESS"`
	Ip_address  string `json:"IP_ADDRESS"`
}

var ShowArpNoResolve_Template string = `Value Required MAC_ADDRESS ([A-Fa-f0-9\:]{17})
Value Required IP_ADDRESS ([A-Fa-f0-9:\.]+)
Value Required INTERFACE (\S+)
Value FLAGS (\S+)

Start
  ^MAC\s+Address\s+Address\s+Interface\s+Flags\s*$$
  ^${MAC_ADDRESS}\s+${IP_ADDRESS}\s+${INTERFACE}\s+${FLAGS} -> Record
  ^Total.*
  ^\s*$$
  ^{master:\d+}
  ^. -> Error
`
