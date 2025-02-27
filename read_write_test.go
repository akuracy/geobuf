package geobuf

import (
	"fmt"
	"github.com/paulmach/go.geojson"
	"io/ioutil"
	"math"
	"os"
	"testing"
)

func DeltaPt(pt []float64, testpt []float64) float64 {
	deltax := math.Abs(pt[0] - testpt[0])
	deltay := math.Abs(pt[1] - testpt[1])
	return deltax + deltay
}

var PrecisionError = math.Pow(10.0, -6.0)

// this function tests whether a all geojson features are maintained on a read write
func TestReadWriteFile(t *testing.T) {
	// reading a geojson
	bytevals, err := ioutil.ReadFile("test_data/county.geojson")
	if err != nil {
		fmt.Println(err)
	}

	// reading geojson
	fc, err := geojson.UnmarshalFeatureCollection(bytevals)
	if err != nil {
		fmt.Println(err)
	}

	featuremap := map[int]*geojson.Feature{}
	for _, feature := range fc.Features {
		featuremap[int(feature.ID.(float64))] = feature
	}

	// gettting the size of the feature collection
	size_fc := len(fc.Features)

	// creating a geobuf file object
	buf := WriterFileNew("test.geobuf")

	// adding each feature
	for _, feature := range fc.Features {
		buf.WriteFeature(feature)
	}

	// converting buffer to reader
	readbuf := buf.Reader()

	// reading each feature
	size_buf := 0
	for readbuf.Next() {
		feature := readbuf.Feature()
		id := int(feature.ID.(int))
		testfeature := featuremap[id]
		if testfeature.Geometry.Type == feature.Geometry.Type {
			for i := range testfeature.Geometry.Polygon {
				testring := testfeature.Geometry.Polygon[i]
				ring := feature.Geometry.Polygon[i]
				if len(ring) != len(testring) {
					t.Errorf("Different ring sizes expected %d got %d", len(testring), len(ring))
				} else {
					for j := range ring {
						pt := ring[j]
						testpt := testring[j]
						deltapt := DeltaPt(pt, testpt)
						if PrecisionError < deltapt {
							t.Errorf("Different Points expected %v %v", testpt, pt)
						}

					}

				}
			}
		} else {
			t.Errorf("Different Types")
		}

		size_buf++
	}

	// chekcing to see if size is the same as feature collection
	if size_buf != size_fc {
		t.Errorf("Error ReadWrite File %d fc %d buf", size_fc, size_buf)
	}
	os.Remove("test.geobuf")
}

// this function tests whether a all geojson features are maintained on a read write
func TestReadWriteBuf(t *testing.T) {
	// reading a geojson
	bytevals, err := ioutil.ReadFile("test_data/county.geojson")
	if err != nil {
		fmt.Println(err)
	}

	// reading geojson
	fc, err := geojson.UnmarshalFeatureCollection(bytevals)
	if err != nil {
		fmt.Println(err)
	}

	featuremap := map[int]*geojson.Feature{}
	for _, feature := range fc.Features {
		featuremap[int(feature.ID.(float64))] = feature
	}

	// gettting the size of the feature collection
	size_fc := len(fc.Features)

	// creating a geobuf file object
	buf := WriterBufNew()

	// adding each feature
	for _, feature := range fc.Features {
		buf.WriteFeature(feature)
	}

	// converting buffer to reader
	readbuf := buf.Reader()

	// reading each feature
	size_buf := 0
	for readbuf.Next() {
		feature := readbuf.Feature()
		id := int(feature.ID.(int))
		testfeature := featuremap[id]
		if testfeature.Geometry.Type == feature.Geometry.Type {
			for i := range testfeature.Geometry.Polygon {
				testring := testfeature.Geometry.Polygon[i]
				ring := feature.Geometry.Polygon[i]
				if len(ring) != len(testring) {
					t.Errorf("Different ring sizes expected %d got %d", len(testring), len(ring))
				} else {
					for j := range ring {
						pt := ring[j]
						testpt := testring[j]
						deltapt := DeltaPt(pt, testpt)
						if PrecisionError < deltapt {
							t.Errorf("Different Points expected %v %v", testpt, pt)
						}

					}

				}
			}
		} else {
			t.Errorf("Different Types")
		}

		size_buf++
	}

	// chekcing to see if size is the same as feature collection
	if size_buf != size_fc {
		t.Errorf("Error ReadWrite File %d fc %d buf", size_fc, size_buf)
	}
}

// given the size of the fc creates indicies for the feature collection
func CreateInds(size int) [][2]int {
	// creating intial indicies
	delta := size / 10
	current := 0
	newlist := []int{current}
	for current < size {
		current += delta
		newlist = append(newlist, current)
	}

	// fixing end
	if newlist[len(newlist)-1] >= size {
		newlist[len(newlist)-1] = size
	}

	// stepping through creaitng the two indicies
	oldi := newlist[0]
	totalinds := [][2]int{}
	for _, i := range newlist[1:] {
		totalinds = append(totalinds, [2]int{oldi, i})
		oldi = i
	}
	return totalinds
}

