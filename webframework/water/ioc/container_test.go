package ioc

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestContainer_AddSingleton(t *testing.T) {
	c := New()

	singletonVal := "singletonValue"
	c.AddSingleton("testSingleton", func(c Container) (any, error) {
		return &singletonVal, nil
	})

	val1 := c.Get("testSingleton")
	val2 := c.Get("testSingleton")

	assert.Equal(t, &singletonVal, val1)
	assert.Equal(t, &singletonVal, val2)
	assert.Same(t, val1, val2)
}

func TestContainer_AddScoped(t *testing.T) {
	c := New()

	c.AddScoped("testScoped", func(c Container) (any, error) {
		return time.Now().UnixNano(), nil
	})

	val1 := c.Get("testScoped")
	val2 := c.Get("testScoped")

	assert.Equal(t, val1, val2) // Scoped should be same within its scope (root container is a scope)

	// Test within a new scope
	scopedCtx := c.Scoped(context.Background())
	scopedContainer := scopedCtx.Value(containerKey).(Container)

	val3 := scopedContainer.Get("testScoped")
	val4 := scopedContainer.Get("testScoped")

	assert.Equal(t, val3, val4)
	assert.NotEqual(t, val1, val3) // Scoped should be different across scopes
}

func TestContainer_Get_UnregisteredDependency(t *testing.T) {
	c := New()

	assert.PanicsWithValue(t, "there is no dependency registered with `nonExistent`", func() {
		c.Get("nonExistent")
	})
}

func TestContainer_Get_CyclicDependency(t *testing.T) {
	c := New()

	c.AddSingleton("A", func(c Container) (any, error) {
		c.Get("B")
		return "A", nil
	})
	c.AddSingleton("B", func(c Container) (any, error) {
		c.Get("A")
		return "B", nil
	})

	assert.Panics(t, func() {
		c.Get("A")
	}, "cyclic dependencies encountered while building `A`, tracked: A,B")
}

func TestContainer_Get_DependencyWithError(t *testing.T) {
	c := New()

	c.AddSingleton("errorDep", func(c Container) (any, error) {
		return nil, errors.New("failed to create dependency")
	})

	assert.PanicsWithValue(t, "error building dependency `errorDep`: failed to create dependency", func() {
		c.Get("errorDep")
	})
}

func TestContainer_Scoped(t *testing.T) {
	c := New()

	c.AddSingleton("singleton", func(c Container) (any, error) {
		return "singletonValue", nil
	})
	c.AddScoped("scoped", func(c Container) (any, error) {
		return time.Now().UnixNano(), nil
	})

	// Root container
	rootSingleton := c.Get("singleton")
	rootScoped1 := c.Get("scoped")
	rootScoped2 := c.Get("scoped")

	assert.Equal(t, rootSingleton, "singletonValue")
	assert.Equal(t, rootScoped1, rootScoped2) // Scoped should be same within its scope (root container is a scope)

	// First scoped container
	ctx1 := c.Scoped(context.Background())
	c1 := ctx1.Value(containerKey).(Container)

	c1Singleton := c1.Get("singleton")
	c1Scoped1 := c1.Get("scoped")
	c1Scoped2 := c1.Get("scoped")

	assert.Equal(t, rootSingleton, c1Singleton) // Singleton should be same across scopes
	assert.Equal(t, c1Scoped1, c1Scoped2)       // Scoped should be same within its scope
	assert.NotEqual(t, rootScoped1, c1Scoped1)  // Scoped should be different across scopes

	// Second scoped container
	ctx2 := c.Scoped(context.Background())
	c2 := ctx2.Value(containerKey).(Container)

	c2Singleton := c2.Get("singleton")
	c2Scoped1 := c2.Get("scoped")
	c2Scoped2 := c2.Get("scoped")

	assert.Equal(t, rootSingleton, c2Singleton)
	assert.Equal(t, c2Scoped1, c2Scoped2)
	assert.NotEqual(t, c1Scoped1, c2Scoped1)
}

func TestContainer_Concurrency(t *testing.T) {
	c := New()

	var counter int64
	c.AddSingleton("concurrentSingleton", func(c Container) (any, error) {
		time.Sleep(10 * time.Millisecond) // Simulate some work
		val := fmt.Sprintf("singleton-%d", counter)
		counter++
		return val, nil
	})

	var wg sync.WaitGroup
	results := make(chan any, 100)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results <- c.Get("concurrentSingleton")
		}()
	}

	wg.Wait()
	close(results)

	firstValue := <-results
	for val := range results {
		assert.Equal(t, firstValue, val, "All concurrent calls to singleton should return the same instance")
	}
	assert.Equal(t, int64(1), counter, "Singleton factory should be called only once")
}

func TestGetFromContext(t *testing.T) {
	c := New()
	c.AddSingleton("myDep", func(c Container) (any, error) {
		return "hello", nil
	})

	ctx := c.Scoped(context.Background())

	val := Get(ctx, "myDep")
	assert.Equal(t, "hello", val)

	assert.PanicsWithValue(t, "container does not exist on context", func() {
		Get(context.Background(), "myDep")
	})
}

func TestContainer_Add(t *testing.T) {
	c := New()

	customStrategyValue := "customValue"
	customStrategy := &testStrategy{value: &customStrategyValue}

	c.Add("customDep", func(c Container) (any, error) {
		return &customStrategyValue, nil
	}, customStrategy)

	val := c.Get("customDep")
	assert.Equal(t, &customStrategyValue, val)
}

type testStrategy struct {
	value any
}

func (ts *testStrategy) Get(c *container, info depInfo) any {
	return ts.value
}
