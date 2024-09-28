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