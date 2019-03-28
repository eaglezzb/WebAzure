package main

import "github.com/gin-gonic/gin"
import "net/http"
func sayHello(c *gin.Context){
}

func main(){
	v1 := router.Group("/v1")
	v2 := router.Group("/v2")

	g := NewGroupGroup([]*gin.RouterGroup { v1, v2 })

	g.handle(http.MethodGet, "hello", sayHello)
	g.handle(http.MethodPost, "goodbye", sayGoodbye)

}

/*
type GroupGroup struct {
    groups []*gin.RouterGroup
}

func NewGroupGroup(groups []*gin.RouterGroup) GroupGroup {
    return GroupGroup {
        groups,
    }
}

func (g *GroupGroup) handle(method string, path string, handler gin.HandlerFunc) {
    for _, group := range g.groups {
        group.Handle(method, path, handler)
    }
}
Then, you can use it like so :

v1 := router.Group("/v1")
v2 := router.Group("/v2")

g := NewGroupGroup([]*gin.RouterGroup { v1, v2 })

g.handle(http.MethodGet, "hello", sayHello)
g.handle(http.MethodPost, "goodbye", sayGoodbye)
*/