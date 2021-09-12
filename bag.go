package config

import (
	"fmt"
	"strings"
	"sync"
)

type NotFoundError struct {
	target string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Parameter \"%s\" was not found", e.target)
}

type DotNotationBag struct {
	sync.Mutex
	bag map[string]interface{}
}

func (b *DotNotationBag) Get(key string) (interface{}, error) {
	keyChain := strings.Split(key, ".")
	current := &b.bag

	for _, link := range keyChain[:len(keyChain)-1] {
		if _, ok := (*current)[link]; !ok {
			return "", &NotFoundError{}
		}

		c := (*current)[link].(map[string]interface{})
		current = &c
	}

	return (*current)[keyChain[len(keyChain)-1]], nil
}

func (b *DotNotationBag) Has(key string) bool {
	keyChain := strings.Split(key, ".")
	current := &b.bag

	for _, link := range keyChain[:len(keyChain)-1] {
		if _, ok := (*current)[link]; !ok {
			return false
		}

		c := (*current)[link].(map[string]interface{})
		current = &c
	}

	if value, ok := (*current)[keyChain[len(keyChain)-1]]; ok {
		return value != nil
	}

	return false
}

func (b *DotNotationBag) Set(key, value string) {
	keyChain := strings.Split(key, ".")

	if b.bag == nil {
		b.bag = make(map[string]interface{})
	}

	current := &b.bag
	b.Lock()

	for _, link := range keyChain[:len(keyChain)-1] {
		if _, ok := (*current)[link]; !ok {
			(*current)[link] = make(map[string]interface{})
		}

		c, _ := ((*current)[link]).(map[string]interface{})
		current = &c
	}

	(*current)[keyChain[len(keyChain)-1]] = value
	b.Unlock()
}

func (b *DotNotationBag) Load(loader func(interface{}) error) error {
	b.Lock()
	defer b.Unlock()

	return loader(b.bag)
}
