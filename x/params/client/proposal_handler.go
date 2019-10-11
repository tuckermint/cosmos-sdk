package client

import (
	govclient "github.com/tuckermint/cosmos-sdk/x/gov/client"
	"github.com/tuckermint/cosmos-sdk/x/params/client/cli"
	"github.com/tuckermint/cosmos-sdk/x/params/client/rest"
)

// param change proposal handler
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
