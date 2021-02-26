package utils

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"runtime"
	"time"
)

func PrintMemoryUsage() {
	v, _ := mem.VirtualMemory()
	fmt.Printf("Total = %vMiB, Free = %vMiB, UsedPercent = %f\n", bToMb(v.Total), bToMb(v.Free), v.UsedPercent)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("TotalAlloc = %v MiB\tAlloc = %v MiB\tSys = %v MiB\n", bToMb(m.TotalAlloc), bToMb(m.Alloc), bToMb(m.Sys))
}

func PrintCpuUsage() {
	last1Ms, _ := cpu.Percent(100*time.Millisecond, false)
	fmt.Printf("Cpu: %f\n", last1Ms[0])
}

func PrintAvgUsage() {
	avg, _ := load.Avg()
	fmt.Printf("1 min = %f\t5 min = %f\t15 min = %f\n", avg.Load1, avg.Load5, avg.Load15)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
