package cisco_asa

type ShowRunningConfigCryptoMap struct {
	Matched_address string `json:"MATCHED_ADDRESS"`
	Map             string `json:"MAP"`
	Seq             string `json:"SEQ"`
	Pfs             string `json:"PFS"`
	Peer            string `json:"PEER"`
	Ike             string `json:"IKE"`
	Transform       string `json:"TRANSFORM"`
	Sa              string `json:"SA"`
}

var ShowRunningConfigCryptoMap_Template string = `Value MATCHED_ADDRESS (\S+)
Value Key MAP (\S+)
Value Key SEQ (\d+)
Value PFS (group\d|\s*)
Value PEER (\S+)
Value IKE (\S+)
Value TRANSFORM (\S+(\s\S+)*?)
Value SA (\d+)
#Future# Value Fillup INTERFACE (\S+)

Start
  # Value's address , start of block
  ^crypto\smap\s${MAP}\s${SEQ}\smatch\saddress\s${MATCHED_ADDRESS}\s*$$ -> ReadBlockLines
  ^. -> Error

ReadBlockLines
  # Fake start, block "match address" to trigger recording of current block
  ^crypto\s+map\s\S+\s\d+\s+match\s+address\s\S+\s*$$ -> Continue.Record
  # Real capture of "match address" start of new record
  ^crypto\s+map\s+${MAP}\s+${SEQ}\s+match\s+address\s+${MATCHED_ADDRESS}\s*$$
  #
  # Read other lines
  ^crypto\s+map\s+${MAP}\s+${SEQ}\s+set\s+pfs\s+${PFS}\s*$$
  ^crypto\s+map\s+${MAP}\s+${SEQ}\s+set\s+peer\s+${PEER}\s*$$
  ^crypto\s+map\s+${MAP}\s+${SEQ}\s+set\s+${IKE}\s+transform-set\s+${TRANSFORM}\s*$$
  ^crypto\s+map\s+${MAP}\s+${SEQ}\s+set\s+security-association\s+lifetime\s+seconds\s+${SA}\s*$$
  #Future# ^crypto\s+map\s+${MAP}\s+interface\s${INTERFACE}\s*
  #
  # Ignore rest
  ^crypto\s+map\s+\S+\s+interface\s+\S+\s*$$
  ^crypto\s+map\s+\S+\s+\d+\sipsec-isakmp\s+dynamic\s+SYSTEM_DEFAULT_CRYPTO_MAP\s*$$
  ^. -> Error
`
