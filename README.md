# kode-task

GET: /note/all - получить список всех заметок.

GET: /note/first - первая добавленная заметка.

GET: /note/last - последняя добавленная заметка.

POST: /note/add - добавить заметку. В тело запроса нужно влить JSON модель заметки, пример ниже:

{"value":"test", "expiration":100}

DELETE: /note/delete/ID - удалить заметку по ее ID'у.
