package models

// NewIDRXRatesRequest creates a rates request for IDRX amount.
// The chainID parameter is optional.
func NewIDRXRatesRequest(idrxAmount string, chainID ...string) *RatesRequest {
	req := &RatesRequest{
		IDRXAmount: &idrxAmount,
	}
	if len(chainID) > 0 && chainID[0] != "" {
		req.ChainID = &chainID[0]
	}
	return req
}

// NewUSDTRatesRequest creates a rates request for USDT amount.
// The chainID parameter is optional.
func NewUSDTRatesRequest(usdtAmount string, chainID ...string) *RatesRequest {
	req := &RatesRequest{
		USDTAmount: &usdtAmount,
	}
	if len(chainID) > 0 && chainID[0] != "" {
		req.ChainID = &chainID[0]
	}
	return req
}

// NewFeesRequest creates a fees request with optional chain ID.
func NewFeesRequest(feeType FeeType, chainID ...string) *FeesRequest {
	req := &FeesRequest{
		FeeType: feeType,
	}
	if len(chainID) > 0 && chainID[0] != "" {
		req.ChainID = &chainID[0]
	}
	return req
}
