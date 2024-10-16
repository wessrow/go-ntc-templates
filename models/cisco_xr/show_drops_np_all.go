package cisco_xr

type ShowDropsNpAll struct {
	Modify_punt_reason_miss_drop                   string `json:"MODIFY_PUNT_REASON_MISS_DROP"`
	Rsv_drop_ipm4_no_olist                         string `json:"RSV_DROP_IPM4_NO_OLIST"`
	Unknown_l2_on_l3_discard                       string `json:"UNKNOWN_L2_ON_L3_DISCARD"`
	Drop_frm_crc_err_sgmii1                        string `json:"DROP_FRM_CRC_ERR_SGMII1"`
	Drop_frm_frm_err_sgmii3                        string `json:"DROP_FRM_FRM_ERR_SGMII3"`
	Drop_frm_frm_err_xaui5                         string `json:"DROP_FRM_FRM_ERR_XAUI5"`
	Punt_ifib_ospf_opt_excd                        string `json:"PUNT_IFIB_OSPF_OPT_EXCD"`
	Rsv_drop_mpls_txadj_no_match                   string `json:"RSV_DROP_MPLS_TXADJ_NO_MATCH"`
	Drop_frm_crc_err_ilkn1                         string `json:"DROP_FRM_CRC_ERR_ILKN1"`
	Drop_frm_frm_err_ilkn1                         string `json:"DROP_FRM_FRM_ERR_ILKN1"`
	Parse_fast_discard_low_priority_drop_0         string `json:"PARSE_FAST_DISCARD_LOW_PRIORITY_DROP_0"`
	Parse_unknown_nph_type_drop                    string `json:"PARSE_UNKNOWN_NPH_TYPE_DROP"`
	Rsv_drop_ipm4_egr_filter_drop                  string `json:"RSV_DROP_IPM4_EGR_FILTER_DROP"`
	Rsv_drop_ipm4_egr_ttl_drop                     string `json:"RSV_DROP_IPM4_EGR_TTL_DROP"`
	Rsv_drop_nhindex                               string `json:"RSV_DROP_NHINDEX"`
	Drop_frm_frm_err_sgmii0                        string `json:"DROP_FRM_FRM_ERR_SGMII0"`
	Drop_frm_frm_err_xaui6                         string `json:"DROP_FRM_FRM_ERR_XAUI6"`
	Rsv_drop_ipm4_no_olist_rep                     string `json:"RSV_DROP_IPM4_NO_OLIST_REP"`
	Drop_frm_crc_err_ilkn0                         string `json:"DROP_FRM_CRC_ERR_ILKN0"`
	Rsv_drop_ipm4_egr_rpf_fail_drop                string `json:"RSV_DROP_IPM4_EGR_RPF_FAIL_DROP"`
	Rsv_drop_egr_uidb_no_match                     string `json:"RSV_DROP_EGR_UIDB_NO_MATCH"`
	Rsv_drop_mpls_rxadj_drop                       string `json:"RSV_DROP_MPLS_RXADJ_DROP"`
	Drop_frm_frm_err_ilkn2                         string `json:"DROP_FRM_FRM_ERR_ILKN2"`
	Parse_drop_in_uidb_tcam_miss                   string `json:"PARSE_DROP_IN_UIDB_TCAM_MISS"`
	Np_section                                     string `json:"NP_SECTION"`
	Rsv_drop_ipv4_rxadj_drop                       string `json:"RSV_DROP_IPV4_RXADJ_DROP"`
	Rsv_drop_mpls_nrldi_no_match                   string `json:"RSV_DROP_MPLS_NRLDI_NO_MATCH"`
	Drop_frm_runt                                  string `json:"DROP_FRM_RUNT"`
	Mpls_ttl_one_punt_excd                         string `json:"MPLS_TTL_ONE_PUNT_EXCD"`
	Parse_drop_ipv4_disabled                       string `json:"PARSE_DROP_IPV4_DISABLED"`
	Parse_l3_tagged_punt_drop                      string `json:"PARSE_L3_TAGGED_PUNT_DROP"`
	Parse_fast_discard_low_priority_drop_1         string `json:"PARSE_FAST_DISCARD_LOW_PRIORITY_DROP_1"`
	Punt_adj_excd                                  string `json:"PUNT_ADJ_EXCD"`
	Rsv_egr_lag_not_local_drop_cnt                 string `json:"RSV_EGR_LAG_NOT_LOCAL_DROP_CNT"`
	Bfd_neighbor_drop                              string `json:"BFD_NEIGHBOR_DROP"`
	Drop_frm_crc_err_xaui5                         string `json:"DROP_FRM_CRC_ERR_XAUI5"`
	Rsv_drop_mpls_leaf_no_match                    string `json:"RSV_DROP_MPLS_LEAF_NO_MATCH"`
	Mdf_punt_police_drop                           string `json:"MDF_PUNT_POLICE_DROP"`
	Rsv_drop_ipv4_drop_null_rte                    string `json:"RSV_DROP_IPV4_DROP_NULL_RTE"`
	Ipv4_ttl_error_excd                            string `json:"IPV4_TTL_ERROR_EXCD"`
	Parse_fast_discard_low_priority_drop_1_monitor string `json:"PARSE_FAST_DISCARD_LOW_PRIORITY_DROP_1_MONITOR"`
	Parse_drop_ipv4_checksum_error                 string `json:"PARSE_DROP_IPV4_CHECKSUM_ERROR"`
	Punt_ipv4_adj_null_rte_excd                    string `json:"PUNT_IPV4_ADJ_NULL_RTE_EXCD"`
	Rsv_drop_egr_lag_no_match                      string `json:"RSV_DROP_EGR_LAG_NO_MATCH"`
	Np                                             string `json:"NP"`
	Drop_frm_crc_err_sgmii3                        string `json:"DROP_FRM_CRC_ERR_SGMII3"`
	Parse_egr_inj_pkt_typ_unknown                  string `json:"PARSE_EGR_INJ_PKT_TYP_UNKNOWN"`
	Rsv_drop_ing_bfd                               string `json:"RSV_DROP_ING_BFD"`
	Rsv_drop_mpls_nrldi_not_local                  string `json:"RSV_DROP_MPLS_NRLDI_NOT_LOCAL"`
	Location                                       string `json:"LOCATION"`
	Parse_drop_in_uidb_down                        string `json:"PARSE_DROP_IN_UIDB_DOWN"`
	Parse_drop_ipv4_length_error                   string `json:"PARSE_DROP_IPV4_LENGTH_ERROR"`
	Rsv_drop_ipv4_nrldi_not_local                  string `json:"RSV_DROP_IPV4_NRLDI_NOT_LOCAL"`
	Rsv_drop_ipv4_txadj_no_match                   string `json:"RSV_DROP_IPV4_TXADJ_NO_MATCH"`
	Rsv_drop_mpls_leaf_no_match_monitor            string `json:"RSV_DROP_MPLS_LEAF_NO_MATCH_MONITOR"`
	Rsv_mldp_egr_drop                              string `json:"RSV_MLDP_EGR_DROP"`
	Drop_frm_frm_err_xaui4                         string `json:"DROP_FRM_FRM_ERR_XAUI4"`
	Mdf_rpf_fail_drop                              string `json:"MDF_RPF_FAIL_DROP"`
	Parse_open_network_service_key_action_unknown  string `json:"PARSE_OPEN_NETWORK_SERVICE_KEY_ACTION_UNKNOWN"`
	Punt_statistics_excd                           string `json:"PUNT_STATISTICS_EXCD"`
	Parse_fast_discard_low_priority_drop_0_monitor string `json:"PARSE_FAST_DISCARD_LOW_PRIORITY_DROP_0_MONITOR"`
	Rsv_drop_egr_uidb_down                         string `json:"RSV_DROP_EGR_UIDB_DOWN"`
	Rsv_drop_ifib_ttl_check                        string `json:"RSV_DROP_IFIB_TTL_CHECK"`
	Drop_frm_crc_err_sgmii2                        string `json:"DROP_FRM_CRC_ERR_SGMII2"`
	Drop_frm_crc_err_xaui4                         string `json:"DROP_FRM_CRC_ERR_XAUI4"`
	Drop_frm_crc_err_xaui6                         string `json:"DROP_FRM_CRC_ERR_XAUI6"`
	Drop_frm_frm_err_ilkn0                         string `json:"DROP_FRM_FRM_ERR_ILKN0"`
	Drop_frm_frm_err_sgmii1                        string `json:"DROP_FRM_FRM_ERR_SGMII1"`
	Punt_no_match_excd                             string `json:"PUNT_NO_MATCH_EXCD"`
	Rsv_drop_in_l3_not_mymac                       string `json:"RSV_DROP_IN_L3_NOT_MYMAC"`
	Drop_frm_crc_err_ilkn2                         string `json:"DROP_FRM_CRC_ERR_ILKN2"`
	Drop_frm_crc_err_sgmii0                        string `json:"DROP_FRM_CRC_ERR_SGMII0"`
	Ipv4_bfd_excd                                  string `json:"IPV4_BFD_EXCD"`
	Ipv4_frag_needed_punt_excd                     string `json:"IPV4_FRAG_NEEDED_PUNT_EXCD"`
	Bfd_remote_punt_disc_0_drop                    string `json:"BFD_REMOTE_PUNT_DISC_0_DROP"`
	Drop_frm_frm_err_sgmii2                        string `json:"DROP_FRM_FRM_ERR_SGMII2"`
}

