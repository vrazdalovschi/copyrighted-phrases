package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MinNamePrice is Initial Starting Price for a name that was never previously owned
var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

type Texts struct {
	Value string         `json:"value" yaml:"value"`
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
	Block int64          `json:"block" yaml:"block"`
}

// implement fmt.Stringer
func (w Texts) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s`, w.Owner, w.Value))
}
