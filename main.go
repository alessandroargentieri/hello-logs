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
"The only way to do great work is to love what you do.",
"Believe you can and you're halfway there.",
"It does not matter how slowly you go as long as you do not stop.",
"Our greatest glory is not in never falling, but in rising every time we fall.",
"You miss 100% of the shots you don't take.",
"I can accept failure, everyone fails at something. But I can't accept not trying.",
"Life is what happens when you're busy making other plans.",
"The future belongs to those who believe in the beauty of their dreams.",
"The best time to plant a tree was 20 years ago. The second best time is now.",
"It's not whether you get knocked down, it's whether you get up.",
"The only limit to our realization of tomorrow will be our doubts of today.",
"Everything you’ve ever wanted is on the other side of fear.",
"Success is not final, failure is not fatal: It is the courage to continue that counts.",
"Hardships often prepare ordinary people for an extraordinary destiny.",
"Your time is limited, don't waste it living someone else's life.",
"Don’t judge each day by the harvest you reap but by the seeds that you plant.",
"The best revenge is massive success.",
"I am not a product of my circumstances. I am a product of my decisions.",
"The mind is everything. What you think you become.",
"An unexamined life is not worth living.",
"Eighty percent of success is showing up.",
"Your life only gets better when you get better.",
"The way to get started is to quit talking and begin doing.",
"Life is 10% what happens to us and 90% how we react to it.",
"Do what you can with all you have, wherever you are.",
"You are never too old to set another goal or to dream a new dream.",
"To see what is right and not do it is a lack of courage.",
"I think goals should never be easy, they should force you to work, even if they are uncomfortable at the time.",
"One of the lessons that I grew up with was to always stay true to yourself and never let what somebody else says distract you from your goals.",
"Today’s accomplishments were yesterday’s impossibilities.",
"The only person you are destined to become is the person you decide to be.",
"Go confidently in the direction of your dreams! Live the life you've imagined.",
"You don’t have to be great to start, but you have to start to be great.",
"I find that the harder I work, the more luck I seem to have.",
"The pessimist sees difficulty in every opportunity. The optimist sees opportunity in every difficulty.",
"Don’t let yesterday take up too much of today.",
"You learn more from failure than from success. Don’t let it stop you. Failure builds character.",
"It’s not about how hard you can hit; it’s about how hard you can get hit and keep moving forward.",
"We may encounter many defeats but we must not be defeated.",
"Whether you think you can or think you can’t, you’re right.",
"The man who has confidence in himself gains the confidence of others.",
"Creativity is intelligence having fun.",
"What you get by achieving your goals is not as important as what you become by achieving your goals.",
"To handle yourself, use your head; to handle others, use your heart.",
"Too many of us are not living our dreams because we are living our fears.",
"I have learned over the years that when one's mind is made up, this diminishes fear.",
"Setting goals is the first step in turning the invisible into the visible.",
"Quality is not an act, it is a habit.",
"The only place where success comes before work is in the dictionary.",
"Life shrinks or expands in proportion to one's courage.",

}
