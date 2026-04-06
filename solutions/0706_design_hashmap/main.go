package main

type MyHashMap struct {
	data [1000001]int
}

func Constructor() MyHashMap {
	m := MyHashMap{}
	for i := range m.data {
		m.data[i] = -1
	}
	return m
}

func (this *MyHashMap) Put(key int, value int) {
	this.data[key] = value
}

func (this *MyHashMap) Get(key int) int {
	return this.data[key]
}

func (this *MyHashMap) Remove(key int) {
	this.data[key] = -1
}
