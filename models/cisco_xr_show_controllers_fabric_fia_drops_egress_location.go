package models

type CiscoXrShowControllersFabricFiaDropsEgressLocation struct {
	Fia	string	`json:"FIA"`
	Category	string	`json:"CATEGORY"`
	From_xbar_uc_crc_0	string	`json:"FROM_XBAR_UC_CRC_0"`
	From_xbar_uc_crc_1	string	`json:"FROM_XBAR_UC_CRC_1"`
	From_xbar_uc_crc_2	string	`json:"FROM_XBAR_UC_CRC_2"`
	From_xbar_uc_crc_3	string	`json:"FROM_XBAR_UC_CRC_3"`
	From_xbar_uc_drp_0	string	`json:"FROM_XBAR_UC_DRP_0"`
	From_xbar_uc_drp_1	string	`json:"FROM_XBAR_UC_DRP_1"`
	From_xbar_uc_drp_2	string	`json:"FROM_XBAR_UC_DRP_2"`
	From_xbar_uc_drp_3	string	`json:"FROM_XBAR_UC_DRP_3"`
	From_xbar_mc_crc_0	string	`json:"FROM_XBAR_MC_CRC_0"`
	From_xbar_mc_crc_1	string	`json:"FROM_XBAR_MC_CRC_1"`
	From_xbar_mc_crc_2	string	`json:"FROM_XBAR_MC_CRC_2"`
	From_xbar_mc_crc_3	string	`json:"FROM_XBAR_MC_CRC_3"`
	From_xbar_mc_drp_0	string	`json:"FROM_XBAR_MC_DRP_0"`
	From_xbar_mc_drp_1	string	`json:"FROM_XBAR_MC_DRP_1"`
	From_xbar_mc_drp_2	string	`json:"FROM_XBAR_MC_DRP_2"`
	From_xbar_mc_drp_3	string	`json:"FROM_XBAR_MC_DRP_3"`
	Uc_dq_pkt_len_crc_ro_seq_len_error_drp	string	`json:"UC_DQ_PKT_LEN_CRC_RO_SEQ_LEN_ERROR_DRP"`
	Uc_eq_pkt_len_crc_lookup_drp	string	`json:"UC_EQ_PKT_LEN_CRC_LOOKUP_DRP"`
	Mc_rf_crc_drp	string	`json:"MC_RF_CRC_DRP"`
	Mc_vl0_src0_buffer_full_drp	string	`json:"MC_VL0_SRC0_BUFFER_FULL_DRP"`
	Mc_vl1_src0_buffer_full_drp	string	`json:"MC_VL1_SRC0_BUFFER_FULL_DRP"`
	Mc_vl2_src0_buffer_full_drp	string	`json:"MC_VL2_SRC0_BUFFER_FULL_DRP"`
	Mc_vl3_src0_buffer_full_drp	string	`json:"MC_VL3_SRC0_BUFFER_FULL_DRP"`
	Mc_vl0_src1_buffer_full_drp	string	`json:"MC_VL0_SRC1_BUFFER_FULL_DRP"`
	Mc_vl1_src1_buffer_full_drp	string	`json:"MC_VL1_SRC1_BUFFER_FULL_DRP"`
	Mc_vl2_src1_buffer_full_drp	string	`json:"MC_VL2_SRC1_BUFFER_FULL_DRP"`
	Mc_vl3_src1_buffer_full_drp	string	`json:"MC_VL3_SRC1_BUFFER_FULL_DRP"`
}

