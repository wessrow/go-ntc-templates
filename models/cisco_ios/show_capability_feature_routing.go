package cisco_ios

type ShowCapabilityFeatureRouting struct {
	Feature string `json:"FEATURE"`
	State   string `json:"STATE"`
}

var ShowCapabilityFeatureRouting_Template string = `Value FEATURE ((\S+(\s)*)+)
Value STATE (\w+)

Start
  ^Displaying capability information for all available features: 
  ^${FEATURE}:\s${STATE} -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
`
