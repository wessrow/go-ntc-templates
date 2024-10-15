package cisco_ios

type ShowCryptoIpsecSaDetail struct {
	Pkts_compressed             string `json:"PKTS_COMPRESSED"`
	Pkts_not_decompressed       string `json:"PKTS_NOT_DECOMPRESSED"`
	Pkts_tagged                 string `json:"PKTS_TAGGED"`
	Ip_mtu_idb                  string `json:"IP_MTU_IDB"`
	Pkts_replay_rollover_send   string `json:"PKTS_REPLAY_ROLLOVER_SEND"`
	Current_outbound_spi_hex    string `json:"CURRENT_OUTBOUND_SPI_HEX"`
	Sa_lifetime_kbytes          string `json:"SA_LIFETIME_KBYTES"`
	Sa_iv_size                  string `json:"SA_IV_SIZE"`
	Pkts_compress_failed        string `json:"PKTS_COMPRESS_FAILED"`
	Pkts_invalid_len            string `json:"PKTS_INVALID_LEN"`
	Sa_in_use_settings          string `json:"SA_IN_USE_SETTINGS"`
	Sa_replay_window_size       string `json:"SA_REPLAY_WINDOW_SIZE"`
	Local_ident_port            string `json:"LOCAL_IDENT_PORT"`
	Pkts_not_compressed         string `json:"PKTS_NOT_COMPRESSED"`
	Pkts_not_untagged           string `json:"PKTS_NOT_UNTAGGED"`
	Sa_replay_detection_support string `json:"SA_REPLAY_DETECTION_SUPPORT"`
	Sa_flow_id                  string `json:"SA_FLOW_ID"`
	Local_addr                  string `json:"LOCAL_ADDR"`
	Local_ident_mask            string `json:"LOCAL_IDENT_MASK"`
	Pkts_decaps_failed          string `json:"PKTS_DECAPS_FAILED"`
	Pkts_internal_err_send      string `json:"PKTS_INTERNAL_ERR_SEND"`
	Dh_group                    string `json:"DH_GROUP"`
	Pkts_decompress_failed      string `json:"PKTS_DECOMPRESS_FAILED"`
	Pkts_verify_failed          string `json:"PKTS_VERIFY_FAILED"`
	Pkts_replay_failed          string `json:"PKTS_REPLAY_FAILED"`
	Sa_spi_hex                  string `json:"SA_SPI_HEX"`
	Pkts_encrypt                string `json:"PKTS_ENCRYPT"`
	Pkts_invalid_prot           string `json:"PKTS_INVALID_PROT"`
	Sa_sibling_flags            string `json:"SA_SIBLING_FLAGS"`
	Sa_orientation              string `json:"SA_ORIENTATION"`
	Sa_spi_dec                  string `json:"SA_SPI_DEC"`
	Interface                   string `json:"INTERFACE"`
	Port                        string `json:"PORT"`
	Pkts_decaps                 string `json:"PKTS_DECAPS"`
	Pkts_encaps_failed          string `json:"PKTS_ENCAPS_FAILED"`
	Current_outbound_spi_dec    string `json:"CURRENT_OUTBOUND_SPI_DEC"`
	Ip_mtu                      string `json:"IP_MTU"`
	Protected_vrf               string `json:"PROTECTED_VRF"`
	Local_ident_addr            string `json:"LOCAL_IDENT_ADDR"`
	Flags                       string `json:"FLAGS"`
	Pkts_invalid_identity       string `json:"PKTS_INVALID_IDENTITY"`
	Pkts_untagged               string `json:"PKTS_UNTAGGED"`
	Current_peer                string `json:"CURRENT_PEER"`
	Pkts_decompressed           string `json:"PKTS_DECOMPRESSED"`
	Pkts_invalid_sa             string `json:"PKTS_INVALID_SA"`
	Pkts_replay_rollover_recv   string `json:"PKTS_REPLAY_ROLLOVER_RECV"`
	Sa_crypto_map               string `json:"SA_CRYPTO_MAP"`
	Sa_transform                string `json:"SA_TRANSFORM"`
	Crypto_map_tag              string `json:"CRYPTO_MAP_TAG"`
	Local_ident_prot            string `json:"LOCAL_IDENT_PROT"`
	Remote_ident_port           string `json:"REMOTE_IDENT_PORT"`
	Pkts_digest                 string `json:"PKTS_DIGEST"`
	Local_crypto_endpt          string `json:"LOCAL_CRYPTO_ENDPT"`
	Pkts_decrypt                string `json:"PKTS_DECRYPT"`
	Pkts_verify                 string `json:"PKTS_VERIFY"`
	Pkts_no_sa                  string `json:"PKTS_NO_SA"`
	Plaintext_mtu               string `json:"PLAINTEXT_MTU"`
	Sa_type                     string `json:"SA_TYPE"`
	Remote_ident_addr           string `json:"REMOTE_IDENT_ADDR"`
	Path_mtu                    string `json:"PATH_MTU"`
	Pfs                         string `json:"PFS"`
	Sa_conn_id                  string `json:"SA_CONN_ID"`
	Sa_lifetime_sec             string `json:"SA_LIFETIME_SEC"`
	Remote_ident_prot           string `json:"REMOTE_IDENT_PROT"`
	Remote_crypto_endpt         string `json:"REMOTE_CRYPTO_ENDPT"`
	Remote_ident_mask           string `json:"REMOTE_IDENT_MASK"`
	Action                      string `json:"ACTION"`
	Sa_status                   string `json:"SA_STATUS"`
	Pkts_encaps                 string `json:"PKTS_ENCAPS"`
	Pkts_not_tagged             string `json:"PKTS_NOT_TAGGED"`
	Pkts_internal_err_recv      string `json:"PKTS_INTERNAL_ERR_RECV"`
}

