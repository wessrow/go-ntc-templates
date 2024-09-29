package vmware_nsxv 

type ShowIpBgpNeighbors struct {
	Neighbor_id	string	`json:"NEIGHBOR_ID"`
	Remote_as	string	`json:"REMOTE_AS"`
	State	string	`json:"STATE"`
	Hold_interval	string	`json:"HOLD_INTERVAL"`
	Keepalive_interval	string	`json:"KEEPALIVE_INTERVAL"`
	Pfx_recv	string	`json:"PFX_RECV"`
	Pfx_sent	string	`json:"PFX_SENT"`
	Pfx_adv	string	`json:"PFX_ADV"`
}

var ShowIpBgpNeighbors_Template = `Value NEIGHBOR_ID (\d+(\.\d+){3})
Value REMOTE_AS (\d+)
Value STATE (\D.*)
Value HOLD_INTERVAL (\d+)
Value KEEPALIVE_INTERVAL (\d+)
Value PFX_RECV (\d+)
Value PFX_SENT (\d+)
Value PFX_ADV (\d+)

Start
	^BGP neighbor.* -> Continue.Record
	^BGP neighbor is ${NEIGHBOR_ID},\s+remote AS ${REMOTE_AS}
	^BGP state = ${STATE}
	^Hold time is ${HOLD_INTERVAL}, Keep alive interval is ${KEEPALIVE_INTERVAL}
	^\s+Prefixes received ${PFX_RECV} sent ${PFX_SENT} advertised ${PFX_ADV}`