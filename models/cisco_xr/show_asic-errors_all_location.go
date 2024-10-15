package cisco_xr

type ShowAsicErrorsAllLocation struct {
	Instance            string `json:"INSTANCE"`
	Mbe_error_count     string `json:"MBE_ERROR_COUNT"`
	Reset_error_count   string `json:"RESET_ERROR_COUNT"`
	Asic                string `json:"ASIC"`
	Number_of_nodes     string `json:"NUMBER_OF_NODES"`
	Sbe_error_count     string `json:"SBE_ERROR_COUNT"`
	Parity_error_count  string `json:"PARITY_ERROR_COUNT"`
	Crc_error_count     string `json:"CRC_ERROR_COUNT"`
	Generic_error_count string `json:"GENERIC_ERROR_COUNT"`
}

var ShowAsicErrorsAllLocation_Template string = `Value Filldown ASIC (\S+)
Value Required INSTANCE (\d+)
Value NUMBER_OF_NODES (\d+)
Value SBE_ERROR_COUNT (\d+)
Value MBE_ERROR_COUNT (\d+)
Value PARITY_ERROR_COUNT (\d+)
Value CRC_ERROR_COUNT (\d+)
Value GENERIC_ERROR_COUNT (\d+)
Value RESET_ERROR_COUNT (\d+)

Start
  ^.+ASIC\s+Error\s+Summary\s+\* -> Continue.Record
  ^\*\s+${ASIC}\s+ASIC\s+Error\s+Summary\s+\*
  ^Instance -> Continue.Record
  ^Instance\s+:\s+${INSTANCE}
  ^Number\s+of\s+nodes\s+:\s+${NUMBER_OF_NODES}
  ^SBE\s+error\s+count\s+:\s+${SBE_ERROR_COUNT}
  ^MBE\s+error\s+count\s+:\s+${MBE_ERROR_COUNT}
  ^Parity\s+error\s+count\s+:\s+${PARITY_ERROR_COUNT}
  ^CRC\s+error\s+count\s+:\s+${CRC_ERROR_COUNT}
  ^Generic\s+error\s+count\s+:\s+${GENERIC_ERROR_COUNT}
  ^Reset\s+error\s+count\s+:\s+${RESET_ERROR_COUNT}
  ^-+$$
  ^\*+$$
  ^\s*$$
  ^. -> Error "LINE NOT FOUND"
`
