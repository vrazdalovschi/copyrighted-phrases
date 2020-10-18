package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/x/params"
)

// Default parameter namespace
const (
	DefaultParamspace = ModuleName
)

// ParamKeyTable for copyrightedphrases module
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

// Params - used for initializing default parameter for copyrightedphrases at genesis
type Params struct{}

// NewParams creates a new Params object
func NewParams() Params {
	return Params{}
}

// String implements the stringer interface for Params
func (p Params) String() string {
	return fmt.Sprintf(``)
}

// ParamSetPairs - Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams()
}
