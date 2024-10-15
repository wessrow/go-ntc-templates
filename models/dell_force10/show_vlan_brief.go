package dell_force10

type ShowVlanBrief struct {
	Mac_aging  string `json:"MAC_AGING"`
	Ip_address string `json:"IP_ADDRESS"`
	Vlan_id    string `json:"VLAN_ID"`
	Vlan_name  string `json:"VLAN_NAME"`
	Stg        string `json:"STG"`
}

var ShowVlanBrief_Template string = `Value VLAN_ID (\d+)
Value VLAN_NAME (\S+(\s\S+)*)
Value STG (\d+)
Value MAC_AGING (\d+)
Value IP_ADDRESS ((?:[0-9]{1,3}\.){3}[0-9]{1,3}(/\d{1,2})?|\w+)

Start
  ^\s*VLAN\s+Name -> VLAN

VLAN
  ^\s*${VLAN_ID}\s+${VLAN_NAME}?\s+${STG}\s+${MAC_AGING}\s+${IP_ADDRESS} -> Record
`
