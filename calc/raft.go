package calc

import (
	"os"
	"time"

	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
)

func NewCalculator(dir, ip, id string) *Calculator {
	calc := new(Calculator)
	os.Mkdir(dir, 0777)
	config := raft.DefaultConfig()
	logStore, _ := raftboltdb.NewBoltStore(dir + "/raft.db")
	snapshotStore, _ := raft.NewFileSnapshotStore(dir, 1, os.Stdout)
	transport, _ := raft.NewTCPTransport(ip, nil, 5, 5*time.Second, os.Stdout)
	peerStore := raft.NewJSONPeers("", transport)
	calc.Raft, _ = raft.NewRaft(config, calc, logStore, logStore, snapshotStore, peerStore, transport)
	return calc
}
