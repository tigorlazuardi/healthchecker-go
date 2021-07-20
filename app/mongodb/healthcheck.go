package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (hc *HealthChecker) check(ctx context.Context) (err error) {
	ctx, done := context.WithTimeout(ctx, time.Second*5)
	defer done()
	return hc.client.Ping(ctx, readpref.PrimaryPreferred())
}
