package tests_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/sproutbro/parserx/internal/parser"
	pb "github.com/sproutbro/parserx/proto/pbparser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
*	parser manager로 변경 JSON YAML ENV 파싱확인
 */
func TestConfigurator(t *testing.T) {
	// 파서 생성
	var decode = make(map[string]interface{})

	// sever 생성
	f := parser.New(nil)

	// ENV 테스트
	resp, err := f.ENV(context.Background(),
		&pb.Req{
			Key: ".env",
		})
	require.NoError(t, err)
	err = json.Unmarshal([]byte(resp.Json), &decode)
	assert.NoError(t, err)
	assert.Equal(t, "dev", decode["ENV"])

}

func TestYAML(t *testing.T) {

	f := parser.New(nil)

	var decode = make(map[string]interface{})

	// YAML 테스트
	resp2, err := f.YAML(context.Background(),
		&pb.Req{
			Key: "app.yaml",
		})
	require.NoError(t, err)
	err = json.Unmarshal([]byte(resp2.Json), &decode)

	assert.NoError(t, err)
	assert.Equal(t, "APP_ENV", decode["env"])
}

func TestJSON(t *testing.T) {

	var decode = make(map[string]interface{})

	f := parser.New(nil)

	// JOSN 파서 테스트
	resp3, err := f.JSON(context.Background(),
		&pb.Req{
			Key: "app.json",
		})
	require.NoError(t, err)
	assert.IsType(t, "json", resp3.Json)

	err = json.Unmarshal([]byte(resp3.Json), &decode)
	assert.NoError(t, err)
	assert.Equal(t, "TEST", decode["ENV"])
	assert.Equal(t, "SECRET", decode["SECRET"])
}
