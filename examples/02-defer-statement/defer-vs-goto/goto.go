package main

import (
	"log"
	"sync"
)

type pmu struct {
	name  string
	type_ int //nolint:revive
}

var (
	pmus          []*pmu
	pmusLock      sync.Mutex
	pmuBusRunning int
)

func perfEventSysfsInitV1() error {
	var err error

	pmusLock.Lock()

	if err = busRegister(); err != nil {
		goto unlock
	}

	for _, pmu := range pmus {
		if pmu.name == "" || pmu.type_ < 0 {
			continue
		}

		if err = pmuDevAlloc(pmu); err != nil {
			log.Printf("Failed to register pmu: %s, reason %v", pmu.name, err)
		}
	}
	pmuBusRunning = 1
	err = nil

unlock:
	pmusLock.Unlock()

	return err
}

func busRegister() error {
	return nil
}

func pmuDevAlloc(_ *pmu) error {
	return nil
}
