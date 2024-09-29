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

var CiscoAsaShowNat_Template = `Value Filldown NAT_SECTION_NUMBER (\d)
Value Required LINE_NUMBER (\d+)
Value SOURCE_INTERFACE (\S+)
Value DESTINATION_INTERFACE (\S+)
Value SOURCE_TYPE (static|dynamic)
Value SOURCE_REAL (any|\S+)
Value SOURCE_MAPPED (any|interface\s+ipv6|interface|pat-pool\s+interface|pat-pool\s+\S+|\S+)
Value DESTINATION_REAL (interface|interface\s+ipv6|\S+)
Value DESTINATION_MAPPED (any|\S+)
Value SERVICE_PROTOCOL (tcp|udp|sctp)
Value SERVICE_REAL (any|\S+)
Value SERVICE_MAPPED (\S+)
Value EXTENDED (extended)
Value FLAT (flat)
Value INCLUDE_RESERVE (include-reserve)
Value ROUND_ROBIN (round-robin)
Value NET_TO_NET (net-to-net)
Value DNS (dns)
Value UNIDIRECTIONAL (unidirectional)
Value NO_PROXY_ARP (no-proxy-arp)
Value ROUTE_LOOKUP (route-lookup)
Value INACTIVE (inactive)
Value DESCRIPTION (.*)
Value TRANSLATE_HITS (\d+)
Value UNTRANSLATE_HITS (\d+)

Start
  # Section 1 = Manual NAT
  # Section 2 = Auto NAT
  # Section 3 = After-auto Manual NAT
  ^(Manual|Auto) NAT Policies \(Section ${NAT_SECTION_NUMBER}\)\s*
  ^${LINE_NUMBER}\s+\(${SOURCE_INTERFACE}\)\s+to\s+\(${DESTINATION_INTERFACE}\)\s+source\s+${SOURCE_TYPE}\s+${SOURCE_REAL}\s+${SOURCE_MAPPED}\s*(destination\s+static\s+${DESTINATION_REAL}\s+${DESTINATION_MAPPED})?\s*(service\s+${SERVICE_PROTOCOL}?\s*${SERVICE_REAL}\s+${SERVICE_MAPPED})?\s*${EXTENDED}?\s*${FLAT}?\s*${INCLUDE_RESERVE}?\s*${ROUND_ROBIN}?\s*${NET_TO_NET}?\s*${DNS}?\s*${UNIDIRECTIONAL}?\s*${NO_PROXY_ARP}?\s*${ROUTE_LOOKUP}?\s*${INACTIVE}?\s*(description\s+${DESCRIPTION})?
  ^\s+translate_hits\s+=\s+${TRANSLATE_HITS},\s+untranslate_hits\s+=\s+${UNTRANSLATE_HITS} -> Record
  ^\s*
  ^. -> Error

`