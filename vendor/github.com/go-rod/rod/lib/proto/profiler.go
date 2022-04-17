// This file is generated by "./lib/proto/generate"

package proto

/*

Profiler

*/

// ProfilerProfileNode Profile node. Holds callsite information, execution statistics and child nodes.
type ProfilerProfileNode struct {

	// ID Unique id of the node.
	ID int `json:"id"`

	// CallFrame Function location.
	CallFrame *RuntimeCallFrame `json:"callFrame"`

	// HitCount (optional) Number of samples where this node was on top of the call stack.
	HitCount *int `json:"hitCount,omitempty"`

	// Children (optional) Child node ids.
	Children []int `json:"children,omitempty"`

	// DeoptReason (optional) The reason of being not optimized. The function may be deoptimized or marked as don't
	// optimize.
	DeoptReason string `json:"deoptReason,omitempty"`

	// PositionTicks (optional) An array of source position ticks.
	PositionTicks []*ProfilerPositionTickInfo `json:"positionTicks,omitempty"`
}

// ProfilerProfile Profile.
type ProfilerProfile struct {

	// Nodes The list of profile nodes. First item is the root node.
	Nodes []*ProfilerProfileNode `json:"nodes"`

	// StartTime Profiling start timestamp in microseconds.
	StartTime float64 `json:"startTime"`

	// EndTime Profiling end timestamp in microseconds.
	EndTime float64 `json:"endTime"`

	// Samples (optional) Ids of samples top nodes.
	Samples []int `json:"samples,omitempty"`

	// TimeDeltas (optional) Time intervals between adjacent samples in microseconds. The first delta is relative to the
	// profile startTime.
	TimeDeltas []int `json:"timeDeltas,omitempty"`
}

// ProfilerPositionTickInfo Specifies a number of samples attributed to a certain source position.
type ProfilerPositionTickInfo struct {

	// Line Source line number (1-based).
	Line int `json:"line"`

	// Ticks Number of samples attributed to the source line.
	Ticks int `json:"ticks"`
}

// ProfilerCoverageRange Coverage data for a source range.
type ProfilerCoverageRange struct {

	// StartOffset JavaScript script source offset for the range start.
	StartOffset int `json:"startOffset"`

	// EndOffset JavaScript script source offset for the range end.
	EndOffset int `json:"endOffset"`

	// Count Collected execution count of the source range.
	Count int `json:"count"`
}

// ProfilerFunctionCoverage Coverage data for a JavaScript function.
type ProfilerFunctionCoverage struct {

	// FunctionName JavaScript function name.
	FunctionName string `json:"functionName"`

	// Ranges Source ranges inside the function with coverage data.
	Ranges []*ProfilerCoverageRange `json:"ranges"`

	// IsBlockCoverage Whether coverage data for this function has block granularity.
	IsBlockCoverage bool `json:"isBlockCoverage"`
}

// ProfilerScriptCoverage Coverage data for a JavaScript script.
type ProfilerScriptCoverage struct {

	// ScriptID JavaScript script id.
	ScriptID RuntimeScriptID `json:"scriptId"`

	// URL JavaScript script name or url.
	URL string `json:"url"`

	// Functions Functions contained in the script that has coverage data.
	Functions []*ProfilerFunctionCoverage `json:"functions"`
}

// ProfilerTypeObject (experimental) Describes a type collected during runtime.
type ProfilerTypeObject struct {

	// Name Name of a type collected with type profiling.
	Name string `json:"name"`
}

// ProfilerTypeProfileEntry (experimental) Source offset and types for a parameter or return value.
type ProfilerTypeProfileEntry struct {

	// Offset Source offset of the parameter or end of function for return values.
	Offset int `json:"offset"`

	// Types The types for this parameter or return value.
	Types []*ProfilerTypeObject `json:"types"`
}

// ProfilerScriptTypeProfile (experimental) Type profile data collected during runtime for a JavaScript script.
type ProfilerScriptTypeProfile struct {

	// ScriptID JavaScript script id.
	ScriptID RuntimeScriptID `json:"scriptId"`

	// URL JavaScript script name or url.
	URL string `json:"url"`

	// Entries Type profile entries for parameters and return values of the functions in the script.
	Entries []*ProfilerTypeProfileEntry `json:"entries"`
}

// ProfilerDisable ...
type ProfilerDisable struct {
}

