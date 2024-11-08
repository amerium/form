package form_test

import (
	"io"
	"net/url"
	"reflect"
	"testing"

	"github.com/amerium/form/v6"
	"github.com/amerium/form/v6/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncoder_Encode_deep_embed(t *testing.T) {
	type S struct {
		io.Writer

		Header string `form:"header"`
	}

	enc := form.NewEncoder()

	enc.SetMode(form.ModeExplicit)
	enc.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Name
	})

	v := S{}
	v.Writer = form.MakeEmbeddedUnexported()
	v.Header = "foo"

	e, err := enc.Encode(v)
	require.NoError(t, err)

	assert.Equal(t, url.Values{"Header": []string{"foo"}}, e)
}

func TestEncoder_Encode_unexported_embed(t *testing.T) {
	type S struct {
		io.Writer
		internal.DeeperEmbedded

		Header string `form:"header"`
	}

	enc := form.NewEncoder()

	enc.SetMode(form.ModeExplicit)

	v := S{}
	v.SetDeeplyEmbedded("baz")
	v.Writer = internal.MakeWriterWithExported()
	v.Header = "foo"

	collect := map[string]interface{}{}
	e, err := enc.Encode(v, collect)
	require.NoError(t, err)

	assert.Equal(t, url.Values{"deeply-embedded": []string{"baz"}, "header": []string{"foo"}, "writer-exported": []string{"bar"}}, e)
	assert.Equal(t, map[string]interface{}{"deeply-embedded": "baz", "header": "foo", "writer-exported": "bar"}, collect)
}
