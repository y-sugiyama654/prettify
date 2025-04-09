package domain

type Formatter interface {
	Format(input []byte) (string, error)
}