package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	dto "server/dto/result"
	usersdto "server/dto/user"
	"server/models"
	"server/repositories"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gorilla/mux"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) FindUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.FindUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: users}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: user}
	json.NewEncoder(w).Encode(response)
}

// func (h *handlerUser) UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

//Mengambil struct dari dto

// id, _ := strconv.Atoi(mux.Vars(r)["id"])
// user, err := h.UserRepository.GetUser(int(id))
// if err != nil {
// 	w.WriteHeader(http.StatusBadRequest)
// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 	json.NewEncoder(w).Encode(response)
// 	return
// }

// dataContex := r.Context().Value("dataFile")
// filepath := dataContex.(string)

// var ctx = context.Background()
// var CLOUD_NAME = "dfxarsquq"
// var API_KEY = "424662388976554"
// var API_SECRET = "izwGO6NvRBu5pNVJoPyp2j1oNC4"

// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
// resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "HalloCorona"})
// fmt.Println(resp.SecureURL)
// if err != nil {
// 	fmt.Println("upload gagal", err.Error())
// }
// request := usersdto.UpdateUserRequest{
// 	Image: resp.SecureURL,
// }

// user = models.User{
// 	Image: request.Image,
// }

// password, err := bcrypt.HashingPassword(request.Password)
// if err != nil {
// 	w.WriteHeader(http.StatusInternalServerError)
// 	response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 	json.NewEncoder(w).Encode(response)
// 	return
// }

//konversi data dari string ke int

//Jika user requestnya tidak kosong maka isi ke models user
// if request.FullName != "" {
// 	user.Fullname = request.FullName
// }

// if request.Email != "" {
// 	user.Email = request.Email
// }

// if request.Password != "" {
// 	user.Password = password
// }

// if request.Phone != "" {
// 	user.Phone = request.Phone
// }

// if request.Address != "" {
// 	user.Address = request.Address
// }

// if request.Role != "" {
// 	user.Role = request.Role
// }

// 	data, err := h.UserRepository.UpdateUser(user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
// 	json.NewEncoder(w).Encode(response)
// }

func (h *handlerUser) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Mengambil struct dari dto
	request := new(usersdto.UpdateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		//Encoder seperti kode morse di pramuka
		json.NewEncoder(w).Encode(response)
		return
	}

	//konversi data dari string ke int
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.UserRepository.GetUser(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)

	var ctx = context.Background()
	var CLOUD_NAME = "dfxarsquq"
	var API_KEY = "424662388976554"
	var API_SECRET = "izwGO6NvRBu5pNVJoPyp2j1oNC4"

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "HalloCorona"})
	fmt.Println(resp.SecureURL)
	if err != nil {
		fmt.Println("upload gagal", err.Error())
	}

	if request.Image != "" {
		user.Image = resp.SecureURL
	}

	data, err := h.UserRepository.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.UserRepository.DeleteUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseDel(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseDel(u models.User) usersdto.UserResponseDel {
	return usersdto.UserResponseDel{
		ID: u.ID,
	}
}
