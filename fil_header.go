// Copyright © 2021 Hidetatsu Yaginuma. All rights reserved.
package innodb

type FilHeader struct {
	Checksum                   [4]byte
	Offset                     [4]byte
	PreviousPage               [4]byte
	NextPage                   [4]byte
	LSNForLastPageModification [8]byte
	PageType                   [2]byte
	FlushLSN                   [8]byte
	SpaceID                    [4]byte
}
