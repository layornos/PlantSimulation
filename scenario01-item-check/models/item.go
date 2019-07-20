package models

import "github.com/layornos/godes"

type Item struct {
	*godes.Runner
	id     string
	defect bool
}
