package common_test

import (
	"testing"

	"github.com/mtkgo/common"
	"github.com/stretchr/testify/assert"
)

func TestIsBIM(t *testing.T) {
	assert.Equal(t, true, common.IsBIM("IFC", false))
	assert.Equal(t, true, common.IsBIM("XXX", true))
	assert.Equal(t, false, common.IsBIM("glTF", false))
}

func TestIsCAD(t *testing.T) {
	assert.Equal(t, true, common.IsCAD("STEP", false))
	assert.Equal(t, true, common.IsCAD("XXX", true))
	assert.Equal(t, false, common.IsCAD("IFC", false))
}
