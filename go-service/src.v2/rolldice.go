package main

import (
	"context"
	"database/sql"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"

	"strconv"

	"go.opentelemetry.io/otel"

	_ "github.com/go-sql-driver/mysql"
)

var (
	tracer = otel.Tracer("otel-go-example")
)

func rolldice(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	rolls, err := strconv.Atoi(r.URL.Query().Get("rolls"))
	if err != nil {
		rolls = 1
	}

	load := r.URL.Query().Get("load")

	resp := ""

	for rolls > 0 {
		roll := rollonce(ctx, load)
		resp = resp + strconv.Itoa(roll) + "\n"
		rolls--
	}

	if _, err := io.WriteString(w, resp); err != nil {
		log.Printf("Write failed: %v\n", err)
	}
}

func rollonce(ctx context.Context, load string) int {

	if strings.Contains(load, "C") {
		arr := make([]int, 1_000_000)
		for i := range arr {
			arr[i] = 1_000_000 - i
		}
		sort.Ints(arr)
	}

	if strings.Contains(load, "E") {

		req, err := http.NewRequestWithContext(ctx, "GET", "http://echo-service:8088/payload?io_msec=10", nil)
		if err != nil {
			log.Printf("Create request failed: %v\n", err)
			return 0
		}

		client := http.Client{Timeout: time.Duration(1) * time.Second}
		res, err := client.Do(req)
		if err != nil {
			log.Printf("Get failed: %v\n", err)
			return 0
		}
		res.Body.Close()
	}

	if strings.Contains(load, "D") {
		dsn := "root:@tcp(pinba:3306)/pinba"
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		var now string

		err = db.QueryRowContext(ctx, "select now()").Scan(&now)
		if err != nil {
			panic(err.Error())
		}

	}

	_, span := tracer.Start(ctx, "getRandom")
	random := rand.Intn(6) + 1
	span.End()

	return random
}
