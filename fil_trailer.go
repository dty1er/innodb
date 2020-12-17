// Copyright © 2021 Hidetatsu Yaginuma. All rights reserved.
package innodb

type FilTrailer struct {
	OldstyleChecksum [4]byte
	Low32BitsOfLSN   [4]byte
}
