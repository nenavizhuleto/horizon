package protocol_test

import (
	"encoding/json"
	"reflect"
	"testing"

	. "github.com/nenavizhuleto/horizon/protocol"
)

func Test_DetectionUnmarshalJSON(t *testing.T) {
	target := Detection{
		Type: ObjectDetection,
		Value: []Object{
			{
				Class: "person",
			},
			{
				Class: "person",
			},
		},
	}

	data, _ := json.Marshal(target)

	var detection Detection

	if err := json.Unmarshal(data, &detection); err != nil {
		t.Fatalf("unmarshal error: %+v", err)
	}

	if !reflect.DeepEqual(target, detection) {
		t.Fatalf("target != detection\n\nexpected: %+v\n\nactual: %+v", target, detection)
	}

}

func Test_DetectionMethods(t *testing.T) {
	detection := Detection{
		Type: MotionDetection,
		Value: []Motion{
			{},
			{},
			{},
			{},
		},
	}

	motions := detection.Motions()
	objects := detection.Objects()
	values := detection.Values()

	if len(motions) != 4 {
		t.Errorf("motions len 4 != %d", len(motions))
	}

	if len(objects) != 0 {
		t.Errorf("objects len 0 != %d", len(objects))
	}

	if len(values) != 0 {
		t.Errorf("values len 0 != %d", len(values))
	}

}
