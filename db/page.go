package db

import "github.com/russross/blackfriday"

var (
	// ErrPageNotFound is returned when a page does not exist.
	ErrPageNotFound = &Error{"page not found", nil}
)

type Page struct {
	Tx   *Tx
	Name string
	Text []byte
}

func (p *Page) get() ([]byte, error) {
	text := p.Tx.Bucket([]byte("pages")).Get([]byte(p.Name))
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
	assert(p.Name != "", "uninitialized page cannot be saved")
	return p.Tx.Bucket([]byte("pages")).Put([]byte(p.Name), p.Text)
}

func (p *Page) TextString(text string) {
	p.Text = []byte(text)
}

func (p *Page) ParsedText() []byte {
	return blackfriday.MarkdownCommon([]byte(p.Text))
}

func (p *Page) ParsedTextString() string {
	return string(p.ParsedText())
}

type Pages []*Page

func (p Pages) Len() int           { return len(p) }
func (p Pages) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Pages) Less(i, j int) bool { return p[i].Name < p[j].Name }
