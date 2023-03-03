package pencairan

import (
	"github.com/julyusmanurung/Kredit/models"
	"regexp"
	"time"
)

func (r *repository) startValidation(customer models.StagingCustomer) {
	r.validateCustomerPPK(customer)
}

func (r *repository) validateCustomerPPK(customer models.StagingCustomer) {
	var recordFound int64
	r.db.Table("customer_data_tab").Where("ppk = ?", customer.CustomerPpk).Count(&recordFound)

	if recordFound == 0 {
		r.validateScCompany(customer)
	} else {
		r.updateCustomerScFlag(customer, false)
		r.insertCustomerToStagingError(customer, "customer ppk is not valid")
	}
}

func (r *repository) validateScCompany(customer models.StagingCustomer) {
	var recordFound int64
	r.db.Table("mst_company_tab").Where("company_short_name = ?", customer.ScCompany).Count(&recordFound)

	if recordFound == 0 {
		r.updateCustomerScFlag(customer, false)
		r.insertCustomerToStagingError(customer, "sc company is not valid")
	} else {
		r.validateScBranchCode(customer)
	}
}

func (r *repository) validateScBranchCode(customer models.StagingCustomer) {
	var recordFound int64
	r.db.Table("branch_tab").Where("code = ?", customer.ScBranchCode).Count(&recordFound)

	if recordFound == 0 {
		r.updateCustomerScFlag(customer, false)
		r.insertCustomerToStagingError(customer, "sc branch code is not valid")
	} else {
		r.validateLoanTglPk(customer)
	}
}

func (r *repository) validateLoanTglPk(customer models.StagingCustomer) {
	currentTime := time.Now()
	currentMonth := int(currentTime.Month())

	date, err := time.Parse("2006-01-02", customer.LoanTglPk)
	if err != nil {
		panic(err)
	}

	if currentMonth == int(date.Month()) {
		r.validateCustomerIDNumber(customer)
	} else {
		r.updateCustomerScFlag(customer, false)
		r.insertCustomerToStagingError(customer, "loan tgl pk is not valid")
	}
}

func (r *repository) validateCustomerIDNumber(customer models.StagingCustomer) {
	if customer.CustomerIDType == "1" {
		if len(customer.CustomerIDNumber) == 0 {
			r.updateCustomerScFlag(customer, false)
			r.insertCustomerToStagingError(customer, "customer id number is not valid")
		} else {
			r.validateCustomerName(customer)
		}
	} else {
		r.updateCustomerScFlag(customer, false)
		r.insertCustomerToStagingError(customer, "customer id number is not valid")
	}
}

func (r *repository) validateCustomerName(customer models.StagingCustomer) {
	regex := regexp.MustCompile(`/[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]/`)
	if matched := regex.MatchString(customer.CustomerName); !matched {
		r.validateVehicleBPKB(customer)
	} else {
		r.insertCustomerToStagingError(customer, "customer name is not valid")
	}
}

func (r *repository) validateVehicleBPKB(customer models.StagingCustomer) {
	if len(customer.VehicleBpkb) == 0 {
		r.updateCustomerScFlag(customer, false)
		r.insertCustomerToStagingError(customer, "vehicle bpkb is not valid")
	} else {
		r.validateVehicleSTNK(customer)
	}
}

func (r *repository) validateVehicleSTNK(customer models.StagingCustomer) {
	if len(customer.VehicleStnk) == 0 {
		r.updateCustomerScFlag(customer, false)
		r.insertCustomerToStagingError(customer, "validate vehicle stnk is not valid")
	} else {
		r.validateVehicleEngineNo(customer)
	}
}

func (r *repository) validateVehicleEngineNo(customer models.StagingCustomer) {
	if len(customer.VehicleEngineNo) == 0 {
		r.updateCustomerScFlag(customer, false)
		r.insertCustomerToStagingError(customer, "validate vehicle engine number is not valid")
	} else {
		var recordFound int64
		r.db.Table("vehicle_data_tab").Where("engine_no = ?", customer.VehicleEngineNo).Count(&recordFound)

		if recordFound == 0 {
			r.validateVehicleChasisNo(customer)
		} else {
			r.updateCustomerScFlag(customer, false)
			r.insertCustomerToStagingError(customer, "validate vehicle engine number is not valid")
		}
	}
}

func (r *repository) validateVehicleChasisNo(customer models.StagingCustomer) {
	if len(customer.VehicleChasisNo) == 0 {
		r.insertCustomerToStagingError(customer, "validate vehicle chasis number is not valid")
	} else {
		var recordFound int64
		r.db.Table("vehicle_data_tab").Where("chasis_no = ?", customer.VehicleChasisNo).Count(&recordFound)

		if recordFound == 0 {
			r.updateCustomerScFlag(customer, true)
			r.startInsertValidCustomer(customer)
		} else {
			r.updateCustomerScFlag(customer, false)
			r.insertCustomerToStagingError(customer, "validate vehicle chasis number is not valid")
		}
	}
}
