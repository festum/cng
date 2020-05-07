package regular

import (
	"testing"

	"github.com/festum/cng"
	"github.com/festum/cng/config"
)

func TestNew(t *testing.T) {
	c := config.LoadConfig()
	db := fate.InitDatabaseWithConfig(*c)
	regular := New(db)
	regular.Run()
}
