package cisco_ios

type ShowIpRouteSummary struct {
	Subnets      string `json:"SUBNETS"`
	Replicates   string `json:"REPLICATES"`
	Overhead     string `json:"OVERHEAD"`
	Memory       string `json:"MEMORY"`
	Route_source string `json:"ROUTE_SOURCE"`
	Name         string `json:"NAME"`
	Networks     string `json:"NETWORKS"`
}

var ShowIpRouteSummary_Template string = `Value Required ROUTE_SOURCE ([^\s][\S ]+?)
Value Filldown NAME (.*)
Value NETWORKS (\d*)
Value SUBNETS (\d*)
Value REPLICATES (\d*)
Value OVERHEAD (\d*)
Value MEMORY (\d*)

Start
  # Checking for header
  ^IP\s+routing\s+table\s+name\s+is\s+${NAME}\s*$$
  ^IP\s+routing\s+table\s+maximum-paths\s+is\s+\d+\s*$$
  ^Route\s+Source\s+Networks\s+Subnets\s+Overhead\s+Memory\s+\(bytes\)\s*$$ -> RouteData1
  ^Route\s+Source\s+Networks\s+Subnets\s+Replicates\s+Overhead\s+Memory\s+\(bytes\)\s*$$ -> RouteData2
  ^\s*$$
  ^. -> Error

RouteData1
  ^${ROUTE_SOURCE}\s{2,16}${NETWORKS}\s{1,12}${SUBNETS}\s{1,12}${OVERHEAD}\s{1,12}${MEMORY}\s*$$ -> Record
  ^\s+External:\s+\d+\s+Internal:\s+\d+\s+Local:\s+\d+\s*$$
  ^\s+Level\s+1:\s+\d+\s+Level\s+2:\s+\d+\s+Inter-area:\s+\d+\s*$$
  ^\s+Intra-area:\s+\d+\s+Inter-area:\s+\d+\s+External-1:\s+\d+\s+External-2:\s+\d+\s*$$
  ^\s+NSSA\s+External-1:\s+\d+\s+NSSA\s+External-2:\s+\d+\s*$$
  ^\s*$$
  ^. -> Error

RouteData2
  ^${ROUTE_SOURCE}\s{2,16}${NETWORKS}\s{1,12}${SUBNETS}\s{1,12}${REPLICATES}\s{1,12}${OVERHEAD}\s{1,12}${MEMORY}\s*$$ -> Record
  ^\s+External:\s+\d+\s+Internal:\s+\d+\s+Local:\s+\d+\s*$$
  ^\s+Level\s+1:\s+\d+\s+Level\s+2:\s+\d+\s+Inter-area:\s+\d+\s*$$
  ^\s+Intra-area:\s+\d+\s+Inter-area:\s+\d+\s+External-1:\s+\d+\s+External-2:\s+\d+\s*$$
  ^\s+NSSA\s+External-1:\s+\d+\s+NSSA\s+External-2:\s+\d+\s*$$
  ^\s*$$
  ^. -> Error
`
