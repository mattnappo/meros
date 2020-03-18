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

	/* -- PUT -- */

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

	/* -- GET -- */

	readFile, err := filedb.GetItem(fid)
	if err != nil {
		t.Logf("%v", file)
		t.Fatal(err)
	}

	v := readFile.(*types.File)
	t.Logf("file '%s': %v\n", fid.String(), v)
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
