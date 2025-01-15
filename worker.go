package pkoderlite

import (
	"context"
	"fmt"
)

// BaseWorker is a foundational struct for a worker. It contains common fields
// required for all types of workers that manage streaming processes.
type BaseWorker struct {
	StreamID  string   // ID of the stream.
	Protocol  string   // Protocol being used for the stream.
	ListenURI string   // URI where the worker listens for data.
	Timeout   int      // Timeout duration for operations.
	RtmpDst   []string // List of destinations for RTMP.
	MpegTs    []string
	Ctx       context.Context // Context to potentially control the lifecycle of the worker.
}

// Worker interface defines a generic worker that performs a specific task.
type Worker interface {
	// Do is the main method where the actual task of the worker is executed.
	Do() error
}

// Do initiates the streaming process for the BaseWorker. It establishes an encoder based on the worker's
// configuration and begins the streaming session. During the session, it continuously monitors messages
// from the encoder and handles errors, interruptions, or timeouts. If a context cancellation is detected,
// or a predefined error code (-60, indicating timeout) is received from the encoder, the method ensures
// a graceful shutdown of the encoder and invokes the appropriate callback methods.
//
// Returns:
//
//	nil if the streaming session ends without errors.
//	error if an issue arises during the streaming process (e.g., timeout, cancellation, etc.)
func (w *BaseWorker) Do() error {
	if len(w.RtmpDst) == 0 {
		return fmt.Errorf("unable to find a valid destination")
	}

	encoder, err := NewEncoder(
		w.Protocol,
		w.ListenURI,
		w.Timeout,
		w.RtmpDst,
		w.MpegTs,
	)
	if err != nil {
		return err
	}

	if err := encoder.Start(); err != nil {
		return err
	}

	for {
		select {
		case <-w.Ctx.Done():
			encoder.Cancel()
			encoder.Wait()
			err = fmt.Errorf("canceled")
			goto exit
		case msg := <-encoder.Chan:
			if val, ok := msg.(int); ok {
				if val == -60 {
					err = fmt.Errorf("timeout")
				} else {
					err = nil
				}
				goto exit
			}
		}
	}
exit:
	return err
}
