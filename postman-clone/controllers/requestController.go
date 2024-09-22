package controllers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type RequestDetails struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Body    string            `json:"body"`
	Headers map[string]string `json:"headers"`
}

func HandleRequest(c *gin.Context) {
	var details RequestDetails
	if err := c.ShouldBindJSON(&details); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parsedUrl, err := url.ParseRequestURI(details.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	req, err := http.NewRequest(details.Method, parsedUrl.String(), bytes.NewBufferString(details.Body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Request creation failed"})
		return
	}

	for key, value := range details.Headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Request failed"})
		return
	}
	defer resp.Body.Close()

	responseHeaders := map[string][]string{}
	for key, values := range resp.Header {
		responseHeaders[key] = values
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  resp.Status,
		"headers": responseHeaders,
		"body":    string(responseBody),
	})
}
