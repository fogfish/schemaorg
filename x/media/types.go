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

// Media Type
type Type string

// Original resolution
type Origin schemaorg.Url

// Blur Hash image placeholder https://blurha.sh
type Blur string

// Half Quarter VGA (240 x 160), use for thumbnails
type HQVGA schemaorg.Url

// Wide Quarter VGA (400 x 240), use for high resolution thumbnails
type WQVGA schemaorg.Url

// Standard Definition (720 x 480)
type SD schemaorg.Url

// High Definition (1280 x 720)
type HD schemaorg.Url

// Full HD (1920 x 1080)
type FHD schemaorg.Url

// Quad HD, 2K (2560 x 1440)
type QHD schemaorg.Url

// Ultra HD, False 4K (3840 x 2160)
type UHD schemaorg.Url

// Hi-Vision System, True 4K (4096 × 2160)
type HV schemaorg.Url

// Super Hi-Vision System, 8K (7680 × 4320)
type SHV schemaorg.Url
