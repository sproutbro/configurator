package tests_test

import (
	"testing"

	"github.com/sproutbro/configurator"
	"github.com/stretchr/testify/assert"
)

/*
*	요청에 응답테스트
*
*	고정형식으로 테스트 완료
 */
func TestConfigurator(t *testing.T) {
	v, f := configurator.GetConfig("ddd")
	assert.Equal(t, "", v)
	assert.Equal(t, false, f)

	v2, f2 := configurator.GetConfig("dev")
	assert.Equal(t, "dev-config", v2)
	assert.Equal(t, true, f2)
}
