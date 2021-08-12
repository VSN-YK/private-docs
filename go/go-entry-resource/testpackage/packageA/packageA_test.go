package packageA

import(
  "testing"
)

func TestPackageAAssert(t *testing.T){
  expect := "From PackageA"
  actual := DispPackageA()
  if expect != actual {
    t.Errorf("%s != %s" , expect , actual)
  }
}
