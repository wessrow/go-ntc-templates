package cisco_apic

type FabricShowVlanExtended struct {
	Vlan_ports  []string `json:"VLAN_PORTS"`
	Node_number string   `json:"NODE_NUMBER"`
	Node_name   string   `json:"NODE_NAME"`
	Vlan_id     string   `json:"VLAN_ID"`
	Vlan_name   []string `json:"VLAN_NAME"`
	Vlan_encap  []string `json:"VLAN_ENCAP"`
}

var FabricShowVlanExtended_Template string = `Value Required VLAN_ID (\d+)
Value List VLAN_NAME ([\w:-]+)
Value List VLAN_ENCAP ([\w:-]+,?)
Value List VLAN_PORTS ((\S+,?\s?)+)
Value Filldown NODE_NUMBER (\d+)
Value Filldown NODE_NAME (\S+)
                     
Start
  ^\s*$$
  ^.+\sfabric\s
  ^---+$$
  ^\s*Node\s${NODE_NUMBER}\s\(${NODE_NAME}\)
  ^\s*VLAN\s+Name\s.+Ports.*$$ -> VLANS
  # Capture time-stamp if vty line has command time-stamping turned on
  ^. -> Error

VLANS
  ^ \d+ -> Continue.Record
  ^\s*--+\s+---
  ^\s{0,3}${VLAN_ID}\s+${VLAN_NAME}\s+${VLAN_ENCAP}\s+${VLAN_PORTS}\s*$$
  ^\s{4,6}${VLAN_NAME}\s+${VLAN_ENCAP}\s+${VLAN_PORTS}
  ^\s{4,6}${VLAN_NAME}\s*$$
  ^\s{38,40}${VLAN_ENCAP}\s*$$
  ^\s{54}\s+${VLAN_PORTS}\s*$$
  ^. -> Error

    `
