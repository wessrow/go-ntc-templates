package models

type CiscoIosShowLicenseStatus struct {
	Sl_enabled	string	`json:"SL_ENABLED"`
	Transport_type	string	`json:"TRANSPORT_TYPE"`
	Url	string	`json:"URL"`
	Proxy_address	string	`json:"PROXY_ADDRESS"`
	Proxy_port	string	`json:"PROXY_PORT"`
	Proxy_username	string	`json:"PROXY_USERNAME"`
	Proxy_password	string	`json:"PROXY_PASSWORD"`
	Vrf	string	`json:"VRF"`
	Reg_status	string	`json:"REG_STATUS"`
	Virt_acct	string	`json:"VIRT_ACCT"`
	Smart_acct	string	`json:"SMART_ACCT"`
	License_auth	string	`json:"LICENSE_AUTH"`
	Last_ack_received	string	`json:"LAST_ACK_RECEIVED"`
	Next_ack_deadline	string	`json:"NEXT_ACK_DEADLINE"`
	Rpt_push_interval	string	`json:"RPT_PUSH_INTERVAL"`
	Next_ack_push_check	string	`json:"NEXT_ACK_PUSH_CHECK"`
	Next_rpt_push	string	`json:"NEXT_RPT_PUSH"`
	Last_rpt_push	string	`json:"LAST_RPT_PUSH"`
	Last_rpt_file_write	string	`json:"LAST_RPT_FILE_WRITE"`
}