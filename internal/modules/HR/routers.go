package hr

func (S *HRService) registerRoutes() {
	// route gorup name
	hrRoute := S.e.Group("/hr-api")

	// admin routes
	hrRoute.POST("/admin", S.createAdmin)
	hrRoute.POST("/admin/tax", S.createTax)
	hrRoute.POST("/admin/payroll", S.createPayroll)
	hrRoute.GET("/admin/payroll", S.getPayroll)
	hrRoute.GET("/admin/tax", S.getTax)
	hrRoute.GET("/admin/payroll/:id", S.getOnePayroll)
	hrRoute.PUT("/admin/payroll/:id", S.updatePayroll)
	hrRoute.PUT("/admin/suspend", S.suspendAdmin)

	//employee routes
	hrRoute.POST("/employee", S.createEmployee)
	hrRoute.GET("/employee", S.getEmployee)
	hrRoute.GET("/employee/:id", S.getEmployeeOne)
	hrRoute.PUT("/employee/:id", S.updateEmployee)
	hrRoute.PUT("/employee/emergency", S.updateEmpEmergencyDetails)
	hrRoute.PUT("/employee/bank", S.updateEmpBankDetails)
	hrRoute.PUT("/employee/salary", S.updateEmpSalary)
	hrRoute.PUT("/employee/certificate", S.updateEmpCertificates)
	hrRoute.PUT("/employee/status", S.updateEmpStatus)
	hrRoute.PUT("/employee/benefits", S.updateEmpBenifits)
	hrRoute.PUT("/employee/user", S.updateEmpUser)
	hrRoute.PUT("/employee/allowances", S.updateEmpAllowances)
	hrRoute.PUT("/employee/expatriate", S.updateEmpExpatriate)
	hrRoute.PUT("/employee/accessibility", S.updateEmpAccessiability)
	hrRoute.DELETE("/employee/:id", S.deleteEmployee)
	hrRoute.GET("/employee/login", S.employeeLogin)

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
