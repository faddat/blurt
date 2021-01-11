package rest

import (
	"net/http"
	"strconv"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/faddat/blurt/x/blurt/types"
    "github.com/gorilla/mux"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createPostRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	Title string `json:"title"`
	Body string `json:"body"`
	Ipfs string `json:"ipfs"`
	Parent string `json:"parent"`
	Category string `json:"category"`
	
}

func createPostHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createPostRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedTitle := req.Title
		
		parsedBody := req.Body
		
		parsedIpfs := req.Ipfs
		
		parsedParent := req.Parent
		
		parsedCategory := req.Category
		

		msg := types.NewMsgCreatePost(
			req.Creator,
			parsedTitle,
			parsedBody,
			parsedIpfs,
			parsedParent,
			parsedCategory,
			
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updatePostRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	Title string `json:"title"`
	Body string `json:"body"`
	Ipfs string `json:"ipfs"`
	Parent string `json:"parent"`
	Category string `json:"category"`
	
}


func updatePostHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        id := mux.Vars(r)["id"]

		var req updatePostRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedTitle := req.Title
		
		parsedBody := req.Body
		
		parsedIpfs := req.Ipfs
		
		parsedParent := req.Parent
		
		parsedCategory := req.Category
		

		msg := types.NewMsgUpdatePost(
			req.Creator,
            id,
			parsedTitle,
			parsedBody,
			parsedIpfs,
			parsedParent,
			parsedCategory,
			
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type deletePostRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
}

func deletePostHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        id := mux.Vars(r)["id"]

		var req deletePostRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgDeletePost(
			req.Creator,
            id,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
