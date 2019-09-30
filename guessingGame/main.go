package main

import(
	"fmt"
	"errors"
	"log"
)

func getHint(secret string, guess string) (string , error){
	if secret == ""{
		return "",errors.New("Entered secret is nil")
	}else if guess == ""{
		return "",errors.New("Entered guess is nil")
	}
	m := make(map[byte]bool)
	for i:=0;i<len(secret);i++{
		if m[secret[i]] == true{
			continue
		}else{
			m[secret[i]] = true
		}
	} // this passes through a map 

	if len(secret) != len(guess){
		count := 0
		for i:=0;i<len(guess);i++{
			if m[guess[i]]{
				count++
			}
		}
		fmt.Println("Count is :", count)
		return "you entered a string that wasnt the right size, but i can tell you that you have: "+string(count) + " in common",nil
	}
	var output string
	for i:=0;i<len(guess);i++{
		if guess[i] == secret[i]{
			output+=string(guess[i])
		}else if m[guess[i]] == true{
			output +="A"
		}else{
			output += "B"
		}
	}
	if output == secret{
		return "You guess correctly",nil
	}

	return output, nil
}

func main(){
	secret := "helloworld"
	fmt.Print("enter a guess:")
	var input string
	fmt.Scan(&input)
	guessed := false

	for guessed == false{
		getHint , err := getHint(secret, input)

		if err != nil{
			log.Println(err)
		}else if getHint == "You guess correctly"{
			fmt.Println(getHint)
			return
		}else{
			fmt.Println(getHint)
			fmt.Print("Enter in your guess")
			fmt.Scan(&input)
		}
	}
	return
}

// take a random number that you dont know, then guess it and continue to return hi or low
