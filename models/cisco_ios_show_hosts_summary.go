package models

type CiscoIosShowHostsSummary struct {
	Default_domain	string	`json:"DEFAULT_DOMAIN"`
	Name_servers	string	`json:"NAME_SERVERS"`
	Local_cache_entries	string	`json:"LOCAL_CACHE_ENTRIES"`
	Dynamic_cache_entries	string	`json:"DYNAMIC_CACHE_ENTRIES"`
}