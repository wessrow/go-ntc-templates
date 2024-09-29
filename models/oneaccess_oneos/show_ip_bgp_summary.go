package oneaccess_oneos 

type ShowIpBgpSummary struct {
	Router_id	string	`json:"ROUTER_ID"`
	Local_as	string	`json:"LOCAL_AS"`
	Vrf	string	`json:"VRF"`
	Rib_entries	string	`json:"RIB_ENTRIES"`
	Rib_memory	string	`json:"RIB_MEMORY"`
	Peers	string	`json:"PEERS"`
	Bgp_memory	string	`json:"BGP_MEMORY"`
	As_path_entries	string	`json:"AS_PATH_ENTRIES"`
	Communtiy_entries	string	`json:"COMMUNTIY_ENTRIES"`
	Bgp_neighbor	string	`json:"BGP_NEIGHBOR"`
	Address_family	string	`json:"ADDRESS_FAMILY"`
	Neighbor_as	string	`json:"NEIGHBOR_AS"`
	Message_received	string	`json:"MESSAGE_RECEIVED"`
	Message_sent	string	`json:"MESSAGE_SENT"`
	Table_version	string	`json:"TABLE_VERSION"`
	In_queue	string	`json:"IN_QUEUE"`
	Out_queue	string	`json:"OUT_QUEUE"`
	Up_down	string	`json:"UP_DOWN"`
	State_prefix_received	string	`json:"STATE_PREFIX_RECEIVED"`
	Total_neighbors	string	`json:"TOTAL_NEIGHBORS"`
}

var ShowIpBgpSummary_Template = `Value Filldown,Required ROUTER_ID ([0-9a-f:\.]+)
Value Filldown LOCAL_AS (\d+(\.\d+)?)
Value Filldown VRF (\w+)
Value Filldown RIB_ENTRIES (\d+)
Value Filldown RIB_MEMORY (\d+\s\S+)
Value Filldown PEERS (\d+)
Value Filldown BGP_MEMORY (\d+\s\S+)
Value Filldown AS_PATH_ENTRIES (\d+)
Value Filldown COMMUNTIY_ENTRIES (\d+)
Value BGP_NEIGHBOR (\d+\.\d+\.\d+\.\d+)
Value ADDRESS_FAMILY (\d+)
Value NEIGHBOR_AS (\d+)
Value MESSAGE_RECEIVED (\d+)
Value MESSAGE_SENT (\d+)
Value TABLE_VERSION (\d+)
Value IN_QUEUE (\d+)
Value OUT_QUEUE (\d+)
Value UP_DOWN (\S+?)
Value STATE_PREFIX_RECEIVED (\S+\s+\S+|\S+)
Value Fillup TOTAL_NEIGHBORS (\d+)

Start
  ^BGP\srouter\sidentifier\s${ROUTER_ID},\slocal\sAS\snumber\s${LOCAL_AS},?\s+vrf\s\(?${VRF}\)?
  ^RIB\sentries\s${RIB_ENTRIES},\susing\s${RIB_MEMORY}\sof\smemory
  ^Peers\s${PEERS},\susing\s${BGP_MEMORY}\sof\smemory
  ^${AS_PATH_ENTRIES}\sBGP\sAS-PATH\sentries
  ^${COMMUNTIY_ENTRIES}\sBGP\scommunity\sentries
  ^Neighbor\s+V\s+AS\s+MsgRcvd\s+MsgSent\s+TblVer\s+InQ\s+OutQ\s+Up/Down\s+State/PfxRcd\s*$$
  ^${BGP_NEIGHBOR}\s+${ADDRESS_FAMILY}\s+${NEIGHBOR_AS}\s+${MESSAGE_RECEIVED}\s+${MESSAGE_SENT}\s+${TABLE_VERSION}\s+${IN_QUEUE}\s+${OUT_QUEUE}\s+${UP_DOWN}\s+${STATE_PREFIX_RECEIVED} -> Record
  ^Total\snumber\sof\sneighbors\s${TOTAL_NEIGHBORS}
  ^\s*$$
  ^. -> Error

EOF
`