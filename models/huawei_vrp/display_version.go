package huawei_vrp

type DisplayVersion struct {
	Patch_version   string `json:"PATCH_VERSION"`
	Vrp_version     string `json:"VRP_VERSION"`
	Product_version string `json:"PRODUCT_VERSION"`
	Model           string `json:"MODEL"`
	Uptime          string `json:"UPTIME"`
}

var DisplayVersion_Template string = `Value VRP_VERSION (\S+)
Value PRODUCT_VERSION (.+)
Value MODEL (((?!\sRouter).)+)
Value UPTIME (.+)
Value PATCH_VERSION (\S+)


Start
  ^Huawei\s+Versatile\s+Routing\s+Platform\s+Software
  ^Copyright.+(Huawei|HUAWEI).+
  ^.*software,\s+Version\s+${VRP_VERSION}\s+\(${PRODUCT_VERSION}\)
  ^Patch\s+[Vv]ersion\s*:\s+${PATCH_VERSION}\s*$$
  ^(HUAWEI|Huawei|Quidway)\s+${MODEL}\s+(Router\s+)?uptime\s+is\s+${UPTIME}$$
  ^BKP\s+\d+\s+version\s+information -> EOF
  ^\s*$$ -> EOF
  ^. -> Error
`
