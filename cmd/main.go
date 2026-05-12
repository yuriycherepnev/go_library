/* TODO
в гайде по чистой архитектуре:
https://www.youtube.com/watch?v=PqQyCFygiZg&t=16s

мы подошли к моменту подключения базы данных (10 минута)
в видео используется монго дб,
мой код надо будет адаптировать под mysql
*/

package main

import (
	authorStorage "go-library/internal/adapters/db/author"
	bookStorage "go-library/internal/adapters/db/book"
	authorService "go-library/internal/usecase/author"
	bookService "go-library/internal/usecase/book"
)

func main() {
	book := bookStorage.NewStorage()
	author := authorStorage.NewStorage()

	bookUseCase := bookService.NewService(book)
	authorUseCase := authorService.NewService(author)

}
