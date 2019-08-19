package data

// PROBLEM
// The code below doesn't compile.
// For some reason usage of github.com/influxdata/influxdb1-client/v2
// breaks compilation.
//
// Speculation
// It may have something to do with cgo, but I'm not savvy enough to know for sure

//import (
//	client "github.com/influxdata/influxdb1-client/v2"
//	"testing"
//	"time"
//)
//
//type ClientMock struct {}
//
//func (c ClientMock) Ping(timeout time.Duration) (time.Duration, string, error) {
//	panic("implement me")
//}
//
//func (c ClientMock) Write(bp client.BatchPoints) error {
//	panic("implement me")
//}
//
//func (c ClientMock) Query(q client.Query) (*client.Response, error) {
//	panic("implement me")
//}
//
//func (c ClientMock) QueryAsChunk(q client.Query) (*client.ChunkedResponse, error) {
//	panic("implement me")
//}
//
//func (c ClientMock) Close() error {
//	panic("implement me")
//}
//
//func TestRequestChannel(t *testing.T) {
//	var _, _ = requestChannel(func() (client client.Client, e error) {
//		return ClientMock{}, nil
//	})
//}
