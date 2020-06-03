package shadowclonefs

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestSerialization(t *testing.T) {
	f, err := ioutil.TempFile("", "TestSerialization_*.msgpack")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	fsmob, err := FileMetadata(f.Name())
	if err != nil {
		t.Error(err)
	}

	blob, err := fsmob.Serialize()
	if err != nil {
		t.Error(err)
	}
	_, err = f.Write(blob)
	if err != nil {
		t.Error(err)
	}

	fsmobnew, err := DeserializeFSObjectMetadata(blob)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(fsmob, fsmobnew) {
		t.Errorf("Objects not equal:\n%#v\n%#v\n", fsmob, fsmobnew)
	}
}
