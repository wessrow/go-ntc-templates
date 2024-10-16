package cisco_xr

type ShowInstallActive struct {
	Boot_partition string   `json:"BOOT_PARTITION"`
	Boot_device    string   `json:"BOOT_DEVICE"`
	Boot_image     string   `json:"BOOT_IMAGE"`
	Packages       []string `json:"PACKAGES"`
	Node           string   `json:"NODE"`
	Node_type      string   `json:"NODE_TYPE"`
}

var ShowInstallActive_Template string = `Value Required NODE (\S+)
Value NODE_TYPE (\w+)
Value BOOT_PARTITION (\S+)
Value BOOT_DEVICE (\S+)
Value BOOT_IMAGE (\S+)
Value List PACKAGES (\S+)

Start
  # Match the timestamp at beginning of command output
  ^\S+\s+\S+\s+\d+\s+\d+:\d+:\d+\.\d+\s+\S+\s*$$
  ^Label\s+:
  ^Secure\s+Domain\s+Router:
  ^\s*Node\s+${NODE}\s+\[${NODE_TYPE}\] -> Node
  ^\s*$$
  ^. -> Error "LINE NOT FOUND"

Node
  ^\s+Boot\s+Partition:\s+${BOOT_PARTITION}
  ^\s+Boot\s+Device:\s+${BOOT_DEVICE}
  ^\s+Boot\s+Image:\s+${BOOT_IMAGE}
  ^\s+Active\s+Packages: -> Packages
  ^. -> Error "LINE NOT FOUND"

Packages
  ^\s+${PACKAGES}
  ^\s*$$ -> Record Start
  ^. -> Error "LINE NOT FOUND"
`
