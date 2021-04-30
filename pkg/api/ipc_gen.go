// Code generated for package api by gen.go DO NOT EDIT. (@generated)
package api

func (c Calls) IpcsPublishBlockchain(blockchainID string) (decisionsURL string, consensusURL string, err error) {
	input := P{
		"blockchainID": blockchainID,
	}
	output := P{}
	err = c.client.NewRequest("ext/ipcs", "ipcs.publishBlockchain", input, &output)
	if err != nil {
		return decisionsURL, consensusURL, err
	}
	err = output.OutStr("decisionsURL", &decisionsURL).OutStr("consensusURL", &consensusURL).Error()
	return decisionsURL, consensusURL, err
}
