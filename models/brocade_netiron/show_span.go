package brocade_netiron

type ShowSpan struct {
	Lasttopochange     string `json:"lasttopochange"`
	Rootport           string `json:"rootport"`
	Rbfwddly           string `json:"rbfwddly"`
	Bridgeid           string `json:"bridgeid"`
	Bridgemaxage       string `json:"bridgemaxage"`
	Bridgefwddly       string `json:"bridgefwddly"`
	Topochangecount    string `json:"topochangecount"`
	Designatedbridgeid string `json:"designatedbridgeid"`
	Rbmaxage           string `json:"rbmaxage"`
	Bridgehello        string `json:"bridgehello"`
	Rootbridgeid       string `json:"rootbridgeid"`
	Rootpathcost       string `json:"rootpathcost"`
	Mastervlan         string `json:"mastervlan"`
	Instance           string `json:"instance"`
	Holdtime           string `json:"holdtime"`
	Rbhello            string `json:"rbhello"`
	Portstates         string `json:"portstates"`
}

var ShowSpan_Template string = `Value mastervlan (\d+)
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
