package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.ibm.com/dash/dash_utils/dashtest"
)

func TestM(t *testing.T) {
	templistenAndServe := listenAndServeFunc
	defer func() {
		listenAndServeFunc = templistenAndServe
	}()
	listenAndServeFunc = func(port string, hanle http.Handler) error {
		panic("customised error")
	}
	assert.PanicsWithValuef(t, "customised error", main, "they should be equal")
}
func TestMain(m *testing.M) {

	dashtest.ControlCoverage(m)
}
