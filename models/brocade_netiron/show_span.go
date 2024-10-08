package brocade_netiron 

type ShowSpan struct {
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

var ShowSpan_Template = `Value mastervlan (\d+)
Value instance (\d+)
Value bridgeid ([0-9a-f]+)
Value bridgemaxage (\d+)
Value bridgehello (\d+)
Value bridgefwddly (\d+)
Value holdtime (\d+)
Value lasttopochange (\d+)
Value topochangecount (\d+)
Value rootbridgeid ([0-9a-f]+)
Value rootpathcost (\d+)
Value designatedbridgeid ([0-9a-f]+)
Value rootport ([0-9a-zA-Z]+)
Value rbmaxage (\d+)
Value rbhello (\d+)
Value rbfwddly (\d+)
Value portstates (\d+)

Start
  ^VLAN ${mastervlan} - STP instance ${instance}
  ^${bridgeid}\s+${bridgemaxage}\s+${bridgehello}\s+${bridgefwddly}\s+${holdtime}\s+${lasttopochange}\s+${topochangecount}
  ^${rootbridgeid}\s+${rootpathcost}\s+${designatedbridgeid}\s+${rootport}\s+${rbmaxage}\s+${rbhello}\s+${rbfwddly} -> Next.Record
`