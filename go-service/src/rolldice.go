package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"net/http"

	//	"sort"
	"strconv"
)

func rolldice(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	rolls, err := strconv.Atoi(r.URL.Query().Get("rolls"))
	if err != nil {
		rolls = 1
	}

	/*
		dsn := "root:@tcp(pinba:3306)/pinba"
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

	*/

	resp := ""

	for rolls > 0 {
		roll := rollonce(ctx)
		resp = resp + strconv.Itoa(roll) + "\n"
		rolls--
	}

	if _, err := io.WriteString(w, resp); err != nil {
		log.Printf("Write failed: %v\n", err)
	}
}

func rollonce(ctx context.Context) int {
	/*
		arr := make([]int, 1_000_000)
		for i := range arr {
			arr[i] = 1_000_000 - i
		}
		sort.Ints(arr)

	*/

	req, err := http.NewRequestWithContext(ctx, "GET", "http://echo-service:8088/payload?io_msec=10", nil)
	if err != nil {
		log.Printf("Create request failed: %v\n", err)
	}

	client := http.DefaultClient
	res, err := client.Do(req)
	//	httpresp, err := client.Get("http://echo-service:8088/payload?io_msec=10")
	if err != nil {
		log.Printf("Get failed: %v\n", err)
	}
	res.Body.Close()

	/*
		var now string

		err = db.QueryRowContext(ctx, "select now()").Scan(&now)
		log.Printf("Get mysql: %s\n", now)

		if err != nil {
			panic(err.Error())
		}


	*/
	return rand.Intn(6) + 1
}
