package serial

import "testing"

func TestConfigGet(t *testing.T) {
	c, e := NewConfigService().Get()
	if e != nil {
		t.Fatalf("%+v", e)
	}
	t.Logf("Value %d, Port %s", c.Value, c.Port)
}

func TestConfigSave(t *testing.T) {
	c := &Config{
		Value: 1,
		Port:  ":234",
	}
	e := NewConfigService().Save(c)
	if e != nil {
		t.Fatalf("%+v", e)
	}
	t.Log("Done.")
}
