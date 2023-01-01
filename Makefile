
App_Flusher ?= ./cdn-flusher/

run:
	go run $(App_Flusher) 

purge:
	go run $(App_Flusher) --purge=true

push:
	go run $(App_Flusher) --push=true
