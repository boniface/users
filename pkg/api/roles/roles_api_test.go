package roles

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRolesRead(t *testing.T) {
	roles, err := GetRoles()

	if err!= nil{
		fmt.Println("This Executes with result ooooo", err.Error())

	}

	println(" The result is ",len(roles))

	assert.True(t, len(roles)>0)

}
func TestRoleRead(t *testing.T) {
	role, err := GetRole("23")
	if err!= nil{
		fmt.Println("This Executes with result ", err.Error())
	}

	fmt.Println(" The Role ",role)

}

func TestCreateRole(t *testing.T) {
	role:= Role{"1","23","TESTROLE","THIS A ROLE TEST "}
	res, err := CreateRole(role)
	if err!= nil{
		fmt.Println("This Executes with result ", err.Error())
	}

	assert.True(t,res)

}



