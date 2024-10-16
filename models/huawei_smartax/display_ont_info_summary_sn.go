package huawei_smartax

type DisplayOntInfoSummarySn struct {
	Serial_number   string `json:"SERIAL_NUMBER"`
	Ont_model_type  string `json:"ONT_MODEL_TYPE"`
	Distance_m      string `json:"DISTANCE_M"`
	Rx_tx_power_dbm string `json:"RX_TX_POWER_DBM"`
	Description     string `json:"DESCRIPTION"`
	Ont_id          string `json:"ONT_ID"`
}

var DisplayOntInfoSummarySn_Template string = `Value Key ONT_ID (\d+)
Value SERIAL_NUMBER (\w+)
Value ONT_MODEL_TYPE (\S+)
Value DISTANCE_M (\d+|-)
Value RX_TX_POWER_DBM (-?\d+\.\d+\/-?\d+\.\d+|\S+)
Value DESCRIPTION (\S+)

Start
  ^\s+-
  ^\s+ONT\s+SN\s+Type\s+Distance\s+Rx\/Tx\s+power\s+Description
  ^\s+ID\s+\(m\)\s+\(dBm\) -> SNs

SNs
  ^\s+${ONT_ID}\s+${SERIAL_NUMBER}\s+${ONT_MODEL_TYPE}\s+${DISTANCE_M}\s+${RX_TX_POWER_DBM}\s+${DESCRIPTION} -> Record
  ^\s+-
  ^\s*$$
  ^. -> Error
`
