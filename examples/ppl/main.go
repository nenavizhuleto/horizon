package main

import (
	"context"
	"log"
	"time"

	"github.com/nenavizhuleto/horizon/ppl"
	"github.com/nenavizhuleto/horizon/ppl/ops"
	"github.com/nenavizhuleto/horizon/producer"
	"github.com/nenavizhuleto/horizon/protocol"
)

func main() {
	object_pdc := producer.
		NewObjectDetectionProducer("", "")

	person_msg := object_pdc.NewDetectionMessage(
		time.Now(),
		protocol.Object{
			Class:      "person",
			Confidence: 0.98,
			UserData:   "john",
		},
	)

	plate_msg := object_pdc.NewDetectionMessage(
		time.Now(),
		protocol.Object{
			Class:      "plate",
			Confidence: 0.54,
			UserData:   "EV322",
		},
	)

	value_msg := producer.NewValueDetectionProducer("", "").NewDetectionMessage(time.Now(), "fuck", "you")
	motion_msg := producer.NewMotionDetectionProducer("", "").NewDetectionMessage(time.Now(), protocol.Motion{})

	ctx := ppl.NewContext(context.Background())

	p := ppl.NewPipeline(
		ops.DetectionFilter().
			Motions(
				ppl.NewElm(func(ctx *ppl.Context, message any) {
					log.Println("motion element")
				}),
			).
			Objects(
				ops.ObjectFilter("plate", func(ctx *ppl.Context, object protocol.Object) {
					log.Println("this is a plate", object)
				}),
				ops.ObjectFilter("person", func(ctx *ppl.Context, object protocol.Object) {
					log.Println("this is a person", object)
				}),
				ops.ObjectElm(func(ctx *ppl.Context, object protocol.Object) {
					log.Println("here goes every object...")
				}),
			).
			Values(
				ppl.Inline(func(ctx *ppl.Context, message any) bool {
					log.Println("this is a value", message)
					return true
				}),
			),
	)

	p.Message(ctx, person_msg)
	p.Message(ctx, plate_msg)
	p.Message(ctx, value_msg)
	p.Message(ctx, motion_msg)
}
