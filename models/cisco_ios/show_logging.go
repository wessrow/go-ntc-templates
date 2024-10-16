package cisco_ios

type ShowLogging struct {
	Severity string   `json:"SEVERITY"`
	Mnemonic string   `json:"MNEMONIC"`
	Time     string   `json:"TIME"`
	Facility string   `json:"FACILITY"`
	Day      string   `json:"DAY"`
	Timezone string   `json:"TIMEZONE"`
	Message  []string `json:"MESSAGE"`
	Number   string   `json:"NUMBER"`
	Month    string   `json:"MONTH"`
}

var ShowLogging_Template string = `Value NUMBER (\d+)
Value MONTH (\S+)
Value DAY (\d{1,2})
Value TIME ((\d+:\d+:\d+\.\d+)|(\d+:\d+:\d+)|(\d{1,2}:\d{1,2}))
Value TIMEZONE (\S{3})
Value FACILITY (\w+)
Value SEVERITY (\d)
Value MNEMONIC (\S+)
Value List MESSAGE (.+)

Start
  ^Log\s+Buffer
  ^(\*)?${NUMBER}:\s+${MONTH}\s+${DAY}\s+${TIME}\s+${TIMEZONE}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> NumberLogs
  # 022701: Jun 19 03:02:31: %LINEPROTO-5-UPDOWN: Line protocol on Interface GigabitEthernet2/0/3, changed state to down
  ^(\*)?${NUMBER}:\s+${MONTH}\s+${DAY}\s+${TIME}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> NumberLogs
  ^(\*)?${NUMBER}:\s+${MONTH}\s+${DAY}\s+${TIME}\s+${TIMEZONE}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> NumberLogs
  # 000024: Dec  2 12:09:21.207: CEF-HWIDB: EDSP0 LES switching vector set to Null
  ^(\*)?${NUMBER}:\s+${MONTH}\s+${DAY}\s+${TIME}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> NumberLogs
  ^(\*)?${MONTH}\s+${DAY}\s+${TIME}\s+${TIMEZONE}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> DateLogs
  # Jan 30 14:11:11.354: %ILPOWER-7-DETECT: Interface Gi4/3: Power Device detected: IEEE PD
  ^(\*)?${MONTH}\s+${DAY}\s+${TIME}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> DateLogs
  ^(\*)?${MONTH}\s+${DAY}\s+${TIME}\s+${TIMEZONE}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> DateLogs
  ^(\*)?${MONTH}\s+${DAY}\s+${TIME}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> DateLogs
  # 7:04: %LINEPROTO-5-UPDOWN: Line protocol on Interface GigabitEthernet2/0/3, changed state to up
  ^(\*)?${TIME}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> TimeLogs
  ^(\*)?${TIME}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> TimeLogs


NumberLogs
  ^(\*)?(\d{6}):\s+(\D\D\D)\s+(\d{1,2})\s+((\d+:\d+:\d+\.\d+)|(\d+:\d+:\d+)) -> Continue.Record
  ^(\*)?(\D\D\D)\s+(\d{1,2})\s+((\d+:\d+:\d+\.\d+)|(\d+:\d+:\d+)) -> Continue.Record
  ^(\*)?\d{1,2}:\d{1,2}: -> Continue.Record
  # NUMBER LOGS
  ^(\*)?${NUMBER}:\s+${MONTH}\s+${DAY}\s+${TIME}\s+${TIMEZONE}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE}$$
  ^(\*)?${NUMBER}:\s+${MONTH}\s+${DAY}\s+${TIME}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE}$$
  # NUMBER LOGS NO SEVERITY
  ^(\*)?${NUMBER}:\s+${MONTH}\s+${DAY}\s+${TIME}\s+${TIMEZONE}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE}$$
  ^(\*)?${NUMBER}:\s+${MONTH}\s+${DAY}\s+${TIME}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE}$$
  # DATE LOGS
  ^(\*)?${MONTH}\s+${DAY}\s+${TIME}\s+${TIMEZONE}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE} -> DateLogs
  ^(\*)?${MONTH}\s+${DAY}\s+${TIME}:\s%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE} -> DateLogs
  # DATE LOGS NO SEVERITY
  ^(\*)?${MONTH}\s+${DAY}\s+${TIME}\s+${TIMEZONE}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE} -> DateLogs
  ^(\*)?${MONTH}\s+${DAY}\s+${TIME}:\s${FACILITY}-${MNEMONIC}:\s+${MESSAGE} -> DateLogs
  # TIME LOGS
  ^(\*)?${TIME}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> TimeLogs
  # TIME LOGS NO SEVERITY
  ^(\*)?${TIME}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> TimeLogs
  ^${MESSAGE}$$
  ^\s*$$
  ^. -> Error

DateLogs
  ^(\*)?(\D\D\D)\s+(\d{1,2})\s+((\d+:\d+:\d+\.\d+)|(\d+:\d+:\d+)) -> Continue.Record
  ^(\*)?(\d{6}):\s+(\D\D\D)\s+(\d{1,2})\s+((\d+:\d+:\d+\.\d+)|(\d+:\d+:\d+)) -> Continue.Record
  ^(\*)?\d{1,2}:\d{1,2}: -> Continue.Record
  # DATE LOGS
  ^(\*)?${MONTH}\s+${DAY}\s+${TIME}\s+${TIMEZONE}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE}
  ^(\*)?${MONTH}\s+${DAY}\s+${TIME}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE}
  # DATE LOGS NO SEVERITY
  ^(\*)?${MONTH}\s+${DAY}\s+${TIME}\s+${TIMEZONE}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE}
  ^(\*)?${MONTH}\s+${DAY}\s+${TIME}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE}
  # NUMBER LOGS
  ^(\*)?${NUMBER}:\s+${MONTH}\s+${DAY}\s+${TIME}\s+${TIMEZONE}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> NumberLogs
  ^(\*)?${NUMBER}:\s+${MONTH}\s+${DAY}\s+${TIME}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> NumberLogs
  # NUMBER LOGS NO SEVERITY
  ^(\*)?${NUMBER}:\s+${MONTH}\s+${DAY}\s+${TIME}\s+${TIMEZONE}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> NumberLogs
  ^(\*)?${NUMBER}:\s+${MONTH}\s+${DAY}\s+${TIME}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> NumberLogs
  # TIME LOGS
  ^(\*)?${TIME}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> TimeLogs
  # TIME LOGS NO SEVERITY
  ^(\*)?${TIME}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> TimeLogs
  ^${MESSAGE}$$
  ^\s*$$
  ^. -> Error

TimeLogs
  ^(\*)?\d{1,2}:\d{1,2}: -> Continue.Record
  ^(\*)?(\D\D\D)\s+(\d{1,2})\s+((\d+:\d+:\d+\.\d+)|(\d+:\d+:\d+)) -> Continue.Record
  ^(\*)?\d+:\s+\S+ -> Continue.Record
  # TIME LOGS
  ^(\*)?${TIME}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE}$$
  # TIME LOGS NO SEVERITY
  ^(\*)?${TIME}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE}$$
  # DATE LOGS
  ^(\*)?${MONTH}\s+${DAY}\s+${TIME}\s+${TIMEZONE}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE} -> DateLogs
  ^(\*)?${MONTH}\s+${DAY}\s+${TIME}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE} -> DateLogs
  # DATE LOGS NO SEVERITY
  ^(\*)?${MONTH}\s+${DAY}\s+${TIME}\s+${TIMEZONE}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE} -> DateLogs
  ^(\*)?${MONTH}\s+${DAY}\s+${TIME}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE} -> DateLogs
  # NUMBER LOGS
  ^(\*)?${NUMBER}:\s+${MONTH}\s+${DAY}\s+${TIME}\s+${TIMEZONE}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> NumberLogs
  ^(\*)?${NUMBER}:\s+${MONTH}\s+${DAY}\s+${TIME}:\s+%${FACILITY}-${SEVERITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> NumberLogs
  # NUMBER LOGS NO SEVERITY
  ^(\*)?${NUMBER}:\s+${MONTH}\s+${DAY}\s+${TIME}\s+${TIMEZONE}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> NumberLogs
  ^(\*)?${NUMBER}:\s+${MONTH}\s+${DAY}\s+${TIME}:\s+${FACILITY}-${MNEMONIC}:\s+${MESSAGE}$$ -> NumberLogs
  ^${MESSAGE}$$
  ^\s*$$
  ^. -> Error`
