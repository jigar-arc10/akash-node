package events

import (
	"testing"

	"github.com/akash-network/node/sdkutil"
	"github.com/akash-network/node/testutil"
	dtypes "github.com/akash-network/node/x/deployment/types/v1beta2"
	mtypes "github.com/akash-network/node/x/market/types/v1beta2"
	ptypes "github.com/akash-network/node/x/provider/types/v1beta2"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func Test_processEvent(t *testing.T) {
	tests := []sdkutil.ModuleEvent{
		// x/deployment events
		dtypes.NewEventDeploymentCreated(testutil.DeploymentID(t), testutil.DeploymentVersion(t)),
		dtypes.NewEventDeploymentUpdated(testutil.DeploymentID(t), testutil.DeploymentVersion(t)),
		dtypes.NewEventDeploymentClosed(testutil.DeploymentID(t)),
		dtypes.NewEventGroupClosed(testutil.GroupID(t)),

		// x/market events
		mtypes.NewEventOrderCreated(testutil.OrderID(t)),
		mtypes.NewEventOrderClosed(testutil.OrderID(t)),
		mtypes.NewEventBidCreated(testutil.BidID(t), testutil.DecCoin(t)),
		mtypes.NewEventBidClosed(testutil.BidID(t), testutil.DecCoin(t)),
		mtypes.NewEventLeaseCreated(testutil.LeaseID(t), testutil.DecCoin(t)),
		mtypes.NewEventLeaseClosed(testutil.LeaseID(t), testutil.DecCoin(t)),

		// x/provider events
		ptypes.NewEventProviderCreated(testutil.AccAddress(t)),
		ptypes.NewEventProviderUpdated(testutil.AccAddress(t)),
		ptypes.NewEventProviderDeleted(testutil.AccAddress(t)),
	}

	for _, test := range tests {
		sdkevs := sdk.Events{
			test.ToSDKEvent(),
		}.ToABCIEvents()

		sdkev := sdkevs[0]

		ev, ok := processEvent(sdkev)
		assert.True(t, ok, test)
		assert.Equal(t, test, ev, test)
	}
}
