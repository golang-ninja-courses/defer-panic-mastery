package main

import "log"

func perfEventSysfsInitV2() error {
	pmusLock.Lock()

	if err := busRegister(); err != nil {
		pmusLock.Unlock()
		return err
	}

	for _, pmu := range pmus {
		if pmu.name == "" || pmu.type_ < 0 {
			continue
		}

		if err := pmuDevAlloc(pmu); err != nil {
			log.Printf("Failed to register pmu: %s, reason %v", pmu.name, err)
		}
	}
	pmuBusRunning = 1

	pmusLock.Unlock()
	return nil
}
