package memeng

// CloseSession implements the Engine interface
func (eng *engine) CloseSession(key string) error {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	if _, exists := eng.sessions[key]; exists {
		delete(eng.sessions, key)
	}

	return nil
}
