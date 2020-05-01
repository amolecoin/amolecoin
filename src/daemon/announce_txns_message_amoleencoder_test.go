// Code generated by github.com/amolecoin/amoleencoder. DO NOT EDIT.

package daemon

import (
	"bytes"
	"fmt"
	mathrand "math/rand"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/amolecoin/encodertest"
	"github.com/amolecoin/amolecoin/src/cipher/encoder"
)

func newEmptyAnnounceTxnsMessageForEncodeTest() *AnnounceTxnsMessage {
	var obj AnnounceTxnsMessage
	return &obj
}

func newRandomAnnounceTxnsMessageForEncodeTest(t *testing.T, rand *mathrand.Rand) *AnnounceTxnsMessage {
	var obj AnnounceTxnsMessage
	err := encodertest.PopulateRandom(&obj, rand, encodertest.PopulateRandomOptions{
		MaxRandLen: 4,
		MinRandLen: 1,
	})
	if err != nil {
		t.Fatalf("encodertest.PopulateRandom failed: %v", err)
	}
	return &obj
}

func newRandomZeroLenAnnounceTxnsMessageForEncodeTest(t *testing.T, rand *mathrand.Rand) *AnnounceTxnsMessage {
	var obj AnnounceTxnsMessage
	err := encodertest.PopulateRandom(&obj, rand, encodertest.PopulateRandomOptions{
		MaxRandLen:    0,
		MinRandLen:    0,
		EmptySliceNil: false,
		EmptyMapNil:   false,
	})
	if err != nil {
		t.Fatalf("encodertest.PopulateRandom failed: %v", err)
	}
	return &obj
}

func newRandomZeroLenNilAnnounceTxnsMessageForEncodeTest(t *testing.T, rand *mathrand.Rand) *AnnounceTxnsMessage {
	var obj AnnounceTxnsMessage
	err := encodertest.PopulateRandom(&obj, rand, encodertest.PopulateRandomOptions{
		MaxRandLen:    0,
		MinRandLen:    0,
		EmptySliceNil: true,
		EmptyMapNil:   true,
	})
	if err != nil {
		t.Fatalf("encodertest.PopulateRandom failed: %v", err)
	}
	return &obj
}

