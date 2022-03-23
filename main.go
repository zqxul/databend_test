package main

import (
	"database/sql"
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"databend.demo/db"
)

type Arg struct {
	H   bool
	Br  int
	Tr  int
	Tn  int
	Mtn int
}

var arg = Arg{}

func init() {
	flag.BoolVar(&arg.H, "h", false, "usage info")
	flag.IntVar(&arg.Br, "br", 1000, "batch rows")
	flag.IntVar(&arg.Tr, "tr", 60000*1, "total rows")
	flag.IntVar(&arg.Tn, "tn", 100, "thread nums")
	flag.IntVar(&arg.Mtn, "mtn", 100, "max thread nums")
}

// Command: ./main -Host 127.0.0.1 -Port 3307 -u db -br 1000 -tr 60000 -tn 1 -mtn 1
func main() {
	flag.Parse()
	if arg.H == true {
		fmt.Println("usage info: \n -h:\tshow help \n -br:\tbatch rows\n -tr:\ttotal rows\n -tn:\tthread nums\n -mtn:\tmax thread nums")
		return
	}
	mdb := db.Open(arg.Mtn)
	fmt.Printf("%+v\n", arg)
	Insert(mdb)
}

var (
	src = []byte("abcdefghijklmnopqrstuvwxyz")
)

func randStr(l int) string {
	if l <= 0 {
		return ""
	}
	str := make([]byte, 0)
	len := len(src)
	for i := 0; i < l; i++ {
		randIndex := rand.Intn(len)
		str = append(str, src[randIndex])
	}
	return string(str)
}

func InsertBatch(mdb *sql.DB, n int) {
	sql := "INSERT INTO db.t(id1,id2,id3,id4,id5,id6,id7,id8,id9,id10,id11,id12,id13,id14,id15,id16,str1,str2,str3,dt) VALUES"
	values := make([]string, 0)
	s := make([]db.T, 0)
	for i := 0; i < n; i++ {
		s = append(s, db.T{rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000),
			rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), fmt.Sprintf("'%s'", randStr(20)), fmt.Sprintf("'%s'", randStr(20)), fmt.Sprintf("'%s'", randStr(40)), time.Now()})
	}
	var finalSql = ""
	for _, t := range s {
		val := fmt.Sprintf("(%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%s,%s,%s,%s)", t.Id1, t.Id2, t.Id3, t.Id4, t.Id5, t.Id6, t.Id7, t.Id8, t.Id9, t.Id10, t.Id11, t.Id12, t.Id13, t.Id14, t.Id15, t.Id16, t.Str1, t.Str2, t.Str3, fmt.Sprintf(`'%s'`, t.Date.UTC().Format("2006-01-02 15:04:05")))
		values = append(values, val)
		finalSql = sql + strings.Join(values, ",")
	}
	startTime := time.Now()
	mdb.Exec(finalSql)
	fmt.Printf("分批耗时\t%f 秒\n", time.Now().Sub(startTime).Seconds())
}

func Insert(mdb *sql.DB) {
	var wg sync.WaitGroup
	wg.Add(arg.Tn)
	startTime := time.Now()
	fmt.Printf("执行开始\t%s\n", startTime.UTC().Format("2006-01-02 15:04:05"))
	for i := 0; i < arg.Tn; i++ {
		go func() {
			start := 0
			if arg.Br < arg.Tr {
				for ; start < arg.Tr; start += arg.Br {
					InsertBatch(mdb, arg.Br)
				}
				InsertBatch(mdb, arg.Tr-start)
			} else {
				InsertBatch(mdb, arg.Tr)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("总耗时\t\t%f 秒 \n", time.Now().Sub(startTime).Seconds())
}
