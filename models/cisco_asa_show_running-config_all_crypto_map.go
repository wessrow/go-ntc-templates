package models

type CiscoAsaShowRunningConfigAllCryptoMap struct {
	Matched_address	string	`json:"MATCHED_ADDRESS"`
	Connection_type	string	`json:"CONNECTION_TYPE"`
	Map	string	`json:"MAP"`
	Seq	string	`json:"SEQ"`
	Pfs	string	`json:"PFS"`
	Peer	string	`json:"PEER"`
	Ikev1_phase1_mode	string	`json:"IKEv1_PHASE1_MODE"`
	Ikev1_transform_set	string	`json:"IKEv1_TRANSFORM_SET"`
	Ikev2_mode	string	`json:"IKEv2_MODE"`
	Isakmp_dynamic	string	`json:"ISAKMP_DYNAMIC"`
	Interface	string	`json:"INTERFACE"`
	Transform	string	`json:"TRANSFORM"`
	Sa_sec	string	`json:"SA_SEC"`
	Sa_kb	string	`json:"SA_KB"`
	Tfc_packets	string	`json:"TFC_PACKETS"`
}