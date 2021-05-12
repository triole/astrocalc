# Astrocalc

<!--- mdtoc: toc begin -->

1.	[Synopsis](#synopsis)
2.	[Howto](#howto)<!--- mdtoc: toc end -->

## Synopsis

A command line tool to do simple astro calculations. Results are printed in a flat toml structure so that they can be parsed easily with [stoml](https://github.com/freshautomations/stoml).

## Howto

Run for example `astrocalc london` (or any other capital) and get the following:

```
Time = 2021-05-11T21:46:04Z

[Location]
  lat = 51.5073509
  lon = -0.1277583
  name = "London"

[Sun]
  [Sun.Light]
    dawn = 2021-05-11T03:35:32Z
    dusk = 2021-05-11T20:20:53Z
    goldenHour = 2021-05-11T18:51:26Z
    goldenHourEnd = 2021-05-11T05:05:00Z
    nadir = 2021-05-10T23:58:13Z
    nauticalDawn = 2021-05-11T02:41:11Z
    nauticalDusk = 2021-05-11T21:15:15Z
    night = 2021-05-11T22:30:57Z
    nightEnd = 2021-05-11T01:25:29Z
    solarNoon = 2021-05-11T11:58:13Z
    sunrise = 2021-05-11T04:16:03Z
    sunriseEnd = 2021-05-11T04:20:02Z
    sunset = 2021-05-11T19:40:23Z
    sunsetStart = 2021-05-11T19:36:24Z
  [Sun.Position]
    altitude = -0.25812499509045206
    azimuth = 2.5793811439637335

[Moon]
  [Moon.Light]
    alwaysDown = false
    alwaysUp = false
    rise = 2021-05-11T04:00:00Z
    set = 2021-05-11T19:00:00Z
  [Moon.Position]
    altitude = -0.2564819897797582
    azimuth = 2.53724845080525
    distance = 405905.7306171807
    parallacticAngle = 0.3784449863649929
  [Moon.Illumination]
    angle = -1.0321882935238267
    fraction = 0.00042794730810746806
    phase = 0.006585310508201714
```
