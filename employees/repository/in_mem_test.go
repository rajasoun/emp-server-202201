package repository_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"algogrit.com/httpex/employees/repository"
	"algogrit.com/httpex/entities"
)

func TestConsistency(t *testing.T) {
	sut := repository.NewInMem()

	initialEmps, err := sut.FindAll()

	assert.Nil(t, err)

	assert.Equal(t, 3, len(initialEmps))

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		newEmp := entities.Employee{0, "Gaurav", "LnD", 1001}
		go func() {
			defer wg.Done()
			sut.Save(newEmp)
			_, _ = sut.FindAll()
		}()
	}

	wg.Wait()
	actualEmps, err := sut.FindAll()

	assert.Nil(t, err)
	assert.Equal(t, 103, len(actualEmps))
}
