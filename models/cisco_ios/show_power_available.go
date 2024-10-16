package cisco_ios

type ShowPowerAvailable struct {
	Total_used_watts      string `json:"TOTAL_USED_WATTS"`
	Max_avail_watts       string `json:"MAX_AVAIL_WATTS"`
	Sys_used_watts        string `json:"SYS_USED_WATTS"`
	Sys_avail_watts       string `json:"SYS_AVAIL_WATTS"`
	Inline_used_watts     string `json:"INLINE_USED_WATTS"`
	Inline_avail_watts    string `json:"INLINE_AVAIL_WATTS"`
	Backplane_used_watts  string `json:"BACKPLANE_USED_WATTS"`
	Backplane_avail_watts string `json:"BACKPLANE_AVAIL_WATTS"`
}

var ShowPowerAvailable_Template string = `####
# For 4500 switches
###
Value SYS_USED_WATTS (\d+)
Value SYS_AVAIL_WATTS (\d+)
Value INLINE_USED_WATTS (\d+)
Value INLINE_AVAIL_WATTS (\d+)
Value BACKPLANE_USED_WATTS (\d+)
Value BACKPLANE_AVAIL_WATTS (\d+)
Value TOTAL_USED_WATTS (\d+)
Value MAX_AVAIL_WATTS (\d+)

Start
  ^Power.*Maximum -> SUMMARY
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is

SUMMARY
  ^System Power \(12V\)\s+${SYS_USED_WATTS}\s+${SYS_AVAIL_WATTS} -> Continue
  ^Inline Power \(-50V\)\s+${INLINE_USED_WATTS}\s+${INLINE_AVAIL_WATTS} -> Continue
  ^Backplane Power \(3.3V\)\s+${BACKPLANE_USED_WATTS}\s+${BACKPLANE_AVAIL_WATTS} -> Continue
  # Capture output that uses a phrase like not to exceed Total Maximum Available
  ^Total\s+${TOTAL_USED_WATTS}.*Available\s=\s${MAX_AVAIL_WATTS}
  # Capture output that only displays digits
  ^Total\s+${TOTAL_USED_WATTS}\s+${MAX_AVAIL_WATTS}

`
