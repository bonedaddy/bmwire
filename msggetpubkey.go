// Copyright (c) 2013-2015 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package bmwire

import (
	"fmt"
	"io"
	"time"
)

const (
	// Starting in version 4, Ripe is derived from the tag and not
	// sent directly
	TagBasedRipeVersion = 4
)

type MsgGetPubKey struct {
	Nonce        uint64
	ExpiresTime  time.Time
	ObjectType   ObjectType
	Version      uint64
	StreamNumber uint64
	Ripe         RipeHash
	Tag          ShaHash
}

// BtcDecode decodes r using the bitcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgGetPubKey) BtcDecode(r io.Reader, pver uint32) error {
	var sec int64
	err := readElements(r, &msg.Nonce, &sec, &msg.ObjectType)
	if err != nil {
		return err
	}

	if msg.ObjectType != ObjectTypeGetPubKey {
		str := fmt.Sprintf("Object Type should be %d, but is %d",
			ObjectTypeGetPubKey, msg.ObjectType)
		return messageError("BtcDecode", str)
	}

	msg.ExpiresTime = time.Unix(sec, 0)
	msg.Version, err = readVarInt(r, pver)
	if err != nil {
		return err
	}

	msg.StreamNumber, err = readVarInt(r, pver)
	if err != nil {
		return err
	}

	if msg.Version >= TagBasedRipeVersion {
		err = readElement(r, &msg.Tag)
		if err != nil {
			return err
		}
	} else {
		err = readElement(r, &msg.Ripe)
		if err != nil {
			return err
		}
	}

	return err
}

// BtcEncode encodes the receiver to w using the bitcoin protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgGetPubKey) BtcEncode(w io.Writer, pver uint32) error {
	err := writeElements(w, msg.Nonce, msg.ExpiresTime.Unix(), msg.ObjectType)
	if err != nil {
		return err
	}

	err = writeVarInt(w, pver, msg.Version)
	if err != nil {
		return err
	}

	err = writeVarInt(w, pver, msg.StreamNumber)
	if err != nil {
		return err
	}

	if msg.Version >= TagBasedRipeVersion {
		err = writeElement(w, msg.Tag)
		if err != nil {
			return err
		}
	} else {
		err = writeElement(w, msg.Ripe)
		if err != nil {
			return err
		}
	}

	return err
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgGetPubKey) Command() string {
	return CmdObject
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgGetPubKey) MaxPayloadLength(pver uint32) uint32 {
	return uint32(8 + 8 + 4 + 8 + 8 + 32)
}

func (msg *MsgGetPubKey) String() string {
	return fmt.Sprintf("msgobject getpubkey: v%d %d %s %d %x %x", msg.Version, msg.Nonce, msg.ExpiresTime, msg.StreamNumber, msg.Ripe, msg.Tag)
}

// NewMsgGetPubKey returns a new object message that conforms to the
// Message interface using the passed parameters and defaults for the remaining
// fields.
func NewMsgGetPubKey(nonce uint64, expires time.Time, version, streamNumber uint64, ripe RipeHash, tag ShaHash) *MsgGetPubKey {

	// Limit the timestamp to one second precision since the protocol
	// doesn't support better.
	return &MsgGetPubKey{
		Nonce:        nonce,
		ExpiresTime:  expires,
		ObjectType:   ObjectTypeGetPubKey,
		Version:      version,
		StreamNumber: streamNumber,
		Ripe:         ripe,
		Tag:          tag,
	}
}