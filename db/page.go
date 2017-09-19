package db

// Errors
var (
	ErrPageNotFound = &Error{"page not found", nil}
	ErrNoPageName   = &Error{"no page name", nil}
)

// Page represents a Wiki page
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
		return nil, ErrPageNotFound
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
	if len(p.Name) == 0 {
		return ErrNoPageName
	}

	return p.Tx.Bucket(p.bucket()).Put(p.Name, p.Text)
}
