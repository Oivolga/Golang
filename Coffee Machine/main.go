package main

import (
	"fmt"	
)

func main() {
    water := 400
    milk := 540
    beans := 120
    cups := 9
    money := 550
    var action string
    for action != "exit" {
        fmt.Println("\nWrite action (buy, fill, take, remaining, exit):")
        fmt.Scan(&action)
        switch(action) {
            case "exit": break
            case "remaining": display(water, milk, beans, cups, money)
            case "fill": {
                var w, m, b, c int
                fmt.Println("Write how many ml of water you want to add:")
                fmt.Scan(&w)
                water += w
                fmt.Println("Write how many ml of milk you want to add:")
                fmt.Scan(&m)
                milk += m
                fmt.Println("Write how many grams of coffee beans you want to add:")
                fmt.Scan(&b)
                beans += b
                fmt.Println("Write how many disposable coffee cups you want to add:")
                fmt.Scan(&c)
                cups += c
            }
            case "take": {
                fmt.Print("I gave you $")
                fmt.Println(money)
                money = 0
            }
            case "buy": {
                var num string
                fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
                fmt.Scan(&num)
                if num != "back" {
                    var w1, m1, b1 int
                    switch(num) {
                        case "1": {
                            w1 = 250
                            m1 = 0
                            b1 = 16
                        }
                        case "2": {
                            w1 = 350
                            m1 = 75
                            b1 = 20
                        }
                        case "3": {
                            w1 = 200
                            m1 = 100
                            b1 = 12
                        }
                    }
                    if water < w1 {
                        fmt.Println("Sorry, not enough water!")
                    } else if milk < m1 {
                        fmt.Println("Sorry, not enough milk!")
                    } else if beans < b1 {
                        fmt.Println("Sorry, not enough coffee beans!")
                    } else if cups < 1 {
                        fmt.Println("Sorry, not enough cups!")
                    } else {
                        fmt.Println("I have enough resources, making you a coffee!")
                        switch(num) {
                            case "1": {
                                water -= 250
                                beans -= 16
                                money += 4  
                            }
                            case "2": {
                                water -= 350
                                milk -= 75
                                beans -= 20
                                money += 7 
                            }
                            case "3": {
                                water -= 200
                                milk -= 100
                                beans -= 12
                                money += 6
                            }
                        }
                        cups -= 1
                    }
                }
            }
        }
    }
}

func display(water int, milk int, beans int, cups int, money int) {
    fmt.Println("The coffee machine has:")
    fmt.Println(water, "of water")
    fmt.Println(milk, "of milk")
    fmt.Println(beans, "of coffee beans")
    fmt.Println(cups, "of disposable cups")
    fmt.Println(money, "of money")
}
