package cmd

type Command interface {
	Name() string
	Description() string
	Execute(args []string) error
}
