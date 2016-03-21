package main

import (
	"time"
	"fmt"
	"github.com/lkj01010/log"
	"github.com/couchbase/gocb"
//	"encoding/json"
)


type User struct {
	Type     string `json:"_type"`
	ID       string `json:"_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Flights  []UserFlight `json:"flights"`
}


type UserFlight struct {
	Type               string `json:"_type"`
	ID                 string `json:"_id"`
	Name               string `json:"name"`
	Flight             string `json:"flight"`
	Date               string `json:"date"`
	Sourceairport      string `json:"sourceairport"`
	Destinationairport string `json:"destinationairport"`
	Bookedon           string `json:"bookedon"`
}



type AirportIntermediary struct {
//	ToAirport   string `json:'toAirport,omitempty'`
//	FromAirport string `json:'fromAirport,omitempty'`
//	Geo         struct {
//					Alt int `json:"alt"`
//					Lat float64 `json:"lat"`
//					Lon float64 `json:"lon"`
//				} `json:"geo"`
	Airportname string `json:"airportname"`
	City string `json:"city"`
	Country string `json:"country"`
	Type string `json:"type"`
	Tz string `json:"tz"`
	Faa interface{} `json:"faa"`
	Icao string `json:"icao"`
	Id string `json:"id"`
}

type TravelSample struct {
	TS string `json:"travel-sample"`
}


var bucket *gocb.Bucket
func main() {
	//	go func() {
	//		cluster1, e := gocb.Connect("couchbase://127.0.0.1")
	//		log.Debugf("connect e=%+v", e)
	//		if e != nil {
	//			fmt.Errorf("%+v", e.Error())
	//		}
	//		bucket1, e := cluster1.OpenBucket("travel-sample", "")
	//
	//
	//		N := 50000
	//		t1 := time.Now()
	//
	//		for i := 0; i < N; i++ {
	//			var curUser User
	//			if _, err := bucket1.Get("lkj01010@163.com", &curUser); err != nil {
	//				fmt.Print("bucket get err=", err.Error())
	//				return
	//			}else {
	//				//			fmt.Printf("%+v", curUser)
	//			}
	//		}
	//		t2 := time.Now()
	//		d := t2.Sub(t1)
	//
	//		log.Infof("BenchmarkGo go, %+v times in %+v", N, d)
	//
	//	}()
	//
	//	go func() {
	//		cluster1, e := gocb.Connect("couchbase://127.0.0.1")
	//		log.Debugf("connect e=%+v", e)
	//		if e != nil {
	//			fmt.Errorf("%+v", e.Error())
	//		}
	//		bucket1, e := cluster1.OpenBucket("travel-sample", "")
	//
	//
	//		N := 50000
	//		t1 := time.Now()
	//
	//		for i := 0; i < N; i++ {
	//			var curUser User
	//			if _, err := bucket1.Get("lkj01010@163.com", &curUser); err != nil {
	//				fmt.Print("bucket get err=", err.Error())
	//				return
	//			}else {
	//				//			fmt.Printf("%+v", curUser)
	//			}
	//		}
	//		t2 := time.Now()
	//		d := t2.Sub(t1)
	//
	//		log.Infof("BenchmarkGo go, %+v times in %+v", N, d)
	//
	//	}()
	//
	//	go func() {
	//		cluster1, e := gocb.Connect("couchbase://127.0.0.1")
	//		log.Debugf("connect e=%+v", e)
	//		if e != nil {
	//			fmt.Errorf("%+v", e.Error())
	//		}
	//		bucket1, e := cluster1.OpenBucket("travel-sample", "")
	//
	//		N := 50000
	//		t1 := time.Now()
	//
	//		for i := 0; i < N; i++ {
	//			var curUser User
	//			if _, err := bucket1.Get("lkj01010@163.com", &curUser); err != nil {
	//				fmt.Print("bucket get err=", err.Error())
	//				return
	//			}else {
	//				//			fmt.Printf("%+v", curUser)
	//			}
	//		}
	//		t2 := time.Now()
	//		d := t2.Sub(t1)
	//
	//		log.Infof("BenchmarkGo go, %+v times in %+v", N, d)
	//
	//	}()
	//
	//	go func() {
	//		cluster1, e := gocb.Connect("couchbase://127.0.0.1")
	//		log.Debugf("connect e=%+v", e)
	//		if e != nil {
	//			fmt.Errorf("%+v", e.Error())
	//		}
	//		bucket1, e := cluster1.OpenBucket("travel-sample", "")
	//
	//
	//		N := 50000
	//		t1 := time.Now()
	//
	//
	//		for i := 0; i < N; i++ {
	//			var curUser User
	//			if _, err := bucket1.Get("lkj01010@163.com", &curUser); err != nil {
	//				fmt.Print("bucket get err=", err.Error())
	//				return
	//			}else {
	//				//			fmt.Printf("%+v", curUser)
	//			}
	//
	//		}
	//		t2 := time.Now()
	//		d := t2.Sub(t1)
	//
	//		log.Infof("BenchmarkGo go, %+v times in %+v", N, d)
	//
	//	}()

	cluster, e := gocb.Connect("couchbase://127.0.0.1")
//	cluster, e := gocb.Connect("couchbase://42.62.101.136")
	log.Debugf("connect e=%+v", e)
	if e != nil {
		fmt.Errorf("%+v", e.Error())
	}
	bucket, e = cluster.OpenBucket("travel-sample", "")
	if e != nil {
		fmt.Errorf("OpenBucket ERROR=%+v", e.Error())
	}
//
//	myQuery := gocb.NewN1qlQuery("select * from `travel-sample`")
	myQuery := gocb.NewN1qlQuery("select * from `travel-sample` where airportname IS NOT NULL limit 1")
	N := 100
	t1 := time.Now()

	//	for i := 0; i < N; i++ {



	rows, err := bucket.ExecuteN1qlQuery(myQuery, nil)
	if err != nil {
		fmt.Println("ERROR EXECUTING N1QL QUERY:", err)
	}
	var airports []AirportIntermediary
	var row AirportIntermediary	// 这里的查询结果要严格对应定义的格式,否则转出来的struct的内部值都是空值
	log.Infof("rows=%#v", rows)
	for rows.Next(&row) {
		airports = append(airports, row)
		log.Debugf("row=%+v", row)
	}
//	_ = rows.Close()
//	bytes, _ := json.Marshal(airports)
	log.Infof("airport = %+v", airports)


//	var row interface{}
//	for rows.Next(&row) {
//		log.Infof("Row: %#v\n", row)
//	}
	if err := rows.Close(); err != nil {
		fmt.Printf("N1QL query error: %s\n", err)
	}




	//	}
	t2 := time.Now()
	d := t2.Sub(t1)

	log.Infof("BenchmarkGo, %+v times in %+v", N, d)


}

/**
结论: (本地127.0.0.1访问)
1. get(JSON)单核7000op/s左右
2. 4~5核20000+op/s
3. n1ql 单核 select *
from `travel-sample`
where airportname IS NOT NULL
limit(1) -> 22ms
limit(大量) -> 350ms    2k op/s 扫描速度
 */
