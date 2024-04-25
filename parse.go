package env

import "github.com/go-zoox/tag"

// Parse parses the environment into the struct of value.
func Parse(value any) error {
	tg := tag.New(templateTag, &EnvDataSource{})
	if err := tg.Decode(value); err != nil {
		return err
	}

	return nil
}
