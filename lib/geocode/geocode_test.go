/**
 * go-mapbox Geocoding Module Tests
 * Wraps the mapbox geocoding API for server side use
 * See https://www.mapbox.com/api-documentation/#geocoding for API information
 *
 * https://github.com/jackmerrill/go-mapbox
 * Copyright 2017 Ryan Kurte
 */

package geocode

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/jackmerrill/go-mapbox/lib/base"
)

func TestGeocoder(t *testing.T) {

	b, err := base.NewBase(os.Getenv("MAPBOX_TOKEN"))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	geocode := NewGeocode(b)

	t.Run("Can geocode", func(t *testing.T) {
		var reqOpt ForwardRequestOpts
		reqOpt.Limit = 1

		place := "2 lincoln memorial circle nw"

		res, err := geocode.Forward(place, &reqOpt)
		if err != nil {
			t.Error(err)
		}

		if res.Type != "FeatureCollection" {
			t.Errorf("Invalid response type: %s", res.Type)
		}

		if !reflect.DeepEqual(res.Query, strings.Split(place, " ")) {
			t.Errorf("Invalid query response: %s", res.Query)
		}

	})

	t.Run("Can reverse geocode", func(t *testing.T) {
		var reqOpt ReverseRequestOpts
		reqOpt.Limit = 1

		loc := &base.Location{72.438939, 34.074122}

		res, err := geocode.Reverse(loc, &reqOpt)
		if err != nil {
			t.Error(err)
		}

		if res.Type != "FeatureCollection" {
			t.Errorf("Invalid response type: %s", res.Type)
		}

	})

}
