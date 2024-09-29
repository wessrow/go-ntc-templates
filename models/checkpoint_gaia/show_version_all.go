package checkpoint_gaia 

type ShowVersionAll struct {
	Version	string	`json:"VERSION"`
	Build	string	`json:"BUILD"`
	Kernel	string	`json:"KERNEL"`
	Architecture	string	`json:"ARCHITECTURE"`
}

var ShowVersionAll_Template = `Value VERSION (\S+)
Value BUILD (\S+)
Value KERNEL (\S+)
Value ARCHITECTURE (\S+)

Start
  ^Product version Check Point Gaia\s${VERSION}
  ^OS build\s${BUILD}
  ^OS kernel version\s${KERNEL}
  ^OS edition\s${ARCHITECTURE}-bit -> Record
`