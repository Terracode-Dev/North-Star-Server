package hr

import (
	rba "github.com/Terracode-Dev/North-Star-Server/internal/pkg/RBA"
)

func (S *HRService) registerRoutes() {
	// route group name
	hrRoute := S.e.Group("/hr-api")

	// admin routes
	hrRoute.GET("/admin", S.getAllAdmin, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.POST("/admin", S.createAdmin, rba.AuthMiddelware([]string{"admin"}))
	hrRoute.PUT("/admin/suspend", S.suspendAdmin, rba.AuthMiddelware([]string{"admin"}))
	hrRoute.PUT("/admin/:id", S.updateAdmin, rba.AuthMiddelware([]string{"admin"}))
	hrRoute.POST("/admin/login", S.adminLogin)

	// employee routes
	hrRoute.POST("/employee", S.createEmployee, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.POST("/employee/all", S.getEmployee, rba.AuthMiddelware([]string{"admin", "emp", "mod"}))
	hrRoute.GET("/employee/:id", S.getEmployeeOne, rba.AuthMiddelware([]string{"admin", "emp", "mod"}))
	hrRoute.PUT("/employee/:id", S.updateEmployee, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.PUT("/employee/emergency", S.updateEmpEmergencyDetails, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.PUT("/employee/bank", S.updateEmpBankDetails, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.PUT("/employee/salary", S.updateEmpSalary, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.PUT("/employee/certificate", S.updateEmpCertificates, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.PUT("/employee/status", S.updateEmpStatus, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.PUT("/employee/benefits", S.updateEmpBenifits, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.PUT("/employee/user", S.updateEmpUser, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.PUT("/employee/allowances", S.updateEmpAllowances, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.PUT("/employee/expatriate", S.updateEmpExpatriate, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.PUT("/employee/accessibility", S.updateEmpAccessiability, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.DELETE("/employee/:id", S.deleteEmployee, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.POST("/employee/login", S.employeeLogin)
	hrRoute.PUT("/employee/empbank", S.empOnlyBankDetailsUpdate, rba.AuthMiddelware([]string{"emp"}))

	// service routes
	hrRoute.POST("/service", S.createAdminService, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/service", S.getAdminServices, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/service/:category", S.getOneAdminService, rba.AuthMiddelware([]string{"admin", "mod", "emp"}))
	hrRoute.PUT("/service/:id", S.updateAdminService, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.DELETE("/service/:id", S.deleteAdminService, rba.AuthMiddelware([]string{"admin", "mod"}))

	// allowance routes
	hrRoute.POST("/allowance", S.createAllowances, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/allowance", S.getAllowances, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/allowance/:id", S.getOneAllowance, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.PUT("/allowance/:id", S.updateAllowance, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.DELETE("/allowance/:id", S.deleteAllowance, rba.AuthMiddelware([]string{"admin", "mod"}))

	hrRoute.GET("/tax", S.getTax, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.POST("/tax", S.createTax, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.POST("/payroll", S.createPayroll, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/payroll", S.getPayroll, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/payroll/:id", S.getOnePayroll, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.PUT("/payroll/:id", S.updatePayroll, rba.AuthMiddelware([]string{"admin", "mod"}))

	hrRoute.POST("/fileupload", S.uploadFile, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.POST("/getfileurl", S.getFileDownloadUrl, rba.AuthMiddelware([]string{"admin", "mod"}))

	hrRoute.GET("/logout", S.Logout)

	hrRoute.POST("/testlogin", S.TestLogin)
	hrRoute.GET("/testauth", S.TestAuth, rba.AuthMiddelware([]string{"admin", "emp"}))
	hrRoute.POST("/testS3upload", S.TestS3Upload)
}
