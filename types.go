package types


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
}