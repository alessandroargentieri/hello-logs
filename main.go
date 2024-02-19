package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	random bool   = os.Getenv("RANDOM") == "true"
	msg    string = os.Getenv("MSG")
	r      *rand.Rand
)

func main() {

	if msg == "" && !random {
		msg = "Hello Logs!"
	}

	if random {
		r = rand.New(rand.NewSource(time.Now().Unix()))
	}

	for {
		time.Sleep(3 * time.Second)
		fmt.Println(getMsg(random))
	}
}

func getMsg(random bool) string {
	if random {
		return messages[r.Intn(len(messages))]
	}
	return msg
}

var messages = []string{
	"The sky is blue.",
	"Dogs bark loudly.",
	"I love eating pizza.",
	"She dances gracefully.",
	"The sun sets in the west.",
	"Birds chirp in the morning.",
	"He plays the guitar beautifully.",
	"Flowers bloom in spring.",
	"The coffee tastes bitter.",
	"She reads novels every night.",
	"The cat purrs softly.",
	"He runs fast.",
	"The raindrops fall gently.",
	"They laugh together.",
	"I enjoy hiking in the mountains.",
	"The children play in the park.",
	"She sings in the choir.",
	"The river flows peacefully.",
	"He paints landscapes.",
	"The wind blows gently.",
	"The stars twinkle in the night sky.",
	"I drink tea with honey.",
	"She writes poetry.",
	"The waves crash against the shore.",
	"They study together.",
	"The fire crackles and pops.",
	"He cooks dinner for his family.",
	"The trees sway in the breeze.",
	"She swims in the ocean.",
	"The clock ticks quietly.",
	"I watch movies on weekends.",
	"He fixes cars for a living.",
	"The snow falls softly.",
	"They build sandcastles on the beach.",
	"She rides her bike to work.",
	"The moon shines brightly.",
	"He plays soccer with his friends.",
	"The butterfly flutters by.",
	"I bake cookies for dessert.",
	"She takes photographs of nature.",
	"The train chugs along the tracks.",
	"They go camping in the woods.",
	"He plays video games all day.",
	"The fog rolls in from the sea.",
	"She practices yoga in the morning.",
	"The birds migrate south for the winter.",
	"I walk my dog in the park.",
	"He writes code for a living.",
	"The snowflakes are delicate.",
	"She paints her nails pink.",
	"The baby giggles and coos.",
	"They go fishing at the lake.",
	"He reads the newspaper every morning.",
	"The flowers bloom in the garden.",
	"She dances in the rain.",
	"The wind howls outside.",
	"I listen to music on my headphones.",
	"He takes a nap in the hammock.",
	"The leaves rustle in the breeze.",
	"She studies for her exams.",
	"The cricket chirps in the night.",
	"They roast marshmallows over the campfire.",
	"He mows the lawn every Saturday.",
	"The thunder rumbles in the distance.",
	"She practices her violin.",
	"The snow covers the ground.",
	"I plant flowers in the garden.",
	"He takes photographs of the sunset.",
	"The cat curls up by the fireplace.",
	"She bakes bread from scratch.",
	"The waves crash against the rocks.",
	"They go for a hike in the forest.",
	"He fixes the leaky faucet.",
	"The clouds drift across the sky.",
	"She sketches in her notebook.",
	"The bees buzz around the flowers.",
	"I take a bubble bath to relax.",
	"He climbs to the top of the mountain.",
	"The fireflies light up the night.",
	"She practices her dance routine.",
	"The sun rises in the east.",
	"They have a picnic in the park.",
	"He builds a sandcastle on the beach.",
	"The rainbows appear after the rain.",
	"She waters the plants in the garden.",
	"The stars twinkle in the night sky.",
	"I cook dinner for my family.",
	"He sketches the landscape.",
	"The waves crash against the shore.",
	"She rides her bike through the countryside.",
	"The birds sing in the trees.",
	"They fly kites in the park.",
	"He studies for his exams.",
	"The flowers bloom in the spring.",
	"She dances in the moonlight.",
	"The wind blows through the trees.",
	"I watch the sunset from my balcony.",
	"He reads a book by the fireplace.",
	"The snowflakes fall gently to the ground.",
	"She takes a walk in the park.",
}
