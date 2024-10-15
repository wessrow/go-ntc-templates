package huawei_smartax

type DisplayServicePortAll struct {
	Flow_type   string `json:"FLOW_TYPE"`
	Tx          string `json:"TX"`
	Port_type   string `json:"PORT_TYPE"`
	Ont_id      string `json:"ONT_ID"`
	Vlan_attr   string `json:"VLAN_ATTR"`
	Fsp         string `json:"FSP"`
	Gem_index   string `json:"GEM_INDEX"`
	Vlan_access string `json:"VLAN_ACCESS"`
	Rx          string `json:"RX"`
	State       string `json:"STATE"`
	Index       string `json:"INDEX"`
	Vlan_id     string `json:"VLAN_ID"`
}

var DisplayServicePortAll_Template string = `Value INDEX (\d+)
Value VLAN_ID (\d+)
Value VLAN_ATTR (\S+)
Value PORT_TYPE (\S+)
Value FSP (\d+\/\d+\s*\/\d+)
Value ONT_ID (\d+)
Value GEM_INDEX (\d+)
Value FLOW_TYPE (\S+)
Value VLAN_ACCESS (\d+)
Value RX (\S+)
Value TX (\S+)
Value STATE (\S+)


Start
  ^Switch-Oriented\s*Flow\s*List
  ^\s*-
  ^\s*INDEX\s+VLAN\s+VLAN\s+PORT\s+F\/\s*S\/\s*P\s*VPI\s*VCI\s*FLOW\s*FLOW\s*RX\s*TX\s*STATE\s*$$
  ^\s*ID\s*ATTR\s*TYPE\s*TYPE\s*PARA\s*$$ -> ServicePorts
  ^\s*$$
  ^. -> Error

ServicePorts
  ^\s*${INDEX}\s*${VLAN_ID}\s*${VLAN_ATTR}\s*${PORT_TYPE}\s*${FSP}\s*${ONT_ID}\s*${GEM_INDEX}\s*${FLOW_TYPE}\s*${VLAN_ACCESS}\s*(-|${RX})\s*(-|${TX})\s*${STATE}\s* -> Record
  ^\s*.
  ^\s*$$
  ^. -> Error
`
