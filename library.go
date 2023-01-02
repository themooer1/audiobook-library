package library

type AudioBookLibrary struct {
	AudioBooksByName map[string]AudioBook
}

func (a *AudioBookLibrary) Initialize() {
	a.AudioBooksByName = make(map[string]AudioBook)
}

func (lib *AudioBookLibrary) Add(book AudioBook) {
	lib.AudioBooksByName[book.Title] = book
}

func (lib *AudioBookLibrary) Get(bookTitle string) AudioBook {
	return lib.AudioBooksByName[bookTitle]
}

func (lib *AudioBookLibrary) ForEach(callback func(string, *AudioBook) error) error {
	var err error = nil

	for name, book := range lib.AudioBooksByName {
		err = callback(name, &book)

		if err != nil {
			break
		}
	}

	return err
}
