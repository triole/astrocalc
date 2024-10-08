# Astrocalc ![build](https://github.com/triole/astrocalc/actions/workflows/build.yaml/badge.svg) ![test](https://github.com/triole/astrocalc/actions/workflows/test.yaml/badge.svg)

<!-- toc -->

- [Synopsis](#synopsis)
- [Howto](#howto)
- [Http server use](#http-server-use)
- [Command line use](#command-line-use)
- [Data output example](#data-output-example)

<!-- /toc -->

## Synopsis

A command line tool to do simple astro calculations. Different output formats are supported.

## Howto

## Http server use

Run `astrocalc -s` to spin up server. Do stuff like...

```shell
curl http://localhost:8080?loc=london

curl http://localhost:8080?lat=51.5073509&lon=-0.1277583
```

## Command line use

Run for example `astrocalc london` or `astrocalc 51.507 -0.127`

## Data output example

```go mdox-exec="r london"
{
  "time": "2024-09-01T10:34:54.166370333+02:00",
  "location": {
    "lat": 51.5073509,
    "lon": -0.1277583,
    "name": "London"
  },
  "sun": {
    "light": {
      "dawn": "2024-09-01T04:40:05.100157952Z",
      "dusk": "2024-09-01T19:23:50.837085952Z",
      "golden_hour": "2024-09-01T18:04:21.184260352Z",
      "golden_hour_end": "2024-09-01T05:59:34.752983296Z",
      "nadir": "2024-09-01T00:01:57.968621824Z",
      "nautical_dawn": "2024-09-01T03:57:18.178896896Z",
      "nautical_dusk": "2024-09-01T20:06:37.758346752Z",
      "night": "2024-09-01T20:53:50.73799168Z",
      "night_end": "2024-09-01T03:10:05.199251968Z",
      "solar_noon": "2024-09-01T12:01:57.968621824Z",
      "sunrise": "2024-09-01T05:14:57.656500224Z",
      "sunrise_end": "2024-09-01T05:18:29.378023424Z",
      "sunset": "2024-09-01T18:48:58.280743424Z",
      "sunset_start": "2024-09-01T18:45:26.559220224Z"
    },
    "position": {
      "altitude": 0.5180177725094327,
      "azimuth": -1.100788368879874
    }
  },
  "moon": {
    "light": {
      "always_down": false,
      "always_up": false,
      "rise": "2024-09-01T04:00:00+02:00",
      "set": "2024-09-01T20:00:00+02:00"
    },
    "position": {
      "altitude": 0.81738638265308,
      "azimuth": -0.9305146132357468,
      "distance": 398068.5939872828,
      "parallactic_angle": -0.5540484171960416
    },
    "illumination": {
      "angle": 2.1114360879257674,
      "fraction": 0.026634652557527283,
      "phase": 0.9478180102947805
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
  -s, --server           run web server, http instead of command line use
  -p, --port=8080        web web server port
  -f, --format="json"    output format, can be: json, toml, yaml
  -V, --version-flag     display version
```
