package main
import "testing"

func TestCommand(t *testing.T) {
  cases := []struct {
	 function func(c *config) error
	 expectedError error
  }{
	
	{
		function : help,
		expectedError : nil,
	},
  }

  con := config{
			commandRegistry : getCommands(),
			previous : "",
			next : "",
	}

  for _,c := range cases {
	actual := c.function(&con)
	if c.expectedError != actual {
		t.Errorf("expected  : %v, got: %v",c.expectedError, actual)
	}
  }
}