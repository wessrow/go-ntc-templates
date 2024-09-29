package cisco_ios 

type ShowTacacs struct {
	Tacacs_server_name	string	`json:"TACACS_SERVER_NAME"`
	Tacacs_server	string	`json:"TACACS_SERVER"`
	Server_port	string	`json:"SERVER_PORT"`
	Socket_opens	string	`json:"SOCKET_OPENS"`
	Socket_closes	string	`json:"SOCKET_CLOSES"`
	Socket_aborts	string	`json:"SOCKET_ABORTS"`
	Socket_errors	string	`json:"SOCKET_ERRORS"`
	Socket_timeouts	string	`json:"SOCKET_TIMEOUTS"`
	Failed_connections	string	`json:"FAILED_CONNECTIONS"`
	Packets_sent	string	`json:"PACKETS_SENT"`
	Packet_received	string	`json:"PACKET_RECEIVED"`
}

var ShowTacacs_Template = `Value TACACS_SERVER_NAME (.+?)
Value TACACS_SERVER (.+?)
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
  ^Tacacs -> Record
  ^\s+Server\s+address:\s+${TACACS_SERVER}\s*$$
  ^\s+Server\s+name:\s+${TACACS_SERVER_NAME}\s*$$
  ^\s+Server\s+port:\s+${SERVER_PORT}\s*$$
  ^\s+Socket\s+opens:\s+ ${SOCKET_OPENS}\s*$$
  ^\s+Socket\s+closes:\s+${SOCKET_CLOSES}\s*$$
  ^\s+Socket\s+aborts:\s+${SOCKET_ABORTS}\s*$$
  ^\s+Socket\s+errors:\s+${SOCKET_ERRORS}\s*$$
  ^\s+Socket\s+Timeouts:\s+${SOCKET_TIMEOUTS}\s*$$
  ^\s+Failed\s+Connect\s+Attempts:\s+${FAILED_CONNECTIONS}\s*$$
  ^\s+Total\s+Packets\s+Sent:\s+${PACKETS_SENT}\s*$$
  ^\s+Total\s+Packets\s+Recv:\s+${PACKET_RECEIVED}\s*$$
  ^\s*$$
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^. -> Error
`