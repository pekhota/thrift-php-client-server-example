package main

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import (
	"context"
	"fmt"
	"go-client/tutorial"
	"log"

	"github.com/apache/thrift/lib/go/thrift"

)

var defaultCtx = context.Background()

func handleClient(client *tutorial.CalculatorClient) (err error) {
	//client.Ping(defaultCtx)
	//fmt.Println("ping()")

	sum, err := client.Add(defaultCtx, 1, 1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("1+1=", sum, "\n")

	//work := tutorial.NewWork()
	//work.Op = tutorial.Operation_DIVIDE
	//work.Num1 = 1
	//work.Num2 = 0
	//quotient, err := client.Calculate(defaultCtx, 1, work)
	//if err != nil {
	//	switch v := err.(type) {
	//	case *tutorial.InvalidOperation:
	//		fmt.Println("Invalid operation:", v)
	//	default:
	//		fmt.Println("Error during operation:", err)
	//	}
	//} else {
	//	fmt.Println("Whoa we can divide by 0 with new value:", quotient)
	//}
	//
	//work.Op = tutorial.Operation_SUBTRACT
	//work.Num1 = 15
	//work.Num2 = 10
	//diff, err := client.Calculate(defaultCtx, 1, work)
	//if err != nil {
	//	switch v := err.(type) {
	//	case *tutorial.InvalidOperation:
	//		fmt.Println("Invalid operation:", v)
	//	default:
	//		fmt.Println("Error during operation:", err)
	//	}
	//	return err
	//} else {
	//	fmt.Print("15-10=", diff, "\n")
	//}
	//
	//log, err := client.GetStruct(defaultCtx, 1)
	//if err != nil {
	//	fmt.Println("Unable to get struct:", err)
	//	return err
	//} else {
	//	fmt.Println("Check log:", log.Value)
	//}
	return err
}

func makeClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool, cfg *thrift.TConfiguration) (*tutorial.CalculatorClient, thrift.TTransport, error) {
	var transport thrift.TTransport
	if secure {
		var err error
		transport, err = thrift.NewTSSLSocketConf(addr, cfg)
		if err != nil {
			return nil, nil, err
		}
	} else {
		var err error
		transport, err = thrift.NewTHttpClientWithOptions(addr, thrift.THttpClientOptions{})
		if err != nil {
			return nil, nil, err
		}
	}
	transport, err := transportFactory.GetTransport(transport)
	if err != nil {
		return nil, nil, err
	}
	//defer transport.Close()
	if err := transport.Open(); err != nil {
		return nil, nil, err
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	client := tutorial.NewCalculatorClient(thrift.NewTStandardClient(iprot, oprot))
	return client, transport, err
	//return handleClient(tutorial.NewCalculatorClient(thrift.NewTStandardClient(iprot, oprot)))
}
