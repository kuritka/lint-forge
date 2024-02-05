package installer

type Installer interface {
	Install(string) *Output
}

type Output struct {
	Output []byte
	Error  error
}

func (o *Output) String() string {
	return string(o.Output)
}
