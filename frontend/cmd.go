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

package frontend

import (
	"github.com/Tricitrus/goctls/frontend/vben"
	"github.com/Tricitrus/goctls/frontend/vben5"
	"github.com/Tricitrus/goctls/frontend/vben5ele"
	"github.com/Tricitrus/goctls/internal/cobrax"
)

var (
	// Cmd describes an api command.
	Cmd         = cobrax.NewCommand("frontend")
	VbenCmd     = cobrax.NewCommand("vben", cobrax.WithRunE(vben.GenCRUDLogic))
	Vben5Cmd    = cobrax.NewCommand("vben5", cobrax.WithRunE(vben5.GenCRUDLogic))
	Vben5EleCmd = cobrax.NewCommand("vben5-ele", cobrax.WithRunE(vben5ele.GenCRUDLogic))
)

func init() {
	vbenCmdFlags := VbenCmd.Flags()
	vben5CmdFlags := Vben5Cmd.Flags()
	vben5EleCmdFlags := Vben5EleCmd.Flags()

	vbenCmdFlags.StringVarPWithDefaultValue(&vben.VarStringOutput, "output", "o", "./")
	vbenCmdFlags.StringVarP(&vben.VarStringApiFile, "api_file", "a")
	vbenCmdFlags.StringVarPWithDefaultValue(&vben.VarStringFolderName, "folder_name", "f", "sys")
	vbenCmdFlags.StringVarP(&vben.VarStringSubFolder, "sub_folder", "s")
	vbenCmdFlags.StringVarPWithDefaultValue(&vben.VarStringApiPrefix, "prefix", "p", "sys-api")
	vbenCmdFlags.StringVarP(&vben.VarStringModelName, "model_name", "m")
	vbenCmdFlags.StringVarPWithDefaultValue(&vben.VarStringFormType, "form_type", "t", "drawer")
	vbenCmdFlags.BoolVarP(&vben.VarBoolOverwrite, "overwrite", "w")
	vbenCmdFlags.StringVar(&vben.VarStringModelChineseName, "model_chinese_name")
	vbenCmdFlags.StringVar(&vben.VarStringModelEnglishName, "model_english_name")

	vben5CmdFlags.StringVarPWithDefaultValue(&vben5.VarStringOutput, "output", "o", "./")
	vben5CmdFlags.StringVarP(&vben5.VarStringApiFile, "api_file", "a")
	vben5CmdFlags.StringVarPWithDefaultValue(&vben5.VarStringFolderName, "folder_name", "f", "sys")
	vben5CmdFlags.StringVarP(&vben5.VarStringSubFolder, "sub_folder", "s")
	vben5CmdFlags.StringVarPWithDefaultValue(&vben5.VarStringApiPrefix, "prefix", "p", "sys-api")
	vben5CmdFlags.StringVarP(&vben5.VarStringModelName, "model_name", "m")
	vben5CmdFlags.StringVarPWithDefaultValue(&vben5.VarStringFormType, "form_type", "t", "modal")
	vben5CmdFlags.BoolVarP(&vben5.VarBoolOverwrite, "overwrite", "w")
	vben5CmdFlags.StringVar(&vben5.VarStringModelChineseName, "model_chinese_name")
	vben5CmdFlags.StringVar(&vben5.VarStringModelEnglishName, "model_english_name")

	vben5EleCmdFlags.StringVarPWithDefaultValue(&vben5ele.VarStringOutput, "output", "o", "./")
	vben5EleCmdFlags.StringVarP(&vben5ele.VarStringApiFile, "api_file", "a")
	vben5EleCmdFlags.StringVarPWithDefaultValue(&vben5ele.VarStringFolderName, "folder_name", "f", "sys")
	vben5EleCmdFlags.StringVarP(&vben5ele.VarStringSubFolder, "sub_folder", "s")
	vben5EleCmdFlags.StringVarPWithDefaultValue(&vben5ele.VarStringApiPrefix, "prefix", "p", "sys-api")
	vben5EleCmdFlags.StringVarP(&vben5ele.VarStringModelName, "model_name", "m")
	vben5EleCmdFlags.StringVarPWithDefaultValue(&vben5ele.VarStringFormType, "form_type", "t", "modal")
	vben5EleCmdFlags.BoolVarP(&vben5ele.VarBoolOverwrite, "overwrite", "w")
	vben5EleCmdFlags.StringVar(&vben5ele.VarStringModelChineseName, "model_chinese_name")
	vben5EleCmdFlags.StringVar(&vben5ele.VarStringModelEnglishName, "model_english_name")

	Cmd.AddCommand(VbenCmd)
	Cmd.AddCommand(Vben5Cmd)
	Cmd.AddCommand(Vben5EleCmd)
}
