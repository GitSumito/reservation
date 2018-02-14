// modify!!

package handler

// APIResponse ...
type APIResponse struct {
    Status   int
    Response interface{}
}

const (
    // AppStatusOK ...
    AppStatusOK = iota
    // AppStatusError ...
    AppStatusError
    // AppStatusNotFuond ...
    AppStatusNotFuond
)