package hr

func (S *HRService) registerRoutes() {
	// route gorup name
	hrRoute := S.e.Group("/hr-api")

	// routes
	hrRoute.GET("/", S.createAdmin)
}
