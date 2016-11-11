package app

import (
	"testing"

	"github.com/kopiczko/mikro/app/apppb"
)

func TestAppHandler(t *testing.T) {
	var _ apppb.AppHandler = new(App) // App should be created with New
}
