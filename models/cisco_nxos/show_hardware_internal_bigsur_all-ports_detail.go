package cisco_nxos 

type ShowHardwareInternalBigsurAllPortsDetail struct {
	Port_name	string	`json:"PORT_NAME"`
	Interface_index	string	`json:"INTERFACE_INDEX"`
	Instance_index	string	`json:"INSTANCE_INDEX"`
	Instance_slot	string	`json:"INSTANCE_SLOT"`
	Instance_asic	string	`json:"INSTANCE_ASIC"`
	Asic_eport	string	`json:"ASIC_EPORT"`
	Logical_port	string	`json:"LOGICAL_PORT"`
	Front_port	string	`json:"FRONT_PORT"`
	State	string	`json:"STATE"`
	Bigsur_eport	string	`json:"BIGSUR_EPORT"`
	Port	string	`json:"PORT"`
	Port_dec	string	`json:"PORT_DEC"`
	Type	string	`json:"TYPE"`
	Firewall_instance	string	`json:"FIREWALL_INSTANCE"`
	Speed	string	`json:"SPEED"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Admin_state_dec	string	`json:"ADMIN_STATE_DEC"`
	Mtu	string	`json:"MTU"`
	Duplexity	string	`json:"DUPLEXITY"`
	Duplexity_dec	string	`json:"DUPLEXITY_DEC"`
	Auto_neg_state	string	`json:"AUTO_NEG_STATE"`
	Auto_neg_state_dec	string	`json:"AUTO_NEG_STATE_DEC"`
	Fast_1g_an_timer	string	`json:"FAST_1G_AN_TIMER"`
	Port_speed	string	`json:"PORT_SPEED"`
	Debounce_time	string	`json:"DEBOUNCE_TIME"`
	Linkup_debounce_time	string	`json:"LINKUP_DEBOUNCE_TIME"`
	Sat_fabric_mode	string	`json:"SAT_FABRIC_MODE"`
	Sat_fabric_mode_hex	string	`json:"SAT_FABRIC_MODE_HEX"`
	Port_flow_ctrl	string	`json:"PORT_FLOW_CTRL"`
	Class_flow_ctrl7_0	string	`json:"CLASS_FLOW_CTRL7_0"`
	Class_flow_ctrl_msk_rx	string	`json:"CLASS_FLOW_CTRL_MSK_RX"`
	Class_flow_ctrl_msk_tx	string	`json:"CLASS_FLOW_CTRL_MSK_TX"`
	Link_state	string	`json:"LINK_STATE"`
	Link_state_dec	string	`json:"LINK_STATE_DEC"`
	Oper_speed	string	`json:"OPER_SPEED"`
	Oper_duplexity	string	`json:"OPER_DUPLEXITY"`
	Oper_duplexity_dec	string	`json:"OPER_DUPLEXITY_DEC"`
	Oper_port_flow_ctrl	string	`json:"OPER_PORT_FLOW_CTRL"`
	Up_count	string	`json:"UP_COUNT"`
	Down_count	string	`json:"DOWN_COUNT"`
	Slot	string	`json:"SLOT"`
	Mm_inst	string	`json:"MM_INST"`
	Mm_mode	string	`json:"MM_MODE"`
	Diag_sts	string	`json:"DIAG_STS"`
	Mac_status_polling	string	`json:"MAC_STATUS_POLLING"`
	An_state	string	`json:"AN_STATE"`
	An_state_dec	string	`json:"AN_STATE_DEC"`
	An_timer	string	`json:"AN_TIMER"`
	Link_fail_inhibit_timer	string	`json:"LINK_FAIL_INHIBIT_TIMER"`
	An_disabled	string	`json:"AN_DISABLED"`
	Mm_port_state	string	`json:"MM_PORT_STATE"`
	Mm_port_state_dec	string	`json:"MM_PORT_STATE_DEC"`
	Mm_port_timer	string	`json:"MM_PORT_TIMER"`
	Remote_fault_disable	string	`json:"REMOTE_FAULT_DISABLE"`
	Adaptive_tuning_disable	string	`json:"ADAPTIVE_TUNING_DISABLE"`
	Link_recovery_disable	string	`json:"LINK_RECOVERY_DISABLE"`
	Link_bring_down_disable	string	`json:"LINK_BRING_DOWN_DISABLE"`
	Dfe_tuning_disable	string	`json:"DFE_TUNING_DISABLE"`
	Force_qsfp_removal	string	`json:"FORCE_QSFP_REMOVAL"`
	Force_qsfp_insertion	string	`json:"FORCE_QSFP_INSERTION"`
	Force_qsfp_fiber_type	string	`json:"FORCE_QSFP_FIBER_TYPE"`
	Force_qsfp_copper_type	string	`json:"FORCE_QSFP_COPPER_TYPE"`
	Serdes_lanes	string	`json:"SERDES_LANES"`
	Dfe_tune_lane_0_dfe_a	string	`json:"DFE_TUNE_LANE_0_DFE_A"`
	Dfe_tune_lane_0_dfe_b	string	`json:"DFE_TUNE_LANE_0_DFE_B"`
	Dfe_tune_lane_0_dfe_tune	string	`json:"DFE_TUNE_LANE_0_DFE_TUNE"`
	Dfe_tune_lane_1_dfe_a	string	`json:"DFE_TUNE_LANE_1_DFE_A"`
	Dfe_tune_lane_1_dfe_b	string	`json:"DFE_TUNE_LANE_1_DFE_B"`
	Dfe_tune_lane_1_dfe_tune	string	`json:"DFE_TUNE_LANE_1_DFE_TUNE"`
	Dfe_tune_lane_2_dfe_a	string	`json:"DFE_TUNE_LANE_2_DFE_A"`
	Dfe_tune_lane_2_dfe_b	string	`json:"DFE_TUNE_LANE_2_DFE_B"`
	Dfe_tune_lane_2_dfe_tune	string	`json:"DFE_TUNE_LANE_2_DFE_TUNE"`
	Dfe_tune_lane_3_dfe_a	string	`json:"DFE_TUNE_LANE_3_DFE_A"`
	Dfe_tune_lane_3_dfe_b	string	`json:"DFE_TUNE_LANE_3_DFE_B"`
	Dfe_tune_lane_3_dfe_tune	string	`json:"DFE_TUNE_LANE_3_DFE_TUNE"`
	Fi_block_lock_0h	string	`json:"FI_BLOCK_LOCK_0H"`
	Fi_block_lock_1h	string	`json:"FI_BLOCK_LOCK_1H"`
	Fi_block_lock_2h	string	`json:"FI_BLOCK_LOCK_2H"`
	Fi_block_lock_3h	string	`json:"FI_BLOCK_LOCK_3H"`
	Tx_pkt_size_lt_64	string	`json:"TX_PKT_SIZE_LT_64"`
	Tx_pkt_size_is_64	string	`json:"TX_PKT_SIZE_IS_64"`
	Tx_pkt_size_is_65_to_127	string	`json:"TX_PKT_SIZE_IS_65_TO_127"`
	Tx_pkt_size_is_128_to_255	string	`json:"TX_PKT_SIZE_IS_128_TO_255"`
	Tx_pkt_size_is_256_to_511	string	`json:"TX_PKT_SIZE_IS_256_TO_511"`
	Tx_pkt_size_is_512_to_1023	string	`json:"TX_PKT_SIZE_IS_512_TO_1023"`
	Tx_pkt_size_is_1024_to_1518	string	`json:"TX_PKT_SIZE_IS_1024_TO_1518"`
	Tx_pkt_size_is_1519_to_2047	string	`json:"TX_PKT_SIZE_IS_1519_TO_2047"`
	Tx_pkt_size_is_2048_to_4095	string	`json:"TX_PKT_SIZE_IS_2048_TO_4095"`
	Tx_pkt_size_is_4095_to_8191	string	`json:"TX_PKT_SIZE_IS_4095_TO_8191"`
	Tx_pkt_size_is_8192_to_9216	string	`json:"TX_PKT_SIZE_IS_8192_TO_9216"`
	Tx_pkt_size_gt_9216	string	`json:"TX_PKT_SIZE_GT_9216"`
	Tx_pkt_total	string	`json:"TX_PKT_TOTAL"`
	Tx_pkt_octets	string	`json:"TX_PKT_OCTETS"`
	Tx_pkt_good	string	`json:"TX_PKT_GOOD"`
	Tx_pkt_ucast	string	`json:"TX_PKT_UCAST"`
	Tx_pkt_mcast	string	`json:"TX_PKT_MCAST"`
	Tx_pkt_bcast	string	`json:"TX_PKT_BCAST"`
	Tx_pkt_vlan	string	`json:"TX_PKT_VLAN"`
	Tx_pkt_802_3x_pause	string	`json:"TX_PKT_802_3X_PAUSE"`
	Tx_pkt_per_priority_pause	string	`json:"TX_PKT_PER_PRIORITY_PAUSE"`
	Tx_pkt_frame_error	string	`json:"TX_PKT_FRAME_ERROR"`
	Rx_pkt_size_lt_64	string	`json:"RX_PKT_SIZE_LT_64"`
	Rx_pkt_size_is_64	string	`json:"RX_PKT_SIZE_IS_64"`
	Rx_pkt_size_is_65_to_127	string	`json:"RX_PKT_SIZE_IS_65_TO_127"`
	Rx_pkt_size_is_128_to_255	string	`json:"RX_PKT_SIZE_IS_128_TO_255"`
	Rx_pkt_size_is_256_to_511	string	`json:"RX_PKT_SIZE_IS_256_TO_511"`
	Rx_pkt_size_is_512_to_1023	string	`json:"RX_PKT_SIZE_IS_512_TO_1023"`
	Rx_pkt_size_is_1024_to_1518	string	`json:"RX_PKT_SIZE_IS_1024_TO_1518"`
	Rx_pkt_size_is_1519_to_2047	string	`json:"RX_PKT_SIZE_IS_1519_TO_2047"`
	Rx_pkt_size_is_2048_to_4095	string	`json:"RX_PKT_SIZE_IS_2048_TO_4095"`
	Rx_pkt_size_is_4095_to_8191	string	`json:"RX_PKT_SIZE_IS_4095_TO_8191"`
	Rx_pkt_size_is_8192_to_9216	string	`json:"RX_PKT_SIZE_IS_8192_TO_9216"`
	Rx_pkt_size_gt_9216	string	`json:"RX_PKT_SIZE_GT_9216"`
	Rx_pkt_total	string	`json:"RX_PKT_TOTAL"`
	Rx_pkt_octets	string	`json:"RX_PKT_OCTETS"`
	Rx_pkt_good	string	`json:"RX_PKT_GOOD"`
	Rx_pkt_ucast	string	`json:"RX_PKT_UCAST"`
	Rx_pkt_mcast	string	`json:"RX_PKT_MCAST"`
	Rx_pkt_bcast	string	`json:"RX_PKT_BCAST"`
	Rx_pkt_vlan	string	`json:"RX_PKT_VLAN"`
	Rx_pkt_oversize	string	`json:"RX_PKT_OVERSIZE"`
	Rx_pkt_toolong	string	`json:"RX_PKT_TOOLONG"`
	Rx_pkt_discard	string	`json:"RX_PKT_DISCARD"`
	Rx_pkt_undersize	string	`json:"RX_PKT_UNDERSIZE"`
	Rx_pkt_fragment	string	`json:"RX_PKT_FRAGMENT"`
	Rx_pkt_crc_not_stomped	string	`json:"RX_PKT_CRC_NOT_STOMPED"`
	Rx_pkt_crc_stomped	string	`json:"RX_PKT_CRC_STOMPED"`
	Rx_pkt_in_range_err	string	`json:"RX_PKT_IN_RANGE_ERR"`
	Rx_pkt_jabber	string	`json:"RX_PKT_JABBER"`
	Rx_pkt_802_3x_pause	string	`json:"RX_PKT_802_3X_PAUSE"`
	Rx_pkt_per_priority_pause	string	`json:"RX_PKT_PER_PRIORITY_PAUSE"`
	Tx_pkt_good_octets	string	`json:"TX_PKT_GOOD_OCTETS"`
	Rx_pkt_good_octets	string	`json:"RX_PKT_GOOD_OCTETS"`
	Sfp_ctrl_bigsur	string	`json:"SFP_CTRL_BIGSUR"`
	Sfp_ctrl_port	string	`json:"SFP_CTRL_PORT"`
	Phy_bigsur	string	`json:"PHY_BIGSUR"`
	Phy_mdio_addr	string	`json:"PHY_MDIO_ADDR"`
	Phy_mdio_master	string	`json:"PHY_MDIO_MASTER"`
	Phy_reset_bigsur	string	`json:"PHY_RESET_BIGSUR"`
	Phy_reset_spio	string	`json:"PHY_RESET_SPIO"`
	Phy_prom_shared	string	`json:"PHY_PROM_SHARED"`
	Serdes_txdata_0	string	`json:"SERDES_TXDATA_0"`
	Serdes_txdata_1	string	`json:"SERDES_TXDATA_1"`
	Serdes_txdata_2	string	`json:"SERDES_TXDATA_2"`
	Serdes_txdata_3	string	`json:"SERDES_TXDATA_3"`
	Phy_pmd_tx_0	string	`json:"PHY_PMD_TX_0"`
	Phy_pmd_tx_1	string	`json:"PHY_PMD_TX_1"`
	Phy_pmd_tx_2	string	`json:"PHY_PMD_TX_2"`
	Pma_pmd_devaddr	string	`json:"PMA_PMD_DEVADDR"`
	Pcs_devaddr	string	`json:"PCS_DEVADDR"`
	Phyxs_devaddr	string	`json:"PHYXS_DEVADDR"`
	Xcvr_devaddr	string	`json:"XCVR_DEVADDR"`
	Phy_present	string	`json:"PHY_PRESENT"`
	Phy_type	string	`json:"PHY_TYPE"`
	Sfp_acc_type	string	`json:"SFP_ACC_TYPE"`
	Dfe_timer	string	`json:"DFE_TIMER"`
	Dfe_timer_running	string	`json:"DFE_TIMER_RUNNING"`
	Init_wait_count	string	`json:"INIT_WAIT_COUNT"`
	Do_tx_disable_clr	string	`json:"DO_TX_DISABLE_CLR"`
	Tx_disable_clr_cnt	string	`json:"TX_DISABLE_CLR_CNT"`
	Qsa_status	string	`json:"QSA_STATUS"`
	Connector_type	string	`json:"CONNECTOR_TYPE"`
	Usb_bus_num	string	`json:"USB_BUS_NUM"`
	Usb_dev_num	string	`json:"USB_DEV_NUM"`
	Usb_priv	string	`json:"USB_PRIV"`
	Reset_done	string	`json:"RESET_DONE"`
	Xcvr_flag_present	string	`json:"XCVR_FLAG_PRESENT"`
	Xcvr_flag_xcvr_not_sfp	string	`json:"XCVR_FLAG_XCVR_NOT_SFP"`
	Xcvr_flag_10g_capable	string	`json:"XCVR_FLAG_10G_CAPABLE"`
	Xcvr_flag_cksum_ok	string	`json:"XCVR_FLAG_CKSUM_OK"`
	Xcvr_flag_security_crc_ok	string	`json:"XCVR_FLAG_SECURITY_CRC_OK"`
	Xcvr_flag_cisco_chk_ok	string	`json:"XCVR_FLAG_CISCO_CHK_OK"`
	Xcvr_flag_ddm_capable	string	`json:"XCVR_FLAG_DDM_CAPABLE"`
	Xcvr_flag_pma_pmd_present	string	`json:"XCVR_FLAG_PMA_PMD_PRESENT"`
	Xcvr_flag_pcs_present	string	`json:"XCVR_FLAG_PCS_PRESENT"`
	Xcvr_flag_phyxs_present	string	`json:"XCVR_FLAG_PHYXS_PRESENT"`
	Xcvr_flag_xcvr_operational	string	`json:"XCVR_FLAG_XCVR_OPERATIONAL"`
	Xcvr_flag_pma_pmd_low_pwr_mode	string	`json:"XCVR_FLAG_PMA_PMD_LOW_PWR_MODE"`
	Xcvr_flag_pcs_low_pwr_mode	string	`json:"XCVR_FLAG_PCS_LOW_PWR_MODE"`
	Xcvr_flag_phyxs_low_pwr_mode	string	`json:"XCVR_FLAG_PHYXS_LOW_PWR_MODE"`
	Xcvr_flag_link_up	string	`json:"XCVR_FLAG_LINK_UP"`
	Xcvr_flag_type_fiber	string	`json:"XCVR_FLAG_TYPE_FIBER"`
	Xcvr_flag_sprom_read_ok	string	`json:"XCVR_FLAG_SPROM_READ_OK"`
	Port_state	string	`json:"PORT_STATE"`
	Xcvr_insert_debounce_timer	string	`json:"XCVR_INSERT_DEBOUNCE_TIMER"`
	Xcvr_link_debounce_timer	string	`json:"XCVR_LINK_DEBOUNCE_TIMER"`
	Xcvr_linkup_debounce_timer	string	`json:"XCVR_LINKUP_DEBOUNCE_TIMER"`
	Xcvr_presence_detect_timer	string	`json:"XCVR_PRESENCE_DETECT_TIMER"`
	Tx_enable_signal	string	`json:"TX_ENABLE_SIGNAL"`
	Debounce_timeout	string	`json:"DEBOUNCE_TIMEOUT"`
	Linkup_debounce_timeout	string	`json:"LINKUP_DEBOUNCE_TIMEOUT"`
	Link_up	string	`json:"LINK_UP"`
	Link_dn_debounce_start	string	`json:"LINK_DN_DEBOUNCE_START"`
	Link_debounce_end	string	`json:"LINK_DEBOUNCE_END"`
	Bit_error_rate	string	`json:"BIT_ERROR_RATE"`
	Bit_error_rate_since_linkup	string	`json:"BIT_ERROR_RATE_SINCE_LINKUP"`
	Error_blocks	string	`json:"ERROR_BLOCKS"`
	Error_blocks_since_linkup	string	`json:"ERROR_BLOCKS_SINCE_LINKUP"`
	Link_up_hex	string	`json:"LINK_UP_HEX"`
	Link_up_dec	string	`json:"LINK_UP_DEC"`
	Link_down_hex	string	`json:"LINK_DOWN_HEX"`
	Link_down_dec	string	`json:"LINK_DOWN_DEC"`
	Link_debounced_with_link_up_hex	string	`json:"LINK_DEBOUNCED_WITH_LINK_UP_HEX"`
	Link_debounced_with_link_up_dec	string	`json:"LINK_DEBOUNCED_WITH_LINK_UP_DEC"`
	Link_debounced_with_link_up_since_last_en_hex	string	`json:"LINK_DEBOUNCED_WITH_LINK_UP_SINCE_LAST_EN_HEX"`
	Link_debounced_with_link_up_since_last_en_dec	string	`json:"LINK_DEBOUNCED_WITH_LINK_UP_SINCE_LAST_EN_DEC"`
}

var ShowHardwareInternalBigsurAllPortsDetail_Template = `Value Required PORT_NAME (\S+)
Value Required INTERFACE_INDEX (\S+)
Value INSTANCE_INDEX (\S+)
Value INSTANCE_SLOT (\S+)
Value INSTANCE_ASIC (\S+)
Value ASIC_EPORT (\S+)
Value LOGICAL_PORT (\S+)
Value FRONT_PORT (\S+)
Value STATE ([\S\s]+?)
Value BIGSUR_EPORT (\S+)
Value PORT (\S+)
Value PORT_DEC (\S+)
Value TYPE (\S+)
Value FIREWALL_INSTANCE (\S+)
Value SPEED (\S+)
Value ADMIN_STATE (\S+)
Value ADMIN_STATE_DEC (\S+)
Value MTU (\S+)
Value DUPLEXITY (\S+)
Value DUPLEXITY_DEC (\S+)
Value AUTO_NEG_STATE (\S+)
Value AUTO_NEG_STATE_DEC (\S+)
Value FAST_1G_AN_TIMER (\S+)
Value PORT_SPEED (\S+)
Value DEBOUNCE_TIME (\S+)
Value LINKUP_DEBOUNCE_TIME (\S+)
Value SAT_FABRIC_MODE (\S+)
Value SAT_FABRIC_MODE_HEX (\S+)
Value PORT_FLOW_CTRL (\S+)
Value CLASS_FLOW_CTRL7_0 ([\S\s]+)
Value CLASS_FLOW_CTRL_MSK_RX (\S+)
Value CLASS_FLOW_CTRL_MSK_TX (\S+)
Value LINK_STATE (\S+)
Value LINK_STATE_DEC (\S+)
Value OPER_SPEED (\S+)
Value OPER_DUPLEXITY (\S+)
Value OPER_DUPLEXITY_DEC (\S+)
Value OPER_PORT_FLOW_CTRL (\S+)
Value UP_COUNT (\S+)
Value DOWN_COUNT (\S+)
Value SLOT (\S+)
Value MM_INST (\S+)
Value MM_MODE (\S+)
Value DIAG_STS (\S+)
Value MAC_STATUS_POLLING (\S+)
Value AN_STATE (\S+)
Value AN_STATE_DEC (\S+)
Value AN_TIMER (\S+)
Value LINK_FAIL_INHIBIT_TIMER (\S+)
Value AN_DISABLED (\S+)
Value MM_PORT_STATE (\S+)
Value MM_PORT_STATE_DEC (\S+)
Value MM_PORT_TIMER (\S+)
Value REMOTE_FAULT_DISABLE (\S+)
Value ADAPTIVE_TUNING_DISABLE (\S+)
Value LINK_RECOVERY_DISABLE (\S+)
Value LINK_BRING_DOWN_DISABLE (\S+)
Value DFE_TUNING_DISABLE (\S+)
Value FORCE_QSFP_REMOVAL (\S+)
Value FORCE_QSFP_INSERTION (\S+)
Value FORCE_QSFP_FIBER_TYPE (\S+)
Value FORCE_QSFP_COPPER_TYPE (\S+)
Value SERDES_LANES ([\S\s]+?)
Value DFE_TUNE_LANE_0_DFE_A ([\S\s]+)
Value DFE_TUNE_LANE_0_DFE_B ([\S\s]+)
Value DFE_TUNE_LANE_0_DFE_TUNE ([\S\s]+)
Value DFE_TUNE_LANE_1_DFE_A ([\S\s]+)
Value DFE_TUNE_LANE_1_DFE_B ([\S\s]+)
Value DFE_TUNE_LANE_1_DFE_TUNE ([\S\s]+)
Value DFE_TUNE_LANE_2_DFE_A ([\S\s]+)
Value DFE_TUNE_LANE_2_DFE_B ([\S\s]+)
Value DFE_TUNE_LANE_2_DFE_TUNE ([\S\s]+)
Value DFE_TUNE_LANE_3_DFE_A ([\S\s]+)
Value DFE_TUNE_LANE_3_DFE_B ([\S\s]+)
Value DFE_TUNE_LANE_3_DFE_TUNE ([\S\s]+)
Value FI_BLOCK_LOCK_0H (\S+)
Value FI_BLOCK_LOCK_1H (\S+)
Value FI_BLOCK_LOCK_2H (\S+)
Value FI_BLOCK_LOCK_3H (\S+)
Value TX_PKT_SIZE_LT_64 (\S+)
Value TX_PKT_SIZE_IS_64 (\S+)
Value TX_PKT_SIZE_IS_65_TO_127 (\S+)
Value TX_PKT_SIZE_IS_128_TO_255 (\S+)
Value TX_PKT_SIZE_IS_256_TO_511 (\S+)
Value TX_PKT_SIZE_IS_512_TO_1023 (\S+)
Value TX_PKT_SIZE_IS_1024_TO_1518 (\S+)
Value TX_PKT_SIZE_IS_1519_TO_2047 (\S+)
Value TX_PKT_SIZE_IS_2048_TO_4095 (\S+)
Value TX_PKT_SIZE_IS_4095_TO_8191 (\S+)
Value TX_PKT_SIZE_IS_8192_TO_9216 (\S+)
Value TX_PKT_SIZE_GT_9216 (\S+)
Value TX_PKT_TOTAL (\S+)
Value TX_PKT_OCTETS (\S+)
Value TX_PKT_GOOD (\S+)
Value TX_PKT_UCAST (\S+)
Value TX_PKT_MCAST (\S+)
Value TX_PKT_BCAST (\S+)
Value TX_PKT_VLAN (\S+)
Value TX_PKT_802_3X_PAUSE (\S+)
Value TX_PKT_PER_PRIORITY_PAUSE (\S+)
Value TX_PKT_FRAME_ERROR (\S+)
Value RX_PKT_SIZE_LT_64 (\S+)
Value RX_PKT_SIZE_IS_64 (\S+)
Value RX_PKT_SIZE_IS_65_TO_127 (\S+)
Value RX_PKT_SIZE_IS_128_TO_255 (\S+)
Value RX_PKT_SIZE_IS_256_TO_511 (\S+)
Value RX_PKT_SIZE_IS_512_TO_1023 (\S+)
Value RX_PKT_SIZE_IS_1024_TO_1518 (\S+)
Value RX_PKT_SIZE_IS_1519_TO_2047 (\S+)
Value RX_PKT_SIZE_IS_2048_TO_4095 (\S+)
Value RX_PKT_SIZE_IS_4095_TO_8191 (\S+)
Value RX_PKT_SIZE_IS_8192_TO_9216 (\S+)
Value RX_PKT_SIZE_GT_9216 (\S+)
Value RX_PKT_TOTAL (\S+)
Value RX_PKT_OCTETS (\S+)
Value RX_PKT_GOOD (\S+)
Value RX_PKT_UCAST (\S+)
Value RX_PKT_MCAST (\S+)
Value RX_PKT_BCAST (\S+)
Value RX_PKT_VLAN (\S+)
Value RX_PKT_OVERSIZE (\S+)
Value RX_PKT_TOOLONG (\S+)
Value RX_PKT_DISCARD (\S+)
Value RX_PKT_UNDERSIZE (\S+)
Value RX_PKT_FRAGMENT (\S+)
Value RX_PKT_CRC_NOT_STOMPED (\S+)
Value RX_PKT_CRC_STOMPED (\S+)
Value RX_PKT_IN_RANGE_ERR (\S+)
Value RX_PKT_JABBER (\S+)
Value RX_PKT_802_3X_PAUSE (\S+)
Value RX_PKT_PER_PRIORITY_PAUSE (\S+)
Value TX_PKT_GOOD_OCTETS (\S+)
Value RX_PKT_GOOD_OCTETS (\S+)
Value SFP_CTRL_BIGSUR (\S+)
Value SFP_CTRL_PORT (\S+)
Value PHY_BIGSUR (\S+)
Value PHY_MDIO_ADDR (\S+)
Value PHY_MDIO_MASTER (\S+)
Value PHY_RESET_BIGSUR (\S+)
Value PHY_RESET_SPIO (\S+)
Value PHY_PROM_SHARED (\S+)
Value SERDES_TXDATA_0 (\S+)
Value SERDES_TXDATA_1 (\S+)
Value SERDES_TXDATA_2 (\S+)
Value SERDES_TXDATA_3 (\S+)
Value PHY_PMD_TX_0 (\S+)
Value PHY_PMD_TX_1 (\S+)
Value PHY_PMD_TX_2 (\S+)
Value PMA_PMD_DEVADDR (\S+)
Value PCS_DEVADDR (\S+)
Value PHYXS_DEVADDR (\S+)
Value XCVR_DEVADDR (\S+)
Value PHY_PRESENT (\S+)
Value PHY_TYPE (\S+)
Value SFP_ACC_TYPE (\S+)
Value DFE_TIMER (\S+)
Value DFE_TIMER_RUNNING (\S+)
Value INIT_WAIT_COUNT (\S+)
Value DO_TX_DISABLE_CLR (\S+)
Value TX_DISABLE_CLR_CNT (\S+)
Value QSA_STATUS (\S+)
Value CONNECTOR_TYPE (\S+)
Value USB_BUS_NUM (\S+)
Value USB_DEV_NUM (\S+)
Value USB_PRIV (\S+)
Value RESET_DONE (\S+)
Value XCVR_FLAG_PRESENT (\S+)
Value XCVR_FLAG_XCVR_NOT_SFP (\S+)
Value XCVR_FLAG_10G_CAPABLE (\S+)
Value XCVR_FLAG_CKSUM_OK (\S+)
Value XCVR_FLAG_SECURITY_CRC_OK (\S+)
Value XCVR_FLAG_CISCO_CHK_OK (\S+)
Value XCVR_FLAG_DDM_CAPABLE (\S+)
Value XCVR_FLAG_PMA_PMD_PRESENT (\S+)
Value XCVR_FLAG_PCS_PRESENT (\S+)
Value XCVR_FLAG_PHYXS_PRESENT (\S+)
Value XCVR_FLAG_XCVR_OPERATIONAL (\S+)
Value XCVR_FLAG_PMA_PMD_LOW_PWR_MODE (\S+)
Value XCVR_FLAG_PCS_LOW_PWR_MODE (\S+)
Value XCVR_FLAG_PHYXS_LOW_PWR_MODE (\S+)
Value XCVR_FLAG_LINK_UP (\S+)
Value XCVR_FLAG_TYPE_FIBER (\S+)
Value XCVR_FLAG_SPROM_READ_OK (\S+)
Value PORT_STATE (\S+)
Value XCVR_INSERT_DEBOUNCE_TIMER ([\S\s]+)
Value XCVR_LINK_DEBOUNCE_TIMER ([\S\s]+)
Value XCVR_LINKUP_DEBOUNCE_TIMER ([\S\s]+)
Value XCVR_PRESENCE_DETECT_TIMER ([\S\s]+)
Value TX_ENABLE_SIGNAL (\S+)
Value DEBOUNCE_TIMEOUT ([\S\s]+)
Value LINKUP_DEBOUNCE_TIMEOUT ([\S\s]+)
Value LINK_UP ([\S\s]+)
Value LINK_DN_DEBOUNCE_START ([\S\s]+)
Value LINK_DEBOUNCE_END ([\S\s]+)
Value BIT_ERROR_RATE (\S+)
Value BIT_ERROR_RATE_SINCE_LINKUP (\S+)
Value ERROR_BLOCKS (\S+)
Value ERROR_BLOCKS_SINCE_LINKUP (\S+)
Value LINK_UP_HEX (\S+)
Value LINK_UP_DEC (\S+)
Value LINK_DOWN_HEX (\S+)
Value LINK_DOWN_DEC (\S+)
Value LINK_DEBOUNCED_WITH_LINK_UP_HEX (\S+)
Value LINK_DEBOUNCED_WITH_LINK_UP_DEC (\S+)
Value LINK_DEBOUNCED_WITH_LINK_UP_SINCE_LAST_EN_HEX (\S+)
Value LINK_DEBOUNCED_WITH_LINK_UP_SINCE_LAST_EN_DEC (\S+)


Start
  ^\s*Bigsur\s*port\s*${PORT_NAME}\s*card-config\s*info:\s*$$
  ^\s*if_index\s*:\s*${INTERFACE_INDEX}\s*$$
  ^\s*instance\s*index\s*:\s*${INSTANCE_INDEX}\s*$$
  ^\s*instance\s*slot\s*:\s*${INSTANCE_SLOT}\s*$$
  ^\s*instance\s*asic\s*:\s*${INSTANCE_ASIC}\s*$$
  ^\s*asic_eport\s*:\s*${ASIC_EPORT}\s*$$
  ^\s*logical_port\s*:\s*${LOGICAL_PORT}\s*$$
  ^\s*front_port\s*:\s*${FRONT_PORT}\s*$$
  ^\s*state\s*:\s*${STATE}\s*$$
  ^\s*bigsur\s*eport\s*:\s*${BIGSUR_EPORT}\s*$$
  ^\s*port_type\s*:\s*${PORT}\(${PORT_DEC}\)\s*$$
  ^\s*type\s*:\s*${TYPE}\s*$$
  ^\s*fw_instance\s*:\s*${FIREWALL_INSTANCE}\s*$$
  ^\s*speed\s*:\s*${SPEED}\s*$$
  ^\s*Bigsur\s*port\s*\S+\s*port-client-config\s*info:\s*$$ -> PortClientConfig
  ^\s*link_state\s*:\s*${LINK_STATE}\(${LINK_STATE_DEC}\)\s*$$
  ^\s*oper\s*speed\s*:\s*${OPER_SPEED}\s*$$
  ^\s*oper\s*duplexity\s*:\s*${OPER_DUPLEXITY}\(${OPER_DUPLEXITY_DEC}\)\s*$$
  ^\s*oper\s*port_flow_ctrl\s*:\s*${OPER_PORT_FLOW_CTRL}\s*$$
  ^\s*up\s*count\s*:\s*${UP_COUNT}\s*$$
  ^\s*down\s*count\s*:\s*${DOWN_COUNT}\s*$$
  ^\s*Bigsur\s*port\s*\S+\s*miscellaneous\s*info:\s*$$
  ^\s*slot\s*:\s*${SLOT}\s*$$
  ^\s*mm_inst\s*:\s*${MM_INST}\s*$$
  ^\s*mm_mode\s*:\s*${MM_MODE}\s*$$
  ^\s*diag\s*sts\s*:\s*${DIAG_STS}\s*$$
  ^\s*mac\s*status\s*polling\s*:\s*${MAC_STATUS_POLLING}\s*$$
  ^\s*Clause\s*\d+\s*AN:\s*$$
  ^\s*AN\s*state\s*:\s*${AN_STATE_DEC}\s*\(${AN_STATE}\)\s*$$
  ^\s*AN\s*timer\s*:\s*${AN_TIMER}\s*$$
  ^\s*Link\s*Fail\s*Inhibit\s*timer\s*:\s*${LINK_FAIL_INHIBIT_TIMER}\s*$$
  ^\s*AN\s*disabled\s*:\s*${AN_DISABLED}\s*$$
  ^\s*MM\s*Port\s*Bring\s*Up\s*state-machine:\s*$$
  ^\s*MM\s*port\s*state\s*:\s*${MM_PORT_STATE_DEC}\s*\(${MM_PORT_STATE}\)\s*$$
  ^\s*MM\s*port\s*timer\s*:\s*${MM_PORT_TIMER}\s*$$
  ^\s*Remote\s*fault\s*disable\s*:\s*${REMOTE_FAULT_DISABLE}\s*$$
  ^\s*Adaptive\s*tuning\s*disable\s*:\s*${ADAPTIVE_TUNING_DISABLE}\s*$$
  ^\s*Link\s*recovery\s*disable\s*:\s*${LINK_RECOVERY_DISABLE}\s*$$
  ^\s*Link\s*bring\s*down\s*disable\s*:\s*${LINK_BRING_DOWN_DISABLE}\s*$$
  ^\s*DFE\s*tuning\s*disable\s*:\s*${DFE_TUNING_DISABLE}\s*$$
  ^\s*Force\s*QSFP\s*removal\s*:\s*${FORCE_QSFP_REMOVAL}\s*$$
  ^\s*Force\s*QSFP\s*insertion\s*:\s*${FORCE_QSFP_INSERTION}\s*$$
  ^\s*Force\s*QSFP\s*fiber\s*type\s*:\s*${FORCE_QSFP_FIBER_TYPE}\s*$$
  ^\s*Force\s*QSFP\s*copper\s*type\s*:\s*${FORCE_QSFP_COPPER_TYPE}\s*$$
  ^\s*SBus\s*addresses:\s*$$
  ^\s*Serdes\s*:\s*${SERDES_LANES}\s*$$
  ^\s*Serdes\s*lanes\s*\d+-\d+\s*:\s*${SERDES_LANES}\s*$$
  ^\s*DFE\s*Tune\s*result\s*:\s*$$ -> Lane0
  ^\s*FI\s*block\s*lock\s*0h:\s*${FI_BLOCK_LOCK_0H}\s*$$
  ^\s*FI\s*block\s*lock\s*1h:\s*${FI_BLOCK_LOCK_1H}\s*$$
  ^\s*FI\s*block\s*lock\s*2h:\s*${FI_BLOCK_LOCK_2H}\s*$$
  ^\s*FI\s*block\s*lock\s*3h:\s*${FI_BLOCK_LOCK_3H}\s*$$
  ^\s*MAC\s*statistics:\s*$$
  ^\s*MAC\s*Statistic\s*\|\s*Value\s*$$
  ^\s*[-+]+\s*$$
  ^\s*TX_PKT_SIZE_LT_64\s*\|\s*${TX_PKT_SIZE_LT_64}\s*$$
  ^\s*TX_PKT_SIZE_IS_64\s*\|\s*${TX_PKT_SIZE_IS_64}\s*$$
  ^\s*TX_PKT_SIZE_IS_65_TO_127\s*\|\s*${TX_PKT_SIZE_IS_65_TO_127}\s*$$
  ^\s*TX_PKT_SIZE_IS_128_TO_255\s*\|\s*${TX_PKT_SIZE_IS_128_TO_255}\s*$$
  ^\s*TX_PKT_SIZE_IS_256_TO_511\s*\|\s*${TX_PKT_SIZE_IS_256_TO_511}\s*$$
  ^\s*TX_PKT_SIZE_IS_512_TO_1023\s*\|\s*${TX_PKT_SIZE_IS_512_TO_1023}\s*$$
  ^\s*TX_PKT_SIZE_IS_1024_TO_1518\s*\|\s*${TX_PKT_SIZE_IS_1024_TO_1518}\s*$$
  ^\s*TX_PKT_SIZE_IS_1519_TO_2047\s*\|\s*${TX_PKT_SIZE_IS_1519_TO_2047}\s*$$
  ^\s*TX_PKT_SIZE_IS_2048_TO_4095\s*\|\s*${TX_PKT_SIZE_IS_2048_TO_4095}\s*$$
  ^\s*TX_PKT_SIZE_IS_4095_TO_8191\s*\|\s*${TX_PKT_SIZE_IS_4095_TO_8191}\s*$$
  ^\s*TX_PKT_SIZE_IS_8192_TO_9216\s*\|\s*${TX_PKT_SIZE_IS_8192_TO_9216}\s*$$
  ^\s*TX_PKT_SIZE_GT_9216\s*\|\s*${TX_PKT_SIZE_GT_9216}\s*$$
  ^\s*TX_PKT_TOTAL\s*\|\s*${TX_PKT_TOTAL}\s*$$
  ^\s*TX_PKT_OCTETS\s*\|\s*${TX_PKT_OCTETS}\s*$$
  ^\s*TX_PKT_GOOD\s*\|\s*${TX_PKT_GOOD}\s*$$
  ^\s*TX_PKT_UCAST\s*\|\s*${TX_PKT_UCAST}\s*$$
  ^\s*TX_PKT_MCAST\s*\|\s*${TX_PKT_MCAST}\s*$$
  ^\s*TX_PKT_BCAST\s*\|\s*${TX_PKT_BCAST}\s*$$
  ^\s*TX_PKT_VLAN\s*\|\s*${TX_PKT_VLAN}\s*$$
  ^\s*TX_PKT_802.3x_PAUSE\s*\|\s*${TX_PKT_802_3X_PAUSE}\s*$$
  ^\s*TX_PKT_PER_PRIORITY_PAUSE\s*\|\s*${TX_PKT_PER_PRIORITY_PAUSE}\s*$$
  ^\s*TX_PKT_FRAME_ERROR\s*\|\s*${TX_PKT_FRAME_ERROR}\s*$$
  ^\s*RX_PKT_SIZE_LT_64\s*\|\s*${RX_PKT_SIZE_LT_64}\s*$$
  ^\s*RX_PKT_SIZE_IS_64\s*\|\s*${RX_PKT_SIZE_IS_64}\s*$$
  ^\s*RX_PKT_SIZE_IS_65_TO_127\s*\|\s*${RX_PKT_SIZE_IS_65_TO_127}\s*$$
  ^\s*RX_PKT_SIZE_IS_128_TO_255\s*\|\s*${RX_PKT_SIZE_IS_128_TO_255}\s*$$
  ^\s*RX_PKT_SIZE_IS_256_TO_511\s*\|\s*${RX_PKT_SIZE_IS_256_TO_511}\s*$$
  ^\s*RX_PKT_SIZE_IS_512_TO_1023\s*\|\s*${RX_PKT_SIZE_IS_512_TO_1023}\s*$$
  ^\s*RX_PKT_SIZE_IS_1024_TO_1518\s*\|\s*${RX_PKT_SIZE_IS_1024_TO_1518}\s*$$
  ^\s*RX_PKT_SIZE_IS_1519_TO_2047\s*\|\s*${RX_PKT_SIZE_IS_1519_TO_2047}\s*$$
  ^\s*RX_PKT_SIZE_IS_2048_TO_4095\s*\|\s*${RX_PKT_SIZE_IS_2048_TO_4095}\s*$$
  ^\s*RX_PKT_SIZE_IS_4095_TO_8191\s*\|\s*${RX_PKT_SIZE_IS_4095_TO_8191}\s*$$
  ^\s*RX_PKT_SIZE_IS_8192_TO_9216\s*\|\s*${RX_PKT_SIZE_IS_8192_TO_9216}\s*$$
  ^\s*RX_PKT_SIZE_GT_9216\s*\|\s*${RX_PKT_SIZE_GT_9216}\s*$$
  ^\s*RX_PKT_TOTAL\s*\|\s*${RX_PKT_TOTAL}\s*$$
  ^\s*RX_PKT_OCTETS\s*\|\s*${RX_PKT_OCTETS}\s*$$
  ^\s*RX_PKT_GOOD\s*\|\s*${RX_PKT_GOOD}\s*$$
  ^\s*RX_PKT_UCAST\s*\|\s*${RX_PKT_UCAST}\s*$$
  ^\s*RX_PKT_MCAST\s*\|\s*${RX_PKT_MCAST}\s*$$
  ^\s*RX_PKT_BCAST\s*\|\s*${RX_PKT_BCAST}\s*$$
  ^\s*RX_PKT_VLAN\s*\|\s*${RX_PKT_VLAN}\s*$$
  ^\s*RX_PKT_OVERSIZE\s*\|\s*${RX_PKT_OVERSIZE}\s*$$
  ^\s*RX_PKT_TOOLONG\s*\|\s*${RX_PKT_TOOLONG}\s*$$
  ^\s*RX_PKT_DISCARD\s*\|\s*${RX_PKT_DISCARD}\s*$$
  ^\s*RX_PKT_UNDERSIZE\s*\|\s*${RX_PKT_UNDERSIZE}\s*$$
  ^\s*RX_PKT_FRAGMENT\s*\|\s*${RX_PKT_FRAGMENT}\s*$$
  ^\s*RX_PKT_CRC_NOT_STOMPED\s*\|\s*${RX_PKT_CRC_NOT_STOMPED}\s*$$
  ^\s*RX_PKT_CRC_STOMPED\s*\|\s*${RX_PKT_CRC_STOMPED}\s*$$
  ^\s*RX_PKT_IN_RANGE_ERR\s*\|\s*${RX_PKT_IN_RANGE_ERR}\s*$$
  ^\s*RX_PKT_JABBER\s*\|\s*${RX_PKT_JABBER}\s*$$
  ^\s*RX_PKT_802.3x_PAUSE\s*\|\s*${RX_PKT_802_3X_PAUSE}\s*$$
  ^\s*RX_PKT_PER_PRIORITY_PAUSE\s*\|\s*${RX_PKT_PER_PRIORITY_PAUSE}\s*$$
  ^\s*TX_PKT_GOOD_OCTETS\s*\|\s*${TX_PKT_GOOD_OCTETS}\s*$$
  ^\s*RX_PKT_GOOD_OCTETS\s*\|\s*${RX_PKT_GOOD_OCTETS}\s*$$
  ^\s*XCVR\s*port\s*info:\s*$$
  ^\s*Card\s*Config\s*info\s*$$
  ^\s*sfp_ctrl_bigsur\s*:\s*${SFP_CTRL_BIGSUR}\s*$$
  ^\s*sfp_ctrl_port\s*:\s*${SFP_CTRL_PORT}\s*$$
  ^\s*phy_bigsur\s*:\s*${PHY_BIGSUR}\s*$$
  ^\s*phy_mdio_addr\s*:\s*${PHY_MDIO_ADDR}\s*$$
  ^\s*phy_mdio_master\s*:\s*${PHY_MDIO_MASTER}\s*$$
  ^\s*phy_reset_bigsur\s*:\s*${PHY_RESET_BIGSUR}\s*$$
  ^\s*phy_reset_spio\s*:\s*${PHY_RESET_SPIO}\s*$$
  ^\s*phy_prom_shared\s*:\s*${PHY_PROM_SHARED}\s*$$
  ^\s*Serdes\s*txdata\[0\]\s*:\s*${SERDES_TXDATA_0}\s*$$
  ^\s*Serdes\s*txdata\[1\]\s*:\s*${SERDES_TXDATA_1}\s*$$
  ^\s*Serdes\s*txdata\[2\]\s*:\s*${SERDES_TXDATA_2}\s*$$
  ^\s*Serdes\s*txdata\[3\]\s*:\s*${SERDES_TXDATA_3}\s*$$
  ^\s*phy_pmd_tx\[0\]\s*:\s*${PHY_PMD_TX_0}\s*$$
  ^\s*phy_pmd_tx\[1\]\s*:\s*${PHY_PMD_TX_1}\s*$$
  ^\s*phy_pmd_tx\[2\]\s*:\s*${PHY_PMD_TX_2}\s*$$
  ^\s*Dev\s*address\s*info\s*$$
  ^\s*pma_pmd_devaddr\s*:\s*${PMA_PMD_DEVADDR}\s*$$
  ^\s*pcs_devaddr\s*:\s*${PCS_DEVADDR}\s*$$
  ^\s*phyxs_devaddr\s*:\s*${PHYXS_DEVADDR}\s*$$
  ^\s*xcvr_devaddr\s*:\s*${XCVR_DEVADDR}\s*$$
  ^\s*Misc\s*info\s*$$
  ^\s*phy_present\s*:\s*${PHY_PRESENT}\s*$$
  ^\s*phy_type\s*:\s*${PHY_TYPE}\s*$$
  ^\s*sfp_acc_type\s*:\s*${SFP_ACC_TYPE}\s*$$
  ^\s*dfe_timer\s*:\s*${DFE_TIMER}\s*$$
  ^\s*dfe_timer_running\s*:\s*${DFE_TIMER_RUNNING}\s*$$
  ^\s*init_wait_count\s*:\s*${INIT_WAIT_COUNT}\s*$$
  ^\s*do_tx_disable_clr\s*:\s*${DO_TX_DISABLE_CLR}\s*$$
  ^\s*tx_disable_clr_cnt\s*:\s*${TX_DISABLE_CLR_CNT}\s*$$
  ^\s*qsa_status\s*:\s*${QSA_STATUS}\s*$$
  ^\s*Connector\s*type\s*:\s*${CONNECTOR_TYPE}\s*$$
  ^\s*usb_bus_num\s*:\s*${USB_BUS_NUM}\s*$$
  ^\s*usb_dev_num\s*:\s*${USB_DEV_NUM}\s*$$
  ^\s*usb_priv\s*:\s*${USB_PRIV}\s*$$
  ^\s*Reset\s*done\s*:\s*${RESET_DONE}\s*$$
  ^\s*Xcvr\s*flags\s*$$
  ^\s*XCVR_FLAG_PRESENT\s*:\s*${XCVR_FLAG_PRESENT}\s*$$
  ^\s*XCVR_FLAG_XCVR_NOT_SFP\s*:\s*${XCVR_FLAG_XCVR_NOT_SFP}\s*$$
  ^\s*XCVR_FLAG_10G_CAPABLE\s*:\s*${XCVR_FLAG_10G_CAPABLE}\s*$$
  ^\s*XCVR_FLAG_CKSUM_OK\s*:\s*${XCVR_FLAG_CKSUM_OK}\s*$$
  ^\s*XCVR_FLAG_SECURITY_CRC_OK\s*:\s*${XCVR_FLAG_SECURITY_CRC_OK}\s*$$
  ^\s*XCVR_FLAG_CISCO_CHK_OK\s*:\s*${XCVR_FLAG_CISCO_CHK_OK}\s*$$
  ^\s*XCVR_FLAG_DDM_CAPABLE\s*:\s*${XCVR_FLAG_DDM_CAPABLE}\s*$$
  ^\s*XCVR_FLAG_PMA_PMD_PRESENT\s*:\s*${XCVR_FLAG_PMA_PMD_PRESENT}\s*$$
  ^\s*XCVR_FLAG_PCS_PRESENT\s*:\s*${XCVR_FLAG_PCS_PRESENT}\s*$$
  ^\s*XCVR_FLAG_PHYXS_PRESENT\s*:\s*${XCVR_FLAG_PHYXS_PRESENT}\s*$$
  ^\s*XCVR_FLAG_XCVR_OPERATIONAL\s*:\s*${XCVR_FLAG_XCVR_OPERATIONAL}\s*$$
  ^\s*XCVR_FLAG_PMA_PMD_LOW_PWR_MODE\s*:\s*${XCVR_FLAG_PMA_PMD_LOW_PWR_MODE}\s*$$
  ^\s*XCVR_FLAG_PCS_LOW_PWR_MODE\s*:\s*${XCVR_FLAG_PCS_LOW_PWR_MODE}\s*$$
  ^\s*XCVR_FLAG_PHYXS_LOW_PWR_MODE\s*:\s*${XCVR_FLAG_PHYXS_LOW_PWR_MODE}\s*$$
  ^\s*XCVR_FLAG_LINK_UP\s*:\s*${XCVR_FLAG_LINK_UP}\s*$$
  ^\s*XCVR_FLAG_TYPE_FIBER\s*:\s*${XCVR_FLAG_TYPE_FIBER}\s*$$
  ^\s*XCVR_FLAG_SPROM_READ_OK\s*:\s*${XCVR_FLAG_SPROM_READ_OK}\s*$$
  ^\s*Link\s*states:[\S\s]*\s*$$
  ^\s*Port\s\S+:\s*$$
  ^\s*State\s*:\s*${PORT_STATE}\s*$$
  ^\s*XCVR\s*insert\s*debounce\s*timer\s*${XCVR_INSERT_DEBOUNCE_TIMER}\s*$$
  ^\s*XCVR\s*link\s*debounce\s*timer\s*${XCVR_LINK_DEBOUNCE_TIMER}\s*$$
  ^\s*XCVR\s*linkup\s*debounce\s*timer\s*${XCVR_LINKUP_DEBOUNCE_TIMER}\s*$$
  ^\s*XCVR\s*presence\s*detect\s*timer\s*${XCVR_PRESENCE_DETECT_TIMER}\s*$$
  ^\s*TX\s*enable\s*signal\s*is\s*${TX_ENABLE_SIGNAL}\s*$$
  ^\s*Debounce\s*timeout\s*:\s*${DEBOUNCE_TIMEOUT}\s*$$
  ^\s*Linkup\s*Debounce\s*timeout\s*:\s*${LINKUP_DEBOUNCE_TIMEOUT}\s*$$
  ^\s*Link\s*up\s*:\s*${LINK_UP}\s*$$
  ^\s*Link\s*dn\s*debounce\s*start\s*:\s*${LINK_DN_DEBOUNCE_START}\s*$$
  ^\s*Link\s*debounce\s*end\s*:\s*${LINK_DEBOUNCE_END}\s*$$
  ^\s*Counters:\s*$$
  ^\s*Interrupt\s*cntrs:\s*$$
  ^\s*Bit\s*error\s*cntrs:\s*$$
  ^\s*Bit\s*Error\s*Rate:\s*${BIT_ERROR_RATE}\s*Bit\s*Error\s*Rate\(since\s*linkup\):\s*${BIT_ERROR_RATE_SINCE_LINKUP}\s*$$
  ^\s*Error\s*blocks\s*:\s*${ERROR_BLOCKS}\s*Error\s*blocks\(since\s*linkup\)\s*:\s*${ERROR_BLOCKS_SINCE_LINKUP}\s*$$
  ^\s*Link\s*cntrs:\s*$$ -> LinkCntrs
  ^\s*Done.\s*$$ -> Record
  ^\s*$$
  ^. -> Error

LinkCntrs
  ^\s*Link\s*up:\s*${LINK_UP_HEX}\s*\(${LINK_UP_DEC}\)\s*$$
  ^\s*Link\s*dn:\s*${LINK_DOWN_HEX}\s*\(${LINK_DOWN_DEC}\)\s*$$
  ^\s*Link\s*debounced\s*with\s*link\s*up:\s*${LINK_DEBOUNCED_WITH_LINK_UP_HEX}\s*\(${LINK_DEBOUNCED_WITH_LINK_UP_DEC}\)\s*$$
  ^\s*Link\s*debounced\s*with\s*link\s*up\s*since\s*last\s*enable\s*:\s*${LINK_DEBOUNCED_WITH_LINK_UP_SINCE_LAST_EN_HEX}\s*\(${LINK_DEBOUNCED_WITH_LINK_UP_SINCE_LAST_EN_DEC}\)\s*$$ -> Start

PortClientConfig
  ^\s*admin_state\s*:\s*${ADMIN_STATE}\(${ADMIN_STATE_DEC}\)\s*$$
  ^\s*mtu\s*:\s*${MTU}\s*$$
  ^\s*duplexity\s*:\s*${DUPLEXITY}\(${DUPLEXITY_DEC}\)\s*$$
  ^\s*auto_neg_state\s*:\s*\(${AUTO_NEG_STATE}\)${AUTO_NEG_STATE_DEC}\s*$$
  ^\s*Fast\s*1G\s*AN\s*timer\s*is\s*:\s*${FAST_1G_AN_TIMER}\s*$$
  ^\s*speed\s*:\s*${PORT_SPEED}\s*$$
  ^\s*debounce\s*time\s*:\s*${DEBOUNCE_TIME}\s*$$
  ^\s*linkup_debounce\s*time\s*:\s*${LINKUP_DEBOUNCE_TIME}\s*$$
  ^\s*sat_fabric_mode\s*:\s*${SAT_FABRIC_MODE}\(${SAT_FABRIC_MODE_HEX}\)\s*$$
  ^\s*port_flow_ctrl\s*:\s*${PORT_FLOW_CTRL}\s*$$
  ^\s*class_flow_ctrl7-0\s*:\s*${CLASS_FLOW_CTRL7_0}\s*$$
  ^\s*class_flow_ctrl\s*msk\s*:\s*rx\s*${CLASS_FLOW_CTRL_MSK_RX},\s*tx\s*${CLASS_FLOW_CTRL_MSK_TX}\s*$$
  ^\s*Bigsur\s*port\s*\S+\s*port-oper\s*info:\s*$$ -> Start
  ^\s*$$
  ^. -> Error


Lane0
  ^\s*Lane\s*\d+:\s*$$
  ^\s*<\d+\s*\S+>:DFE\s*A:\s*${DFE_TUNE_LANE_0_DFE_A}\s*$$
  ^\s*<\d+\s*\S+>:DFE\s*B:\s*${DFE_TUNE_LANE_0_DFE_B}\s*$$
  ^\s*<\d+\s*\S+>\s*DFE\s*Tune\s*\[\S+\]:\s*${DFE_TUNE_LANE_0_DFE_TUNE}\s*$$ -> Lane1
  ^\s*$$ -> Start
  ^. -> Error


Lane1
  ^\s*Lane\s*\d+:\s*$$
  ^\s*<\d+\s*\S+>:DFE\s*A:\s*${DFE_TUNE_LANE_1_DFE_A}\s*$$
  ^\s*<\d+\s*\S+>:DFE\s*B:\s*${DFE_TUNE_LANE_1_DFE_B}\s*$$
  ^\s*<\d+\s*\S+>\s*DFE\s*Tune\s*\[\S+\]:\s*${DFE_TUNE_LANE_1_DFE_TUNE}\s*$$ -> Lane2
  ^\s*$$ -> Start
  ^. -> Error


Lane2
  ^\s*Lane\s*\d+:\s*$$
  ^\s*<\d+\s*\S+>:DFE\s*A:\s*${DFE_TUNE_LANE_2_DFE_A}\s*$$
  ^\s*<\d+\s*\S+>:DFE\s*B:\s*${DFE_TUNE_LANE_2_DFE_B}\s*$$
  ^\s*<\d+\s*\S+>\s*DFE\s*Tune\s*\[\S+\]:\s*${DFE_TUNE_LANE_2_DFE_TUNE}\s*$$ -> Lane3
  ^\s*$$ -> Start
  ^. -> Error


Lane3
  ^\s*Lane\s*\d+:\s*$$
  ^\s*<\d+\s*\S+>:DFE\s*A:\s*${DFE_TUNE_LANE_3_DFE_A}\s*$$
  ^\s*<\d+\s*\S+>:DFE\s*B:\s*${DFE_TUNE_LANE_3_DFE_B}\s*$$
  ^\s*<\d+\s*\S+>\s*DFE\s*Tune\s*\[\S+\]:\s*${DFE_TUNE_LANE_3_DFE_TUNE}\s*$$ -> Start
  ^\s*$$ -> Start
  ^. -> Error


EOF`