package blockchain

import (
	"testing"
)

func TestNetworkConstants(t *testing.T) {
	// Test that all network constants exist
	expectedNetworks := []string{
		BaseMainnet,
		PolygonMainnet,
		BSCMainnet,
		LiskMainnet,
		KaiaMainnet,
		WorldChainMainnet,
		EtherlinkMainnet,
		GnosisMainnet,
	}

	for _, networkName := range expectedNetworks {
		config, exists := GetNetworkConfig(networkName)
		if !exists {
			t.Errorf("Network %s should exist in SupportedNetworks", networkName)
			continue
		}

		if config.ChainID == 0 {
			t.Errorf("Network %s should have a valid chain ID", networkName)
		}

		if config.Name == "" {
			t.Errorf("Network %s should have a name", networkName)
		}
	}
}

func TestGetNetworkConfigByChainID(t *testing.T) {
	// Test reverse lookup by chain ID
	testCases := map[uint64]string{
		8453:  BaseMainnet,
		137:   PolygonMainnet,
		56:    BSCMainnet,
		1135:  LiskMainnet,
		8217:  KaiaMainnet,
		480:   WorldChainMainnet,
		42793: EtherlinkMainnet,
		100:   GnosisMainnet,
	}

	for chainID, expectedNetwork := range testCases {
		config, networkName, exists := GetNetworkConfigByChainID(chainID)
		if !exists {
			t.Errorf("Chain ID %d should exist", chainID)
			continue
		}

		if networkName != expectedNetwork {
			t.Errorf("Chain ID %d should map to %s, got %s", chainID, expectedNetwork, networkName)
		}

		if config.ChainID != chainID {
			t.Errorf("Config chain ID should match input: expected %d, got %d", chainID, config.ChainID)
		}
	}
}

func TestIsNetworkSupported(t *testing.T) {
	// Test valid networks
	validNetworks := []string{
		BaseMainnet,
		PolygonMainnet,
		BSCMainnet,
		LiskMainnet,
		KaiaMainnet,
		WorldChainMainnet,
		EtherlinkMainnet,
		GnosisMainnet,
	}
	for _, network := range validNetworks {
		if !IsNetworkSupported(network) {
			t.Errorf("Network %s should be supported", network)
		}
	}

	// Test invalid network
	if IsNetworkSupported("InvalidNetwork") {
		t.Error("Invalid network should not be supported")
	}
}

func TestIsChainSupported(t *testing.T) {
	// Test valid chain IDs
	validChainIDs := []uint64{8453, 137, 56, 1135, 8217, 480, 42793, 100}
	for _, chainID := range validChainIDs {
		if !IsChainSupported(chainID) {
			t.Errorf("Chain ID %d should be supported", chainID)
		}
	}

	// Test invalid chain ID
	if IsChainSupported(999999) {
		t.Error("Invalid chain ID should not be supported")
	}
}

func TestGetSupportedNetworks(t *testing.T) {
	networks := GetSupportedNetworks()
	if len(networks) == 0 {
		t.Error("Should have at least one supported network")
	}

	// Check that all expected networks are present
	expectedNetworks := map[string]bool{
		BaseMainnet:       false,
		PolygonMainnet:    false,
		BSCMainnet:        false,
		LiskMainnet:       false,
		KaiaMainnet:       false,
		WorldChainMainnet: false,
		EtherlinkMainnet:  false,
		GnosisMainnet:     false,
	}

	for _, network := range networks {
		if _, exists := expectedNetworks[network]; exists {
			expectedNetworks[network] = true
		}
	}

	for network, found := range expectedNetworks {
		if !found {
			t.Errorf("Expected network %s not found in supported networks", network)
		}
	}
}

func TestNetworkNameByChainID(t *testing.T) {
	// Test valid mappings
	testCases := map[uint64]string{
		8453:  BaseMainnet,
		137:   PolygonMainnet,
		56:    BSCMainnet,
		1135:  LiskMainnet,
		8217:  KaiaMainnet,
		480:   WorldChainMainnet,
		42793: EtherlinkMainnet,
		100:   GnosisMainnet,
	}

	for chainID, expectedName := range testCases {
		name, exists := GetNetworkNameByChainID(chainID)
		if !exists {
			t.Errorf("Chain ID %d should have a network name", chainID)
			continue
		}

		if name != expectedName {
			t.Errorf("Chain ID %d should map to %s, got %s", chainID, expectedName, name)
		}
	}

	// Test invalid chain ID
	_, exists := GetNetworkNameByChainID(999999)
	if exists {
		t.Error("Invalid chain ID should not have a network name")
	}
}
