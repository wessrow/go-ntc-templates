package mikrotik_routeros

type RoutingOspfNeighborPrintTerseWithoutPaging struct {
	Router_id         string `json:"ROUTER_ID"`
	Dr_address        string `json:"DR_ADDRESS"`
	State             string `json:"STATE"`
	State_changes     string `json:"STATE_CHANGES"`
	Ls_retransmits    string `json:"LS_RETRANSMITS"`
	Adjacency         string `json:"ADJACENCY"`
	Index             string `json:"INDEX"`
	Address           string `json:"ADDRESS"`
	Interface         string `json:"INTERFACE"`
	Backup_dr_address string `json:"BACKUP_DR_ADDRESS"`
	Db_summaries      string `json:"DB_SUMMARIES"`
	Instance          string `json:"INSTANCE"`
	Priority          string `json:"PRIORITY"`
	Ls_requests       string `json:"LS_REQUESTS"`
}

var RoutingOspfNeighborPrintTerseWithoutPaging_Template string = `Value INDEX (\d+)
Value INSTANCE (\S+)
Value ROUTER_ID (\S+)
Value ADDRESS (\S+)
Value INTERFACE (\S+)
Value PRIORITY (\d+)
Value DR_ADDRESS (\S+)
Value BACKUP_DR_ADDRESS (\S+)
Value STATE (\S+)
Value STATE_CHANGES (\d+)
Value LS_RETRANSMITS (\d+)
Value LS_REQUESTS (\d+)
Value DB_SUMMARIES (\d+)
Value ADJACENCY (\S+)

Start
  ^\s*${INDEX}\s+instance=${INSTANCE}\s+router-id=${ROUTER_ID}\s+address=${ADDRESS}\s+interface=${INTERFACE}\s+priority=${PRIORITY}(\s+dr-address=${DR_ADDRESS})?(\s+backup-dr-address=${BACKUP_DR_ADDRESS})?(\s+state=${STATE})?(\s+state-changes=${STATE_CHANGES})?(\s+ls-retransmits=${LS_RETRANSMITS})?(\s+ls-requests=${LS_REQUESTS})?(\s+db-summaries=${DB_SUMMARIES})?(\s+adjacency=${ADJACENCY})?\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
