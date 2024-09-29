package models

type MikrotikRouterosRoutingBgpPeerPrintStatus struct {
	Bgp_neighbor	string	`json:"BGP_NEIGHBOR"`
	Neighbor_as	string	`json:"NEIGHBOR_AS"`
	Up_down	string	`json:"UP_DOWN"`
	State	string	`json:"STATE"`
}

var MikrotikRouterosRoutingBgpPeerPrintStatus_Template = `Value BGP_NEIGHBOR (\S+)
Value NEIGHBOR_AS (\d+)
Value UP_DOWN (\S+)
Value STATE (\S+)

Start
  ^\s*Flags:\s+X\s+-\s+disabled,\s+E\s+-\s+established\s*$$
  ^.*remote-address=${BGP_NEIGHBOR}.*remote-as=${NEIGHBOR_AS}.*(?:uptime=${UP_DOWN})?.*$$
  ^.*state=${STATE}.*$$ -> Record
  ^\s*$$
  ^. -> Error

`