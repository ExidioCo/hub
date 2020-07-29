// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/provider/types
// ALIASGEN: github.com/sentinel-official/hub/x/provider/keeper
// ALIASGEN: github.com/sentinel-official/hub/x/provider/querier
package provider

import (
	"github.com/sentinel-official/hub/x/provider/keeper"
	"github.com/sentinel-official/hub/x/provider/querier"
	"github.com/sentinel-official/hub/x/provider/types"
)

const (
	Codespace               = types.Codespace
	EventTypeSetProvider    = types.EventTypeSetProvider
	EventTypeUpdateProvider = types.EventTypeUpdateProvider
	AttributeKeyAddress     = types.AttributeKeyAddress
	ModuleName              = types.ModuleName
	QuerierRoute            = types.QuerierRoute
	QueryProvider           = types.QueryProvider
	QueryProviders          = types.QueryProviders
)

var (
	// functions aliases
	RegisterCodec             = types.RegisterCodec
	ErrorMarshal              = types.ErrorMarshal
	ErrorUnmarshal            = types.ErrorUnmarshal
	ErrorUnknownMsgType       = types.ErrorUnknownMsgType
	ErrorUnknownQueryType     = types.ErrorUnknownQueryType
	ErrorInvalidField         = types.ErrorInvalidField
	ErrorDuplicateProvider    = types.ErrorDuplicateProvider
	ErrorProviderDoesNotExist = types.ErrorProviderDoesNotExist
	NewGenesisState           = types.NewGenesisState
	DefaultGenesisState       = types.DefaultGenesisState
	ProviderKey               = types.ProviderKey
	NewMsgRegisterProvider    = types.NewMsgRegisterProvider
	NewMsgUpdateProvider      = types.NewMsgUpdateProvider
	NewQueryProviderParams    = types.NewQueryProviderParams
	NewKeeper                 = keeper.NewKeeper
	Querier                   = querier.Querier

	// variable aliases
	ModuleCdc         = types.ModuleCdc
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	ProviderKeyPrefix = types.ProviderKeyPrefix
)

type (
	GenesisState        = types.GenesisState
	MsgRegisterProvider = types.MsgRegisterProvider
	MsgUpdateProvider   = types.MsgUpdateProvider
	Provider            = types.Provider
	Providers           = types.Providers
	QueryProviderParams = types.QueryProviderParams
	Keeper              = keeper.Keeper
)