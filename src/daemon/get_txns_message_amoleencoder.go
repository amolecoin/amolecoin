// Code generated by github.com/amolecoin/amoleencoder. DO NOT EDIT.

package daemon

import (
	"errors"
	"math"

	"github.com/amolecoin/amolecoin/src/cipher"
	"github.com/amolecoin/amolecoin/src/cipher/encoder"
)

// encodeSizeGetTxnsMessage computes the size of an encoded object of type GetTxnsMessage
func encodeSizeGetTxnsMessage(obj *GetTxnsMessage) uint64 {
	i0 := uint64(0)

	// obj.Transactions
	i0 += 4
	{
		i1 := uint64(0)

		// x1
		i1 += 32

		i0 += uint64(len(obj.Transactions)) * i1
	}

	return i0
}

// encodeGetTxnsMessage encodes an object of type GetTxnsMessage to a buffer allocated to the exact size
// required to encode the object.
func encodeGetTxnsMessage(obj *GetTxnsMessage) ([]byte, error) {
	n := encodeSizeGetTxnsMessage(obj)
	buf := make([]byte, n)

	if err := encodeGetTxnsMessageToBuffer(buf, obj); err != nil {
		return nil, err
	}

	return buf, nil
}

// encodeGetTxnsMessageToBuffer encodes an object of type GetTxnsMessage to a []byte buffer.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func encodeGetTxnsMessageToBuffer(buf []byte, obj *GetTxnsMessage) error {
	if uint64(len(buf)) < encodeSizeGetTxnsMessage(obj) {
		return encoder.ErrBufferUnderflow
	}

	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.Transactions maxlen check
	if len(obj.Transactions) > 256 {
		return encoder.ErrMaxLenExceeded
	}

	// obj.Transactions length check
	if uint64(len(obj.Transactions)) > math.MaxUint32 {
		return errors.New("obj.Transactions length exceeds math.MaxUint32")
	}

	// obj.Transactions length
	e.Uint32(uint32(len(obj.Transactions)))

	// obj.Transactions
	for _, x := range obj.Transactions {

		// x
		e.CopyBytes(x[:])

	}

	return nil
}

// decodeGetTxnsMessage decodes an object of type GetTxnsMessage from a buffer.
// Returns the number of bytes used from the buffer to decode the object.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
func decodeGetTxnsMessage(buf []byte, obj *GetTxnsMessage) (uint64, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.Transactions

		ul, err := d.Uint32()
		if err != nil {
			return 0, err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return 0, encoder.ErrBufferUnderflow
		}

		if length > 256 {
			return 0, encoder.ErrMaxLenExceeded
		}

		if length != 0 {
			obj.Transactions = make([]cipher.SHA256, length)

			for z1 := range obj.Transactions {
				{
					// obj.Transactions[z1]
					if len(d.Buffer) < len(obj.Transactions[z1]) {
						return 0, encoder.ErrBufferUnderflow
					}
					copy(obj.Transactions[z1][:], d.Buffer[:len(obj.Transactions[z1])])
					d.Buffer = d.Buffer[len(obj.Transactions[z1]):]
				}

			}
		}
	}

	return uint64(len(buf) - len(d.Buffer)), nil
}

// decodeGetTxnsMessageExact decodes an object of type GetTxnsMessage from a buffer.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
// If the buffer is longer than required to decode the object, returns encoder.ErrRemainingBytes.
func decodeGetTxnsMessageExact(buf []byte, obj *GetTxnsMessage) error {
	if n, err := decodeGetTxnsMessage(buf, obj); err != nil {
		return err
	} else if n != uint64(len(buf)) {
		return encoder.ErrRemainingBytes
	}

	return nil
}
