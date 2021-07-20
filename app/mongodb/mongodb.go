package mongodb

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/tigorlazuardi/healthchecker/pkg"
	"go.mongodb.org/mongo-driver/mongo"
)

type HealthChecker struct {
	client *mongo.Client
	mu     *sync.RWMutex
	state  *pkg.PublishMessage
	ctx    context.Context
	done   chan pkg.Done
}

func NewHealthChecker(ctx context.Context, client *mongo.Client) *HealthChecker {
	hc := &HealthChecker{
		client: client,
		mu:     &sync.RWMutex{},
		state:  &pkg.PublishMessage{},
		ctx:    ctx,
	}
	go hc.loop()
	return hc
}

func (hc *HealthChecker) Publish(msgChan chan<- pkg.PublishMessage) {
	hc.mu.RLock()
	defer hc.mu.RUnlock()
	msgChan <- *hc.state
}

func (hc *HealthChecker) Name() string {
	return "mongodb"
}

func (hc *HealthChecker) loop() {
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-hc.ctx.Done():
			return
		case <-ticker.C:
			err := hc.check(hc.ctx)
			hc.mu.Lock()
			if err != nil {
				hc.state.Status = "error"
				hc.state.Code = 1
				hc.state.Message = err.Error()
			} else {
				hc.state.Code = 0
				hc.state.Status = "ok"
				hc.state.Message = "healthy"
			}
			hc.mu.Unlock()
		}
	}
}

func (hc *HealthChecker) Close() <-chan pkg.Done {
	go func() {
		ctx, done := context.WithTimeout(context.Background(), time.Second*5)
		defer done()
		err := hc.client.Disconnect(ctx)
		if err != nil {
			log.Println(err.Error())
		}
		hc.done <- pkg.Done{}
	}()
	return hc.done
}
