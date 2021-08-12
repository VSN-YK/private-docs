/******************Testing Package B************************/

package packageB

import (
  "testing"
)

func TestPackageAAssert(t *testing.T){
  expect := "From PackageB"
  actual := DispPackageB()
  if expect != actual{
    t.Errorf("%s != %s" , expect , actual)
  }
}
