package ocrlib

type OCR interface {
	GetResult(path string) (string, error)
}
