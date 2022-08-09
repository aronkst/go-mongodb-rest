run:
	docker compose up

start:
	docker compose up -d

stop:
	docker compose stop

logs:
	docker compose logs -f

test-list:
	curl --request GET --url 'http://localhost:3000/collection/test'

test-list-find:
	curl --request GET --url 'http://localhost:3000/collection/test?abc=123'

test-show:
	curl --request GET --url 'http://localhost:3000/collection/test/$(_id)'

test-insert:
	curl --request POST --url 'http://localhost:3000/collection/test' --header 'Content-Type: application/json' --data '{"abc": 123}'

test-replace:
	curl --request PUT --url 'http://localhost:3000/collection/test' --header 'Content-Type: application/json' --data '{"cba": 321}'

test-replace-find:
	curl --request PUT --url 'http://localhost:3000/collection/test?abc=123' --header 'Content-Type: application/json' --data '{"abc": 123}'

test-update:
	curl --request PATCH --url 'http://localhost:3000/collection/test' --header 'Content-Type: application/json' --data '{"xyz": 987}'

test-update-find:
	curl --request PATCH --url 'http://localhost:3000/collection/test?abc=123' --header 'Content-Type: application/json' --data '{"xyz": 987}'

test-delete:
	curl --request DELETE --url 'http://localhost:3000/collection/test'

test-delete-find:
	curl --request DELETE --url 'http://localhost:3000/collection/test?abc=123'

test-query:
	curl --request POST --url 'http://localhost:3000/query/test' --header 'Content-Type: application/json' --data '{"find": {"abc": {"$$eq": 123}},"limit": 100,"skip": 0,"sort": {"abc": 1}}'

test-count:
	curl --request POST --url 'http://localhost:3000/count/test' --header 'Content-Type: application/json' --data '{"find": {"abc": {"$$eq": 123}}}'

test-paginate:
	curl --request POST --url 'http://localhost:3000/paginate/test' --header 'Content-Type: application/json' --data '{"find": {"abc": {"$$eq": 123}},"sort":{"abc": 1},"page": 1,"perpage": 5}'
