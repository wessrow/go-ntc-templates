package cisco_xr 

type ShowControllersFabricFiaDropsIngressLocation struct {
	Fia	string	`json:"FIA"`
	Category	string	`json:"CATEGORY"`
	From_spaui_drop_0	string	`json:"FROM_SPAUI_DROP_0"`
	Accpt_tbl_0	string	`json:"ACCPT_TBL_0"`
	Ctl_len_0	string	`json:"CTL_LEN_0"`
	Short_pkt_0	string	`json:"SHORT_PKT_0"`
	Max_pkt_len_0	string	`json:"MAX_PKT_LEN_0"`
	Min_pkt_len_0	string	`json:"MIN_PKT_LEN_0"`
	From_spaui_drop_1	string	`json:"FROM_SPAUI_DROP_1"`
	Accpt_tbl_1	string	`json:"ACCPT_TBL_1"`
	Ctl_len_1	string	`json:"CTL_LEN_1"`
	Short_pkt_1	string	`json:"SHORT_PKT_1"`
	Max_pkt_len_1	string	`json:"MAX_PKT_LEN_1"`
	Min_pkt_len_1	string	`json:"MIN_PKT_LEN_1"`
	Tail_drp	string	`json:"TAIL_DRP"`
	Vqi_drp	string	`json:"VQI_DRP"`
	Header_parsing_drp	string	`json:"HEADER_PARSING_DRP"`
	Pw_to_ni_drp	string	`json:"PW_TO_NI_DRP"`
	Ni_from_pw_drp	string	`json:"NI_FROM_PW_DRP"`
	Sp0_crc_err	string	`json:"SP0_CRC_ERR"`
	Sp0_bad_align	string	`json:"SP0_BAD_ALIGN"`
	Sp0_bad_code	string	`json:"SP0_BAD_CODE"`
	Sp0_align_fail	string	`json:"SP0_ALIGN_FAIL"`
	Sp0_prot_err	string	`json:"SP0_PROT_ERR"`
	Sp1_crc_err	string	`json:"SP1_CRC_ERR"`
	Sp1_bad_align	string	`json:"SP1_BAD_ALIGN"`
	Sp1_bad_code	string	`json:"SP1_BAD_CODE"`
	Sp1_align_fail	string	`json:"SP1_ALIGN_FAIL"`
	Sp1_prot_err	string	`json:"SP1_PROT_ERR"`
}

var ShowControllersFabricFiaDropsIngressLocation_Template = `Value FIA (\S+)
Value CATEGORY (\S+)
Value FROM_SPAUI_DROP_0 (\d+)
Value ACCPT_TBL_0 (\d+)
Value CTL_LEN_0 (\d+)
Value SHORT_PKT_0 (\d+)
Value MAX_PKT_LEN_0 (\d+)
Value MIN_PKT_LEN_0 (\d+)
Value FROM_SPAUI_DROP_1 (\d+)
Value ACCPT_TBL_1 (\d+)
Value CTL_LEN_1 (\d+)
Value SHORT_PKT_1 (\d+)
Value MAX_PKT_LEN_1 (\d+)
Value MIN_PKT_LEN_1 (\d+)
Value TAIL_DRP (\d+)
Value VQI_DRP (\d+)
Value HEADER_PARSING_DRP (\d+)
Value PW_TO_NI_DRP (\d+)
Value NI_FROM_PW_DRP (\d+)
Value SP0_CRC_ERR (\d+)
Value SP0_BAD_ALIGN (\d+)
Value SP0_BAD_CODE (\d+)
Value SP0_ALIGN_FAIL (\d+)
Value SP0_PROT_ERR (\d+)
Value SP1_CRC_ERR (\d+)
Value SP1_BAD_ALIGN (\d+)
Value SP1_BAD_CODE (\d+)
Value SP1_ALIGN_FAIL (\d+)
Value SP1_PROT_ERR (\d+)

Start
  ^\s*\*+\s+\S+ -> Continue.Record
  ^\s*\*+\s+${FIA}\s+\*+
  ^Category:\s+${CATEGORY}
  ^\s+From\s+Spaui\s+Drop-0\s+${FROM_SPAUI_DROP_0}
  ^\s+accpt\s+tbl-0\s+${ACCPT_TBL_0}
  ^\s+ctl\s+len-0\s+${CTL_LEN_0}
  ^\s+short\s+pkt-0\s+${SHORT_PKT_0}
  ^\s+max\s+pkt\s+len-0\s+${MAX_PKT_LEN_0}
  ^\s+min\s+pkt\s+len-0\s+${MIN_PKT_LEN_0}
  ^\s+From\s+Spaui\s+Drop-1\s+${FROM_SPAUI_DROP_1}
  ^\s+accpt\s+tbl-1\s+${ACCPT_TBL_1}
  ^\s+ctl\s+len-1\s+${CTL_LEN_1}
  ^\s+short\s+pkt-1\s+${SHORT_PKT_1}
  ^\s+max\s+pkt\s+len-1\s+${MAX_PKT_LEN_1}
  ^\s+min\s+pkt\s+len-1\s+${MIN_PKT_LEN_1}
  ^\s+Tail\s+drp\s+${TAIL_DRP}
  ^\s+Vqi\s+drp\s+${VQI_DRP}
  ^\s+Header\s+parsing drp\s+${HEADER_PARSING_DRP}
  ^\s+pw\s+to\s+ni\s+drp\s+${PW_TO_NI_DRP}
  ^\s+ni\s+from\s+pw\s+drp\s+${NI_FROM_PW_DRP}
  ^\s+sp0\s+crc\s+err\s+${SP0_CRC_ERR}
  ^\s+sp0\s+bad\s+align\s+${SP0_BAD_ALIGN}
  ^\s+sp0\s+bad\s+code\s+${SP0_BAD_CODE}
  ^\s+sp0\s+align\s+fail\s+${SP0_ALIGN_FAIL}
  ^\s+sp0\s+prot\s+err\s+${SP0_PROT_ERR}
  ^\s+sp1\s+crc\s+err\s+${SP1_CRC_ERR}
  ^\s+sp1\s+bad\s+align\s+${SP1_BAD_ALIGN}
  ^\s+sp1\s+bad\s+code\s+${SP1_BAD_CODE}
  ^\s+sp1\s+align\s+fail\s+${SP1_ALIGN_FAIL}
  ^\s+sp1\s+prot\s+err\s+${SP1_PROT_ERR}
  ^\s*$$
  ^.* -> Error "LINE NOT FOUND"
`