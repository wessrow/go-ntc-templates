package cisco_asa

type ShowPortChannelSummary struct {
	Member_interface_status []string `json:"MEMBER_INTERFACE_STATUS"`
	Group                   string   `json:"GROUP"`
	Bundle_name             string   `json:"BUNDLE_NAME"`
	Bundle_status           string   `json:"BUNDLE_STATUS"`
	Bundle_protocol         string   `json:"BUNDLE_PROTOCOL"`
	Member_interface        []string `json:"MEMBER_INTERFACE"`
}

var ShowPortChannelSummary_Template string = `Value Required GROUP (\d+)
Value BUNDLE_NAME (\S+)
Value BUNDLE_STATUS (\D+)
Value BUNDLE_PROTOCOL (\S+)
Value List MEMBER_INTERFACE ([\w\.\/]+)
Value List MEMBER_INTERFACE_STATUS (\D+)

Start
  ^Group\s+Port-channel\s+Protocol\s+Span-cluster\s+Ports -> PortChannel
  ^\s*$$

PortChannel
  ^\d+ -> Continue.Record
  ^${GROUP}\s+${BUNDLE_NAME}\(${BUNDLE_STATUS}\)\s+${BUNDLE_PROTOCOL}\s* -> Continue
  ^.+(-|LACP|PAgP)\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\) -> Continue
  ^.+(-|LACP|PAgP)\s+[\w\.\/\(\)]+\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\) -> Continue
  ^.+(-|LACP|PAgP)\s+[\w\.\/\(\)]+\s+[\w\.\/\(\)]+\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\) -> Continue
  ^.+(-|LACP|PAgP)\s+[\w\.\/\(\)]+\s+[\w\.\/\(\)]+\s+[\w\.\/\(\)]+\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\) -> Continue
  ^\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\) -> Continue
  ^\s+[\w\.\/\(\)]+\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\) -> Continue
  ^\s+[\w\.\/\(\)]+\s+[\w\.\/\(\)]+\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\) -> Continue
  ^\s+[\w\.\/\(\)]+\s+[\w\.\/\(\)]+\s+[\w\.\/\(\)]+\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\) -> Continue
  ^\d+
  ^\s+
  ^-+\++
  ^(RU|SU)\s+-\s+L(2|3)\s+port-channel\s+UP\s+(s|S)tate
  ^(P|S)/(bndl|susp)\s+-\s+(Bundled|Suspended)
  ^\s*$$
  ^. -> Error
  `
