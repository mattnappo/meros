package database

import (
	"testing"

	"github.com/xoreo/meros/types"
)

func TestPutFile(t *testing.T) {
	filedb, err := Open("myfiledb", FILEDB)
	if err != nil {
		t.Fatal(err)
	}
	defer filedb.Close()

	file, err := types.NewFile("test_file.txt")
	if err != nil {
		t.Fatal(err)
	}

	fid, err := filedb.PutItem(*file)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("fileID: %s\n", fid.String())
}

func TestGetFile(t *testing.T) {
	filedb, err := Open("myfiledb", FILEDB)
	if err != nil {
		t.Fatal(err)
	}
	defer filedb.Close()

	fileid, err := IDFromString("8ade293f4d76ba0e0b1101c1274dda59925ecb178b763c12126a70d54d1b5469")
	if err != nil {
		t.Fatal(err)
	}

	file, err := filedb.GetItem(fileid)
	if err != nil {
		t.Logf("%v", file)
		t.Fatal(err)
	}
	// try dereffing the whole 47
	v := file.(*types.File)
	t.Logf("file '%s': %v\n", fileid.String(), v)
}
