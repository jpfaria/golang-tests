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
	"github.com/imdario/mergo"
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
	fmt.Printf("Binomial took %s\n", elapsed1)

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
	fmt.Printf("Binomial took %s\n", elapsed2)

	start3 := time.Now()

	map1 := make(map[string]interface{})

	ctx1, cancel1 := context.WithTimeout(context.Background(), requestTimeout)
	resp1, err := cli.Get(ctx1, "_config/hefesto", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	cancel1()
	if err != nil {
		log.Fatal(err)
	}

	elapsed3 := time.Since(start3)
	fmt.Printf("Binomial took %s\n", elapsed3)

	for _, ev := range resp1.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
		fields := strings.Split(string(ev.Key), "/")
		key := fields[len(fields)-1]
		map1[key] = string(ev.Value)
	}

	tree1 := unflatten.Unflatten(map1, func(k string) []string { return strings.Split(k, ".") })

	unflat1, _ := json.Marshal(tree1)

	fmt.Printf("JSON: %s\n", unflat1)


	/*
	t := map[string]interface{}{
		"a": "b",
		"c": map[string]interface{}{
			"d": "e",
			"f": "g",
		},
		"z": 1.4567,
	}
	flat1, err := flatten.Flatten(t, "", flatten.DotStyle)
	*/

	//flat1, err := flatten.FlattenString(string(unflat1), "", flatten.DotStyle)

	flat1, err := flatten.Flatten(tree1, "", flatten.DotStyle)

	json1, _ := json.Marshal(flat1)

	fmt.Printf("FLAT: %s\n", json1)

	map2 := make(map[string]interface{})

	ctx2, cancel2 := context.WithTimeout(context.Background(), requestTimeout)
	resp2, err := cli.Get(ctx2, "_revision/hefesto", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	cancel2()
	if err != nil {
		log.Fatal(err)
	}

	for _, ev := range resp2.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
		fields := strings.Split(string(ev.Key), "/")
		key := fields[len(fields)-1]
		map2[key] = string(ev.Value)
	}

	tree2 := unflatten.Unflatten(map2, func(k string) []string { return strings.Split(k, ".") })

	unflat2, _ := json.Marshal(tree2)

	fmt.Printf("JSON: %s\n", unflat2)

	flat2, err := flatten.Flatten(tree2, "", flatten.DotStyle)

	json2, _ := json.Marshal(flat2)

	fmt.Printf("FLAT: %s\n", json2)

	mergo.Merge(&tree2, tree1)

	json3, _ := json.Marshal(tree2)

	fmt.Printf("MERGED: %s\n", json3)


}
