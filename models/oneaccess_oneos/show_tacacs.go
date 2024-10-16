package oneaccess_oneos

type ShowTacacs struct {
	Server_port        string `json:"SERVER_PORT"`
	Socket_closes      string `json:"SOCKET_CLOSES"`
	Socket_timeouts    string `json:"SOCKET_TIMEOUTS"`
	Failed_connections string `json:"FAILED_CONNECTIONS"`
	Tacacs_server      string `json:"TACACS_SERVER"`
	Socket_opens       string `json:"SOCKET_OPENS"`
	Socket_aborts      string `json:"SOCKET_ABORTS"`
	Socket_errors      string `json:"SOCKET_ERRORS"`
	Packets_sent       string `json:"PACKETS_SENT"`
	Packet_received    string `json:"PACKET_RECEIVED"`
}

var ShowTacacs_Template string = `Value TACACS_SERVER (.+?)
Value SERVER_PORT (\d+)
Value SOCKET_OPENS (\d+)
Value SOCKET_CLOSES (\d+)
Value SOCKET_ABORTS (\d+)
Value SOCKET_ERRORS (\d+)
Value SOCKET_TIMEOUTS (\d+)
Value FAILED_CONNECTIONS (\d+)
Value PACKETS_SENT (\d+)
Value PACKET_RECEIVED (\d+)

Start
  ^\s+Tacacs\+\s+Server\s+Address\s+:\s.*$$ -> Continue.Record
  ^\s+Tacacs\+\s+Server\s+Address\s+:\s+${TACACS_SERVER}\s*$$
  ^\s+Server\s+port\s+:\s+${SERVER_PORT}\s*$$
  ^\s+Number\sof\ssockets\sopen\s+:\s+${SOCKET_OPENS}\s*$$
  ^\s+Number\sof\ssockets\sclosed\s+:\s+${SOCKET_CLOSES}\s*$$
  ^\s+Number\sof\ssockets\saborted\s+:\s+${SOCKET_ABORTS}\s*$$
  ^\s+Number\sof\ssockets\serror\s+:\s+${SOCKET_ERRORS}\s*$$
  ^\s+Number\sof\ssockets\stimeout\s+:\s+${SOCKET_TIMEOUTS}\s*$$
  ^\s+Number\sof\sconnect\sfails\s+:\s+${FAILED_CONNECTIONS}\s*$$
  ^\s+Number\sof\spackets\ssent\s+:\s+${PACKETS_SENT}\s*$$
  ^\s+Number\sof\spackets\sreceived\s+:\s+${PACKET_RECEIVED}\s*$$
  ^\s*TACACS\+\s+SERVER\s+Statistics\s*$$
  ^\s*-+\s*$$
  ^\s*$$
  ^. -> Error
`
