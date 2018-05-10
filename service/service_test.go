package service

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestServiceGetWhatYouPut(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testKey := int64(42)
	testVal := "hello"

	kvs := NewMockKeyValueStore(ctrl)
	kvs.EXPECT().Put(testKey, testVal).Return(nil)
	kvs.EXPECT().Get(testKey).Return(testVal, nil)

	svc := New(kvs)

	err := svc.PutString(testKey, testVal)
	if err != nil {
		t.Error(err)
	}

	val, err := svc.GetString(testKey)
	if err != nil {
		t.Error(err)
	}

	if val != testVal {
		t.Errorf("expected %s, got %s", testVal, val)
	}

}
