package cisco_ios

type ShowInterfaceLink struct {
	Since    string `json:"SINCE"`
	Uptime   string `json:"UPTIME"`
	Port     string `json:"PORT"`
	Name     string `json:"NAME"`
	Downtime string `json:"DOWNTIME"`
}

var ShowInterfaceLink_Template string = `Value PORT (\S+)
Value NAME (.*?)
Value DOWNTIME ((00:00:00)|(([0-9]{1,2} year[s]? , )?([0-9]{1,2} week[s]?, )?([0-9]{1,2} day[s]?, )?([0-9]{1,2} hour[s]?, )?([0-9]{1,2} minute[s]? )?([0-9]{1,2} secs))|((([0-9]{1,2}y)?([0-9]{1,2}w)?([0-9]{1,2}d)([0-9]{1,2}h)?|[0-9]{1,2}:[0-9]{1,2}:[0-9]{1,2})))
Value SINCE (.+?)
Value UPTIME ((([0-9]{1,2}m)?([0-9]{1,2}w)?([0-9]{1,2}d)([0-9]{1,2}h)?|[0-9]{1,2}:[0-9]{1,2}:[0-9]{1,2}))

Start
  ^\s*\^\s* -> EOF
  ^.*Invalid input detected.* -> EOF
  ^Load\s+for\s+
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Time\s+source\s+is
  ^-+\s*$$
  ^\s*Port\s+Name\s+Down Time\s+Down Since\s*$$ -> DownSince
  ^\s*Port\s+Name\s+Down Time\s+Up Time\s*$$ -> UpTime
  ^\s*$$
  ^. -> Error

DownSince
  ^\s*${PORT}\s+${DOWNTIME}\s*$$ -> Record
  ^\s*${PORT}\s+${DOWNTIME}\s*${SINCE}\s*$$ -> Record
  ^\s*${PORT}\s+${NAME}\s*${DOWNTIME}\s*$$ -> Record
  ^\s*${PORT}\s+${NAME}\s*${DOWNTIME}\s+${SINCE}\s*$$ -> Record
  ^\s*$$
  ^. -> Error

UpTime
  ^\s*${PORT}\s+${NAME}\s+${DOWNTIME}\s+${UPTIME}?\s*$$ -> Record
  ^\s*${PORT}\s+${DOWNTIME}\s+${UPTIME}?\s*$$ -> Record
  ^\s*${PORT}\s+${NAME}\s+${DOWNTIME}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
