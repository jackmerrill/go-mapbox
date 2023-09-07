# go-mapbox

Mapbox API wrappers for Golang

[![Documentation](https://img.shields.io/badge/docs-godoc-blue.svg)](https://godoc.org/github.com/jackmerrill/go-mapbox/lib)
[![GitHub tag](https://img.shields.io/github/tag/jackmerrill/go-mapbox.svg)](https://github.com/jackmerrill/go-mapbox)
[![Build Status](https://travis-ci.org/jackmerrill/go-mapbox.svg?branch=master)](https://travis-ci.org/jackmerrill/go-mapbox)

See [here](https://golanglibs.com/top?q=mapbox) for other golang/mapbox projects.

## Status

Very early WIP, pull requests and issues are most welcome. See [lib/geocode/](lib/geocode) or [lib/directions/](lib/directions) for an example module to mimic.

Because Travis-CI does not expose the build environment to untrusted branches (ie. Pull Requests) tests have to be manually prompted by a repository admin then force merged. Don't panic when your local tests pass but travis fails with "Mapbox API token not found", we will manually run them as soon as possible. See issue [#10](https://github.com/jackmerrill/go-mapbox/issues/10) for more information.

### Modules

- [x] Geocoding
- [x] Directions
- [x] Directions Matrix
- [x] Map Matching
- [ ] Styles
- [x] Maps
- [ ] Static
- [ ] Datasets

## Examples

### Initialisation

```go
// Import the core module (and any required APIs)
import (
    "gopkg.in/jackmerrill/go-mapbox.v0/lib"
    "gopkg.in/jackmerrill/go-mapbox.v0/lib/base"
)

// Fetch token from somewhere
token := os.Getenv("MAPBOX_TOKEN")

// Create new mapbox instance
mapBox := mapbox.NewMapbox(token)

```

### Map API

```go
import (
    "gopkg.in/jackmerrill/go-mapbox.v0/lib/maps"
)

img, err := mapBox.Maps.GetTiles(maps.MapIDSatellite, 1, 0, 0, maps.MapFormatJpg90, true)
```

### Geocoding

```go
import (
    "gopkg.in/jackmerrill/go-mapbox.v0/lib/geocode"
)

// Forward Geocoding
var forwardOpts geocode.ForwardRequestOpts
forwardOpts.Limit = 1

place := "2 lincoln memorial circle nw"

forward, err := mapBox.Geocode.Forward(place, &forwardOpts)


// Reverse Geocoding
var reverseOpts geocode.ReverseRequestOpts
reverseOpts.Limit = 1

loc := &base.Location{72.438939, 34.074122}

reverse, err := mapBox.Geocode.Reverse(loc, &reverseOpts)

```

### Directions

```go
import (
    "gopkg.in/jackmerrill/go-mapbox.v0/lib/directions"
)

var directionOpts directions.RequestOpts

locs := []base.Location{{-122.42, 37.78}, {-77.03, 38.91}}

directions, err := mapBox.Directions.GetDirections(locs, directions.RoutingCycling, &directionOpts)

```

## Layout

- [lib/base](lib/base/) contains a common base for API modules
- [lib/maps](lib/maps/) contains the maps API module
- [lib/directions](lib/directions/) contains the directions API module
- [lib/geocode](lib/geocode/) contains the geocoding API module

---

If you have any questions, comments, or suggestions, feel free to open an issue or a pull request.
