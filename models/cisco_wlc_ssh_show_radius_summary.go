package models

type CiscoWlcSshShowRadiusSummary struct {
	Vendor_id_backward_compat	string	`json:"VENDOR_ID_BACKWARD_COMPAT"`
	Call_station_id_case	string	`json:"CALL_STATION_ID_CASE"`
	Accounting_call_station_id_type	string	`json:"ACCOUNTING_CALL_STATION_ID_TYPE"`
	Authentication_call_station_id_type	string	`json:"AUTHENTICATION_CALL_STATION_ID_TYPE"`
	Extended_source_ports_support	string	`json:"EXTENDED_SOURCE_PORTS_SUPPORT"`
	Aggressive_failover	string	`json:"AGGRESSIVE_FAILOVER"`
	Keywrap	string	`json:"KEYWRAP"`
	Fallback_test_mode	string	`json:"FALLBACK_TEST_MODE"`
	Fallback_test_username	string	`json:"FALLBACK_TEST_USERNAME"`
	Fallback_test_interval	string	`json:"FALLBACK_TEST_INTERVAL"`
	Mac_delim_authentication_msg	string	`json:"MAC_DELIM_AUTHENTICATION_MSG"`
	Mac_delim_accounting_msg	string	`json:"MAC_DELIM_ACCOUNTING_MSG"`
	Authentication_framed_mtu	string	`json:"AUTHENTICATION_FRAMED_MTU"`
	Ap_events_accounting	string	`json:"AP_EVENTS_ACCOUNTING"`
	Index	string	`json:"INDEX"`
	Authentication_addr	string	`json:"AUTHENTICATION_ADDR"`
	Accounting_addr	string	`json:"ACCOUNTING_ADDR"`
	Type	string	`json:"TYPE"`
	Port	string	`json:"PORT"`
	State	string	`json:"STATE"`
	Timeout	string	`json:"TIMEOUT"`
	Mgmt_timeout	string	`json:"MGMT_TIMEOUT"`
	Rfc3576	string	`json:"RFC3576"`
	Ipsec_state	string	`json:"IPSEC_STATE"`
	Ipsec_profile	string	`json:"IPSEC_PROFILE"`
	Ipsec_region	string	`json:"IPSEC_REGION"`
}

var CiscoWlcSshShowRadiusSummary_Template = `Value Filldown VENDOR_ID_BACKWARD_COMPAT (\S+)
Value Filldown CALL_STATION_ID_CASE (\S+)
Value Filldown ACCOUNTING_CALL_STATION_ID_TYPE (.+)
Value Filldown AUTHENTICATION_CALL_STATION_ID_TYPE (.+)
Value Filldown EXTENDED_SOURCE_PORTS_SUPPORT (\S+)
Value Filldown AGGRESSIVE_FAILOVER (\S+)
Value Filldown KEYWRAP (\S+)
Value Filldown FALLBACK_TEST_MODE (\S+)
Value Filldown FALLBACK_TEST_USERNAME (.+)
Value Filldown FALLBACK_TEST_INTERVAL (\d+)
Value Filldown MAC_DELIM_AUTHENTICATION_MSG (\S+)
Value Filldown MAC_DELIM_ACCOUNTING_MSG (\S+)
Value Filldown AUTHENTICATION_FRAMED_MTU (\d+)
Value Filldown AP_EVENTS_ACCOUNTING (\S+)
Value Required INDEX (\d+)
Value AUTHENTICATION_ADDR (\S+)
Value ACCOUNTING_ADDR (\S+)
Value TYPE (\S+)
Value PORT (\d+)
Value STATE (\S+)
Value TIMEOUT (\d+)
Value MGMT_TIMEOUT (\d+)
Value RFC3576 (\S+)
Value IPSEC_STATE (\S+)
Value IPSEC_PROFILE (\S+)
Value IPSEC_REGION (\S+)

Start
  ^\s*Vendor Id Backward Compatibility\s*\.+\s+${VENDOR_ID_BACKWARD_COMPAT}\s*$$
  ^\s*Call Station I[dD] Case\s*\.+\s+${CALL_STATION_ID_CASE}\s*$$
  ^\s*Accounting Call Station I[dD] Type\s*\.+\s+${ACCOUNTING_CALL_STATION_ID_TYPE}\s*$$
  ^\s*Auth Call Station I[dD] Type\s*\.+\s+${AUTHENTICATION_CALL_STATION_ID_TYPE}\s*$$
  ^\s*Extended Source Ports Support\s*\.+\s+${EXTENDED_SOURCE_PORTS_SUPPORT}\s*$$
  ^\s*Aggressive Failover\s*\.+\s+${AGGRESSIVE_FAILOVER}\s*$$
  ^\s*Keywrap\s*\.+\s+${KEYWRAP}\s*$$
  ^\s*Fallback Test:\s*$$
  ^\s+Test Mode\s*\.+\s+${FALLBACK_TEST_MODE}\s*$$
  ^\s+Probe User Name\s*\.+\s+${FALLBACK_TEST_USERNAME}\s*$$
  ^\s+Interval \(in seconds\)\.+\s+${FALLBACK_TEST_INTERVAL}\s*$$
  ^\s*MAC Delimiter for Authentication Messages\s*\.+\s+${MAC_DELIM_AUTHENTICATION_MSG}\s*$$
  ^\s*MAC Delimiter for Accounting Messages\s*\.+\s+${MAC_DELIM_ACCOUNTING_MSG}\s*$$
  ^\s*RADIUS Authentication Framed-MTU\s*\.+\s+${AUTHENTICATION_FRAMED_MTU}\s[bB]ytes\s*$$
  ^\s*AP Events Accounting\s*\.+\s+${AP_EVENTS_ACCOUNTING}\s*$$
  ^\s*Authentication Servers\s*$$ -> AuthC_Servers
  ^\s*Accounting Servers\s*$$ -> Acct_Servers
  ^\s*$$
  ^. -> Error

AuthC_Servers
  ^\s*Idx\s+Type\s+Server Address\s+Port\s+State\s+Tout\s+MgmtTout\s+RFC3576\s+IPSec - state\/Profile Name\/RadiusRegionString\s*$$
  ^-+
  ^\s*${INDEX}\s+(\*)?\s+${TYPE}\s+${AUTHENTICATION_ADDR}\s+${PORT}\s+${STATE}\s+${TIMEOUT}\s+${MGMT_TIMEOUT}\s+${RFC3576}\s+${IPSEC_STATE}\s+${IPSEC_PROFILE}\s+\/${IPSEC_REGION}\s*$$ -> Record
  ^\s*$$
  ^\s*Accounting Servers\s*$$ -> Acct_Servers
  ^. -> Error

Acct_Servers
  ^\s*Idx\s+Type\s+Server Address\s+Port\s+State\s+Tout\s+MgmtTout\s+RFC3576\s+IPSec - state\/Profile Name\/RadiusRegionString\s*$$
  ^-+
  ^\s*${INDEX}\s+(\*)?\s+${TYPE}\s+${ACCOUNTING_ADDR}\s+${PORT}\s+${STATE}\s+${TIMEOUT}\s+${MGMT_TIMEOUT}\s+${RFC3576}\s+${IPSEC_STATE}\s+${IPSEC_PROFILE}\s+\/${IPSEC_REGION}\s*$$ -> Record
  ^\s*Authentication Servers\s*$$ -> AuthC_Servers
  ^\s*$$
  ^. -> Error

`