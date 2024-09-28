package models

type CiscoIosShowAdjacency struct {
	Interface	string	`json:"INTERFACE"`
	Endpoint	string	`json:"ENDPOINT"`
	Rewrite_headers	[]string	`json:"REWRITE_HEADERS"`
	Recursive_interface	string	`json:"RECURSIVE_INTERFACE"`
	Recursive_nexthop	string	`json:"RECURSIVE_NEXTHOP"`
}