package arista_eos 

type ShowMacSecurityMkaCounters struct {
	Interface	string	`json:"INTERFACE"`
	Rx_success	string	`json:"RX_SUCCESS"`
	Rx_failure	string	`json:"RX_FAILURE"`
	Tx_success	string	`json:"TX_SUCCESS"`
	Tx_failure	string	`json:"TX_FAILURE"`
}

var ShowMacSecurityMkaCounters_Template = `Value INTERFACE (\S+)
Value RX_SUCCESS (\S+)
Value RX_FAILURE (\S+)
Value TX_SUCCESS (\S+)
Value TX_FAILURE (\S+)

Start
  ^Interface.*Failure -> Data

Data
  ^${INTERFACE}\s+${RX_SUCCESS}\s+${RX_FAILURE}\s+${TX_SUCCESS}\s+${TX_FAILURE} -> Record
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
  
EOF  
`