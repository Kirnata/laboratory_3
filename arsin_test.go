package main

import "testing"

func TestSuperAsin(t *testing.T) {

	t.Run("zero", func(t *testing.T) {
		var x, result = 0.0, 0.0

		realResult := superAsin(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("positive test 1", func(t *testing.T) {
		var x, result = 1.0, -1.5707963267948966

		realResult := superAsin(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("positive test 2", func(t *testing.T) {
		var x, result = 0.8, -0.9272952180016123

		realResult := superAsin(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("positive test 3", func(t *testing.T) {
		var x, result = 0.3, -0.3046926540153975

		realResult := superAsin(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("positive test 4", func(t *testing.T) {
		var x, result = 0.1, -0.1001674211615598

		realResult := superAsin(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("negative test 1", func(t *testing.T) {
		var x, result = -1.0, 1.5707963267948966

		realResult := superAsin(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("negative test 2", func(t *testing.T) {
		var x, result = -0.7, 0.775397496610753

		realResult := superAsin(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("negative test 3", func(t *testing.T) {
		var x, result = -0.4, 0.41151684606748806

		realResult := superAsin(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("negative test 4", func(t *testing.T) {
		var x, result = -0.3, 0.3046926540153975

		realResult := superAsin(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("negative test 5", func(t *testing.T) {
		var x, result = -0.05, 0.050020856805770016

		realResult := superAsin(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})
}
