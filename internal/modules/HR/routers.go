package hr

func (S *HRService) registerRoutes() {
	// route gorup name
	hrRoute := S.e.Group("/hr-api")

	// routes
	hrRoute.GET("/", S.createAdmin)

	//employee routes
	hrRoute.POST("/employee", S.createEmployee)
	hrRoute.GET("/employee", S.getEmployee)
	hrRoute.GET("/employee/:id", S.getEmployeeOne)
	hrRoute.PUT("/employee/:id", S.updateEmployee)
	hrRoute.DELETE("/employee/:id", S.deleteEmployee)

	//service routes
	hrRoute.POST("/service", S.createAdminService)
	hrRoute.GET("/service", S.getAdminServices)
	hrRoute.GET("/service/:category", S.getOneAdminService)
	hrRoute.PUT("/service/:id", S.updateAdminService)
	hrRoute.DELETE("/service/:id", S.deleteAdminService)

	//allowance routes
	hrRoute.POST("/allowance", S.createAllowances)
	hrRoute.GET("/allowance", S.getAllowances)
	hrRoute.GET("/allowance/:id", S.getOneAllowance)
	hrRoute.PUT("/allowance/:id", S.updateAllowance)
	hrRoute.DELETE("/allowance/:id", S.deleteAllowance)

}
