package bca

import (
	"testing"
)

// func TestReadFixture(t *testing.T) {

// 	reader := NewReader()

// 	data, err := reader.Read("../../../../testdata/bca/credit_card_transaction.eml")

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(data) == 0 {
// 		t.Fatal("fixture should not be empty")
// 	}

// }

// func TestReadMIME(t *testing.T) {

// 	reader := NewReader()

// 	data, err := reader.Read("../../../../testdata/bca/credit_card_transaction.eml")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	mimeReader := NewMIMEReader()

// 	entity, err := mimeReader.Parse(data)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if entity == nil {
// 		t.Fatal("entity should not be nil")
// 	}
// }

// func TestInspectMIME(t *testing.T) {

// 	reader := NewReader()

// 	data, err := reader.Read("../../../../testdata/bca/credit_card_transaction.eml")

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	mime := NewMIMEReader()

// 	entity, err := mime.Parse(data)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	Inspect(entity)
// }

// func TestReadBody(t *testing.T) {

// 	reader := NewReader()

// 	data, err := reader.Read("../../../../testdata/bca/credit_card_transaction.eml")

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	mime := NewMIMEReader()

// 	entity, err := mime.Parse(data)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	bodyReader := NewBodyReader()

// 	body, err := bodyReader.Read(entity)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(body) == 0 {
// 		t.Fatal("body should not be empty")
// 	}

// 	t.Log(string(body))
// }

// func TestExtractHTML(t *testing.T) {

// 	reader := NewReader()

// 	data, err := reader.Read("../../../../testdata/bca/credit_card_transaction.eml")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	mime := NewMIMEReader()

// 	entity, err := mime.Parse(data)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	extractor := NewHTMLExtractor()

// 	html, err := extractor.Extract(entity)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(html) == 0 {
// 		t.Fatal("html should not be empty")
// 	}

// 	t.Log(html[:500])
// }

// func TestExtractHTMLHeader(t *testing.T) {

// 	reader := NewReader()

// 	data, err := reader.Read("../../../../testdata/bca/credit_card_transaction.eml")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	mime := NewMIMEReader()

// 	entity, err := mime.Parse(data)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	extractor := NewHTMLExtractor()

// 	_, err = extractor.ExtractHTML(entity)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// }

// func TestExtractHTMLOnly(t *testing.T) {

// 	reader := NewReader()

// 	data, err := reader.Read("../../../../testdata/bca/credit_card_transaction.eml")

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	mime := NewMIMEReader()

// 	entity, err := mime.Parse(data)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	htmlExtractor := NewHTMLExtractor()

// 	html, err := htmlExtractor.ExtractHTML(entity)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if html == "" {
// 		t.Fatal("html should not be empty")
// 	}

// 	t.Log(html[:500])
// }

// func TestParseHTML(t *testing.T) {

// 	reader := NewReader()

// 	data, err := reader.Read("../../../../testdata/bca/credit_card_transaction.eml")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	mime := NewMIMEReader()

// 	entity, err := mime.Parse(data)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	htmlExtractor := NewHTMLExtractor()

// 	htmlContent, err := htmlExtractor.ExtractHTML(entity)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	dom := NewDOMParser()

// 	node, err := dom.Parse(htmlContent)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if node == nil {
// 		t.Fatal("DOM should not be nil")
// 	}
// }

// func TestExtractText(t *testing.T) {

// 	reader := NewReader()

// 	data, err := reader.Read("../../../../testdata/bca/credit_card_transaction.eml")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	mime := NewMIMEReader()

// 	entity, err := mime.Parse(data)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	htmlExtractor := NewHTMLExtractor()

// 	htmlContent, err := htmlExtractor.ExtractHTML(entity)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	dom := NewDOMParser()

// 	node, err := dom.Parse(htmlContent)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	textExtractor := NewTextExtractor()

// 	text := textExtractor.Extract(node)

// 	if text == "" {
// 		t.Fatal("text should not be empty")
// 	}

// 	t.Log(text)
// }

func TestFindMerchant(t *testing.T) {

	reader := NewReader()

	data, err := reader.Read("../../../../testdata/bca/credit_card_transaction.eml")
	if err != nil {
		t.Fatal(err)
	}

	mime := NewMIMEReader()

	entity, err := mime.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	htmlExtractor := NewHTMLExtractor()

	htmlContent, err := htmlExtractor.ExtractHTML(entity)
	if err != nil {
		t.Fatal(err)
	}

	dom := NewDOMParser()

	node, err := dom.Parse(htmlContent)
	if err != nil {
		t.Fatal(err)
	}

	textExtractor := NewTextExtractor()

	text := textExtractor.Extract(node)

	parser := New()

	merchant := parser.FindValue(text, "Merchant / ATM")

	if merchant == "" {
		t.Fatal("merchant not found")
	}

	t.Log("Merchant:", merchant)
}

func TestFindAmount(t *testing.T) {

	reader := NewReader()

	data, err := reader.Read("../../../../testdata/bca/credit_card_transaction.eml")
	if err != nil {
		t.Fatal(err)
	}

	mime := NewMIMEReader()

	entity, err := mime.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	htmlExtractor := NewHTMLExtractor()

	htmlContent, err := htmlExtractor.ExtractHTML(entity)
	if err != nil {
		t.Fatal(err)
	}

	dom := NewDOMParser()

	node, err := dom.Parse(htmlContent)
	if err != nil {
		t.Fatal(err)
	}

	textExtractor := NewTextExtractor()

	text := textExtractor.Extract(node)

	parser := New()

	amount := parser.FindValue(text, "Sejumlah")

	if amount == "" {
		t.Fatal("amount not found")
	}

	t.Log("Amount:", amount)
}

func TestNormalizeAmount(t *testing.T) {

	n := NewNormalizer()

	amount, err := n.NormalizeAmount("Rp201.000,00")

	if err != nil {
		t.Fatal(err)
	}

	if amount != 201000 {
		t.Fatalf("expected 201000 got %f", amount)
	}
}

func TestNormalizeDate(t *testing.T) {

	n := NewNormalizer()

	date, err := n.NormalizeDate(
		"25-06-2026 19:32:45 WIB",
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(date)
}
