package keeper

import (
	"interchange/x/dex/types"

	ibctransfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SaveVoucherDenom(ctx sdk.Context, port string, channel string, denom string) {
	voucher := VoucherDenom(port, channel, denom)

	_, saved := k.GetDenomTrace(ctx, voucher)
	if !saved {
		k.SetDenomTrace(ctx, types.DenomTrace{
			Index:   voucher,
			Port:    port,
			Channel: channel,
			Origin:  denom,
		})
	}
}

func VoucherDenom(port string, channel string, denom string) string {
	sourcePrefix := ibctransfertypes.GetDenomPrefix(port, channel)

	prefixDenom := sourcePrefix + denom

	denomTrace := ibctransfertypes.ParseDenomTrace(prefixDenom)
	voucher := denomTrace.IBCDenom()
	return voucher[:16]
}

func (k Keeper) OriginalDenom(ctx sdk.Context, port string, channel string, voucher string) (string, bool) {
	trace, exist := k.GetDenomTrace(ctx, voucher)
	if exist {
		// Check if original port and channel
		if trace.Port == port && trace.Channel == channel {
			return trace.Origin, true
		}
	}

	// Not the original chain
	return "", false
}
