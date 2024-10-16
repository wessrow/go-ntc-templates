package hp_comware

type DisplayVlanAll struct {
	Description     string `json:"DESCRIPTION"`
	Type            string `json:"TYPE"`
	Route_interface string `json:"ROUTE_INTERFACE"`
	Ipv4_address    string `json:"IPV4_ADDRESS"`
	Ipv4_subnet     string `json:"IPV4_SUBNET"`
	Vlan_id         string `json:"VLAN_ID"`
	Vlan_name       string `json:"VLAN_NAME"`
}

var DisplayVlanAll_Template string = `Value Required VLAN_ID (\d+)
Value Required VLAN_NAME (.*)
Value DESCRIPTION (.*)
Value TYPE (.*)
Value ROUTE_INTERFACE (.*)
Value IPV4_ADDRESS (\S+)
Value IPV4_SUBNET (\S+)

Start
  ^\s*VLAN\s+ID\s*: -> Continue.Record
  ^\s*VLAN\s+ID\s*:\s*${VLAN_ID}
  ^\s*VLAN\s+[Tt]ype\s*:\s*${TYPE}
  ^\s*Route\s+[Ii]nterface\s*:\s*${ROUTE_INTERFACE}
  ^\s*IPv4\s+[Aa]ddress\s*:\s*${IPV4_ADDRESS}
  ^\s*IPv4\s+[Ss]ubnet\s+mask\s*:\s*${IPV4_SUBNET}
  ^\s*Description\s*:\s*${DESCRIPTION}
  ^\s*Name\s*:\s*${VLAN_NAME}
  ^\s*Tagged\s+[Pp]orts
  ^\s*Untagged\s+[Pp]orts
  ^\s{3,}\S+\s*$$
  ^\s{3,}\S+\s+\S+\s*$$
  ^\s{3,}\S+\s+\S+\s+\S+\s*$$
  ^. -> Error
`
