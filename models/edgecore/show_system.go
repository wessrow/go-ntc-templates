package edgecore

type ShowSystem struct {
	Oid                  string   `json:"OID"`
	Location             string   `json:"LOCATION"`
	Secure_server_port   string   `json:"SECURE_SERVER_PORT"`
	Telnet_server_port   string   `json:"TELNET_SERVER_PORT"`
	Auth_login           string   `json:"AUTH_LOGIN"`
	Description          string   `json:"DESCRIPTION"`
	Hostname             string   `json:"HOSTNAME"`
	Server_port          string   `json:"SERVER_PORT"`
	Secure_server_status string   `json:"SECURE_SERVER_STATUS"`
	Contact              string   `json:"CONTACT"`
	Jumbo_frame_status   string   `json:"JUMBO_FRAME_STATUS"`
	Post_results         []string `json:"POST_RESULTS"`
	Auth_enabled         string   `json:"AUTH_ENABLED"`
	Uptime               string   `json:"UPTIME"`
	Address              string   `json:"ADDRESS"`
	Server_status        string   `json:"SERVER_STATUS"`
	Telnet_server_status string   `json:"TELNET_SERVER_STATUS"`
}

var ShowSystem_Template string = `Value DESCRIPTION (.*)
Value OID (\d+(\.\d+)*)
Value UPTIME (.*)
Value HOSTNAME (\S+)
Value LOCATION (.*)
Value CONTACT (.*)
Value ADDRESS ([a-zA-Z0-9]{2}(?:-[a-zA-Z0-9]{2}){5})
Value SERVER_STATUS (\S+)
Value SERVER_PORT (\d+)
Value SECURE_SERVER_STATUS (\S+)
Value SECURE_SERVER_PORT (\d+)
Value TELNET_SERVER_STATUS (\S+)
Value TELNET_SERVER_PORT (\d+)
Value AUTH_LOGIN (.*)
Value AUTH_ENABLED (.*)
Value JUMBO_FRAME_STATUS (\S+)
Value List POST_RESULTS (.*)

Start
  ^\s*System\s+Description:\s+${DESCRIPTION}\s*$$
  ^\s*System\s+OID\s+String:\s+${OID}\s*$$
  ^\s*System\s+Information\s*$$ -> SystemInfo
  ^\s*$$
  ^. -> Error

SystemInfo
  ^\s*System\s+Up\s+Time:\s+${UPTIME}\s*$$
  ^\s*System\s+Name:\s+${HOSTNAME}\s*$$
  ^\s*System\s+Location:\s+${LOCATION}\s*$$
  ^\s*System\s+Contact:\s+${CONTACT}\s*$$
  ^\s*MAC\s+Address\s+\(Unit\d+\):\s+${ADDRESS}\s*$$
  ^\s*Web\s+Server:\s+${SERVER_STATUS}\s*$$
  ^\s*Web\s+Server\s+Port:\s+${SERVER_PORT}\s*$$
  ^\s*Web\s+Secure\s+Server:\s+${SECURE_SERVER_STATUS}\s*$$
  ^\s*Web\s+Secure\s+Server\s+Port:\s+${SECURE_SERVER_PORT}\s*$$
  ^\s*Telnet\s+Server:\s+${TELNET_SERVER_STATUS}\s*$$
  ^\s*Telnet\s+Server\s+Port:\s+${TELNET_SERVER_PORT}\s*$$
  ^\s*Authentication\s+Login:\s+${AUTH_LOGIN}\s*$$
  ^\s*Authentication\s+Enabled:\s+${AUTH_ENABLED}\s*$$
  ^\s*Jumbo\s+Frame:\s+${JUMBO_FRAME_STATUS}\s*$$
  ^\s*POST\s+Result:\s*$$ -> PostResults
  ^\s*$$
  ^. -> Error

PostResults
  ^\s*Done\s+All\s+Pass.\s*$$
  ^\s*$$
  ^\s*${POST_RESULTS}\s*$$ -> PostResults
  ^. -> Error
`
