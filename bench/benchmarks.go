package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"

	"gitlab.com/NebulousLabs/fastrand"
)

func fr() uint64 {
	return fastrand.Uint64n(1e9)
}

func rr() uint64 {
	return uint64(rand.Int63n(1e9))
}

func s(ch chan struct{}) {
	select {
	case <-ch:
	default:
	}
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
}
