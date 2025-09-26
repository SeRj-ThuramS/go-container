package container

import (
	"testing"
)

type testStruct1 struct {
	Name string
}

type testStruct2 struct {
	Name string
}

func msgRegister(status RegisterStatus) string {
	if status == RegisterExists {
		return "the container already exists"
	} else if status == RegisterCreate {
		return "the container is registered without being initialized"
	} else if status == RegisterCreateInit {
		return "the container is registered from its initialization"
	} else if status == RegisterUnknown {
		return "the container is not registered"
	}

	return "the container is not registered"
}

func Test_RegisterRef(t *testing.T) {
	t.Run("without initialization", func(_t *testing.T) {
		ref, err := RegisterRef[testStruct1]()
		if err != nil {
			msg := msgRegister(ref)
			_t.Errorf("%s error = %v", msg, err)
			return
		}

		_t.Logf(msgRegister(ref))
	})

	t.Run("from initialization", func(_t *testing.T) {
		ref, err := RegisterRef[testStruct2](&testStruct2{
			Name: "test",
		})

		if err != nil {
			msg := msgRegister(ref)
			_t.Errorf("%s error = %v", msg, err)
			return
		}

		_t.Logf(msgRegister(ref))
	})

	t.Run("existence check", func(_t *testing.T) {
		ref, err := RegisterRef[testStruct1]()

		if err != nil {
			msg := msgRegister(ref)
			_t.Errorf("%s error = %v", msg, err)
			return
		}

		_t.Logf(msgRegister(ref))
	})
}

func Test_AssignRef(t *testing.T) {
	ref, err := RegisterRef[testStruct1](&testStruct1{
		Name: "Container 1",
	})
	if err != nil {
		msg := msgRegister(ref)
		t.Errorf("%s error = %v", msg, err)
		return
	}

	t.Logf(msgRegister(ref))

	t.Run("assign an existing container", func(_t *testing.T) {
		provider, err := AssignRef[testStruct1]()
		if err != nil {
			_t.Errorf("%v", err)
			return
		}

		_t.Logf("container = %s", provider.Name)
	})

	t.Run("assign a non-existent container", func(_t *testing.T) {
		_, err := AssignRef[testStruct2]()
		if err != nil {
			_t.Logf("%v", err)
			return
		}
	})
}
