package lock

import (
	"context"

	"gopkg.in/errgo.v1"
)

func (l *EtcdLock) Release() error {
	if l == nil {
		return errgo.New("nil lock")
	}

	l.Lock()
	defer l.Unlock()

	defer l.session.Close()
	if !l.released {
		err := l.mutex.Unlock(context.Background())
		if err != nil {
			return err
		}
		l.released = true
	}
	return nil
}
