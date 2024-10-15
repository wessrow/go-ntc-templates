package brocade_fastiron

type ShowTopo struct {
	Topogroup    string `json:"topogroup"`
	Mastervlan   string `json:"mastervlan"`
	L2proto      string `json:"l2proto"`
	Membervlans  string `json:"membervlans"`
	Controlports string `json:"controlports"`
	Freeports    string `json:"freeports"`
}

var ShowTopo_Template string = `Value topogroup (\d+)
Value mastervlan (\d+)
Value l2proto (MRP|STP|RSTP|VSRP|ERP)
Value membervlans ([0-9\ to]+)
Value controlports ([0-9\ \/eto]+)
Value freeports ([0-9\ \/eto]+)

Start
  ^Topology Group\s+${topogroup} -> Continue
  ^\s+master-vlan\s+${mastervlan} -> Continue
  ^\s+ethernet\s+([0-9\/]+)\s+${l2proto} -> Continue
  ^\s+Per -> Record`
