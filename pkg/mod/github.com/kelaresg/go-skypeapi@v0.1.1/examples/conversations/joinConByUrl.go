package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"github.com/kelaresg/go-skypeapi"
	"strings"
)

/**
Retrieve the join URL for a group conversation, if it is currently public.
Join by code need 3 steps;
1. get conversation id by use JoinConByCode
2. add a member by use AddMember, the member id is user self id, and role value is "Admin", and the response will get code 207.
2. then add a member by use AddMember too, the member id is user self id, but role value is "User". and done
 */
func main5() {
	cli, err := skype.NewClient()
	if err != nil {
		fmt.Println(err)
	}
	skype.GetConfigYamlForBuildExample()
	username := viper.GetString("user.username")
	pwd := viper.GetString("user.password")
	err  = cli.Login(username, pwd)
	c := skype.ConversationsClient{}

	fmt.Printf("\niuput url and enter to join:")
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}
	inputArr := strings.Split(input, " ")
	joinUrl := inputArr[0]
	err, rsp := c.JoinConByCode(cli.LoginInfo.SkypeToken, cli.LoginInfo.RegistrationtokensStr, joinUrl)
	member1 := skype.Member{
		Id: "8:"+cli.UserProfile.Username,
		Role: "Admin",
	}
	Members := skype.Members{}
	Members.Members = append(Members.Members, member1)
	c.AddMember(cli.LoginInfo.LocationHost, cli.LoginInfo.SkypeToken, cli.LoginInfo.RegistrationtokensStr, Members, rsp.Resource)
	member2 := skype.Member{
		Id: "8:"+cli.UserProfile.Username,
		Role: "User",
	}
	mewMembers := skype.Members{}
	mewMembers.Members = append(mewMembers.Members, member2)
	c.AddMember(cli.LoginInfo.LocationHost, cli.LoginInfo.SkypeToken, cli.LoginInfo.RegistrationtokensStr, mewMembers, rsp.Resource)
	fmt.Println("-----------------------------end-------------------------------")
}
