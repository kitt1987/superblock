package xfs

import (
	"encoding/binary"
	"fmt"
	"os"
)

// from kernel/fs/xfs/xfs_sb.h

type xfs_drfsbno_t uint64
type xfs_drtbno_t uint64
type uuid_t [16]byte
type xfs_dfsbno_t uint64
type xfs_ino_t uint64
type xfs_agblock_t uint32
type xfs_agnumber_t uint32
type xfs_extlen_t uint32
type xfs_lsn_t int64

type SuperBlock struct {
	SB_magicnum              uint32
	SB_blocksize             uint32
	SB_dblocks               xfs_drfsbno_t
	SB_rblocks               xfs_drfsbno_t
	SB_rextents              xfs_drtbno_t
	SB_uuid                  uuid_t
	SB_logstart              xfs_dfsbno_t
	SB_rootino               xfs_ino_t
	SB_rbmino                xfs_ino_t
	SB_rsumino               xfs_ino_t
	SB_rextsize              xfs_agblock_t
	SB_agblocks              xfs_agblock_t
	SB_agcount               xfs_agnumber_t
	SB_rbmblocks             xfs_extlen_t
	SB_logblocks             xfs_extlen_t
	SB_versionnum            uint16
	SB_sectsize              uint16
	SB_inodesize             uint16
	SB_inopblock             uint16
	SB_fname                 [12]byte
	SB_blocklog              uint8
	SB_sectlog               uint8
	SB_inodelog              uint8
	SB_inopblog              uint8
	SB_agblklog              uint8
	SB_rextslog              uint8
	SB_inprogress            uint8
	SB_imax_pct              uint8
	SB_icount                uint64
	SB_ifree                 uint64
	SB_fdblocks              uint64
	SB_frextents             uint64
	SB_uquotino              xfs_ino_t
	SB_gquotino              xfs_ino_t
	SB_qflags                uint16
	SB_flags                 uint8
	SB_shared_vn             uint8
	SB_inoalignmt            xfs_extlen_t
	SB_unit                  uint32
	SB_width                 uint32
	SB_dirblklog             uint8
	SB_logsectlog            uint16
	SB_logsunit              uint32
	SB_features2             uint32
	SB_bad_features2         uint32
	SB_features_comapt       uint32
	SB_features_ro_comapt    uint32
	SB_features_incompat     uint32
	SB_features_log_incompat uint32
	SB_crc                   uint32
	SB_pad                   uint32
	SB_pquotino              xfs_ino_t
	SB_lsn                   xfs_lsn_t
}

func GetSuperBlock(blkPath string) (sb *SuperBlock, err error) {
	blkFD, err := os.Open(blkPath)
	if err != nil {
		return
	}

	defer blkFD.Close()

	sb = &SuperBlock{}
	err = binary.Read(blkFD, binary.BigEndian, sb)
	if err != nil {
		return
	}

	if sb.SB_magicnum != 0x58465342 {
		err = fmt.Errorf("magic number not matched")
	}

	return
}
