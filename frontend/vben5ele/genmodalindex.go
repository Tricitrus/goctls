// Copyright (C) 2023  Ryan SU (https://github.com/suyuan32)

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package vben5ele

import (
	"fmt"
	"path/filepath"

	"github.com/iancoleman/strcase"

	"github.com/Tricitrus/goctls/util"
)

func genModalIndex(g *GenContext) error {
	if g.FormType != "modal" {
		return nil
	}

	if err := util.With("modelIndexTpl").Parse(modalIndexTpl).SaveTo(map[string]any{
		"modelName":           g.ModelName,
		"modelNameLowerCamel": strcase.ToLowerCamel(g.ModelName),
		"folderName":          g.FolderName,
		"addButtonTitle":      fmt.Sprintf("{{ $t('%s.%s.add%s') }}", g.FolderName, strcase.ToLowerCamel(g.ModelName), g.ModelName),
	},
		filepath.Join(g.ViewDir, "index.vue"), g.Overwrite); err != nil {
		return err
	}
	return nil
}
