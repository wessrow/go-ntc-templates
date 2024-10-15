package cisco_asa

type ShowCryptoIpsecSa struct {
	Local_identity_addr                   string `json:"LOCAL_IDENTITY_ADDR"`
	Pmtus_received                        string `json:"PMTUS_RECEIVED"`
	Inbound_remaining_lifetime_kilobytes  string `json:"INBOUND_REMAINING_LIFETIME_KILOBYTES"`
	Outbound_authentication               string `json:"OUTBOUND_AUTHENTICATION"`
	Packets_not_compressed                string `json:"PACKETS_NOT_COMPRESSED"`
	Sequence_number                       string `json:"SEQUENCE_NUMBER"`
	Remote_identity_protocol              string `json:"REMOTE_IDENTITY_PROTOCOL"`
	Inbound_encryption                    string `json:"INBOUND_ENCRYPTION"`
	Outbound_connection_id                string `json:"OUTBOUND_CONNECTION_ID"`
	Pre_fragment_failures                 string `json:"PRE_FRAGMENT_FAILURES"`
	Remote_crypto_endpoint_name           string `json:"REMOTE_CRYPTO_ENDPOINT_NAME"`
	Current_outbound_spi                  string `json:"CURRENT_OUTBOUND_SPI"`
	Outbound_replay_detection             string `json:"OUTBOUND_REPLAY_DETECTION"`
	Current_inbound_spi                   string `json:"CURRENT_INBOUND_SPI"`
	Outbound_crypto_map                   string `json:"OUTBOUND_CRYPTO_MAP"`
	Outbound_remaining_lifetime_kilobytes string `json:"OUTBOUND_REMAINING_LIFETIME_KILOBYTES"`
	Local_identity_port                   string `json:"LOCAL_IDENTITY_PORT"`
	Dynamic_peer_name                     string `json:"DYNAMIC_PEER_NAME"`
	Packets_compress_failed               string `json:"PACKETS_COMPRESS_FAILED"`
	Packets_decompress_failed             string `json:"PACKETS_DECOMPRESS_FAILED"`
	Pmtus_sent                            string `json:"PMTUS_SENT"`
	Path_mtu                              string `json:"PATH_MTU"`
	Inbound_crypto_map                    string `json:"INBOUND_CRYPTO_MAP"`
	Inbound_iv_size                       string `json:"INBOUND_IV_SIZE"`
	Outbound_spi_hex                      string `json:"OUTBOUND_SPI_HEX"`
	Crypto_map_tag                        string `json:"CRYPTO_MAP_TAG"`
	Current_peer                          string `json:"CURRENT_PEER"`
	Packets_encapsulated                  string `json:"PACKETS_ENCAPSULATED"`
	Local_crypto_endpoint_name            string `json:"LOCAL_CRYPTO_ENDPOINT_NAME"`
	Local_address                         string `json:"LOCAL_ADDRESS"`
	Ipsec_overhead                        string `json:"IPSEC_OVERHEAD"`
	Inbound_spi_hex                       string `json:"INBOUND_SPI_HEX"`
	Inbound_remaining_lifetime            string `json:"INBOUND_REMAINING_LIFETIME"`
	Outbound_spi_integer                  string `json:"OUTBOUND_SPI_INTEGER"`
	Outbound_settings_in_use              string `json:"OUTBOUND_SETTINGS_IN_USE"`
	Remote_identity_addr                  string `json:"REMOTE_IDENTITY_ADDR"`
	Packets_verified                      string `json:"PACKETS_VERIFIED"`
	Decap_frags_needing_reassembly        string `json:"DECAP_FRAGS_NEEDING_REASSEMBLY"`
	Send_errors                           string `json:"SEND_ERRORS"`
	Inbound_spi_integer                   string `json:"INBOUND_SPI_INTEGER"`
	Outbound_encryption                   string `json:"OUTBOUND_ENCRYPTION"`
	Interface                             string `json:"INTERFACE"`
	Packets_decrypted                     string `json:"PACKETS_DECRYPTED"`
	Packets_compressed                    string `json:"PACKETS_COMPRESSED"`
	Local_crypto_endpoint                 string `json:"LOCAL_CRYPTO_ENDPOINT"`
	Remote_identity_port                  string `json:"REMOTE_IDENTITY_PORT"`
	Current_peer_name                     string `json:"CURRENT_PEER_NAME"`
	Packets_decapsulated                  string `json:"PACKETS_DECAPSULATED"`
	Remote_crypto_endpoint                string `json:"REMOTE_CRYPTO_ENDPOINT"`
	Local_identity_protocol               string `json:"LOCAL_IDENTITY_PROTOCOL"`
	Outbound_remaining_lifetime           string `json:"OUTBOUND_REMAINING_LIFETIME"`
	Outbound_iv_size                      string `json:"OUTBOUND_IV_SIZE"`
	Dynamic_peer                          string `json:"DYNAMIC_PEER"`
	Fragments_created                     string `json:"FRAGMENTS_CREATED"`
	Receive_errors                        string `json:"RECEIVE_ERRORS"`
	Inbound_settings_in_use               string `json:"INBOUND_SETTINGS_IN_USE"`
	Inbound_authentication                string `json:"INBOUND_AUTHENTICATION"`
	Inbound_connection_id                 string `json:"INBOUND_CONNECTION_ID"`
	Inbound_replay_detection              string `json:"INBOUND_REPLAY_DETECTION"`
	Local_identity_mask                   string `json:"LOCAL_IDENTITY_MASK"`
	Packets_digested                      string `json:"PACKETS_DIGESTED"`
	Packets_decompressed                  string `json:"PACKETS_DECOMPRESSED"`
	Pre_fragment_success                  string `json:"PRE_FRAGMENT_SUCCESS"`
	Inbound_slot                          string `json:"INBOUND_SLOT"`
	Outbound_slot                         string `json:"OUTBOUND_SLOT"`
	Local_address_name                    string `json:"LOCAL_ADDRESS_NAME"`
	Remote_identity_mask                  string `json:"REMOTE_IDENTITY_MASK"`
	Packets_encrypted                     string `json:"PACKETS_ENCRYPTED"`
	Media_mtu                             string `json:"MEDIA_MTU"`
}

