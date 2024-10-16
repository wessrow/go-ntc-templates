package cisco_ios

type ShowWirelessTagPolicySummary struct {
	Description string `json:"DESCRIPTION"`
	Name        string `json:"NAME"`
}

var ShowWirelessTagPolicySummary_Template string = `Value Required NAME (\S+)
Value DESCRIPTION (.+?)

Start
  # Skip the header lines
  ^Number of Policy Tags:.*
  ^Policy Tag Name\s+Description\s*$$
  ^-+
  # Capturing output
  ^${NAME}\s+${DESCRIPTION}\s*$$ -> Record
  # Handle cases where the description is missing
  ^${NAME} -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  # Match blank lines
  ^\s*$$
  # Error out if raw data does not match any above rules.
  ^. -> Error
`
