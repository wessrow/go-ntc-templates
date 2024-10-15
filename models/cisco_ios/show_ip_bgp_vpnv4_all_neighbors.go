package cisco_ios

type ShowIpBgpVpnv4AllNeighbors struct {
	Local_as          string `json:"LOCAL_AS"`
	Peer_group        string `json:"PEER_GROUP"`
	Bgp_state         string `json:"BGP_STATE"`
	Localhost_ip      string `json:"LOCALHOST_IP"`
	Remote_ip         string `json:"REMOTE_IP"`
	Remote_port       string `json:"REMOTE_PORT"`
	Outbound_routemap string `json:"OUTBOUND_ROUTEMAP"`
	Neighbor          string `json:"NEIGHBOR"`
	Vrf               string `json:"VRF"`
	Remote_as         string `json:"REMOTE_AS"`
	Remote_router_id  string `json:"REMOTE_ROUTER_ID"`
	Localhost_port    string `json:"LOCALHOST_PORT"`
	Inbound_routemap  string `json:"INBOUND_ROUTEMAP"`
}

var ShowIpBgpVpnv4AllNeighbors_Template string = `Value NEIGHBOR (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value VRF (\S+)
Value REMOTE_AS (\d+)
Value LOCAL_AS (\d+)
Value PEER_GROUP (\S+)
Value REMOTE_ROUTER_ID (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value BGP_STATE (\w+)
Value LOCALHOST_IP (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value LOCALHOST_PORT (\d+)
Value REMOTE_IP (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value REMOTE_PORT (\d+)
Value INBOUND_ROUTEMAP (\S+)
Value OUTBOUND_ROUTEMAP (\S+)

Start
  # Capture first line, which shows the BGP neighor and remote AS number
  ^BGP\s+neighbor\s+is -> Continue.Record
  ^BGP\s+neighbor\s+is\s+${NEIGHBOR},\s+vrf\s+${VRF},\s+remote\s+AS\s+${REMOTE_AS},\s+local\s+AS\s+${LOCAL_AS}
  # Capture BGP peer group
  # Example: 'Member of peer-group RR_SERVERS for session parameters'
  ^\s*Member\s+of\s+peer-group\s+${PEER_GROUP}
  #
  # Capture remote router ID
  # Example: ' BGP version 4, remote router ID 10.10.255.14'
  ^.+remote\s+router\s+ID\s+${REMOTE_ROUTER_ID}
  #
  # Capture BGP state
  # Example: ' BGP state = Established, up for 7w3d'
  ^\s+BGP\s+state\s+=\s+${BGP_STATE}
  #
  # Capture Inbound/Outbound Route Maps
  # Example: 'Route map for incoming advertisements is BGP_Vendor_in'
  # Example: 'Route map for outgoing advertisements is BGP_Vendor_out'
  ^\s+Route\s+map\s+for\s+incoming\s+advertisements\s+is\s+${INBOUND_ROUTEMAP}
  ^\s+Route\s+map\s+for\s+outgoing\s+advertisements\s+is\s+${OUTBOUND_ROUTEMAP}
  #
  # Match local host and port
  # Example: 'Local host: 10.10.255.13, Local port: 39443'
  ^Local\s+host:\s+${LOCALHOST_IP},\s+Local\s+port:\s+${LOCALHOST_PORT}
  #
  # Match foreign host and port
  # Example: 'Foreign host: 10.10.255.14, Local port: 39443'
  ^Foreign\s+host:\s+${REMOTE_IP},\s+Foreign\s+port:\s+${REMOTE_PORT}
  #^.+ -> Error`
