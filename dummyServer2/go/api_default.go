package swagger

 import (
 
	 "net/http"
	 "encoding/json"
	 "bytes"
	 "fmt"
	 "io/ioutil"
	 "github.com/buger/jsonparser"
	 //"github.com/gin-gonic/gin"
 
 );


func queryForTimeoff(user string, team string) {

 query :=
    `{"query":"query queryForTimeoff {` +
                `usermanagementdb_organization_employee(distinct_on: slackId, where:` +
                    `{slackId: {_eq: \"U03D6Q82139\"}})` +
                        `{employeeId
                    isAccountLinked
                    organizationId
                    role
                    slackId}` +
                `usermanagementdb_Organization(where:` +
                    `{organization_employees: {workspaceId: {_eq: \"T03E01NGG81\"}}})` +
                        `{slackChannel}` +
                `timeoffdb_LeaveType(where:` +
                    `{Organization: {organization_employees: {workspaceId: {_eq: \"T03E01NGG81\"}}}})` +
                        `{leaveTypeName
                    organizationId
                    leaveTypeEmoticon}` +
                `usermanagementdb_SlackBot(where:` +
                    `{workspaceId: {_eq: \"T03E01NGG81\"}})` +
                    `{botToken
                    workspaceId}` +
                `timeoffdb_GeneralSettings(where:` +
                    `{Organization: {organization_employees: {workspaceId: {_eq: \"T03E01NGG81\"}}}})` +
                    `{workingHours}
		}"
    }`


	r, err := http.NewRequest("POST",  "https://0c4d-2405-201-6815-f820-221-86ff-fe2a-9287.in.ngrok.io/v1/graphql", bytes.NewBuffer([]byte(query)))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("x-hasura-admin-secret", "5035f53ab015ff91b54d8a4063451107dc55aefcdc1bfdcc4315fe33d9f47b69")

	client1 := &http.Client{}
	res1, err := client1.Do(r)



	if err != nil {
		panic(err)
	}

    data, _ := ioutil.ReadAll(res1.Body)

	data1 := string(data)
	fmt.Println("response", data1)

	//  text := jsonparser.Get(data, "usermanagementdb_organization_employee", "employeeId")
	//  fmt.Println("text", text)

	// jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
	// 	fmt.Println(jsonparser.Get(value, "role"))
	// }, "usermanagementdb_organization_employee", "role")

}

func displayHome(user string, team string) {
	// myToken := ""
	// myRole  := ""
	// isAccountLinked := false
	// leaveTypes := []
	// try {
        
        // fmt.println("Query result")

    //     leaveTypes := (dbResData, "data", "timeoffdb_LeaveType");

    //     if(dbResData.data.usermanagementdb_SlackBot.length > 0) {
    //         myToken = (dbResData, "data", "usermanagementdb_SlackBot[0]", "botToken")
    //         myRole = (dbResData, "data","usermanagementdb_organization_employee[0]", "role")
    //         isAccountLinked = (dbResData, "data", "usermanagementdb_organization_employee[0]", "isAccountLinked")
    //     }
    //     else {
    //         fmt.Println("some error in dbResData");
    //     }
    // } catch (er) {
    //     fmt.Println("Caught in error while fetching bot token from db: ")
    // }

	// let webClient := new WebClient(myToken)

    // if(isAccountLinked) {
    //     try {
    //          startDate := new Date()
    //          month := ("0" + (startDate.getMonth() + 1)).slice(-2)
    //         currentDate := startDate.getFullYear() + '-' + month + '-' + startDate.getDate()
    
    //          jsonData := await fetchTodaysTimeoffs(myRole, currentDate, team)
    //         fmt.Println("Query result")

    //          fetchedLeaves := jsonData.data.timeoffdb_Leaves
    //          displayResponse := await displayTimeoffs("timeoffFeature", fetchedLeaves, "Today", leaveTypes, webClient, user, team, myRole)
    //         fmt.Println("Display response in events: ")
    //     }
    //     catch (error) {
    //         fmt.Println("Error in display home function")
    //     }
    // }
    // else {
    //     fmt.Println("User account is not linked")


	var blockData = []byte(fmt.Sprintf(`
		{
			"type": "home",
			"blocks": [
				{
					"type": "section",
					"text": {
						"type": "mrkdwn",
						"text": "*Welcome home %s, :house:*"
					}
				},
				{
					"type": "section",
					"text": {
						"type": "mrkdwn",
						"text": "You need to Sign In to your TeamPlus account!"
					}
				}
			]
		}`, user))

			jsonparser.ArrayEach(blockData, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				fmt.Println(jsonparser.GetString(value, "text", "text"))
			},"blocks")
		// result := await webClient.views.publish(homeBlock);
       
}
			

func jsonEscape(i string) string {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	s := string(b)
	return s[1 : len(s)-1]
			
}
   
func AddLeave(w http.ResponseWriter, r *http.Request) {


	 data, _ := ioutil.ReadAll(r.Body)

	 if value, err := jsonparser.GetString(data, "type"); err == nil {
		 if value == "url_verification" {
			 fmt.Println("url verification received")
 
		 }
		
	}

	 if event, err := jsonparser.GetString(data, "event", "type"); err == nil {
			if event == "app_home_opened" {
				fmt.Println("App Home Opened");
			   displayHome("kavita", "aipan");
			   queryForTimeoff("kavita", "aipan")
				
			}
	 }

	 if event, err := jsonparser.GetString(data, "event", "type"); err == nil {
		if event == "app-uninstalled" {
			fmt.Println("App is uninstalled");
		}

	} else {
			eventType, err := jsonparser.GetString(data, "event", "type")
			fmt.Println("Some other event invoked ", eventType, err)
	}
	
	
	 w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	 w.WriteHeader(http.StatusOK)
	 w.Write([]byte("{\"a\":\"b\"}"))
 
}