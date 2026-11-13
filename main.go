package main

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

const (
	simplePattern  = `\W`
	complexPattern = `^(\d{4})-(\d{2})-(\d{2})[T ](\d{2}):(\d{2}):(\d{2})(?:\.(\d+))?(?:Z|[+-]\d{2}:\d{2})?\s+\[(ERROR|WARN|INFO|DEBUG)\]\s+(.*)$`
	emailPattern   = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
)

var (
	simpleRE  = regexp.MustCompile(simplePattern)
	complexRE = regexp.MustCompile(complexPattern)
	emailRE   = regexp.MustCompile(emailPattern)
)

var (
	sinkString   string
	sinkSubmatch []string
	sinkBool     bool
)

type BenchResult struct {
	Name        string
	Iters       int
	NsPerOp     float64
	BytesPerOp  int64
	AllocsPerOp int64
}

func run(name string, fn func()) BenchResult {
	r := testing.Benchmark(func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			fn()
		}
	})
	return BenchResult{
		Name:        name,
		Iters:       r.N,
		NsPerOp:     float64(r.NsPerOp()),
		BytesPerOp:  r.AllocedBytesPerOp(),
		AllocsPerOp: r.AllocsPerOp(),
	}
}

// printTable uses go-pretty to render the comparison table.
func printTable(pairs [][2]BenchResult) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{
		"Benchmark",
		"iters",
		"ns/op",
		"B/op",
		"allocs/op",
		"time gain",
		"alloc gain",
	})

	// Right-align numeric columns; left-align the benchmark name.
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignLeft},  // Benchmark
		{Number: 2, Align: text.AlignRight}, // iters
		{Number: 3, Align: text.AlignRight}, // ns/op
		{Number: 4, Align: text.AlignRight}, // B/op
		{Number: 5, Align: text.AlignRight}, // allocs/op
		{Number: 6, Align: text.AlignRight}, // time gain
		{Number: 7, Align: text.AlignRight}, // alloc gain
	})

	for _, pair := range pairs {
		naive, cached := pair[0], pair[1]

		timeGain := naive.NsPerOp / cached.NsPerOp
		allocGain := float64(naive.AllocsPerOp) / float64(cached.AllocsPerOp)

		allocGainStr := fmt.Sprintf("%.1fx", allocGain)
		if cached.AllocsPerOp == 0 {
			allocGainStr = "∞"
		}

		// Render the naive row (first of the pair).
		t.AppendRow(table.Row{
			naive.Name,
			naive.Iters,
			fmt.Sprintf("%.2f", naive.NsPerOp),
			naive.BytesPerOp,
			naive.AllocsPerOp,
			fmt.Sprintf("%.1fx", timeGain),
			allocGainStr,
		})

		// Render the cached row (second of the pair).
		t.AppendRow(table.Row{
			cached.Name,
			cached.Iters,
			fmt.Sprintf("%.2f", cached.NsPerOp),
			cached.BytesPerOp,
			cached.AllocsPerOp,
			fmt.Sprintf("%.1fx", timeGain),
			allocGainStr,
		})

		t.AppendSeparator()
	}

	t.Render()
}

func main() {
	fmt.Println("Regexp: compile-per-request vs precompiled (cached)")

	simpleNaive := run("compile per req (simple)", func() {
		re := regexp.MustCompile(simplePattern)
		sinkString = re.ReplaceAllString(" the string ", "")
	})
	simpleCached := run("precompiled (simple)", func() {
		sinkString = simpleRE.ReplaceAllString(" the string ", "")
	})

	complexNaive := run("compile per req (complex)", func() {
		re := regexp.MustCompile(complexPattern)
		sinkSubmatch = re.FindStringSubmatch("2025-01-15 12:34:56 [ERROR] connection timeout")
	})
	complexCached := run("precompiled (complex)", func() {
		sinkSubmatch = complexRE.FindStringSubmatch("2025-01-15 12:34:56 [ERROR] connection timeout")
	})

	emailNaive := run("compile per req (email)", func() {
		re := regexp.MustCompile(emailPattern)
		sinkBool = re.MatchString("alice.smith+tag@example.co.uk")
	})
	emailCached := run("precompiled (email)", func() {
		sinkBool = emailRE.MatchString("alice.smith+tag@example.co.uk")
	})

	compileOnly := run("compile only (complex)", func() {
		regexp.MustCompile(complexPattern)
	})
	matchOnly := run("match only (precompiled)", func() {
		sinkSubmatch = complexRE.FindStringSubmatch("2025-01-15 12:34:56 [ERROR] connection timeout")
	})

	printTable([][2]BenchResult{
		{simpleNaive, simpleCached},
		{complexNaive, complexCached},
		{emailNaive, emailCached},
		{compileOnly, matchOnly},
	})
	fmt.Println(systemInfo())
}

func systemInfo() string {
	cpuModel := "unknown"
	if data, err := os.ReadFile("/proc/cpuinfo"); err == nil {
		for _, line := range strings.Split(string(data), "\n") {
			if strings.HasPrefix(line, "model name") {
				if parts := strings.SplitN(line, ":", 2); len(parts) == 2 {
					cpuModel = strings.TrimSpace(parts[1])
				}
				break
			}
		}
	}

	return fmt.Sprintf(
		"System: %s/%s\nCPU: %s\nCores: %d\nGOMAXPROCS: %d\nGo: %s",
		runtime.GOOS, runtime.GOARCH, cpuModel,
		runtime.NumCPU(), runtime.GOMAXPROCS(0), runtime.Version(),
	)
}
