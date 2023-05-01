package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/kings-crown/cosmos-sdk/tree/main/x/auction/types"
)

const (
	restAuctionID = "auction-id"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc("/auction/{auction-id}", getAuctionHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/auction/bid", submitBidHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/auction/reveal", revealBidHandler(clientCtx)).Methods("POST")
}

func getAuctionHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		auctionID := vars[restAuctionID]

		res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/auction/get/%s", auctionID), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, clientCtx, res)
	}
}

func submitBidHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SubmitBidReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgSubmitBid(req.Bidder, req.AuctionId, req.Commitment)
		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func revealBidHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RevealBidReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgRevealBid(req.Bidder, req.AuctionId, req.BidAmount, req.Nonce)
		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
