package manager

import "context"

type DummyStateLoader struct {
}

func (l *DummyStateLoader) StreamUpdate(ctx context.Context, lastSync uint64, iter ItemIterator) error {
	return nil
}
