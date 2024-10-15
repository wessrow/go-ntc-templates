package hp_procurve

type ShowVlans struct {
	Vlan_id   string `json:"VLAN_ID"`
	Vlan_name string `json:"VLAN_NAME"`
	Status    string `json:"STATUS"`
	Voice     string `json:"VOICE"`
	Jumbo     string `json:"JUMBO"`
}

var ShowVlans_Template string = `Value VLAN_ID (\d+)
Value VLAN_NAME (.*?)
Value STATUS (\S+)
Value VOICE (Yes|No)
Value JUMBO (Yes|No)

Start
  ^.*VLAN ID -> VLAN

VLAN
  ^\s+${VLAN_ID}\s+${VLAN_NAME}\s+(\|\s+|)${STATUS}\s+${VOICE}(\s+${JUMBO}|)\s*$$ -> Record
  ^\s+-+\s+-+\s+(\+\s+|)-+\s+-+(\s+-+|)\s*$$
  ^\s*$$
  ^.*$$ -> Error
`
