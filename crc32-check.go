package main

import (
	"encoding/json"
	"hash/crc32"
	"log"
	"math/bits"
	"strconv"
	"strings"

	typev3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	ep "github.com/wrossmorrow/envoy-extproc-sdk-go"
)

const (
	kCRC  = "crc32"
	kData = "data"
	kPoly = "poly"
)

type crc32CheckRequestProcessor struct {
	opts *ep.ProcessingOptions
	poly uint32
}

func (s *crc32CheckRequestProcessor) GetName() string {
	return "crc32-check"
}

func (s *crc32CheckRequestProcessor) GetOptions() *ep.ProcessingOptions {
	return s.opts
}

func (s *crc32CheckRequestProcessor) ProcessRequestHeaders(ctx *ep.RequestContext, headers ep.AllHeaders) error {
	return ctx.ContinueRequest()
}

func extract(m map[string]any, k string) string {
	vv, ok := m[k]
	if ok {
		return vv.(string)
	}
	return ""
}

func (s *crc32CheckRequestProcessor) crc(input []byte) uint32 {
	revPoly := bits.Reverse32(s.poly)
	t := crc32.MakeTable(revPoly)
	return crc32.Checksum(input, t)
}

func (s *crc32CheckRequestProcessor) ProcessRequestBody(ctx *ep.RequestContext, body []byte) error {
	cancel := func(code int32) error {
		return ctx.CancelRequest(code, map[string]ep.HeaderValue{}, typev3.StatusCode_name[code])
	}

	var unstructure map[string]any

	err := json.Unmarshal(body, &unstructure)
	if err != nil {
		log.Printf("parse the request is failed: %v", err.Error())
		return cancel(400)
	}

	actual, _ := strconv.ParseUint(extract(unstructure, kCRC), 16, 32)

	want := s.crc([]byte(extract(unstructure, kData)))
	if want != uint32(actual) {
		log.Printf("verify the checksum is failed, want=0x%0x actual=0x%0x", want, actual)
		return cancel(403)
	}

	return ctx.ContinueRequest()
}

func (s *crc32CheckRequestProcessor) ProcessRequestTrailers(ctx *ep.RequestContext, trailers ep.AllHeaders) error {
	return ctx.ContinueRequest()
}

func (s *crc32CheckRequestProcessor) ProcessResponseHeaders(ctx *ep.RequestContext, headers ep.AllHeaders) error {
	return ctx.ContinueRequest()
}

func (s *crc32CheckRequestProcessor) ProcessResponseBody(ctx *ep.RequestContext, body []byte) error {
	return ctx.ContinueRequest()
}

func (s *crc32CheckRequestProcessor) ProcessResponseTrailers(ctx *ep.RequestContext, trailers ep.AllHeaders) error {
	return ctx.ContinueRequest()
}

func (s *crc32CheckRequestProcessor) Init(opts *ep.ProcessingOptions, nonFlagArgs []string) error {
	s.opts = opts
	s.poly = crc32.IEEE

	var i int
	nArgs := len(nonFlagArgs)
	for ; i < nArgs-1; i++ {
		if nonFlagArgs[i] == kPoly {
			break
		}
	}

	if i == nArgs {
		log.Printf("the argument: 'poly' is missing, use the IEEE.\n")
	} else {
		polyStr := strings.ToLower(nonFlagArgs[i+1])

		if ok := strings.HasPrefix(polyStr, "0x"); ok && len(polyStr) > 2 {
			polyStr = polyStr[2:]
		}

		poly, _ := strconv.ParseInt(polyStr, 16, 64)
		if poly == 0 {
			log.Printf("parse the value for parameter: 'poly' is failed, use the IEEE.\n")
		} else {
			s.poly = uint32(poly)
			log.Printf("the poly is: 0x%0x.\n", s.poly)
		}
	}
	return nil
}

func (s *crc32CheckRequestProcessor) Finish() {}
