package main

import (
	"context"
	"crypto/tls"
	"github.com/apache/thrift/lib/go/thrift"
	"go-client/load"
)

func main() {
	load.DoLoad(func() (int64, error) {
		secure := new(bool); *secure = false
		addr   := new(string); *addr = "http://nginx:80/PhpServer.php"

		//var protocolFactory thrift.TProtocolFactory
		protocolFactory := thrift.NewTBinaryProtocolFactoryConf(nil)

		var transportFactory thrift.TTransportFactory
		cfg := &thrift.TConfiguration{
			TLSConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}

		transportFactory = thrift.NewTBufferedTransportFactory(8192)
		client, transport, err := makeClient(transportFactory, protocolFactory, *addr, *secure, cfg);
		if err != nil {
			return 0, err
		}
		defer transport.Close()

		defaultCtx := context.Background()
		sum, err := client.Add(defaultCtx, 1, 1)
		if err != nil {
			return 0, err
		}

		return int64(sum), err
	})

	load.PrintResults()
}