package flag

import (
	"time"
)

var (
	WorkersFlag = Flag{
		Name:       "workers",
		ConfigName: "worker.workers",
		Value:      1,
		Usage:      "number of processing workers to spawn",
		Persistent: true,
	}
	TimeoutFlag = Flag{
		Name:       "timeout",
		ConfigName: "worker.timeout",
		Value:      10 * time.Minute,
		Usage:      "time allowed to complete scan",
		Persistent: true,
	}
	TimeoutFileMinimumFlag = Flag{
		Name:       "timeout-file-min",
		ConfigName: "worker.timeout-file-min",
		Value:      5 * time.Second,
		Usage:      "minimum timeout assigned to scanning file, this config superseeds timeout-second-per-bytes",
		Persistent: true,
	}
	TimeoutFileMaximumFlag = Flag{
		Name:       "timeout-file-max",
		ConfigName: "worker.timeout-file-max",
		Value:      300 * time.Second, // 5 mins
		Usage:      "maximum timeout assigned to scanning file, this config superseeds timeout-second-per-bytes",
		Persistent: true,
	}
	TimeoutFileSecondPerBytesFlag = Flag{
		Name:       "timeout-file-second-per-bytes",
		ConfigName: "worker.timeout-file-second-per-bytes",
		Value:      10 * 1000, // 10kb/s
		Usage:      "number of file size bytes producing a second of timeout assigned to scanning a file",
		Persistent: true,
	}
	TimeoutWorkerOnlineFlag = Flag{
		Name:       "timeout-worker-online",
		ConfigName: "worker.timeout-worker-online",
		Value:      60 * time.Second,
		Usage:      "maximum time for worker process to come online",
		Persistent: true,
	}
	FileSizeMaximumFlag = Flag{
		Name:       "file-size-max",
		ConfigName: "worker.file-size-max",
		Value:      25 * 1000 * 1000, // 25 MB
		Usage:      "ignore files with file size larger than this config",
		Persistent: true,
	}
	FilesToBatchFlag = Flag{
		Name:       "files-to-batch",
		ConfigName: "worker.files-to-batch",
		Value:      1,
		Usage:      "number of files to batch to worker",
		Persistent: true,
	}
	MemoryMaximumFlag = Flag{
		Name:       "memory-max",
		ConfigName: "worker.memory-max",
		Value:      800 * 1000 * 1000, // 800 MB
		Usage:      "if memory needed to scan a file surpasses this limit, skip the file",
		Persistent: true,
	}
)

type WorkerFlagGroup struct {
	Workers                   *Flag
	Timeout                   *Flag
	TimeoutFileMinimum        *Flag
	TimeoutFileMaximum        *Flag
	TimeoutFileSecondPerBytes *Flag
	TimeoutWorkerOnline       *Flag
	FileSizeMaximum           *Flag
	FilesToBatch              *Flag
	MemoryMaximum             *Flag
}

// GlobalOptions defines flags and other configuration parameters for all the subcommands
type WorkerOptions struct {
	Workers                   int
	Timeout                   time.Duration
	TimeoutFileMinimum        time.Duration
	TimeoutFileMaximum        time.Duration
	TimeoutFileSecondPerBytes int
	TimeoutWorkerOnline       time.Duration
	FileSizeMaximum           int
	FilesToBatch              int
	MemoryMaximum             int
}

func NewWorkerFlagGroup() *WorkerFlagGroup {
	return &WorkerFlagGroup{
		Workers:                   &WorkersFlag,
		Timeout:                   &TimeoutFlag,
		TimeoutFileMinimum:        &TimeoutFileMinimumFlag,
		TimeoutFileMaximum:        &TimeoutFileMaximumFlag,
		TimeoutFileSecondPerBytes: &TimeoutFileSecondPerBytesFlag,
		TimeoutWorkerOnline:       &TimeoutWorkerOnlineFlag,
		FileSizeMaximum:           &FileSizeMaximumFlag,
		FilesToBatch:              &FilesToBatchFlag,
		MemoryMaximum:             &MemoryMaximumFlag,
	}
}

func (f *WorkerFlagGroup) Name() string {
	return "Worker"
}

func (f *WorkerFlagGroup) Flags() []*Flag {
	return []*Flag{
		f.Workers,
		f.Timeout,
		f.TimeoutFileMinimum,
		f.TimeoutFileMaximum,
		f.TimeoutFileSecondPerBytes,
		f.TimeoutWorkerOnline,
		f.FileSizeMaximum,
		f.FilesToBatch,
		f.MemoryMaximum,
	}
}

func (f *WorkerFlagGroup) ToOptions() WorkerOptions {
	return WorkerOptions{
		Workers:                   getInt(f.Workers),
		Timeout:                   getDuration(f.Timeout),
		TimeoutFileMinimum:        getDuration(f.TimeoutFileMinimum),
		TimeoutFileMaximum:        getDuration(f.TimeoutFileMaximum),
		TimeoutFileSecondPerBytes: getInt(f.TimeoutFileSecondPerBytes),
		TimeoutWorkerOnline:       getDuration(f.TimeoutWorkerOnline),
		FilesToBatch:              getInt(f.FilesToBatch),
		FileSizeMaximum:           getInt(f.FileSizeMaximum),
		MemoryMaximum:             getInt(f.MemoryMaximum),
	}
}