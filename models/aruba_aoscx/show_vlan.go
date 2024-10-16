package aruba_aoscx

type ShowVlan struct {
	Reason     string   `json:"REASON"`
	Type       string   `json:"TYPE"`
	Interfaces []string `json:"INTERFACES"`
	Vlan_id    string   `json:"VLAN_ID"`
	Vlan_name  string   `json:"VLAN_NAME"`
	Status     string   `json:"STATUS"`
}

var ShowVlan_Template string = `Value Required VLAN_ID (\d+)
Value VLAN_NAME (\S+)
Value STATUS (\S+)
Value REASON (\S+)
Value TYPE (\S+)
Value List INTERFACES ([^,]+)

Start
  ^---
  ^VLAN\s+Name\s+Status -> VLANS
  ^. -> Error

VLANS
  ^---
  ^\d+ -> Continue.Record
  # Lines starting with VLAN ID
  ^${VLAN_ID}\s+${VLAN_NAME}\s+${STATUS}\s+${REASON}\s+${TYPE}\s*$$
  ^${VLAN_ID}\s+${VLAN_NAME}\s+${STATUS}\s+${REASON}\s+${TYPE}\s+${INTERFACES},* -> Continue
  ^\d+\s+(?:\S+\s+){4}(?:[^,]+,){1}${INTERFACES},* -> Continue
  ^\d+\s+(?:\S+\s+){4}(?:[^,]+,){2}${INTERFACES},* -> Continue
  ^\d+\s+(?:\S+\s+){4}(?:[^,]+,){3}${INTERFACES},* -> Continue
  ^\d+\s+(?:\S+\s+){4}(?:[^,]+,){4}${INTERFACES},* -> Continue
  ^\d+\s+(?:\S+\s+){4}(?:[^,]+,){5}${INTERFACES},* -> Continue
  # Lines starting with multiple spaces
  ^\s{6,}${INTERFACES},* -> Continue
  ^\s{6,}(?:[^,]+,){1}${INTERFACES},* -> Continue
  ^\s{6,}(?:[^,]+,){2}${INTERFACES},* -> Continue
  ^\s{6,}(?:[^,]+,){3}${INTERFACES},* -> Continue
  ^\s{6,}(?:[^,]+,){4}${INTERFACES},* -> Continue
  ^\s{6,}(?:[^,]+,){5}${INTERFACES},* -> Continue
  # Dropping lines
  ^\d+\s+\S+\s+\S+\s+\S+\s+\S+\s*
  ^\s{6,}
  ^.+ -> Error
`
