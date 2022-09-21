package main

type IObservable interface {
	Subscribe(Observer Observer)
	Unsubscribe(Observer Observer)
	NotifyAll()
}
