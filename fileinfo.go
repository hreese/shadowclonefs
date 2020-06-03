package shadowclonefs

import (
	"github.com/vmihailenco/msgpack"
	"golang.org/x/sys/unix"
)

type FSObjectMetadata struct {
	// Error reading this file
	Error   error
	Name    string
	Mode    uint32
	Dev     uint64
	Inode   uint64
	Nlink   uint64
	Uid     uint32
	Gid     uint32
	Rdev    int32
	Atim    unix.Timespec
	Mtim    unix.Timespec
	Ctim    unix.Timespec
	Size    int64
	Blocks  int64
	Blksize int32
	Flags   uint32
	Gen     uint32
}

func (fom *FSObjectMetadata) Serialize() ([]byte, error) {
	return msgpack.Marshal(fom)
}

func DeserializeFSObjectMetadata(input []byte) (FSObjectMetadata, error) {
	var fom FSObjectMetadata
	err := msgpack.Unmarshal(input, &fom)
	return fom, err
}

func FileMetadata(filename string) (FSObjectMetadata, error) {
	var (
		metadata = FSObjectMetadata{
			Name: filename,
		}
		lstatResult unix.Stat_t
		err         error = nil
	)

	// read file metadata using lstat()
	err = unix.Lstat(filename, &lstatResult)
	if err != nil {
		metadata.Error = err
	}

	metadata.Mode = lstatResult.Mode
	metadata.Dev = lstatResult.Dev
	metadata.Inode = lstatResult.Ino
	metadata.Nlink = lstatResult.Nlink
	// TODO: continue here

	return metadata, err

	// "golang.org/x/sys/unix"
	/*
		unix.Listxattr()
		unix.Lstat()
		unix.Getxattr()
		unix.Statfs()
	*/

}
