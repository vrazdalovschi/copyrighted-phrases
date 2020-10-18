package types

import "fmt"

// GenesisState - all copyrightedphrases state that must be provided at genesis
type GenesisState struct {
	// TODO: Fill out what is needed by the module for genesis
	CopyrightedTextRecords []Texts `json:"copyrighted_text_records"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState( /* TODO: Fill out with what is needed for genesis state */ ) GenesisState {
	return GenesisState{
		// TODO: Fill out according to your genesis state
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
	// TODO: Create a sanity check to make sure the state conforms to the modules needs
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
