# Astrocalc ![build](https://github.com/triole/astrocalc/actions/workflows/build.yaml/badge.svg) ![test](https://github.com/triole/astrocalc/actions/workflows/test.yaml/badge.svg)

<!-- toc -->

- [Synopsis](#synopsis)
- [Howto](#howto)

<!-- /toc -->

## Synopsis

A command line tool to do simple astro calculations. Different output formats are supported.

## Howto

Run for example `astrocalc london` (or any other capital) and get the following:

```go mdox-exec="r london"
{
  "time": "2024-08-29T00:03:06.87743037+02:00",
  "location": {
    "lat": 51.5073509,
    "lon": -0.1277583,
    "name": "London"
  },
  "sun": {
    "light": {
      "dawn": "2024-08-28T04:33:12.426516224Z",
      "dusk": "2024-08-28T19:33:11.403226624Z",
      "goldenHour": "2024-08-28T18:12:49.817309184Z",
      "goldenHourEnd": "2024-08-28T05:53:34.012433664Z",
      "nadir": "2024-08-28T00:03:11.914871296Z",
      "nauticalDawn": "2024-08-28T03:49:25.702360832Z",
      "nauticalDusk": "2024-08-28T20:16:58.12738176Z",
      "night": "2024-08-28T21:06:02.825229824Z",
      "nightEnd": "2024-08-28T03:00:21.004513024Z",
      "solarNoon": "2024-08-28T12:03:11.914871296Z",
      "sunrise": "2024-08-28T05:08:35.498203136Z",
      "sunriseEnd": "2024-08-28T05:12:09.582095104Z",
      "sunset": "2024-08-28T18:57:48.331539456Z",
      "sunsetStart": "2024-08-28T18:54:14.247647488Z"
    },
    "position": {
      "altitude": -0.41692918810510243,
      "azimuth": 2.577407411869789
    }
  },
  "moon": {
    "light": {
      "alwaysDown": false,
      "alwaysUp": false,
      "rise": "2024-08-29T01:00:00+02:00",
      "set": "2024-08-29T19:00:00+02:00"
    },
    "position": {
      "altitude": -0.09379953186346796,
      "azimuth": -2.6811070547214033,
      "distance": 382729.4771832017,
      "parallacticAngle": -0.31968796547776823
    },
    "illumination": {
      "angle": 1.6769018312367283,
      "fraction": 0.24407281063319503,
      "phase": 0.835520633623858
    }
  }
}
```

## Help

```go mdox-exec="r -h"
Usage: astrocalc [<location> ...] [flags]

simple astronomical calculation

Arguments:
  [<location> ...]    location to use

Flags:
  -h, --help             Show context-sensitive help.
  -f, --format="json"    output format, can be: json, toml, yaml
  -V, --version-flag     display version
```
