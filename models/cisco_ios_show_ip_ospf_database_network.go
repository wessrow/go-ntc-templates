package models

type CiscoIosShowIpOspfDatabaseNetwork struct {
	Router_id	string	`json:"ROUTER_ID"`
	Process_id	string	`json:"PROCESS_ID"`
	Area	string	`json:"AREA"`
	Lsa_age	string	`json:"LSA_AGE"`
	Lsa_options	string	`json:"LSA_OPTIONS"`
	Lsa_type	string	`json:"LSA_TYPE"`
	Lsa_id	string	`json:"LSA_ID"`
	Lsa_adv_router	string	`json:"LSA_ADV_ROUTER"`
	Lsa_seq_number	string	`json:"LSA_SEQ_NUMBER"`
	Lsa_checksum	string	`json:"LSA_CHECKSUM"`
	Lsa_length	string	`json:"LSA_LENGTH"`
	Lsa_network_mask	string	`json:"LSA_NETWORK_MASK"`
	Lsa_abr	string	`json:"LSA_ABR"`
	Lsa_asbr	string	`json:"LSA_ASBR"`
	Ls_att_router	string	`json:"LS_ATT_ROUTER"`
}