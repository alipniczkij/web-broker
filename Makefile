build:
	docker build -t web-broker .

run:
	docker run --name=go-web-broker -p 8000:8000 web-broker