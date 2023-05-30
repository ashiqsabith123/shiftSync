package pdf

import (
	"fmt"
	"shiftsync/pkg/helper/response"
	"strconv"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func CreatePdf(data response.SalarySlip) error {

	month := time.Now().Month().String()
	cyear := time.Now().Year()

	year := strconv.Itoa(cyear)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetDrawColor(0, 0, 0)
	pdf.AddPage()

	pageWidth, pageHeight := pdf.GetPageSize()

	pageWidth = pageWidth - 16
	pageHeight = pageHeight - 16

	pdf.SetLineWidth(.6)
	pdf.Rect(8, 8, pageWidth, pageHeight, "D")

	pdf.AddUTF8Font("Josefin", "", "pdf/fonts/Kanit-Regular.ttf")

	pdf.SetFont("Josefin", "", 20)
	pdf.Text(pageWidth/2-4, 19, "shiftSync")

	pdf.SetFontSize(10)
	pdf.Text(pageWidth/2-4, 24, "www.shiftSync.org")

	pdf.SetFontSize(18)
	pdf.Text(18, 45, "Salary Slip")

	pdf.SetFontSize(14)
	pdf.Text(18, 51, month+"-"+year)

	pdf.Text(18, 71, "Employee Id: "+data.Employee_id)
	pdf.Text(18, 78, "Employee Name: "+data.Name)
	pdf.Text(18, 85, "Designation: "+data.Designation)
	pdf.Text(18, 100, "Account Number: "+data.Account_no)

	pdf.Text(pageWidth-34, 71, "Grade: "+data.Grade+" Level")
	pdf.Text(pageWidth-38, 78, "Worked Days: "+data.Duties)
	pdf.Text(pageWidth-22, 85, "Leaves: "+data.Leave_count)

	pdf.SetFontSize(20)
	pdf.Text(pageWidth/2-4, 120, "Pay Details")

	pdf.SetFillColor(170, 186, 217) // Set cell fill color
	// Set text color
	pdf.SetFont("Josefin", "", 13) // Set font style and size

	// Define table headers
	headers := []string{"Earnings", "Amount"}
	rows := []string{"Basic Pay", "DA", "SPA", "Leave Pay", "MA", "Over Time"}
	datas := []string{data.Base_salary, data.D_allowance, data.Sp_allowance, data.Leave_pay, data.M_allowance, data.Over_time}

	columnWidths := []float64{60, 60}

	pdf.SetXY(46.5, 128)
	for i, header := range headers {
		pdf.CellFormat(columnWidths[i], 7, header, "1B", 0, "C", true, 0, "")
	}

	pdf.SetFillColor(186, 184, 215)

	pdf.SetXY(46.5, 135.5)
	for _, row := range rows {
		pdf.CellFormat(65, 7, row, "1B", 2, "C", true, 0, "")
	}

	pdf.SetXY(101.2, 135.5)
	for _, data := range datas {
		pdf.CellFormat(65, 7, data, "1B", 2, "C", true, 0, "")
	}

	pdf.SetFontSize(16)
	pdf.Text(120, 188, "Earnings: "+data.Gross_salary)

	dheaders := []string{"Deductions", "Amount"}
	drows := []string{"Professional Tax", "Provident Fund"}
	ddatas := []string{data.Tax, data.Provident_fund}

	pdf.SetXY(46.5, 197)
	pdf.SetFillColor(170, 186, 217)

	pdf.SetFont("Josefin", "", 13)
	for i, header := range dheaders {
		pdf.CellFormat(columnWidths[i], 7, header, "1B", 0, "C", true, 0, "")
	}

	pdf.SetFillColor(186, 184, 215)

	pdf.SetXY(46.5, 204.3)
	for _, row := range drows {
		pdf.CellFormat(65, 7, row, "1B", 2, "C", true, 0, "")
	}

	pdf.SetXY(101.2, 204.3)
	for _, data := range ddatas {
		pdf.CellFormat(65, 7, data, "1B", 2, "C", true, 0, "")
	}

	pdf.SetFontSize(16)
	pdf.Text(115, 230, "Deductions: "+data.Deductions)

	pdf.SetFont("Josefin", "", 16)
	pdf.Text(140, pageHeight-22, "Total: "+data.Net_salary+" \u20B9")

	pdf.SetFontSize(16)
	pdf.Text(130, 250, "Earnings - Deductions ")

	pdf.SetLineWidth(.5)
	pdf.Line(130.0, 252.5, 186.0, 252.5)

	pdf.SetFontSize(10)
	pdf.Text(pageWidth/2-37, pageHeight+5, "This is a system generated slip, For any queries contact admin")

	err := pdf.OutputFileAndClose("pdf/generated/salary_slip" + data.Employee_id + ".pdf")
	if err != nil {
		fmt.Println("Error creating PDF:", err)
		return err
	}

	fmt.Println("PDF with page border created successfully.")

	return nil
}
