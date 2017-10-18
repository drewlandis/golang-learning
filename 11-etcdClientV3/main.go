package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/coreos/etcd/clientv3"

	"golang.org/x/net/context"
)

/* start etcd container
   cd /d3/hyperdark/testing/etcd
   docker-compose up
	 cd /learning/golang-learning/11-etcdClientV3
	 go run main.go
*/

func main() {
	cli, err := clientv3.New(clientv3.Config{
		//Endpoints:   []string{"0.0.0.0:23791"}, // /d3/hyperdark/testing/etcd
		Endpoints:   []string{"0.0.0.0:2379"}, // test environment
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("err 1")
		fmt.Println(err)
		// handle error!
	}
	defer cli.Close()

	timeout := 5 * time.Second

	// Put
	/*
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		_, err = cli.Put(ctx, "high-threat-score-foo3", "bar3")
		cancel()
		if err != nil {
			fmt.Println("Put err:")
			fmt.Println(err)
		}
	*/
	/*
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		//pResp, err := cli.Put(ctx, "sample_key", "sample_value")
		//pResp, err := cli.Put(ctx, "3/sample_key2", "3_sample_value2")
		//_, err = cli.Put(ctx, "settingsFileConverted", "notABoolValue")
		//_, err = cli.Put(ctx, "test-byte", "string")
		cancel()
		if err != nil {
			fmt.Println("err 2")
			fmt.Println(err)
			// handle error!
			switch err {
			case context.Canceled:
				log.Fatalf("ctx is canceled by another routine: %v", err)
			case context.DeadlineExceeded:
				log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
			case rpctypes.ErrEmptyKey:
				log.Fatalf("client-side error: %v", err)
			default:
				log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
			}
		}
		// use the response
		//fmt.Println("pResp")
		//fmt.Println(pResp)
	*/

	// Get
	ctx2, cancel2 := context.WithTimeout(context.Background(), timeout)
	//gResp, err := cli.Get(ctx2, "3", clientv3.WithPrefix())
	//gResp, err := cli.Get(ctx2, "dab03e56-c7a8-11e6-a4e1-33ec88ecbc05-", clientv3.WithPrefix())
	gResp, err := cli.Get(ctx2, "", clientv3.WithPrefix())
	cancel2()
	if err != nil {
		fmt.Println("err 3")
		fmt.Println(err)
		// handle error!
	}
	fmt.Println("gResp")
	//fmt.Println(string(gResp.Kvs[0].Value))

	//fmt.Println(gResp.Kvs)
	//fmt.Println("gResp.Count")
	//fmt.Println(gResp.Count)
	for _, ev := range gResp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
	fmt.Println("len:" + strconv.Itoa(len(gResp.Kvs)))

	/* * /
	// Delete
	ctx3, cancel3 := context.WithTimeout(context.Background(), timeout)
	//dResp, err := cli.Delete(ctx3, "settingsFileConverted")
	//fmt.Println("dResp.Deleted settingsFileConverted")
	//fmt.Println(dResp.Deleted)
	dResp, err := cli.Delete(ctx3, "high-threat-score-", clientv3.WithPrefix())
	//dResp, err = cli.Delete(ctx3, "dab03e56-c7a8-11e6-a4e1-33ec88ecbc05", clientv3.WithPrefix())
	fmt.Println("dResp.Deleted applianceUUID")
	fmt.Println(dResp.Deleted)
	cancel3()
	/* */

	fmt.Println("\n")
}
