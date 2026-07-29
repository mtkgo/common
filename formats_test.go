package common_test

import (
	"testing"

	"github.com/mtkgo/common"
	"github.com/stretchr/testify/assert"
)

func TestIsBIM(t *testing.T) {
	assert.Equal(t, true, common.IsBIM("RVM"))
	assert.Equal(t, true, common.IsBIM("IFC"))

	assert.Equal(t, false, common.IsBIM("glTF"))
	assert.Equal(t, false, common.IsBIM("XXX"))
}

func TestIsCAD(t *testing.T) {
	assert.Equal(t, true, common.IsCAD("BREP"))
	assert.Equal(t, true, common.IsCAD("IGES"))
	assert.Equal(t, true, common.IsCAD("STEP"))

	assert.Equal(t, false, common.IsCAD("IFC"))
	assert.Equal(t, false, common.IsCAD("XXX"))
}

func TestGetFormatID(t *testing.T) {
	assert.Equal(t, common.FormatID("FBX"), common.GetFormatID("a.fbx"))
	assert.Equal(t, common.FormatID("FBX"), common.GetFormatID("a.FBX"))
	assert.Equal(t, common.FormatID(""), common.GetFormatID("a.ABC"))
}
