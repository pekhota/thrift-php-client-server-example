install:
	composer install

thrift-gen:
	mkdir -p ./app/thrift
	docker compose exec thrift /usr/local/bin/thrift -out /app/app/thrift --gen php:server -r /app/thrift/tutorial.thrift
	#thrift -out /app/app/thrift --gen php:server -r /app/thrift/tutorial.thrift

go-client:
	docker compose exec go-client go build
	docker compose exec go-client ./go-client

php-client:
	php ./clients/php-client/PhpClient.php

php-server:
	php -S 127.0.0.1:8080 -t ./public