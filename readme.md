# Apache Thrift PHP client-server example

Thrift golang client example also available here.

## Getting started 

1. Run composer install. Tested on PHP 8.0.10 (cli). 
2. Run `docker compose up -d`
3. Run `make php-client`
4. Enjoy

## Golang client

1. Run `make go-client`

## Update thrift files

1. Run `make thrift-gen`

## Other

1. directory: `./app/lib` relates to this one: `https://github.com/apache/thrift/tree/master/lib/php/lib`.
If you want to use it inside your project - download it via wget or something.
2. You can use `makefile`, `docker-compose.yml` as some kind of documentation as they are quite self explaining. 