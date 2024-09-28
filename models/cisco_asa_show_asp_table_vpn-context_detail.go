package models

type CiscoAsaShowAspTableVpnContextDetail struct {
	Vpn_ctx	string	`json:"VPN_CTX"`
	Peer_ip	string	`json:"PEER_IP"`
	Pointer	string	`json:"POINTER"`
	State	string	`json:"STATE"`
	Flags	string	`json:"FLAGS"`
	Sa	string	`json:"SA"`
	Spi	string	`json:"SPI"`
	Group	string	`json:"GROUP"`
	Pkts	string	`json:"PKTS"`
	Bad_pkts	string	`json:"BAD_PKTS"`
	Bad_spi	string	`json:"BAD_SPI"`
	Spoof	string	`json:"SPOOF"`
	Bad_crypto	string	`json:"BAD_CRYPTO"`
	Rekey_pkt	string	`json:"REKEY_PKT"`
	Rekey_call	string	`json:"REKEY_CALL"`
	Vpn_filter	string	`json:"VPN_FILTER"`
}