IMAGE=midorigreen/pwg-server
NAME=pwg-server
HOSTPORT=3333

build:
	docker build -t $(IMAGE) .

run-front:
	docker run -it --rm --name $(NAME) -p $(HOSTPORT):80 -e PORT=80 -e LBP_SECRET_TOKEN=$(LBP_SECRET_TOKEN) -e LBP_ACCESS_TOKEN=$(LBP_ACCESS_TOKEN) $(IMAGE)

run:
	docker run -d -it --rm --name $(NAME) -p $(HOSTPORT):80 -e PORT=80 -e LBP_SECRET_TOKEN=$(LBP_SECRET_TOKEN) -e LBP_ACCESS_TOKEN=$(LBP_ACCESS_TOKEN) $(IMAGE)
	docker ps

restart: build
	docker stop $(NAME)
	docker run -d -it --rm --name $(NAME) -p $(HOSTPORT):80 -e PORT=80 -e LBP_SECRET_TOKEN=$(LBP_SECRET_TOKEN) -e LBP_ACCESS_TOKEN=$(LBP_ACCESS_TOKEN) $(IMAGE)
	docker ps

deploy:
	heroku container:push --app linebot-pwg web
	heroku container:release web

heroku-config:
	heroku config:set LBP_ACCESS_TOKEN=$(LBP_ACCESS_TOKEN)
	heroku config:set LBP_SECRET_TOKEN=$(LBP_SECRET_TOKEN)

local:
	ngrok http $(HOSTPORT)
