IMAGE=midorigreen/pwg-server

build:
	docker build -t $(IMAGE) .

run:
	docker run -it --rm --name pwg-server -p 8000:80 -e PORT=80 -e LBP_SECRET_TOKEN=$(LBP_SECRET_TOKEN) -e LBP_ACCESS_TOKEN=$(LBP_ACCESS_TOKEN) $(IMAGE)

deploy:
	heroku container:push --app linebot-pwg web
	heroku container:release web

heroku-config:
	heroku config:set LBP_ACCESS_TOKEN=$(LBP_ACCESS_TOKEN)
	heroku config:set LBP_SECRET_TOKEN=$(LBP_SECRET_TOKEN)
