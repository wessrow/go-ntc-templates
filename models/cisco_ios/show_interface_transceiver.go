package cisco_ios 

type ShowInterfaceTransceiver struct {
	Interface	string	`json:"INTERFACE"`
	Temperature	string	`json:"TEMPERATURE"`
	Voltage	string	`json:"VOLTAGE"`
	Tx_pwr	string	`json:"TX_PWR"`
	Rx_pwr	string	`json:"RX_PWR"`
}

var ShowInterfaceTransceiver_Template = `Value INTERFACE (\w+\d+\/\S+)
Value TEMPERATURE (\d+.\d+)
Value VOLTAGE (\d+.\d+)
Value TX_PWR (\S+)
Value RX_PWR (\S+)

Start
  ^${INTERFACE}\s+${TEMPERATURE}\s+${VOLTAGE}\s+${TX_PWR}\s+${RX_PWR}\s+ -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
`