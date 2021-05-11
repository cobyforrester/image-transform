package scalars

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
	"image"
	"image/png"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

type Image struct {
	Image image.Image
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (i *Image) UnmarshalGQL(v interface{}) error {
	image, ok := v.(string)
	if !ok {
		return fmt.Errorf("Image must be a string")
	}
	// Decode image
	// image := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABkAAAAqCAMAAABx2QBSAAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAACcUExURQAAAAAAAAAAAAUFBQAAAAAAAAAAAAAAABISEgAAAAAAAAAAAAAAAAAAAAAAAAcHBwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgICBkZGRoaGgUFBQAAAAcHB9nZ2QAAANbW1sPDwxoaGtXV1Xl5eaioqGVlZW1tba2trampqa+vr6SkpHh4eNTU1Kenp8rKymxsbHp6eqWlpaamplmx7/0AAAAedFJOUwDM/PMwrreN/E7lyCbuSPBRIo4zNCPvJfP+/vTL8wbHZEgAAADESURBVDjL7ZRHGoIwEEaHEkITFcXuBALEAvb7382CfCIknsC3fYvZvH8AXpjUc3XGmO561IQPQ7IQRRIjYpwUYk6GtQj65QE/HMp+UAnH3uI3W9t5iqmRY5vc8B/GyrBLZgGMeCoxKR9BeEYZGwreUWpyDdxEak490GOpQQYM/+ZvHkZZSE9Z1UpZYrhR1TtWFD/+sRLwjV1H7JZ+tcZ9S+yrNQIM1lnzVnqZDepxR4SL63v1N8FJ1PgIJtUm7MlEqz/FHZEZb+pMu1/4AAAAAElFTkSuQmCC"
	image = strings.ReplaceAll(image, "data:image/png;base64,", "")
	unbased, err := b64.StdEncoding.DecodeString(image)

	if err != nil {
		// panic("Cannot decode b64")
		return err
	}

	r := bytes.NewReader(unbased)
	im, err := png.Decode(r)
	if err != nil {
		// panic("Bad PNG")
		return err
	}
	// FOR TESTING
	// _, err = os.OpenFile("example.png", os.O_WRONLY|os.O_CREATE, 0777)
	// if err != nil {
	// 	panic("Cannot open file")
	// 	return err
	// }
	*i = Image{im}

	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (i Image) MarshalGQL(w io.Writer) {
	// Convert image to bytes.Buffer
	buf := new(bytes.Buffer)
	// Convert image to io.Writer
	err := png.Encode(buf, i.Image)
	if err != nil {
		panic("Invalid PNG")
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

	// Write to w
	io.WriteString(w, strconv.Quote(base64Encoding))
}

// MarshalID Lets redefine the base ID type to use a uuid as string
func MarshalID(id uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(id.String()))
	})
}

// UnmarshalID And the same for the unmarshaler
func UnmarshalID(v interface{}) (uuid.UUID, error) {
	switch v := v.(type) {
	case string:
		return uuid.MustParse(v), nil
	default:
		return uuid.Nil, fmt.Errorf("%s is not a string", v)
	}
}