var ShowDropsNpAll_Template string = `Value Filldown LOCATION (\S+)
Value Required NP (\d+)
Value BFD_NEIGHBOR_DROP (\d+)
Value BFD_REMOTE_PUNT_DISC_0_DROP (\d+)
Value DROP_FRM_CRC_ERR_ILKN0 (\d+)
Value DROP_FRM_CRC_ERR_ILKN1 (\d+)
Value DROP_FRM_CRC_ERR_ILKN2 (\d+)
Value DROP_FRM_CRC_ERR_SGMII0 (\d+)
Value DROP_FRM_CRC_ERR_SGMII1 (\d+)
Value DROP_FRM_CRC_ERR_SGMII2 (\d+)
Value DROP_FRM_CRC_ERR_SGMII3 (\d+)
Value DROP_FRM_CRC_ERR_XAUI4 (\d+)
Value DROP_FRM_CRC_ERR_XAUI5 (\d+)
Value DROP_FRM_CRC_ERR_XAUI6 (\d+)
Value DROP_FRM_FRM_ERR_ILKN0 (\d+)
Value DROP_FRM_FRM_ERR_ILKN1 (\d+)
Value DROP_FRM_FRM_ERR_ILKN2 (\d+)
Value DROP_FRM_FRM_ERR_SGMII0 (\d+)
Value DROP_FRM_FRM_ERR_SGMII1 (\d+)
Value DROP_FRM_FRM_ERR_SGMII2 (\d+)
Value DROP_FRM_FRM_ERR_SGMII3 (\d+)
Value DROP_FRM_FRM_ERR_XAUI4 (\d+)
Value DROP_FRM_FRM_ERR_XAUI5 (\d+)
Value DROP_FRM_FRM_ERR_XAUI6 (\d+)
Value DROP_FRM_RUNT (\d+)
Value IPV4_BFD_EXCD (\d+)
Value IPV4_FRAG_NEEDED_PUNT_EXCD (\d+)
Value IPV4_TTL_ERROR_EXCD (\d+)
Value MDF_PUNT_POLICE_DROP (\d+)
Value MDF_RPF_FAIL_DROP (\d+)
Value MODIFY_PUNT_REASON_MISS_DROP (\d+)
Value MPLS_TTL_ONE_PUNT_EXCD (\d+)
Value NP_SECTION (\d+)
Value PARSE_FAST_DISCARD_LOW_PRIORITY_DROP_0 (\d+)
Value PARSE_FAST_DISCARD_LOW_PRIORITY_DROP_1 (\d+)
Value PARSE_FAST_DISCARD_LOW_PRIORITY_DROP_0_MONITOR (\d+)
Value PARSE_FAST_DISCARD_LOW_PRIORITY_DROP_1_MONITOR (\d+)
Value PARSE_DROP_IN_UIDB_DOWN (\d+)
Value PARSE_DROP_IN_UIDB_TCAM_MISS (\d+)
Value PARSE_DROP_IPV4_DISABLED (\d+)
Value PARSE_DROP_IPV4_LENGTH_ERROR (\d+)
Value PARSE_DROP_IPV4_CHECKSUM_ERROR (\d+)
Value PARSE_EGR_INJ_PKT_TYP_UNKNOWN (\d+)
Value PARSE_L3_TAGGED_PUNT_DROP (\d+)
Value PARSE_OPEN_NETWORK_SERVICE_KEY_ACTION_UNKNOWN (\d+)
Value PARSE_UNKNOWN_NPH_TYPE_DROP (\d+)
Value PUNT_ADJ_EXCD (\d+)
Value PUNT_IFIB_OSPF_OPT_EXCD (\d+)
Value PUNT_IPV4_ADJ_NULL_RTE_EXCD (\d+)
Value PUNT_NO_MATCH_EXCD (\d+)
Value PUNT_STATISTICS_EXCD (\d+)
Value RSV_DROP_EGR_LAG_NO_MATCH (\d+)
Value RSV_DROP_EGR_UIDB_DOWN (\d+)
Value RSV_DROP_EGR_UIDB_NO_MATCH (\d+)
Value RSV_DROP_IFIB_TTL_CHECK (\d+)
Value RSV_DROP_IN_L3_NOT_MYMAC (\d+)
Value RSV_DROP_ING_BFD (\d+)
Value RSV_DROP_IPM4_EGR_FILTER_DROP (\d+)
Value RSV_DROP_IPM4_EGR_RPF_FAIL_DROP (\d+)
Value RSV_DROP_IPM4_EGR_TTL_DROP (\d+)
Value RSV_DROP_IPM4_NO_OLIST (\d+)
Value RSV_DROP_IPM4_NO_OLIST_REP (\d+)
Value RSV_DROP_IPV4_DROP_NULL_RTE (\d+)
Value RSV_DROP_IPV4_NRLDI_NOT_LOCAL (\d+)
Value RSV_DROP_IPV4_RXADJ_DROP (\d+)
Value RSV_DROP_IPV4_TXADJ_NO_MATCH (\d+)
Value RSV_DROP_MPLS_LEAF_NO_MATCH (\d+)
Value RSV_DROP_MPLS_LEAF_NO_MATCH_MONITOR (\d+)
Value RSV_DROP_MPLS_NRLDI_NO_MATCH (\d+)
Value RSV_DROP_MPLS_NRLDI_NOT_LOCAL (\d+)
Value RSV_DROP_MPLS_RXADJ_DROP (\d+)
Value RSV_DROP_MPLS_TXADJ_NO_MATCH (\d+)
Value RSV_DROP_NHINDEX (\d+)
Value RSV_EGR_LAG_NOT_LOCAL_DROP_CNT (\d+)
Value RSV_MLDP_EGR_DROP (\d+)
Value UNKNOWN_L2_ON_L3_DISCARD (\d+)

Start
  ^\s*Node:\s+\S+ -> Continue.Record
  ^\s*Node:\s+${LOCATION}:
  ^-+
  ^(?:No\s+)?NP -> Continue.Record
  ^No\s+NP\s+${NP}\s+Drops
  ^NP\s+${NP}\s+Drops:
  ^BFD_NEIGHBOR_DROP\s+${BFD_NEIGHBOR_DROP}
  ^BFD_REMOTE_PUNT_DISC_0_DROP\s+${BFD_REMOTE_PUNT_DISC_0_DROP}
  ^DROP_FRM_CRC_ERR_ILKN0\s+${DROP_FRM_CRC_ERR_ILKN0}
  ^DROP_FRM_CRC_ERR_ILKN1\s+${DROP_FRM_CRC_ERR_ILKN1}
  ^DROP_FRM_CRC_ERR_ILKN2\s+${DROP_FRM_CRC_ERR_ILKN2}
  ^DROP_FRM_CRC_ERR_SGMII0\s+${DROP_FRM_CRC_ERR_SGMII0}
  ^DROP_FRM_CRC_ERR_SGMII1\s+${DROP_FRM_CRC_ERR_SGMII1}
  ^DROP_FRM_CRC_ERR_SGMII2\s+${DROP_FRM_CRC_ERR_SGMII2}
  ^DROP_FRM_CRC_ERR_SGMII3\s+${DROP_FRM_CRC_ERR_SGMII3}
  ^DROP_FRM_CRC_ERR_XAUI4\s+${DROP_FRM_CRC_ERR_XAUI4}
  ^DROP_FRM_CRC_ERR_XAUI5\s+${DROP_FRM_CRC_ERR_XAUI5}
  ^DROP_FRM_CRC_ERR_XAUI6\s+${DROP_FRM_CRC_ERR_XAUI6}
  ^DROP_FRM_FRM_ERR_ILKN0\s+${DROP_FRM_FRM_ERR_ILKN0}
  ^DROP_FRM_FRM_ERR_ILKN1\s+${DROP_FRM_FRM_ERR_ILKN1}
  ^DROP_FRM_FRM_ERR_ILKN2\s+${DROP_FRM_FRM_ERR_ILKN2}
  ^DROP_FRM_FRM_ERR_SGMII0\s+${DROP_FRM_FRM_ERR_SGMII0}
  ^DROP_FRM_FRM_ERR_SGMII1\s+${DROP_FRM_FRM_ERR_SGMII1}
  ^DROP_FRM_FRM_ERR_SGMII2\s+${DROP_FRM_FRM_ERR_SGMII2}
  ^DROP_FRM_FRM_ERR_SGMII3\s+${DROP_FRM_FRM_ERR_SGMII3}
  ^DROP_FRM_FRM_ERR_XAUI4\s+${DROP_FRM_FRM_ERR_XAUI4}
  ^DROP_FRM_FRM_ERR_XAUI5\s+${DROP_FRM_FRM_ERR_XAUI5}
  ^DROP_FRM_FRM_ERR_XAUI6\s+${DROP_FRM_FRM_ERR_XAUI6}
  ^DROP_FRM_RUNT\s+${DROP_FRM_RUNT}
  ^IPV4_BFD_EXCD\s+${IPV4_BFD_EXCD}
  ^IPV4_FRAG_NEEDED_PUNT_EXCD\s+${IPV4_FRAG_NEEDED_PUNT_EXCD}
  ^IPV4_TTL_ERROR_EXCD\s+${IPV4_TTL_ERROR_EXCD}
  ^MDF_PUNT_POLICE_DROP\s+${MDF_PUNT_POLICE_DROP}
  ^MDF_RPF_FAIL_DROP\s+${MDF_RPF_FAIL_DROP}
  ^MODIFY_PUNT_REASON_MISS_DROP\s+${MODIFY_PUNT_REASON_MISS_DROP}
  ^MPLS_TTL_ONE_PUNT_EXCD\s+${MPLS_TTL_ONE_PUNT_EXCD}
  ^NP_SECTION\s+${NP_SECTION}
  ^PARSE_FAST_DISCARD_LOW_PRIORITY_DROP_0\s+${PARSE_FAST_DISCARD_LOW_PRIORITY_DROP_0}
  ^PARSE_FAST_DISCARD_LOW_PRIORITY_DROP_1\s+${PARSE_FAST_DISCARD_LOW_PRIORITY_DROP_1}
  ^PARSE_FAST_DISCARD_LOW_PRIORITY_DROP_0_MONITOR\s+${PARSE_FAST_DISCARD_LOW_PRIORITY_DROP_0_MONITOR}
  ^PARSE_FAST_DISCARD_LOW_PRIORITY_DROP_1_MONITOR\s+${PARSE_FAST_DISCARD_LOW_PRIORITY_DROP_1_MONITOR}
  ^PARSE_DROP_IN_UIDB_DOWN\s+${PARSE_DROP_IN_UIDB_DOWN}
  ^PARSE_DROP_IN_UIDB_TCAM_MISS\s+${PARSE_DROP_IN_UIDB_TCAM_MISS}
  ^PARSE_DROP_IPV4_DISABLED\s+${PARSE_DROP_IPV4_DISABLED}
  ^PARSE_DROP_IPV4_CHECKSUM_ERROR\s+${PARSE_DROP_IPV4_CHECKSUM_ERROR}
  ^PARSE_DROP_IPV4_LENGTH_ERROR\s+${PARSE_DROP_IPV4_LENGTH_ERROR}
  ^PARSE_L3_TAGGED_PUNT_DROP\s+${PARSE_L3_TAGGED_PUNT_DROP}
  ^PARSE_OPEN_NETWORK_SERVICE_KEY_ACTION_UNKNOWN\s+${PARSE_OPEN_NETWORK_SERVICE_KEY_ACTION_UNKNOWN}
  ^PARSE_EGR_INJ_PKT_TYP_UNKNOWN\s+${PARSE_EGR_INJ_PKT_TYP_UNKNOWN}
  ^PARSE_UNKNOWN_NPH_TYPE_DROP\s+${PARSE_UNKNOWN_NPH_TYPE_DROP}
  ^PUNT_ADJ_EXCD\s+${PUNT_ADJ_EXCD}
  ^PUNT_IFIB_OSPF_OPT_EXCD\s+${PUNT_IFIB_OSPF_OPT_EXCD}
  ^PUNT_IPV4_ADJ_NULL_RTE_EXCD\s+${PUNT_IPV4_ADJ_NULL_RTE_EXCD}
  ^PUNT_NO_MATCH_EXCD\s+${PUNT_NO_MATCH_EXCD}
  ^PUNT_STATISTICS_EXCD\s+${PUNT_STATISTICS_EXCD}
  ^RSV_DROP_EGR_LAG_NO_MATCH\s+${RSV_DROP_EGR_LAG_NO_MATCH}
  ^RSV_DROP_EGR_UIDB_DOWN\s+${RSV_DROP_EGR_UIDB_DOWN}
  ^RSV_DROP_EGR_UIDB_NO_MATCH\s+${RSV_DROP_EGR_UIDB_NO_MATCH}
  ^RSV_DROP_IFIB_TTL_CHECK\s+${RSV_DROP_IFIB_TTL_CHECK}
  ^RSV_DROP_IN_L3_NOT_MYMAC\s+${RSV_DROP_IN_L3_NOT_MYMAC}
  ^RSV_DROP_ING_BFD\s+${RSV_DROP_ING_BFD}
  ^RSV_DROP_IPM4_EGR_FILTER_DROP\s+${RSV_DROP_IPM4_EGR_FILTER_DROP}
  ^RSV_DROP_IPM4_EGR_RPF_FAIL_DROP\s+${RSV_DROP_IPM4_EGR_RPF_FAIL_DROP}
  ^RSV_DROP_IPM4_EGR_TTL_DROP\s+${RSV_DROP_IPM4_EGR_TTL_DROP}
  ^RSV_DROP_IPM4_NO_OLIST\s+${RSV_DROP_IPM4_NO_OLIST}
  ^RSV_DROP_IPM4_NO_OLIST_REP\s+${RSV_DROP_IPM4_NO_OLIST_REP}
  ^RSV_DROP_IPV4_DROP_NULL_RTE\s+${RSV_DROP_IPV4_DROP_NULL_RTE}
  ^RSV_DROP_IPV4_NRLDI_NOT_LOCAL\s+${RSV_DROP_IPV4_NRLDI_NOT_LOCAL}
  ^RSV_DROP_IPV4_RXADJ_DROP\s+${RSV_DROP_IPV4_RXADJ_DROP}
  ^RSV_DROP_IPV4_TXADJ_NO_MATCH\s+${RSV_DROP_IPV4_TXADJ_NO_MATCH}
  ^RSV_DROP_MPLS_LEAF_NO_MATCH\s+${RSV_DROP_MPLS_LEAF_NO_MATCH}
  ^RSV_DROP_MPLS_LEAF_NO_MATCH_MONITOR\s+${RSV_DROP_MPLS_LEAF_NO_MATCH_MONITOR}
  ^RSV_DROP_MPLS_NRLDI_NO_MATCH\s+${RSV_DROP_MPLS_NRLDI_NO_MATCH}
  ^RSV_DROP_MPLS_NRLDI_NOT_LOCAL\s+${RSV_DROP_MPLS_NRLDI_NOT_LOCAL}
  ^RSV_DROP_MPLS_RXADJ_DROP\s+${RSV_DROP_MPLS_RXADJ_DROP}
  ^RSV_DROP_MPLS_TXADJ_NO_MATCH\s+${RSV_DROP_MPLS_TXADJ_NO_MATCH}
  ^RSV_DROP_NHINDEX\s+${RSV_DROP_NHINDEX}
  ^RSV_EGR_LAG_NOT_LOCAL_DROP_CNT\s+${RSV_EGR_LAG_NOT_LOCAL_DROP_CNT}
  ^RSV_MLDP_EGR_DROP\s+${RSV_MLDP_EGR_DROP}
  ^UNKNOWN_L2_ON_L3_DISCARD\s+${UNKNOWN_L2_ON_L3_DISCARD}
  ^\s*$$
  ^. -> Error "LINE NOT FOUND"
`
