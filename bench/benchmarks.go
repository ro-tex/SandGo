package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	netpprof "net/http/pprof"
	"os"
	"runtime/pprof"
	"sync"
	"sync/atomic"

	"gitlab.com/NebulousLabs/fastrand"
)

// fr generates a random number using fastrand
func fr() uint64 {
	return fastrand.Uint64n(1e9)
}

// rr generates a random number using math.rand
func rr() uint64 {
	return uint64(rand.Int63n(1e9))
}

// s makes a single select on a channel
func s(ch chan struct{}) {
	select {
	case <-ch:
	default:
	}
}

func atomicCount() uint64 {
	var ops uint64
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return ops
}

func mutexCount() uint64 {
	var ops uint64
	var opsMu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				opsMu.Lock()
				ops += 1
				opsMu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return ops
}

/* Disable CPU frequency scaling:
$ sudo bash
# for i in /sys/devices/system/cpu/cpu[0-7]
do
echo performance > $i/cpufreq/scaling_governor
done
#
*/

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var web = flag.String("web", "", "start a web server at :6060")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		_ = pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	for i := 0; i < 1e7; i++ {
		_ = fr()
		_ = rr()
	}

	if *web != "" {
		go func() {
			fmt.Println(
				http.ListenAndServe(
					"localhost:6060",
					http.HandlerFunc(netpprof.Index),
				),
			)
		}()
		// go tool pprof http://localhost:6060/debug/pprof/heap
	}
}
