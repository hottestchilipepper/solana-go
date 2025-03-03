package rpc

import (
	"context"
)

// GetMaxShredInsertSlot returns the max slot seen from after shred insert.
func (cl *Client) GetMaxShredInsertSlot(ctx context.Context) (out uint64, err error) {
	err = cl.rpcClient.CallForInto(ctx, &out, "getMaxShredInsertSlot", nil)
	return
}
