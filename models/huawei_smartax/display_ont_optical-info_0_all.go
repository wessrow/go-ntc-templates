package huawei_smartax 

type DisplayOntOpticalInfo0All struct {
	Ont_id	string	`json:"ONT_ID"`
	Rx_power	string	`json:"RX_POWER"`
	Tx_power	string	`json:"TX_POWER"`
	Olt_rx_ont	string	`json:"OLT_RX_ONT"`
	Temperature	string	`json:"TEMPERATURE"`
	Voltage	string	`json:"VOLTAGE"`
	Current	string	`json:"CURRENT"`
	Distance	string	`json:"DISTANCE"`
}

var DisplayOntOpticalInfo0All_Template = `Value Key ONT_ID (\d+)
Value RX_POWER (\S+)
Value TX_POWER (\S+)
Value OLT_RX_ONT (\S+)
Value TEMPERATURE (\d+)
Value VOLTAGE (\S+)
Value CURRENT (\d+)
Value DISTANCE (\d+)

Start
  ^\s+ONT\s+Rx\s*power\s+Tx\s*power\s+OLT\s*Rx\s*ONT\s+Temperature\s+Voltage\s+Current\s+Distance
  ^\s+ID\s+\(dBm\)\s+\(dBm\)\s+power\(dBm\)\s+\(C\)\s+\(V\)\s+\(mA\)\s+\(m\)
  ^\s+${ONT_ID}\s+${RX_POWER}\s+${TX_POWER}\s+${OLT_RX_ONT}\s+${TEMPERATURE}\s+${VOLTAGE}\s+${CURRENT}\s+${DISTANCE} -> Record
  ^\s+-
  ^. -> Error
`