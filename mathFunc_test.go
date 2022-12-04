package main

import (
	"fmt"
	"testing"
)

//go test -v -run /zero
//go test -run null (название любого несуществующего теста) - посмотреть ошибки компиляции

func TestMathFunc(t *testing.T) {
	fmt.Println("SETUP") // создаем все что нужно для теста

	//чистим все после тестов
	t.Cleanup(func() {
		fmt.Println("CLEANUP")
	})

	t.Run("positive + zero group", func(t *testing.T) {
		t.Log("positive + zero group")
		t.Run("zero", func(t *testing.T) {
			t.Parallel()
			t.Log("zero")
			var x, result = 0.0, 0.0

			realResult := mathFunc(x)
			if realResult != result {
				t.Errorf("expected result %f != real result %f", result, realResult)
			}
		})

		t.Run("positive test 1", func(t *testing.T) {
			t.Parallel()
			t.Log("positive test 1")
			var x, result = 0.1, 1.1998325788384403

			realResult := mathFunc(x)
			if realResult != result {
				t.Errorf("expected result %f != real result %f", result, realResult)
			}
		})

		t.Run("positive test 2", func(t *testing.T) {
			t.Parallel()
			t.Log("positive test 2")
			var x, result = 0.3, 3.5953073459846023

			realResult := mathFunc(x)
			if realResult != result {
				t.Errorf("expected result %f != real result %f", result, realResult)
			}
		})

		t.Run("positive test 3", func(t *testing.T) {
			t.Parallel()
			t.Log("positive test 3")
			var x, result = 0.6, 7.156498891206716
			realResult := mathFunc(x)
			if realResult != result {
				t.Errorf("expected result %f != real result %f", result, realResult)
			}
		})

		t.Run("positive test 4", func(t *testing.T) {
			t.Parallel()
			t.Log("positive test 4")
			var x, result = 0.8, 9.472704781998388

			realResult := mathFunc(x)
			if realResult != result {
				t.Errorf("expected result %f != real result %f", result, realResult)
			}
		})
	})

	t.Run("negative group", func(t *testing.T) {
		t.Log("negative group")
		t.Run("negative test 1", func(t *testing.T) {
			t.Parallel()
			t.Log("negative test 1")
			var x, result = -0.1, -1.1998325788384403

			realResult := mathFunc(x)
			if realResult != result {
				t.Errorf("expected result %f != real result %f", result, realResult)
			}
		})

		t.Run("negative test 2", func(t *testing.T) {
			t.Parallel()
			t.Log("negative test 2")
			var x, result = -0.3, -3.5953073459846023

			realResult := mathFunc(x)
			if realResult != result {
				t.Errorf("expected result %f != real result %f", result, realResult)
			}
		})

		t.Run("negative test 3", func(t *testing.T) {
			t.Parallel()
			t.Log("negative test 3")
			var x, result = -0.6, -7.156498891206716

			realResult := mathFunc(x)
			if realResult != result {
				t.Errorf("expected result %f != real result %f", result, realResult)
			}
		})

		t.Run("negative test 4", func(t *testing.T) {
			t.Parallel()
			t.Log("negative test 4")
			var x, result = -0.8, -9.472704781998388

			realResult := mathFunc(x)
			if realResult != result {
				t.Errorf("expected result %f != real result %f", result, realResult)
			}
		})

		t.Run("negative test 5", func(t *testing.T) {
			t.Parallel()
			t.Log("negative test 5")
			var x, result = -1.0, -11.429203673205103

			realResult := mathFunc(x)
			if realResult != result {
				t.Errorf("expected result %f != real result %f", result, realResult)
			}
		})
	})
}
