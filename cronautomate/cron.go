package cronautomate

import (
	"fmt"
	"log"
	"math"
	"shiftsync/pkg/db"
	"shiftsync/pkg/encrypt"
	"shiftsync/pkg/helper"
	"shiftsync/pkg/helper/response"
	"shiftsync/razorpay"
	"time"

	"github.com/robfig/cron"
)

func AutomateCreditSalary() {
	fmt.Println("automate")

	automate := cron.New()
	automate.AddFunc("0 0 0 1 *", CreditSalary)

	automate.Start()

	//select {}
}

func CreditSalary() {

	IDs, idErr := getAllIds()

	if idErr != nil || len(IDs) == 0 {
		return
	}

	for _, v := range IDs {
		currentMonth := time.Now().Local().Format("2006-01-02")

		var totalAmount float32

		hours, calulateError := calculateTotalWorkingHours(v, currentMonth)
		if calulateError != nil {
			continue
		}

		duties := hours / 8

		grade, gradeError := getGradeOfTheEmployee(v)
		if gradeError != nil {
			continue
		}

		switch grade {
		case "A":
			totalAmount = hours * 150.0
		case "B":
			totalAmount = hours * 125.0
		case "C":
			totalAmount = hours * 100.0
		}

		allowance, allowanceError := addAllAlowances(v)
		if allowanceError != nil {
			continue
		}

		var getCount response.LeaveCount

		getCount.Id = v
		getCount.Date = int(time.Now().Month())

		paidLeave, leaveError := getCountOfPaidLeave(getCount)
		if leaveError != nil {
			fmt.Println(leaveError)
		}

		leaveCount := paidLeave

		if paidLeave != 0 {
			paidLeave = paidLeave * 100
		}

		deductions, deductionError := caculateDeductions(v)
		if deductionError != nil {
			continue
		}

		grossSalary := totalAmount + allowance + float32(paidLeave)
		netSalary := grossSalary - deductions

		if err := updateFullSalary(v, paidLeave, leaveCount, int(duties), grossSalary, netSalary); err != nil {
			continue
		}

		if netSalary < 1 {
			continue
		}

		accDetails := fetchAccDetailsById(v)

		accDetails.ID = v

		accDetails.Account_no = string(encrypt.Decrypt(helper.Decode(accDetails.Account_no)))

		accDetails.Reference_id = helper.GenerateTransactionID()

		accDetails.Amount = math.Floor(float64(netSalary))

		if err := razorpay.CreatePayouts(accDetails); err != nil {
			continue
		}
		time.Sleep(3 * time.Second)

	}

}

func calculateTotalWorkingHours(id int, month string) (float32, error) {
	var hours float32
	if err := db.GetDatabaseInstance().Raw("SELECT CAST(SUM(EXTRACT(epoch FROM (TO_TIMESTAMP(punch_out, 'HH24:MI:SS') - TO_TIMESTAMP(punch_in, 'HH24:MI:SS')))) AS FLOAT) / 3600 AS hours FROM attendances WHERE employee_id = ? AND date_trunc('month', TO_DATE(date, 'YYYY-MM-DD')) = date_trunc('month', TO_DATE(?, 'YYYY-MM-DD'));", id, month).Scan(&hours).Error; err != nil {
		return 0, err
	}

	return hours, nil
}

func getGradeOfTheEmployee(id int) (string, error) {
	var grade string
	if err := db.GetDatabaseInstance().Raw("SELECT grade FROM salaries 	WHERE employee_id = ?", id).Scan(&grade).Error; err != nil {
		return "", err
	}

	return grade, nil
}

func addAllAlowances(id int) (float32, error) {
	var allowance float32
	if err := db.GetDatabaseInstance().Raw("SELECT (d_allowance + sp_allowance + m_allowance) AS allowance FROM salaries WHERE employee_id = ?", id).Scan(&allowance).Error; err != nil {
		return 0, err
	}

	return allowance, nil
}

func caculateDeductions(id int) (float32, error) {
	var deductions float32
	if err := db.GetDatabaseInstance().Raw("SELECT (tax + provident_fund ) AS deductions FROM salaries WHERE employee_id = ?", id).Scan(&deductions).Error; err != nil {
		return 0, err
	}

	return deductions, nil
}

func updateFullSalary(id, leave, leave_count, duties int, gross, net float32) error {
	if err := db.GetDatabaseInstance().Exec("UPDATE salaries SET gross_salary = ? ,net_salary = ? , leave_pay = ? , duties = ?, leave_count = ? WHERE employee_id = ?", gross, net, leave, duties, leave_count, id).Error; err != nil {
		return err
	}
	return nil
}

func fetchAccDetailsById(id int) response.AccDetails {
	var details response.AccDetails
	if err := db.GetDatabaseInstance().Raw("SELECT forms.account_no, razorpays.contact_id FROM forms INNER JOIN razorpays ON forms.form_id = razorpays.id WHERE forms.form_id = ?;", id).Scan(&details).Error; err != nil {
		log.Fatal(err)
	}

	return details
}

func getCountOfPaidLeave(reqCount response.LeaveCount) (int, error) {
	var count int

	if err := db.GetDatabaseInstance().Raw("SELECT SUM(DATE_PART('day', leaves.to::timestamp - leaves.from::timestamp)) AS count FROM leaves WHERE employee_id = ? AND mode = 'P' AND status = 'A' AND EXTRACT(MONTH FROM to_date(leaves.from, 'DD-MM-YYYY')) =?;", reqCount.Id, reqCount.Date).Scan(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func getAllIds() ([]int, error) {
	var id []int
	if err := db.GetDatabaseInstance().Raw("SELECT form_id FROM FORMS WHERE status = 'A'").Scan(&id).Error; err != nil {
		return id, err
	}

	return id, nil
}
