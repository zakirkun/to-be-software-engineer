package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"imzakir.dev/e-commerce/pkg/logstash"
	"imzakir.dev/e-commerce/utils"
)

type bodyDumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *bodyDumpResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	n, err := w.Writer.Write(b)
	if err != nil {
		return n, fmt.Errorf("bodyDumpResponseWriter: failed to write response body: %w", err)
	}
	return n, nil
}

func Logging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Record the start time
		start := time.Now()

		// Request Body Handling
		var reqBody []byte
		if c.Request().Body != nil {
			reqBody, _ = io.ReadAll(c.Request().Body)
		}
		c.Request().Body = io.NopCloser(bytes.NewBuffer(reqBody)) // Reset

		// Parse Form Params
		params, _ := c.FormParams()
		jsonMap := make(map[string]interface{})

		ok := utils.JsonToSruct(reqBody, &jsonMap)
		if !ok {
			for key := range params {
				jsonMap[key] = c.FormValue(key)
			}
		}

		// Response Body Handling
		resBody := new(bytes.Buffer)
		mw := io.MultiWriter(c.Response().Writer, resBody)
		writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: c.Response().Writer}
		c.Response().Writer = writer

		var _header string

		for k, v := range c.Request().Header.Clone() {
			_header += fmt.Sprintf("Header Key: %v - Header Value: %v\n", k, strings.Join(v, ","))
		}

		// Proceed to the next handler
		exec := next(c)

		// Record the stop time
		stop := time.Now()

		// Calculate response time
		useTime := stop.Sub(start)

		// Create LogStash connection
		conn, err := logstash.LOGSTASH.Open()
		if err != nil {
			log.Printf("Error connecting to Logstash: %v", err.Error())
			// Skip if Logstash connection fails
			return exec
		}
		defer func(conn net.Conn) {
			err = conn.Close()
			if err != nil {
				log.Printf("Error closing Logstash connection: %v", err.Error())
			}
		}(conn)

		log.Infof("Request Header: %v", _header)
		log.Infof("Request: %v, Response: %v, Response Time: %v", utils.StructToJson(jsonMap), resBody.String(), useTime.String())

		setData := make(map[string]interface{})
		setData["FunctionName"] = c.Request().RequestURI
		setData["RequestBody"] = utils.StructToJson(jsonMap)
		setData["ResponseBody"] = resBody.String()
		setData["ResponseTime"] = useTime.String()

		logs := logstash.LogPayload{}
		// Send data to Logstash
		err = logs.SetIndex("request-response").SetData(&setData).WriteCaller("INFO", "Dump Request & Response")
		if err != nil {
			log.Printf("Error writing to Logstash: %v", err)
		}

		return exec
	}
}
