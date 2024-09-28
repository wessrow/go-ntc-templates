package models

type BrocadeFastironShowSpan struct {
	Mastervlan	string	`json:"mastervlan"`
	Vlan	string	`json:"vlan"`
	Rootbridgeid	string	`json:"rootbridgeid"`
	Rootpathcost	string	`json:"rootpathcost"`
	Rootport	string	`json:"rootport"`
	Priorityhex	string	`json:"priorityhex"`
	Bridgemaxage	string	`json:"bridgemaxage"`
	Bridgehello	string	`json:"bridgehello"`
	Holdtime	string	`json:"holdtime"`
	Bridgefwddly	string	`json:"bridgefwddly"`
	Lasttopochange	string	`json:"lasttopochange"`
	Topochangecount	string	`json:"topochangecount"`
	Bridgeaddress	string	`json:"bridgeaddress"`
	Portstates	string	`json:"portstates"`
}