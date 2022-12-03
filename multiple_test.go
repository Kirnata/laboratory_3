package main

import "testing"

func TestSuperMultiple(t *testing.T) {

	t.Run("zero", func(t *testing.T) {
		var x, result = 0.0, 0.0

		realResult := superMultiple(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("positive test 1", func(t *testing.T) {
		var x, result = 1.0, 13.0

		realResult := superMultiple(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("positive test 2", func(t *testing.T) {
		var x, result = 4.0, 52.0

		realResult := superMultiple(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("positive test 3", func(t *testing.T) {
		var x, result = 50.0, 650.0

		realResult := superMultiple(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("positive test 4", func(t *testing.T) {
		var x, result = 100.0, 1300.0

		realResult := superMultiple(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("negative test 1", func(t *testing.T) {
		var x, result = -1.0, -13.0

		realResult := superMultiple(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("negative test 2", func(t *testing.T) {
		var x, result = -4.0, -52.0

		realResult := superMultiple(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("negative test 3", func(t *testing.T) {
		var x, result = -50.0, -650.0

		realResult := superMultiple(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("negative test 4", func(t *testing.T) {
		var x, result = -100.0, -1300.0

		realResult := superMultiple(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("negative test 5", func(t *testing.T) {
		var x, result = -0.02, -0.26

		realResult := superMultiple(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})
}
