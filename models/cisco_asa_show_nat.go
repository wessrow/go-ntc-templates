package models

type CiscoAsaShowNat struct {
	Nat_section_number	string	`json:"NAT_SECTION_NUMBER"`
	Line_number	string	`json:"LINE_NUMBER"`
	Source_interface	string	`json:"SOURCE_INTERFACE"`
	Destination_interface	string	`json:"DESTINATION_INTERFACE"`
	Source_type	string	`json:"SOURCE_TYPE"`
	Source_real	string	`json:"SOURCE_REAL"`
	Source_mapped	string	`json:"SOURCE_MAPPED"`
	Destination_real	string	`json:"DESTINATION_REAL"`
	Destination_mapped	string	`json:"DESTINATION_MAPPED"`
	Service_protocol	string	`json:"SERVICE_PROTOCOL"`
	Service_real	string	`json:"SERVICE_REAL"`
	Service_mapped	string	`json:"SERVICE_MAPPED"`
	Extended	string	`json:"EXTENDED"`
	Flat	string	`json:"FLAT"`
	Include_reserve	string	`json:"INCLUDE_RESERVE"`
	Round_robin	string	`json:"ROUND_ROBIN"`
	Net_to_net	string	`json:"NET_TO_NET"`
	Dns	string	`json:"DNS"`
	Unidirectional	string	`json:"UNIDIRECTIONAL"`
	No_proxy_arp	string	`json:"NO_PROXY_ARP"`
	Route_lookup	string	`json:"ROUTE_LOOKUP"`
	Inactive	string	`json:"INACTIVE"`
	Description	string	`json:"DESCRIPTION"`
	Translate_hits	string	`json:"TRANSLATE_HITS"`
	Untranslate_hits	string	`json:"UNTRANSLATE_HITS"`
}