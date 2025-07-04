package hr

import (
	rba "github.com/Terracode-Dev/North-Star-Server/internal/pkg/RBA"
)

func (S *HRService) registerRoutes() {
	// route group name
	hrRoute := S.e.Group("/hr-api")

	// admin routes
	hrRoute.POST("/admin/all", S.getAllAdmin, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.POST("/admin", S.createAdmin, rba.AuthMiddelware([]string{"admin"}))
	hrRoute.PUT("/admin/suspend", S.suspendAdmin, rba.AuthMiddelware([]string{"admin"}))
	hrRoute.PUT("/admin/:id", S.updateAdmin, rba.AuthMiddelware([]string{"admin"}))
	hrRoute.POST("/admin/login", S.adminLogin)

	// employee routes
	hrRoute.POST("/employee", S.createEmployee, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.POST("/employee/all", S.getEmployee, rba.AuthMiddelware([]string{"admin", "emp", "mod"}))
	hrRoute.GET("/employee/:id", S.getEmployeeOne, rba.AuthMiddelware([]string{"admin", "emp", "mod"}))
	hrRoute.GET("/employeefrombranch", S.getEmployeeByBranch, rba.AuthMiddelware([]string{"admin", "emp", "mod"}))
	hrRoute.GET("/employeesalary/:id", S.getEmployeeSalary, rba.AuthMiddelware([]string{"admin", "emp", "mod"}))
	hrRoute.GET("/employeeallowances/:id", S.getEmployeeAllowances, rba.AuthMiddelware([]string{"admin", "emp", "mod"}))
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
	// update trainer commission
	hrRoute.PUT("/employee/trainerCom", S.UpdateEmpCommission, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.DELETE("/employee/:id", S.deleteEmployee, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.POST("/employee/login", S.employeeLogin)
	hrRoute.PUT("/employee/empbank", S.empOnlyBankDetailsUpdate, rba.AuthMiddelware([]string{"emp"}))
	hrRoute.POST("/checkTrainer", S.CheckIfEMPIsTrainer, rba.AuthMiddelware([]string{"admin", "mod"}))
	// Delete employee certificates from the HR_FileSub,it table and S3
	hrRoute.DELETE("/employee/deletefiles", S.DeleteEmployeeFiles, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/employeefiles/:id", S.GetFileData, rba.AuthMiddelware([]string{"admin", "mod", "emp"}))
	hrRoute.POST("/employee/checktodaysession", S.CheckFortodayTrainerClientSession)

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
	hrRoute.DELETE("/tax/:id", S.deleteTax, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.POST("/payroll", S.createPayroll, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/payroll", S.getPayroll, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/payroll/:id", S.getOnePayroll, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.PUT("/payroll/:id", S.updatePayroll, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/calculatetrainercom/:trainer_id", S.CalculateTrainerCommision)

	hrRoute.POST("/hrbranch", S.addHRBranch, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/hrbranch", S.getAllHRBranch, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/hrbranch/protect", S.getProtectedHRBranch, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.DELETE("/hrbranch/:id", S.deleteHRBranch, rba.AuthMiddelware([]string{"admin", "mod"}))

	hrRoute.POST("/fileupload", S.uploadFile)
	hrRoute.POST("/getfileurl", S.getFileDownloadUrl, rba.AuthMiddelware([]string{"admin", "mod", "emp"}))

	hrRoute.GET("/verify-auth-emp", S.empVerifyAuth, rba.AuthMiddelware([]string{"emp"}))
	hrRoute.GET("/verify-auth", S.verifyAuth, rba.AuthMiddelware([]string{"admin", "emp", "mod", "floor_manager"}))
	hrRoute.GET("/logout", S.Logout)

	hrRoute.POST("/testlogin", S.TestLogin)
	hrRoute.GET("/testauth", S.TestAuth, rba.AuthMiddelware([]string{"admin", "emp"}))
	hrRoute.POST("/testS3upload", S.TestS3Upload)

	hrRoute.POST("/exchange-rate", S.CreateExchangeRate, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.DELETE("/exchange-rate/:id", S.DeleteExchangeRate, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/exchange-rate/:type", S.GetExchangeRate, rba.AuthMiddelware([]string{"admin", "mod", "emp"}))
	hrRoute.GET("/exchange-rate", S.GetExchangeRateAll, rba.AuthMiddelware([]string{"admin", "mod", "emp"}))

	// v2 routes
	hrRoute.POST("/V2/preset", S.CreatePreset, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/V2/preset", S.GetAllPresets, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/V2/preset/:id", S.GetPresetByID, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.PUT("/V2/preset/:id", S.UpdatePreset, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.DELETE("/V2/preset/:id", S.DeletePreset, rba.AuthMiddelware([]string{"admin", "mod"}))

	hrRoute.POST("/V2/preset-workout", S.CreatePresetWorkout, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/V2/preset-workout", S.GetAllPresetWorkouts, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.PUT("/V2/preset-workout/:id", S.UpdatePresetWorkout, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.DELETE("/V2/preset-workout/:id", S.DeletePresetWorkout, rba.AuthMiddelware([]string{"admin", "mod"}))

	hrRoute.POST("/V2/session", S.CreateSession, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/V2/session", S.GetAllSessions, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.DELETE("/V2/session/:id", S.DeleteSession, rba.AuthMiddelware([]string{"admin", "mod"}))

	hrRoute.POST("/V2/preset-session", S.CreatePresetSession, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/V2/preset-session", S.GetAllPresetSession, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.PUT("/V2/preset-session/:id", S.UpdatePresetSession, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.DELETE("/V2/preset-session/:id", S.DeletePresetSession, rba.AuthMiddelware([]string{"admin", "mod"}))

	hrRoute.POST("/V2/session-workout", S.CreateSessionWorkout, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.GET("/V2/session-workout", S.GetAllSessionWorkouts, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.PUT("/V2/session-workout/:id", S.UpdateSessionWorkout, rba.AuthMiddelware([]string{"admin", "mod"}))
	hrRoute.DELETE("/V2/session-workout/:id", S.DeleteSessionWorkout, rba.AuthMiddelware([]string{"admin", "mod"}))

	leaves := hrRoute.Group("/leaves")

	leaves.POST("", S.CreateLeaveHandler)
	leaves.GET("", S.GetAllLeavesHandler)
	leaves.POST("/checkuser", S.CheckValideteEmp)
	leaves.GET("/:id", S.GetLeaveByIdHandler)
	leaves.PUT("/:id", S.UpdateLeaveHandler)
	leaves.DELETE("/:id", S.DeleteLeaveHandler)
	leaves.GET("/:id/emp-leave", S.GetEmployeeLeavesHandler)
	leaves.GET("/:id/benifit-leave", S.GetEmployeeLeaveBenefitsHandler)

	empSchedule := hrRoute.Group("/schedule")

	// Employee lookup
	empSchedule.POST("/employee/id-by-email", S.GetEmployeeIdByEmail)

	// Schedule CRUD operations
	empSchedule.POST("/employee/schedule", S.CreateEmployeeSchedule)
	empSchedule.PUT("/employee/:id/schedule", S.UpdateEmployeeSchedule)
	empSchedule.DELETE("/employee/:id/schedule", S.DeleteEmployeeSchedule)
	empSchedule.DELETE("/employee/:id/schedule/additional/:date", S.DeleteAdditionalSchedule)

	// Employee list with work days
	empSchedule.GET("/employees", S.GetEmployeeList)
	empSchedule.GET("/employee/:id/workdays-breakdown", S.GetEmployeeWorkDaysBreakdown)
}
