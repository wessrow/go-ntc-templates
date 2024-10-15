package alcatel_sros

type ShowRouterBgpSummaryFamily struct {
	Inq      string `json:"INQ"`
	Up_down  string `json:"UP_DOWN"`
	Act      string `json:"ACT"`
	Sent     string `json:"SENT"`
	Pktrcvd  string `json:"PKTRCVD"`
	Pktsent  string `json:"PKTSENT"`
	Outq     string `json:"OUTQ"`
	State    string `json:"STATE"`
	Rcv      string `json:"RCV"`
	Neighbor string `json:"NEIGHBOR"`
	As       string `json:"AS"`
}

var ShowRouterBgpSummaryFamily_Template string = `Value NEIGHBOR (\d+\.\d+\.\d+\.\d+|[0-9a-f:]*)
Value Required AS (\d+\**)
Value PKTRCVD (\d+\**)
Value PKTSENT (\d+\**)
Value INQ (\d+\**)
Value OUTQ (\d+\**)
Value UP_DOWN (\S+)
Value STATE (\w+)
Value RCV (\d+)
Value ACT (\d+)
Value SENT (\d+)

Start
  ^----------- -> Neighbor

Neighbor
  ^${NEIGHBOR}(\s|$$) -> NeighborData

NeighborData
  ^\s+${AS}\s+${PKTRCVD}\s+${PKTSENT}\s+${INQ}\s+${OUTQ}\s+${UP_DOWN}\s+(${STATE}|${RCV}/${ACT}/${SENT})(\s|$$) -> Record Neighbor
  ^.*indicates
  ^-----------
  ^\s*$$
  ^. -> Error
`
