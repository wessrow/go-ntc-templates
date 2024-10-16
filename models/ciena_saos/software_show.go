package ciena_saos

type SoftwareShow struct {
	Kernel            string `json:"KERNEL"`
	Version_installed string `json:"VERSION_INSTALLED"`
	Version_running   string `json:"VERSION_RUNNING"`
}

var SoftwareShow_Template string = `Value VERSION_INSTALLED (\S+)
Value VERSION_RUNNING (\S+)
Value KERNEL (\S+)

Start
  ^\| *Installed Package +\: +${VERSION_INSTALLED}
  ^\| *Running Package +\: +${VERSION_RUNNING}
  ^\| *Running Kernel +\: +${KERNEL}
  ^\+-+
  ^\|\s+
  ^\|\s+Running\s+Package
  ^\|\s+Application\s+Build
  ^\|\s+Package\s+Build\s+Info
  ^\|\s+Running\s+Kernel
  ^\|\s+Running\s+MIB\s+Version
  ^\|\s+(?:Release|Banks)\s+Status
  ^\|\s+(?:Running|Standby) bank
  ^\|\s+Bank\s+package\s+version
  ^\|\s+Bootloader\s+(?:version|status)
  ^\|\s+Bank\s+package\s+version
  ^\|\s+Last\s+(?:command|configuration)\s+file
  ^\s*$$
  ^. -> Error
`
