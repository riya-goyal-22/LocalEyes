package interfaces

import "context"

type CursorInterface interface {
	//ID() int64
	Next(ctx context.Context) bool
	//TryNext(ctx context.Context) bool
	//next(ctx context.Context, nonBlocking bool) bool
	Decode(val interface{}) error
	Err() error
	Close(ctx context.Context) error
	//All(ctx context.Context, results interface{}) error
	//RemainingBatchLength() int
	//addFromBatch(sliceVal reflect.Value, elemType reflect.Type, batch *bsoncore.DocumentSequence, index int) (reflect.Value, int, error)
	//closeImplicitSession()
	//SetBatchSize(batchSize int32)
	//SetMaxTime(dur time.Duration)
	//SetComment(comment interface{})
}
