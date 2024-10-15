package cisco_ios

type ShowHostsSummary struct {
	Local_cache_entries   string `json:"LOCAL_CACHE_ENTRIES"`
	Dynamic_cache_entries string `json:"DYNAMIC_CACHE_ENTRIES"`
	Default_domain        string `json:"DEFAULT_DOMAIN"`
	Name_servers          string `json:"NAME_SERVERS"`
}

var ShowHostsSummary_Template string = `Value DEFAULT_DOMAIN (\S+)
Value NAME_SERVERS (.+)
Value LOCAL_CACHE_ENTRIES (\d+)
Value DYNAMIC_CACHE_ENTRIES (\d+)

Start
  ^Default\s*domain\s*is\s*${DEFAULT_DOMAIN}
  ^Name\s*servers\s*are\s*${NAME_SERVERS}
  ^Local\s*cache\s*entries:\s*${LOCAL_CACHE_ENTRIES}
  ^Dynamic\s*cache\s*entries:\s*${DYNAMIC_CACHE_ENTRIES}
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^\s*$$
  ^. -> Error
`
