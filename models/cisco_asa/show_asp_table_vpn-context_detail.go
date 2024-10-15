package cisco_asa

type ShowAspTableVpnContextDetail struct {
	Vpn_ctx    string `json:"VPN_CTX"`
	State      string `json:"STATE"`
	Bad_pkts   string `json:"BAD_PKTS"`
	Pointer    string `json:"POINTER"`
	Sa         string `json:"SA"`
	Group      string `json:"GROUP"`
	Rekey_call string `json:"REKEY_CALL"`
	Vpn_filter string `json:"VPN_FILTER"`
	Spoof      string `json:"SPOOF"`
	Bad_crypto string `json:"BAD_CRYPTO"`
	Peer_ip    string `json:"PEER_IP"`
	Flags      string `json:"FLAGS"`
	Spi        string `json:"SPI"`
	Pkts       string `json:"PKTS"`
	Bad_spi    string `json:"BAD_SPI"`
	Rekey_pkt  string `json:"REKEY_PKT"`
}

var ShowAspTableVpnContextDetail_Template string = `Value Required VPN_CTX (\S+)
Value PEER_IP (\S+)
Value POINTER (\S+)
Value STATE (\S+)
Value FLAGS (\S+)
Value SA (\S+)
Value SPI (\S+)
Value GROUP (\d+)
Value PKTS (\d+)
Value BAD_PKTS (\d+)
Value BAD_SPI (\d+)
Value SPOOF (\d+)
Value BAD_CRYPTO (\d+)
Value REKEY_PKT (\d+)
Value REKEY_CALL (\d+)
Value VPN_FILTER (\S+)

Start
  ^VPN\s+CTX -> Continue.Record
  ^VPN\s+CTX\s+=\s+${VPN_CTX}
  ^Peer\s+IP\s+=\s+${PEER_IP}
  ^Pointer\s+=\s+${POINTER}
  ^State\s+=\s+${STATE}
  ^Flags\s+=\s+${FLAGS}
  ^SA\s+=\s+${SA}
  ^SPI\s+=\s+${SPI}
  ^Group\s+=\s+${GROUP}
  ^Pkts\s+=\s+${PKTS}
  ^Bad\s+Pkts\s+=\s+${BAD_PKTS}
  ^Bad\s+SPI\s+=\s+${BAD_SPI}
  ^Spoof\s+=\s+${SPOOF}
  ^Bad\s+Crypto\s+=\s+${BAD_CRYPTO}
  ^Rekey\s+Pkt\s+=\s+${REKEY_PKT}
  ^Rekey\s+Call\s+=\s+${REKEY_CALL}
  ^VPN\s+Filter\s+=\s+${VPN_FILTER}
  ^\s*
  ^. -> Error
`
