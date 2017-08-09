package goroutines

import (
	"errors"
	"os"
	"sync"
)

type Data []byte

type DataFile interface {
	Read() (rsn int64, d Data, err error)
	Write(d Data) (wsn int64, err error)
	RSN() int64
	WSN() int64
	DataLen() uint32
	Close() error
}

type dataFile struct {
	f       *os.File
	fmutex  sync.RWMutex // 文件读写锁
	woffset int64
	roffset int64
	wm      sync.Mutex
	rm      sync.Mutex
	dataLen uint32
}

func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)

	if err != nil {
		return nil, err
	}

	if dataLen == 0 {
		return nil, errors.New("invalid data length")
	}

	df := &dataFile{f: f, dataLen: dataLen}
	return df, nil
}

func (df *dataFile) Read() (rsn int64, d Data, err error) {
	var offset int64

	df.rm.Lock()
	offset = df.roffset
	df.roffset += int64(df.dataLen)
	df.rm.Unlock()

	rsn = offset / int64(df.dataLen)

	df.fmutex.RLock()
	defer df.fmutex.RUnlock()

	bytes := make([]byte, df.dataLen)
	_, err = df.f.ReadAt(bytes, offset)
	if err != nil {
		return
	}

	d = bytes
	return
}

func (df *dataFile) Write(d Data) (wsn int64, err error) {
	var offset int64
	df.wm.Lock()
	offset = df.woffset
	df.woffset += int64(df.dataLen)
	df.wm.Unlock()

	wsn = offset / int64(df.dataLen)
	var bytes []byte

	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}

	df.fmutex.Lock()
	defer df.fmutex.Unlock()

	_, err = df.f.Write(bytes)
	return
}

func (df *dataFile) RSN() int64 {
	df.rm.Lock()
	defer df.rm.Unlock()
	return df.roffset / int64(df.dataLen)
}

func (df *dataFile) WSN() int64 {
	df.wm.Lock()
	defer df.wm.Unlock()
	return df.woffset / int64(df.dataLen)
}

func (df *dataFile) DataLen() uint32 {
	return df.dataLen
}

func (df *dataFile) Close() error {
	return df.f.Close()
}
