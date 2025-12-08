package lib

import (
	"os"
	"runtime/pprof"
)

func CpuProfile() func() {
	if p := os.Getenv("PROFILE_CPU"); p != "" {
		f, err := os.Create(p)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		return func() {
			pprof.StopCPUProfile()
			f.Close()
		}
	}

	return func() {}
}