var ShowCryptoIpsecSa_Template string = `Value Filldown INTERFACE (\S+)
Value Filldown CRYPTO_MAP_TAG (\S+)
Value Filldown SEQUENCE_NUMBER (\d+)
Value Filldown LOCAL_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value Filldown LOCAL_ADDRESS_NAME (\S+)
Value LOCAL_IDENTITY_ADDR (\d+\.\d+\.\d+\.\d+)
Value LOCAL_IDENTITY_MASK (\d+\.\d+\.\d+\.\d+)
Value LOCAL_IDENTITY_PROTOCOL (\d+)
Value LOCAL_IDENTITY_PORT (\d+)
Value REMOTE_IDENTITY_ADDR (\d+\.\d+\.\d+\.\d+)
Value REMOTE_IDENTITY_MASK (\d+\.\d+\.\d+\.\d+)
Value REMOTE_IDENTITY_PROTOCOL (\d+)
Value REMOTE_IDENTITY_PORT (\d+)
Value CURRENT_PEER (\d+\.\d+\.\d+\.\d+)
Value DYNAMIC_PEER (\d+\.\d+\.\d+\.\d+)
Value CURRENT_PEER_NAME (\S+)
Value DYNAMIC_PEER_NAME (\S+)
Value PACKETS_ENCAPSULATED (\d+)
Value PACKETS_ENCRYPTED (\d+)
Value PACKETS_DIGESTED (\d+)
Value PACKETS_DECAPSULATED (\d+)
Value PACKETS_DECRYPTED (\d+)
Value PACKETS_VERIFIED (\d+)
Value PACKETS_COMPRESSED (\d+)
Value PACKETS_DECOMPRESSED (\d+)
Value PACKETS_NOT_COMPRESSED (\d+)
Value PACKETS_COMPRESS_FAILED (\d+)
Value PACKETS_DECOMPRESS_FAILED (\d+)
Value PRE_FRAGMENT_SUCCESS (\d+)
Value PRE_FRAGMENT_FAILURES (\d+)
Value FRAGMENTS_CREATED (\d+)
Value PMTUS_SENT (\d+)
Value PMTUS_RECEIVED (\d+)
Value DECAP_FRAGS_NEEDING_REASSEMBLY (\d+)
Value SEND_ERRORS (\d+)
Value RECEIVE_ERRORS (\d+)
Value LOCAL_CRYPTO_ENDPOINT (\d+\.\d+\.\d+\.\d+)
Value REMOTE_CRYPTO_ENDPOINT (\d+\.\d+\.\d+\.\d+)
Value LOCAL_CRYPTO_ENDPOINT_NAME (\S+)
Value REMOTE_CRYPTO_ENDPOINT_NAME (\S+)
Value PATH_MTU (\d+)
Value IPSEC_OVERHEAD (\d+)
Value MEDIA_MTU (\d+)
Value CURRENT_INBOUND_SPI (\w+)
Value CURRENT_OUTBOUND_SPI (\w+)
Value INBOUND_SPI_HEX (\w+)
Value INBOUND_SPI_INTEGER (\d+)
Value INBOUND_ENCRYPTION (\S+)
Value INBOUND_AUTHENTICATION (\S+)
Value INBOUND_SETTINGS_IN_USE (.*)
Value INBOUND_SLOT (\d+)
Value INBOUND_CONNECTION_ID (\d+)
Value INBOUND_CRYPTO_MAP (\S+)
Value INBOUND_REMAINING_LIFETIME (\d+)
Value INBOUND_REMAINING_LIFETIME_KILOBYTES (\d+)
Value INBOUND_IV_SIZE (\d+\s+\w+)
Value INBOUND_REPLAY_DETECTION (\w+)
Value OUTBOUND_SPI_HEX (\w+)
Value OUTBOUND_SPI_INTEGER (\d+)
Value OUTBOUND_ENCRYPTION (\S+)
Value OUTBOUND_AUTHENTICATION (\S+)
Value OUTBOUND_SETTINGS_IN_USE (.*)
Value OUTBOUND_SLOT (\d+)
Value OUTBOUND_CONNECTION_ID (\d+)
Value OUTBOUND_CRYPTO_MAP (\S+)
Value OUTBOUND_REMAINING_LIFETIME (\d+)
Value OUTBOUND_REMAINING_LIFETIME_KILOBYTES (\d+)
Value OUTBOUND_IV_SIZE (\d+\s+\w+)
Value OUTBOUND_REPLAY_DETECTION (\w+)

Start
  ^interface:\s+${INTERFACE}\s*
  ^\s+Crypto map tag:\s+${CRYPTO_MAP_TAG},\s+local addr:\s+(?:${LOCAL_ADDRESS}|${LOCAL_ADDRESS_NAME})\s*
  ^\s+Crypto map tag:\s+${CRYPTO_MAP_TAG}, seq num:\s+${SEQUENCE_NUMBER},\s+local addr:\s+(?:${LOCAL_ADDRESS}|${LOCAL_ADDRESS_NAME})\s*
  ^\s+local\s+ident\s+\(addr\/mask\/prot\/port\):\s+\(${LOCAL_IDENTITY_ADDR}\/${LOCAL_IDENTITY_MASK}\/${LOCAL_IDENTITY_PROTOCOL}\/${LOCAL_IDENTITY_PORT}\)\s*
  ^\s+remote\s+ident\s+\(addr/mask/prot/port\):\s+\(${REMOTE_IDENTITY_ADDR}\/${REMOTE_IDENTITY_MASK}\/${REMOTE_IDENTITY_PROTOCOL}\/${REMOTE_IDENTITY_PORT}\)\s*
  ^\s+current_peer:\s+(?:${CURRENT_PEER}|${CURRENT_PEER_NAME})\s*
  ^\s+dynamic\s+allocated\s+peer\s+ip:\s+(?:${DYNAMIC_PEER}|${DYNAMIC_PEER_NAME})\s*
  ^\s+#pkts\s+encaps:\s+${PACKETS_ENCAPSULATED},\s+#pkts\s+encrypt:\s+${PACKETS_ENCRYPTED},\s+#pkts\s+digest:\s+${PACKETS_DIGESTED}\s*
  ^\s+#pkts\s+decaps:\s+${PACKETS_DECAPSULATED},\s+#pkts\s+decrypt:\s+${PACKETS_DECRYPTED},\s+#pkts\s+verify:\s+${PACKETS_VERIFIED}\s*
  ^\s+#pkts\s+compressed:\s+${PACKETS_COMPRESSED},\s+#pkts\s+decompressed:\s+${PACKETS_DECOMPRESSED}\s*
  ^\s+#pkts\s+not\s+compressed:\s+${PACKETS_NOT_COMPRESSED},\s+#pkts\s+comp\s+failed:\s+${PACKETS_COMPRESS_FAILED},\s+#pkts\s+decomp\s+failed:\s+${PACKETS_DECOMPRESS_FAILED}\s*
  ^\s+#pre-frag\s+successes:\s+${PRE_FRAGMENT_SUCCESS},\s+#pre-frag\s+failures:\s+${PRE_FRAGMENT_FAILURES},\s+#fragments\s+created:\s+${FRAGMENTS_CREATED}\s*
  ^\s+#PMTUs\s+sent:\s+${PMTUS_SENT},\s+#PMTUs\s+rcvd:\s+${PMTUS_RECEIVED},\s+#decapsulated\s+fra?gs\s+needing\s+reassembly:\s+${DECAP_FRAGS_NEEDING_REASSEMBLY}\s*
  ^\s+#send\s+errors:\s+${SEND_ERRORS},\s+#recv\s+errors:\s+${RECEIVE_ERRORS}\s*
  ^\s+local\s+crypto\s+endpt\.:\s+${LOCAL_CRYPTO_ENDPOINT},\s+remote\s+crypto\s+endpt\.:\s+${REMOTE_CRYPTO_ENDPOINT}\s*
  ^\s+local\s+crypto\s+endpt\.:\s+(?:${LOCAL_CRYPTO_ENDPOINT}|${LOCAL_CRYPTO_ENDPOINT_NAME})(\/\d+),\s+remote\s+crypto\s+endpt\.:\s+(?:${REMOTE_CRYPTO_ENDPOINT}|${REMOTE_CRYPTO_ENDPOINT_NAME})(\/\d+)\s*
  ^\s+path\s+mtu\s+${PATH_MTU},\s+ipsec\s+overhead\s+${IPSEC_OVERHEAD}(\(\d+\))?,\s+media\s+mtu\s+${MEDIA_MTU}\s*
  ^\s+current\s+outbound\s+spi:\s+${CURRENT_OUTBOUND_SPI}\s*
  ^\s+current\s+inbound\s+spi\s+:\s+${CURRENT_INBOUND_SPI}\s*
  ^\s+inbound\s+esp\s+sas:\s* -> Inbound
  ^\s+outbound\s+esp\s+sas:\s* -> Outbound

Inbound
  ^\s+spi:\s+${INBOUND_SPI_HEX}\s+\(${INBOUND_SPI_INTEGER}\)\s*
  ^\s+transform:\s+${INBOUND_ENCRYPTION}\s+${INBOUND_AUTHENTICATION}\s*
  ^\s+in\s+use\s+settings\s+=\{${INBOUND_SETTINGS_IN_USE},\s+\}\s*
  ^\s+slot:\s+${INBOUND_SLOT},\s+conn_id:\s+${INBOUND_CONNECTION_ID},\s+crypto-map:\s+${INBOUND_CRYPTO_MAP}\s*
  ^\s+sa\s+timing:\s+remaining\s+key\s+lifetime\s+\(sec\):\s+${INBOUND_REMAINING_LIFETIME}\s*
  ^\s+sa\s+timing:\s+remaining\s+key\s+lifetime\s+\(kB\/sec\):\s+\(${INBOUND_REMAINING_LIFETIME_KILOBYTES}\/${INBOUND_REMAINING_LIFETIME}\)\s*
  ^\s+IV\s+size:\s+${INBOUND_IV_SIZE}\s*
  ^\s+replay\s+detection\s+support:\s+${INBOUND_REPLAY_DETECTION}\s* -> Start

Outbound
  ^\s+spi:\s+${OUTBOUND_SPI_HEX}\s+\(${OUTBOUND_SPI_INTEGER}\)\s*
  ^\s+transform:\s+${OUTBOUND_ENCRYPTION}\s+${OUTBOUND_AUTHENTICATION}\s*
  ^\s+in\s+use\s+settings\s+=\{${OUTBOUND_SETTINGS_IN_USE},\s+\}\s*
  ^\s+slot:\s+${OUTBOUND_SLOT},\s+conn_id:\s+${OUTBOUND_CONNECTION_ID},\s+crypto-map:\s+${OUTBOUND_CRYPTO_MAP}\s*
  ^\s+sa\s+timing:\s+remaining\s+key\s+lifetime\s+\(sec\):\s+${OUTBOUND_REMAINING_LIFETIME}\s*
  ^\s+sa\s+timing:\s+remaining\s+key\s+lifetime\s+\(kB\/sec\):\s+\(${OUTBOUND_REMAINING_LIFETIME_KILOBYTES}\/${OUTBOUND_REMAINING_LIFETIME}\)\s*
  ^\s+IV\s+size:\s+${OUTBOUND_IV_SIZE}\s*
  ^\s+replay\s+detection\s+support:\s+${OUTBOUND_REPLAY_DETECTION}\s* -> Record Start

EOF
`
