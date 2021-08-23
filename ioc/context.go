package ioc

var contextInstance = newContext()

type Context struct {
	// 存放bean
	beanMap map[string]interface{}
}

type ContextFacade interface {
	// 增加bean
	add(id string, bean interface{}) (bool, error)

	// 删除bean
	delete(id string) (bool, error)

	// 获取bean
	getBeanById(id string) interface{}

	// 获取bean
	getBeanByName(id string) interface{}

	// 获取bean
	getBeanByType(class interface{}) interface{}
}

// 构建context
func newContext() *Context {
	return &Context{
		beanMap: map[string]interface{}{},
	}
}

func (c *Context) add(id string, bean interface{}) (bool, error) {
	c.beanMap[id] = bean
	return true, nil
}

func (c *Context) delete(id string) (bool, error) {
	c.beanMap[id] = nil
	return true, nil
}

func (c *Context) getBeanById(id string) interface{} {
	return c.beanMap[id]
}

func (c *Context) getBeanByName(id string) interface{} {
	panic("implement me")
}

func (c *Context) getBeanByType(class interface{}) interface{} {
	panic("implement me")
}

// 获取context
func GetContext() *Context {
	return contextInstance
}
