package danmu_rest

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

//Controller ...controller of MlController
type Controller struct {
	MlController MlController
}

func setupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-TOken, Authorization")
}

//MostSimilar get most similar words from word2vec
func (c *Controller) MostSimilar(w http.ResponseWriter, r *http.Request) {

	//setupResponse(&w)
//	w.Header().Add("Access-Control-Allow-Origin", "*")
//	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
//	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

	var similarWordRequest SimilarWordRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	log.Println(body)

	if err != nil {
		log.Fatalln("request body parse error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &similarWordRequest); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error occured when encoding MlRequest unmarshalling error message.", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	success, out, stderr := c.MlController.getMostSimilarWords(
		similarWordRequest.Word)

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if stderr != "" {
		errMessage := SimilarWordsResponse{"Ooops, something wrong happens when searching similar words."}
		js, err := json.Marshal(errMessage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.Write(js)
		return
	}

	similarWordsResponse := SimilarWordsResponse{out}
	js, err := json.Marshal(similarWordsResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(js)
}

//RecentTopic
func (c *Controller) RecentTopic(w http.ResponseWriter, r *http.Request) {

//	setupResponse(&w)
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
//	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var recentTopicRequest RecentTopicRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	log.Println(body)

	if err != nil {
		log.Fatalln("request body parse error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &recentTopicRequest); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error occured when encoding MlRequest unmarshalling error message.", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	success, out, stderr := c.MlController.getRecentTopic(
		recentTopicRequest.Interval)

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if stderr != "" {
		//errMessage := RecentTopicResponse{"Ooops, something wrong happens when searching recent topics."}
		errMessage := RecentTopicResponse{stderr}
		js, err := json.Marshal(errMessage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.Write(js)
		return
	}

	recentTopicResponse := RecentTopicResponse{out}
	js, err := json.Marshal(recentTopicResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Write(js)
}