func testAmoleencoderAnnounceTxnsMessage(t *testing.T, obj *AnnounceTxnsMessage) {
	isEncodableField := func(f reflect.StructField) bool {
		// Skip unexported fields
		if f.PkgPath != "" {
			return false
		}

		// Skip fields disabled with and enc:"- struct tag
		tag := f.Tag.Get("enc")
		return !strings.HasPrefix(tag, "-,") && tag != "-"
	}

	hasOmitEmptyField := func(obj interface{}) bool {
		v := reflect.ValueOf(obj)
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		}

		switch v.Kind() {
		case reflect.Struct:
			t := v.Type()
			n := v.NumField()
			f := t.Field(n - 1)
			tag := f.Tag.Get("enc")
			return isEncodableField(f) && strings.Contains(tag, ",omitempty")
		default:
			return false
		}
	}

	// returns the number of bytes encoded by an omitempty field on a given object
	omitEmptyLen := func(obj interface{}) uint64 {
		if !hasOmitEmptyField(obj) {
			return 0
		}

		v := reflect.ValueOf(obj)
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		}

		switch v.Kind() {
		case reflect.Struct:
			n := v.NumField()
			f := v.Field(n - 1)
			if f.Len() == 0 {
				return 0
			}
			return uint64(4 + f.Len())

		default:
			return 0
		}
	}

	// encodeSize

	n1 := encoder.Size(obj)
	n2 := encodeSizeAnnounceTxnsMessage(obj)

	if uint64(n1) != n2 {
		t.Fatalf("encoder.Size() != encodeSizeAnnounceTxnsMessage() (%d != %d)", n1, n2)
	}

	// Encode

	// encoder.Serialize
	data1 := encoder.Serialize(obj)

	// Encode
	data2, err := encodeAnnounceTxnsMessage(obj)
	if err != nil {
		t.Fatalf("encodeAnnounceTxnsMessage failed: %v", err)
	}
	if uint64(len(data2)) != n2 {
		t.Fatal("encodeAnnounceTxnsMessage produced bytes of unexpected length")
	}
	if len(data1) != len(data2) {
		t.Fatalf("len(encoder.Serialize()) != len(encodeAnnounceTxnsMessage()) (%d != %d)", len(data1), len(data2))
	}

	// EncodeToBuffer
	data3 := make([]byte, n2+5)
	if err := encodeAnnounceTxnsMessageToBuffer(data3, obj); err != nil {
		t.Fatalf("encodeAnnounceTxnsMessageToBuffer failed: %v", err)
	}

	if !bytes.Equal(data1, data2) {
		t.Fatal("encoder.Serialize() != encode[1]s()")
	}

	// Decode

	// encoder.DeserializeRaw
	var obj2 AnnounceTxnsMessage
	if n, err := encoder.DeserializeRaw(data1, &obj2); err != nil {
		t.Fatalf("encoder.DeserializeRaw failed: %v", err)
	} else if n != uint64(len(data1)) {
		t.Fatalf("encoder.DeserializeRaw failed: %v", encoder.ErrRemainingBytes)
	}
	if !cmp.Equal(*obj, obj2, cmpopts.EquateEmpty(), encodertest.IgnoreAllUnexported()) {
		t.Fatal("encoder.DeserializeRaw result wrong")
	}

	// Decode
	var obj3 AnnounceTxnsMessage
	if n, err := decodeAnnounceTxnsMessage(data2, &obj3); err != nil {
		t.Fatalf("decodeAnnounceTxnsMessage failed: %v", err)
	} else if n != uint64(len(data2)) {
		t.Fatalf("decodeAnnounceTxnsMessage bytes read length should be %d, is %d", len(data2), n)
	}
	if !cmp.Equal(obj2, obj3, cmpopts.EquateEmpty(), encodertest.IgnoreAllUnexported()) {
		t.Fatal("encoder.DeserializeRaw() != decodeAnnounceTxnsMessage()")
	}

	// Decode, excess buffer
	var obj4 AnnounceTxnsMessage
	n, err := decodeAnnounceTxnsMessage(data3, &obj4)
	if err != nil {
		t.Fatalf("decodeAnnounceTxnsMessage failed: %v", err)
	}

	if hasOmitEmptyField(&obj4) && omitEmptyLen(&obj4) == 0 {
		// 4 bytes read for the omitEmpty length, which should be zero (see the 5 bytes added above)
		if n != n2+4 {
			t.Fatalf("decodeAnnounceTxnsMessage bytes read length should be %d, is %d", n2+4, n)
		}
	} else {
		if n != n2 {
			t.Fatalf("decodeAnnounceTxnsMessage bytes read length should be %d, is %d", n2, n)
		}
	}
	if !cmp.Equal(obj2, obj4, cmpopts.EquateEmpty(), encodertest.IgnoreAllUnexported()) {
		t.Fatal("encoder.DeserializeRaw() != decodeAnnounceTxnsMessage()")
	}

	// DecodeExact
	var obj5 AnnounceTxnsMessage
	if err := decodeAnnounceTxnsMessageExact(data2, &obj5); err != nil {
		t.Fatalf("decodeAnnounceTxnsMessage failed: %v", err)
	}
	if !cmp.Equal(obj2, obj5, cmpopts.EquateEmpty(), encodertest.IgnoreAllUnexported()) {
		t.Fatal("encoder.DeserializeRaw() != decodeAnnounceTxnsMessage()")
	}

	// Check that the bytes read value is correct when providing an extended buffer
	if !hasOmitEmptyField(&obj3) || omitEmptyLen(&obj3) > 0 {
		padding := []byte{0xFF, 0xFE, 0xFD, 0xFC}
		data4 := append(data2[:], padding...)
		if n, err := decodeAnnounceTxnsMessage(data4, &obj3); err != nil {
			t.Fatalf("decodeAnnounceTxnsMessage failed: %v", err)
		} else if n != uint64(len(data2)) {
			t.Fatalf("decodeAnnounceTxnsMessage bytes read length should be %d, is %d", len(data2), n)
		}
	}
}

func TestAmoleencoderAnnounceTxnsMessage(t *testing.T) {
	rand := mathrand.New(mathrand.NewSource(time.Now().Unix()))

	type testCase struct {
		name string
		obj  *AnnounceTxnsMessage
	}

	cases := []testCase{
		{
			name: "empty object",
			obj:  newEmptyAnnounceTxnsMessageForEncodeTest(),
		},
	}

	nRandom := 10

	for i := 0; i < nRandom; i++ {
		cases = append(cases, testCase{
			name: fmt.Sprintf("randomly populated object %d", i),
			obj:  newRandomAnnounceTxnsMessageForEncodeTest(t, rand),
		})
		cases = append(cases, testCase{
			name: fmt.Sprintf("randomly populated object %d with zero length variable length contents", i),
			obj:  newRandomZeroLenAnnounceTxnsMessageForEncodeTest(t, rand),
		})
		cases = append(cases, testCase{
			name: fmt.Sprintf("randomly populated object %d with zero length variable length contents set to nil", i),
			obj:  newRandomZeroLenNilAnnounceTxnsMessageForEncodeTest(t, rand),
		})
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			testAmoleencoderAnnounceTxnsMessage(t, tc.obj)
		})
	}
}

func decodeAnnounceTxnsMessageExpectError(t *testing.T, buf []byte, expectedErr error) {
	var obj AnnounceTxnsMessage
	if _, err := decodeAnnounceTxnsMessage(buf, &obj); err == nil {
		t.Fatal("decodeAnnounceTxnsMessage: expected error, got nil")
	} else if err != expectedErr {
		t.Fatalf("decodeAnnounceTxnsMessage: expected error %q, got %q", expectedErr, err)
	}
}

