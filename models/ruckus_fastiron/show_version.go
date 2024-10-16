package ruckus_fastiron

type ShowVersion struct {
	Version  []string `json:"VERSION"`
	Bootcode []string `json:"BOOTCODE"`
	Hardware string   `json:"HARDWARE"`
	Serial   []string `json:"SERIAL"`
	Uptime   []string `json:"UPTIME"`
}

var ShowVersion_Template string = `Value List VERSION (\S+)
Value List BOOTCODE (\S+\s+\S+)
Value HARDWARE (.*)
Value List SERIAL (\S+)
Value List UPTIME (.*)


Start
  ^\s+Copyright\s+\(c\)\s+Ruckus\s+Networks,\s+Inc\.\s+All\s+rights\s+reserved\.
  ^\s+UNIT\s+[0-9]*:\s+compiled\s+on\s+\S+\s+[0-9]*\s+[0-9]*\s+at\s+[0-9]*:[0-9]*:[0-9]*\s+labeled\s+as\s+\S+
  ^\s+\([0-9]*\s+bytes\)\s+from\s+Primary\s+\S+\s+\(UFI\)
  ^\s+SW:\s+Version\s+${VERSION}
  ^\s+Compressed\s+Primary\s+Boot\s+Code\s+size\s+=\s+[0-9]*,\s+Version:${BOOTCODE}
  ^\s+Compiled\s+on\s+\S+\s+\S+\s+[0-9]*\s+[0-9]*:[0-9]*:[0-9]*\s+[0-9]*
  ^\s+HW:\s+${HARDWARE}\s*$$
  ^UNIT\s+[0-9]*:\s+SL\s+[0-9]*:\s+.*
  ^\s+Serial\s+#:${SERIAL}
  ^\s+Software\s+Package:\s+\S+
  ^\s+Current\s+License:\s+\S+
  ^\s+P-ASIC\s+[0-9]*:\s+.*
  ^\s+[0-9]* [KMGT](Hz|B)\s+.*
  ^STACKID\s+[0-9]*\s+system\s+uptime\s+is\s+${UPTIME}\s*$$
  ^The\s+system\s+started\s+at\s+.*
  ^The\s+system\s+:\s+started=.*
  ^My\s+stack\s+unit\s+ID\s+=\s+[0-9]*,\s+bootup\s+role\s+=\s+\S+ -> Record
  ^=+
  ^\s*$$
  ^. -> Error
`
