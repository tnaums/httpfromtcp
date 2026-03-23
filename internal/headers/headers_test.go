package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"	
)

func TestParse(t *testing.T) {
	// Test: Valid single header
	headers := NewHeaders()
	data := []byte("Host: localhost:42069\r\n\r\n")
	n, done, err := headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "localhost:42069", headers["host"])
	assert.Equal(t, 23, n)
	assert.False(t, done)

	// Test: Valid single header begining with whitespace
	headers = NewHeaders()
	data = []byte("  Host: localhost:42069\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "localhost:42069", headers["host"])
	assert.Equal(t, 25, n)
	assert.False(t, done)

	// Test: Invalid character in header
	headers = NewHeaders()
	data = []byte("H@st: localhost:42069\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)

	
	// Test: Invalid spacing header
	headers = NewHeaders()
	data = []byte("       Host : localhost:42069       \r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	// Test: incomplete header line
	headers = NewHeaders()
	data = []byte("Host: localhost:42")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	// Test: end of headers
	headers = NewHeaders()
	data = []byte("\r\n{tag: somevalue}")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	assert.Equal(t, 2, n)
	assert.True(t, done)

	// Test: two valid headers
	headers = NewHeaders()
	data = []byte("Host: localhost:42069\r\nApplication-type: json\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	assert.Equal(t, 23, n)
	assert.False(t, done)
	assert.Equal(t, "localhost:42069", headers["host"])	
	n, done, err = headers.Parse(data[23:])
	require.NoError(t, err)
	assert.Equal(t, 24, n)
	assert.False(t, done)
	n, done, err = headers.Parse(data[47:])
	require.NoError(t, err)
	assert.Equal(t, 2, n)
	assert.True(t, done)

	// Test: Multiple Set-Person lines
	headers = NewHeaders()
	data = []byte("Host: localhost:42069\r\nSet-Person: lane-loves-go\r\nSet-Person: prime-loves-zig\r\nSet-Person: tj-loves-ocaml\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	assert.Equal(t, 23, n)
	assert.False(t, done)
	assert.Equal(t, "localhost:42069", headers["host"])
	n, done, err = headers.Parse(data[23:])
	assert.Equal(t, 27, n)
	assert.False(t, done)
	assert.Equal(t, "lane-loves-go", headers["set-person"])
	n, done, err = headers.Parse(data[50:])
	assert.Equal(t, 29, n)
	assert.False(t, done)
	assert.Equal(t, "lane-loves-go, prime-loves-zig", headers["set-person"])	
}
