package models

type AristaEosShowInterfacesTransceiver struct {
	Port	string	`json:"PORT"`
	Temp	string	`json:"TEMP"`
	Voltage	string	`json:"VOLTAGE"`
	Bias_current	string	`json:"BIAS_CURRENT"`
	Tx_power	string	`json:"TX_POWER"`
	Rx_power	string	`json:"RX_POWER"`
}

var AristaEosShowInterfacesTransceiver_Template = `Value PORT (\S+)
Value TEMP (\S+)
Value VOLTAGE (\S+)
Value BIAS_CURRENT (\S+)
Value TX_POWER (\S+)
Value RX_POWER (\S+)


Start
  ^---- -> Begin

Begin
  ^${PORT}\s+${TEMP}\s+${VOLTAGE}\s+${BIAS_CURRENT}\s+${TX_POWER}\s+${RX_POWER}\s+(\S+ ago|N/A) -> Record
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
  
EOF

`