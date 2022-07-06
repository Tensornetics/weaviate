package test

import (
	"testing"

	"github.com/semi-technologies/weaviate/entities/models"
)

func addTestDataL2(t *testing.T) {
	createObject(t, &models.Object{
		Class: "L2Squared_Class",
		Properties: map[string]interface{}{
			"name": "object_1",
		},
		Vector: []float32{
			10, 11, 12,
		},
	})

	createObject(t, &models.Object{
		Class: "L2Squared_Class",
		Properties: map[string]interface{}{
			"name": "object_2",
		},
		Vector: []float32{
			13, 15, 17,
		},
	})

	createObject(t, &models.Object{
		Class: "L2Squared_Class",
		Properties: map[string]interface{}{
			"name": "object_3",
		},
		Vector: []float32{
			0, 0, 0, // a zero vecto
		},
	})
}

func testL2(t *testing.T) {
	t.Run("without any limiting parameters", func(t *testing.T) {
		res := AssertGraphQL(t, nil, `
		{
			Get{
				L2Squared_Class(nearVector:{vector: [10,11,12]}){
					name 
					_additional{distance certainty}
				}
			}
		}
		`)
		results := res.Get("Get", "L2Squared_Class").AsSlice()
		expectedDistances := []float32{
			0,   // the same vector as the query
			50,  // distance to the second vector
			365, // l2 squared distance to the root
		}

		compareDistances(t, expectedDistances, results)
	})

	t.Run("with a certainty", func(t *testing.T) {
		// not supported for non-cosine distances
		ErrorGraphQL(t, nil, `
		{
			Get{
				L2Squared_Class(nearVector:{vector: [10,11,12], certainty:0.3}){
					name 
					_additional{distance certainty}
				}
			}
		}
		`)
	})

	t.Run("a high distance that includes all elements", func(t *testing.T) {
		res := AssertGraphQL(t, nil, `
		{
			Get{
				L2Squared_Class(nearVector:{vector: [10,11,12], distance: 365}){
					name 
					_additional{distance certainty}
				}
			}
		}
		`)
		results := res.Get("Get", "L2Squared_Class").AsSlice()
		expectedDistances := []float32{
			0,   // the same vector as the query
			50,  // distance to the second vector
			365, // l2 squared distance to the root
		}

		compareDistances(t, expectedDistances, results)
	})

	t.Run("a distance that is too low for the last element", func(t *testing.T) {
		res := AssertGraphQL(t, nil, `
		{
			Get{
				L2Squared_Class(nearVector:{vector: [10,11,12], distance: 364}){
					name 
					_additional{distance certainty}
				}
			}
		}
		`)
		results := res.Get("Get", "L2Squared_Class").AsSlice()
		expectedDistances := []float32{
			0,  // the same vector as the query
			50, // distance to the second vector
			// last eleme skipped, because 365 > 364
		}

		compareDistances(t, expectedDistances, results)
	})

	t.Run("a distance that is too low for the second element", func(t *testing.T) {
		res := AssertGraphQL(t, nil, `
		{
			Get{
				L2Squared_Class(nearVector:{vector: [10,11,12], distance: 49}){
					name 
					_additional{distance certainty}
				}
			}
		}
		`)
		results := res.Get("Get", "L2Squared_Class").AsSlice()
		expectedDistances := []float32{
			0, // the same vector as the query
			// second elem skipped, because 50 > 49
			// last eleme skipped, because 365 > 364
		}

		compareDistances(t, expectedDistances, results)
	})

	t.Run("a really low distance that only matches one elem", func(t *testing.T) {
		res := AssertGraphQL(t, nil, `
		{
			Get{
				L2Squared_Class(nearVector:{vector: [10,11,12], distance: 0.001}){
					name 
					_additional{distance certainty}
				}
			}
		}
		`)
		results := res.Get("Get", "L2Squared_Class").AsSlice()
		expectedDistances := []float32{
			0, // the same vector as the query
			// second elem skipped, because 50 > 0.001
			// last eleme skipped, because 365 > 0.001
		}

		compareDistances(t, expectedDistances, results)
	})

	t.Run("a distance of 0 only matches exact elements", func(t *testing.T) {
		res := AssertGraphQL(t, nil, `
		{
			Get{
				L2Squared_Class(nearVector:{vector: [10,11,12], distance: 0}){
					name 
					_additional{distance certainty}
				}
			}
		}
		`)
		results := res.Get("Get", "L2Squared_Class").AsSlice()
		expectedDistances := []float32{
			0, // the same vector as the query
			// second elem skipped, because 50 > 0
			// last eleme skipped, because 365 > 0
		}

		compareDistances(t, expectedDistances, results)
	})
}