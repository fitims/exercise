.PHONY: deps clean live

deps:
	go get -u ./...

clean:
	rm -rf ./output

live: clean
	GOOS=linux go build -o output/historic_api ./historic_api.go
	cp config.yml output/config.yml
	cp -rf assets/* output/assets

deploy: live
	scp output/historic_api nzxApi:/go-web/historic_api
	scp config.yml nzxApi:/go-web/config.yml
	scp -pr output/assets nzxApi:/go-web/assets

cleaner: clean
	GOOS=linux go build -o output/cleaner ./cleaner.go

deploy-cleaner: cleaner
	scp output/cleaner nzxApi:/home/ubuntu/cleaner
