package scanner

import (
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

// Ensures that a positive number can be scanned.
func TestScanPositiveNumber(t *testing.T) {
	tok, b, err := NewScanner(strings.NewReader("100")).Scan()
	assert.NoError(t, err)
	assert.Equal(t, tok, TNUMBER)
	assert.Equal(t, string(b), "100")
}

// Ensures that a negative number can be scanned.
func TestScanNegativeNumber(t *testing.T) {
	tok, b, err := NewScanner(strings.NewReader("-1")).Scan()
	assert.NoError(t, err)
	assert.Equal(t, tok, TNUMBER)
	assert.Equal(t, string(b), "-1")
}

// Ensures that a fractional number can be scanned.
func TestScanFloat(t *testing.T) {
	tok, b, err := NewScanner(strings.NewReader("120.12931")).Scan()
	assert.NoError(t, err)
	assert.Equal(t, tok, TNUMBER)
	assert.Equal(t, string(b), "120.12931")
}

// Ensures that a quoted string can be scanned.
func TestScanString(t *testing.T) {
	tok, b, err := NewScanner(strings.NewReader(`"hello world"`)).Scan()
	assert.NoError(t, err)
	assert.Equal(t, tok, TSTRING)
	assert.Equal(t, string(b), "hello world")
}

// Ensures that a quoted string with escaped characters can be scanned.
func TestScanEscapedString(t *testing.T) {
	tok, b, err := NewScanner(strings.NewReader(`"\"\\\/\b\f\n\r\t"`)).Scan()
	assert.NoError(t, err)
	assert.Equal(t, tok, TSTRING)
	assert.Equal(t, string(b), "\"\\/\b\f\n\r\t")
}

// Ensures that a true value can be scanned.
func TestScanTrue(t *testing.T) {
	tok, _, err := NewScanner(strings.NewReader(`true`)).Scan()
	assert.NoError(t, err)
	assert.Equal(t, tok, TTRUE)
}

// Ensures that a false value can be scanned.
func TestScanFalse(t *testing.T) {
	tok, _, err := NewScanner(strings.NewReader(`false`)).Scan()
	assert.NoError(t, err)
	assert.Equal(t, tok, TFALSE)
}

// Ensures that a null value can be scanned.
func TestScanNull(t *testing.T) {
	tok, _, err := NewScanner(strings.NewReader(`null`)).Scan()
	assert.NoError(t, err)
	assert.Equal(t, tok, TNULL)
}

// Ensures that an EOF gets returned.
func TestScanEOF(t *testing.T) {
	_, _, err := NewScanner(strings.NewReader(``)).Scan()
	assert.Equal(t, err, io.EOF)
}


// Ensures that a string can be read into a field.
func TestReadString(t *testing.T) {
	var v string
	err := NewScanner(strings.NewReader(`"foo"`)).ReadString(&v)
	assert.NoError(t, err)
	assert.Equal(t, v, "foo")
}

// Ensures that a non-string value is read into a string field as blank.
func TestReadNonStringAsString(t *testing.T) {
	var v string
	err := NewScanner(strings.NewReader(`12`)).ReadString(&v)
	assert.NoError(t, err)
	assert.Equal(t, v, "")
}

// Ensures that a non-value returns a read error.
func TestReadNonValueAsString(t *testing.T) {
	var v string
	err := NewScanner(strings.NewReader(`{`)).ReadString(&v)
	assert.Error(t, err)
}

// Ensures that an int can be read into a field.
func TestReadInt(t *testing.T) {
	var v int
	err := NewScanner(strings.NewReader(`100`)).ReadInt(&v)
	assert.NoError(t, err)
	assert.Equal(t, v, 100)
}

// Ensures that a non-number value is read into an int field as zero.
func TestReadNonNumberAsInt(t *testing.T) {
	var v int
	err := NewScanner(strings.NewReader(`"foo"`)).ReadInt(&v)
	assert.NoError(t, err)
	assert.Equal(t, v, 0)
}

// Ensures that an int64 can be read into a field.
func TestReadInt64(t *testing.T) {
	var v int64
	err := NewScanner(strings.NewReader(`-100`)).ReadInt64(&v)
	assert.NoError(t, err)
	assert.Equal(t, v, -100)
}

// Ensures that a uint can be read into a field.
func TestReadUint(t *testing.T) {
	var v uint
	err := NewScanner(strings.NewReader(`100`)).ReadUint(&v)
	assert.NoError(t, err)
	assert.Equal(t, v, uint(100))
}

// Ensures that an uint64 can be read into a field.
func TestReadUint64(t *testing.T) {
	var v uint64
	err := NewScanner(strings.NewReader(`1024`)).ReadUint64(&v)
	assert.NoError(t, err)
	assert.Equal(t, v, uint(1024))
}

// Ensures that a float32 can be read into a field.
func TestReadFloat32(t *testing.T) {
	var v float32
	err := NewScanner(strings.NewReader(`1293.123`)).ReadFloat32(&v)
	assert.NoError(t, err)
	assert.Equal(t, v, float32(1293.123))
}

// Ensures that a float64 can be read into a field.
func TestReadFloat64(t *testing.T) {
	var v float64
	err := NewScanner(strings.NewReader(`9871293.414123`)).ReadFloat64(&v)
	assert.NoError(t, err)
	assert.Equal(t, v, 9871293.414123)
}

// Ensures that a boolean can be read into a field.
func TestReadBoolTrue(t *testing.T) {
	var v bool
	err := NewScanner(strings.NewReader(`true`)).ReadBool(&v)
	assert.NoError(t, err)
	assert.Equal(t, v, true)
}
