## Motion Detection Message

```json
    {
          "type": "detection.motion",
          "producer": {
            "camera": {
              "name": "fuck",
              "group": "you"
            },
            "id": "id",
            "name": "producer"
          },
          "timestamp": "2024-04-25T17:12:44.458118064+05:00",
          "motions": [
            {
              "source": {
                "uri": "",
                "dimensions": {
                  "width": 0,
                  "height": 0
                }
              },
              "position": {
                "x": 0,
                "y": 0,
                "width": 0,
                "height": 0
              }
            },
            {
              "source": {
                "uri": "",
                "dimensions": {
                  "width": 0,
                  "height": 0
                }
              },
              "position": {
                "x": 0,
                "y": 0,
                "width": 0,
                "height": 0
              }
            }
          ]
        }
```

## Object Detection Message

```json
    {
          "frame_location": {
            "partition": 123,
            "offset": 332,
            "topic": ""
          },
          "type": "detection.object",
          "producer": {
            "id": "id",
            "name": "producer"
          },
          "timestamp": "2024-04-25T17:12:44.458279588+05:00",
          "objects": [
            {
              "class": "person",
              "bounding_box": {
                "x": 0,
                "y": 0,
                "width": 0,
                "height": 0
              },
              "confidence": 0,
              "user_data": null
            },
            {
              "class": "weapon",
              "bounding_box": {
                "x": 0,
                "y": 0,
                "width": 0,
                "height": 0
              },
              "confidence": 0,
              "user_data": null
            },
            {
              "class": "plate",
              "bounding_box": {
                "x": 0,
                "y": 0,
                "width": 0,
                "height": 0
              },
              "confidence": 0,
              "user_data": null
            }
          ]
        }
```

## Value Detection Message

```json
    {
          "type": "detection.value",
          "producer": {
            "id": "id",
            "name": "producer"
          },
          "timestamp": "2024-04-25T17:12:44.458355336+05:00",
          "values": [
            123,
            555,
            "hello",
            "omh",
            true,
            {}
          ]
        }
```

## Frame Message

```json
    {
          "dimensions": {
            "width": 100,
            "height": 2333
          },
          "regions": [
            {
              "x": 0,
              "y": 0,
              "width": 100,
              "height": 200
            }
          ],
          "type": "frame",
          "producer": {
            "id": "id",
            "name": "producer"
          },
          "frame": "AAECBAUGBwg="
        }
```

## Motion Detection Analysis

```json
    {
          "type": "analysis.detection.motion",
          "producer": {
            "camera": {
              "name": "fuck",
              "group": "you"
            },
            "id": "id",
            "name": "producer"
          },
          "severity": "warning",
          "report": {
            "producer": {
              "id": "new_id",
              "name": "new_producer"
            },
            "timestamp": "2024-04-25T17:12:44.45845688+05:00",
            "user_data": null,
            "motion": {
              "source": {
                "uri": "helloworld",
                "dimensions": {
                  "width": 0,
                  "height": 0
                }
              },
              "position": {
                "x": 0,
                "y": 0,
                "width": 0,
                "height": 0
              }
            }
          }
        }
```

## Object Detection Analysis

```json
    {
          "type": "analysis.detection.object",
          "producer": {
            "id": "id",
            "name": "producer"
          },
          "severity": "panic",
          "report": {
            "frame_location": {
              "partition": 1234,
              "offset": 1234,
              "topic": "topic"
            },
            "producer": {
              "id": "",
              "name": ""
            },
            "timestamp": "2024-04-25T17:12:44.458507604+05:00",
            "user_data": {
              "direction": "forward",
              "name": "john"
            },
            "object": {
              "class": "person",
              "bounding_box": {
                "x": 0,
                "y": 0,
                "width": 0,
                "height": 0
              },
              "confidence": 0,
              "user_data": null
            }
          }
        }
```

## Value Detection Analysis

```json
    {
          "type": "analysis.detection.value",
          "producer": {
            "id": "id",
            "name": "producer"
          },
          "severity": "info",
          "report": {
            "producer": {
              "id": "",
              "name": ""
            },
            "timestamp": "2024-04-25T17:12:44.45855635+05:00",
            "user_data": null,
            "value": 1259812
          }
        }
```

## Event Analysis

```json
    {
          "type": "analysis.event",
          "producer": {
            "id": "id",
            "name": "producer"
          },
          "severity": "panic",
          "report": {
            "producer": {
              "id": "",
              "name": ""
            },
            "timestamp": "2024-04-25T17:12:44.458604651+05:00",
            "user_data": null,
            "event": {
              "id": "event_id",
              "trigger": null,
              "start": "0001-01-01T00:00:00Z",
              "end": "0001-01-01T00:00:00Z",
              "duration": 0
            },
            "detections": [
              {
                "source": {
                  "uri": "",
                  "dimensions": {
                    "width": 0,
                    "height": 0
                  }
                },
                "position": {
                  "x": 0,
                  "y": 0,
                  "width": 0,
                  "height": 0
                }
              },
              {
                "class": "",
                "bounding_box": {
                  "x": 0,
                  "y": 0,
                  "width": 0,
                  "height": 0
                },
                "confidence": 0,
                "user_data": null
              },
              123,
              {
                "source": {
                  "uri": "",
                  "dimensions": {
                    "width": 0,
                    "height": 0
                  }
                },
                "position": {
                  "x": 0,
                  "y": 0,
                  "width": 0,
                  "height": 0
                }
              }
            ]
          }
        }
```
