package dao

import (
	passwordvalidator "github.com/wagslane/go-password-validator"

	"github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	logs "github.com/mt1976/ebEstimates/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/dao/credentialspassword.go"
// ----------------------------------------------------------------
// Package              : dao
// Object 			    : CredentialsPassword (credentialspassword)
// Endpoint 	        : CredentialsPassword (ID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 24/02/2023 at 16:16:46
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------
// The following functions should be created in credentialspassword_impl.go
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// CredentialsPassword_ObjectValidation_impl provides Record/Object level validation for CredentialsPassword
func CredentialsPassword_ObjectValidation_impl(iAction string, iId string, iRec dm.CredentialsPassword) (dm.CredentialsPassword, string, error) {
	logs.Callout("CredentialsPassword", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("CredentialsPassword" + " - Invalid Action [" + iAction + "]")
	}
	return iRec, "", nil
}

//
// ----------------------------------------------------------------
// These are the the default implementations, which do nothing

// ----------------------------------------------------------------
// CredentialsPassword_UserName_validate_impl provides validation/actions for UserName
func CredentialsPassword_UserName_validate_impl(iAction string, iId string, iValue string, iRec dm.CredentialsPassword, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("CredentialsPassword", dm.CredentialsPassword_UserName_scrn, VAL+"-"+iAction, iId)
	//fmt.Println("CredentialsPassword_UserName_validate_impl: PUT")
	//fmt.Println("CredentialsPassword_UserName_validate_impl: iValue", iValue)
	//fmt.Printf("iRec: %v\n", iRec)
	if iValue == "" {
		fP = SetFieldError(fP, "user identity is required")
		return iValue, fP
	}
	if iValue != iRec.UserName {
		fP = SetFieldError(fP, "user identity does not match")
		return iValue, fP
	}
	if iValue != "" {
		//Check that the user 'name' is found in the database
		noFound, _, err := Credentials_GetByID(iValue)
		if err != nil {
			fP = SetFieldError(fP, "user identity not found")
			return iValue, fP
		}
		if noFound == 0 {
			fP = SetFieldError(fP, "user identity not found")
			return iValue, fP
		}
	}
	return iValue, fP
}

// ----------------------------------------------------------------
// CredentialsPassword_PasswordOld_validate_impl provides validation/actions for PasswordOld
func CredentialsPassword_PasswordOld_validate_impl(iAction string, iId string, iValue string, iRec dm.CredentialsPassword, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("CredentialsPassword", dm.CredentialsPassword_PasswordOld_scrn, VAL+"-"+iAction, iId)
	if iValue == "" {
		fP = SetFieldError(fP, "old password is required")
		return iValue, fP
	}
	return iValue, fP
}

// ----------------------------------------------------------------
// CredentialsPassword_PasswordNew_validate_impl provides validation/actions for PasswordNew
func CredentialsPassword_PasswordNew_validate_impl(iAction string, iId string, iValue string, iRec dm.CredentialsPassword, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("CredentialsPassword", dm.CredentialsPassword_PasswordNew_scrn, VAL+"-"+iAction, iId)
	if iValue == "" {
		fP = SetFieldError(fP, "new password is required")
		return iValue, fP
	}
	if iValue != "" && iValue != iRec.PasswordConfirm {
		fP = SetFieldError(fP, "new passwords do not match")
		return iValue, fP
	}

	iValue, fP = ComplexityCheck(iValue, fP)

	return iValue, fP
}

// ----------------------------------------------------------------
// CredentialsPassword_PasswordConfirm_validate_impl provides validation/actions for PasswordConfirm
func CredentialsPassword_PasswordConfirm_validate_impl(iAction string, iId string, iValue string, iRec dm.CredentialsPassword, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("CredentialsPassword", dm.CredentialsPassword_PasswordConfirm_scrn, VAL+"-"+iAction, iId)
	if iValue == "" {
		fP = SetFieldError(fP, "confirm password is required")
		return iValue, fP
	}
	if iValue != "" && iValue != iRec.PasswordNew {
		fP = SetFieldError(fP, "new passwords do not match")
	}

	iValue, fP = ComplexityCheck(iValue, fP)

	return iValue, fP
}

// ----------------------------------------------------------------
// Automatically generated code ends here
// ----------------------------------------------------------------

func ComplexityCheck(iValue string, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("CredentialsPassword", dm.CredentialsPassword_PasswordNew_scrn, "ComplexityCheck", "")
	//Check for password complexity
	if iValue != "" {
		//sampleEntropy := core.GetApplicationProperty("passwordEntropy")
		//entropy := passwordvalidator.GetEntropy(sampleEntropy)
		// entropy is a float64, representing the strength in base 2 (bits)

		minEntropyBits := core.StringToFloat(core.GetApplicationProperty("passwordMinEntropyBits"))
		err := passwordvalidator.Validate(iValue, minEntropyBits)
		// if the password has enough entropy, err is nil
		// otherwise, a formatted error message is provided explaining
		// how to increase the strength of the password
		// (safe to show to the client)
		if err != nil {
			fP = SetFieldError(fP, err.Error())
			return iValue, fP
		}
	}
	logs.Information("CredentialsPassword", "Password is valid")
	return iValue, fP
}
