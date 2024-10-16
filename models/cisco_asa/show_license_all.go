package cisco_asa

type ShowLicenseAll struct {
	License_reservation                      string `json:"LICENSE_RESERVATION"`
	Export_controlled                        string `json:"EXPORT_CONTROLLED"`
	Registration_initial_status              string `json:"REGISTRATION_INITIAL_STATUS"`
	Registration_last_renewal_date           string `json:"REGISTRATION_LAST_RENEWAL_DATE"`
	Privacy_version_privacy                  string `json:"PRIVACY_VERSION_PRIVACY"`
	Transport_type                           string `json:"TRANSPORT_TYPE"`
	Registration_last_renewal_error          string `json:"REGISTRATION_LAST_RENEWAL_ERROR"`
	License                                  string `json:"LICENSE"`
	License_pid                              string `json:"LICENSE_PID"`
	License_count                            string `json:"LICENSE_COUNT"`
	Pid                                      string `json:"PID"`
	Registration_expiration_date             string `json:"REGISTRATION_EXPIRATION_DATE"`
	Auth_last_communication_date             string `json:"AUTH_LAST_COMMUNICATION_DATE"`
	Auth_deadline                            string `json:"AUTH_DEADLINE"`
	License_export_status                    string `json:"LICENSE_EXPORT_STATUS"`
	Smart_account                            string `json:"SMART_ACCOUNT"`
	Registration_initial_date                string `json:"REGISTRATION_INITIAL_DATE"`
	Registration_last_renewal                string `json:"REGISTRATION_LAST_RENEWAL"`
	Privacy_smart_licensing_hostname_privacy string `json:"PRIVACY_SMART_LICENSING_HOSTNAME_PRIVACY"`
	Smart_status                             string `json:"SMART_STATUS"`
	Auth_date                                string `json:"AUTH_DATE"`
	Sn                                       string `json:"SN"`
	Smart_agent_version                      string `json:"SMART_AGENT_VERSION"`
	License_status                           string `json:"LICENSE_STATUS"`
	Registration_status                      string `json:"REGISTRATION_STATUS"`
	Virtual_account                          string `json:"VIRTUAL_ACCOUNT"`
	Registration_next_renewal_date           string `json:"REGISTRATION_NEXT_RENEWAL_DATE"`
	Auth_status                              string `json:"AUTH_STATUS"`
	Privacy_sending_hostname                 string `json:"PRIVACY_SENDING_HOSTNAME"`
	Auth_fail_error                          string `json:"AUTH_FAIL_ERROR"`
	Auth_next_communication_date             string `json:"AUTH_NEXT_COMMUNICATION_DATE"`
	License_version                          string `json:"LICENSE_VERSION"`
	Utility_status                           string `json:"UTILITY_STATUS"`
	Eval_remaining                           string `json:"EVAL_REMAINING"`
	Auth_last_communication                  string `json:"AUTH_LAST_COMMUNICATION"`
	License_description                      string `json:"LICENSE_DESCRIPTION"`
	Privacy_callhome_hostname_privacy        string `json:"PRIVACY_CALLHOME_HOSTNAME_PRIVACY"`
}

