package cisco_ios

type ShowIpEigrpTopology struct {
	Prefix_length string   `json:"PREFIX_LENGTH"`
	Successors    string   `json:"SUCCESSORS"`
	Adv_fd        []string `json:"ADV_FD"`
	Adv_rd        []string `json:"ADV_RD"`
	Process_id    string   `json:"PROCESS_ID"`
	Network       string   `json:"NETWORK"`
	Fd            string   `json:"FD"`
	Tag           string   `json:"TAG"`
	Adv_router    []string `json:"ADV_ROUTER"`
	Out_interface []string `json:"OUT_INTERFACE"`
	Source        string   `json:"SOURCE"`
	Router_id     string   `json:"ROUTER_ID"`
	Code          string   `json:"CODE"`
}

var ShowIpEigrpTopology_Template string = `Value Filldown PROCESS_ID (\d+)
Value Filldown ROUTER_ID (\d+\.\d+\.\d+\.\d+)
Value Required CODE (\S+)
Value NETWORK (\d+\.\d+\.\d+\.\d+)
Value PREFIX_LENGTH (\d+)
Value SUCCESSORS (\d+)
Value FD (\d+|Inaccessible)
Value TAG (\d+)
Value List ADV_ROUTER (\d+\.\d+\.\d+\.\d+|\w+)
Value List ADV_FD (\d+)
Value List ADV_RD (\d+)
Value List OUT_INTERFACE (\S+)
Value SOURCE (R\S+)

Start
  # Captures Process ID and Router ID
  ^.+AS\(${PROCESS_ID}\)/ID\(${ROUTER_ID}\)
  # Skips over the code line that explains what each code means
  ^Codes:
  # Skips over the definitions for the codes
  ^\s*\S\s-\s\S+
  # Matches a route and captures if ${TAG} is use for the route and then moves to Gateway section
  ^${CODE}\s+${NETWORK}/${PREFIX_LENGTH},\s+${SUCCESSORS}\s+successors,\s+FD\s+is\s+${FD},\s+tag\s+is\s+${TAG} -> Gateway
  # Matches a route and captures it and then moves to Gateway section
  ^${CODE}\s+${NETWORK}/${PREFIX_LENGTH},\s+${SUCCESSORS}\s+successors,\s+FD\s+is\s+${FD} -> Gateway
  # If it doesn't match anything above, this just acknowledges a new line
  ^\s*$$
  # This will throw an error if there are no matches
  ^. -> Error

Gateway
  # This captures adv router, FD to router, and Reported Distance(RD)
  ^\s*via\s+${ADV_ROUTER}\s+\(${ADV_FD}/${ADV_RD}\),\s+${OUT_INTERFACE}
  # This captures the advertising router and outgoing interface
  ^\s*via\s+${ADV_ROUTER},\s+${OUT_INTERFACE}
  # This captures the scenarion where the route is injected via Redistribution.
  ^\s*via\s+${SOURCE}
  # This will not capture anything but if it encounters another route, it will continue and record what it already captured
  ^\S+\s+(?:\d+(?:\.|)){4}/\d+,\s+\d+\s+successors -> Continue.Record
  # These are the same as above and capture the next set of routes
  ^${CODE}\s+${NETWORK}/${PREFIX_LENGTH},\s+${SUCCESSORS}\s+successors,\s+FD\s+is\s+${FD},\s+tag\s+is\s+${TAG}
  ^${CODE}\s+${NETWORK}/${PREFIX_LENGTH},\s+${SUCCESSORS}\s+successors,\s+FD\s+is\s+${FD}
  # If it encounters another AS/Router ID it will continue and record what it just captured
  ^.+AS\(\d+\)/ID -> Continue.Record
  # If it encounters it again, it will continue, but clear all captured data other than Filldown Values
  ^.+AS\(\d+\)/ID -> Continue.Clearall
  # It will start the process over again if it encounters a new process ID / router id
  ^.+AS\(${PROCESS_ID}\)/ID\(${ROUTER_ID}\)
  ^Codes: -> Start
  # If it doesn't match anything above, this just acknowledges a new line
  ^\s*$$
  # This will throw an error if there are no matches
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^. -> Error

`
