package memeng

import wwr "github.com/qbeon/webwire-go"

// SaveSession implements the Engine interface
func (eng *engine) SaveSession(newSession *wwr.Session) error {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	// Save a clone of the provided session object
	eng.sessions[newSession.Key] = newSession.Clone()
	return nil
}
