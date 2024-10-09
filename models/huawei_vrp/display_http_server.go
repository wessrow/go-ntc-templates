package huawei_vrp 

type DisplayHttpServer struct {
	Http_server_status	string	`json:"HTTP_SERVER_STATUS"`
	Http_server_port	string	`json:"HTTP_SERVER_PORT"`
	Http_timeout_interval	string	`json:"HTTP_TIMEOUT_INTERVAL"`
	Current_online_users	string	`json:"CURRENT_ONLINE_USERS"`
	Maximum_users_allowed	string	`json:"MAXIMUM_USERS_ALLOWED"`
	Https_server_status	string	`json:"HTTPS_SERVER_STATUS"`
	Https_server_port	string	`json:"HTTPS_SERVER_PORT"`
	Http_ssl_policy	string	`json:"HTTP_SSL_POLICY"`
	Http_ipv6_server_status	string	`json:"HTTP_IPV6_SERVER_STATUS"`
	Http_ipv6_server_port	string	`json:"HTTP_IPV6_SERVER_PORT"`
	Https_ipv6_status	string	`json:"HTTPS_IPV6_STATUS"`
	Http_server_source_address	string	`json:"HTTP_SERVER_SOURCE_ADDRESS"`
}

var DisplayHttpServer_Template = `Value HTTP_SERVER_STATUS (\S+)
Value HTTP_SERVER_PORT (disabled|\d+\(\d+\))
Value HTTP_TIMEOUT_INTERVAL (\d+)
Value CURRENT_ONLINE_USERS (\d+)
Value MAXIMUM_USERS_ALLOWED (\d+)
Value HTTPS_SERVER_STATUS (\S+)
Value HTTPS_SERVER_PORT (disabled|\d+\(\d+\))
Value HTTP_SSL_POLICY (\S+)
Value HTTP_IPV6_SERVER_STATUS (\S+)
Value HTTP_IPV6_SERVER_PORT (disabled|\d+\(\d+\))
Value HTTPS_IPV6_STATUS (disabled|\d+\(\d+\))
Value HTTP_SERVER_SOURCE_ADDRESS (\S+)

Start
  ^\s*HTTP\s*Server\s*Status\s*:\s*${HTTP_SERVER_STATUS}
  ^\s*HTTP\s*Server\s*Port\s*:\s*${HTTP_SERVER_PORT}
  ^\s*HTTP\s*Timeout\s*Interval\s*:\s*${HTTP_TIMEOUT_INTERVAL}
  ^\s*Current\s*Online\s*Users\s*:\s*${CURRENT_ONLINE_USERS}
  ^\s*Maximum\s*Users\s*Allowed\s*:\s*${MAXIMUM_USERS_ALLOWED}
  ^\s*HTTP\s*Secure-server\s*Status\s*:\s*${HTTPS_SERVER_STATUS}
  ^\s*HTTP\s*Secure-server\s*Port\s*:\s*${HTTPS_SERVER_PORT}
  ^\s*HTTP\s*SSL\s*Policy\s*:\s*${HTTP_SSL_POLICY}
  ^\s*HTTP\s*IPv6\s*Server\s*Status\s*:\s*${HTTP_IPV6_SERVER_STATUS}
  ^\s*HTTP\s*IPv6\s*Server\s*Port\s*:\s*${HTTP_IPV6_SERVER_PORT}
  ^\s*HTTP\s*IPv6\s*Secure-server\s*Status\s*:\s*${HTTP_IPV6_SERVER_PORT}
  ^\s*HTTP\s*IPv6\s*Secure-server\s*Port\s*:\s*${HTTPS_IPV6_STATUS}
  ^\s*HTTP\s*server\s*source\s*address\s*:\s*${HTTP_SERVER_SOURCE_ADDRESS}
  ^\s*$$
  ^. -> Error
`