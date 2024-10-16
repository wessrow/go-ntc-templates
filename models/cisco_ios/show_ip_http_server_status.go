package cisco_ios

type ShowIpHttpServerStatus struct {
	Ipv4_access_class              string   `json:"IPV4_ACCESS_CLASS"`
	Server_life_timeout            string   `json:"SERVER_LIFE_TIMEOUT"`
	Secure_client_auth             string   `json:"SECURE_CLIENT_AUTH"`
	Server_session_timeout         string   `json:"SERVER_SESSION_TIMEOUT"`
	Server_linger_time             string   `json:"SERVER_LINGER_TIME"`
	Secure_server_port             string   `json:"SECURE_SERVER_PORT"`
	Secure_ciphersuite             []string `json:"SECURE_CIPHERSUITE"`
	Secure_piv_authn               string   `json:"SECURE_PIV_AUTHN"`
	Secure_server_capability       string   `json:"SECURE_SERVER_CAPABILITY"`
	File_upload_status             string   `json:"FILE_UPLOAD_STATUS"`
	Secondary_connections_allowed  string   `json:"SECONDARY_CONNECTIONS_ALLOWED"`
	Server_idle_timeout            string   `json:"SERVER_IDLE_TIMEOUT"`
	Secure_active_session_modules  string   `json:"SECURE_ACTIVE_SESSION_MODULES"`
	Max_requests_per_connection    string   `json:"MAX_REQUESTS_PER_CONNECTION"`
	Secure_piv_authz               string   `json:"SECURE_PIV_AUTHZ"`
	Secure_ecdhe_curve             string   `json:"SECURE_ECDHE_CURVE"`
	Auth_method                    string   `json:"AUTH_METHOD"`
	Auth_retry                     string   `json:"AUTH_RETRY"`
	Active_session_modules         string   `json:"ACTIVE_SESSION_MODULES"`
	Time_window                    string   `json:"TIME_WINDOW"`
	Ipv6_access_class              string   `json:"IPV6_ACCESS_CLASS"`
	Concurrent_connections_allowed string   `json:"CONCURRENT_CONNECTIONS_ALLOWED"`
	Tls_versions                   []string `json:"TLS_VERSIONS"`
	Server_port                    string   `json:"SERVER_PORT"`
	Server_digest                  string   `json:"SERVER_DIGEST"`
	Secure_server_status           string   `json:"SECURE_SERVER_STATUS"`
	Server_status                  string   `json:"SERVER_STATUS"`
	Access_class                   string   `json:"ACCESS_CLASS"`
}