// ProtoReq name
func (m ProfilerDisable) ProtoReq() string { return "Profiler.disable" }

// Call sends the request
func (m ProfilerDisable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// ProfilerEnable ...
type ProfilerEnable struct {
}

// ProtoReq name
func (m ProfilerEnable) ProtoReq() string { return "Profiler.enable" }

// Call sends the request
func (m ProfilerEnable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// ProfilerGetBestEffortCoverage Collect coverage data for the current isolate. The coverage data may be incomplete due to
// garbage collection.
type ProfilerGetBestEffortCoverage struct {
}

// ProtoReq name
func (m ProfilerGetBestEffortCoverage) ProtoReq() string { return "Profiler.getBestEffortCoverage" }

// Call the request
func (m ProfilerGetBestEffortCoverage) Call(c Client) (*ProfilerGetBestEffortCoverageResult, error) {
	var res ProfilerGetBestEffortCoverageResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// ProfilerGetBestEffortCoverageResult Collect coverage data for the current isolate. The coverage data may be incomplete due to
// garbage collection.
type ProfilerGetBestEffortCoverageResult struct {

	// Result Coverage data for the current isolate.
	Result []*ProfilerScriptCoverage `json:"result"`
}

// ProfilerSetSamplingInterval Changes CPU profiler sampling interval. Must be called before CPU profiles recording started.
type ProfilerSetSamplingInterval struct {

	// Interval New sampling interval in microseconds.
	Interval int `json:"interval"`
}

// ProtoReq name
func (m ProfilerSetSamplingInterval) ProtoReq() string { return "Profiler.setSamplingInterval" }

// Call sends the request
func (m ProfilerSetSamplingInterval) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// ProfilerStart ...
type ProfilerStart struct {
}

// ProtoReq name
func (m ProfilerStart) ProtoReq() string { return "Profiler.start" }

// Call sends the request
func (m ProfilerStart) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// ProfilerStartPreciseCoverage Enable precise code coverage. Coverage data for JavaScript executed before enabling precise code
// coverage may be incomplete. Enabling prevents running optimized code and resets execution
// counters.
type ProfilerStartPreciseCoverage struct {

	// CallCount (optional) Collect accurate call counts beyond simple 'covered' or 'not covered'.
	CallCount bool `json:"callCount,omitempty"`

	// Detailed (optional) Collect block-based coverage.
	Detailed bool `json:"detailed,omitempty"`

	// AllowTriggeredUpdates (optional) Allow the backend to send updates on its own initiative
	AllowTriggeredUpdates bool `json:"allowTriggeredUpdates,omitempty"`
}

// ProtoReq name
func (m ProfilerStartPreciseCoverage) ProtoReq() string { return "Profiler.startPreciseCoverage" }

// Call the request
func (m ProfilerStartPreciseCoverage) Call(c Client) (*ProfilerStartPreciseCoverageResult, error) {
	var res ProfilerStartPreciseCoverageResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// ProfilerStartPreciseCoverageResult Enable precise code coverage. Coverage data for JavaScript executed before enabling precise code
// coverage may be incomplete. Enabling prevents running optimized code and resets execution
// counters.
type ProfilerStartPreciseCoverageResult struct {

	// Timestamp Monotonically increasing time (in seconds) when the coverage update was taken in the backend.
	Timestamp float64 `json:"timestamp"`
}

// ProfilerStartTypeProfile (experimental) Enable type profile.
type ProfilerStartTypeProfile struct {
}

// ProtoReq name
func (m ProfilerStartTypeProfile) ProtoReq() string { return "Profiler.startTypeProfile" }

// Call sends the request
func (m ProfilerStartTypeProfile) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// ProfilerStop ...
type ProfilerStop struct {
}

// ProtoReq name
func (m ProfilerStop) ProtoReq() string { return "Profiler.stop" }

// Call the request
func (m ProfilerStop) Call(c Client) (*ProfilerStopResult, error) {
	var res ProfilerStopResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// ProfilerStopResult ...
type ProfilerStopResult struct {

	// Profile Recorded profile.
	Profile *ProfilerProfile `json:"profile"`
}

// ProfilerStopPreciseCoverage Disable precise code coverage. Disabling releases unnecessary execution count records and allows
// executing optimized code.
type ProfilerStopPreciseCoverage struct {
}

// ProtoReq name
func (m ProfilerStopPreciseCoverage) ProtoReq() string { return "Profiler.stopPreciseCoverage" }

// Call sends the request
func (m ProfilerStopPreciseCoverage) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// ProfilerStopTypeProfile (experimental) Disable type profile. Disabling releases type profile data collected so far.
type ProfilerStopTypeProfile struct {
}

// ProtoReq name
func (m ProfilerStopTypeProfile) ProtoReq() string { return "Profiler.stopTypeProfile" }

// Call sends the request
func (m ProfilerStopTypeProfile) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// ProfilerTakePreciseCoverage Collect coverage data for the current isolate, and resets execution counters. Precise code
// coverage needs to have started.
type ProfilerTakePreciseCoverage struct {
}

// ProtoReq name
func (m ProfilerTakePreciseCoverage) ProtoReq() string { return "Profiler.takePreciseCoverage" }

// Call the request
func (m ProfilerTakePreciseCoverage) Call(c Client) (*ProfilerTakePreciseCoverageResult, error) {
	var res ProfilerTakePreciseCoverageResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// ProfilerTakePreciseCoverageResult Collect coverage data for the current isolate, and resets execution counters. Precise code
// coverage needs to have started.
type ProfilerTakePreciseCoverageResult struct {

	// Result Coverage data for the current isolate.
	Result []*ProfilerScriptCoverage `json:"result"`

	// Timestamp Monotonically increasing time (in seconds) when the coverage update was taken in the backend.
	Timestamp float64 `json:"timestamp"`
}

// ProfilerTakeTypeProfile (experimental) Collect type profile.
type ProfilerTakeTypeProfile struct {
}

// ProtoReq name
func (m ProfilerTakeTypeProfile) ProtoReq() string { return "Profiler.takeTypeProfile" }

// Call the request
func (m ProfilerTakeTypeProfile) Call(c Client) (*ProfilerTakeTypeProfileResult, error) {
	var res ProfilerTakeTypeProfileResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// ProfilerTakeTypeProfileResult (experimental) Collect type profile.
type ProfilerTakeTypeProfileResult struct {

	// Result Type profile for all scripts since startTypeProfile() was turned on.
	Result []*ProfilerScriptTypeProfile `json:"result"`
}

// ProfilerConsoleProfileFinished ...
type ProfilerConsoleProfileFinished struct {

	// ID ...
	ID string `json:"id"`

	// Location Location of console.profileEnd().
	Location *DebuggerLocation `json:"location"`

	// Profile ...
	Profile *ProfilerProfile `json:"profile"`

	// Title (optional) Profile title passed as an argument to console.profile().
	Title string `json:"title,omitempty"`
}

// ProtoEvent name
func (evt ProfilerConsoleProfileFinished) ProtoEvent() string {
	return "Profiler.consoleProfileFinished"
}

// ProfilerConsoleProfileStarted Sent when new profile recording is started using console.profile() call.
type ProfilerConsoleProfileStarted struct {

	// ID ...
	ID string `json:"id"`

	// Location Location of console.profile().
	Location *DebuggerLocation `json:"location"`

	// Title (optional) Profile title passed as an argument to console.profile().
	Title string `json:"title,omitempty"`
}

// ProtoEvent name
func (evt ProfilerConsoleProfileStarted) ProtoEvent() string {
	return "Profiler.consoleProfileStarted"
}

// ProfilerPreciseCoverageDeltaUpdate (experimental) Reports coverage delta since the last poll (either from an event like this, or from
// `takePreciseCoverage` for the current isolate. May only be sent if precise code
// coverage has been started. This event can be trigged by the embedder to, for example,
// trigger collection of coverage data immediately at a certain point in time.
type ProfilerPreciseCoverageDeltaUpdate struct {

	// Timestamp Monotonically increasing time (in seconds) when the coverage update was taken in the backend.
	Timestamp float64 `json:"timestamp"`

	// Occasion Identifier for distinguishing coverage events.
	Occasion string `json:"occasion"`

	// Result Coverage data for the current isolate.
	Result []*ProfilerScriptCoverage `json:"result"`
}

// ProtoEvent name
func (evt ProfilerPreciseCoverageDeltaUpdate) ProtoEvent() string {
	return "Profiler.preciseCoverageDeltaUpdate"
}
