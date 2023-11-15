package SLinkedList

import (
	"DsaGo/dsago/geom/Point"
	"DsaGo/utils"
	"fmt"
)

func SListDemo1() {
	slist := NewSLinkedList(nil)
	for i := 0; i < 20; i++ {
		slist.Add(i)
	}
	println(slist.ToString(nil))
}

func SListDemo2() {
	slist := NewSLinkedList(Point.PointEqual)

	for i := 0; i < 20; i++ {
		slist.Add(Point.NewPoint(float64(i), float64(i)))
	}

	println(slist.ToString(utils.Point2str))
}

func SListDemo3() {
	slist := NewSLinkedList(Point.PointEqual)

	for i := 0; i < 20; i++ {
		slist.Add(Point.NewPoint(float64(i), float64(i)))
	}

	println(slist.ToString(utils.Point2str))

	_, err := slist.RemoveAt(0)
	if err != nil {
		return
	}
	println(slist.ToString(utils.Point2str))
}

func SListDemo4() {
	slist := NewSLinkedList(Point.PointEqual)

	for i := 0; i < 20; i++ {
		slist.Add(Point.NewPoint(float64(i), float64(i)))
	}

	println(slist.ToString(utils.Point2str))

	err := slist.AddAtIndex(20, Point.NewPoint(100, 100))
	if err != nil {
		return
	}
	err = slist.AddAtIndex(slist.count, Point.NewPoint(200, 200))
	if err != nil {
		return
	}
	err = slist.AddAtIndex(0, Point.NewPoint(300, 300))
	if err != nil {
		return
	}
	err = slist.AddAtIndex(1, Point.NewPoint(400, 400))
	if err != nil {
		return
	}
	println(slist.ToString(utils.Point2str))
}

// test PointEqual
func SlistDemo5() {
	list1 := NewSLinkedList(Point.PointEqual)

	list1.Add(Point.NewPoint(23.2, 25.4))
	list1.Add(Point.NewPoint(24.6, 23.1))
	list1.Add(Point.NewPoint(12.5, 22.3))

	println(list1.ToString(utils.Point2str))

	p1 := Point.NewPoint(23.2, 25.4)   // Yes
	p2 := Point.NewPoint(123.6, 233.1) // No

	fmt.Println(*p1, "=>", list1.Contains(p1), "\n", "indexOf returns:", list1.IndexOf(p1))
	fmt.Println(*p2, "=>", list1.Contains(p2), "\n", "indexOf returns:", list1.IndexOf(p2))

	// Different result without PointEqual
	list2 := NewSLinkedList(nil)

	list2.Add(Point.NewPoint(23.2, 25.4))
	list2.Add(Point.NewPoint(24.6, 23.1))
	list2.Add(Point.NewPoint(12.5, 22.3))
	println(list2.ToString(utils.Point2str))

	fmt.Println(*p1, "=>", list2.Contains(p1), "\n", "indexOf returns:", list2.IndexOf(p1))
	fmt.Println(*p2, "=>", list2.Contains(p2), "\n", "indexOf returns:", list2.IndexOf(p2))
}
