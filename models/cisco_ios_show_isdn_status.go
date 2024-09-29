package models

type CiscoIosShowIsdnStatus struct {
	Isdn_switchtype	string	`json:"ISDN_SWITCHTYPE"`
	Interface	string	`json:"INTERFACE"`
	L1_status	string	`json:"L1_STATUS"`
	Tei	string	`json:"TEI"`
	Ces	string	`json:"CES"`
	Sapi	string	`json:"SAPI"`
	L2_state	string	`json:"L2_STATE"`
	L3_active_calls	string	`json:"L3_ACTIVE_CALLS"`
	Active_dsl_ccbs	string	`json:"ACTIVE_DSL_CCBS"`
	Free_channel_mask	string	`json:"FREE_CHANNEL_MASK"`
	L2_discards	string	`json:"L2_DISCARDS"`
	L2_session_id	string	`json:"L2_SESSION_ID"`
}

var CiscoIosShowIsdnStatus_Template = `Value ISDN_SWITCHTYPE ([A-Za-z-]+)
Value Required INTERFACE ([A-Za-z0-9\/\:]+)
Value L1_STATUS (DEACTIVATED|SHUTDOWN|ACTIVE)
Value TEI (\d+)
Value CES (\d+)
Value SAPI (\d+)
Value L2_STATE ([A-Za-z_]+)
Value L3_ACTIVE_CALLS (\d+)
Value ACTIVE_DSL_CCBS (\d+)
Value FREE_CHANNEL_MASK (\w+)
Value L2_DISCARDS (\d+)
Value L2_SESSION_ID (\d+)


Start
  ^ISDN ${INTERFACE} interface
  ^\s+dsl \d+, interface ISDN Switchtype = ${ISDN_SWITCHTYPE}
  ^\s+Layer 1 Status:
  ^\s+${L1_STATUS}
  ^\s+Layer 2 Status:
  ^\s+TEI =\s+${TEI}, Ces =\s+${CES}, SAPI =\s+${SAPI}, State =\s+${L2_STATE}
  ^\s+Layer 3 Status:
  ^\s+${L3_ACTIVE_CALLS} Active Layer 3 Call\(s\)
  ^\s+Active dsl \d+ CCBs =\s+${ACTIVE_DSL_CCBS}
  ^\s+The Free Channel Mask:\s+${FREE_CHANNEL_MASK}
  ^\s+Number of L2 Discards =\s+${L2_DISCARDS}, L2 Session ID =\s+${L2_SESSION_ID} -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is

`