func decodeAnnounceTxnsMessageExactExpectError(t *testing.T, buf []byte, expectedErr error) {
	var obj AnnounceTxnsMessage
	if err := decodeAnnounceTxnsMessageExact(buf, &obj); err == nil {
		t.Fatal("decodeAnnounceTxnsMessageExact: expected error, got nil")
	} else if err != expectedErr {
		t.Fatalf("decodeAnnounceTxnsMessageExact: expected error %q, got %q", expectedErr, err)
	}
}

func testAmoleencoderAnnounceTxnsMessageDecodeErrors(t *testing.T, k int, tag string, obj *AnnounceTxnsMessage) {
	isEncodableField := func(f reflect.StructField) bool {
		// Skip unexported fields
		if f.PkgPath != "" {
			return false
		}

		// Skip fields disabled with and enc:"- struct tag
		tag := f.Tag.Get("enc")
		return !strings.HasPrefix(tag, "-,") && tag != "-"
	}

	numEncodableFields := func(obj interface{}) int {
		v := reflect.ValueOf(obj)
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		}

		switch v.Kind() {
		case reflect.Struct:
			t := v.Type()

			n := 0
			for i := 0; i < v.NumField(); i++ {
				f := t.Field(i)
				if !isEncodableField(f) {
					continue
				}
				n++
			}
			return n
		default:
			return 0
		}
	}

	hasOmitEmptyField := func(obj interface{}) bool {
		v := reflect.ValueOf(obj)
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		}

		switch v.Kind() {
		case reflect.Struct:
			t := v.Type()
			n := v.NumField()
			f := t.Field(n - 1)
			tag := f.Tag.Get("enc")
			return isEncodableField(f) && strings.Contains(tag, ",omitempty")
		default:
			return false
		}
	}

	// returns the number of bytes encoded by an omitempty field on a given object
	omitEmptyLen := func(obj interface{}) uint64 {
		if !hasOmitEmptyField(obj) {
			return 0
		}

		v := reflect.ValueOf(obj)
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		}

		switch v.Kind() {
		case reflect.Struct:
			n := v.NumField()
			f := v.Field(n - 1)
			if f.Len() == 0 {
				return 0
			}
			return uint64(4 + f.Len())

		default:
			return 0
		}
	}

	n := encodeSizeAnnounceTxnsMessage(obj)
	buf, err := encodeAnnounceTxnsMessage(obj)
	if err != nil {
		t.Fatalf("encodeAnnounceTxnsMessage failed: %v", err)
	}

	// A nil buffer cannot decode, unless the object is a struct with a single omitempty field
	if hasOmitEmptyField(obj) && numEncodableFields(obj) > 1 {
		t.Run(fmt.Sprintf("%d %s buffer underflow nil", k, tag), func(t *testing.T) {
			decodeAnnounceTxnsMessageExpectError(t, nil, encoder.ErrBufferUnderflow)
		})

		t.Run(fmt.Sprintf("%d %s exact buffer underflow nil", k, tag), func(t *testing.T) {
			decodeAnnounceTxnsMessageExactExpectError(t, nil, encoder.ErrBufferUnderflow)
		})
	}

	// Test all possible truncations of the encoded byte array, but skip
	// a truncation that would be valid where omitempty is removed
	skipN := n - omitEmptyLen(obj)
	for i := uint64(0); i < n; i++ {
		if i == skipN {
			continue
		}

		t.Run(fmt.Sprintf("%d %s buffer underflow bytes=%d", k, tag, i), func(t *testing.T) {
			decodeAnnounceTxnsMessageExpectError(t, buf[:i], encoder.ErrBufferUnderflow)
		})

		t.Run(fmt.Sprintf("%d %s exact buffer underflow bytes=%d", k, tag, i), func(t *testing.T) {
			decodeAnnounceTxnsMessageExactExpectError(t, buf[:i], encoder.ErrBufferUnderflow)
		})
	}

	// Append 5 bytes for omit empty with a 0 length prefix, to cause an ErrRemainingBytes.
	// If only 1 byte is appended, the decoder will try to read the 4-byte length prefix,
	// and return an ErrBufferUnderflow instead
	if hasOmitEmptyField(obj) {
		buf = append(buf, []byte{0, 0, 0, 0, 0}...)
	} else {
		buf = append(buf, 0)
	}

	t.Run(fmt.Sprintf("%d %s exact buffer remaining bytes", k, tag), func(t *testing.T) {
		decodeAnnounceTxnsMessageExactExpectError(t, buf, encoder.ErrRemainingBytes)
	})
}

func TestAmoleencoderAnnounceTxnsMessageDecodeErrors(t *testing.T) {
	rand := mathrand.New(mathrand.NewSource(time.Now().Unix()))
	n := 10

	for i := 0; i < n; i++ {
		emptyObj := newEmptyAnnounceTxnsMessageForEncodeTest()
		fullObj := newRandomAnnounceTxnsMessageForEncodeTest(t, rand)
		testAmoleencoderAnnounceTxnsMessageDecodeErrors(t, i, "empty", emptyObj)
		testAmoleencoderAnnounceTxnsMessageDecodeErrors(t, i, "full", fullObj)
	}
}
