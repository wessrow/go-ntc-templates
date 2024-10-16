package arista_eos

type ShowIpRoute struct {
	Distance      string   `json:"DISTANCE"`
	Next_hop      []string `json:"NEXT_HOP"`
	Interface     []string `json:"INTERFACE"`
	Direct        string   `json:"DIRECT"`
	Vrf           string   `json:"VRF"`
	Protocol      string   `json:"PROTOCOL"`
	Network       string   `json:"NETWORK"`
	Prefix_length string   `json:"PREFIX_LENGTH"`
	Metric        string   `json:"METRIC"`
}

var ShowIpRoute_Template string = `Value Filldown VRF (\S+)
Value PROTOCOL (\S+\s\S+?|\w?)
Value Required NETWORK (\d+\.\d+\.\d+\.\d+)
Value PREFIX_LENGTH (\d+)
Value DISTANCE (\d+)
Value METRIC (\d+)
Value DIRECT (directly)
Value Required,List NEXT_HOP (connected|\d+\.\d+\.\d+\.\d+)
#Value INTERFACE (\S+)
Value List INTERFACE (.+)

Start
  ^VRF(\s+name)?:\s+${VRF}\s*$$
  ^WARNING
  ^kernel
  ^Codes:
  # Match for codes
  ^\s+.+-.+
  ^\s*$$ -> Routes
  # Ignore IP routing not enabled
  ^\! IP routing not enabled
  # Error on everything else
  ^. -> Error

Routes
  ^\s+(\S+\s\S+?|\w?)\s+(\d+\.\d+\.\d+\.\d+)/(\d+)\s -> Continue.Record
  ^\s+${PROTOCOL}\s+${NETWORK}/${PREFIX_LENGTH}\s+is\s+${DIRECT}\s+${NEXT_HOP},\s+${INTERFACE}$$
  ^\s+${PROTOCOL}\s+${NETWORK}/${PREFIX_LENGTH}\s+(?:\[${DISTANCE}/${METRIC}\]|is\s+${DIRECT})(?:.+?)${NEXT_HOP},\s+${INTERFACE}$$
  ^\s+via\s+${NEXT_HOP},\s+${INTERFACE}
  ^\s*$$ -> Record
  ^VRF(\s+name)?:\s+${VRF}\s*$$ -> Start
  ^Gateway\s+of\s+last
  # Ignore IP routing not enabled
  ^\! IP routing not enabled
  # Error on everything else
  ^. -> Error`
