package utils

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	"os"
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
	process, _ := process.NewProcess(int32(os.Getpid()))
	cpuPercent, _ := process.CPUPercent()

	last1Ms, _ := cpu.Percent(100*time.Millisecond, false)
	fmt.Printf("Cpu: %f. Process: %f\n", last1Ms[0], cpuPercent)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