var ShowLicenseAll_Template string = `Value SMART_STATUS (.+?)
Value REGISTRATION_STATUS (.+?)
Value SMART_ACCOUNT (.+?)
Value VIRTUAL_ACCOUNT (.+?)
Value EXPORT_CONTROLLED (.+?)
Value REGISTRATION_INITIAL_STATUS (.+?)
Value REGISTRATION_INITIAL_DATE (.+?)
Value REGISTRATION_LAST_RENEWAL (.+?)
Value REGISTRATION_LAST_RENEWAL_DATE (.+?)
Value REGISTRATION_LAST_RENEWAL_ERROR (.+?)
Value REGISTRATION_NEXT_RENEWAL_DATE (.+?)
Value REGISTRATION_EXPIRATION_DATE (.+?)
Value AUTH_STATUS (.+?)
Value AUTH_DATE (.+?)
Value EVAL_REMAINING (.+?)
Value AUTH_LAST_COMMUNICATION (.+?)
Value AUTH_LAST_COMMUNICATION_DATE (.+?)
Value AUTH_FAIL_ERROR (.+?)
Value AUTH_NEXT_COMMUNICATION_DATE (.+?)
Value AUTH_DEADLINE (.+?)
Value LICENSE (\S.*?)
Value LICENSE_PID (\S+?)
Value LICENSE_DESCRIPTION (.+?)
Value LICENSE_COUNT (\d+)
Value LICENSE_VERSION (\S+)
Value LICENSE_STATUS (.+?)
Value LICENSE_EXPORT_STATUS (.+?)
Value PID (.+?)
Value SN (.+?)
Value SMART_AGENT_VERSION (.+?)
Value UTILITY_STATUS (.+?)
Value PRIVACY_SENDING_HOSTNAME (.+?)
Value PRIVACY_CALLHOME_HOSTNAME_PRIVACY (.+?)
Value PRIVACY_SMART_LICENSING_HOSTNAME_PRIVACY (.+?)
Value PRIVACY_VERSION_PRIVACY (.+?)
Value TRANSPORT_TYPE (.+?)
Value LICENSE_RESERVATION (.+?)

Start
  ^\s*$$
  ^Smart\s+Licensing\s+Status
  ^Smart\s+Licensing\s+is\s+${SMART_STATUS}\s*$$
  ^Registration:\s*$$ -> REGISTRATION
  ^License\s+Authorization:\s*$$ -> AUTHORIZATION
  ^License\s+Usage\s*$$ -> LICENSE_USAGE
  ^Product\s+Information\s*$$ -> PRODUCT_INFORMATION
  ^Agent\s+Version\s*$$ -> AGENT_VERSION
  ^Utility:\s*$$ -> UTILITY
  ^Data\s+Privacy:\s*$$ -> PRIVACY
  ^Transport:\s*$$ -> TRANSPORT
  ^No\s+licenses\s+in\s+use\s*$$
  ^=+\s*$$
  ^. -> Error

REGISTRATION
  ^\s*$$
  ^\s+Status:\s+${REGISTRATION_STATUS}\s*$$
  ^\s+Smart\s+Account:\s+${SMART_ACCOUNT}\s*$$
  ^\s+Virtual\s+Account:\s+${VIRTUAL_ACCOUNT}\s*$$
  ^\s+Export-Controlled\s+Functionality:\s+${EXPORT_CONTROLLED}\s*$$
  ^\s+Initial\s+Registration:\s+${REGISTRATION_INITIAL_STATUS}(?:\s+on\s+${REGISTRATION_INITIAL_DATE})?\s*$$
  ^\s+Last\s+Renewal\s+Attempt:\s+${REGISTRATION_LAST_RENEWAL}(?:\s+on\s+${REGISTRATION_LAST_RENEWAL_DATE})?\s*$$
  ^\s+Failure\s+reason:\s+${REGISTRATION_LAST_RENEWAL_ERROR}\s*$$
  ^\s+Next\s+Renewal\s+Attempt:\s+${REGISTRATION_NEXT_RENEWAL_DATE}\s*$$
  ^\s+Registration\s+Expires:\s+${REGISTRATION_EXPIRATION_DATE}\s*$$
  ^License\s+Authorization:\s*$$ -> AUTHORIZATION
  ^License\s+Usage\s*$$ -> LICENSE_USAGE
  ^Product\s+Information\s*$$ -> PRODUCT_INFORMATION
  ^Agent\s+Version\s*$$ -> AGENT_VERSION
  ^Smart\s+Licensing\s+is\s+${SMART_STATUS}\s*$$
  ^Registration:\s*$$ -> REGISTRATION
  ^=+\s*$$
  ^. -> Error

AUTHORIZATION
  ^\s*$$
  ^\s+Status:\s+${AUTH_STATUS}(?:\s+on\s+${AUTH_DATE})?\s*$$
  ^\s+Evaluation\s+Period\s+Remaining:\s+${EVAL_REMAINING}\s*$$
  ^\s+Last\s+Communication\s+Attempt:\s+${AUTH_LAST_COMMUNICATION}(?:\s+on\s+${AUTH_LAST_COMMUNICATION_DATE})?\s*$$
  ^\s+Failure\s+reason:\s+${AUTH_FAIL_ERROR}\s*$$
  ^\s+Next\s+Communication\s+Attempt:\s+${AUTH_NEXT_COMMUNICATION_DATE}\s*$$
  ^\s+Communication\s+Deadline:\s+${AUTH_DEADLINE}\s*$$
  ^License\s+Usage\s*$$ -> LICENSE_USAGE
  ^Product\s+Information\s*$$ -> PRODUCT_INFORMATION
  ^Agent\s+Version\s*$$ -> AGENT_VERSION
  ^Smart\s+Licensing\s+is\s+${SMART_STATUS}\s*$$
  ^Registration:\s*$$ -> REGISTRATION
  ^License\s+Authorization:\s*$$ -> AUTHORIZATION
  ^Utility:\s*$$ -> UTILITY
  ^Export\s+Authorization\s+Key:\s*$$ -> EXPORT_AUTHORIZATION_KEY
  ^=+\s*$$
  ^. -> Error

LICENSE_USAGE
  ^\s*$$
  ^${LICENSE}?\s*\(${LICENSE_PID}\):\s*$$
  ^\s+Description:(?:\s+${LICENSE_DESCRIPTION})?\s*$$
  ^\s+Count:\s+${LICENSE_COUNT}\s*$$
  ^\s+Version:\s+${LICENSE_VERSION}\s*$$
  ^\s+Status:\s+${LICENSE_STATUS}\s*$$
  ^\s+Export\s+status:\s+${LICENSE_EXPORT_STATUS}\s*$$
  ^Product\s+Information\s*$$ -> PRODUCT_INFORMATION
  ^Agent\s+Version\s*$$ -> AGENT_VERSION
  ^Smart\s+Licensing\s+is\s+${SMART_STATUS}\s*$$
  ^Registration:\s*$$ -> REGISTRATION
  ^License\s+Authorization:\s*$$ -> AUTHORIZATION
  ^License\s+Usage\s*$$ -> LICENSE_USAGE
  ^No\s+licenses\s+in\s+use\s*$$
  ^=+\s*$$
  ^. -> Error

PRODUCT_INFORMATION
  ^\s*$$
  ^UDI:.*PID:${PID},\s*SN:${SN}\s*$$ -> Start
  ^Agent\s+Version\s*$$ -> AGENT_VERSION
  ^Smart\s+Licensing\s+is\s+${SMART_STATUS}\s*$$
  ^Registration:\s*$$ -> REGISTRATION
  ^License\s+Authorization:\s*$$ -> AUTHORIZATION
  ^License\s+Usage\s*$$ -> LICENSE_USAGE
  ^Product\s+Information\s*$$ -> PRODUCT_INFORMATION
  ^=+\s*$$
  ^. -> Error

AGENT_VERSION
  ^\s*$$
  ^Smart\s+Agent\s+for\s+Licensing:\s+${SMART_AGENT_VERSION}\s*$$
  ^Smart\s+Licensing\s+is\s+${SMART_STATUS}\s*$$
  ^Registration:\s*$$ -> REGISTRATION
  ^License\s+Authorization:\s*$$ -> AUTHORIZATION
  ^License\s+Usage\s*$$ -> LICENSE_USAGE
  ^Product\s+Information\s*$$ -> PRODUCT_INFORMATION
  ^Agent\s+Version\s*$$ -> AGENT_VERSION
  ^Reservation\s+Info\s*$$ -> RESERVATION_INFO
  ^=+\s*$$
  ^. -> Error

RESERVATION_INFO
  ^\s*$$
  ^\s*License\s+reservation:\s+${LICENSE_RESERVATION}\s*$$
  ^=+\s*$$
  ^. -> Error

UTILITY
  ^\s*$$
  ^\s+Status:\s+${UTILITY_STATUS}\s*$$
  ^Data\s+Privacy:\s*$$ -> PRIVACY
  ^=+\s*$$
  ^. -> Error

PRIVACY
  ^\s*$$
  ^\s+Sending\s+Hostname:\s+${PRIVACY_SENDING_HOSTNAME}\s*$$
  ^\s+Callhome\s+hostname\s+privacy:\s+${PRIVACY_CALLHOME_HOSTNAME_PRIVACY}\s*$$
  ^\s+Smart\s+Licensing\s+hostname\s+privacy:\s+${PRIVACY_SMART_LICENSING_HOSTNAME_PRIVACY}\s*$$
  ^\s+Version\s+privacy:\s+${PRIVACY_VERSION_PRIVACY}\s*$$
  ^Transport:\s*$$ -> TRANSPORT
  ^=+\s*$$
  ^. -> Error

TRANSPORT
  ^\s*$$
  ^\s+Type:\s+${TRANSPORT_TYPE}\s*$$
  ^License\s+Usage\s*$$ -> LICENSE_USAGE
  ^=+\s*$$
  ^. -> Error

EXPORT_AUTHORIZATION_KEY
  ^\s*$$
  ^\s+Features\s+Authorized:\s*$$
  ^\s+<none>\s*$$
  ^License\s+Usage\s*$$ -> LICENSE_USAGE
  ^Product\s+Information\s*$$ -> PRODUCT_INFORMATION
  ^Agent\s+Version\s*$$ -> AGENT_VERSION
  ^Smart\s+Licensing\s+is\s+${SMART_STATUS}\s*$$
  ^Registration:\s*$$ -> REGISTRATION
  ^License\s+Authorization:\s*$$ -> AUTHORIZATION
  ^Utility:\s*$$ -> UTILITY
  ^=+\s*$$
  ^. -> Error
`
