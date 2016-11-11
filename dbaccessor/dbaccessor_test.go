package dbaccessor

import (
	"testing"

	"github.com/kopiczko/mikro/dbaccessor/dbaccessorpb"
)

func TestDBAccessorHandler(t *testing.T) {
	var _ dbaccessorpb.DBAccessorHandler = new(DBAccessor)
}
