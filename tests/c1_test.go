package tests_test

import (
	"testing"

	"github.com/sproutbro/filex/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
*	parser manager로 변경 JSON YAML ENV 파싱확인
 */
func TestConfigurator(t *testing.T) {
	// 파서 생성
	f := internal.New()

	// ENV 디코딩 테스트
	decode, err := f.ENV(".env")
	require.NoError(t, err)
	assert.Equal(t, "dev", decode["ENV"])

	// json 디코딩 테스트
	dcodejson, err := f.JSON("app.json")
	require.NoError(t, err)
	assert.Equal(t, "TEST", dcodejson["ENV"])
	assert.Equal(t, "SECRET", dcodejson["SECRET"])

	// YMAL 디코딩 테스트
	decodeYAML, err := f.YAML("app.yaml")
	require.NoError(t, err)
	assert.Equal(t, "APP_ENV", decodeYAML["env"])
}
