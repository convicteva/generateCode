package err

type Comerr struct {
	Msg string
}

func (err *Comerr) Error() string {
	return err.Msg
}
