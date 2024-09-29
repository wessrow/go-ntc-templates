package hp_procurve 

type ShowInterfaces struct {
	Port	string	`json:"PORT"`
	Total_bytes	string	`json:"TOTAL_BYTES"`
	Total_frames	string	`json:"TOTAL_FRAMES"`
	Errors_rx	string	`json:"ERRORS_RX"`
	Drops_tx	string	`json:"DROPS_TX"`
	Flow_ctrl	string	`json:"FLOW_CTRL"`
	Bcast_limit	string	`json:"BCAST_LIMIT"`
}

var ShowInterfaces_Template = `Value PORT (\S+)
Value TOTAL_BYTES (\S+)
Value TOTAL_FRAMES (\S+)
Value ERRORS_RX (\S+)
Value DROPS_TX (\S+)
Value FLOW_CTRL (off|on)
Value BCAST_LIMIT (\d+)

Start
  ^\s+Status.*Counters\s*
  ^\s*$$
  ^\s+Flow\s*
  ^\s+Port.*Ctrl\s*
  ^\s+-+\s+-+\s+-+\s+-+\s+-+\s-+\s*$$ -> ShowInterfaces1
  ^\s+Flow\s+Bcast\s*
  ^\s+Port.*Limit\s*
  ^\s+-+\s+-+\s+-+\s+-+\s+-+\s+-+\s-+\s*$$ -> ShowInterfaces2


ShowInterfaces1
  ^\s+${PORT}\s+${TOTAL_BYTES}\s+${TOTAL_FRAMES}\s+${ERRORS_RX}\s+${DROPS_TX}\s+${FLOW_CTRL}\s*$$ -> Record
  ^\s*$$
  ^. -> Error

ShowInterfaces2
  ^\s+${PORT}\s+${TOTAL_BYTES}\s+${TOTAL_FRAMES}\s+${ERRORS_RX}\s+${DROPS_TX}\s+${FLOW_CTRL}\s+${BCAST_LIMIT}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`