IMAGE=midorigreen/pwg-server

build:
	docker build -t $(IMAGE) .

run:
	docker run -it --rm -d --name $(IMAGE) -p 8000:80 -e LBP_SECRET_TOKEN=$(LBP_SECRET_TOKEN) -e LBP_ACCESS_TOKEN=$(LBP_ACCESS_TOKEN) $(IMAGE)
