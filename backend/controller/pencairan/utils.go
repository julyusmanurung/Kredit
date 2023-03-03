package pencairan

import (
	"fmt"
	"github.com/julyusmanurung/Kredit/models"
	"strconv"
	"time"
)

func (r *repository) startInsertValidCustomer(customer models.StagingCustomer) {
	customerCode := r.generateCustomerCode(customer)

	r.insertCustomerToDataTab(customer, customerCode)
	r.insertCustomerToLoanDataTab(customer, customerCode)
	r.insertCustomerToVehicleDataTab(customer, customerCode)
}

func (r *repository) insertCustomerToDataTab(customer models.StagingCustomer, customerCode string) {
	convertedBirthDate, err := time.Parse("2006-01-02 15:04:05", customer.CustomerBirthDate)
	if err != nil {
		fmt.Println(err)
	}

	convertedIDType, err := strconv.ParseInt(customer.CustomerIDType, 10, 8)
	if err != nil {
		fmt.Println(err)
	}

	convertedLoanTglPk, err := time.Parse("2006-01-02", customer.LoanTglPk)
	if err != nil {
		fmt.Println(err)
	}

	convertedTglPkChanneling, err := time.Parse("2006-01-02", customer.LoanTglPkChanneling)
	if err != nil {
		fmt.Println(err)
	}
	verifiedCustomer := models.CustomerDataTab{
		Custcode:          customerCode,
		PPK:               customer.CustomerPpk,
		Name:              customer.CustomerName,
		Address1:          customer.CustomerAddress1,
		Address2:          customer.CustomerAddress2,
		City:              customer.CustomerCity,
		Zip:               customer.CustomerZip,
		BirthPlace:        customer.CustomerBirthPlace,
		BirthDate:         convertedBirthDate,
		IdType:            int8(convertedIDType),
		IdNumber:          customer.CustomerIDNumber,
		MobileNo:          customer.CustomerMobileNo,
		DrawdownDate:      convertedLoanTglPk,
		TglPkChanneling:   convertedTglPkChanneling,
		MotherMaidenName:  customer.CustomerMotherMaidenName,
		ChannelingCompany: customer.ScCompany,
		ApprovalStatus:    "9",
	}
	r.db.Create(&verifiedCustomer)
}

func (r *repository) insertCustomerToLoanDataTab(customer models.StagingCustomer, customerCode string) {
	convertedLoanInterestFlatChanneling, err := strconv.ParseFloat(customer.LoanInterestFlatChanneling, 32)
	if err != nil {
		fmt.Println(err)
	}

	convertedLoanInterestEffectiveChanneling, err := strconv.ParseFloat(customer.LoanInterestEffectiveChanneling, 32)
	if err != nil {
		fmt.Println(err)
	}

	convertedLoanEffectivePaymentType, err := strconv.ParseInt(customer.LoanEffectivePaymentType, 10, 8)
	if err != nil {
		fmt.Println(err)
	}

	convertedLoanOTR, err := strconv.ParseFloat(customer.LoanOtr, 32)
	if err != nil {
		fmt.Println(err)
	}

	convertedLoanDownPayment, err := strconv.ParseFloat(customer.LoanDownPayment, 32)
	if err != nil {
		fmt.Println(err)
	}

	convertedLoanAmountChanneling, err := strconv.ParseFloat(customer.LoanLoanAmountChanneling, 32)
	if err != nil {
		fmt.Println(err)
	}

	convertedLoanMonthlyPaymentChanneling, err := strconv.ParseFloat(customer.LoanMonthlyPaymentChanneling, 32)
	if err != nil {
		fmt.Println(err)
	}

	verifiedCustomer := models.LoanDataTab{
		Custcode:             customerCode,
		Branch:               customer.ScBranchCode,
		OTR:                  convertedLoanOTR,
		DownPayment:          convertedLoanDownPayment,
		LoanAmount:           convertedLoanAmountChanneling,
		LoanPeriod:           customer.LoanLoanPeriodChanneling,
		InterestType:         1,
		InterestFlat:         float32(convertedLoanInterestFlatChanneling),
		InterestEffective:    float32(convertedLoanInterestEffectiveChanneling),
		EffectivePaymentType: int8(convertedLoanEffectivePaymentType),
		AdminFee:             30,
		MonthlyPayment:       convertedLoanMonthlyPaymentChanneling,
		InputDate:            customer.ScCreateDate,
		LastModified:         time.Now(),
		ModifiedBy:           "system",
		InputDate2:           customer.ScCreateDate,
		InputBy:              "system",
		LastModified2:        time.Now(),
		ModifiedBy2:          "system",
	}
	r.db.Create(&verifiedCustomer)
}

