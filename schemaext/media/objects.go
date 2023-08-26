//
// Copyright (C) 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/schemaorg
//

package media

import (
	"github.com/fogfish/schemaorg"
)

// Urls to media object
type Urls struct {
	Blur   Blur          `json:"blur,omitempty"`
	HQVGA  HQVGA         `json:"hqvga,omitempty"`
	WQVGA  WQVGA         `json:"wqvga,omitempty"`
	SD     SD            `json:"sd,omitempty"`
	HD     HD            `json:"hd,omitempty"`
	FHD    FHD           `json:"fhd,omitempty"`
	QHD    QHD           `json:"qhd,omitempty"`
	UHD    UHD           `json:"uhd,omitempty"`
	HV     HV            `json:"hv,omitempty"`
	SHV    SHV           `json:"shv,omitempty"`
	Origin Origin        `json:"origin,omitempty"`
	Url    schemaorg.Url `json:"url,omitempty"`
}