var ShowIpHttpServerStatus_Template string = `Value SERVER_STATUS (\S+)
Value SERVER_PORT (\d+)
Value AUTH_METHOD (\S+)
Value AUTH_RETRY (\d+)
Value TIME_WINDOW (\d+)
Value SERVER_DIGEST (\S+)
Value ACCESS_CLASS (\S+)
Value IPV4_ACCESS_CLASS (\S+)
Value IPV6_ACCESS_CLASS (\S+)
Value FILE_UPLOAD_STATUS (\S+)
Value CONCURRENT_CONNECTIONS_ALLOWED (\d+)
Value SECONDARY_CONNECTIONS_ALLOWED (\d+)
Value SERVER_IDLE_TIMEOUT (.*)
Value SERVER_LIFE_TIMEOUT (.*)
Value SERVER_SESSION_TIMEOUT (.*)
Value MAX_REQUESTS_PER_CONNECTION (\d+)
Value SERVER_LINGER_TIME (.*)
Value ACTIVE_SESSION_MODULES (\S+)
Value SECURE_SERVER_CAPABILITY (\S+)
Value SECURE_SERVER_STATUS (\S+)
Value SECURE_SERVER_PORT (\d+)
Value List SECURE_CIPHERSUITE (\S+)
Value List TLS_VERSIONS (\S+)
Value SECURE_CLIENT_AUTH (\S+)
Value SECURE_PIV_AUTHN (\S+)
Value SECURE_PIV_AUTHZ (\S+)
Value SECURE_ECDHE_CURVE (\S+)
Value SECURE_ACTIVE_SESSION_MODULES (.+)

				  
Start
  ^HTTP\s+server\s+status:\s+${SERVER_STATUS}
  ^HTTP\s+server\s+port:\s+${SERVER_PORT}
  ^HTTP\s+server\s+active\s+supplementary\s+listener\s+ports:
  ^HTTP\s+server\s+authentication\s+method:\s+${AUTH_METHOD}
  ^HTTP\s+server\s+auth-retry\s+${AUTH_RETRY}\s+time-window\s+${TIME_WINDOW}
  ^HTTP\s+server\s+digest\s+algorithm:\s+${SERVER_DIGEST}
  ^HTTP\s+server\s+access\s+class:\s+${ACCESS_CLASS}
  ^HTTP\s+server\s+IPv4\s+access\s+class:\s+${IPV4_ACCESS_CLASS}
  ^HTTP\s+server\s+IPv6\s+access\s+class:\s+${IPV6_ACCESS_CLASS}
  ^HTTP\s+server\s+base\s+path:
  ^HTTP\s+File\s+Upload\s+status:\s+${FILE_UPLOAD_STATUS}
  ^HTTP\s+server\s+upload\s+path:
  ^HTTP\s+server\s+help\s+root:
  ^Maximum\s+number\s+of\s+concurrent\s+server\s+connections\s+allowed:\s+${CONCURRENT_CONNECTIONS_ALLOWED}
  ^Maximum\s+number\s+of\s+secondary\s+server\s+connections\s+allowed:\s+${SECONDARY_CONNECTIONS_ALLOWED}
  ^Server\s+idle\s+time-out:\s+${SERVER_IDLE_TIMEOUT}
  ^Server\s+life\s+time-out:\s+${SERVER_LIFE_TIMEOUT}
  ^Server\s+session\s+idle\s+time-out:\s+${SERVER_SESSION_TIMEOUT}
  ^Maximum\s+number\s+of\s+requests\s+allowed\s+on\s+a\s+connection:\s+${MAX_REQUESTS_PER_CONNECTION}
  ^Server\s+linger\s+time\s+:\s+${SERVER_LINGER_TIME}
  ^HTTP\s+server\s+active\s+session\s+modules:\s+${ACTIVE_SESSION_MODULES}
  ^HTTP\s+secure\s+server\s+capability:\s+${SECURE_SERVER_CAPABILITY}
  ^HTTP\s+secure\s+server\s+status:\s+${SECURE_SERVER_STATUS}
  ^HTTP\s+secure\s+server\s+port:\s+${SECURE_SERVER_PORT}
  ^HTTP\s+secure\s+server\s+ciphersuite:\s+${SECURE_CIPHERSUITE}\s*$$
  ^HTTP\s+secure\s+server\s+ciphersuite:\s+${SECURE_CIPHERSUITE} -> Continue
  ^HTTP\s+secure\s+server\s+ciphersuite:\s+(?:\S+\s+)${SECURE_CIPHERSUITE}\s*$$
  ^HTTP\s+secure\s+server\s+ciphersuite:\s+(?:\S+\s+)${SECURE_CIPHERSUITE} -> Continue
  ^HTTP\s+secure\s+server\s+ciphersuite:\s+(?:\S+\s+){2}${SECURE_CIPHERSUITE}\s*$$
  ^HTTP\s+secure\s+server\s+ciphersuite:\s+(?:\S+\s+){2}${SECURE_CIPHERSUITE} -> Continue
  ^HTTP\s+secure\s+server\s+ciphersuite:\s+(?:\S+\s+){3}${SECURE_CIPHERSUITE}\s*$$
  ^HTTP\s+secure\s+server\s+ciphersuite:\s+(?:\S+\s+){3}${SECURE_CIPHERSUITE} -> Continue
  ^\s+${SECURE_CIPHERSUITE}\s*$$
  ^\s+${SECURE_CIPHERSUITE} -> Continue
  ^\s+(?:\S+\s+)${SECURE_CIPHERSUITE}\s*$$
  ^\s+(?:\S+\s+)${SECURE_CIPHERSUITE} -> Continue
  ^\s+(?:\S+\s+){2}${SECURE_CIPHERSUITE}\s*$$
  ^\s+(?:\S+\s+){2}${SECURE_CIPHERSUITE} -> Continue
  ^\s+(?:\S+\s+){3}${SECURE_CIPHERSUITE}\s*$$
  ^\s+(?:\S+\s+){3}${SECURE_CIPHERSUITE} -> Continue
  ^\s+(?:\S+\s+){4}${SECURE_CIPHERSUITE}\s*$$
  ^\s+(?:\S+\s+){4}${SECURE_CIPHERSUITE} -> Continue
  ^\s+(?:\S+\s+){5}${SECURE_CIPHERSUITE}\s*$$
  ^\s+(?:\S+\s+){5}${SECURE_CIPHERSUITE} -> Continue
  ^HTTP\s+secure\s+server\s+TLS\s+version:\s+${TLS_VERSIONS} -> Continue
  ^HTTP\s+secure\s+server\s+TLS\s+version:\s+(?:\S+\s+)${TLS_VERSIONS}
  ^HTTP\s+secure\s+server\s+client\s+authentication:\s+${SECURE_CLIENT_AUTH}
  ^HTTP\s+secure\s+server\s+PIV\s+authentication:\s+${SECURE_PIV_AUTHN}
  ^HTTP\s+secure\s+server\s+PIV\s+authorization\s+only:\s+${SECURE_PIV_AUTHZ}
  ^HTTP\s+secure\s+server\s+trustpoint:
  ^HTTP\s+secure\s+server\s+peer\s+validation\s+trustpoint:
  ^HTTP\s+secure\s+server\s+ECDHE\s+curve:\s+${SECURE_ECDHE_CURVE}
  ^HTTP\s+secure\s+server\s+active\s+session\s+modules:\s+${SECURE_ACTIVE_SESSION_MODULES}
  ^\s*$$
  ^. -> Error
`
