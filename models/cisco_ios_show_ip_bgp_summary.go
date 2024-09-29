package models

type CiscoIosShowIpBgpSummary struct {
	Router_id	string	`json:"ROUTER_ID"`
	Local_as	string	`json:"LOCAL_AS"`
	Address_family	string	`json:"ADDRESS_FAMILY"`
	Bgp_neighbor	string	`json:"BGP_NEIGHBOR"`
	Bgp_version	string	`json:"BGP_VERSION"`
	Neighbor_as	string	`json:"NEIGHBOR_AS"`
	Messages_received	string	`json:"MESSAGES_RECEIVED"`
	Messages_sent	string	`json:"MESSAGES_SENT"`
	Table_version	string	`json:"TABLE_VERSION"`
	Input_queue	string	`json:"INPUT_QUEUE"`
	Output_queue	string	`json:"OUTPUT_QUEUE"`
	Up_down	string	`json:"UP_DOWN"`
	State_or_prefixes_received	string	`json:"STATE_OR_PREFIXES_RECEIVED"`
}

var CiscoIosShowIpBgpSummary_Template = `Value Filldown,Required ROUTER_ID ([0-9a-f:\.]+)
Value Filldown LOCAL_AS (\d+(\.\d+)?)
Value Filldown ADDRESS_FAMILY (.+?)
Value BGP_NEIGHBOR (\d+?\.\d+?\.\d+?\.\d+?)
Value BGP_VERSION (\d+)
Value NEIGHBOR_AS (\d+)
Value MESSAGES_RECEIVED (\d+)
Value MESSAGES_SENT (\d+)
Value TABLE_VERSION (\d+)
Value INPUT_QUEUE (\d+)
Value OUTPUT_QUEUE (\d+)
Value UP_DOWN (\S+?)
Value STATE_OR_PREFIXES_RECEIVED (\S+?\s+\S+?|\S+?)

Start
  ^For\s+address\s+family:\s+${ADDRESS_FAMILY}$$
  ^BGP\s+router\s+identifier\s+${ROUTER_ID},\s+local\s+AS\s+number\s+${LOCAL_AS}\s*$$
  ^${BGP_NEIGHBOR}\s+${BGP_VERSION}\s+${NEIGHBOR_AS}\s+${MESSAGES_RECEIVED}\s+${MESSAGES_SENT}\s+${TABLE_VERSION}\s+${INPUT_QUEUE}\s+${OUTPUT_QUEUE}\s+${UP_DOWN}\s+${STATE_OR_PREFIXES_RECEIVED}\s*$$ -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is

EOF
`