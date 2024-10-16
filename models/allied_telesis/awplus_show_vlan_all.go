package allied_telesis

type AwplusShowVlanAll struct {
	Vlan_id    string   `json:"VLAN_ID"`
	Name       string   `json:"NAME"`
	Status     string   `json:"STATUS"`
	Interfaces []string `json:"INTERFACES"`
}

var AwplusShowVlanAll_Template string = `Value Required VLAN_ID (\d+)
Value NAME (\S+)
Value STATUS (\S+)
Value List INTERFACES (\w+\.\d\.\d+\(\S\)|\S+\(\S\))

Start
  ^\s*VLAN\s+ID\s+Name\s+Type\s+State\s+Member\s+ports
  ^\s{8,}\(u\)-Untagged,\s+\(t\)-Tagged
  ^=== -> VLANS
  ^. -> Error

VLANS
  ^\d+ -> Continue.Record
  ^${VLAN_ID}\s+${NAME}\s+\w+\s+${STATUS}\s*$$
  ^${VLAN_ID}\s+${NAME}\s+\w+\s+${STATUS}\s+${INTERFACES}\s\S* -> Continue
  ^\d+\s+(?:\S+\s+){4}${INTERFACES}\s* -> Continue
  ^\d+\s+(?:\S+\s+){5}${INTERFACES}\s* -> Continue
  ^\d+\s+(?:\S+\s+){6}${INTERFACES}\s* -> Continue
  ^\d+\s+(?:\S+\s+){7}${INTERFACES}\s* -> Continue
  ^\d+\s+(?:\S+\s+){8}${INTERFACES}\s* -> Continue
  ^\d+\s+(?:\S+\s+){9}${INTERFACES}\s* -> Continue
  ^\d+\s+(?:\S+\s+){10}${INTERFACES}\s* -> Continue
  ^\d+\s+
  ^\s{8,}${INTERFACES}\s* -> Continue
  ^\s{8,}(?:\S+\s+){1}${INTERFACES}\s* -> Continue
  ^\s{8,}(?:\S+\s+){2}${INTERFACES}\s* -> Continue
  ^\s{8,}(?:\S+\s+){3}${INTERFACES}\s* -> Continue
  ^\s{8,}(?:\S+\s+){4}${INTERFACES}\s* -> Continue
  ^\s{8,}(?:\S+\s+){5}${INTERFACES}\s* -> Continue
  ^\s{8,}(?:\S+\s+){6}${INTERFACES}\s* -> Continue
  ^\s{8,}(?:\S+\s+){7}${INTERFACES}\s* -> Continue
  ^\s{8,}(?:\S+\s+){8}${INTERFACES}\s* -> Continue
  ^\s{8,}
  ^. -> Error
`
