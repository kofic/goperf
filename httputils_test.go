package main

import (
	//"bytes"
	"fmt"
	"github.com/fatih/color"
	"github.com/gnulnx/goperf/httputils"
	"io/ioutil"
	"reflect"
	"testing"
)

func get_test_body() string {
	body_bytes, _ := ioutil.ReadFile("test_data/test.html")
	num_bytes := len(body_bytes)
	body := string(body_bytes[:num_bytes])
	return body
}

func test_deep_equal(input []string, testdata []string, t *testing.T) bool {
	if reflect.DeepEqual(input, testdata) != true {
		fmt.Println(input)
		fmt.Println(testdata)
		t.Error("Slices above are not equal")
		return false
	}
	return true
}

func TestGetjs(t *testing.T) {
	color.Green("~~ TestGetjs ~~")
	body := get_test_body()

	jsfiles := httputils.Getjs(body)
	test_data := []string{
		`/static/tcart/js/test1.min.js`,
		`/static/tcart/js/bundle_kldsf2334.min.js`,
	}
	test_deep_equal(*jsfiles, test_data, t)
}

func TestGetimg(t *testing.T) {
	color.Green("~~ TestGetimg ~~")
	body := get_test_body()
	imgfiles := httputils.Getimg(body)
	test_data := []string{
		`/media/cart.svg`,
		`/static/tcart/img/stripe_badges/outline_dark/powered_by_stripe.png`,
	}
	test_deep_equal(*imgfiles, test_data, t)
}

func TestGetcss(t *testing.T) {
	color.Green("~~ TestGetcss ~~")
	body := get_test_body()
	cssfiles := httputils.Getcss(body)
	test_data := []string{
		`/media/manifest.webmanifest`,
		`/static/vendor/icomoon/style.css`,
		`/media/favicon_94S_icon.ico`,
		`/static/vendor/bootstrap/bootstrap.min.css`,
		`/static/tcart/css/styles.min.css`,
	}
	test_deep_equal(*cssfiles, test_data, t)
}