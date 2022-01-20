package main

type NotEnoughFoodError struct {}

func (e NotEnoughFoodError) Error() string {
    return "Not enough food"
}

type NotEnoughEggsError struct {}

func (e NotEnoughEggsError) Error() string {
    return "Not enough eggs"
}
