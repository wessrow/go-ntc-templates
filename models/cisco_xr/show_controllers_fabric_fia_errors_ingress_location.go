package cisco_xr 

type ShowControllersFabricFiaErrorsIngressLocation struct {
	Fia	string	`json:"FIA"`
	Category	string	`json:"CATEGORY"`
	To_xbar_uc_crc_0	string	`json:"TO_XBAR_UC_CRC_0"`
	To_xbar_uc_crc_1	string	`json:"TO_XBAR_UC_CRC_1"`
	To_xbar_uc_crc_2	string	`json:"TO_XBAR_UC_CRC_2"`
	To_xbar_uc_crc_3	string	`json:"TO_XBAR_UC_CRC_3"`
	To_xbar_mc_crc_0	string	`json:"TO_XBAR_MC_CRC_0"`
	To_xbar_mc_crc_1	string	`json:"TO_XBAR_MC_CRC_1"`
	To_xbar_mc_crc_2	string	`json:"TO_XBAR_MC_CRC_2"`
	To_xbar_mc_crc_3	string	`json:"TO_XBAR_MC_CRC_3"`
	Nb_pa_read_data_err	string	`json:"NB_PA_READ_DATA_ERR"`
	Pa_header_err	string	`json:"PA_HEADER_ERR"`
	Pa_crc16_err	string	`json:"PA_CRC16_ERR"`
	Pa_crc32_err	string	`json:"PA_CRC32_ERR"`
	Pa_to_tf_err	string	`json:"PA_TO_TF_ERR"`
	Ab_overflow_req_lost	string	`json:"AB_OVERFLOW_REQ_LOST"`
	Ni_bad_crc32	string	`json:"NI_BAD_CRC32"`
	Ni_crc32_corrupt	string	`json:"NI_CRC32_CORRUPT"`
}

var ShowControllersFabricFiaErrorsIngressLocation_Template = `Value FIA (\S+)
Value CATEGORY (\S+)
Value TO_XBAR_UC_CRC_0 (\d+)
Value TO_XBAR_UC_CRC_1 (\d+)
Value TO_XBAR_UC_CRC_2 (\d+)
Value TO_XBAR_UC_CRC_3 (\d+)
Value TO_XBAR_MC_CRC_0 (\d+)
Value TO_XBAR_MC_CRC_1 (\d+)
Value TO_XBAR_MC_CRC_2 (\d+)
Value TO_XBAR_MC_CRC_3 (\d+)
Value NB_PA_READ_DATA_ERR (\d+)
Value PA_HEADER_ERR (\d+)
Value PA_CRC16_ERR (\d+)
Value PA_CRC32_ERR (\d+)
Value PA_TO_TF_ERR (\d+)
Value AB_OVERFLOW_REQ_LOST (\d+)
Value NI_BAD_CRC32 (\d+)
Value NI_CRC32_CORRUPT (\d+)

Start
  ^\s*\*+\s+\S+ -> Continue.Record
  ^\s*\*+\s+${FIA}\s+\*+
  ^Category:\s+${CATEGORY}
  ^\s+To\s+Xbar\s+Uc\s+Crc-0\s+${TO_XBAR_UC_CRC_0}
  ^\s+To\s+Xbar\s+Uc\s+Crc-1\s+${TO_XBAR_UC_CRC_1}
  ^\s+To\s+Xbar\s+Uc\s+Crc-2\s+${TO_XBAR_UC_CRC_2}
  ^\s+To\s+Xbar\s+Uc\s+Crc-3\s+${TO_XBAR_UC_CRC_3}
  ^\s+To\s+Xbar\s+Mc\s+Crc-0\s+${TO_XBAR_MC_CRC_0}
  ^\s+To\s+Xbar\s+Mc\s+Crc-1\s+${TO_XBAR_MC_CRC_1}
  ^\s+To\s+Xbar\s+Mc\s+Crc-2\s+${TO_XBAR_MC_CRC_2}
  ^\s+To\s+Xbar\s+Mc\s+Crc-3\s+${TO_XBAR_MC_CRC_3}
  ^\s+nb\s+pa\s+read\s+data\s+err\s+${NB_PA_READ_DATA_ERR}
  ^\s+pa\s+header\s+err\s+${PA_HEADER_ERR}
  ^\s+pa\s+crc16\s+err\s+${PA_CRC16_ERR}
  ^\s+pa\s+crc32\s+err\s+${PA_CRC32_ERR}
  ^\s+pa\s+to\s+tf\s+err\s+${PA_TO_TF_ERR}
  ^\s+ab\s+overflow\s+req lost\s+${AB_OVERFLOW_REQ_LOST}
  ^\s+ni\s+bad\s+crc32\s+${NI_BAD_CRC32}
  ^\s+ni\s+crc32\s+corrupt\s+${NI_CRC32_CORRUPT}
  ^\s*$$
  ^.* -> Error "LINE NOT FOUND"
`