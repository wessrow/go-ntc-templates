package models

type CiscoIosShowDot1xAll struct {
	Sysauthcontrol	string	`json:"SYSAUTHCONTROL"`
	Dot1xversion	string	`json:"DOT1XVERSION"`
	Critical_recovery_delay	string	`json:"CRITICAL_RECOVERY_DELAY"`
	Critical_eapol	string	`json:"CRITICAL_EAPOL"`
	Interface	string	`json:"INTERFACE"`
	Pae	string	`json:"PAE"`
	Portcontrol	string	`json:"PORTCONTROL"`
	Controldirection	string	`json:"CONTROLDIRECTION"`
	Hostmode	string	`json:"HOSTMODE"`
	Reauthentication	string	`json:"REAUTHENTICATION"`
	Quietperiod	string	`json:"QUIETPERIOD"`
	Servertimeout	string	`json:"SERVERTIMEOUT"`
	Supptimeout	string	`json:"SUPPTIMEOUT"`
	Reauthperiod	string	`json:"REAUTHPERIOD"`
	Reauthmax	string	`json:"REAUTHMAX"`
	Masreq	string	`json:"MASREQ"`
	Txperiod	string	`json:"TXPERIOD"`
	Ratelimitperiod	string	`json:"RATELIMITPERIOD"`
}