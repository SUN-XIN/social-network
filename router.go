package main

type Router struct {
	DB *Store
}

// NewRouter returns a default router.
func NewRouter() *Router {
	return &Router{
		DB: NewStore(),
	}
}
