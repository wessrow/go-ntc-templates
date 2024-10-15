package huawei_smartax

type DisplayOntGemport0Ontid0 struct {
	Upstream_pq         string `json:"UPSTREAM_PQ"`
	Downstream_pq       string `json:"DOWNSTREAM_PQ"`
	Traffic_table_index string `json:"TRAFFIC_TABLE_INDEX"`
	Gem_port            string `json:"GEM_PORT"`
	T_cont_id           string `json:"T_CONT_ID"`
	Service_type        string `json:"SERVICE_TYPE"`
	Encrypt             string `json:"ENCRYPT"`
}

var DisplayOntGemport0Ontid0_Template string = `Value Key GEM_PORT (\d+)
Value T_CONT_ID (\d+)
Value SERVICE_TYPE (\S+)
Value ENCRYPT (\S+)
Value UPSTREAM_PQ (\S+)
Value DOWNSTREAM_PQ (\S+)
Value TRAFFIC_TABLE_INDEX (\S+)

Start
  ^\s*GEM\s*port\s*T-CONT\s*Service\s*Encrypt\s*Up\s*Down\s*Traffic\s*
  ^\s*ID\s*ID\s*type\s*PQ\s*PQ\s*table\s*index
  ^\s+${GEM_PORT}\s+${T_CONT_ID}\s+${SERVICE_TYPE}\s+${ENCRYPT}\s+(-|${UPSTREAM_PQ})\s+(-|${DOWNSTREAM_PQ})\s+(-|${TRAFFIC_TABLE_INDEX}) -> Record
  ^\s*The\s*number\s*of\s*GEM\s*ports\s*is:\s*\d+
  ^\s*
  ^. -> Error
`
