package bca

type Service struct {
	reader     *Reader
	mime       *MIMEReader
	html       *HTMLExtractor
	dom        *DOMParser
	text       *TextExtractor
	parser     *Parser
	normalizer *Normalizer
}

func NewService() *Service {
	return &Service{
		reader:     NewReader(),
		mime:       NewMIMEReader(),
		html:       NewHTMLExtractor(),
		dom:        NewDOMParser(),
		text:       NewTextExtractor(),
		parser:     New(),
		normalizer: NewNormalizer(),
	}
}

func (s *Service) ParseFile(path string) (*TransactionData, error) {

	data, err := s.reader.Read(path)
	if err != nil {
		return nil, err
	}

	entity, err := s.mime.Parse(data)
	if err != nil {
		return nil, err
	}

	htmlContent, err := s.html.ExtractHTML(entity)
	if err != nil {
		return nil, err
	}

	root, err := s.dom.Parse(htmlContent)
	if err != nil {
		return nil, err
	}

	plainText := s.text.Extract(root)

	raw, err := s.parser.Parse(plainText)
	if err != nil {
		return nil, err
	}

	return s.normalizer.Normalize(raw)
}