// this test splits a feature collection into 10 geobufs
// each geobuf is then appended to each other
// and then tested to see if all features are maintained
func TestReadWriteMultiBufFile(t *testing.T) {
	// reading a geojson
	bytevals, err := ioutil.ReadFile("test_data/county.geojson")
	if err != nil {
		fmt.Println(err)
	}

	// reading geojson
	fc, err := geojson.UnmarshalFeatureCollection(bytevals)
	if err != nil {
		fmt.Println(err)
	}

	featuremap := map[int]*geojson.Feature{}
	for _, feature := range fc.Features {
		featuremap[int(feature.ID.(float64))] = feature
	}

	// gettting the size of the feature collection
	size_fc := len(fc.Features)

	// getting the indicies for each buffer split
	inds := CreateInds(size_fc)

	// creating each buffer
	buffers := []*Writer{}
	total_split := 0
	for _, ind := range inds {
		// getting features from indicies
		i1, i2 := ind[0], ind[1]
		features := fc.Features[i1:i2]
		total_split += len(features)

		// creating a geobuf file object
		buf := WriterBufNew()

		// adding each feature
		for _, feature := range features {
			buf.WriteFeature(feature)
		}
		buffers = append(buffers, buf)
	}

	// checking to see if the total split is the same as fc
	if total_split != size_fc {
		t.Errorf("Split function not workign %d %d", size_fc, total_split)
	}

	// adding the geobufs into one big buffer
	bigbuffer := WriterFileNew("test.geobuf")

	// iterating through each buffer
	for _, buf := range buffers {
		bigbuffer.AddGeobuf(buf)
	}

	// converting buffer to reader
	bigbufreader := bigbuffer.Reader()

	// reading each feature
	// reading each feature
	size_buf := 0
	for bigbufreader.Next() {
		feature := bigbufreader.Feature()
		id := int(feature.ID.(int))
		testfeature := featuremap[id]
		if testfeature.Geometry.Type == feature.Geometry.Type {
			for i := range testfeature.Geometry.Polygon {
				testring := testfeature.Geometry.Polygon[i]
				ring := feature.Geometry.Polygon[i]
				if len(ring) != len(testring) {
					t.Errorf("Different ring sizes expected %d got %d", len(testring), len(ring))
				} else {
					for j := range ring {
						pt := ring[j]
						testpt := testring[j]
						deltapt := DeltaPt(pt, testpt)
						if PrecisionError < deltapt {
							t.Errorf("Different Points expected %v %v", testpt, pt)
						}

					}

				}
			}
		} else {
			t.Errorf("Different Types")
		}

		size_buf++
	}

	// chekcing to see if size is the same as feature collection
	if size_buf != size_fc {
		t.Errorf("Error ReadWrite File %d fc %d buf", size_fc, size_buf)
	}
	os.Remove("test.geobuf")
}

// this test splits a feature collection into 10 geobufs
// each geobuf is then appended to each other
// and then tested to see if all features are maintained
func TestReadWriteMultiBuf(t *testing.T) {
	// reading a geojson
	bytevals, err := ioutil.ReadFile("test_data/county.geojson")
	if err != nil {
		fmt.Println(err)
	}

	// reading geojson
	fc, err := geojson.UnmarshalFeatureCollection(bytevals)
	if err != nil {
		fmt.Println(err)
	}

	featuremap := map[int]*geojson.Feature{}
	for _, feature := range fc.Features {
		featuremap[int(feature.ID.(float64))] = feature
	}

	// gettting the size of the feature collection
	size_fc := len(fc.Features)

	// getting the indicies for each buffer split
	inds := CreateInds(size_fc)

	// creating each buffer
	buffers := []*Writer{}
	total_split := 0
	for _, ind := range inds {
		// getting features from indicies
		i1, i2 := ind[0], ind[1]
		features := fc.Features[i1:i2]
		total_split += len(features)

		// creating a geobuf file object
		buf := WriterBufNew()

		// adding each feature
		for _, feature := range features {
			buf.WriteFeature(feature)
		}
		buffers = append(buffers, buf)
	}

	// checking to see if the total split is the same as fc
	if total_split != size_fc {
		t.Errorf("Split function not workign %d %d", size_fc, total_split)
	}

	// adding the geobufs into one big buffer
	bigbuffer := WriterBufNew()

	// iterating through each buffer
	for _, buf := range buffers {
		bigbuffer.AddGeobuf(buf)
	}

	// converting buffer to reader
	bigbufreader := bigbuffer.Reader()

	// reading each feature
	size_buf := 0
	for bigbufreader.Next() {
		feature := bigbufreader.Feature()
		id := int(feature.ID.(int))
		testfeature := featuremap[id]
		if testfeature.Geometry.Type == feature.Geometry.Type {
			for i := range testfeature.Geometry.Polygon {
				testring := testfeature.Geometry.Polygon[i]
				ring := feature.Geometry.Polygon[i]
				if len(ring) != len(testring) {
					t.Errorf("Different ring sizes expected %d got %d", len(testring), len(ring))
				} else {
					for j := range ring {
						pt := ring[j]
						testpt := testring[j]
						deltapt := DeltaPt(pt, testpt)
						if PrecisionError < deltapt {
							t.Errorf("Different Points expected %v %v", testpt, pt)
						}

					}

				}
			}
		} else {
			t.Errorf("Different Types")
		}

		size_buf++
	}

	// chekcing to see if size is the same as feature collection
	if size_buf != size_fc {
		t.Errorf("Error ReadWrite File %d fc %d buf", size_fc, size_buf)
	}
}
