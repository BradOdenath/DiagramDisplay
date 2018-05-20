package main

import (
	"fmt"
	"log"
)

const (
	THENAME string = "Bradley Michael Odenath"
	NIL_STRING = ""
	INDENT_SIZE uint64 = 2
)

var layerID uint64 = 0
var displayIDs[] uint64

func main() {
	head := CreateBubble("Start Application", "Service: Startup")

	//FOCUS: Login

	loginPage := CreateBubble("IO: Username & Password or Sign Up", "Login Page")
	head.AppendPossibility(loginPage)

	loginSuccess := CreateBubble("Login Success!", "Service: Login")
	loginPage.AppendPossibility(loginSuccess)

	loginShitty := CreateBubble("Login Failed", "Retry: Login")
	loginPage.AppendPossibility(loginShitty)
	loginShitty.AppendPossibility(loginPage)

	//FOCUS: Sign-Up

	signUpPage := CreateBubble("Enter Details", "Page: Sign-Up")

	signUpSuccess := CreateBubble("Sign-Up Success!", "Service: Sign-Up")
	signUpPage.AppendPossibility(signUpSuccess)

	signUpShitty := CreateBubble("Sign-Up Failure!", "Retry: Sign-Up")
	signUpPage.AppendPossibility(signUpShitty)
	signUpShitty.AppendPossibility(signUpPage)

	verifyEmail := CreateBubble("Verify E-Mail To Continue!", "Service: Verify")
	signUpSuccess.AppendPossibility(verifyEmail)

	//FOCUS: Forgot Password

	forgotPWPage := CreateBubble("Enter E-Mail", "Page: Forgot Password!")
	loginPage.AppendPossibility(forgotPWPage)

	forgotPWSuccess := CreateBubble("E-Mail Sent!", "Service: Password Reset")
	forgotPWPage.AppendPossibility(forgotPWSuccess)

	forgotPWShitty := CreateBubble("Retry E-Mail!", "Retry: Email")
	forgotPWPage.AppendPossibility(forgotPWShitty)
	forgotPWShitty.AppendPossibility(forgotPWPage)

	setupPage := CreateBubble("Goto Parent / Lock phone as Child", "Page: Setup")
	loginSuccess.AppendPossibility(setupPage)

	managementPage := CreateBubble("Manage Children / ML Play", "Page: Parent View")
	setupPage.AppendPossibility(managementPage)

	childrenListPage := CreateBubble("Choose a Child", "Page: Children")
	managementPage.AppendPossibility(childrenListPage)

	machinePage := CreateBubble("Enter Phrase for Machine", "Page: Machine")
	managementPage.AppendPossibility(machinePage)

	lockPhonePage := CreateBubble("Enter a master password", "Page: Lock-Live as Child")
	setupPage.AppendPossibility(lockPhonePage)

	lockPhoneSuccess := CreateBubble("Locking Phone!", "Service: Lock-Live as Child")
	lockPhonePage.AppendPossibility(lockPhoneSuccess)

	childPage := CreateBubble("No functionality for you!", "Page: Child View")
	lockPhoneSuccess.AppendPossibility(childPage)

	monitoringService := CreateBubble("Monitoring Child", "Service: SMS Snooping")
	childPage.AppendPossibility(monitoringService)

	parentNotification := CreateBubble("Something wrong bro!", "Service: Notification/Alert")
	monitoringService.AppendPossibility(parentNotification)
	head.AppendPossibility(parentNotification)

	//WOAW, LOG IT PRINTS OUT OF ORDER... FANCY!!
	//log.Println("<DiagramV1>")
	head.RepresentBubble(0)
	//log.Println("</DiagramV1>")
}

//Representation of linkage between bubbles.
type BubbleReference struct {
	children[] Bubble
	//parent Bubble
}

//BubbleReference: Add a child node.
func (br *Bubble) AppendPossibility(b Bubble) {
	br.connection.children = append(br.connection.children, b)
	log.Println(br.connection.children)
	//b.connection.parent = br
}

//Representation of the machine's state.
type Bubble struct {
	id uint64
	blurb string
	statement string
	connection BubbleReference
}

func CreateBubble(blur string, statement string) Bubble {
	layerID += 1 //Append Unique ID
	return Bubble{layerID, blur, statement, BubbleReference{}}
}

func (b Bubble) RepresentBubble(ind uint64) {
	if !(b.IsDisplayed()) {
		fmt.Println(Indent(ind), b.blurb, "-", b.statement)
		//fmt.Println(Indent(ind*2), ":Directions:")

		//Append the ID, probably should be it's own function.
		displayIDs = append(displayIDs, b.id)

		//Recursively print the rest ignoring the ones that exist.
		if b.HasPossibilities() {
			for i, c := range b.connection.children {
				c.connection.children[i].RepresentBubble(ind+INDENT_SIZE)
			}
		}
	}
}

//Multiple children would indicate a condition
func (b Bubble) IsConditionalBubble() bool {
	return len(b.connection.children) > 1
}

func (b Bubble) HasPossibilities() bool {
	return len(b.connection.children) <= 0
}

func (b Bubble) IsDisplayed() bool {
	set := make(map[uint64]struct{}, len(displayIDs))
	for _, s := range displayIDs {
		set[s] = struct {}{}
	}
	_, ok := set[b.id]
	return ok
}

func Indent(size uint64) string {
	str := ""
	for i := 0; i <= int(size); i++ {
		str += " "
	}
	return str
}