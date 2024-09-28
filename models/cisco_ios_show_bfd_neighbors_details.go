package models

type CiscoIosShowBfdNeighborsDetails struct {
	Our_addr	string	`json:"OUR_ADDR"`
	Neighbor_addr	string	`json:"NEIGHBOR_ADDR"`
	Local_discrim	string	`json:"LOCAL_DISCRIM"`
	Remote_discrim	string	`json:"REMOTE_DISCRIM"`
	Remote_heard	string	`json:"REMOTE_HEARD"`
	Holddown	string	`json:"HOLDDOWN"`
	Multiplier	string	`json:"MULTIPLIER"`
	State	string	`json:"STATE"`
	Interface	string	`json:"INTERFACE"`
	Echo_function	string	`json:"ECHO_FUNCTION"`
	Echo_interval	string	`json:"ECHO_INTERVAL"`
	Diagnostic_bit	string	`json:"DIAGNOSTIC_BIT"`
	Demand_bit	string	`json:"DEMAND_BIT"`
	Poll_bit	string	`json:"POLL_BIT"`
	Min_tx_interval	string	`json:"MIN_TX_INTERVAL"`
	Min_rx_interval	string	`json:"MIN_RX_INTERVAL"`
	Received_min_rx_interval	string	`json:"RECEIVED_MIN_RX_INTERVAL"`
	Received_multiplier	string	`json:"RECEIVED_MULTIPLIER"`
	Holddown_hits	string	`json:"HOLDDOWN_HITS"`
	Hello_interval	string	`json:"HELLO_INTERVAL"`
	Hello_hits	string	`json:"HELLO_HITS"`
	Rx_count	string	`json:"RX_COUNT"`
	Rx_last	string	`json:"RX_LAST"`
	Tx_count	string	`json:"TX_COUNT"`
	Tx_last	string	`json:"TX_LAST"`
	Registered_protocols	string	`json:"REGISTERED_PROTOCOLS"`
	Template	string	`json:"TEMPLATE"`
	Uptime	string	`json:"UPTIME"`
	Version	string	`json:"VERSION"`
}