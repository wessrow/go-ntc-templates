package mikrotik_routeros

type RoutingOspfNeighborPrint struct {
	Neighbor_id string `json:"NEIGHBOR_ID"`
	Priority    string `json:"PRIORITY"`
	State       string `json:"STATE"`
	Up_time     string `json:"UP_TIME"`
	Address     string `json:"ADDRESS"`
	Interface   string `json:"INTERFACE"`
}

var RoutingOspfNeighborPrint_Template string = `Value NEIGHBOR_ID (\S+)
Value PRIORITY (\d+)
Value STATE (\S+)
Value UP_TIME (\S+)
Value ADDRESS (\S+)
Value INTERFACE (\S+)

Start
  ^.*router-id=${NEIGHBOR_ID}.*address=${ADDRESS}.*interface=${INTERFACE}.*priority=${PRIORITY}.*state=\"${STATE}\".*adjacency=${UP_TIME}.*$$ -> Record
  ^\s*$$
  ^. -> Error
`
