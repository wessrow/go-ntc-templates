package avaya_ers

type ShowVlan struct {
	Vlan_id           string   `json:"VLAN_ID"`
	Vlan_name         string   `json:"VLAN_NAME"`
	Vlan_type         string   `json:"VLAN_TYPE"`
	Vlan_protocol     string   `json:"VLAN_PROTOCOL"`
	Vlan_active       string   `json:"VLAN_ACTIVE"`
	Vlan_port_members []string `json:"VLAN_PORT_MEMBERS"`
	Vlan_pid          string   `json:"VLAN_PID"`
	Vlan_ivl_svl      string   `json:"VLAN_IVL_SVL"`
	Vlan_mgmt         string   `json:"VLAN_MGMT"`
}

var ShowVlan_Template string = `Value VLAN_ID (\d+)
Value VLAN_NAME (\S.*?)
Value VLAN_TYPE (\S+)
Value VLAN_PROTOCOL (\S+)
Value VLAN_PID (\S+)
Value VLAN_ACTIVE (\S+)
Value VLAN_IVL_SVL (\S+)
Value VLAN_MGMT (\S+)
Value List VLAN_PORT_MEMBERS (\S+)

Start
  ^Total\s+VLAN_s: -> Done
  ^\d+ -> Continue.Record
  ^${VLAN_ID}\s+${VLAN_NAME}\s+${VLAN_TYPE}\s+${VLAN_PROTOCOL}\s+${VLAN_PID}\s+${VLAN_ACTIVE}\s+${VLAN_IVL_SVL}\s+${VLAN_MGMT}$$
  ^\t+Port Members:\s+${VLAN_PORT_MEMBERS}$$
  ^\t+\s{14}${VLAN_PORT_MEMBERS}$$


Done
`
