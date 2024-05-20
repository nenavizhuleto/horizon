
# Messages

## Message type: `detection.motion`

```go
{
  "producer": {
    "id": "PRODUCER_ID",
    "name": "PRODUCER_NAME",
    "group": "PRODUCER_GROUP",
    "modules": []
  },
  "type": "detection.motion",
  "body": {
    "timestamp": "2024-05-20T15:41:57.504626875+05:00",
    "detections": [
      {
        "source": {
          "uri": "http://localhost:1234/hello_world",
          "dimensions": {
            "width": 4123,
            "height": 1523
          }
        },
        "position": {
          "x": 123,
          "y": 321,
          "width": 400,
          "height": 600
        }
      }
    ]
  }
}
```

## Message type: `detection.object`

```go
{
  "producer": {
    "id": "PRODUCER_ID",
    "name": "PRODUCER_NAME",
    "group": "PRODUCER_GROUP",
    "modules": []
  },
  "type": "detection.object",
  "body": {
    "timestamp": "2024-05-20T15:41:57.504626875+05:00",
    "frame_location": {
      "partition": 124,
      "offset": 5123512,
      "topic": "TOPIC_NAME"
    },
    "detections": [
      {
        "id": "e422664a-d910-4a75-905c-1fdff28cd876",
        "class": "person",
        "bounding_box": {
          "x": 123,
          "y": 321,
          "width": 400,
          "height": 600
        },
        "confidence": 0.51
      },
      {
        "id": "ff74a087-5878-474c-9990-fb08b195a506",
        "class": "car",
        "bounding_box": {
          "x": 123,
          "y": 321,
          "width": 400,
          "height": 600
        },
        "confidence": 0.87
      }
    ]
  }
}
```

## Message type: `analysis.object`

```go
{
  "producer": {
    "id": "PRODUCER_ID",
    "name": "PRODUCER_NAME",
    "group": "PRODUCER_GROUP",
    "modules": []
  },
  "type": "analysis.object",
  "body": {
    "event_id": "EVENT_ID",
    "timestamp": "2024-05-20T15:41:57.504626875+05:00",
    "severity": "info",
    "frame_location": {
      "partition": 124,
      "offset": 5123512,
      "topic": "TOPIC_NAME"
    },
    "analyses": [
      {
        "id": "OBJECT_ANALYSIS_ID",
        "report": "OBJECT_ANALYSIS_REPORT",
        "subject": {
          "id": "e422664a-d910-4a75-905c-1fdff28cd876",
          "class": "person",
          "bounding_box": {
            "x": 123,
            "y": 321,
            "width": 400,
            "height": 600
          },
          "confidence": 0.51
        }
      }
    ]
  }
}
```

## Message type: `detection.plate`

```go
{
  "producer": {
    "id": "PRODUCER_ID",
    "name": "PRODUCER_NAME",
    "group": "PRODUCER_GROUP",
    "modules": []
  },
  "type": "detection.plate",
  "body": {
    "timestamp": "2024-05-20T15:41:57.504626875+05:00",
    "frame_location": {
      "partition": 124,
      "offset": 5123512,
      "topic": "TOPIC_NAME"
    },
    "detections": [
      {
        "bounding_box": {
          "x": 123,
          "y": 321,
          "width": 400,
          "height": 600
        },
        "confidence": 0.99,
        "plate": "EM322"
      },
      {
        "bounding_box": {
          "x": 123,
          "y": 321,
          "width": 400,
          "height": 600
        },
        "confidence": 0.64,
        "plate": "AD231"
      }
    ]
  }
}
```

## Message type: `analysis.plate`

```go
{
  "producer": {
    "id": "PRODUCER_ID",
    "name": "PRODUCER_NAME",
    "group": "PRODUCER_GROUP",
    "modules": []
  },
  "type": "analysis.plate",
  "body": {
    "event_id": "EVENT_ID",
    "timestamp": "2024-05-20T15:41:57.504626875+05:00",
    "severity": "warning",
    "frame_location": {
      "partition": 124,
      "offset": 5123512,
      "topic": "TOPIC_NAME"
    },
    "analyses": [
      {
        "id": "PLATE_ANALYSIS_ID",
        "report": [
          {
            "id": "LIST_ID",
            "name": "LIST_NAME",
            "severity": "warning",
            "color": "#00ff44",
            "vehicles": [
              {
                "id": "VEHICLE_ID",
                "plate": "EM322",
                "brand": "Hyundai",
                "model": "Solaris",
                "color": "white"
              }
            ]
          }
        ],
        "subject": {
          "bounding_box": {
            "x": 123,
            "y": 321,
            "width": 400,
            "height": 600
          },
          "confidence": 0.99,
          "plate": "EM322"
        }
      }
    ]
  }
}
```

## Message type: `frame`

```go
{
  "producer": {
    "id": "PRODUCER_ID",
    "name": "PRODUCER_NAME",
    "group": "PRODUCER_GROUP",
    "modules": []
  },
  "type": "frame",
  "body": {
    "timestamp": "2024-05-20T15:41:57.504626875+05:00",
    "dimensions": {
      "width": 4123,
      "height": 1523
    },
    "regions": [
      {
        "x": 10,
        "y": 20,
        "width": 400,
        "height": 300
      },
      {
        "x": 20,
        "y": 10,
        "width": 300,
        "height": 400
      }
    ],
    "data": "AAECAw=="
  }
}
```