var CiscoXrShowControllersFabricFiaDropsEgressLocation_Template = `Value FIA (\S+)
Value CATEGORY (\S+)
Value FROM_XBAR_UC_CRC_0 (\d+)
Value FROM_XBAR_UC_CRC_1 (\d+)
Value FROM_XBAR_UC_CRC_2 (\d+)
Value FROM_XBAR_UC_CRC_3 (\d+)
Value FROM_XBAR_UC_DRP_0 (\d+)
Value FROM_XBAR_UC_DRP_1 (\d+)
Value FROM_XBAR_UC_DRP_2 (\d+)
Value FROM_XBAR_UC_DRP_3 (\d+)
Value FROM_XBAR_MC_CRC_0 (\d+)
Value FROM_XBAR_MC_CRC_1 (\d+)
Value FROM_XBAR_MC_CRC_2 (\d+)
Value FROM_XBAR_MC_CRC_3 (\d+)
Value FROM_XBAR_MC_DRP_0 (\d+)
Value FROM_XBAR_MC_DRP_1 (\d+)
Value FROM_XBAR_MC_DRP_2 (\d+)
Value FROM_XBAR_MC_DRP_3 (\d+)
Value UC_DQ_PKT_LEN_CRC_RO_SEQ_LEN_ERROR_DRP (\d+)
Value UC_EQ_PKT_LEN_CRC_LOOKUP_DRP (\d+)
Value MC_RF_CRC_DRP (\d+)
Value MC_VL0_SRC0_BUFFER_FULL_DRP (\d+)
Value MC_VL1_SRC0_BUFFER_FULL_DRP (\d+)
Value MC_VL2_SRC0_BUFFER_FULL_DRP (\d+)
Value MC_VL3_SRC0_BUFFER_FULL_DRP (\d+)
Value MC_VL0_SRC1_BUFFER_FULL_DRP (\d+)
Value MC_VL1_SRC1_BUFFER_FULL_DRP (\d+)
Value MC_VL2_SRC1_BUFFER_FULL_DRP (\d+)
Value MC_VL3_SRC1_BUFFER_FULL_DRP (\d+)

Start
  ^\s*\*+\s+\S+\s+\*+ -> Continue.Record
  ^\s*\*+\s+${FIA}\s+\*+
  ^Category:\s+${CATEGORY}
  ^\s+From\s+Xbar\s+Uc\s+Crc-0\s+${FROM_XBAR_UC_CRC_0}
  ^\s+From\s+Xbar\s+Uc\s+Crc-1\s+${FROM_XBAR_UC_CRC_1}
  ^\s+From\s+Xbar\s+Uc\s+Crc-2\s+${FROM_XBAR_UC_CRC_2}
  ^\s+From\s+Xbar\s+Uc\s+Crc-3\s+${FROM_XBAR_UC_CRC_3}
  ^\s+From\s+Xbar\s+Uc\s+Drp-0\s+${FROM_XBAR_UC_DRP_0}
  ^\s+From\s+Xbar\s+Uc\s+Drp-1\s+${FROM_XBAR_UC_DRP_1}
  ^\s+From\s+Xbar\s+Uc\s+Drp-2\s+${FROM_XBAR_UC_DRP_2}
  ^\s+From\s+Xbar\s+Uc\s+Drp-3\s+${FROM_XBAR_UC_DRP_3}
  ^\s+From\s+Xbar\s+Mc\s+Crc-0\s+${FROM_XBAR_MC_CRC_0}
  ^\s+From\s+Xbar\s+Mc\s+Crc-1\s+${FROM_XBAR_MC_CRC_1}
  ^\s+From\s+Xbar\s+Mc\s+Crc-2\s+${FROM_XBAR_MC_CRC_2}
  ^\s+From\s+Xbar\s+Mc\s+Crc-3\s+${FROM_XBAR_MC_CRC_3}
  ^\s+From\s+Xbar\s+Mc\s+Drp-0\s+${FROM_XBAR_MC_DRP_0}
  ^\s+From\s+Xbar\s+Mc\s+Drp-1\s+${FROM_XBAR_MC_DRP_1}
  ^\s+From\s+Xbar\s+Mc\s+Drp-2\s+${FROM_XBAR_MC_DRP_2}
  ^\s+From\s+Xbar\s+Mc\s+Drp-3\s+${FROM_XBAR_MC_DRP_3}
  ^\s+Uc\s+dq\s+pkt-len-crc/RO-seq/len\s+error\s+drp\s+${UC_DQ_PKT_LEN_CRC_RO_SEQ_LEN_ERROR_DRP}
  ^\s+Uc\s+eq\s+pkt-len-crc/lookup-drp\s+${UC_EQ_PKT_LEN_CRC_LOOKUP_DRP}
  ^\s+Mc\s+rf\s+crc\s+drp\s+${MC_RF_CRC_DRP}
  ^\s+Mc\s+vl0\s+src0\s+buffer\s+full\s+drp\s+${MC_VL0_SRC0_BUFFER_FULL_DRP}
  ^\s+Mc\s+vl1\s+src0\s+buffer\s+full\s+drp\s+${MC_VL1_SRC0_BUFFER_FULL_DRP}
  ^\s+Mc\s+vl2\s+src0\s+buffer\s+full\s+drp\s+${MC_VL2_SRC0_BUFFER_FULL_DRP}
  ^\s+Mc\s+vl3\s+src0\s+buffer\s+full\s+drp\s+${MC_VL3_SRC0_BUFFER_FULL_DRP}
  ^\s+Mc\s+vl0\s+src1\s+buffer\s+full\s+drp\s+${MC_VL0_SRC1_BUFFER_FULL_DRP}
  ^\s+Mc\s+vl1\s+src1\s+buffer\s+full\s+drp\s+${MC_VL1_SRC1_BUFFER_FULL_DRP}
  ^\s+Mc\s+vl2\s+src1\s+buffer\s+full\s+drp\s+${MC_VL2_SRC1_BUFFER_FULL_DRP}
  ^\s+Mc\s+vl3\s+src1\s+buffer\s+full\s+drp\s+${MC_VL3_SRC1_BUFFER_FULL_DRP}
  ^\s*$$
  ^.* -> Error "LINE NOT FOUND"

`