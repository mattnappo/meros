package filedb

import (
	"testing"

	"github.com/xoreo/meros/types"
)

func TestPutFile(t *testing.T) {
	filedb, err := Open("myfiledb")
	if err != nil {
		t.Fatal(err)
	}
	defer filedb.Close()

	file, err := types.NewFile("test_file.txt")
	if err != nil {
		t.Fatal(err)
	}

	fid, err := filedb.PutFile(*file)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("fileID: %s\n", fid.String())
}

func TestGetFile(t *testing.T) {
	filedb, err := Open("myfiledb")
	if err != nil {
		t.Fatal(err)
	}
	defer filedb.Close()

	fileid, err := FileIDFromString("0c644a9a8745e7c160af8fa985801d0fabdcaf0627aa8ba2bb2f11ab1a0f8ee9")
	if err != nil {
		t.Fatal(err)
	}

	file, err := filedb.GetFile(fileid)
	if err != nil {
		t.Logf("%v", file)
		t.Fatal(err)
	}

	t.Logf("file '%s': %v\n", fileid.String(), *file)
}
