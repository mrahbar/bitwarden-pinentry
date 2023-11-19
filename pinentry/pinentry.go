package pinentry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/foxcpp/go-assuan/common"
	"github.com/foxcpp/go-assuan/pinentry"
	"github.com/mrahbar/bitwarden-pinentry/bitwarden"
	"os/exec"
)

type BitwardenClient struct {
	Session string
	ItemId  string
	Auditor *Auditor
}

// GetPIN shows window with password textbox, Cancel and Ok buttons.
// Error is returned if Cancel is pressed.
func (c *BitwardenClient) GetPIN(settings pinentry.Settings) (string, *common.Error){
	c.Auditor.Println("GETPIN got called with settings:")
	c.Auditor.Printf("GETPIN: %+v\n", settings)
	c.Auditor.Println("GETPIN: executing bw get item")
	cmd := exec.Command("bw", "get", "item", c.ItemId, "--session", c.Session)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	var cmdErr *common.Error
	if err != nil {
		cmdErr = &common.Error{
			Src: common.ErrSrcUnknown, Code: common.ErrGeneral,
			SrcName: "system", Message: fmt.Sprint(err),
		}
	}

	var item bitwarden.ItemResponse
	if err := json.Unmarshal(out.Bytes(), &item); err != nil {
		cmdErr = &common.Error{
			Src: common.ErrSrcUnknown, Code: common.ErrGeneral,
			SrcName: "system", Message: fmt.Sprint(err),
		}
	}

	if cmdErr != nil {
		c.Auditor.Printf("GETPIN: err result - %s", cmdErr.Message)
		return "", cmdErr
	} else {
		c.Auditor.Println("GETPIN: result ok")
		return item.Login.Password, nil
	}
}


// TODO
func (c *BitwardenClient) Confirm(settings pinentry.Settings) (bool, *common.Error) {
	c.Auditor.Println("CONFIRM got called")
	return true, nil
}

// TODO
func (c *BitwardenClient) Message(settings pinentry.Settings) *common.Error {
	c.Auditor.Println("MESSAGE got called")
	return nil
}
