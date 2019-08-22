package parentpointertree_test

import (
	"testing"
	"github.com/Yalhash/scheduler/parentpointertree"
)

func TestGetValue(t *testing.T){
    n := new(parentpointertree.Node)
    n.SetValue(6)
    if n.GetValue() != 6 {
        t.Error("expected node value of 6")
    }
}

func TestSetValue(t *testing.T) {
    n := new(parentpointertree.Node)
    n.SetValue(6)
    if n.GetValue() != 6 {
        t.Error("expected node value of 6")
    }
    n.SetValue(8)
    if n.GetValue() != 8 {
        t.Error("expected node value of 8")
    }
}


func TestGetParentSetParent(t *testing.T) {
    n1 := new(parentpointertree.Node)
    n2 := new(parentpointertree.Node)
    n2.SetParent(n1)
    if n2.GetParent() != n1 {
	t.Error("expected reference to n1")
    }

}



func TestAddChildGetChild(t * testing.T) {
    testPPT := new(parentpointertree.PPT)
    testPPT.AddChild(10)
    if testPPT.GetChild(0).GetValue() != 10 {
	t.Error("expected node of value 10")
    }
    testPPT.AddChild(20)
    if testPPT.GetChild(0).GetValue() != 10 {
	t.Error("expected node of value 10")
    }
    if testPPT.GetChild(1).GetValue() != 20 {
	t.Error("expected node of value 20")
    }
}

