package handlers

import (
	"GO-08/database"
	"GO-08/models"
	"GO-08/providers"
	"GO-08/utils"
	"fmt"
	"net/http"
)

var EmailProvider providers.EmailProvider

func ContactUser(w http.ResponseWriter, r *http.Request) {
	var request models.User
	if err := utils.ParseBody(r.Body, &request); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "Failed to parse request")
		return
	}

	template, _ := EmailProvider.GetEmailTemplate(providers.EmailTypeContactUser)

	template.AddRecipient(request.FirstName, request.Email)

	template.DynamicData["fullName"] = fmt.Sprintf("%s %s", request.FirstName, request.LastName)
	template.DynamicData["phoneNumber"] = fmt.Sprintf("%s %s", request.CountryCode, request.Number)
	template.DynamicData["email"] = request.Email
	EmailProvider.Send(template)

	result, err := database.UserCollection.InsertOne(database.MongoCtx, request)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "Failed to save user")
		return
	}

	utils.RespondJSON(w, http.StatusOK, result)
}
