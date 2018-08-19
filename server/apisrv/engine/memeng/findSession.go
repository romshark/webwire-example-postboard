package memeng

import (
	"time"

	wwr "github.com/qbeon/webwire-go"
)

// FindSession implements the Engine interface
func (eng *engine) FindSession(key string) (wwr.SessionLookupResult, error) {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	if session, exists := eng.sessions[key]; exists {
		// Session found, update last lookup time
		// to prevent it from being garbage collected too early
		session.LastLookup = time.Now().UTC()
		eng.sessions[key] = session

		info := session.Info
		return wwr.SessionLookupResult{
			Creation:   session.Creation,
			LastLookup: session.LastLookup,
			Info:       wwr.SessionInfoToVarMap(info),
		}, nil
	}

	// Not found
	return wwr.SessionLookupResult{}, wwr.SessNotFoundErr{}

}
