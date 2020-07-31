package standalone_storage

import (
	"github.com/pingcap-incubator/tinykv/kv/config"
	"github.com/pingcap-incubator/tinykv/kv/storage"
	"github.com/pingcap-incubator/tinykv/kv/util/engine_util"
	"github.com/pingcap-incubator/tinykv/proto/pkg/kvrpcpb"
	badger "github.com/Connor1996/badger"

)

// StandAloneStorage is an implementation of `Storage` for a single-node TinyKV instance. It does not
// communicate with other nodes and all data is stored locally.
type StandAloneStorage struct {
	// Your Data Here (1).
	Config *config.Config
	Engine	engine_util.Engines

}

func NewStandAloneStorage(conf *config.Config) *StandAloneStorage {
	// Your Code Here (1).
	engine := engine_util.NewEngines()
	return  &StandAloneStorage{
		Config: conf,

	}
	return nil
}

func (s *StandAloneStorage) Start() error {
	// Your Code Here (1).
	return nil
}

func (s *StandAloneStorage) Stop() error {
	// Your Code Here (1).
	return nil
}

func (s *StandAloneStorage) Reader(ctx *kvrpcpb.Context) (storage.StorageReader, error) {
	// Your Code Here (1).
	return nil, nil
}

func (s *StandAloneStorage) Write(ctx *kvrpcpb.Context, batch []storage.Modify) error {
	// Your Code Here (1).
	return nil
}

type StandAloneStorageReader struct {

}

func (sr *StandAloneStorageReader) GetCF(cf string, key []byte) ([]byte, error) {

}
func (sr *StandAloneStorageReader) IterCF(cf string) engine_util.DBIterator {

}
func (sr *StandAloneStorageReader) Close() {

}
