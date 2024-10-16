package arista_eos

type ShowPortChannelSummary struct {
	Bundle_name             string   `json:"BUNDLE_NAME"`
	Bundle_status           string   `json:"BUNDLE_STATUS"`
	Bundle_protocol         string   `json:"BUNDLE_PROTOCOL"`
	Bundle_protocol_state   string   `json:"BUNDLE_PROTOCOL_STATE"`
	Member_interface        []string `json:"MEMBER_INTERFACE"`
	Member_interface_status []string `json:"MEMBER_INTERFACE_STATUS"`
}

var ShowPortChannelSummary_Template string = `Value BUNDLE_NAME (Po\d+)
Value BUNDLE_STATUS (\S+)
Value BUNDLE_PROTOCOL (\w+)
Value BUNDLE_PROTOCOL_STATE (\S+)
Value List MEMBER_INTERFACE (Et\S+)
Value List MEMBER_INTERFACE_STATUS (\S+)


Start
  ^\s*Flags
  ^-+
  ^\s+\S\s+-
  ^Number
  ^\s+Port-Channel\s+Protocol\s+Ports\s+$$ -> PCS
  ^\s*$$
  ^. -> Error



PCS
  ^\s*Po\d+ -> Continue.Record
  ^\s*${BUNDLE_NAME}\(${BUNDLE_STATUS}\)\s+${BUNDLE_PROTOCOL}\(${BUNDLE_PROTOCOL_STATE}\)\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\s*Po\d+.+(Et.+?\(.+?\)\s){1}${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\s*Po\d+.+(Et.+?\(.+?\)\s){2}${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\s*Po\d+.+(Et.+?\(.+?\)\s){3}${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\s*Po\d+.+(Et.+?\(.+?\)\s){4}${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\s+(Et.+?\(.+?\)\s){1}${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\s+(Et.+?\(.+?\)\s){2}${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\s+(Et.+?\(.+?\)\s){3}${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\s+(Et.+?\(.+?\)\s){4}${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue

`
