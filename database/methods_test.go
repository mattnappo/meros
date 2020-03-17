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

	v := file.(*types.File)
	t.Logf("file '%s': %v\n", fileid.String(), v)
}

func TestPutShard(t *testing.T) {
	sharddb, err := Open("mynodesharddb", NSHARDDB)
	if err != nil {
		t.Fatal(err)
	}
	defer sharddb.Close()

	shard, err := types.NewShard([]byte("I represent some bytes in a shard"))
	if err != nil {
		t.Fatal(err)
	}

	sid, err := sharddb.PutItem(*shard)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("shardID: %s\n", sid.String())
}

func TestGetShard(t *testing.T) {
	sharddb, err := Open("mynodesharddb", NSHARDDB)
	if err != nil {
		t.Fatal(err)
	}
	defer sharddb.Close()

	sid, err := IDFromString("f910670bbb0012b2eb6c4f321a02251d2f38c97a8c28acdda16bb0a3b79c1ab5")
	if err != nil {
		t.Fatal(err)
	}

	shard, err := sharddb.GetItem(sid)
	if err != nil {
		t.Fatal(err)
	}

	v := shard.(*types.Shard)
	t.Logf("shard '%s': %v\n", sid, v)
}
