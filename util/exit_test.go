package util
import "testing"
import "time"


func TestSetExit(t *testing.T) {
	go SetExit()
	go func() {
		for{
			t.Logf("exit: %v\n", exit)
			time.Sleep(1* time.Second)
		}
	}()
	time.Sleep(10 * time.Second)
	//这里模拟输入到控制台一个 q
	for{
			t.Logf("exit: %v\n", exit)
			time.Sleep(1* time.Second)
		}
}
