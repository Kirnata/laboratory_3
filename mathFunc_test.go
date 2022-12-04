package main

import "testing"

func TestMathFunc(t *testing.T) {

	t.Run("zero", func(t *testing.T) {
		var x, result = 0.0, 0.0

		realResult := mathFunc(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("positive test 1", func(t *testing.T) {
		var x, result = 0.1, 1.1998325788384403

		realResult := mathFunc(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("positive test 2", func(t *testing.T) {
		var x, result = 0.3, 3.5953073459846023

		realResult := mathFunc(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("positive test 3", func(t *testing.T) {
		var x, result = 0.6, 7.156498891206716
		realResult := mathFunc(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("positive test 4", func(t *testing.T) {
		var x, result = 0.8, 9.472704781998388

		realResult := mathFunc(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("negative test 1", func(t *testing.T) {
		var x, result = -0.1, -1.1998325788384403

		realResult := mathFunc(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("negative test 2", func(t *testing.T) {
		var x, result = -0.3, -3.5953073459846023

		realResult := mathFunc(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("negative test 3", func(t *testing.T) {
		var x, result = -0.6, -7.156498891206716

		realResult := mathFunc(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("negative test 4", func(t *testing.T) {
		var x, result = -0.8, -9.472704781998388

		realResult := mathFunc(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})

	t.Run("negative test 5", func(t *testing.T) {
		var x, result = -1.0, -11.429203673205103

		realResult := mathFunc(x)
		if realResult != result {
			t.Errorf("expected result %f != real result %f", result, realResult)
		}
	})
}
