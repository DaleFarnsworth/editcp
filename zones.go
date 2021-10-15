// Copyright 2017-2021 Dale Farnsworth. All rights reserved.

// Dale Farnsworth
// 1007 W Mendoza Ave
// Mesa, AZ  85210
// USA
//
// dale@farnsworth.org

// This file is part of Editcp.
//
// Editcp is free software: you can redistribute it and/or modify
// it under the terms of version 3 of the GNU General Public License
// as published by the Free Software Foundation.
//
// Editcp is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Editcp.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"github.com/dalefarnsworth-dmr/codeplug"
	"github.com/dalefarnsworth-dmr/ui"
)

func zones(edt *editor) {
	writable := true
	edt.newRecordWindow(codeplug.RtZones_md380, writable, ziRecord)
}

func ziRecord(edt *editor, recordBox *ui.HBox) {
	r := currentRecord(recordBox.Window())

	column := recordBox.AddVbox()
	row := column.AddHbox()
	form := row.AddForm()
	form.AddFieldTypeRows(r, codeplug.FtZiName)

	if r.HasFieldType(codeplug.FtZiChannel_md380) {
		column.AddHbox().AddGroupbox("Channels").AddFieldMembers(r,
			codeplug.FtZiChannel_md380,
			&settings.sortAvailableChannels)
	}

	//func(r *codeplug.Record) bool {
	//	f := r.Field(codeplug.FtCiRxFrequency)
	//	freq, _ := strconv.ParseFloat(f.String(), 64)
	//	return r.Codeplug().FrequencyValidA(freq)
	//}

	if r.HasFieldType(codeplug.FtZiChannelA_uv380) {
		hBox := column.AddHbox()

		hBox.AddGroupbox("A Channels").AddHbox().AddFieldMembers(r,
			codeplug.FtZiChannelA_uv380,
			&settings.sortAvailableChannels)

		hBox.AddSpace(5)

		hBox.AddGroupbox("B Channels").AddHbox().AddFieldMembers(r,
			codeplug.FtZiChannelB_uv380,
			&settings.sortAvailableChannelsB)
	}
}
