package cisco_ios

type ShowBfdNeighborsDetails struct {
	Multiplier               string `json:"MULTIPLIER"`
	State                    string `json:"STATE"`
	Tx_last                  string `json:"TX_LAST"`
	Holddown                 string `json:"HOLDDOWN"`
	Rx_last                  string `json:"RX_LAST"`
	Tx_count                 string `json:"TX_COUNT"`
	Uptime                   string `json:"UPTIME"`
	Local_discrim            string `json:"LOCAL_DISCRIM"`
	Interface                string `json:"INTERFACE"`
	Diagnostic_bit           string `json:"DIAGNOSTIC_BIT"`
	Hello_hits               string `json:"HELLO_HITS"`
	Min_tx_interval          string `json:"MIN_TX_INTERVAL"`
	Remote_heard             string `json:"REMOTE_HEARD"`
	Hello_interval           string `json:"HELLO_INTERVAL"`
	Version                  string `json:"VERSION"`
	Echo_function            string `json:"ECHO_FUNCTION"`
	Received_multiplier      string `json:"RECEIVED_MULTIPLIER"`
	Holddown_hits            string `json:"HOLDDOWN_HITS"`
	Poll_bit                 string `json:"POLL_BIT"`
	Min_rx_interval          string `json:"MIN_RX_INTERVAL"`
	Our_addr                 string `json:"OUR_ADDR"`
	Neighbor_addr            string `json:"NEIGHBOR_ADDR"`
	Echo_interval            string `json:"ECHO_INTERVAL"`
	Demand_bit               string `json:"DEMAND_BIT"`
	Template                 string `json:"TEMPLATE"`
	Remote_discrim           string `json:"REMOTE_DISCRIM"`
	Received_min_rx_interval string `json:"RECEIVED_MIN_RX_INTERVAL"`
	Rx_count                 string `json:"RX_COUNT"`
	Registered_protocols     string `json:"REGISTERED_PROTOCOLS"`
}

var ShowBfdNeighborsDetails_Template string = `Value OUR_ADDR (\S+)
Value NEIGHBOR_ADDR (\S+)
Value LOCAL_DISCRIM (\d+)
Value REMOTE_DISCRIM (\d+)
Value REMOTE_HEARD (\S+)
Value HOLDDOWN (\d+)
Value MULTIPLIER (\d{1,2})
Value STATE (\w+)
Value INTERFACE (\S+)
Value ECHO_FUNCTION ((?:\w+\s+)?\w+\s+echo\s+function)
Value ECHO_INTERVAL (\d+)
Value DIAGNOSTIC_BIT (\d)
Value DEMAND_BIT (\d)
Value POLL_BIT (\d)
Value MIN_TX_INTERVAL (\d+)
Value MIN_RX_INTERVAL (\d+)
Value RECEIVED_MIN_RX_INTERVAL (\d+)
Value RECEIVED_MULTIPLIER (\d{1,2})
Value HOLDDOWN_HITS (\d+)
Value HELLO_INTERVAL (\d+)
Value HELLO_HITS (\d+)
Value RX_COUNT (\d+)
Value RX_LAST (\d+)
Value TX_COUNT (\d+)
Value TX_LAST (\d+)
Value REGISTERED_PROTOCOLS (\S+(?:\s+\S+)*)
Value TEMPLATE (\S+)
Value UPTIME (\S+)
Value VERSION (\d+)

Start
  ^\s*IPv\d+\s+Sessions\s*$$
  ^\s*(?:\S+\s+)?NeighAddr.*$$ -> BFD
  #
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^\s*$$
  ^. -> Error

BFD
  ^\s*${OUR_ADDR}\s*${NEIGHBOR_ADDR}\s+${LOCAL_DISCRIM}\/${REMOTE_DISCRIM}\s+${REMOTE_HEARD}\s+\d+\s+\(\s*\d+\s*\)\s+${STATE}\s+${INTERFACE}\s*$$
  #
  ^\s*${NEIGHBOR_ADDR}\s+${LOCAL_DISCRIM}\/${REMOTE_DISCRIM}\s+${REMOTE_HEARD}\s+${STATE}\s+${INTERFACE}\s*$$
  ^\s*OurAddr:\s+${OUR_ADDR}\s*$$
  #
  # common BFD details
  ^\s*Session\s+state\s+is\s+\S+\s+and\s+${ECHO_FUNCTION}(?:\s+with\s+${ECHO_INTERVAL}\s+)?.*$$
  ^\s*Local\s+Diag:\s+${DIAGNOSTIC_BIT},\s+Demand\s+mode:\s+${DEMAND_BIT},\s+Poll\s+bit:\s+${POLL_BIT}\s*$$
  ^\s*MinTxInt:\s+${MIN_TX_INTERVAL},\s+MinRxInt:\s+${MIN_RX_INTERVAL},\s+Multiplier:\s+${MULTIPLIER}\s*$$
  ^\s*Received\s+MinRxInt:\s+${RECEIVED_MIN_RX_INTERVAL},\s+Received Multiplier:\s+${RECEIVED_MULTIPLIER}\s*$$
  #
  # some legacy output spells holddown with one d
  ^\s*Holdd?own\s+\(hits\):\s+${HOLDDOWN}\(${HOLDDOWN_HITS}\),\s+Hello\s+\(hits\):\s+${HELLO_INTERVAL}\(${HELLO_HITS}\)\s*$$
  #
  ^\s*Rx\s+Count:\s+${RX_COUNT},(?:\s+\S+){5}\s*$$
  ^\s*Tx\s+Count:\s+${TX_COUNT},(?:\s+\S+){5}\s*$$
  ^\s*Rx\s+Count:\s+${RX_COUNT},(?:\s+\S+){6}\s+${RX_LAST}(?:\s+\S+){2}\s*$$
  ^\s*Tx\s+Count:\s+${TX_COUNT},(?:\s+\S+){6}\s+${TX_LAST}(?:\s+\S+){2}\s*$$
  ^\s*Registered\s+protocols:\s+${REGISTERED_PROTOCOLS}\s*$$
  ^\s*Template:\s+${TEMPLATE}\s*$$
  ^\s*Uptime:\s+${UPTIME}\s*$$
  ^\s*Last\s+packet:\s+Version:\s+${VERSION}\s+.*$$ -> Record
  #
  # lines of no interest for the time being
  #
  ^\s*Session\s+Host.*$$
  ^\s*Handle.*$$
  ^\s*Elapsed\s+time\s+watermarks.*$$
  ^\s+(?:\S+\s+){1,3}bit.*$$
  ^\s+Multiplier.*$$
  ^\s+My\s+Discr.*$$
  ^\s+Min\s+\S+\s+interval.*$$
  #
  ^\s*$$ -> Start
  ^. -> Error
`
