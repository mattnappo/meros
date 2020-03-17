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

	fileid, err := IDFromString("0c644a9a8745e7c160af8fa985801d0fabdcaf0627aa8ba2bb2f11ab1a0f8ee9")
	if err != nil {
		t.Fatal(err)
	}

	file, err := filedb.GetItem(fileid)
	if err != nil {
		t.Logf("%v", file)
		t.Fatal(err)
	}

	// t.Logf("file '%s': %v\n", fileid.String(), *file)
}
