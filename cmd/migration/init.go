package migration

import (
	"ariga.io/atlas-provider-gorm/gormschema"
	"fmt"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"io"
	"os"
)

func loadSchemas() {
	stmts, err := gormschema.New("postgresql").Load(model.Models...)
	if err != nil {
		fmt.Errorf("failed to load gorm schema: %v", err)
	}

	io.WriteString(os.Stdout, stmts)
}
