package notification

import (
	"fmt"
	"log"
	"shiftsync/pkg/config"
	"shiftsync/pkg/db"
	"shiftsync/pkg/helper/response"
	"time"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendNotification(cn config.Config) {
	fmt.Println("notification")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cn.Twilio.Account_sid,
		Password: cn.Twilio.Auth_token,
	})

	var resForgot []response.Forgot

	for {
		now := time.Now()

		time1 := time.Date(now.Year(), now.Month(), now.Day(), 15, 5, 0, 0, now.Location()) // 3:00 PM
		time2 := time.Date(now.Year(), now.Month(), now.Day(), 22, 5, 0, 0, now.Location()) // 10:00 PM
		time3 := time.Date(now.Year(), now.Month(), now.Day(), 5, 5, 0, 0, now.Location())  // 5:00 AM

		if now.Equal(time1) || now.Equal(time2) || now.Equal(time3) {
			if err := db.GetDatabaseInstance().Raw("select first_name as name, phone from employees inner join attendances on employees.id = attendances.employee_id where extract(hour from (to_timestamp(punch_out, 'HH24:MI:SS') - to_timestamp(punch_in, 'HH24:MI:SS'))) > 8").Scan(&resForgot).Error; err != nil {
				log.Fatal(err)
			}

			fmt.Println(resForgot)

			if len(resForgot) > 0 {
				for i := 0; i < len(resForgot); i++ {
					phone := fmt.Sprint(resForgot[i].Phone)
					params := &api.CreateMessageParams{}
					params.SetBody("Hi, " + resForgot[i].Name + " please remember to punch out for your shift as you forgot to do so. Thank you. ")
					params.SetFrom("+12543212748")
					params.SetTo("+91" + phone)

					resp, err := client.Api.CreateMessage(params)
					if err != nil {
						fmt.Println(err.Error())
					} else {
						if resp.Sid != nil {
							fmt.Println(*resp.Sid)
						} else {
							fmt.Println(resp.Sid)
						}
					}
				}
			}

		}

		time.Sleep(3 * time.Hour)
	}

}
