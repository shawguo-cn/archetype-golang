.PHONY: gin
gin:
	rm ./build/gin -rf --verbose
	mkdir -p ./build/gin ;\
	go build -o ./build/gin/gin ./dominus/webapp/gin ;\
	cp ./dominus/config/config.json ./build/gin/config.json

spider:
	rm ./build/spider -rf --verbose
	mkdir -p ./build/spider ;\
	go build -o ./build/spider/spider ./dominus/intg/dht-spider/spider.go ;\