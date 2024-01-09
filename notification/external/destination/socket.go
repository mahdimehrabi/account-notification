package destination

import (
	"encoding/json"
	"errors"
	"net"
	"sync"
)

var (
	ErrForbidden  = errors.New("access denied")
	ErrInternal   = errors.New("internal server error in challange-notification")
	ErrBadRequest = errors.New("bad request,please check your request data")
	ErrUndefined  = errors.New("undefined error")
)

type Socket struct {
	address       string
	conn          net.Conn
	connected     bool
	results       map[string]Result
	getResultCond *sync.Cond
}

func NewSocket(address string) *Socket {
	return &Socket{
		address:       address,
		connected:     false,
		results:       make(map[string]Result, 0),
		getResultCond: sync.NewCond(&sync.Mutex{}),
	}
}

func (s *Socket) Connect() error {
	conn, err := net.Dial("tcp", s.address)
	if err != nil {
		return err
	}
	s.conn = conn
	s.connected = true
	go s.SaveResults()
	return nil
}

func (s *Socket) Disconnect() error {
	s.connected = false
	return s.conn.Close()
}

func (s *Socket) SaveResults() {
	for {
		if !s.connected {
			return
		}
		d := json.NewDecoder(s.conn)
		result := Result{}
		if err := d.Decode(&result); err != nil {
			continue
		}
		s.results[result.ID] = result
		s.getResultCond.Broadcast()
	}
}

func (s *Socket) SendWaitJSON(data any, method string, msgID string) (*Result, error) {
	req := Request{
		Method: method,
		Data:   data,
		ID:     msgID,
	}
	if !s.connected {
		if err := s.Connect(); err != nil {
			return nil, err
		}
	}
	e := json.NewEncoder(s.conn)
	if err := e.Encode(req); err != nil {
		s.connected = false
		return nil, err
	}
	// get response
	return s.getResult(msgID)
}

func (s *Socket) getResult(msgID string) (*Result, error) {
	for {
		s.getResultCond.L.Lock()
		s.getResultCond.Wait()
		data, ok := s.results[msgID]
		if !ok {
			// its result of another request(pass)
			s.getResultCond.L.Unlock()
			continue
		}

		delete(s.results, msgID)
		s.getResultCond.L.Unlock()
		switch data.Status {
		case StatusOk:
			return &data, nil
		case StatusForbidden:
			return &data, ErrForbidden
		case StatusInternalError:
			return &data, ErrInternal
		case StatusBadRequest:
			return &data, ErrBadRequest
		default:
			return &data, ErrUndefined
		}
	}
}
