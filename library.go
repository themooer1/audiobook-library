package library

type AudioBookLibrary struct {
	audiobooks         map[string]AudioBook
	AudioBooksByName   map[string][]*AudioBook
	AudioBooksByAuthor map[string][]*AudioBook
}

func (a *AudioBookLibrary) Initialize() {
	a.audiobooks = make(map[string]AudioBook, 0)
	a.AudioBooksByName = make(map[string][]*AudioBook)
	a.AudioBooksByAuthor = make(map[string][]*AudioBook)
}

func GetAudioBookHashKey(book *AudioBook) string {
	return book.Title + book.Author
}

func (lib *AudioBookLibrary) Add(book AudioBook) {
	var hashKey = GetAudioBookHashKey(&book)
	lib.audiobooks[hashKey] = book

	lib.AudioBooksByName[book.Title] = append(lib.AudioBooksByName[book.Title], &book)
	lib.AudioBooksByAuthor[book.Author] = append(lib.AudioBooksByAuthor[book.Author], &book)

}

func (lib *AudioBookLibrary) Get(bookTitle string) *AudioBook {
	return lib.AudioBooksByName[bookTitle][0]
}

func (lib *AudioBookLibrary) GetByAuthor(authorName string) []*AudioBook {
	return lib.AudioBooksByAuthor[authorName]
}

func (lib *AudioBookLibrary) GetByName(bookName string) []*AudioBook {
	return lib.AudioBooksByName[bookName]
}

func (lib *AudioBookLibrary) ForEach(callback func(string, *AudioBook) error) error {
	var err error = nil

	for name, book := range lib.audiobooks {
		err = callback(name, &book)

		if err != nil {
			break
		}
	}

	return err
}
