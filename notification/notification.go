package notification

import (
	"fmt"
	"shiftsync/pkg/db"
	"shiftsync/pkg/helper/response"
	"time"
)

func SendNotification() {

	var resForgot []response.Forgot

	for {
		if err := db.GetDatabaseInstance().Raw("select phone from employees inner join attendances on employees.id = attendances.employee_id where extract(hour from (to_timestamp(punch_out, 'HH24:MI:SS') - to_timestamp(punch_in, 'HH24:MI:SS'))) > 8").Scan(&resForgot).Error; err != nil {
			fmt.Println(err)

		}

		fmt.Println("responce", resForgot)

		time.Sleep(30 * time.Second)
	}

}
