package allied_telesis

type AwplusShowInterfaceSwitchport struct {
	Interface      string   `json:"INTERFACE"`
	Mode           string   `json:"MODE"`
	Native_vlan    string   `json:"NATIVE_VLAN"`
	Trunking_vlans []string `json:"TRUNKING_VLANS"`
}

var AwplusShowInterfaceSwitchport_Template string = `Value Required INTERFACE (\S+)
Value MODE (\S+)
Value NATIVE_VLAN (\d+|None)
Value List TRUNKING_VLANS (\d+|\d)

Start
  ^\s*Interface\s+name -> Continue.Record
  ^\s*Interface\s+name\s*:\s*${INTERFACE}
  ^\s*Switchport\s+mode\s*:\s*${MODE}
  ^\s*Ingress\s+filter
  ^\s*Acceptable\s+frame\s+types
  ^\s*Default\s+Vlan\s*:\s*${NATIVE_VLAN}
  ^\s*Configured\s+Vlans\s*:\s*${TRUNKING_VLANS} -> Continue
  ^\s*Configured\s+Vlans\s*:\s*(?:\d+\s+){1}${TRUNKING_VLANS} -> Continue
  ^\s*Configured\s+Vlans\s*:\s*(?:\d+\s+){2}${TRUNKING_VLANS} -> Continue
  ^\s*Configured\s+Vlans\s*:\s*(?:\d+\s+){3}${TRUNKING_VLANS} -> Continue
  ^\s*Configured\s+Vlans\s*:\s*(?:\d+\s+){4}${TRUNKING_VLANS} -> Continue
  ^\s*Configured\s+Vlans\s*:\s*(?:\d+\s+){5}${TRUNKING_VLANS} -> Continue
  ^\s*Configured\s+Vlans\s*:\s*(?:\d+\s+){6}${TRUNKING_VLANS} -> Continue
  ^\s*Configured\s+Vlans\s*:\s*(?:\d+\s+){7}${TRUNKING_VLANS} -> Continue
  ^\s*Configured\s+Vlans\s*:\s*(?:\d+\s+){8}${TRUNKING_VLANS} -> Continue
  ^\s*Configured\s+Vlans\s*:\s*(?:\d+\s+){9}${TRUNKING_VLANS} -> Continue
  ^\s*Configured\s+Vlans\s*:\s*(?:\d+\s+){10}${TRUNKING_VLANS} -> Continue
  ^\s*Configured\s+Vlans
  ^\s*Translated\s+Vlans
  ^\s*Translated\s+Default
  ^\s{8,}${TRUNKING_VLANS} -> Continue
  ^\s{8,}(?:\d+\s+){1}${TRUNKING_VLANS} -> Continue
  ^\s{8,}(?:\d+\s+){2}${TRUNKING_VLANS} -> Continue
  ^\s{8,}(?:\d+\s+){3}${TRUNKING_VLANS} -> Continue
  ^\s{8,}(?:\d+\s+){4}${TRUNKING_VLANS} -> Continue
  ^\s{8,}(?:\d+\s+){5}${TRUNKING_VLANS} -> Continue
  ^\s{8,}(?:\d+\s+){6}${TRUNKING_VLANS} -> Continue
  ^\s{8,}(?:\d+\s+){7}${TRUNKING_VLANS} -> Continue
  ^\s{8,}(?:\d+\s+){8}${TRUNKING_VLANS} -> Continue
  ^\s{8,}(?:\d+\s+){9}${TRUNKING_VLANS} -> Continue
  ^\s{8,}(?:\d+\s+){10}${TRUNKING_VLANS} -> Continue
  ^\s{8,}
  ^\s*Dynamic\s+Vlans
  ^. -> Error
`
