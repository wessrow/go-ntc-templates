package cisco_nxos

type ShowPortChannelSummary struct {
	Bundle_name             string   `json:"BUNDLE_NAME"`
	Bundle_status           string   `json:"BUNDLE_STATUS"`
	Bundle_protocol         string   `json:"BUNDLE_PROTOCOL"`
	Bundle_protocol_state   string   `json:"BUNDLE_PROTOCOL_STATE"`
	Member_interface        []string `json:"MEMBER_INTERFACE"`
	Member_interface_status []string `json:"MEMBER_INTERFACE_STATUS"`
}

var ShowPortChannelSummary_Template string = `Value Required,Filldown BUNDLE_NAME (Po\d+)
Value Filldown BUNDLE_STATUS (\w+)
Value Filldown BUNDLE_PROTOCOL (\w+)
Value Filldown BUNDLE_PROTOCOL_STATE (\w+)
Value List MEMBER_INTERFACE ([\w\/]+|\-\-)
Value List MEMBER_INTERFACE_STATUS (.+?)


Start
  ^\s+Flag -> CASE1
  ^Flags: -> CASE2


CASE1
  ^.*------------ -> CASE1_RTE

CASE1_RTE
  ^Po -> Continue.Record
  ^${BUNDLE_NAME}\(${BUNDLE_STATUS}\)\s+${BUNDLE_PROTOCOL}\(${BUNDLE_PROTOCOL_STATE}\)\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^Po\d+\(\w+\)\s+\w+\(\w+\)\s+Et.+?\(.+?\)\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^Po\d+\(\w+\)\s+\w+\(\w+\)\s+Et.+?\(.+?\)\s+Et.+?\(.+?\)\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^Po\d+\(\w+\)\s+\w+\(\w+\)\s+Et.+?\(.+?\)\s+Et.+?\(.+?\)\s+Et.+?\(.+?\)\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^Po\d+\(\w+\)\s+\w+\(\w+\)\s+Et.+?\(.+?\)\s+Et.+?\(.+?\)\s+Et.+?\(.+?\)\s+Et.+?\(.+?\)\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^Po\d+\(\w+\)\s+\w+\(\w+\)\s+Et.+?\(.+?\)\s+Et.+?\(.+?\)\s+Et.+?\(.+?\)\s+Et.+?\(.+?\)\s+Et.+?\(.+?\)\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\s{24}${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\s{24}Et.+?\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\s{24}Et.+?\s+Et.+?\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\s{24}Et.+?\s+Et.+?\s+Et.+?\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\s{24}Et.+?\s+Et.+?\s+Et.+?\s+Et.+?\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue


CASE2
  ^Group\s+Port -> CASE2_RTE

CASE2_RTE
  ^------------ -> Next
  ^\d+ -> Continue.Record
  ^\d+\s+${BUNDLE_NAME}\(${BUNDLE_STATUS}\)\s+\w+\s+${BUNDLE_PROTOCOL}\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\d+\s+Po\d+\(\w+\)\s+\w+\s+\w+\s+Et.+?\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\d+\s+Po\d+\(\w+\)\s+\w+\s+\w+\s+Et.+?\s+Et.+?\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\d+\s+Po\d+\(\w+\)\s+\w+\s+\w+\s+Et.+?\s+Et.+?\s+Et.+?\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\s{37}${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\s{37}Et.+?\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\s{37}Et.+?\s+Et.+?\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\s{37}Et.+?\s+Et.+?\s+Et.+?\s+${MEMBER_INTERFACE}\(${MEMBER_INTERFACE_STATUS}\)(\s|$$) -> Continue
  ^\d+\s+${BUNDLE_NAME}\(${BUNDLE_STATUS}\)\s+\w+\s+${BUNDLE_PROTOCOL}\s+${MEMBER_INTERFACE}($$) -> Next
`
