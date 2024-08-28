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
  "time": "2024-08-29T00:09:44.36448346+02:00",
  "location": {
    "lat": 51.5073509,
    "lon": -0.1277583,
    "name": "London"
  },
  "sun": {
    "light": {
      "dawn": "2024-08-28T04:33:12.426516224Z",
      "dusk": "2024-08-28T19:33:11.403226624Z",
      "golden_hour": "2024-08-28T18:12:49.817309184Z",
      "golden_hour_end": "2024-08-28T05:53:34.012433664Z",
      "nadir": "2024-08-28T00:03:11.914871296Z",
      "nautical_dawn": "2024-08-28T03:49:25.702360832Z",
      "nautical_dusk": "2024-08-28T20:16:58.12738176Z",
      "night": "2024-08-28T21:06:02.825229824Z",
      "night_end": "2024-08-28T03:00:21.004513024Z",
      "solar_noon": "2024-08-28T12:03:11.914871296Z",
      "sunrise": "2024-08-28T05:08:35.498203136Z",
      "sunrise_end": "2024-08-28T05:12:09.582095104Z",
      "sunset": "2024-08-28T18:57:48.331539456Z",
      "sunset_start": "2024-08-28T18:54:14.247647488Z"
    },
    "position": {
      "altitude": -0.4263533541260498,
      "azimuth": 2.6069105817646965
    }
  },
  "moon": {
    "light": {
      "always_down": false,
      "always_up": false,
      "rise": "2024-08-29T01:00:00+02:00",
      "set": "2024-08-29T19:00:00+02:00"
    },
    "position": {
      "altitude": -0.08597925031873344,
      "azimuth": -2.657834491697706,
      "distance": 382751.27894795325,
      "parallactic_angle": -0.33515996911638857
    },
    "illumination": {
      "angle": 1.6775088410544998,
      "fraction": 0.24364853048063495,
      "phase": 0.8356778874020945
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