## Message type: `event.start`

```go
{
  "producer": {
    "id": "PRODUCER_ID",
    "name": "PRODUCER_NAME",
    "group": "PRODUCER_GROUP",
    "modules": []
  },
  "type": "event.start",
  "body": {
    "id": "EVENT_ID",
    "start": "2024-05-20T15:41:57.504626875+05:00"
  }
}
```

## Message type: `event.end`

```go
{
  "producer": {
    "id": "PRODUCER_ID",
    "name": "PRODUCER_NAME",
    "group": "PRODUCER_GROUP",
    "modules": []
  },
  "type": "event.end",
  "body": {
    "id": "EVENT_ID",
    "end": "2024-05-20T15:41:57.504626875+05:00",
    "duration": 10000000000
  }
}
```

## Message type: `media`

```go
{
  "producer": {
    "id": "PRODUCER_ID",
    "name": "PRODUCER_NAME",
    "group": "PRODUCER_GROUP",
    "modules": []
  },
  "type": "media",
  "body": {
    "event_id": "EVENT_ID",
    "recording": {
      "file": "recording.mp4",
      "url": "http://localhost:8912/recording.mp4"
    },
    "media": [
      {
        "analysis_id": "ANALYSIS_ID",
        "timestamp": "2024-05-20T15:41:57.504626875+05:00",
        "file": "object.jpeg",
        "url": "http://localhost:8912/object.jpeg"
      }
    ]
  }
}
```

# Entities

## Entity type: `protocol.Message[interface {}]`

```go
{
  "producer": {
    "id": "PRODUCER_ID",
    "name": "PRODUCER_NAME",
    "group": "PRODUCER_GROUP",
    "modules": []
  },
  "type": "detection.object",
  "body": "BODY"
}
```

## Entity type: `[]protocol.Position`

```go
[
  {
    "x": 10,
    "y": 20,
    "width": 400,
    "height": 300
  },
  {
    "x": 20,
    "y": 10,
    "width": 300,
    "height": 400
  }
]
```

## Entity type: `protocol.FrameLocation`

```go
{
  "partition": 124,
  "offset": 5123512,
  "topic": "TOPIC_NAME"
}
```

## Entity type: `protocol.Position`

```go
{
  "x": 123,
  "y": 321,
  "width": 400,
  "height": 600
}
```

## Entity type: `protocol.Frame`

```go
"AAECAw=="
```

## Entity type: `protocol.Dimensions`

```go
{
  "width": 4123,
  "height": 1523
}
```

## Entity type: `protocol.Recording`

```go
{
  "file": "recording.mp4",
  "url": "http://localhost:8912/recording.mp4"
}
```

## Entity type: `[]protocol.Media`

```go
[
  {
    "analysis_id": "ANALYSIS_ID",
    "timestamp": "2024-05-20T15:41:57.504626875+05:00",
    "file": "object.jpeg",
    "url": "http://localhost:8912/object.jpeg"
  }
]
```

## Entity type: `protocol.Event`

```go
{
  "id": "EVENT_ID",
  "start": "2024-05-20T15:41:57.504626875+05:00",
  "end": "2024-05-20T15:41:57.504626875+05:00",
  "duration": 10000000000
}
```

## Entity type: `[]protocol.MotionDetection`

```go
[
  {
    "source": {
      "uri": "http://localhost:1234/hello_world",
      "dimensions": {
        "width": 4123,
        "height": 1523
      }
    },
    "position": {
      "x": 123,
      "y": 321,
      "width": 400,
      "height": 600
    }
  }
]
```

## Entity type: `[]protocol.ObjectDetection`

```go
[
  {
    "id": "e422664a-d910-4a75-905c-1fdff28cd876",
    "class": "person",
    "bounding_box": {
      "x": 123,
      "y": 321,
      "width": 400,
      "height": 600
    },
    "confidence": 0.51
  },
  {
    "id": "ff74a087-5878-474c-9990-fb08b195a506",
    "class": "car",
    "bounding_box": {
      "x": 123,
      "y": 321,
      "width": 400,
      "height": 600
    },
    "confidence": 0.87
  }
]
```

## Entity type: `[]protocol.PlateDetection`

```go
[
  {
    "bounding_box": {
      "x": 123,
      "y": 321,
      "width": 400,
      "height": 600
    },
    "confidence": 0.99,
    "plate": "EM322"
  },
  {
    "bounding_box": {
      "x": 123,
      "y": 321,
      "width": 400,
      "height": 600
    },
    "confidence": 0.64,
    "plate": "AD231"
  }
]
```

## Entity type: `[]protocol.Vehicle`

```go
[
  {
    "id": "VEHICLE_ID",
    "plate": "EM322",
    "brand": "Hyundai",
    "model": "Solaris",
    "color": "white"
  }
]
```
