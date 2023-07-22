package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/repeale/fp-go"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			http.Error(res, "Bad Request", http.StatusBadRequest)
			return
		}

		date := req.URL.Query().Get("date")
		if date == "" {
			http.Error(res, "Bad Request", http.StatusBadRequest)
			return
		}

		hour, err := getHourOfDay(date)
		if err != nil {
			http.Error(res, "Bad Request", http.StatusBadRequest)
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Header().Set("Content-Type", "text/plain")
		x := fp.Pipe2(DeterminePartOfDayFromHour, MapPartOfDayToMotivationalMessage)(hour)
		res.Write([]byte(x))

		return
	})

	http.ListenAndServe(":8080", r)
}

func getHourOfDay(date string) (int, error) {
	isoLayout := "2006-01-02T15:04:05Z"
	time, err := time.Parse(isoLayout, date)
	time.Local().Hour()
	if err != nil {
		return 0, err
	} else {
		return time.Local().Hour(), nil
	}
}

type PartOfDay int

const (
	Morning PartOfDay = iota
	Afternoon
	Evening
	Night
)

func DeterminePartOfDayFromHour(hour int) PartOfDay {
	if hour >= 5 && hour < 12 {
		return Morning
	} else if hour >= 12 && hour < 16 {
		return Afternoon
	} else if hour >= 16 && hour < 21 {
		return Evening
	} else {
		return Night
	}
}

func MapPartOfDayToMotivationalMessage(partOfDay PartOfDay) string {
	switch partOfDay {
	case Morning:
		return generateMorningMessage()
	case Afternoon:
		return generateAfternoonMessage()
	case Evening:
		return generateEveningMessage()
	case Night:
		return generateNightMessage()
	default:
		return ""
	}
}

func generateMorningMessage() string {
	return pickOneFromList([]string{
		"Seize the day, your masterpiece awaits.",
		"Create your best today!",
		"The world awaits your brilliance.",
		"Today's pages are blank. Fill them well.",
		"Make today extraordinary!",
		"Dawn's here - so is your potential.",
		"Productivity is one sunrise away.",
		"Life's a canvas, paint your day.",
		"Be the energy you want to attract.",
		"Embrace the hustle, savor the outcome.",
	})
}

func generateAfternoonMessage() string {
	return pickOneFromList([]string{
		"Slow and steady wins the day. Keep going!",
		"Step by step, you're moving mountains.",
		"Remember: The sun also rises slowly, yet it brightens the world.",
		"Small progress is still progress. You're doing great!",
		"Remember the tortoise - slow, steady, unstoppable.",
		"Relax. Breathe. Your pace, your race.",
		"Every effort is a building block. You're creating greatness.",
		"It's okay to be tired. It's not okay to give up. You've got this!",
		"Tired? That's just proof of your effort. Keep pushing.",
		"Slow doesn't mean stopped. Every step matters.",
	})
}

func generateEveningMessage() string {
	return pickOneFromList([]string{
		"As the day ends, remember: every effort adds up. Be proud.",
		"Even the sun sets in paradise. Relax, you've done well today.",
		"Your dedication today deserves an evening of peace. Enjoy.",
		"Take pride in how far you've come. Have a peaceful night.",
		"Another day conquered. Rest and recharge, warrior.",
		"Feel the day's weight lifting. You've earned your calm.",
		"You gave today your all. Now, give yourself permission to relax.",
		"Productivity is knowing when to relax. You've done enough today.",
		"The stars are out, time to rest. Tomorrow, shine brighter.",
		"Tonight, let your mind rest. Tomorrow, let it create anew.",
	})
}

func generateNightMessage() string {
	return pickOneFromList([]string{
		"Embrace the night's silence. Let go of your worries, sleep deep.",
		"Rest your mind. Tomorrow holds new promise.",
		"Let the moon guide you to peaceful dreams.",
		"Stars can't shine without darkness. Embrace rest, recharge for tomorrow.",
		"Sleep is the best meditation. Embrace tranquility.",
		"Nighttime whispers peace. Listen, rest, rejuvenate.",
		"Like the calm sea under the moon, let your mind find peace tonight.",
		"Breathe out worries, breathe in peace. Sleep awaits.",
		"Tomorrow is a new canvas. Rest now, paint tomorrow.",
		"Trade your fears for dreams. Goodnight, peaceful sleep.",
	})
}

func pickOneFromList(list []string) string {
	i := rand.Intn(len(list))
	return list[i]
}
