package db

type Page struct {
	Tx   *Tx
	Name []byte
	Text []byte
}

func (p *Page) bucket() []byte {
	return []byte("pages")
}

func (p *Page) get() ([]byte, error) {
	text := p.Tx.Bucket(p.bucket()).Get(p.Name)
	if text == nil {
		return nil, &Error{"page not found", nil}
	}
	return text, nil
}

// Load retrieves a page from the database.
func (p *Page) Load() error {
	text, err := p.get()
	if err != nil {
		return err
	}

	p.Text = text

	return nil
}

// Save commits the Page to the database.
func (p *Page) Save() error {
	assert(len(p.Name) != 0, "uninitialized page cannot be saved")
	return p.Tx.Bucket(p.bucket()).Put(p.Name, p.Text)
}
