package models

import "testing"

func TestNewIDRXRatesRequest(t *testing.T) {
	t.Run("with chain ID", func(t *testing.T) {
		req := NewIDRXRatesRequest("100000", "137")

		if req.IDRXAmount == nil || *req.IDRXAmount != "100000" {
			t.Errorf("Expected IDRXAmount to be 100000, got %v", req.IDRXAmount)
		}

		if req.ChainID == nil || *req.ChainID != "137" {
			t.Errorf("Expected ChainID to be 137, got %v", req.ChainID)
		}

		if req.USDTAmount != nil {
			t.Errorf("Expected USDTAmount to be nil, got %v", req.USDTAmount)
		}
	})

	t.Run("without chain ID", func(t *testing.T) {
		req := NewIDRXRatesRequest("100000")

		if req.IDRXAmount == nil || *req.IDRXAmount != "100000" {
			t.Errorf("Expected IDRXAmount to be 100000, got %v", req.IDRXAmount)
		}

		if req.ChainID != nil {
			t.Errorf("Expected ChainID to be nil, got %v", *req.ChainID)
		}
	})

	t.Run("with empty chain ID", func(t *testing.T) {
		req := NewIDRXRatesRequest("100000", "")

		if req.ChainID != nil {
			t.Errorf("Expected ChainID to be nil for empty string, got %v", *req.ChainID)
		}
	})
}

func TestNewUSDTRatesRequest(t *testing.T) {
	t.Run("with chain ID", func(t *testing.T) {
		req := NewUSDTRatesRequest("50000", "56")

		if req.USDTAmount == nil || *req.USDTAmount != "50000" {
			t.Errorf("Expected USDTAmount to be 50000, got %v", req.USDTAmount)
		}

		if req.ChainID == nil || *req.ChainID != "56" {
			t.Errorf("Expected ChainID to be 56, got %v", req.ChainID)
		}

		if req.IDRXAmount != nil {
			t.Errorf("Expected IDRXAmount to be nil, got %v", req.IDRXAmount)
		}
	})

	t.Run("without chain ID", func(t *testing.T) {
		req := NewUSDTRatesRequest("50000")

		if req.USDTAmount == nil || *req.USDTAmount != "50000" {
			t.Errorf("Expected USDTAmount to be 50000, got %v", req.USDTAmount)
		}

		if req.ChainID != nil {
			t.Errorf("Expected ChainID to be nil, got %v", *req.ChainID)
		}
	})

	t.Run("with empty chain ID", func(t *testing.T) {
		req := NewUSDTRatesRequest("50000", "")

		if req.ChainID != nil {
			t.Errorf("Expected ChainID to be nil for empty string, got %v", *req.ChainID)
		}
	})
}

func TestNewFeesRequest(t *testing.T) {
	t.Run("with chain ID", func(t *testing.T) {
		req := NewFeesRequest(FeeTypeMint, "137")

		if req.FeeType != FeeTypeMint {
			t.Errorf("Expected FeeType to be MINT, got %v", req.FeeType)
		}

		if req.ChainID == nil || *req.ChainID != "137" {
			t.Errorf("Expected ChainID to be 137, got %v", req.ChainID)
		}
	})

	t.Run("without chain ID", func(t *testing.T) {
		req := NewFeesRequest(FeeTypeRedeem)

		if req.FeeType != FeeTypeRedeem {
			t.Errorf("Expected FeeType to be REDEEM, got %v", req.FeeType)
		}

		if req.ChainID != nil {
			t.Errorf("Expected ChainID to be nil, got %v", *req.ChainID)
		}
	})

	t.Run("with empty chain ID", func(t *testing.T) {
		req := NewFeesRequest(FeeTypeRefund, "")

		if req.ChainID != nil {
			t.Errorf("Expected ChainID to be nil for empty string, got %v", *req.ChainID)
		}
	})
}
