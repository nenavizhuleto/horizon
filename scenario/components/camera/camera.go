package camera

import (
	"github.com/nenavizhuleto/horizon/scenario"
)

type Input string
type Output string
type Parameter string

// Outputs
const (
	Video = Output("video")
	Audio = Output("audio")
)

// Parameters
const (
	Source = Parameter("source")
)

type component struct {
	source string
}

func New() scenario.Component[Input, Output, Parameter] {
	return &component{}
}

func (c *component) Name() string {
	return "camera"
}

func (c *component) Inputs() []Input {
	return nil
}

func (c *component) SetInput(_ scenario.Channel)       {}
func (c *component) GetInput(_ Input) scenario.Channel { return nil }

func (c *component) Outputs() []Output {
	return []Output{
		Video,
		Audio,
	}
}

func (c *component) SetOutput(output scenario.Output) {
}
func (c *component) GetOutput(label Output) scenario.Output { return nil }

func (c *component) GetParameters() []Parameter {
	return []Parameter{
		Source,
	}
}

func (c *component) SetParameter(label Parameter, value any) {
	switch label {
	case Source:
		c.source = value.(string)
	}
}
func (c *component) GetParameter(label Parameter) any { return nil }

func Test() {
	camera := New()
	camera.SetParameter(Source, "rtsp://localhost:8554/live.stream")
	video := camera.GetOutput(Video)
}
