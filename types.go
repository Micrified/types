package types

import (
	"fmt"
	"io/ioutil"
	"os"
	"errors"
	"encoding/json"
)


/*
 *******************************************************************************
 *                    Type Definitions for Graph Generation                    *
 *******************************************************************************
*/


// Rules for generation of random ROS programs
type Rules struct {
	Name               string
	Directory          string
	Chain_count        int
	Chain_avg_len      int
	Chain_merge_p      float64
	Chain_sync_p       float64
	Chain_variance     float64
	Util_total         float64
	Min_period_us      int
	Max_period_us      int
	Period_step_us     float64
	Hyperperiod_count  int
	Max_duration_us    int
	PPE                bool
	Executor_count     int
	Random_seed        int
	Logging_mode       int       // Log (0: none, 1: callbacks, 2: chains)
}


/*
 *******************************************************************************
 *                        Type Definitions for Analysis                        *
 *******************************************************************************
*/


// Describes a trace from a chain execution
type Trace struct {
	ID                  int      // Chain Identifier
	Priority            int      // Chain priority value
	Length              int      // Length of the chain
	Period              int64    // Chain-specific period
	Utilisation         float64  // Chain-specific utilisation
	BCRT_us             int64    // Best-case response time
	WCRT_us             int64    // Worst-case response time
	ACRT_us             int64    // Average-case response time
	Chain_count         int      // Number of concurrent chains
	Avg_chain_length    int      // Average chain length of system
	Seed                int      // Seed used in generating system
	Merge_p             float64  // Merge probability used in system
	Sync_p              float64  // Sync probability used in system
	Variance            float64  // Variance (chain length) used
	PPE                 int      // Nonzero if PPE was used
	Executors           int      // Number of executors in the system
}

func (t *Trace) Serialise (to *os.File) error {
	sfmt := "%d %d %d %d %f %d %d %d %d %d %d %f %f %f %d %d\n"
	output_string := fmt.Sprintf(sfmt,
		t.ID, t.Priority, t.Length, t.Period, t.Utilisation,
	    t.BCRT_us, t.WCRT_us, t.ACRT_us, t.Chain_count, 
		t.Avg_chain_length, t.Seed, t.Merge_p, t.Sync_p,
		t.Variance, t.PPE, t.Executors)
	output_buffer := []byte(output_string)
	_, err := to.Write(output_buffer)
	if nil != err {
		return err
	}
	return nil
}

func (t *Trace) ParseFrom (line []byte) error {
	input_string := string(line)
	sfmt := "%d %d %d %d %f %d %d %d %d %d %d %f %f %f %d %d\n"
	n, err := fmt.Sscanf(input_string, sfmt,
		&(t.ID), &(t.Priority), &(t.Length), &(t.Period), &(t.Utilisation),
		&(t.BCRT_us), &(t.WCRT_us), &(t.ACRT_us), &(t.Chain_count),
	    &(t.Avg_chain_length), &(t.Seed), &(t.Merge_p), &(t.Sync_p),
	    &(t.Variance), &(t.PPE), &(t.Executors))
	if nil != err || n != 16 {
		return errors.New("Unable to parse \"" + input_string + "\": " +
			err.Error())
	}
	return nil
}


/*
 *******************************************************************************
 *                       Type Definitions for Benchmarks                       *
 *******************************************************************************
*/


// Describes a benchmark
type Benchmark struct {
	Name              string
	Execution_time_us int64
}

// Describes multiple benchmarks
type Benchmarks []Benchmark

// Describes work assigned to a callback
type Work struct {
	Benchmark_p       *Benchmark
	Iterations         int
}

func (b *Benchmarks) ReadFrom (filename string) error {
	buffer, err := ioutil.ReadFile(filename)
	if nil != err {
		return errors.New("Unable to read \"" + filename + "\": " + 
			err.Error())
	}
	return json.Unmarshal(buffer, b)
}