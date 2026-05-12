/* TODO
в гайде по чистой архитектуре:
https://www.youtube.com/watch?v=PqQyCFygiZg&t=16s

мы подошли к моменту подключения базы данных (10 минута)
в видео используется монго дб,
мой код надо будет адаптировать под mysql
*/

package cmd

import (
	authorStorage "ca-template/internal/adapters/db/author"
	bookStorage "ca-template/internal/adapters/db/book"
	authorService "ca-template/internal/domain/author"
	bookService "ca-template/internal/domain/book"
)

func main() {
	book := bookStorage.NewStorage()
	author := authorStorage.NewStorage()

	bookUseCase := bookService.NewService(book)
	authorUseCase := authorService.NewService(author)

}
