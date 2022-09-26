package service

import (
	"encoding/json"
	"fmt"
	"os"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type ChecklistPreventiveServiceInterface interface {
	CreateChecklistPreventive(request *model.CreateChecklistPreventiveRequest, context *gin.Context) (model.CreateChecklistPreventiveResponse, error)
	GetChecklistPreventive(request *string) (model.GetChecklistPreventiveResponse, error)
}

type checklistPreventiveService struct {
	checklistHwRepository             repository.ChecklistHwRepositoryInterface
	headerPreventiveRepository        repository.HeaderChecklistPreventiveRepositoryInterface
	userChecklistPreventiveRepository repository.UserChecklistPreventiveRepositoryInterface
}

func ChecklistPreventiveService(checklistHwRepository repository.ChecklistHwRepositoryInterface, headerPreventiveRepository repository.HeaderChecklistPreventiveRepositoryInterface, userChecklistPreventiveRepository repository.UserChecklistPreventiveRepositoryInterface) *checklistPreventiveService {
	return &checklistPreventiveService{checklistHwRepository, headerPreventiveRepository, userChecklistPreventiveRepository}
}

func (checklistPreventiveService *checklistPreventiveService) CreateChecklistPreventive(request *model.CreateChecklistPreventiveRequest, context *gin.Context) (model.CreateChecklistPreventiveResponse, error) {
	var response model.CreateChecklistPreventiveResponse
	var header_req *entity.HeaderChecklistPreventive
	var items_req []*entity.ChecklistHw
	var user_trilogi_req *entity.UserChecklistPreventive
	var user_tsel_req *entity.UserChecklistPreventive
	var user_req []*entity.UserChecklistPreventive

	json.Unmarshal([]byte(request.Header), &header_req)
	json.Unmarshal([]byte(request.Items), &items_req)
	json.Unmarshal([]byte(request.UserTrilogi), &user_trilogi_req)
	json.Unmarshal([]byte(request.UserTsel), &user_tsel_req)

	date_now := time.Now()
	dir := os.Getenv("FILE_DIR")
	path := dir + "signature/" + header_req.PrevCode + "/" + date_now.Format("2006-01-02")
	error := fmt.Errorf("error")
	user_trilogi_signature := ""
	user_tsel_signature := ""

	_, check_dir_error := os.Stat(path)

	if os.IsNotExist(check_dir_error) {
		check_dir_error := os.MkdirAll(path, 0755)

		if check_dir_error != nil {
			error = check_dir_error
		}
	}

	if request.UserTrilogiSignature != nil {
		user_trilogi_signature = request.UserTrilogiSignature.Filename
		error = context.SaveUploadedFile(request.UserTrilogiSignature, path+"/"+user_trilogi_signature)
	} else {
		error = fmt.Errorf("User trilogi signature must be filled!")
	}

	if request.UserTselSignature != nil {
		user_tsel_signature = request.UserTselSignature.Filename
		error = context.SaveUploadedFile(request.UserTselSignature, path+"/"+user_tsel_signature)
	} else {
		error = fmt.Errorf("User tsel signature must be filled!")
	}

	if error == nil {

		header_req.CreatedAt = date_now
		_, error = checklistPreventiveService.headerPreventiveRepository.CreateHeaderChecklistPreventive(header_req)

	}

	if error == nil {

		_, error = checklistPreventiveService.checklistHwRepository.CreateChecklistHw(items_req)

	}

	if error == nil {

		user_trilogi_req.Signature = user_trilogi_signature
		user_tsel_req.Signature = user_tsel_signature

		user_req = append(user_req, user_trilogi_req)
		user_req = append(user_req, user_tsel_req)

		_, error = checklistPreventiveService.userChecklistPreventiveRepository.CreateUserChecklistPreventive(user_req)

	}

	if error == nil {
		response.Header = header_req
		response.Items = items_req

		url := os.Getenv("FILE_URL")

		for index := range user_req {
			date := user_req[index].CreatedAt.Format("2006-01-02")
			path := url + "signature/" + header_req.PrevCode + "/" + date + "/"
			if user_req[index].Signature != "" {
				file_name := user_req[index].Signature
				user_req[index].Signature = path + file_name
			}
		}

		response.User = user_req
	}

	return response, error

}

func (checklistPreventiveService *checklistPreventiveService) GetChecklistPreventive(request *string) (model.GetChecklistPreventiveResponse, error) {
	var response model.GetChecklistPreventiveResponse
	var header_res entity.HeaderChecklistPreventive
	var items_res []entity.ChecklistHw
	var user_res []entity.UserChecklistPreventive
	error := fmt.Errorf("error")

	header_res, error = checklistPreventiveService.headerPreventiveRepository.GetHeaderChecklistPreventive(request)

	if error == nil {
		items_res, error = checklistPreventiveService.checklistHwRepository.GetChecklistHw(request)
	}

	if error == nil {
		user_res, error = checklistPreventiveService.userChecklistPreventiveRepository.GetUserChecklistPreventive(request)
		url := os.Getenv("FILE_URL")
		for index := range user_res {
			date := user_res[index].CreatedAt.Format("2006-01-02")
			path := url + "signature/" + *request + "/" + date + "/"
			if user_res[index].Signature != "" {
				file_name := user_res[index].Signature
				user_res[index].Signature = path + file_name
			}
		}
	}

	if error == nil {
		response.Header = header_res
		response.Items = items_res
		response.User = user_res
	}

	return response, error
}
