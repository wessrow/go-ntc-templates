package models

type MikrotikRouterosSystemClockPrint struct {
	Time	string	`json:"TIME"`
	Date	string	`json:"DATE"`
	Time_zone_autodetect	string	`json:"TIME_ZONE_AUTODETECT"`
	Time_zone_name	string	`json:"TIME_ZONE_NAME"`
	Gmt_offset	string	`json:"GMT_OFFSET"`
	Dst_active	string	`json:"DST_ACTIVE"`
}

var MikrotikRouterosSystemClockPrint_Template = `Value TIME (\d{2}\:\d{2}\:\d{2})
Value DATE ([a-z]{3}\/\d{2}\/\d{4})
Value TIME_ZONE_AUTODETECT (yes|no)
Value TIME_ZONE_NAME (\S+)
Value GMT_OFFSET (\+\d{2}\:\d{2})
Value DST_ACTIVE (yes|no)

Start
  ^\s*time:\s+${TIME}\s*$$
  ^\s*date:\s+${DATE}\s*$$
  ^\s*time-zone-autodetect:\s+${TIME_ZONE_AUTODETECT}\s*$$
  ^\s*time-zone-name:\s+${TIME_ZONE_NAME}\s*$$
  ^\s*gmt-offset:\s+${GMT_OFFSET}\s*$$
  ^\s*dst-active:\s+${DST_ACTIVE}\s*$$ -> Record
  ^\s*$$
  ^. -> Error

`