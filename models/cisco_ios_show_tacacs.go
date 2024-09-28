package models

type CiscoIosShowTacacs struct {
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