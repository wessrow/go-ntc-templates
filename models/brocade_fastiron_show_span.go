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

var BrocadeFastironShowSpan_Template = `Value Filldown mastervlan (\d+)
Value Required vlan (\d+)
Value rootbridgeid ([0-9a-f]+)
Value rootpathcost (\d+)
Value rootport ([0-9a-zA-Z]+)
Value priorityhex ([0-9a-f]+)
Value bridgemaxage (\d+)
Value bridgehello (\d+)
Value holdtime (\d+)
Value bridgefwddly (\d+)
Value lasttopochange (\d+)
Value topochangecount (\d+)
Value bridgeaddress ([0-9a-f]+)
Value portstates (\d+)

Start
  ^STP instance owned by VLAN ${mastervlan}
  ^\s+${vlan}\s+${rootbridgeid}\s+${rootpathcost}\s+${rootport}\s+${priorityhex}\s+${bridgemaxage}\s+${bridgehello}\s+${holdtime}\s+${bridgefwddly}\s+${lasttopochange}\s+${bridgeaddress} -> Continue
  ^\s+${vlan}\s+${rootbridgeid}\s+${rootpathcost}\s+${rootport}\s+${priorityhex}\s+${bridgemaxage}\s+${bridgehello}\s+${holdtime}\s+${bridgefwddly}\s+${lasttopochange}\s+${topochangecount}\s+${bridgeaddress} -> Next.Record
`