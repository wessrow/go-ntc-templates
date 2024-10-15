package cisco_xr

type ShowRunningConfigHostname struct {
	Hostname string `json:"HOSTNAME"`
}

var ShowRunningConfigHostname_Template string = `Value HOSTNAME (\S+)

Start
  # Capture date output by XR by 'show run hostname' first
  ^\S+ \S+ \S+ .*$$
  # Capture hostname
  ^hostname ${HOSTNAME}$$
  ^\s*$$
  ^. -> Error
`
