package models

import "sync"

func ConnectToDatabases() {
	//? Connecting to databases simultaneously
	var connectToDatabasesWG sync.WaitGroup

	connectToDatabasesWG.Add(1)

	go ConnectToPostgres(&connectToDatabasesWG)
	// go ConnectToRedis(&connectToDatabasesWG)

	connectToDatabasesWG.Wait()
}