var ShowCryptoIpsecSaDetail_Template string = `Value Filldown INTERFACE (\S+)
Value Filldown CRYPTO_MAP_TAG ([\S-]+)
Value Filldown LOCAL_ADDR (\S+)
Value Filldown PROTECTED_VRF (\S+)
Value Filldown LOCAL_IDENT_ADDR (\S+)
Value Filldown LOCAL_IDENT_MASK (\S+)
Value Filldown LOCAL_IDENT_PROT (\d+)
Value Filldown LOCAL_IDENT_PORT (\d+)
Value Filldown REMOTE_IDENT_ADDR (\S+)
Value Filldown REMOTE_IDENT_MASK (\S+)
Value Filldown REMOTE_IDENT_PROT (\d+)
Value Filldown REMOTE_IDENT_PORT (\d+)
Value Filldown CURRENT_PEER (\S+)
Value Filldown PORT (\d+)
Value Filldown ACTION (\S+)
Value Filldown FLAGS ([\S\s]*)
Value Filldown PKTS_ENCAPS (\d+)
Value Filldown PKTS_ENCRYPT (\d+)
Value Filldown PKTS_DIGEST (\d+)
Value Filldown PKTS_DECAPS (\d+)
Value Filldown PKTS_DECRYPT (\d+)
Value Filldown PKTS_VERIFY (\d+)
Value Filldown PKTS_COMPRESSED (\d+)
Value Filldown PKTS_DECOMPRESSED (\d+)
Value Filldown PKTS_NOT_COMPRESSED (\d+)
Value Filldown PKTS_NOT_DECOMPRESSED (\d+)
Value Filldown PKTS_COMPRESS_FAILED (\d+)
Value Filldown PKTS_DECOMPRESS_FAILED (\d+)
Value Filldown PKTS_NO_SA (\d+)
Value Filldown PKTS_INVALID_SA (\d+)
Value Filldown PKTS_ENCAPS_FAILED (\d+)
Value Filldown PKTS_DECAPS_FAILED (\d+)
Value Filldown PKTS_INVALID_PROT (\d+)
Value Filldown PKTS_VERIFY_FAILED (\d+)
Value Filldown PKTS_INVALID_IDENTITY (\d+)
Value Filldown PKTS_INVALID_LEN (\d+)
Value Filldown PKTS_REPLAY_ROLLOVER_SEND (\d+)
Value Filldown PKTS_REPLAY_ROLLOVER_RECV (\d+)
Value Filldown PKTS_REPLAY_FAILED (\d+)
Value Filldown PKTS_TAGGED (\d+)
Value Filldown PKTS_UNTAGGED (\d+)
Value Filldown PKTS_NOT_TAGGED (\d+)
Value Filldown PKTS_NOT_UNTAGGED (\d+)
Value Filldown PKTS_INTERNAL_ERR_SEND (\d+)
Value Filldown PKTS_INTERNAL_ERR_RECV (\d+)
Value Filldown LOCAL_CRYPTO_ENDPT (\S+)
Value Filldown REMOTE_CRYPTO_ENDPT (\S+)
Value Filldown PLAINTEXT_MTU (\d+)
Value Filldown PATH_MTU (\d+)
Value Filldown IP_MTU (\d+)
Value Filldown IP_MTU_IDB (\S+)
Value Filldown CURRENT_OUTBOUND_SPI_HEX (\S+)
Value Filldown CURRENT_OUTBOUND_SPI_DEC (\d+)
Value Filldown PFS (\S+)
Value Filldown DH_GROUP (\S+)
Value Filldown SA_ORIENTATION (\S+)
Value Filldown SA_TYPE (\S+)
Value Required SA_SPI_HEX (\S+)
Value Required SA_SPI_DEC (\d+)
Value SA_TRANSFORM ([\S\s]+)
Value SA_IN_USE_SETTINGS ([\S\s]+?)
Value SA_CONN_ID (\d+)
Value SA_FLOW_ID ([\S\s]+)
Value SA_SIBLING_FLAGS ([\d\w]+)
Value SA_CRYPTO_MAP (\S+)
Value SA_LIFETIME_KBYTES (\d+)
Value SA_LIFETIME_SEC (\d+)
Value SA_IV_SIZE (\d+)
Value SA_REPLAY_DETECTION_SUPPORT (\S+)
Value SA_REPLAY_WINDOW_SIZE (\d+)
Value SA_STATUS (\S+)

Start
  #^\s*interface:\s+\S+\s*$$ -> Continue.Record
  ^\s*interface:\s+${INTERFACE}\s*$$
  ^\s*Crypto\s+map\s+tag:\s+${CRYPTO_MAP_TAG},\s+local\s+addr\s+${LOCAL_ADDR}\s*$$
  ^\s*protected\s+vrf:\s+${PROTECTED_VRF}\s*$$
  ^\s*local\s+ident\s+\(addr/mask/prot/port\):\s+\(${LOCAL_IDENT_ADDR}/${LOCAL_IDENT_MASK}/${LOCAL_IDENT_PROT}/${LOCAL_IDENT_PORT}\)\s*$$
  ^\s*remote\s+ident\s+\(addr/mask/prot/port\):\s+\(${REMOTE_IDENT_ADDR}/${REMOTE_IDENT_MASK}/${REMOTE_IDENT_PROT}/${REMOTE_IDENT_PORT}\)\s*$$
  ^\s*current_peer\s+${CURRENT_PEER}\s+port\s+${PORT}\s*$$
  ^\s*${ACTION},\s+flags={${FLAGS}}\s*$$
  ^\s*#pkts\s+encaps:\s+${PKTS_ENCAPS},\s+#pkts\s+encrypt:\s+${PKTS_ENCRYPT},\s+#pkts\s+digest:\s+${PKTS_DIGEST}\s*$$
  ^\s*#pkts\s+decaps:\s+${PKTS_DECAPS},\s+#pkts\s+decrypt:\s+${PKTS_DECRYPT},\s+#pkts\s+verify:\s+${PKTS_VERIFY}\s*$$
  ^\s*#pkts\s+compressed:\s+${PKTS_COMPRESSED},\s+#pkts\s+decompressed:\s+${PKTS_DECOMPRESSED}\s*$$
  ^\s*#pkts\s+not\s+compressed:\s+${PKTS_NOT_COMPRESSED},\s+#pkts\s+compr.\s+failed:\s+${PKTS_COMPRESS_FAILED}\s*$$
  ^\s*#pkts\s+not\s+decompressed:\s+${PKTS_NOT_DECOMPRESSED},\s+#pkts\s+decompress\s+failed:\s+${PKTS_DECOMPRESS_FAILED}\s*$$
  ^\s*#pkts\s+no\s+sa\s+\(send\)\s+${PKTS_NO_SA},\s+#pkts\s+invalid\s+sa\s\(rcv\)\s+${PKTS_INVALID_SA}\s*$$
  ^\s*#pkts\sencaps\s+failed\s\(send\)\s+${PKTS_ENCAPS_FAILED},\s+#pkts\s+decaps\s+failed\s+\(rcv\)\s+${PKTS_DECAPS_FAILED}\s*$$
  ^\s*#pkts\s+invalid\s+prot\s+\(recv\)\s+${PKTS_INVALID_PROT},\s+#pkts\s+verify\s+failed:\s+${PKTS_VERIFY_FAILED}\s*$$
  ^\s*#pkts\s+invalid\s+identity\s+\(recv\)\s+${PKTS_INVALID_IDENTITY},\s+#pkts\s+invalid\s+len\s+\(rcv\)\s+${PKTS_INVALID_LEN}\s*$$
  ^\s*#pkts\s+replay\s+rollover\s\(send\):\s+${PKTS_REPLAY_ROLLOVER_SEND},\s+#pkts\s+replay\s+rollover\s+\(rcv\)\s+${PKTS_REPLAY_ROLLOVER_RECV}\s*$$
  ^\s*##pkts\s+replay\s+failed\s+\(rcv\):\s+${PKTS_REPLAY_FAILED}\s*$$
  ^\s*#pkts\s+tagged\s+\(send\):\s+${PKTS_TAGGED},\s+#pkts\s+untagged\s\(rcv\):\s+${PKTS_UNTAGGED}\s*$$
  ^\s*#pkts\s+not\s+tagged\s\(send\):\s+${PKTS_NOT_TAGGED},\s+#pkts\s+not\s+untagged\s+\(rcv\):\s+${PKTS_NOT_UNTAGGED}\s*$$
  ^\s*#pkts\s+internal\s+err\s+\(send\):\s+${PKTS_INTERNAL_ERR_SEND},\s+#pkts\s+internal\s+err\s+\(recv\)\s+${PKTS_INTERNAL_ERR_RECV}\s*$$
  ^\s*local\s+crypto\s+endpt.:\s+${LOCAL_CRYPTO_ENDPT},\s+remote\s+crypto\s+endpt.:\s+${REMOTE_CRYPTO_ENDPT}\s*$$
  ^\s*plaintext\s+mtu\s+${PLAINTEXT_MTU},\s+path\s+mtu\s+${PATH_MTU},\s+ip\s+mtu\s+${IP_MTU},\s+ip\s+mtu\s+idb\s+${IP_MTU_IDB}\s*$$
  ^\s*current\s+outbound\s+spi:\s+${CURRENT_OUTBOUND_SPI_HEX}\(${CURRENT_OUTBOUND_SPI_DEC}\)\s*$$
  ^\s*PFS\s+\(Y/N\):\s+${PFS},\s+DH\s+group:\s+${DH_GROUP}\s*$$
  ^\s*${SA_ORIENTATION}\s+${SA_TYPE}\s+sas:\s*$$ -> SAs
  ^\s*$$
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^. -> Error

SAs
  ^\s*${SA_ORIENTATION}\s+${SA_TYPE}\s+sas:\s*$$
  ^\s*spi:\s+${SA_SPI_HEX}\(${SA_SPI_DEC}\)\s*$$
  ^\s*transform:\s+${SA_TRANSFORM}\s+,\s*$$
  ^\s*in\s+use\s+settings\s+={${SA_IN_USE_SETTINGS},*\s+}\s*$$
  ^\s*conn\s+id:\s+${SA_CONN_ID},\s+flow_id:\s+${SA_FLOW_ID},\s+sibling_flags\s+${SA_SIBLING_FLAGS},\s+crypto\s+map:\s+${SA_CRYPTO_MAP}\s*$$
  ^\s*sa\s+timing:\s+remaining\s+key\s+lifetime\s+\(k/sec\):\s+\(${SA_LIFETIME_KBYTES}/${SA_LIFETIME_SEC}\)\s*$$
  ^\s*IV\s+size:\s+${SA_IV_SIZE}\s+bytes\s*$$
  ^\s*replay\s+detection\s+support:\s+${SA_REPLAY_DETECTION_SUPPORT}\s+replay\s+window\s+size:\s+${SA_REPLAY_WINDOW_SIZE}\s*$$
  ^\s*Status:\s+${SA_STATUS}\s*$$ -> Record
  ^\s*interface:\s+${INTERFACE}\s*$$ -> Start
  ^\s*$$
  ^. -> Error
`
