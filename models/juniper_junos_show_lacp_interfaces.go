package models

type JuniperJunosShowLacpInterfaces struct {
	Bundle_name	string	`json:"BUNDLE_NAME"`
	Member_interface	[]string	`json:"MEMBER_INTERFACE"`
	Receive_state	[]string	`json:"RECEIVE_STATE"`
	Transmit_state	[]string	`json:"TRANSMIT_STATE"`
	Mux_state	[]string	`json:"MUX_STATE"`
}

var JuniperJunosShowLacpInterfaces_Template = `Value Required BUNDLE_NAME (\S+)
Value List MEMBER_INTERFACE (\S+)
Value List RECEIVE_STATE (\S+\s?\w+)
Value List TRANSMIT_STATE (\S+\s?periodic)
Value List MUX_STATE (\S+\s?\S*)

Start
  ^Aggregated -> Continue.Record
  ^Aggregated\sinterface:\s${BUNDLE_NAME}(\s|$$)
  ^\s+LACP\s+state:
  ^\s+LACP\sprotocol: -> LACP_PROTO
  ^\s+\S+
  ^\s*$$
  ^{master:\d+}
  ^. -> Error

LACP_PROTO
  ^\s+${MEMBER_INTERFACE}\s+${RECEIVE_STATE}\s+${TRANSMIT_STATE}\s+${MUX_STATE}(\s|$$)
  ^\s*$$
  ^Aggregated -> Continue.Record
  ^Aggregated\sinterface:\s${BUNDLE_NAME}(\s|$$) -> Start
  ^{master:\d+}
  ^. -> Error

`