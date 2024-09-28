package models

type CiscoIosShowIpHttpServerStatus struct {
	Server_status	string	`json:"SERVER_STATUS"`
	Server_port	string	`json:"SERVER_PORT"`
	Auth_method	string	`json:"AUTH_METHOD"`
	Auth_retry	string	`json:"AUTH_RETRY"`
	Time_window	string	`json:"TIME_WINDOW"`
	Server_digest	string	`json:"SERVER_DIGEST"`
	Access_class	string	`json:"ACCESS_CLASS"`
	Ipv4_access_class	string	`json:"IPV4_ACCESS_CLASS"`
	Ipv6_access_class	string	`json:"IPV6_ACCESS_CLASS"`
	File_upload_status	string	`json:"FILE_UPLOAD_STATUS"`
	Concurrent_connections_allowed	string	`json:"CONCURRENT_CONNECTIONS_ALLOWED"`
	Secondary_connections_allowed	string	`json:"SECONDARY_CONNECTIONS_ALLOWED"`
	Server_idle_timeout	string	`json:"SERVER_IDLE_TIMEOUT"`
	Server_life_timeout	string	`json:"SERVER_LIFE_TIMEOUT"`
	Server_session_timeout	string	`json:"SERVER_SESSION_TIMEOUT"`
	Max_requests_per_connection	string	`json:"MAX_REQUESTS_PER_CONNECTION"`
	Server_linger_time	string	`json:"SERVER_LINGER_TIME"`
	Active_session_modules	string	`json:"ACTIVE_SESSION_MODULES"`
	Secure_server_capability	string	`json:"SECURE_SERVER_CAPABILITY"`
	Secure_server_status	string	`json:"SECURE_SERVER_STATUS"`
	Secure_server_port	string	`json:"SECURE_SERVER_PORT"`
	Secure_ciphersuite	[]string	`json:"SECURE_CIPHERSUITE"`
	Tls_versions	[]string	`json:"TLS_VERSIONS"`
	Secure_client_auth	string	`json:"SECURE_CLIENT_AUTH"`
	Secure_piv_authn	string	`json:"SECURE_PIV_AUTHN"`
	Secure_piv_authz	string	`json:"SECURE_PIV_AUTHZ"`
	Secure_ecdhe_curve	string	`json:"SECURE_ECDHE_CURVE"`
	Secure_active_session_modules	string	`json:"SECURE_ACTIVE_SESSION_MODULES"`
}