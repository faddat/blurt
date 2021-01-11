package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/faddat/blurt/x/blurt/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strconv"
)

// GetPostCount get the total number of post
func (k Keeper) GetPostCount(ctx sdk.Context) int64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostCountKey))
	byteKey := types.KeyPrefix(types.PostCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetPostCount set the total number of post
func (k Keeper) SetPostCount(ctx sdk.Context, count int64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostCountKey))
	byteKey := types.KeyPrefix(types.PostCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreatePost(ctx sdk.Context, msg types.MsgCreatePost) {
	// Create the post
    count := k.GetPostCount(ctx)
    var post = types.Post{
        Creator: msg.Creator,
        Id:      strconv.FormatInt(count, 10),
        Title: msg.Title,
        Body: msg.Body,
        Ipfs: msg.Ipfs,
        Category: msg.Category,
    }

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
    key := types.KeyPrefix(types.PostKey + post.Id)
    value := k.cdc.MustMarshalBinaryBare(&post)
    store.Set(key, value)

    // Update post count
    k.SetPostCount(ctx, count+1)
}

func (k Keeper) UpdatePost(ctx sdk.Context, post types.Post) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	b := k.cdc.MustMarshalBinaryBare(&post)
	store.Set(types.KeyPrefix(types.PostKey + post.Id), b)
}

func (k Keeper) GetPost(ctx sdk.Context, key string) types.Post {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	var post types.Post
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.PostKey + key)), &post)
	return post
}

func (k Keeper) HasPost(ctx sdk.Context, id string) bool {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	return store.Has(types.KeyPrefix(types.PostKey + id))
}

func (k Keeper) GetPostOwner(ctx sdk.Context, key string) string {
    return k.GetPost(ctx, key).Creator
}

// DeletePost deletes a post
func (k Keeper) DeletePost(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	store.Delete(types.KeyPrefix(types.PostKey + key))
}

func (k Keeper) GetAllPost(ctx sdk.Context) (msgs []types.Post) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PostKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Post
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
