package ch22

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

/**
对象池
*/

type ReusableObj struct {
}
type ObjPoll struct {
	bufChan chan *ReusableObj
}

/**
预先创建些连接
*/
func NewObjPool(numOfObj int) *ObjPoll {
	objPoll := ObjPoll{}
	objPoll.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPoll.bufChan <- new(ReusableObj)
	}
	return &objPoll
}

func (p *ObjPoll) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout): // 超时控制
		return nil, errors.New("time out")
	}
}

func (p *ObjPoll) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestObjPoll(t *testing.T) {
	pool := NewObjPool(10)
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Print(v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
	t.Log("done")
}
