package common

type Property struct {
	ID    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Value any    `json:"value"`
}

type NodePatch struct {
	Matrix      *[16]float64 `json:"matrix,omitempty"`
	Translation *[3]float64  `json:"translation,omitempty"`
	Rotation    *[4]float64  `json:"rotation,omitempty"`
	Scale       *[3]float64  `json:"scale,omitempty"`
	Fields      []Property   `json:"fields,omitempty"`
}

type ModelPatch struct {
	NodePatches map[string]NodePatch `json:"nodePatches,omitempty"`
}

type TaskSimplifyParams struct {
	SimplifyFactor float32 `arg:"--simplify-factor" json:"simplifyFactor,omitempty"`
	MergeMaterials bool    `arg:"--merge-materials" json:"mergeMaterials,omitempty"`
	MaxTextureSize uint32  `arg:"--maximum-texture-size" json:"maximumTextureSize,omitempty"`
	TextureScale   float32 `arg:"--texture-scale" json:"textureScale,omitempty"`
	UseWebP        bool    `arg:"--use-webp" json:"useWebP,omitempty"`
	Draco          bool    `arg:"--draco" json:"draco,omitempty"`
	DracoPosBits   int32   `arg:"--draco-pos-quantization-bits" json:"dracoPosBits,omitempty"`

	OnlyExteriorNodes bool `arg:"--only-keep-exterior-nodes" json:"onlyExteriorNodes,omitempty"`
}

type TaskType string

const (
	TaskSlice    TaskType = "slice"
	TaskSimplify TaskType = "simplify"
)

type TaskInfo struct {
	Type  TaskType
	ID    string `json:"id"`
	ReRun bool   `json:"rerun,omitempty"`
	Async bool   `json:"async,omitempty"`

	ModelPatch *ModelPatch `json:"modelPatch,omitempty"`

	SimplifyParams *TaskSimplifyParams `json:"simplifyParams,omitempty"`
}
