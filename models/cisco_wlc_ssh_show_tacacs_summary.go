package models

type CiscoWlcSshShowTacacsSummary struct {
	Fallback_test_interval	string	`json:"FALLBACK_TEST_INTERVAL"`
	Index	string	`json:"INDEX"`
	Authentication_addr	string	`json:"AUTHENTICATION_ADDR"`
	Authorization_addr	string	`json:"AUTHORIZATION_ADDR"`
	Accounting_addr	string	`json:"ACCOUNTING_ADDR"`
	Port	string	`json:"PORT"`
	State	string	`json:"STATE"`
	Timeout	string	`json:"TIMEOUT"`
	Mgmt_timeout	string	`json:"MGMT_TIMEOUT"`
}

var CiscoWlcSshShowTacacsSummary_Template = `Value Filldown FALLBACK_TEST_INTERVAL (\d+)
Value Required INDEX (\d+)
Value AUTHENTICATION_ADDR (\S+)
Value AUTHORIZATION_ADDR (\S+)
Value ACCOUNTING_ADDR (\S+)
Value PORT (\d+)
Value STATE (\S+)
Value TIMEOUT (\d+)
Value MGMT_TIMEOUT (\d+)

Start
  ^\s*Fallback Test:\s*$$
  ^\s+Interval \(in seconds\)\s*\.+\s+${FALLBACK_TEST_INTERVAL}\s*$$
  ^\s*Authentication Servers\s*$$ -> AuthC_Servers
  ^\s*Authorization Servers\s*$$ -> AuthZ_Servers
  ^\s*Accounting Servers\s*$$ -> Acct_Servers
  ^\s*$$
  ^. -> Error

AuthC_Servers
  ^\s*Idx\s+Server Address\s+Port\s+State\s+Tout\s+MgmtTout\s*$$
  ^-+
  ^\s*${INDEX}\s+${AUTHENTICATION_ADDR}\s+${PORT}\s+${STATE}\s+${TIMEOUT}\s+${MGMT_TIMEOUT}\s*$$ -> Record
  ^\s*Authorization Servers\s*$$ -> AuthZ_Servers
  ^\s*Accounting Servers\s*$$ -> Acct_Servers
  ^\s*$$
  ^. -> Error

AuthZ_Servers
  ^\s*Idx\s+Server Address\s+Port\s+State\s+Tout\s+MgmtTout\s*$$
  ^-+
  ^\s*${INDEX}\s+${AUTHORIZATION_ADDR}\s+${PORT}\s+${STATE}\s+${TIMEOUT}\s+${MGMT_TIMEOUT}\s*$$ -> Record
  ^\s*Authentication Servers\s*$$ -> AuthC_Servers
  ^\s*Accounting Servers\s*$$ -> Acct_Servers
  ^\s*$$
  ^. -> Error

Acct_Servers
  ^\s*Idx\s+Server Address\s+Port\s+State\s+Tout\s+MgmtTout\s*$$
  ^-+
  ^\s*${INDEX}\s+${ACCOUNTING_ADDR}\s+${PORT}\s+${STATE}\s+${TIMEOUT}\s+${MGMT_TIMEOUT}\s*$$ -> Record
  ^\s*Authentication Servers\s*$$ -> AuthC_Servers
  ^\s*Authorization Servers\s*$$ -> AuthZ_Servers
  ^\s*$$
  ^. -> Error

`