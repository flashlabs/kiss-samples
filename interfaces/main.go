package interfaces

type Runner interface {
	Run() error
	Stop() error
}

var _ Runner = (*Worker)(nil)

type Worker struct {
}

func (w Worker) Stop() error {
	return nil
}

func (w Worker) Run() error {
	return nil
}
