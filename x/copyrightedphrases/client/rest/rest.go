package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers copyrightedphrases-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/copyrightedphrases/text", registerTextHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/copyrightedphrases/text", listCopyrightedTextHandler(cliCtx, "copyrightedphrases")).Methods("GET")
	r.HandleFunc("/copyrightedphrases/text/{key}", getCopyrightedTextHandler(cliCtx, "copyrightedphrases")).Methods("GET")
	r.HandleFunc("/copyrightedphrases/text", deleteCopyrightedTextHandler(cliCtx)).Methods("DELETE")

}
