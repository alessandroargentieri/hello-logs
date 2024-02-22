package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"net"
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
		h := fnv.New32a()
		h.Write([]byte(getIP()))
		ipHash := h.Sum32()
		r = rand.New(rand.NewSource(int64(ipHash) + time.Now().Unix()))
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

func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "unknown"
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return "unknown"

}

var messages = []string{
	"[warn] The only way to do great work is to love what you do.",
	"[debug] Believe you can and you're halfway there.",
	"[trace] It does not matter how slowly you go as long as you do not stop.",
	"[info] Our greatest glory is not in never falling, but in rising every time we fall.",
	"[info] You miss 100% of the shots you don't take.",
	"[error] I can accept failure, everyone fails at something. But I can't accept not trying.",
	"[info] Life is what happens when you're busy making other plans.",
	"[info] The future belongs to those who believe in the beauty of their dreams.",
	"[warn] The best time to plant a tree was 20 years ago. The second best time is now.",
	"[debug] It's not whether you get knocked down, it's whether you get up.",
	"[info] The only limit to our realization of tomorrow will be our doubts of today.",
	"[info] Everything you’ve ever wanted is on the other side of fear.",
	"[info] Success is not final, failure is not fatal: It is the courage to continue that counts.",
	"[info] Hardships often prepare ordinary people for an extraordinary destiny.",
	"[info] Your time is limited, don't waste it living someone else's life.",
	"[warn] Don’t judge each day by the harvest you reap but by the seeds that you plant.",
	"[warn] The best revenge is massive success.",
	"[info] I am not a product of my circumstances. I am a product of my decisions.",
	"[debug] The mind is everything. What you think you become.",
	"[error] An unexamined life is not worth living.",
	"[info] Eighty percent of success is showing up.",
	"[debug] Your life only gets better when you get better.",
	"[info] The way to get started is to quit talking and begin doing.",
	"[info] Life is 10% what happens to us and 90% how we react to it.",
	"[info] Do what you can with all you have, wherever you are.",
	"[info] You are never too old to set another goal or to dream a new dream.",
	"[warn] To see what is right and not do it is a lack of courage.",
	"[trace] I think goals should never be easy, they should force you to work, even if they are uncomfortable at the time.",
	"[info] One of the lessons that I grew up with was to always stay true to yourself and never let what somebody else says distract you from your goals.",
	"[info] Today’s accomplishments were yesterday’s impossibilities.",
	"[debug] The only person you are destined to become is the person you decide to be.",
	"[info] Go confidently in the direction of your dreams! Live the life you've imagined.",
	"[info] You don’t have to be great to start, but you have to start to be great.",
	"[warn] I find that the harder I work, the more luck I seem to have.",
	"[debug] The pessimist sees difficulty in every opportunity. The optimist sees opportunity in every difficulty.",
	"[info] Don’t let yesterday take up too much of today.",
	"[info] You learn more from failure than from success. Don’t let it stop you. Failure builds character.",
	"[debug] It’s not about how hard you can hit; it’s about how hard you can get hit and keep moving forward.",
	"[warn] We may encounter many defeats but we must not be defeated.",
	"[warn] Whether you think you can or think you can’t, you’re right.",
	"[info] The man who has confidence in himself gains the confidence of others.",
	"[error] Creativity is intelligence having fun.",
	"[info] What you get by achieving your goals is not as important as what you become by achieving your goals.",
	"[warn] To handle yourself, use your head; to handle others, use your heart.",
	"[debug] Too many of us are not living our dreams because we are living our fears.",
	"[info] I have learned over the years that when one's mind is made up, this diminishes fear.",
	"[info] Setting goals is the first step in turning the invisible into the visible.",
	"[warn] Quality is not an act, it is a habit.",
	"[debug] The only place where success comes before work is in the dictionary.",
	"[info] Life shrinks or expands in proportion to one's courage.",
}