func (r *repository) insertCustomerToVehicleDataTab(customer models.StagingCustomer, customerCode string) {
	convertedVehicleType, err := strconv.ParseInt(customer.VehicleType, 10, 8)
	if err != nil {
		fmt.Println(err)
	}

	convertedVehicleStatus, err := strconv.ParseInt(customer.VehicleStatus, 10, 8)
	if err != nil {
		fmt.Println(err)
	}

	convertedVehicleDealerID := 0
	if customer.VehicleDealerID != "" {
		vehicleDealerID, err := strconv.ParseInt(customer.VehicleDealerID, 10, 8)
		if err != nil {
			fmt.Println(err)
		}
		convertedVehicleDealerID = int(vehicleDealerID)
	}

	convertedVehicleTglSTNK, err := time.Parse("2006-01-02 15:04:05", customer.VehicleTglStnk)
	if err != nil {
		fmt.Println(err)
	}

	convertedVehicleTglBPKP, err := time.Parse("2006-01-02 15:04:05", customer.VehicleTglStnk)
	if err != nil {
		fmt.Println(err)
	}

	convertedCollateralTypeID, err := strconv.ParseInt(customer.CollateralTypeID, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	verifiedCustomer := models.VehicleDataTab{
		Custcode:       customerCode,
		Brand:          int(convertedVehicleType),
		Type:           customer.VehicleBrand,
		Year:           customer.VehicleYear,
		Golongan:       1,
		Jenis:          customer.VehicleJenis,
		Status:         int8(convertedVehicleStatus),
		Color:          customer.VehicleColor,
		PoliceNo:       customer.VehiclePoliceNo,
		EngineNo:       customer.VehicleEngineNo,
		ChasisNo:       customer.VehicleChasisNo,
		Bpkb:           customer.VehicleBpkb,
		RegisterNo:     "",
		Stnk:           customer.VehicleStnk,
		StnkAddress1:   "",
		StnkAddress2:   "",
		StnkCity:       "",
		DealerID:       int(convertedVehicleDealerID),
		Inputdate:      time.Now(),
		Inputby:        "system",
		Lastmodified:   time.Now(),
		Modifiedby:     "system",
		TglStnk:        convertedVehicleTglSTNK,
		TglBpkb:        convertedVehicleTglBPKP,
		TglPolis:       time.Now(),
		PolisNo:        "",
		CollateralID:   convertedCollateralTypeID,
		Ketagunan:      "",
		AgunanLbu:      "",
		Dealer:         customer.VehicleDealer,
		AddressDealer1: "",
		AddressDealer2: "",
		CityDealer:     customer.VehicleCityDealer}
	r.db.Create(&verifiedCustomer)
}

func (r *repository) insertCustomerToStagingError(customer models.StagingCustomer, unverifiedReason string) {
	unVerifiedCustomer := models.StagingError{
		Id:           customer.ID,
		SeReff:       customer.ScReff,
		SeCreateDate: customer.ScCreateDate,
		BranchCode:   customer.ScBranchCode,
		Company:      customer.ScCompany,
		Ppk:          customer.CustomerPpk,
		Name:         customer.CustomerName,
		ErrorDesc:    unverifiedReason}
	r.db.Create(&unVerifiedCustomer)
}

func (r *repository) updateCustomerScFlag(customer models.StagingCustomer, isValid bool) {
	var stagingCustomer []models.StagingCustomer
	if isValid {
		r.db.Model(&stagingCustomer).Where("id=?", customer.ID).Update("sc_flag", "1")
	} else {
		r.db.Model(&stagingCustomer).Where("id=?", customer.ID).Update("sc_flag", "8")
	}
}

func (r *repository) generateCustomerCode(customer models.StagingCustomer) string {
	type Company struct {
		company_code string
	}

	var idTab models.IdTab
	var companyCode Company
	NewCustomerCode := ""

	resIDTab := r.db.First(&idTab)
	if resIDTab.Error != nil {
		fmt.Println("error get data from id_table", resIDTab.Error)
		return "error"
	}

	resCompanyCode := r.db.Table("mst_company_tab").Where("company_short_name = ? ", customer.ScCompany).Scan(&companyCode)
	if resCompanyCode.Error != nil {
		fmt.Println("error get data from mst_company_tab", resCompanyCode.Error)
		return "error"
	}

	now := time.Now()
	currentMonth := convertMonth(int(now.Month()))
	currentYear := fmt.Sprintf("%d", int(now.Year()))

	AppCustomerCode := idTab.CODE
	AppCompanyCode := companyCode
	CustomerCodeSeq := idTab.VALUE
	CustomerCodeLen := idTab.DIGIT

	appCustomerCodeSeq := "0000000000" + fmt.Sprintf("%d", CustomerCodeSeq)
	appCustomerCodeSeq = appCustomerCodeSeq[len(appCustomerCodeSeq)-CustomerCodeLen:]

	NewCustomerCode = AppCustomerCode + AppCompanyCode.company_code + currentYear + currentMonth + appCustomerCodeSeq

	// update value column
	r.db.Model(&models.IdTab{}).Where("code = ?", idTab.CODE).Update("value", idTab.VALUE+1)

	return NewCustomerCode
}

func convertMonth(month int) string {
	if month < 10 {
		return "0" + fmt.Sprintf("%d", month)
	} else {
		return fmt.Sprintf("%d", month)
	}
}
