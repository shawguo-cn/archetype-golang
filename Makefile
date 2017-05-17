.PHONY: gin
gin:
	rm ./build/gin -rf --verbose
	mkdir -p ./build/gin ;\
	go build -o ./build/gin/gin ./dominus/webapp/gin ;\
	cp ./dominus/config/config.json ./build/gin/config.json