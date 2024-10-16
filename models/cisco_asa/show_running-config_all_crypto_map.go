package cisco_asa

type ShowRunningConfigAllCryptoMap struct {
	Matched_address     string `json:"MATCHED_ADDRESS"`
	Map                 string `json:"MAP"`
	Seq                 string `json:"SEQ"`
	Ikev2_mode          string `json:"IKEv2_MODE"`
	Interface           string `json:"INTERFACE"`
	Transform           string `json:"TRANSFORM"`
	Connection_type     string `json:"CONNECTION_TYPE"`
	Ikev1_phase1_mode   string `json:"IKEv1_PHASE1_MODE"`
	Ikev1_transform_set string `json:"IKEv1_TRANSFORM_SET"`
	Isakmp_dynamic      string `json:"ISAKMP_DYNAMIC"`
	Sa_kb               string `json:"SA_KB"`
	Peer                string `json:"PEER"`
	Sa_sec              string `json:"SA_SEC"`
	Tfc_packets         string `json:"TFC_PACKETS"`
	Pfs                 string `json:"PFS"`
}

var ShowRunningConfigAllCryptoMap_Template string = `Value MATCHED_ADDRESS (\S+)
Value CONNECTION_TYPE (\S+)
Value Required MAP (\S+)
Value Required SEQ (\d+)
Value PFS (group\d|\s*)
Value PEER (\S+)
Value IKEv1_PHASE1_MODE (\S+)
Value IKEv1_TRANSFORM_SET (\S+(\s\S+)*?)
Value IKEv2_MODE (\S+)
Value ISAKMP_DYNAMIC (\S+)
Value Fillup INTERFACE (\S+)
Value TRANSFORM (\S+)
Value SA_SEC (\d+)
Value SA_KB (\d+)
Value TFC_PACKETS (\S\S)

Start
  # Value's address , start of block
  ^crypto\smap\s${MAP}\s${SEQ}\smatch\saddress\s${MATCHED_ADDRESS}\s*$$ -> ReadBlockLines
  ^. -> Error

ReadBlockLines
  #1 Fake start, block "match address" to trigger recording of current block
  ^crypto\s+map\s\S+\s\d+\s+match\s+address\s\S+\s*$$ -> Continue.Record
  #1 Real capture of "match address" start of new record
  ^crypto\smap\s${MAP}\s${SEQ}\smatch\saddress\s${MATCHED_ADDRESS}\s*$$
  #
  ^crypto\smap\s${MAP}\s${SEQ}\sset\sconnection-type\s${CONNECTION_TYPE}\s*$$
  ^crypto\smap\s${MAP}\s${SEQ}\sset\spfs\s${PFS}\s*$$
  ^crypto\smap\s${MAP}\s${SEQ}\sset\speer\s${PEER}\s*$$
  ^crypto\smap\s${MAP}\s${SEQ}\sset\sikev1\sphase1-mode\s${IKEv1_PHASE1_MODE}\s*$$
  ^crypto\smap\s${MAP}\s${SEQ}\sset\sikev1\stransform-set\s${IKEv1_TRANSFORM_SET}\s*$$
  ^crypto\smap\s${MAP}\s${SEQ}\sset\sikev2\smode\s${IKEv2_MODE}\s*$$
  # SA Second/Byte alone or in different combinations
  ^crypto\smap\s\S+\s\d+\sset\ssecurity-association\slifetime\sseconds\s${SA_SEC}\s*$$
  ^crypto\smap\s\S+\s\d+\sset\ssecurity-association\slifetime\skilobytes\s${SA_KB}\s*$$
  ^crypto\smap\s\S+\s\d+\sset\ssecurity-association\slifetime\skilobytes\s${SA_KB}\sseconds\s${SA_SEC}\s*$$
  ^crypto\smap\s\S+\s\d+\sset\ssecurity-association\slifetime\sseconds\s${SA_SEC}\skilobytes\s${SA_KB}\s*$$
  #2 Fake start, block "match address" to trigger recording of current block
  ^crypto\s+map\s\S+\s\d+\sipsec-isakmp\sdynamic\s${ISAKMP_DYNAMIC}\s*$$ -> Continue.Record
  #2 Real capture of "match address" start of new record
  ^crypto\smap\s${MAP}\s${SEQ}\sipsec-isakmp\sdynamic\s${ISAKMP_DYNAMIC}\s*$$
  #
  #3 no crypto map only at end of each block, if unset (add -> Record for safety)
  ^${TFC_PACKETS}\scrypto\smap\s${MAP}\s${SEQ}\sset\stfc-packets\s*$$ -> Record
  #4 Interface only after multiple blocks, FillUp
  ^crypto\smap\s${MAP}\sinterface\s${INTERFACE}\s*$$
  #
  ^\s*$$
  ^. -> Error
`
