package cisco_ios

type ShowLicenseStatus struct {
	Next_ack_deadline   string `json:"NEXT_ACK_DEADLINE"`
	Last_rpt_push       string `json:"LAST_RPT_PUSH"`
	Transport_type      string `json:"TRANSPORT_TYPE"`
	Url                 string `json:"URL"`
	Proxy_address       string `json:"PROXY_ADDRESS"`
	Proxy_port          string `json:"PROXY_PORT"`
	Next_ack_push_check string `json:"NEXT_ACK_PUSH_CHECK"`
	Next_rpt_push       string `json:"NEXT_RPT_PUSH"`
	Last_rpt_file_write string `json:"LAST_RPT_FILE_WRITE"`
	Sl_enabled          string `json:"SL_ENABLED"`
	Proxy_username      string `json:"PROXY_USERNAME"`
	Reg_status          string `json:"REG_STATUS"`
	Virt_acct           string `json:"VIRT_ACCT"`
	Smart_acct          string `json:"SMART_ACCT"`
	License_auth        string `json:"LICENSE_AUTH"`
	Last_ack_received   string `json:"LAST_ACK_RECEIVED"`
	Rpt_push_interval   string `json:"RPT_PUSH_INTERVAL"`
	Proxy_password      string `json:"PROXY_PASSWORD"`
	Vrf                 string `json:"VRF"`
}

var ShowLicenseStatus_Template string = `Value SL_ENABLED (\S+)
Value TRANSPORT_TYPE (.*)
Value URL (\S+)
Value PROXY_ADDRESS (\S+)
Value PROXY_PORT (\d+)
Value PROXY_USERNAME (\S+)
Value PROXY_PASSWORD (\S+)
Value VRF (\S+)
Value REG_STATUS (\S+)
Value VIRT_ACCT (.*)
Value SMART_ACCT (.*?)
Value LICENSE_AUTH (.*)
Value LAST_ACK_RECEIVED (.*)
Value NEXT_ACK_DEADLINE (.*)
Value RPT_PUSH_INTERVAL (.*)
Value NEXT_ACK_PUSH_CHECK (.*)
Value NEXT_RPT_PUSH (.*)
Value LAST_RPT_PUSH (.*)
Value LAST_RPT_FILE_WRITE (.*)

Start
  ^Smart\s+Licensing\s+is\s+${SL_ENABLED}\s*$$
  ^Smart\s+Licensing\s+Using\s+Policy:\s*$$ -> UsingPolicy
  ^Transport:\s*$$ -> Transport
  ^Registration:\s*$$ -> Registration
  ^Account\s+Information:\s*$$ -> Registration
  ^License\s+Authorization:\s*$$ -> LicenseAuth
  ^Usage\s+Reporting:\s*$$ -> UsageReporting
  # match other sections
  ^\s*(\S+\s*)+$$
  ^\s*$$
  ^. -> Error

UsingPolicy
  ^\s+Status:\s+${SL_ENABLED}\s*$$
  ^\s*$$ -> Start
  ^. -> Error

Transport
  ^\s+Type:\s+${TRANSPORT_TYPE}\s*$$
  ^\s+URL:\s+${URL}\s*$$
  ^\s+Address:\s+${PROXY_ADDRESS}\s*$$
  ^\s+Port:\s+${PROXY_PORT}\s*$$
  ^\s+Username:\s+${PROXY_USERNAME}\s*$$
  ^\s+Password:\s+${PROXY_PASSWORD}\s*$$
  ^\s+VRF:\s+${VRF}\s*$$
  ^\s+(\S+\s*)+$$
  ^\s*$$ -> Start
  ^. -> Error

Registration
  ^\s+Status:\s+${REG_STATUS}\s*$$
  ^\s+Smart\s+Account:\s+${SMART_ACCT}(\s+As\s+of(\s+\S+)+)?\s*$$
  ^\s+Virtual\s+Account:\s+${VIRT_ACCT}\s*$$
  # match other lines
  ^\s+(\S+\s*)+$$
  ^\s*$$ -> Start
  ^. -> Error

LicenseAuth
  ^\s+Status:\s+${LICENSE_AUTH}(\s+\S+){6}$$
  # match other lines
  ^\s+(\S+\s*)+$$
  ^\s*$$ -> Start
  ^. -> Error

UsageReporting
  ^\s+Last\s+ACK\s+received:\s+${LAST_ACK_RECEIVED}\s*$$
  ^\s+Next\s+ACK\s+deadline:\s+${NEXT_ACK_DEADLINE}\s*$$
  ^\s+Reporting\s+push\s+interval:\s+${RPT_PUSH_INTERVAL}\s*$$
  ^\s+Next\s+ACK\s+push\s+check:\s+${NEXT_ACK_PUSH_CHECK}\s*$$
  ^\s+Next\s+report\s+push:\s+${NEXT_RPT_PUSH}\s*$$
  ^\s+Last\s+report\s+push:\s+${LAST_RPT_PUSH}\s*$$
  ^\s+Last\s+report\s+file\s+write:\s+${LAST_RPT_FILE_WRITE}\s*$$
  ^\s*$$ -> Start
  ^. -> Error
`
