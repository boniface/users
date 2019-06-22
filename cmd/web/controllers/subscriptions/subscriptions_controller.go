package subscriptions

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"strconv"
	"time"
	"users/cmd/config"
	"users/pkg/api/subscription"
)

func Subscriptions(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", subscriptionsHanler(app))
	r.Get("/", subscriptionsHanler(app))
	// SubscritionType
	r.Post("/types/create", createSubscriptionTypeHandler(app))
	r.Post("/create", createSubscriptionHandler(app))
	return r

}

func createSubscriptionTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//id := xid.New().String()
		id := r.PostFormValue("id")
		subscritionTypeName := r.PostFormValue("subscritionTypeName")
		subscriptionType := subscription.SubscriptionTypes{id, subscritionTypeName}
		_, err := subscription.CreateSubscriptionType(subscriptionType)
		if err != nil {
			app.ServerError(w, err)
		}
		http.Redirect(w, r, "/subscriptions", 301)
	}

}
func createSubscriptionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		id := r.PostFormValue("id")
		subscriptionType := r.PostFormValue("subscriptionType")

		value, _ := strconv.ParseFloat(r.PostFormValue("SubscriptionValue"), 10)
		duration, _ := strconv.Atoi(r.PostFormValue("duration"))
		description := r.PostFormValue("description")
		dateCreated := time.Now()
		status := r.PostFormValue("status")

		fmt.Println("The Time is ", r.PostFormValue("dateCreated"))
		fmt.Println(" the VAlue for Site ", r.PostFormValue("SubscriptionValue"))

		sub := subscription.Subscription{id, subscriptionType, value, duration, description, dateCreated, status}
		_, err := subscription.CreateSubscription(sub)
		if err != nil {
			app.ServerError(w, err)
		}
		fmt.Println(" the Data in IS ", sub)
		http.Redirect(w, r, "/subscriptions", 301)
	}
}

func subscriptionsHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		subs, err := subscription.GetSubscriptions()
		if err != nil {
			app.ServerError(w, err)
		}

		usesubs, err := subscription.GetUserSubscriptions()

		if err != nil {

			app.ServerError(w, err)
		}

		siteSubs, err := subscription.GetSiteSubscriptions()

		if err != nil {
			app.ServerError(w, err)
		}

		subTypes, err := subscription.GetSubscriptionTypes()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			Subs     []subscription.Subscription
			SubTypes []subscription.SubscriptionTypes
			SiteSubs []subscription.SiteSubscription
			UserSubs []subscription.UserSubscriptions
		}
		data := PageData{subs, subTypes, siteSubs, usesubs}

		files := []string{
			app.Path + "/subscriptions/subscriptions.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebar.page.html",
			app.Path + "/base/footer.page.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
