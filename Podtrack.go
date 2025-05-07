//This is the Golang code for Podtrack - a system for podcast publishers and subscribers alike! 
//As Publisher, you can manage your podcasts and track performance
//The system recommends the right podcast for your (customer's) mood
package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// Structs for Publisher (CRUD) & Productivity Tracking
type Podcast struct {
	ID     int
	Title  string
	Mood   string
	Tags   []string
	Seats  int
}

type Task struct {
	Name       string
	AssignedTo string
	Completed  bool
}

type Publisher struct {
	Podcasts []Podcast
	Tasks    []Task
}

// Structs for Customer Side (Recommendation)
type Mood struct {
	Name        string
	Recommended []string // List of recommended podcasts based on this mood
}

// Global data (for simplicity)
var publisher = Publisher{}
var customerMoods = []Mood{
	{"Tired", []string{"Calm Talks", "Relaxing Beats", "Evening Chill"}},
	{"Melancholic", []string{"Sad Songs", "Blue Skies", "Reflective Moments"}},
	{"Happy", []string{"Energetic Vibes", "Party Time", "Adventure Begins"}},
	{"Grounded", []string{"Spiritual Awakening", "Peaceful Meditation", "Deep Reflections"}},
}

// Function to display tables (nice text format)
func displayTable(headers []string, data [][]string) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', 0)
	fmt.Fprintln(w, strings.Join(headers, "\t"))
	for _, row := range data {
		fmt.Fprintln(w, strings.Join(row, "\t"))
	}
	w.Flush()
}

// Publisher Operations
func (p *Publisher) createPodcast(id int, title, mood string, tags []string, seats int) {
	p.Podcasts = append(p.Podcasts, Podcast{ID: id, Title: title, Mood: mood, Tags: tags, Seats: seats})
	fmt.Println("Podcast Created:", title)
}

func (p *Publisher) readPodcasts() {
	var data [][]string
	for _, pod := range p.Podcasts {
		data = append(data, []string{fmt.Sprintf("%d", pod.ID), pod.Title, pod.Mood, fmt.Sprintf("%d seats", pod.Seats)})
	}
	displayTable([]string{"ID", "Title", "Mood", "Seats Available"}, data)
}

func (p *Publisher) updatePodcast(id int, newTitle, newMood string) {
	for i := range p.Podcasts {
		if p.Podcasts[i].ID == id {
			p.Podcasts[i].Title = newTitle
			p.Podcasts[i].Mood = newMood
			fmt.Println("Podcast Updated:", newTitle)
			return
		}
	}
	fmt.Println("Podcast not found.")
}

func (p *Publisher) deletePodcast(id int) {
	for i := range p.Podcasts {
		if p.Podcasts[i].ID == id {
			p.Podcasts = append(p.Podcasts[:i], p.Podcasts[i+1:]...)
			fmt.Println("Podcast Deleted")
			return
		}
	}
	fmt.Println("Podcast not found.")
}

// Customer Operations (Mood-based Recommendations)
func recommendPodcastBasedOnMood(moodInput string) {
	for _, mood := range customerMoods {
		if strings.ToLower(mood.Name) == strings.ToLower(moodInput) {
			fmt.Println("Recommended Podcasts based on your mood (" + moodInput + "):")
			for _, rec := range mood.Recommended {
				fmt.Println("- " + rec)
			}
			return
		}
	}
	fmt.Println("Sorry, no recommendations available for that mood.")
}

// CLI Flow
func main() {
	// Initial Data
	publisher.createPodcast(1, "Calm Talks", "Tired", []string{"relaxing", "calm", "meditation"}, 0)
	publisher.createPodcast(2, "Energetic Vibes", "Happy", []string{"energetic", "party", "dance"}, 5)
	publisher.createPodcast(3, "Spiritual Awakening", "Grounded", []string{"spiritual", "meditation", "calm"}, 3)

	// CLI Loop
	for {
		fmt.Println("\n--------------------------")
		fmt.Println("Welcome to PodTrack System!")
		fmt.Println("--------------------------")
		fmt.Println("1. Publisher Mode")
		fmt.Println("2. Customer Mode")
		fmt.Println("3. Exit")
		fmt.Print("Select an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Publisher Mode
			var publisherChoice int
			fmt.Println("\nPublisher Mode - Manage Podcasts")
			fmt.Println("1. Create Podcast")
			fmt.Println("2. Read Podcasts")
			fmt.Println("3. Update Podcast")
			fmt.Println("4. Delete Podcast")
			fmt.Println("5. Exit Publisher Mode")
			fmt.Print("Select an option: ")
			fmt.Scanln(&publisherChoice)

			switch publisherChoice {
			case 1:
				var id, seats int
				var title, mood, tagsInput string
				fmt.Print("Enter podcast ID: ")
				fmt.Scanln(&id)
				fmt.Print("Enter podcast title: ")
				fmt.Scanln(&title)
				fmt.Print("Enter podcast mood: ")
				fmt.Scanln(&mood)
				fmt.Print("Enter tags (comma-separated): ")
				fmt.Scanln(&tagsInput)
				tags := strings.Split(tagsInput, ",")
				fmt.Print("Enter number of seats: ")
				fmt.Scanln(&seats)
				publisher.createPodcast(id, title, mood, tags, seats)
			case 2:
				publisher.readPodcasts()
			case 3:
				var id int
				var newTitle, newMood string
				fmt.Print("Enter podcast ID to update: ")
				fmt.Scanln(&id)
				fmt.Print("Enter new title: ")
				fmt.Scanln(&newTitle)
				fmt.Print("Enter new mood: ")
				fmt.Scanln(&newMood)
				publisher.updatePodcast(id, newTitle, newMood)
			case 4:
				var id int
				fmt.Print("Enter podcast ID to delete: ")
				fmt.Scanln(&id)
				publisher.deletePodcast(id)
			case 5:
				continue
			}

		case 2:
			// Customer Mode (Mood-based recommendations)
			var moodInput string
			fmt.Println("\nCustomer Mode - Recommend Podcast based on Mood")
			fmt.Println("How do you feel today?")
			fmt.Println("A. Tired")
			fmt.Println("B. Melancholic")
			fmt.Println("C. Happy")
			fmt.Println("D. Grounded")
			fmt.Print("Enter your mood (A/B/C/D): ")
			fmt.Scanln(&moodInput)

			switch moodInput {
			case "A":
				recommendPodcastBasedOnMood("Tired")
			case "B":
				recommendPodcastBasedOnMood("Melancholic")
			case "C":
				recommendPodcastBasedOnMood("Happy")
			case "D":
				recommendPodcastBasedOnMood("Grounded")
			default:
				fmt.Println("Invalid mood input!")
			}

		case 3:
			// Exit
			fmt.Println("Exiting PodTrack System...")
			return
		}
	}
}
