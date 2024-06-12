package controller

import (
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/helper"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/helper/ghupload"
	"github.com/whatsauth/itmodel"
)

func PostUploadGithub(respw http.ResponseWriter, req *http.Request) {
	var respn itmodel.Response
	// _, err := watoken.Decode(config.PublicKeyWhatsAuth, helper.GetLoginFromHeader(req))
	// if err != nil {
	// 	respn.Info = helper.GetSecretFromHeader(req)
	// 	respn.Response = err.Error()
	// 	helper.WriteJSON(respw, http.StatusForbidden, respn)
	// 	return
	// }
	// Parse the form file
	_, header, err := req.FormFile("image")
	if err != nil {
		
		respn.Response = err.Error()
		helper.WriteJSON(respw, http.StatusBadRequest, respn)
		return
	}

	//folder := ctx.Params("folder")
	folder := helper.GetParam(req)
	var pathFile string
	if folder != "" {
		pathFile = folder + "/" + header.Filename
	} else {
		pathFile = header.Filename
	}

	// save to github
	gh, err:=atdb.GetOneDoc[model.Ghcreates](config.Mongoconn, "github", bson.M{})
	if err != nil{
	respn.Info = helper.GetSecretFromHeader(req)
		respn.Response = err.Error()
		helper.WriteJSON(respw, http.StatusConflict, respn)
		return
	}

	content, _, err := ghupload.GithubUpload(gh.GitHubAccessToken, gh.GitHubAuthorName, gh.GitHubAuthorEmail, header, "alittifaq", "cdn", pathFile, false)
	if err != nil {
		respn.Info = "gagal upload github"
		respn.Response = err.Error()
		helper.WriteJSON(respw, http.StatusEarlyHints, content)
		return
	}

	helper.WriteJSON(respw, http.StatusOK, respn)

}
