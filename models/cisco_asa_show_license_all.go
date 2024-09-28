package models

type CiscoAsaShowLicenseAll struct {
	Smart_status	string	`json:"SMART_STATUS"`
	Registration_status	string	`json:"REGISTRATION_STATUS"`
	Smart_account	string	`json:"SMART_ACCOUNT"`
	Virtual_account	string	`json:"VIRTUAL_ACCOUNT"`
	Export_controlled	string	`json:"EXPORT_CONTROLLED"`
	Registration_initial_status	string	`json:"REGISTRATION_INITIAL_STATUS"`
	Registration_initial_date	string	`json:"REGISTRATION_INITIAL_DATE"`
	Registration_last_renewal	string	`json:"REGISTRATION_LAST_RENEWAL"`
	Registration_last_renewal_date	string	`json:"REGISTRATION_LAST_RENEWAL_DATE"`
	Registration_last_renewal_error	string	`json:"REGISTRATION_LAST_RENEWAL_ERROR"`
	Registration_next_renewal_date	string	`json:"REGISTRATION_NEXT_RENEWAL_DATE"`
	Registration_expiration_date	string	`json:"REGISTRATION_EXPIRATION_DATE"`
	Auth_status	string	`json:"AUTH_STATUS"`
	Auth_date	string	`json:"AUTH_DATE"`
	Eval_remaining	string	`json:"EVAL_REMAINING"`
	Auth_last_communication	string	`json:"AUTH_LAST_COMMUNICATION"`
	Auth_last_communication_date	string	`json:"AUTH_LAST_COMMUNICATION_DATE"`
	Auth_fail_error	string	`json:"AUTH_FAIL_ERROR"`
	Auth_next_communication_date	string	`json:"AUTH_NEXT_COMMUNICATION_DATE"`
	Auth_deadline	string	`json:"AUTH_DEADLINE"`
	License	string	`json:"LICENSE"`
	License_pid	string	`json:"LICENSE_PID"`
	License_description	string	`json:"LICENSE_DESCRIPTION"`
	License_count	string	`json:"LICENSE_COUNT"`
	License_version	string	`json:"LICENSE_VERSION"`
	License_status	string	`json:"LICENSE_STATUS"`
	License_export_status	string	`json:"LICENSE_EXPORT_STATUS"`
	Pid	string	`json:"PID"`
	Sn	string	`json:"SN"`
	Smart_agent_version	string	`json:"SMART_AGENT_VERSION"`
	Utility_status	string	`json:"UTILITY_STATUS"`
	Privacy_sending_hostname	string	`json:"PRIVACY_SENDING_HOSTNAME"`
	Privacy_callhome_hostname_privacy	string	`json:"PRIVACY_CALLHOME_HOSTNAME_PRIVACY"`
	Privacy_smart_licensing_hostname_privacy	string	`json:"PRIVACY_SMART_LICENSING_HOSTNAME_PRIVACY"`
	Privacy_version_privacy	string	`json:"PRIVACY_VERSION_PRIVACY"`
	Transport_type	string	`json:"TRANSPORT_TYPE"`
	License_reservation	string	`json:"LICENSE_RESERVATION"`
}