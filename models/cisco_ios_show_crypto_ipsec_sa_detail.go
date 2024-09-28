package models

type CiscoIosShowCryptoIpsecSaDetail struct {
	Interface	string	`json:"INTERFACE"`
	Crypto_map_tag	string	`json:"CRYPTO_MAP_TAG"`
	Local_addr	string	`json:"LOCAL_ADDR"`
	Protected_vrf	string	`json:"PROTECTED_VRF"`
	Local_ident_addr	string	`json:"LOCAL_IDENT_ADDR"`
	Local_ident_mask	string	`json:"LOCAL_IDENT_MASK"`
	Local_ident_prot	string	`json:"LOCAL_IDENT_PROT"`
	Local_ident_port	string	`json:"LOCAL_IDENT_PORT"`
	Remote_ident_addr	string	`json:"REMOTE_IDENT_ADDR"`
	Remote_ident_mask	string	`json:"REMOTE_IDENT_MASK"`
	Remote_ident_prot	string	`json:"REMOTE_IDENT_PROT"`
	Remote_ident_port	string	`json:"REMOTE_IDENT_PORT"`
	Current_peer	string	`json:"CURRENT_PEER"`
	Port	string	`json:"PORT"`
	Action	string	`json:"ACTION"`
	Flags	string	`json:"FLAGS"`
	Pkts_encaps	string	`json:"PKTS_ENCAPS"`
	Pkts_encrypt	string	`json:"PKTS_ENCRYPT"`
	Pkts_digest	string	`json:"PKTS_DIGEST"`
	Pkts_decaps	string	`json:"PKTS_DECAPS"`
	Pkts_decrypt	string	`json:"PKTS_DECRYPT"`
	Pkts_verify	string	`json:"PKTS_VERIFY"`
	Pkts_compressed	string	`json:"PKTS_COMPRESSED"`
	Pkts_decompressed	string	`json:"PKTS_DECOMPRESSED"`
	Pkts_not_compressed	string	`json:"PKTS_NOT_COMPRESSED"`
	Pkts_not_decompressed	string	`json:"PKTS_NOT_DECOMPRESSED"`
	Pkts_compress_failed	string	`json:"PKTS_COMPRESS_FAILED"`
	Pkts_decompress_failed	string	`json:"PKTS_DECOMPRESS_FAILED"`
	Pkts_no_sa	string	`json:"PKTS_NO_SA"`
	Pkts_invalid_sa	string	`json:"PKTS_INVALID_SA"`
	Pkts_encaps_failed	string	`json:"PKTS_ENCAPS_FAILED"`
	Pkts_decaps_failed	string	`json:"PKTS_DECAPS_FAILED"`
	Pkts_invalid_prot	string	`json:"PKTS_INVALID_PROT"`
	Pkts_verify_failed	string	`json:"PKTS_VERIFY_FAILED"`
	Pkts_invalid_identity	string	`json:"PKTS_INVALID_IDENTITY"`
	Pkts_invalid_len	string	`json:"PKTS_INVALID_LEN"`
	Pkts_replay_rollover_send	string	`json:"PKTS_REPLAY_ROLLOVER_SEND"`
	Pkts_replay_rollover_recv	string	`json:"PKTS_REPLAY_ROLLOVER_RECV"`
	Pkts_replay_failed	string	`json:"PKTS_REPLAY_FAILED"`
	Pkts_tagged	string	`json:"PKTS_TAGGED"`
	Pkts_untagged	string	`json:"PKTS_UNTAGGED"`
	Pkts_not_tagged	string	`json:"PKTS_NOT_TAGGED"`
	Pkts_not_untagged	string	`json:"PKTS_NOT_UNTAGGED"`
	Pkts_internal_err_send	string	`json:"PKTS_INTERNAL_ERR_SEND"`
	Pkts_internal_err_recv	string	`json:"PKTS_INTERNAL_ERR_RECV"`
	Local_crypto_endpt	string	`json:"LOCAL_CRYPTO_ENDPT"`
	Remote_crypto_endpt	string	`json:"REMOTE_CRYPTO_ENDPT"`
	Plaintext_mtu	string	`json:"PLAINTEXT_MTU"`
	Path_mtu	string	`json:"PATH_MTU"`
	Ip_mtu	string	`json:"IP_MTU"`
	Ip_mtu_idb	string	`json:"IP_MTU_IDB"`
	Current_outbound_spi_hex	string	`json:"CURRENT_OUTBOUND_SPI_HEX"`
	Current_outbound_spi_dec	string	`json:"CURRENT_OUTBOUND_SPI_DEC"`
	Pfs	string	`json:"PFS"`
	Dh_group	string	`json:"DH_GROUP"`
	Sa_orientation	string	`json:"SA_ORIENTATION"`
	Sa_type	string	`json:"SA_TYPE"`
	Sa_spi_hex	string	`json:"SA_SPI_HEX"`
	Sa_spi_dec	string	`json:"SA_SPI_DEC"`
	Sa_transform	string	`json:"SA_TRANSFORM"`
	Sa_in_use_settings	string	`json:"SA_IN_USE_SETTINGS"`
	Sa_conn_id	string	`json:"SA_CONN_ID"`
	Sa_flow_id	string	`json:"SA_FLOW_ID"`
	Sa_sibling_flags	string	`json:"SA_SIBLING_FLAGS"`
	Sa_crypto_map	string	`json:"SA_CRYPTO_MAP"`
	Sa_lifetime_kbytes	string	`json:"SA_LIFETIME_KBYTES"`
	Sa_lifetime_sec	string	`json:"SA_LIFETIME_SEC"`
	Sa_iv_size	string	`json:"SA_IV_SIZE"`
	Sa_replay_detection_support	string	`json:"SA_REPLAY_DETECTION_SUPPORT"`
	Sa_replay_window_size	string	`json:"SA_REPLAY_WINDOW_SIZE"`
	Sa_status	string	`json:"SA_STATUS"`
}