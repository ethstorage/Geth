// Code generated by rlpgen. DO NOT EDIT.

//go:build !norlpgen
// +build !norlpgen

package types

import "github.com/ethereum/go-ethereum/rlp"
import "io"

func (obj *Withdrawal) EncodeRLP(_w io.Writer) error {
	w := rlp.NewEncoderBuffer(_w)
	_tmp0 := w.List()
	w.WriteUint64(obj.Index)
	w.WriteUint64(obj.Validator)
	w.WriteBytes(obj.Address[:])
	if obj.Amount == nil {
		w.Write(rlp.EmptyString)
	} else {
		if obj.Amount.Sign() == -1 {
			return rlp.ErrNegativeBigInt
		}
		w.WriteBigInt(obj.Amount)
	}
	w.ListEnd(_tmp0)
	return w.Flush()
}