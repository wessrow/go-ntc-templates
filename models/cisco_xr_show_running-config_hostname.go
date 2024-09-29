package models

type CiscoXrShowRunningConfigHostname struct {
	Hostname	string	`json:"HOSTNAME"`
}

var CiscoXrShowRunningConfigHostname_Template = `Value HOSTNAME (\S+)

Start
  # Capture date output by XR by 'show run hostname' first
  ^\S+ \S+ \S+ .*$$
  # Capture hostname
  ^hostname ${HOSTNAME}$$
  ^\s*$$
  ^. -> Error

`