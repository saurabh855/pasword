package main

import (
	"encoding/json"
	"fmt"
	"os"
)












type Generated struct {
	Email     string          `json:"email"`
	EmailCc   string          `json:"email_cc"`
	EmailBcc  string          `json:"email_bcc"`
	FromEmail string          `json:"from_email"`
	FromName  string          `json:"from_name"`
	Subject   string          `json:"subject"`
	ItemList  [][]interface{} `json:"item_list"`

	OrderTotal       string `json:"order_total"`
	PaymentLink      string `json:"payment_link"`
	EmailAttachment  string `json:"email_attachment"`
	Sendtime         string `json:"sendtime"`
	BatchID          string `json:"batch_id"`
	Status           string `json:"status"`
	Isactive         string `json:"isactive"`
	UniqueCode       string `json:"unique_code"`
	BookingID        string `json:"booking_id"`
	ServiceBookingID string `json:"service_booking_id"`
}

func main() {

	ab := `  {"email":"hemant.singh@healthians.com","email_cc":"","email_bcc":"","from_email":"shop@healthians.com","from_name":"Healthians","subject":"Hrebveda Payment Link","item_list":[["Powerup",177],["Saunfs",59.83]],"order_total":"236.83","payment_link":"http:\/\/172.19.5.180\/omniorder\/payment-link\/HLC61A4AA6831F88","email_attachment":"","sendtime":"2021-11-29 10:24:44","batch_id":"","status":"1","isactive":"1","unique_code":"shop_payment_link","booking_id":"HLC61A4AA6831F88","service_booking_id":""}`
	var body Generated
	fmt.Printf(ab)
	bodyErr := json.Unmarshal([]byte(ab), &body)

	if bodyErr != nil {
		fmt.Errorf("Unmarshal Error: %s", bodyErr.Error())

	}
	fmt.Println("ffg", body.OrderTotal)
	fmt.Println(body)
	fmt.Println("price:", body.ItemList[0][1])
	var productname interface{}
	var productprice interface{}
	grandtotal := body.OrderTotal
	fmt.Println("ttt", body.OrderTotal)
	var twotemp string
	//var onetemp string

	for i := 0; i < len(body.ItemList); i++ {
		productname = body.ItemList[i][0]
		str := fmt.Sprintf("%v", productname)

		productprice = body.ItemList[i][1]
		stri := fmt.Sprintf("%v", productprice)
		twotemp += "<table width=\"98%\" align=\"center\" cellspacing=\"0\" cellpadding=\"0\" align=\"center\" style=\"margin-top:0px;\"><tr><td valign=\"top\" width=\"100%\" style=\"padding-left:15px;\"><table cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" border=\"0\"><tr><td><span style=\"display:block;padding-bottom:5px;padding-top:5px;\"><p style=\"margin:0px;padding:0px;font-size:1rem;color:#000;line-height:19px;\">" + str + "<br></p></span></td><td><span style=\"display:block;padding-bottom:5px;padding-top:5px;text-align:right;\"><p style=\"margin:0px;padding:0px;font-size:0.8rem;;;color:#000;line-height:19px;\"><span style=\"font-size:20px;\">Rs." + stri + "</span></p></span></td></tr></table></td></tr></table>"

	}
	var mail string
	mail += "<body style=\"margin:0px;padding:0px;background:#eaeaea;font-family: sans-serif;\">"
	mail += "<center> <table width=\"600\" bgcolor=\"#ccc\" align=\"center\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" id=\"bodyTable\" height=\"100%\" style=\"min-width:600px; table-layout: fixed;max-width:100% !important;background-color:#fff;\"> <tr> <td valign=\"top\"><img src=\" https://helma.healthians.com/stationery/mailer-assets/614833ecc043b.jpg\"></td> </tr> <tr> <td align=\"center\" valign=\"top\" style=\"width:100%;\"> <table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" style=\"border-collapse: separate;mso-table-lspace: 0pt; mso-table-rspace:0pt;width: 100%;background:#007D83;\"> <tr> <td> <table width=\"92%\" cellspacing=\"0\" cellpadding=\"0\" align=\"center\"> <tr> "
	mail += "<td valign=\"top\"> <h2 style=\"margin:0px;padding:40px 0px 20px 0px;color:#fff;font-family: sans-serif;font-size:1.5em;\"> Thank you for choosing Herbveda+, the all-natural road to a healthier you. </h2>"
	mail += "<p style=\"margin:0px;padding-top:6px;color:#fff;font-family: sans-serif;font-size:1em;line-height:20px;\">We re delighted to have you on board. Use Pay Now   to proceed with the payment for the purchase of the below products. </p> </td> </tr> </table> </td> </tr> <tr> <td> <table width=\"92%\" cellspacing=\"0\" cellpadding=\"0\" align=\"center\"> <tr> <td valign=\"top\"> <p style=\"margin:0px;padding-top:5px;padding-bottom:5px;color:#dcfdff;font-family: sans-serif;font-size:1em;line-height:15px;\">&nbsp; </p> </td> </tr> </table> </td> </tr> <tr> <td valign=\"top\"> <table width=\"92%\" cellspacing=\"0\" cellpadding=\"0\" align=\"center\"> <tr> <td bgcolor=\"#fff\" valign=\"top\" style=\"background:#fff;border-radius:3px;box-shadow:0px 2px 3px #cecece;height:auto;padding-bottom:10px;\"> "
	mail += "<table width=\"100%\" cellspacing=\"0\" cellpadding=\"0\" align=\"center\" style=\"margin-top:10px;\"> <tr> <td> <table width=\"98%\" cellspacing=\"0\" cellpadding=\"0\" align=\"center\"> <tr> <td valign=\"top\"> <table width=\"100%\" cellspacing=\"0\" cellpadding=\"0\" align=\"center\"> <tr> <td> <h3 style=\"padding-top:15px;padding-left:15px;text-transform:none;font-size:16px;color:#000;text-transform: uppercase;font-weight:600;\">Product Summary</h3> </td> </tr> </table>"
	mail += " " + twotemp + " <!-- item list end --> <!--Grand total section start--> <table style=\"padding-top:10px;border-top:1px dotted #eaeaea;padding-bottom:10px;\" width=\"96%\" align=\"center\" cellspacing=\"0\" cellpadding=\"0\"> <tr> <td style=\"font-weight:bold;\" valign=\"top\" style=\"padding-left:10px;\"> Total </td> <td style=\"text-align:right;font-weight:bold;\">Rs. " + grandtotal + "</td> </tr> </table> <!--Grand total section end--> </td> </tr> </table> </td> </tr> </table> <!-- --> </td> </tr> </table> </td> </tr> </table> </td> </tr> <tr> <td align=\"center\" valign=\"top\" style=\"width:100%;\"> <p style=\"margin:15px 20px; padding-top: 6px; color: #000; font-family: sans-serif; font-size: 1rem; line-height: 20px;\">Please call us at <customer care number> for any payment-related support and we  ll be happy to address your questions and queries. </p> <div style=\"display:inline-block;background:#F27D27;border-radius:30px;color:#fff;font-weight:bold;\"><a href=\" " + body.PaymentLink + "\" style=\"color:#fff;padding:12px 36px;display:block;text-decoration:none;\">Pay now</a></div> <br><br> </td> </tr> <tr> <td style=\"background:#f4f4f4;padding:0px 0px;text-align:center;\"> <p style=\"font-size:0.9rem;text-align:center;padding-bottom:0px;padding-top:15px;\">Follow us on</p> <span style=\"display:inline-block;margin-right:10px;margin-bottom:20px;\">"
	mail += "<a href=\"https://www.facebook.com/Healthians?ref=hl\" target=\"_blank\"><img src=\"https://helma.healthians.com/stationery/mailer-assets/6141c5c6c0987.png\" width=\"36\"></a></span> <span style=\"display:inline-block;margin-right:10px;margin-bottom:20px;\"><a href=\"https://twitter.com/healthians\" target=\"_blank\"><img src=\"https://helma.healthians.com/stationery/mailer-assets/6141c61944de3.png\" width=\"36\"></a></span> <span style=\"display:inline-block;margin-right:10px;margin-bottom:20px;\"><a href=\"https://www.linkedin.com/company/healthians\" target=\"_blank\"><img src=\"https://helma.healthians.com/stationery/mailer-assets/6141c67237c16.png\" width=\"36\"></a></span> <span style=\"display:inline-block;margin-right:10px;margin-bottom:20px;\"><a href=\"https://www.youtube.com/user/Healthians\" target=\"_blank\"><img src=\"https://helma.healthians.com/stationery/mailer-assets/6141c6ebd8026.png\" width=\"36\"></a></span> <span style=\"display:inline-block;margin-right:10px;margin-bottom:20px;\"><a href=\"https://www.instagram.com/healthians/?utm_source=ig_profile_share&amp;igshid=qc3h7c2udi3n\" target=\"_blank\"><img src=\"https://helma.healthians.com/stationery/mailer-assets/6141c63bb00eb.png\" width=\"36\"></a></span> </td> </tr> <tr> <td style=\"background:#fff;border-top:2px solid #FFAE41;padding:4px 0px;\"> <p style=\"padding:5px 15px;margin:0px;color:#333;font-size:9px;font-family: Arial, Helvetica, sans-serif;line-height:11px;text-align:justify;\">Healthians is India s leading health test at home service offering a wide range of health tests across 80 cities of India. We have a network of state-of-the-art fully automated laboratories and a large team of highly skilled phlebotomists who specialize in sample collection from homes.</p> </td> </tr> </table> </center> </body>"
	f, err := os.Create("hulu.html")
	fmt.Println(err)

	fmt.Println("Temp file name:", f.Name())

	_, err = f.Write([]byte(mail))
	fmt.Println(err)

}
