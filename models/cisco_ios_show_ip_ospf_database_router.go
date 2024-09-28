package models

type CiscoIosShowIpOspfDatabaseRouter struct {
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
	Lsa_num_links	string	`json:"LSA_NUM_LINKS"`
	Lsa_abr	string	`json:"LSA_ABR"`
	Lsa_asbr	string	`json:"LSA_ASBR"`
	Ls_link_type	string	`json:"LS_LINK_TYPE"`
	Ls_link_id	string	`json:"LS_LINK_ID"`
	Ls_link_data	string	`json:"LS_LINK_DATA"`
	Ls_mtid_metrics	string	`json:"LS_MTID_METRICS"`
	Ls_tos_0_metrics	string	`json:"LS_TOS_0_METRICS"`
}