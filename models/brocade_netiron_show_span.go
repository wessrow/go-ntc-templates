package models

type BrocadeNetironShowSpan struct {
	Mastervlan	string	`json:"mastervlan"`
	Instance	string	`json:"instance"`
	Bridgeid	string	`json:"bridgeid"`
	Bridgemaxage	string	`json:"bridgemaxage"`
	Bridgehello	string	`json:"bridgehello"`
	Bridgefwddly	string	`json:"bridgefwddly"`
	Holdtime	string	`json:"holdtime"`
	Lasttopochange	string	`json:"lasttopochange"`
	Topochangecount	string	`json:"topochangecount"`
	Rootbridgeid	string	`json:"rootbridgeid"`
	Rootpathcost	string	`json:"rootpathcost"`
	Designatedbridgeid	string	`json:"designatedbridgeid"`
	Rootport	string	`json:"rootport"`
	Rbmaxage	string	`json:"rbmaxage"`
	Rbhello	string	`json:"rbhello"`
	Rbfwddly	string	`json:"rbfwddly"`
	Portstates	string	`json:"portstates"`
}