package bmwire_test

import (
	"io"

	"github.com/jimmysong/bmwire"
)

// fakeMessage implements the bmwire.Message interface and is used to force
// encode errors in messages.
type fakeMessage struct {
	command        string
	payload        []byte
	forceEncodeErr bool
	forceLenErr    bool
}

// Decode doesn't do anything.  It just satisfies the bmwire.Message
// interface.
func (msg *fakeMessage) Decode(r io.Reader) error {
	return nil
}

// Encode writes the payload field of the fake message or forces an error
// if the forceEncodeErr flag of the fake message is set.  It also satisfies the
// bmwire.Message interface.
func (msg *fakeMessage) Encode(w io.Writer) error {
	if msg.forceEncodeErr {
		err := &bmwire.MessageError{
			Func:        "fakeMessage.Encode",
			Description: "intentional error",
		}
		return err
	}

	_, err := w.Write(msg.payload)
	return err
}

// Command returns the command field of the fake message and satisfies the
// bmwire.Message interface.
func (msg *fakeMessage) Command() string {
	return msg.command
}

// MaxPayloadLength returns the length of the payload field of fake message
// or a smaller value if the forceLenErr flag of the fake message is set.  It
// satisfies the bmwire.Message interface.
func (msg *fakeMessage) MaxPayloadLength() uint32 {
	lenp := uint32(len(msg.payload))
	if msg.forceLenErr {
		return lenp - 1
	}

	return lenp
}
