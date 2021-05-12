package helper

import (
	"bytes"
	b64 "encoding/base64"
	"image"
	"image/png"
	"net/http"
	"strings"

	"github.com/disintegration/imaging"
)

func B64ToImage(s string) (image.Image, error) {
	s = s[strings.Index(s, ",")+1:] // removes: data:image/png;base64,
	unbased, err := b64.StdEncoding.DecodeString(s)
	if err != nil {
		// panic("Cannot decode b64")
		return nil, err
	}

	r := bytes.NewReader(unbased)
	im, err := imaging.Decode(r)
	if err != nil {
		// panic("Bad PNG")
		return nil, err
	}
	return im, nil
}

func ImageToB64(image image.Image) (string, error) {
	// Convert image to bytes.Buffer
	buf := new(bytes.Buffer)
	// Convert image to io.Writer
	err := png.Encode(buf, image)
	if err != nil {
		// panic("Invalid PNG")
		return "", err
	}

	// Convert io.Writer to correct string
	var base64Encoding string
	mimeType := http.DetectContentType(buf.Bytes())
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}
	base64Encoding += b64.StdEncoding.EncodeToString(buf.Bytes())

	return base64Encoding, nil
}

func B64ToImageConfig(s string) (image.Config, error) {
	s = s[strings.Index(s, ",")+1:] // removes: data:image/png;base64,
	unbased, err := b64.StdEncoding.DecodeString(s)
	if err != nil {
		panic("Cannot decode b64")
		// return nil, err
	}
	r := bytes.NewReader(unbased)
	im, _, err := image.DecodeConfig(r)
	if err != nil {
		panic("Bad PNG")
		// return nil, err
	}
	return im, nil
}
