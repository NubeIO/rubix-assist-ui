package main

import (
	"context"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/utilstime"
	"time"
)

// sendTimeToUI
func (init App) sendTimeToUI(ctx context.Context) {
	for {
		time.Sleep(time.Millisecond * 1000)
		MsgToUI(ctx, "os-time", utilstime.TimeStamp())
	}
}
