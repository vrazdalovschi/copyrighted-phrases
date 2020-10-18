package types

import "fmt"

// GenesisState - all copyrightedphrases state that must be provided at genesis
type GenesisState struct {
	CopyrightedTextRecords []Texts `json:"copyrighted_text_records"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState() GenesisState {
	return GenesisState{
		CopyrightedTextRecords: nil,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		CopyrightedTextRecords: []Texts{},
	}
}

// ValidateGenesis validates the copyrightedphrases genesis parameters
func ValidateGenesis(data GenesisState) error {
	for _, record := range data.CopyrightedTextRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid CopyrightedTextRecord: Value: %s. Error: Missing Owner", record.Value)
		}
		if record.Value == "" {
			return fmt.Errorf("invalid CopyrightedTextRecord: Owner: %s. Error: Missing Value", record.Owner)
		}
	}
	return nil
}
