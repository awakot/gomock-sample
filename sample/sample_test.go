package sample

import (
	"testing"

	"github.com/awakot/mock-sample/sample/mock_sample"
	"github.com/golang/mock/gomock"
)

func TestSample1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSample := mock_sample.NewMockSample(ctrl)
	mockSample.EXPECT().Method("hoge").Return(1)

	t.Log("result:", mockSample.Method("hoge"))
}

func TestSample2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSample := mock_sample.NewMockSample(ctrl)

	// 呼び出される順番を指定
	gomock.InOrder(
		mockSample.EXPECT().Method("hoge").Return(1),
		mockSample.EXPECT().Method("fuga").Return(2),
	)
	/* // 上記と同じ
	   mockSample.EXPECT().Method("hoge").Return(1).After(
	       mockSample.EXPECT().Method("fuga").Return(2),
	   )
	*/

	t.Log("result", mockSample.Method("hoge"))
	t.Log("result", mockSample.Method("fuga"))
}
