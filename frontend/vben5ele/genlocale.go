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
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/Tricitrus/goctls/api/spec"
	"github.com/Tricitrus/goctls/util"
	"github.com/Tricitrus/goctls/util/pathx"
)

func genLocale(g *GenContext) error {
	var localeEnData, localeZhData strings.Builder
	var enLocaleFileName, zhLocaleFileName string
	enLocaleFileName = filepath.Join(g.LocaleDir, "en-US", fmt.Sprintf("%s.json", g.FolderName))
	zhLocaleFileName = filepath.Join(g.LocaleDir, "zh-CN", fmt.Sprintf("%s.json", g.FolderName))

	modelChineseName, modelEnglishName := g.ModelName, g.ModelName

	if g.ModelChineseName != "" {
		modelChineseName = g.ModelChineseName
	}

	if g.ModelEnglishName != "" {
		modelEnglishName = g.ModelEnglishName
	}

	for _, v := range g.ApiSpec.Types {
		if v.Name() == fmt.Sprintf("%sInfo", g.ModelName) {
			specData, ok := v.(spec.DefineStruct)
			if !ok {
				return errors.New("cannot get the field")
			}

			localeEnData.WriteString(fmt.Sprintf("  \"%s\": {\n", strcase.ToLowerCamel(g.ModelName)))
			localeZhData.WriteString(fmt.Sprintf("  \"%s\": {\n", strcase.ToLowerCamel(g.ModelName)))

			for _, val := range specData.Members {
				if val.Name != "" {
					localeEnData.WriteString(fmt.Sprintf("    \"%s\": \"%s\",\n",
						strcase.ToLowerCamel(val.Name), strcase.ToCamel(val.Name)))

					localeZhData.WriteString(fmt.Sprintf("    \"%s\": \"%s\",\n",
						strcase.ToLowerCamel(val.Name), strcase.ToCamel(val.Name)))
				}
			}

			localeEnData.WriteString(fmt.Sprintf("    \"add%s\": \"Add %s\",\n", g.ModelName, modelEnglishName))
			localeEnData.WriteString(fmt.Sprintf("    \"edit%s\": \"Edit %s\",\n", g.ModelName, modelEnglishName))
			localeEnData.WriteString(fmt.Sprintf("    \"%sList\": \"%s List\"\n", strcase.ToLowerCamel(g.ModelName), modelEnglishName))
			localeEnData.WriteString("  }")

			localeZhData.WriteString(fmt.Sprintf("    \"add%s\": \"添加 %s\",\n", g.ModelName, modelChineseName))
			localeZhData.WriteString(fmt.Sprintf("    \"edit%s\": \"编辑 %s\",\n", g.ModelName, modelChineseName))
			localeZhData.WriteString(fmt.Sprintf("    \"%sList\": \"%s 列表\"\n", strcase.ToLowerCamel(g.ModelName), modelChineseName))
			localeZhData.WriteString("  }")
		}
	}

	if !pathx.FileExists(enLocaleFileName) {
		if err := util.With("localeTpl").Parse(localeTpl).SaveTo(map[string]any{
			"localeData": localeEnData.String(),
		},
			enLocaleFileName, false); err != nil {
			return err
		}
	} else {
		file, err := os.ReadFile(enLocaleFileName)
		if err != nil {
			return err
		}

		data := string(file)

		if !strings.Contains(data, fmt.Sprintf("\"%s\":", strings.ToLower(g.ModelName))) && strings.Contains(data, ":") {
			data = data[:len(data)-3] + ",\n" + localeEnData.String() + data[len(data)-3:]
		} else if g.Overwrite {
			begin, end := FindBeginEndOfLocaleField(data, strings.ToLower(g.ModelName))
			data = data[:begin-2] + localeEnData.String() + data[end+1:]
		}

		err = os.WriteFile(enLocaleFileName, []byte(data), os.ModePerm)
		if err != nil {
			return err
		}
	}

	if !pathx.FileExists(zhLocaleFileName) {
		if err := util.With("localeTpl").Parse(localeTpl).SaveTo(map[string]any{
			"localeData": localeZhData.String(),
		},
			zhLocaleFileName, false); err != nil {
			return err
		}
	} else {
		file, err := os.ReadFile(zhLocaleFileName)
		if err != nil {
			return err
		}

		data := string(file)

		if !strings.Contains(data, fmt.Sprintf("\"%s\":", strcase.ToLowerCamel(g.ModelName))) && strings.Contains(data, ":") {
			data = data[:len(data)-3] + ",\n" + localeZhData.String() + data[len(data)-3:]
		} else if g.Overwrite {
			begin, end := FindBeginEndOfLocaleField(data, fmt.Sprintf("\"%s\"", strcase.ToLowerCamel(g.ModelName)))
			data = data[:begin-2] + localeZhData.String() + data[end+1:]
		}

		err = os.WriteFile(zhLocaleFileName, []byte(data), os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
