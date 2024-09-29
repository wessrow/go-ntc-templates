package models

type CiscoWlcSshShowRedundancyDetail struct {
	Redundancy_mgmt_addr	string	`json:"REDUNDANCY_MGMT_ADDR"`
	Peer_redundancy_mgmt_addr	string	`json:"PEER_REDUNDANCY_MGMT_ADDR"`
	Redundancy_port_addr	string	`json:"REDUNDANCY_PORT_ADDR"`
	Peer_redundancy_port_addr	string	`json:"PEER_REDUNDANCY_PORT_ADDR"`
	Peer_service_port_addr	string	`json:"PEER_SERVICE_PORT_ADDR"`
	Keep_alive_timeout	string	`json:"KEEP_ALIVE_TIMEOUT"`
	Peer_search_timeout	string	`json:"PEER_SEARCH_TIMEOUT"`
}

var CiscoWlcSshShowRedundancyDetail_Template = `Value REDUNDANCY_MGMT_ADDR (\S+)
Value PEER_REDUNDANCY_MGMT_ADDR (\S+)
Value REDUNDANCY_PORT_ADDR (\S+)
Value PEER_REDUNDANCY_PORT_ADDR (\S+)
Value PEER_SERVICE_PORT_ADDR (\S+)
Value KEEP_ALIVE_TIMEOUT (\d+)
Value PEER_SEARCH_TIMEOUT (\d+)

Start
  ^\s*Redundancy Management IP Address\.*\s+${REDUNDANCY_MGMT_ADDR}\s*$$
  ^\s*Peer Redundancy Management IP Address\.*\s+${PEER_REDUNDANCY_MGMT_ADDR}\s*$$
  ^\s*Redundancy Port IP Address\.*\s+${REDUNDANCY_PORT_ADDR}\s*$$
  ^\s*Peer Redundancy Port IP Address\.*\s+${PEER_REDUNDANCY_PORT_ADDR}\s*$$
  ^\s*Peer Service Port IP Address\.*\s+${PEER_SERVICE_PORT_ADDR}\s*$$
  # presently not parsing Switchover History
  ^\s*Keep Alive Timeout\s+:\s+${KEEP_ALIVE_TIMEOUT}\s+msecs\s*$$
  ^\s*Peer Search Timeout\s+:\s+${PEER_SEARCH_TIMEOUT}\s+secs\s*$$
  # presently not parsing Peer Network Routes

`