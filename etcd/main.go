package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/coreos/etcd/clientv3"

	"google.golang.org/grpc/grpclog"
	"fmt"
	"github.com/wolfeidau/unflatten"
	"strings"
	"encoding/json"
	"github.com/jeremywohl/flatten"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 10 * time.Second
	endpoints      = []string{"localhost:2379"}
	username       = "root"
	password       = "root"
)

func main() {

	clientv3.SetLogger(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))

	start1 := time.Now()

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
		Username:    username,
		Password:    password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close() // make sure to close the client

	elapsed1 := time.Since(start1)
	log.Printf("Binomial took %s", elapsed1)

	start2 := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.Get(ctx, "_config/hefesto/level1a/level2a/level3a/key3a")
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}

	elapsed2 := time.Since(start2)
	log.Printf("Binomial took %s", elapsed2)

	start3 := time.Now()

	m := make(map[string]interface{})

	ctx1, cancel1 := context.WithTimeout(context.Background(), requestTimeout)
	resp1, err := cli.Get(ctx1, "_revision/hefesto", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	cancel1()
	if err != nil {
		log.Fatal(err)
	}
	for _, ev := range resp1.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
		fields := strings.Split(string(ev.Key), "/")
		key := fields[len(fields)-1]
		m[key] = string(ev.Value)
	}

	tree := unflatten.Unflatten(m, func(k string) []string { return strings.Split(k, ".") })

	unflat, _ := json.Marshal(tree)

	log.Printf("JSON: %s", unflat)

	elapsed3 := time.Since(start3)
	log.Printf("Binomial took %s", elapsed3)

	/*
	t := map[string]interface{}{
		"a": "b",
		"c": map[string]interface{}{
			"d": "e",
			"f": "g",
		},
		"z": 1.4567,
	}
	flat, err := flatten.Flatten(t, "", flatten.DotStyle)
	*/

	//flat, err := flatten.FlattenString(string(unflat), "", flatten.DotStyle)

	flat, err := flatten.Flatten(tree, "", flatten.DotStyle)

	jsonString, _ := json.Marshal(flat)

	log.Printf("FLAT: %s", jsonString)

}
