package querier

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

func querySession(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySessionParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	session, found := k.GetSession(ctx, params.ID)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(session)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func querySessions(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySessionsParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	sessions := k.GetSessions(ctx)

	start, end := client.Paginate(len(sessions), params.Page, params.Limit, len(sessions))
	if start < 0 || end < 0 {
		sessions = types.Sessions{}
	} else {
		sessions = sessions[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(sessions)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func querySessionsForSubscription(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySessionsForSubscriptionParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	sessions := k.GetSessionsForSubscription(ctx, params.ID)

	start, end := client.Paginate(len(sessions), params.Page, params.Limit, len(sessions))
	if start < 0 || end < 0 {
		sessions = types.Sessions{}
	} else {
		sessions = sessions[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(sessions)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func querySessionsForNode(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySessionsForNodeParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	sessions := k.GetSessionsForNode(ctx, params.Address)

	start, end := client.Paginate(len(sessions), params.Page, params.Limit, len(sessions))
	if start < 0 || end < 0 {
		sessions = types.Sessions{}
	} else {
		sessions = sessions[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(sessions)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func querySessionsForAddress(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySessionsForAddressParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	sessions := k.GetSessionsForAddress(ctx, params.Address)

	start, end := client.Paginate(len(sessions), params.Page, params.Limit, len(sessions))
	if start < 0 || end < 0 {
		sessions = types.Sessions{}
	} else {
		sessions = sessions[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(sessions)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryOngoingSession(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryOngoingSessionParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	session, found := k.GetOngoingSession(ctx, params.ID, params.Address)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(session)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
