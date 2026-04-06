package main

type MyHashSet struct {
	data [1000001]bool
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
	this.data[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.data[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
	return this.data[key]
}
