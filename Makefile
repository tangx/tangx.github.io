
App_Flusher ?= ./cdn-flusher/

purge: download
	./qcloud-cdn-flusher -c purge.flusher.yml

push: download
	./qcloud-cdn-flusher -c push.flusher.yml

download:
	wget -c https://github.com/tangx/qcloud-cdn-flusher/releases/latest/download/qcloud-cdn-flusher-linux-amd64 -O ./qcloud-cdn-flusher && chmod +x ./qcloud-cdn-flusher
