package adapter

import "context"

type gate chan struct{}

func newGate() gate {
	result := make(gate, 1)
	result <- struct{}{}
	return result
}
func (g gate) Lock()   { <-g }
func (g gate) Unlock() { g <- struct{}{} }
func (g gate) lock(control *ControlState, root context.Context) error {
	if control == nil {
		return ErrArgument
	}
	if err := control.ctx.Err(); err != nil {
		return err
	}
	select {
	case <-g:
		if err := control.ctx.Err(); err != nil {
			g.Unlock()
			return err
		}
		if root.Err() != nil {
			g.Unlock()
			return ErrClosed
		}
		return nil
	case <-control.ctx.Done():
		return control.ctx.Err()
	case <-root.Done():
		return ErrClosed
	}
}
