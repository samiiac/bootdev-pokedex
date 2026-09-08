package main
import "testing"

func TestCommand(t *testing.T) {
  cases := []struct {
	 function func() error
	 expectedError error
  }{
	
	{
		function : help,
		expectedError : nil,
	},
  }


  for _,c := range cases {
	actual := c.function()
	if c.expectedError != actual {
		t.Errorf("expected  : %v, got: %v",c.expectedError, actual)
	}
  }
}