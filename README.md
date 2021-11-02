# USAGE: go run cmd/main.go

GET: /note/ - получить список всех заметок.

GET: /note/first - первая добавленная заметка.

GET: /note/last - последняя добавленная заметка.

POST: /note/ - добавить заметку. В тело запроса нужно влить JSON модель заметки, пример ниже:

{"value":"test", "expiration":100}

DELETE: /note/{ID} - удалить заметку по ее ID'у.


